// Modified by DefenseStation on 2024-06-06
// Changes: Updated ElasticSearch client to OpenSearch client, changed package name to 'osquery',
// updated references to OpenSearch documentation, and modified examples accordingly.

package osquery

import (
	"testing"
)

func TestConstantScore(t *testing.T) {
	runMapTests(t, []mapTest{
		{
			"constant_score query without boost",
			ConstantScore(Term("user", "kimchy")),
			map[string]interface{}{
				"constant_score": map[string]interface{}{
					"filter": map[string]interface{}{
						"term": map[string]interface{}{
							"user": map[string]interface{}{
								"value": "kimchy",
							},
						},
					},
				},
			},
		},
		{
			"constant_score query with boost",
			ConstantScore(Term("user", "kimchy")).Boost(2.2),
			map[string]interface{}{
				"constant_score": map[string]interface{}{
					"filter": map[string]interface{}{
						"term": map[string]interface{}{
							"user": map[string]interface{}{
								"value": "kimchy",
							},
						},
					},
					"boost": 2.2,
				},
			},
		},
		{
			"constant_score query with name",
			ConstantScore(Term("user", "kimchy")).Boost(10).Name("test"),
			map[string]interface{}{
				"constant_score": map[string]interface{}{
					"filter": map[string]interface{}{
						"term": map[string]interface{}{
							"user": map[string]interface{}{
								"value": "kimchy",
							},
						},
					},
					"boost": 10.0,
					"_name": "test",
				},
			},
		},
	})
}
