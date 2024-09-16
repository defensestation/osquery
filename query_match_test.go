// Modified by DefenseStation on 2024-06-06
// Changes: Updated ElasticSearch client to OpenSearch client, changed package name to 'osquery',
// updated references to OpenSearch documentation, and modified examples accordingly.

package osquery

import (
	"testing"
)

func TestMatch(t *testing.T) {
	runMapTests(t, []mapTest{
		{
			"simple match",
			Match("title", "sample text"),
			map[string]interface{}{
				"match": map[string]interface{}{
					"title": map[string]interface{}{
						"query": "sample text",
					},
				},
			},
		},
		{
			"match with more params",
			Match("issue_number").Query(16).FuzzyTranspositions(false).MaxExpansions(32).Operator(OperatorAnd),
			map[string]interface{}{
				"match": map[string]interface{}{
					"issue_number": map[string]interface{}{
						"query":                16,
						"max_expansions":       32,
						"fuzzy_transpositions": false,
						"operator":             "AND",
					},
				},
			},
		},
		{
			"match_bool_prefix",
			MatchBoolPrefix("title", "sample text"),
			map[string]interface{}{
				"match_bool_prefix": map[string]interface{}{
					"title": map[string]interface{}{
						"query": "sample text",
					},
				},
			},
		},
		{
			"match_phrase",
			MatchPhrase("title", "sample text"),
			map[string]interface{}{
				"match_phrase": map[string]interface{}{
					"title": map[string]interface{}{
						"query": "sample text",
					},
				},
			},
		},
		{
			"match_phrase_prefix",
			MatchPhrasePrefix("title", "sample text"),
			map[string]interface{}{
				"match_phrase_prefix": map[string]interface{}{
					"title": map[string]interface{}{
						"query": "sample text",
					},
				},
			},
		},
	})
}
