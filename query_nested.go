// Package osquery Modified by bforbhartiii on 2025-02-24
// Changes: Added nested query support
package osquery

import "github.com/fatih/structs"

type ScoreModeType string

const (
	// ScoreModeAvg Uses the average relevance score of all matching inner documents
	ScoreModeAvg ScoreModeType = "avg"

	// ScoreModeMax Assigns the highest relevance score from the matching inner documents to the parent.
	ScoreModeMax ScoreModeType = "max"

	// ScoreModeMin Assigns the lowest relevance score from the matching inner documents to the parent.
	ScoreModeMin ScoreModeType = "min"

	// ScoreModeSum Sums the relevance scores of all matching inner documents.
	ScoreModeSum ScoreModeType = "sum"

	// ScoreModeNone Ignores the relevance scores of inner documents and assigns a score of 0 to the parent document.
	ScoreModeNone ScoreModeType = "none"
)

// NestedQuery represents a compound query of type "nested",
// as described in https://opensearch.org/docs/latest/query-dsl/joining/nested/
type NestedQuery struct {
	path      string
	query     Mappable
	name      string
	scoreMode string
	innerHits map[string]interface{}
}

// Nested creates a new query of type "nested" with the provided path and query.
func Nested(path string, query Mappable) *NestedQuery {
	return &NestedQuery{
		path:  path,
		query: query,
	}
}

// ScoreMode sets the score mode of the query.
func (q *NestedQuery) ScoreMode(mode ScoreModeType) *NestedQuery {
	q.scoreMode = string(mode)
	return q
}

// InnerHits sets the inner_hits field of the query.
func (q *NestedQuery) InnerHits(innerHits map[string]interface{}) *NestedQuery {
	q.innerHits = innerHits
	return q
}

func (q *NestedQuery) Name(name string) *NestedQuery {
	q.name = name
	return q
}

// Map returns a map representation of the query, implementing the Mappable interface.
func (q *NestedQuery) Map() map[string]interface{} {
	return map[string]interface{}{
		"nested": structs.Map(struct {
			Path      string                 `structs:"path"`
			Query     map[string]interface{} `structs:"query"`
			Name      string                 `structs:"_name,omitempty"`
			ScoreMode string                 `structs:"score_mode,omitempty"`
			InnerHits map[string]interface{} `structs:"inner_hits,omitempty"`
		}{q.path, q.query.Map(), q.name, q.scoreMode, q.innerHits}),
	}
}
