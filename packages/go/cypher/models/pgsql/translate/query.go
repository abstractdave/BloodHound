package translate

import (
	"fmt"
	"github.com/specterops/bloodhound/cypher/models/cypher"
	"github.com/specterops/bloodhound/cypher/models/pgsql"
)

func (s *Translator) buildSinglePartQuery(singlePartQuery *cypher.SinglePartQuery) error {
	if s.query.CurrentPart().HasMutations() {
		if err := s.translateUpdates(s.query.Scope); err != nil {
			s.SetError(err)
		}

		if err := s.buildUpdates(s.query.Scope); err != nil {
			s.SetError(err)
		}
	}

	if s.query.CurrentPart().HasDeletions() {
		if err := s.buildDeletions(s.query.Scope); err != nil {
			s.SetError(err)
		}
	}

	// If there was no return specified end the CTE chain with a bare select
	if singlePartQuery.Return == nil {
		if literalReturn, err := pgsql.AsLiteral(1); err != nil {
			s.SetError(err)
		} else {
			s.query.CurrentPart().Model.Body = pgsql.Select{
				Projection: []pgsql.SelectItem{literalReturn},
			}
		}
	} else if err := s.buildTailProjection(s.query.Scope); err != nil {
		s.SetError(err)
	}

	return nil
}

func (s *Translator) buildMultiPartQuery(singlePartQuery *cypher.SinglePartQuery) error {
	topLevelQuery := pgsql.Query{
		CommonTableExpressions: &pgsql.With{},
	}

	for _, part := range s.query.Parts[:len(s.query.Parts)-1] {
		var nextCTE pgsql.CommonTableExpression

		if part.frame != nil {
			nextCTE.Alias = pgsql.TableAlias{
				Name: part.frame.Binding.Identifier,
			}
		}

		if inlineSelect, err := s.buildInlineProjection(s.query.Scope, part); err != nil {
			return err
		} else {
			nextCTE.Query.Body = inlineSelect
		}

		topLevelQuery.CommonTableExpressions.Expressions = append(topLevelQuery.CommonTableExpressions.Expressions, nextCTE)
	}

	if err := s.buildSinglePartQuery(singlePartQuery); err != nil {
		return err
	}

	topLevelQuery.Body = *s.query.CurrentPart().Model

	s.translation.Statement = topLevelQuery
	return nil
}

func (s *Translator) translateWith(with *cypher.With) error {
	currentPart := s.query.CurrentPart()

	if currentPart.HasProjections() {
		for _, projectionItem := range currentPart.projections.Items {
			if !projectionItem.Alias.Set {
				return fmt.Errorf("expected all with statement projections to contain an alias definition")
			} else {
				currentPart.frame.Export(projectionItem.Alias.Value)
			}
		}
	}

	return nil
}

func (s *Translator) translateMultiPartQueryPart(scope *Scope, part *cypher.MultiPartQueryPart) error {
	queryPart := s.query.CurrentPart()

	//if boundProjections, err := buildVisibleScopeProjections(scope, nil); err != nil {
	//	return err
	//} else {
	//	for _, boundProjection := range boundProjections.Items {
	//		queryPart.projections.Add(&Projection{
	//			SelectItem: boundProjection,
	//		})
	//	}
	//}

	// Unwind nested frames
	if err := scope.UnwindToFrame(queryPart.frame); err != nil {
		return err
	}

	// Pop the multipart query part's frame last
	return scope.PopFrame()
}
