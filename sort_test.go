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
			Search().Sort(Field("field").Order(OrderAsc)),
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
			Search().Sort(Field("field").Order(OrderDesc).Mode(SortModeAvg)),
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
			Search().Sort(Field("nested.field").Order(OrderAsc).NestedPath("nested")),
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
			Search().Sort(
				Field("nested.field").
					Order(OrderAsc).
					NestedPath("nested").
					NestedFilter(Match("nested.type").Query("value")),
			),
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
			Search().Sort(
				Field("nested.field").
					Order(OrderDesc).
					Mode(SortModeMax).
					NestedPath("nested").
					NestedFilter(Match("nested.type").Query("value")),
			),
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
			Search().
				Sort(Field("field1").Order(OrderAsc)).
				Sort(
					Field("nested.field").
						Order(OrderDesc).
						Mode(SortModeMin).
						NestedPath("nested").
						NestedFilter(Match("nested.type").Query("value")),
				),
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
	scriptField1 := Script("test_script").
		Source("doc['field_name'].value").
		Lang("painless")

	scriptField2 := Script("test_script").
		Source("doc['field_name'].value * params.factor").
		Lang("painless").
		Params(ScriptParams{"factor": 1.5})

	scriptField3 := Script("test_script").
		Source("if (doc['parent_obj.score_field'].size()!=0) { return ( Math.log(doc['parent_obj.score_field'].value*100 + 10 ) * _score ) } else { return _score }").
		Lang("painless")

	// Create script sort params for reuse
	scriptSortParams1 := ScriptSort(scriptField1, "number").Order(OrderDesc)
	scriptSortParams2 := ScriptSort(scriptField2, "number").Order(OrderAsc)
	scriptSortParams3 := ScriptSort(scriptField3, "number").Order(OrderDesc)

	runMapTests(t, []mapTest{
		{
			"sort with script",
			Search().SortScript(*scriptSortParams1),
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
			Search().SortScript(*scriptSortParams2),
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
			Search().
				Sort(Field("_score")).
				SortScript(*scriptSortParams3),
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
			Search().
				Sort(Field("regular_field").Order(OrderAsc)).
				SortScript(*scriptSortParams1),
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
