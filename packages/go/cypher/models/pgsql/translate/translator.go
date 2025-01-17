// Copyright 2024 Specter Ops, Inc.
//
// Licensed under the Apache License, Version 2.0
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: Apache-2.0

package translate

import (
	"context"
	"github.com/specterops/bloodhound/cypher/models"
	"github.com/specterops/bloodhound/cypher/models/cypher"
	"github.com/specterops/bloodhound/cypher/models/pgsql"
	"github.com/specterops/bloodhound/cypher/models/walk"
	"github.com/specterops/bloodhound/dawgs/graph"
)

type intermediates struct {
	properties  map[string]pgsql.Expression
	pattern     *Pattern
	match       *Match
	projections *Projections
	mutations   *Mutations
}

func (s *intermediates) PrepareProjections(distinct bool) {
	s.projections = &Projections{
		Distinct: distinct,
	}
}

func (s *intermediates) PrepareMutations() {
	if s.mutations == nil {
		s.mutations = NewMutations()
	}
}

func (s *intermediates) HasMutations() bool {
	return s.mutations != nil && s.mutations.Assignments.Len() > 0
}

func (s *intermediates) HasDeletions() bool {
	return s.mutations != nil && s.mutations.Deletions.Len() > 0
}

func (s *intermediates) PrepareProjection() {
	s.projections.Items = append(s.projections.Items, &Projection{})
}

func (s *intermediates) CurrentProjection() *Projection {
	return s.projections.Current()
}

func (s *intermediates) PrepareProperties() {
	if s.properties != nil {
		clear(s.properties)
	} else {
		s.properties = map[string]pgsql.Expression{}
	}
}

type Translator struct {
	walk.HierarchicalVisitor[cypher.SyntaxNode]

	ctx            context.Context
	kindMapper     pgsql.KindMapper
	translation    Result
	treeTranslator *ExpressionTreeTranslator
	intermediates  *intermediates
	query          *Query
}

func NewTranslator(ctx context.Context, kindMapper pgsql.KindMapper, parameters map[string]any) *Translator {
	if parameters == nil {
		parameters = map[string]any{}
	}

	return &Translator{
		HierarchicalVisitor: walk.NewComposableHierarchicalVisitor[cypher.SyntaxNode](),
		translation: Result{
			Parameters: parameters,
		},
		ctx:            ctx,
		kindMapper:     kindMapper,
		treeTranslator: NewExpressionTreeTranslator(),
		intermediates:  &intermediates{},
		query: &Query{
			Scope: NewScope(),
		},
	}
}

