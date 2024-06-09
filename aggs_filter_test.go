// Modified by DefenseStation on 2024-06-06
// Changes: Updated ElasticSearch client to OpenSearch client, changed package name to 'osquery',
// updated references to OpenSearch documentation, and modified examples accordingly.

package osquery

import "testing"

func TestFilterAggs(t *testing.T) {
	runMapTests(t, []mapTest{
		{
			"filter agg: simple",
			FilterAgg("filtered", Term("type", "t-shirt")),
			map[string]interface{}{
				"filter": map[string]interface{}{
					"term": map[string]interface{}{
						"type": map[string]interface{}{
							"value": "t-shirt",
						},
					},
				},
			},
		},
		{
			"filter agg: with aggs",
			FilterAgg("filtered", Term("type", "t-shirt")).
				Aggs(Avg("avg_price", "price")),
			map[string]interface{}{
				"filter": map[string]interface{}{
					"term": map[string]interface{}{
						"type": map[string]interface{}{
							"value": "t-shirt",
						},
					},
				},
				"aggs": map[string]interface{}{
					"avg_price": map[string]interface{}{
						"avg": map[string]interface{}{
							"field": "price",
						},
					},
				},
			},
		},
	})
}
