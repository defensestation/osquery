// Package osquery Modified by harshit98 on 2025-05-07
// Changes: Added sort params support like mode, nested_path, nested_filter
package osquery

import (
	"testing"
)

func TestSortExtensions(t *testing.T) {
	runMapTests(t, []mapTest{
		{
			"sort with basic order only",
			Search().Sort(SortParams{Field: "field", Order: OrderAsc}),
			map[string]interface{}{
				"sort": []map[string]interface{}{
					{
						"field": map[string]interface{}{
							"order": "asc",
						},
					},
				},
			},
		},
		{
			"sort with mode",
			Search().Sort(SortParams{Field: "field", Order: OrderDesc, Mode: SortModeAvg}),
			map[string]interface{}{
				"sort": []map[string]interface{}{
					{
						"field": map[string]interface{}{
							"order": "desc",
							"mode":  "avg",
						},
					},
				},
			},
		},
		{
			"sort with nested_path",
			Search().Sort(SortParams{Field: "nested.field", Order: OrderAsc, NestedPath: "nested"}),
			map[string]interface{}{
				"sort": []map[string]interface{}{
					{
						"nested.field": map[string]interface{}{
							"order":       "asc",
							"nested_path": "nested",
						},
					},
				},
			},
		},
		{
			"sort with nested_path and nested_filter",
			Search().Sort(SortParams{
				Field:        "nested.field",
				Order:        OrderAsc,
				NestedPath:   "nested",
				NestedFilter: Match("nested.type").Query("value"),
			}),
			map[string]interface{}{
				"sort": []map[string]interface{}{
					{
						"nested.field": map[string]interface{}{
							"order":       "asc",
							"nested_path": "nested",
							"nested_filter": map[string]interface{}{
								"match": map[string]interface{}{
									"nested.type": map[string]interface{}{
										"query": "value",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			"sort with mode, nested_path and nested_filter",
			Search().Sort(SortParams{
				Field:        "nested.field",
				Order:        OrderDesc,
				Mode:         SortModeMax,
				NestedPath:   "nested",
				NestedFilter: Match("nested.type").Query("value"),
			}),
			map[string]interface{}{
				"sort": []map[string]interface{}{
					{
						"nested.field": map[string]interface{}{
							"order":       "desc",
							"mode":        "max",
							"nested_path": "nested",
							"nested_filter": map[string]interface{}{
								"match": map[string]interface{}{
									"nested.type": map[string]interface{}{
										"query": "value",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			"multiple sorts with different options",
			Search().
				Sort(SortParams{Field: "field1", Order: OrderAsc}).
				Sort(SortParams{
					Field:        "nested.field",
					Order:        OrderDesc,
					Mode:         SortModeMin,
					NestedPath:   "nested",
					NestedFilter: Match("nested.type").Query("value"),
				}),
			map[string]interface{}{
				"sort": []map[string]interface{}{
					{
						"field1": map[string]interface{}{
							"order": "asc",
						},
					},
					{
						"nested.field": map[string]interface{}{
							"order":       "desc",
							"mode":        "min",
							"nested_path": "nested",
							"nested_filter": map[string]interface{}{
								"match": map[string]interface{}{
									"nested.type": map[string]interface{}{
										"query": "value",
									},
								},
							},
						},
					},
				},
			},
		},
	})
}