func (s *Translator) Enter(expression cypher.SyntaxNode) {
	switch typedExpression := expression.(type) {
	case *cypher.MultiPartQuery:

	case *cypher.RegularQuery, *cypher.SingleQuery, *cypher.PatternElement, *cypher.Return,
		*cypher.Comparison, *cypher.Skip, *cypher.Limit, cypher.Operator, *cypher.ArithmeticExpression,
		*cypher.NodePattern, *cypher.RelationshipPattern, *cypher.Remove, *cypher.Set,
		*cypher.ReadingClause, *cypher.UnaryAddOrSubtractExpression, *cypher.PropertyLookup,
		*cypher.Negation, *cypher.Create, *cypher.Where, *cypher.ListLiteral,
		*cypher.FunctionInvocation, *cypher.Order, *cypher.RemoveItem, *cypher.SetItem,
		*cypher.MapItem:
	// No operation for these syntax nodes

	case *cypher.SinglePartQuery:
		s.query.PrepareTail()

	case *cypher.Match:
		// Start with a fresh match and where clause. Instantiation of the where clause here is necessary since
		// cypher will store identifier constraints in the query pattern which precedes the query where clause.
		s.intermediates.pattern = &Pattern{}
		s.intermediates.match = &Match{
			Scope:   s.query.Scope,
			Pattern: s.intermediates.pattern,
		}

	case graph.Kinds:
		s.treeTranslator.Push(pgsql.KindListLiteral{
			Values: typedExpression,
		})

	case *cypher.KindMatcher:
		if err := s.translateKindMatcher(typedExpression); err != nil {
			s.SetError(err)
		}

	case *cypher.Parameter:
		var (
			cypherIdentifier = pgsql.Identifier(typedExpression.Symbol)
			binding, bound   = s.query.Scope.AliasedLookup(cypherIdentifier)
		)

		if !bound {
			if parameterBinding, err := s.query.Scope.DefineNew(pgsql.ParameterIdentifier); err != nil {
				s.SetError(err)
			} else {
				// Alias the old parameter identifier to the synthetic one
				if cypherIdentifier != "" {
					s.query.Scope.Alias(cypherIdentifier, parameterBinding)
				}

				// Create a new container for the parameter and its value
				if newParameter, err := pgsql.AsParameter(parameterBinding.Identifier, typedExpression.Value); err != nil {
					s.SetError(err)
				} else if negotiatedValue, err := pgsql.NegotiateValue(typedExpression.Value); err != nil {
					s.SetError(err)
				} else {
					// Lift the parameter value into the parameters map
					s.translation.Parameters[parameterBinding.Identifier.String()] = negotiatedValue
					parameterBinding.Parameter = models.ValueOptional(newParameter)
				}

				// Set the outer reference
				binding = parameterBinding
			}
		}

		s.treeTranslator.Push(binding.Parameter.Value)

	case *cypher.Variable:
		if binding, resolved := s.query.Scope.LookupString(typedExpression.Symbol); !resolved {
			s.SetErrorf("unable to find identifier %s", typedExpression.Symbol)
		} else {
			s.treeTranslator.Push(binding.Identifier)
		}

	case *cypher.Literal:
		literalValue := typedExpression.Value

		if stringValue, isString := typedExpression.Value.(string); isString {
			// Cypher parser wraps string literals with ' characters
			literalValue = stringValue[1 : len(stringValue)-1]
		}

		if newLiteral, err := pgsql.AsLiteral(literalValue); err != nil {
			s.SetError(err)
		} else {
			newLiteral.Null = typedExpression.Null
			s.treeTranslator.Push(newLiteral)
		}

	case *cypher.Parenthetical:
		s.treeTranslator.PushParenthetical()

	case *cypher.SortItem:
		s.query.Tail.OrderBy = append(s.query.Tail.OrderBy, pgsql.OrderBy{
			Ascending: typedExpression.Ascending,
		})

	case *cypher.Projection:
		if err := s.translateProjection(typedExpression); err != nil {
			s.SetError(err)
		}

	case *cypher.ProjectionItem:
		s.intermediates.PrepareProjection()

	case *cypher.PatternPredicate:
		s.intermediates.pattern = &Pattern{}

		if err := s.translatePatternPredicate(s.query.Scope); err != nil {
			s.SetError(err)
		}

	case *cypher.PatternPart:
		if err := s.translatePatternPart(s.query.Scope, typedExpression); err != nil {
			s.SetError(err)
		}

	case *cypher.PartialComparison:
		s.treeTranslator.PushOperator(pgsql.Operator(typedExpression.Operator))

	case *cypher.PartialArithmeticExpression:
		s.treeTranslator.PushOperator(pgsql.Operator(typedExpression.Operator))

	case *cypher.Disjunction:
		for idx := 0; idx < typedExpression.Len()-1; idx++ {
			s.treeTranslator.PushOperator(pgsql.OperatorOr)
		}

	case *cypher.Conjunction:
		for idx := 0; idx < typedExpression.Len()-1; idx++ {
			s.treeTranslator.PushOperator(pgsql.OperatorAnd)
		}

	case *cypher.UpdatingClause:
		s.intermediates.PrepareMutations()

	case *cypher.Properties:
		s.intermediates.PrepareProperties()

	case *cypher.Delete:
		s.intermediates.PrepareMutations()

	default:
		s.SetErrorf("unable to translate cypher type: %T", expression)
	}
}

