package osquery

import "testing"

func TestScriptField(t *testing.T) {
	runMapTests(t, []mapTest{
		{
			"script with source",
			Script("my_script").
				Source("doc['my_field'].value * params.factor").
				Params(ScriptParams{"factor": 2}).
				Lang("painless"),
			map[string]interface{}{
				"script": map[string]interface{}{
					"source": "doc['my_field'].value * params.factor",
					"params": map[string]interface{}{
						"factor": 2,
					},
					"lang": "painless",
				},
			},
		},
		{
			"script with all fields",
			Script("my_script").
				ID("my_id").
				Params(ScriptParams{"factor": 2}).
				Lang("painless"),
			map[string]interface{}{
				"script": map[string]interface{}{
					"id":     "my_id",
					"params": map[string]interface{}{"factor": 2},
					"lang":   "painless",
				},
			},
		},
	})
}
