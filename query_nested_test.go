package osquery

import "testing"

// Test cases for NestedQuery
func TestNestedQuery(t *testing.T) {
	runMapTests(t, []mapTest{
		{
			"nested query without optional fields",
			Nested("comments", Term("user", "kimchy")),
			map[string]interface{}{
				"nested": map[string]interface{}{
					"path": "comments",
					"query": map[string]interface{}{
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
			"nested query with score_mode",
			Nested("comments", Term("user", "kimchy")).ScoreMode(ScoreModeMax),
			map[string]interface{}{
				"nested": map[string]interface{}{
					"path": "comments",
					"query": map[string]interface{}{
						"term": map[string]interface{}{
							"user": map[string]interface{}{
								"value": "kimchy",
							},
						},
					},
					"score_mode": "max",
				},
			},
		},
		{
			"nested query with inner_hits",
			Nested("comments", Term("user", "kimchy")).InnerHits(map[string]interface{}{"size": 3}),
			map[string]interface{}{
				"nested": map[string]interface{}{
					"path": "comments",
					"query": map[string]interface{}{
						"term": map[string]interface{}{
							"user": map[string]interface{}{
								"value": "kimchy",
							},
						},
					},
					"inner_hits": map[string]interface{}{
						"size": 3,
					},
				},
			},
		},
	})
}
