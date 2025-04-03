package osquery

import "github.com/fatih/structs"

// KNNQuery represents a k-nearest-neighbors query for OpenSearch, as described in
// https://opensearch.org/docs/latest/query-dsl/specialized/k-nn/
type KNNQuery struct {
	field            string
	vector           []float64
	k                *int
	maxDistance      *float64
	minScore         *float64
	filter           map[string]interface{}
	methodParameters map[string]interface{}
	rescore          interface{}
	expandNestedDocs *bool
}

// KNN creates a new KNNQuery.
func KNN(field string, vector []float64) *KNNQuery {
	return &KNNQuery{
		field:  field,
		vector: vector,
	}
}

// K sets the k parameter.
func (q *KNNQuery) K(k int) *KNNQuery {
	q.k = &k
	return q
}

// MaxDistance sets the max_distance parameter.
func (q *KNNQuery) MaxDistance(d float64) *KNNQuery {
	q.maxDistance = &d
	return q
}

// MinScore sets the min_score parameter.
func (q *KNNQuery) MinScore(s float64) *KNNQuery {
	q.minScore = &s
	return q
}

// Filter sets the filter parameter.
func (q *KNNQuery) Filter(f map[string]interface{}) *KNNQuery {
	q.filter = f
	return q
}

// MethodParameters sets the method_parameters.
func (q *KNNQuery) MethodParameters(params map[string]interface{}) *KNNQuery {
	q.methodParameters = params
	return q
}

// Rescore sets the rescore parameter.
func (q *KNNQuery) Rescore(rescore interface{}) *KNNQuery {
	q.rescore = rescore
	return q
}

// ExpandNestedDocs sets the expand_nested_docs flag.
func (q *KNNQuery) ExpandNestedDocs(expand bool) *KNNQuery {
	q.expandNestedDocs = &expand
	return q
}

// Map returns a map representation of the query, thus implementing the
// Mappable interface.
func (q *KNNQuery) Map() map[string]interface{} {
	return map[string]interface{}{
		"knn": map[string]interface{}{
			q.field: structs.Map(struct {
				Vector           []float64              `structs:"vector"`
				K                *int                   `structs:"k,omitempty"`
				MaxDistance      *float64               `structs:"max_distance,omitempty"`
				MinScore         *float64               `structs:"min_score,omitempty"`
				Filter           map[string]interface{} `structs:"filter,omitempty"`
				MethodParameters map[string]interface{} `structs:"method_parameters,omitempty"`
				Rescore          interface{}            `structs:"rescore,omitempty"`
				ExpandNestedDocs *bool                  `structs:"expand_nested_docs,omitempty"`
			}{
				Vector:           q.vector,
				K:                q.k,
				MaxDistance:      q.maxDistance,
				MinScore:         q.minScore,
				Filter:           q.filter,
				MethodParameters: q.methodParameters,
				Rescore:          q.rescore,
				ExpandNestedDocs: q.expandNestedDocs,
			}),
		},
	}
}
