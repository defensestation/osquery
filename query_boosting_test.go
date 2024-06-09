// Modified by DefenseStation on 2024-06-06
// Changes: Updated ElasticSearch client to OpenSearch client, changed package name to 'osquery',
// updated references to OpenSearch documentation, and modified examples accordingly.

package osquery

import (
	"testing"
)

func TestBoosting(t *testing.T) {
	runMapTests(t, []mapTest{
		{
			"boosting query",
			Boosting().
				Positive(Term("text", "apple")).
				Negative(Term("text", "pie tart")).
				NegativeBoost(0.5),
			map[string]interface{}{
				"boosting": map[string]interface{}{
					"positive": map[string]interface{}{
						"term": map[string]interface{}{
							"text": map[string]interface{}{
								"value": "apple",
							},
						},
					},
					"negative": map[string]interface{}{
						"term": map[string]interface{}{
							"text": map[string]interface{}{
								"value": "pie tart",
							},
						},
					},
					"negative_boost": 0.5,
				},
			},
		},
	})
}
