// Modified by DefenseStation on 2024-06-06
// Changes: Updated ElasticSearch client to OpenSearch client, changed package name to 'osquery',
// updated references to OpenSearch documentation, and modified examples accordingly.

package osquery

import "github.com/fatih/structs"

// DisMaxQuery represents a compound query of type "dis_max", as described in
// https://opensearch.org/docs/1.3/query-dsl/compound/disjunction-max/
type DisMaxQuery struct {
	queries    []Mappable
	tieBreaker float32
}

// DisMax creates a new compound query of type "dis_max" with the provided
// queries.
func DisMax(queries ...Mappable) *DisMaxQuery {
	return &DisMaxQuery{
		queries: queries,
	}
}

// TieBreaker sets the "tie_breaker" value for the query.
func (q *DisMaxQuery) TieBreaker(b float32) *DisMaxQuery {
	q.tieBreaker = b
	return q
}

// Map returns a map representation of the dis_max query, thus implementing
// the Mappable interface.
func (q *DisMaxQuery) Map() map[string]interface{} {
	inner := make([]map[string]interface{}, len(q.queries))
	for i, iq := range q.queries {
		inner[i] = iq.Map()
	}
	return map[string]interface{}{
		"dis_max": structs.Map(struct {
			Queries    []map[string]interface{} `structs:"queries"`
			TieBreaker float32                  `structs:"tie_breaker,omitempty"`
		}{inner, q.tieBreaker}),
	}
}
