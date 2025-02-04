// Modified by Sushmita on 2025-01-31
// Changes: Updated ElasticSearch client to OpenSearch client, changed package name to 'osquery',
// updated references to OpenSearch documentation, and modified examples accordingly.

package osquery

import "github.com/fatih/structs"

// ScriptScoreQuery represents a compound query of type "script_score", as
// described in
// https://opensearch.org/docs/latest/query-dsl/specialized/script-score/
type ScriptScoreQuery struct {
	query    Mappable
	script   ScriptField
	minScore float32
	boost    float32
}

// ScriptScore creates a new query of type "script_score" with the provided
// query and script.
func ScriptScore(query Mappable, script *ScriptField) *ScriptScoreQuery {
	return &ScriptScoreQuery{
		query:  query,
		script: *script,
	}
}

// Boost sets the boost value of the query.
func (q *ScriptScoreQuery) Boost(b float32) *ScriptScoreQuery {
	q.boost = b
	return q
}

// MinScore sets the minimum score of the query.
func (q *ScriptScoreQuery) MinScore(min float32) *ScriptScoreQuery {
	q.minScore = min
	return q
}

// Map returns a map representation of the query, thus implementing the
// Mappable interface.
func (q *ScriptScoreQuery) Map() map[string]interface{} {
	script := q.script.Map()["script"].(map[string]interface{})
	return map[string]interface{}{
		"script_score": structs.Map(struct {
			Query    map[string]interface{} `structs:"query"`
			Script   map[string]interface{} `structs:"script"`
			Boost    float32                `structs:"boost,omitempty"`
			MinScore float32                `structs:"min_score,omitempty"`
		}{q.query.Map(), script, q.boost, q.minScore}),
	}
}
