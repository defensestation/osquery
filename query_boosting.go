// Modified by DefenseStation on 2024-06-06
// Changes: Updated ElasticSearch client to OpenSearch client, changed package name to 'osquery',
// updated references to OpenSearch documentation, and modified examples accordingly.

package osquery

// BoostingQuery represents a compound query of type "boosting", as described in
// https://opensearch.org/docs/latest/query-dsl/compound/boosting/
type BoostingQuery struct {
	// Pos is the positive part of the query.
	Pos Mappable
	// Neg is the negative part of the query.
	Neg Mappable
	// NegBoost is the negative boost value.
	NegBoost float32
}

// Boosting creates a new compound query of type "boosting".
func Boosting() *BoostingQuery {
	return &BoostingQuery{}
}

// Positive sets the positive part of the boosting query.
func (q *BoostingQuery) Positive(p Mappable) *BoostingQuery {
	q.Pos = p
	return q
}

// Negative sets the negative part of the boosting query.
func (q *BoostingQuery) Negative(p Mappable) *BoostingQuery {
	q.Neg = p
	return q
}

// NegativeBoost sets the negative boost value.
func (q *BoostingQuery) NegativeBoost(b float32) *BoostingQuery {
	q.NegBoost = b
	return q
}

// Map returns a map representation of the boosting query, thus implementing
// the Mappable interface.
func (q *BoostingQuery) Map() map[string]interface{} {
	return map[string]interface{}{
		"boosting": map[string]interface{}{
			"positive":       q.Pos.Map(),
			"negative":       q.Neg.Map(),
			"negative_boost": q.NegBoost,
		},
	}
}
