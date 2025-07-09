package osquery

import "testing"

func TestReverseNestedAggs(t *testing.T) {
	runMapTests(t, []mapTest{
		{
			"reverse_nested agg: basic",
			ReverseNestedAgg("to_parent"),
			map[string]interface{}{
				"reverse_nested": map[string]interface{}{},
			},
		},
		{
			"reverse_nested agg: with path",
			ReverseNestedAgg("to_root").Path("some.nested.path"),
			map[string]interface{}{
				"reverse_nested": map[string]interface{}{
					"path": "some.nested.path",
				},
			},
		},
		{
			"reverse_nested agg: with sub-aggregations",
			ReverseNestedAgg("to_parent").
				Aggs(Cardinality("product_count", "group_id")),
			map[string]interface{}{
				"reverse_nested": map[string]interface{}{},
				"aggs": map[string]interface{}{
					"product_count": map[string]interface{}{
						"cardinality": map[string]interface{}{
							"field": "group_id",
						},
					},
				},
			},
		},
	})
}
