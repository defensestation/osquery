// Modified by DefenseStation on 2024-06-06
// Changes: Updated ElasticSearch client to OpenSearch client, changed package name to 'osquery',
// updated references to OpenSearch documentation, and modified examples accordingly.

package osquery

import "testing"

func TestNestedAggs(t *testing.T) {
	runMapTests(t, []mapTest{
		{
			"nested agg: simple",
			NestedAgg("simple", "categories"),
			map[string]interface{}{
				"nested": map[string]interface{}{
					"path": "categories",
				},
			},
		},
		{
			"nested agg: with aggs",
			NestedAgg("more_nested", "authors").
				Aggs(TermsAgg("authors", "name")),
			map[string]interface{}{
				"nested": map[string]interface{}{
					"path": "authors",
				},
				"aggs": map[string]interface{}{
					"authors": map[string]interface{}{
						"terms": map[string]interface{}{
							"field": "name",
						},
					},
				},
			},
		},
	})
}
