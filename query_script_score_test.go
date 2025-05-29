// Modified by Sushmita on 2025-01-31
// Changes: Updated ElasticSearch client to OpenSearch client, changed package name to 'osquery',
// updated references to OpenSearch documentation, and modified examples accordingly.

package osquery

import (
	"testing"
)

func TestScriptScore(t *testing.T) {
	runMapTests(t, []mapTest{
		{
			"script_score query without boost or min_score",
			ScriptScore(Term("user", "kimchy"), Script("my_script").Source("doc['field_name'].value * _score")),
			map[string]interface{}{
				"script_score": map[string]interface{}{
					"query": map[string]interface{}{
						"term": map[string]interface{}{
							"user": map[string]interface{}{
								"value": "kimchy",
							},
						},
					},
					"script": map[string]interface{}{
						"source": "doc['field_name'].value * _score",
					},
				},
			},
		},
		{
			"script_score query with boost",
			ScriptScore(Term("user", "kimchy"), Script("my_script").Source("doc['field_name'].value * _score")).Boost(1.0),
			map[string]interface{}{
				"script_score": map[string]interface{}{
					"query": map[string]interface{}{
						"term": map[string]interface{}{
							"user": map[string]interface{}{
								"value": "kimchy",
							},
						},
					},
					"script": map[string]interface{}{
						"source": "doc['field_name'].value * _score",
					},
					"boost": 1.0,
				},
			},
		},
		{
			"script_score query with min_score",
			ScriptScore(Term("user", "kimchy"), Script("my_script").Source("doc['field_name'].value * _score")).MinScore(2.2),
			map[string]interface{}{
				"script_score": map[string]interface{}{
					"query": map[string]interface{}{
						"term": map[string]interface{}{
							"user": map[string]interface{}{
								"value": "kimchy",
							},
						},
					},
					"script": map[string]interface{}{
						"source": "doc['field_name'].value * _score",
					},
					"min_score": 2.2,
				},
			},
		},
	})
}