func (s *Translator) Exit(expression cypher.SyntaxNode) {
	switch typedExpression := expression.(type) {
	case *cypher.NodePattern:
		if err := s.translateNodePattern(s.query.Scope, typedExpression, s.intermediates.pattern.CurrentPart()); err != nil {
			s.SetError(err)
		}

	case *cypher.RelationshipPattern:
		if err := s.translateRelationshipPattern(s.query.Scope, typedExpression, s.intermediates.pattern.CurrentPart()); err != nil {
			s.SetError(err)
		}

	case *cypher.MapItem:
		if value, err := s.treeTranslator.Pop(); err != nil {
			s.SetError(err)
		} else {
			s.intermediates.properties[typedExpression.Key] = value
		}

	case *cypher.PatternPredicate:
		// Retire the predicate scope frames and build the predicate
		for range s.intermediates.pattern.CurrentPart().TraversalSteps {
			s.query.Scope.PopFrame()
		}

		if err := s.buildPatternPredicate(); err != nil {
			s.SetError(err)
		}

	case *cypher.RemoveItem:
		if err := s.translateRemoveItem(typedExpression); err != nil {
			s.SetError(err)
		}

	case *cypher.Delete:
		if err := s.translateDelete(s.query.Scope, typedExpression); err != nil {
			s.SetError(err)
		}

	case *cypher.SetItem:
		if err := s.translateSetItem(typedExpression); err != nil {
			s.SetError(err)
		}

	case *cypher.ListLiteral:
		var (
			numExpressions = len(typedExpression.Expressions())
			literal        = pgsql.ArrayLiteral{
				Values:   make([]pgsql.Expression, numExpressions),
				CastType: pgsql.UnsetDataType,
			}
		)

		for idx := numExpressions - 1; idx >= 0; idx-- {
			if nextExpression, err := s.treeTranslator.Pop(); err != nil {
				s.SetError(err)
			} else {
				if typeHint, isTypeHinted := nextExpression.(pgsql.TypeHinted); isTypeHinted {
					if arrayCastType, err := typeHint.TypeHint().ToArrayType(); err != nil {
						s.SetError(err)
					} else if literal.CastType != pgsql.UnsetDataType && literal.CastType != arrayCastType {
						s.SetErrorf("expected array literal value type %s at index %d but found type %s", literal.CastType, idx, arrayCastType)
					} else {
						literal.CastType = arrayCastType
					}
				}

				literal.Values[idx] = nextExpression
			}
		}

		if literal.CastType == pgsql.UnsetDataType {
			s.SetErrorf("array literal has no available type hints")
		} else {
			s.treeTranslator.Push(literal)
		}

	case *cypher.SortItem:
		// Rewrite the order by constraints
		if lookupExpression, err := s.treeTranslator.Pop(); err != nil {
			s.SetError(err)
		} else if err := RewriteExpressionIdentifiers(lookupExpression, s.query.Scope.CurrentFrameBinding().Identifier, s.query.Scope.Visible()); err != nil {
			s.SetError(err)
		} else {
			if propertyLookup, isPropertyLookup := asPropertyLookup(lookupExpression); isPropertyLookup {
				// If sorting, use the raw type of the JSONB field
				propertyLookup.Operator = pgsql.OperatorJSONField
			}

			s.query.Tail.CurrentOrderBy().Expression = lookupExpression
		}

	case *cypher.KindMatcher:
		if matcher, err := s.treeTranslator.Pop(); err != nil {
			s.SetError(err)
		} else {
			s.treeTranslator.Push(matcher)
		}

	case *cypher.Parenthetical:
		// Pull the sub-expression we wrap
		if wrappedExpression, err := s.treeTranslator.Pop(); err != nil {
			s.SetError(err)
		} else if parenthetical, err := s.treeTranslator.PopParenthetical(); err != nil {
			s.SetError(err)
		} else {
			parenthetical.Expression = wrappedExpression
			s.treeTranslator.Push(*parenthetical)
		}

	case *cypher.FunctionInvocation:
		s.translateFunction(typedExpression)

	case *cypher.UnaryAddOrSubtractExpression:
		if operand, err := s.treeTranslator.Pop(); err != nil {
			s.SetError(err)
		} else {
			s.treeTranslator.Push(&pgsql.UnaryExpression{
				Operator: pgsql.Operator(typedExpression.Operator),
				Operand:  operand,
			})
		}

	case *cypher.Negation:
		if operand, err := s.treeTranslator.Pop(); err != nil {
			s.SetError(err)
		} else {
			for cursor := operand; cursor != nil; {
				switch typedCursor := cursor.(type) {
				case pgsql.Parenthetical:
					// Unwrap parentheticals
					cursor = typedCursor.Expression
					continue

				case *pgsql.BinaryExpression:
					switch typedCursor.Operator {
					case pgsql.OperatorLike, pgsql.OperatorILike:
						// If this is a string comparison operation then the negation requires wrapping the
						// operand references in coalesce functions. While this will kick out index acceleration
						// the negation will already damage the query planner's ability to utilize an index lookup.

						if leftPropertyLookup, isPropertyLookup := asPropertyLookup(typedCursor.LOperand); isPropertyLookup {
							typedCursor.LOperand = pgsql.FunctionCall{
								Function: pgsql.FunctionCoalesce,
								Parameters: []pgsql.Expression{
									leftPropertyLookup,
									pgsql.NewLiteral("", pgsql.Text),
								},
								CastType: pgsql.Text,
							}
						}

						if rightPropertyLookup, isPropertyLookup := asPropertyLookup(typedCursor.ROperand); isPropertyLookup {
							typedCursor.ROperand = pgsql.FunctionCall{
								Function: pgsql.FunctionCoalesce,
								Parameters: []pgsql.Expression{
									rightPropertyLookup,
									pgsql.NewLiteral("", pgsql.Text),
								},
								CastType: pgsql.Text,
							}
						}
					}
				}

				break
			}

			s.treeTranslator.Push(&pgsql.UnaryExpression{
				Operator: pgsql.OperatorNot,
				Operand:  operand,
			})
		}

	case *cypher.Where:
		// Assign the last operands as identifier set constraints
		if err := s.treeTranslator.PopRemainingExpressionsAsConstraints(); err != nil {
			s.SetError(err)
		}

	case *cypher.PropertyLookup:
		s.translatePropertyLookup(typedExpression)

	case *cypher.PartialComparison:
		if err := s.treeTranslator.PopPushOperator(s.query.Scope, pgsql.Operator(typedExpression.Operator)); err != nil {
			s.SetError(err)
		}


	case *cypher.PartialArithmeticExpression:
		if err := s.treeTranslator.PopPushOperator(s.query.Scope, pgsql.Operator(typedExpression.Operator)); err != nil {
			s.SetError(err)
		}

	case *cypher.Disjunction:
		for idx := 0; idx < typedExpression.Len()-1; idx++ {
			if err := s.treeTranslator.PopPushOperator(s.query.Scope, pgsql.OperatorOr); err != nil {
				s.SetError(err)
			}
		}

	case *cypher.Conjunction:
		for idx := 0; idx < typedExpression.Len()-1; idx++ {
			if err := s.treeTranslator.PopPushOperator(s.query.Scope, pgsql.OperatorAnd); err != nil {
				s.SetError(err)
			}
		}

	case *cypher.ProjectionItem:
		if err := s.translateProjectionItem(s.query.Scope, typedExpression); err != nil {
			s.SetError(err)
		}


	case *cypher.Match:
		if err := s.buildMatch(s.intermediates.match.Scope); err != nil {
			s.SetError(err)
		}

	case *cypher.SinglePartQuery:
		if s.intermediates.HasMutations() {
			if err := s.translateUpdates(s.query.Scope); err != nil {
				s.SetError(err)
			}

			if err := s.buildUpdates(s.query.Scope); err != nil {
				s.SetError(err)
			}
		}

		if s.intermediates.HasDeletions() {
			if err := s.buildDeletions(s.query.Scope); err != nil {
				s.SetError(err)
			}
		}

		// If there was no return specified end the CTE chain with a bare select
		if typedExpression.Return == nil {
			if literalReturn, err := pgsql.AsLiteral(1); err != nil {
				s.SetError(err)
			} else {
				s.query.Tail.Model.Body = pgsql.Select{
					Projection: []pgsql.SelectItem{literalReturn},
				}
			}
		} else if err := s.buildTailProjection(s.query.Scope); err != nil {
			s.SetError(err)
		}

		s.translation.Statement = *s.query.Tail.Model
	}
}

type Result struct {
	Statement  pgsql.Statement
	Parameters map[string]any
}

func Translate(ctx context.Context, cypherQuery *cypher.RegularQuery, kindMapper pgsql.KindMapper, parameters map[string]any) (Result, error) {
	translator := NewTranslator(ctx, kindMapper, parameters)

	if err := walk.WalkCypher(cypherQuery, translator); err != nil {
		return Result{}, err
	}

	return translator.translation, nil
}
