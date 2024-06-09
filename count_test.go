// Modified by DefenseStation on 2024-06-06
// Changes: Updated ElasticSearch client to OpenSearch client, changed package name to 'osquery',
// updated references to OpenSearch documentation, and modified examples accordingly.

package osquery

import "testing"

func TestCount(t *testing.T) {
	runMapTests(t, []mapTest{
		{
			"a simple count request",
			Count(MatchAll()),
			map[string]interface{}{
				"query": map[string]interface{}{
					"match_all": map[string]interface{}{},
				},
			},
		},
	})
}
