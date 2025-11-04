// Modified by DefenseStation on 2024-06-06
// Changes: Updated ElasticSearch client to OpenSearch client, changed package name to 'osquery',
// updated references to OpenSearch documentation, and modified examples accordingly.

package osquery

// Source represents the "_source" option which is commonly accepted in OS
// queries. Currently, only the "includes" option is supported.
type Source struct {
	includes []string
	excludes []string
	disabled bool
}

// Map returns a map representation of the Source object.
func (source Source) Map() map[string]interface{} {
	m := make(map[string]interface{})
	if len(source.includes) > 0 {
		m["includes"] = source.includes
	}
	if len(source.excludes) > 0 {
		m["excludes"] = source.excludes
	}
	if source.disabled {
		m["enabled"] = false
	}
	return m
}
