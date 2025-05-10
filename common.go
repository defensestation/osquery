// Modified by DefenseStation on 2024-06-06
// Changes: Updated ElasticSearch client to OpenSearch client, changed package name to 'osquery',
// updated references to OpenSearch documentation, and modified examples accordingly.

package osquery

// Source represents the "_source" option which is commonly accepted in OS
// queries. Currently, only the "includes" option is supported.
type Source struct {
	includes []string
	excludes []string
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
	return m
}

// Sort represents a list of SortParams for sorting purpose.
type Sort []SortParams

// Order is the ordering for a sort key (ascending, descending).
type Order string

const (
	// OrderAsc represents sorting in ascending order.
	OrderAsc Order = "asc"

	// OrderDesc represents sorting in descending order.
	OrderDesc Order = "desc"
)

// Mode is the mode for a sort key (min, max, sum, avg, median).
type Mode string

const (
	// SortModeMin represents the minimum value.
	SortModeMin Mode = "min"

	// SortModeMax represents the maximum value.
	SortModeMax Mode = "max"

	// SortModeSum represents the sum of values.
	SortModeSum Mode = "sum"

	// SortModeAvg represents the average of values.
	SortModeAvg Mode = "avg"

	// SortModeMedian represents the median of values.
	SortModeMedian Mode = "median"
)

type SortParams struct {
	Field        string
	Order        Order
	Mode         Mode
	NestedPath   string
	NestedFilter Mappable
}

func (s SortParams) Map() map[string]interface{} {
	sortOptions := map[string]interface{}{}

	if s.Order != "" {
		sortOptions["order"] = s.Order
	}
	if s.Mode != "" {
		sortOptions["mode"] = s.Mode
	}
	if s.NestedPath != "" {
		sortOptions["nested_path"] = s.NestedPath

		if s.NestedFilter != nil {
			sortOptions["nested_filter"] = s.NestedFilter.Map()
		}
	}
	return map[string]interface{}{
		s.Field: sortOptions,
	}
}
