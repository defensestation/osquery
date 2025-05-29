// Package osquery modified by harshit98 on 2025-05-29
// Changes: Added sort params support like order, mode, nested_path, nested_filter, script-based sort
package osquery

import (
	"testing"
)

func TestSortExtensions(t *testing.T) {
	fieldSortWithOrder := FieldSort("field").Order(OrderAsc)
	fieldSortWithOrderAndMode := FieldSort("field").Order(OrderDesc).Mode(SortModeAvg)

	nestedFieldSort := FieldSort("nested.field").Order(OrderAsc).NestedPath("nested")

	nestedFieldSortWithFilter := FieldSort("nested.field").
		Order(OrderAsc).
		NestedPath("nested").
		NestedFilter(Match("nested.type").Query("value"))

	nestedFieldSortWithOrderAndMode := FieldSort("nested.field").
		Order(OrderDesc).
		Mode(SortModeMax).
		NestedPath("nested").
		NestedFilter(Match("nested.type").Query("value"))

	multipleSort1 := FieldSort("field1").Order(OrderAsc)

	multipleSort2 := FieldSort("nested.field").
		Order(OrderDesc).
		Mode(SortModeMin).
		NestedPath("nested").
		NestedFilter(Match("nested.type").Query("value"))

	runMapTests(t, []mapTest{
		{
			"sort with basic order only",
			Search().Sort(fieldSortWithOrder),
			map[string]any{
				"sort": []map[string]any{
					{
						"field": map[string]any{
							"order": "asc",
						},
					},
				},
			},
		},
		{
			"sort with mode",
			Search().Sort(fieldSortWithOrderAndMode),
			map[string]any{
				"sort": []map[string]any{
					{
						"field": map[string]any{
							"order": "desc",
							"mode":  "avg",
						},
					},
				},
			},
		},
		{
			"sort with nested_path",
			Search().Sort(nestedFieldSort),
			map[string]any{
				"sort": []map[string]any{
					{
						"nested.field": map[string]any{
							"order":       "asc",
							"nested_path": "nested",
						},
					},
				},
			},
		},
		{
			"sort with nested_path and nested_filter",
			Search().Sort(nestedFieldSortWithFilter),
			map[string]any{
				"sort": []map[string]any{
					{
						"nested.field": map[string]any{
							"order":       "asc",
							"nested_path": "nested",
							"nested_filter": map[string]any{
								"match": map[string]any{
									"nested.type": map[string]any{
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
			Search().Sort(nestedFieldSortWithOrderAndMode),
			map[string]any{
				"sort": []map[string]any{
					{
						"nested.field": map[string]any{
							"order":       "desc",
							"mode":        "max",
							"nested_path": "nested",
							"nested_filter": map[string]any{
								"match": map[string]any{
									"nested.type": map[string]any{
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
			Search().Sort(multipleSort1, multipleSort2),
			map[string]any{
				"sort": []map[string]any{
					{
						"field1": map[string]any{
							"order": "asc",
						},
					},
					{
						"nested.field": map[string]any{
							"order":       "desc",
							"mode":        "min",
							"nested_path": "nested",
							"nested_filter": map[string]any{
								"match": map[string]any{
									"nested.type": map[string]any{
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

func TestScriptSortExtensions(t *testing.T) {
	// Create script fields for reuse
	scriptFieldSort1 := Script("test_script").
		Source("doc['field_name'].value").
		Lang("painless")

	scriptFieldSort2 := Script("test_script").
		Source("doc['field_name'].value * params.factor").
		Lang("painless").
		Params(ScriptParams{"factor": 1.5})

	scriptFieldSort3 := Script("test_script").
		Source("if (doc['parent_obj.score_field'].size()!=0) { return ( Math.log(doc['parent_obj.score_field'].value*100 + 10 ) * _score ) } else { return _score }").
		Lang("painless")

	// Create script sort params for reuse
	scriptSortParams1 := ScriptSort(scriptFieldSort1, "number").Order(OrderDesc)
	scriptSortParams2 := ScriptSort(scriptFieldSort2, "number").Order(OrderAsc)
	scriptSortParams3 := ScriptSort(scriptFieldSort3, "number").Order(OrderDesc)

	docScoreFieldSort := FieldSort("_score")
	regularFieldSort := FieldSort("regular_field").Order(OrderAsc)

	runMapTests(t, []mapTest{
		{
			"sort with script",
			Search().Sort(scriptSortParams1),
			map[string]any{
				"sort": []map[string]any{
					{
						"_script": map[string]any{
							"type": "number",
							"script": map[string]any{
								"source": "doc['field_name'].value",
								"lang":   "painless",
							},
							"order": "desc",
						},
					},
				},
			},
		},
		{
			"sort with script and params",
			Search().Sort(scriptSortParams2),
			map[string]any{
				"sort": []map[string]any{
					{
						"_script": map[string]any{
							"type": "number",
							"script": map[string]any{
								"source": "doc['field_name'].value * params.factor",
								"lang":   "painless",
								"params": map[string]any{
									"factor": 1.5,
								},
							},
							"order": "asc",
						},
					},
				},
			},
		},
		{
			"sort with raw field and script",
			Search().Sort(docScoreFieldSort, scriptSortParams3),
			map[string]any{
				"sort": []any{
					map[string]any{
						"_score": map[string]any{},
					},
					map[string]any{
						"_script": map[string]any{
							"type": "number",
							"script": map[string]any{
								"source": "if (doc['parent_obj.score_field'].size()!=0) { return ( Math.log(doc['parent_obj.score_field'].value*100 + 10 ) * _score ) } else { return _score }",
								"lang":   "painless",
							},
							"order": "desc",
						},
					},
				},
			},
		},
		{
			"mixed sort with field and script",
			Search().Sort(regularFieldSort, scriptSortParams1),
			map[string]any{
				"sort": []map[string]any{
					{
						"regular_field": map[string]any{
							"order": "asc",
						},
					},
					{
						"_script": map[string]any{
							"type": "number",
							"script": map[string]any{
								"source": "doc['field_name'].value",
								"lang":   "painless",
							},
							"order": "desc",
						},
					},
				},
			},
		},
	})
}
