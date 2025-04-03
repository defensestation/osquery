package osquery

import (
	"testing"
)

func TestKNNQuery_Map(t *testing.T) {
	tests := []mapTest{
		{
			"knn query without additional options",
			KNN("vector_field", []float64{1.0, 2.0, 3.0}).K(5),
			map[string]interface{}{
				"knn": map[string]interface{}{
					"vector_field": map[string]interface{}{
						"vector": []float64{1.0, 2.0, 3.0},
						"k":      5,
					},
				},
			},
		},
		{
			"knn query with max_distance option",
			KNN("vector_field", []float64{1.0, 2.0, 3.0}).K(5).MaxDistance(5.5),
			map[string]interface{}{
				"knn": map[string]interface{}{
					"vector_field": map[string]interface{}{
						"vector":       []float64{1.0, 2.0, 3.0},
						"k":            5,
						"max_distance": 5.5,
					},
				},
			},
		},
		{
			"knn query with min_score and filter",
			KNN("vector_field", []float64{0.1, 0.2}).K(5).MinScore(0.75).Filter(Term("status", "active").Map()),
			map[string]interface{}{
				"knn": map[string]interface{}{
					"vector_field": map[string]interface{}{
						"vector":    []float64{0.1, 0.2},
						"k":         5,
						"min_score": 0.75,
						"filter": map[string]interface{}{
							"term": map[string]interface{}{
								"status": map[string]interface{}{
									"value": "active",
								},
							},
						},
					},
				},
			},
		},
		{
			"knn query with method_parameters and expand_nested_docs",
			KNN("vector_field", []float64{4.0, 5.0, 6.0}).K(7).MethodParameters(
				map[string]interface{}{
					"nprobes": 3,
				},
			).ExpandNestedDocs(true),
			map[string]interface{}{
				"knn": map[string]interface{}{
					"vector_field": map[string]interface{}{
						"vector":             []float64{4.0, 5.0, 6.0},
						"k":                  7,
						"method_parameters":  map[string]interface{}{"nprobes": 3},
						"expand_nested_docs": true,
					},
				},
			},
		},
		{
			"knn query with rescore",
			KNN("vector_field", []float64{7, 8, 9}).K(3).Rescore(true),
			map[string]interface{}{
				"knn": map[string]interface{}{
					"vector_field": map[string]interface{}{
						"vector":  []float64{7, 8, 9},
						"k":       3,
						"rescore": true,
					},
				},
			},
		},
	}

	runMapTests(t, tests)
}
