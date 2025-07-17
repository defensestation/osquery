package osquery

import "testing"

func TestHistogramAggs(t *testing.T) {
	runMapTests(t, []mapTest{
		{
			"histogram agg: basic",
			HistogramAgg("hist", "price", 10),
			map[string]interface{}{
				"histogram": map[string]interface{}{
					"field":    "price",
					"interval": 10.0,
				},
			},
		},
		{
			"histogram agg: with offset and min_doc_count",
			HistogramAgg("hist", "price", 5).
				Offset(2).
				MinDocCount(1),
			map[string]interface{}{
				"histogram": map[string]interface{}{
					"field":         "price",
					"interval":      5.0,
					"offset":        2.0,
					"min_doc_count": 1,
				},
			},
		},
		{
			"histogram agg: with sub-aggs",
			HistogramAgg("hist", "price", 20).
				Aggs(Avg("avg_price", "price")),
			map[string]interface{}{
				"histogram": map[string]interface{}{
					"field":    "price",
					"interval": 20.0,
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
