package translate

import (
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

	s.translation.Statement = *s.query.CurrentPart().Model
	return nil
}

func (s *Translator) buildMultiPartQuery() error {
	topLevelQuery := pgsql.Query{
		CommonTableExpressions: &pgsql.With{},
		Body:                   pgsql.Select{},
	}

	for _, part := range s.query.Parts {
		topLevelQuery.CommonTableExpressions.Expressions = append(topLevelQuery.CommonTableExpressions.Expressions, pgsql.CommonTableExpression{
			Alias: pgsql.TableAlias{
				Name: part.frame.Binding.Identifier,
			},
			Query: *part.Model,
		})
	}

	// Stopped here - still gotta keep wiring up the chained CTEs
	s.translation.Statement = topLevelQuery
	return nil
}

func (s *Translator) translateMultiPartQueryPart(scope *Scope, part *cypher.MultiPartQueryPart) error {
	if boundProjections, err := buildVisibleScopeProjections(scope, nil); err != nil {
		return err
	} else {
		for _, boundProjection := range boundProjections.Items {
			s.query.CurrentPart().projections.Add(&Projection{
				SelectItem: boundProjection,
			})
		}
	}

	return nil
}
