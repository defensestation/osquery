package osquery

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

// Missing allows you to control where documents with missing values are placed in the sorted results
type Missing string

const (
	// MissingLast represents Documents missing the field are treated as having the lowest possible value and are sorted to the end of the results
	MissingLast Missing = "_last"

	// MissingFirst represents Documents missing the field are treated as having the highest possible value and are sorted to the beginning of the results
	MissingFirst Missing = "_first"
)

// SortOption is an interface for different types of sort options
type SortOption interface {
	Map() map[string]any
}

// ScriptSortOption represents a script-based sort option for elasticsearch
type ScriptSortOption struct {
	sortType string
	script   *ScriptField
	order    Order
}

// ScriptSort creates a new query of type "_script" with the provided
// type and script.
func ScriptSort(scriptField *ScriptField, sortType string) *ScriptSortOption {
	return &ScriptSortOption{
		script:   scriptField,
		sortType: sortType,
	}
}

func (s *ScriptSortOption) Order(order Order) *ScriptSortOption {
	s.order = order
	return s
}

func (s *ScriptSortOption) Map() map[string]any {
	scriptMapRaw, ok := s.script.Map()["script"]
	if !ok {
		return nil
	}

	scriptMap, ok := scriptMapRaw.(map[string]any)
	if !ok {
		return nil
	}

	sortOptions := map[string]any{
		"type":   s.sortType,
		"script": scriptMap,
	}

	if s.order != "" {
		sortOptions["order"] = s.order
	}

	return map[string]any{
		"_script": sortOptions,
	}
}

type FieldSortOption struct {
	field        string
	order        Order
	mode         Mode
	missing      Missing
	nestedPath   string
	nestedFilter Mappable
}

func FieldSort(field string) *FieldSortOption {
	return &FieldSortOption{
		field: field,
	}
}

func (f *FieldSortOption) Order(order Order) *FieldSortOption {
	f.order = order
	return f
}

func (f *FieldSortOption) Mode(mode Mode) *FieldSortOption {
	f.mode = mode
	return f
}

func (f *FieldSortOption) Missing(missing Missing) *FieldSortOption {
	f.missing = missing
	return f
}

func (f *FieldSortOption) NestedPath(nestedPath string) *FieldSortOption {
	f.nestedPath = nestedPath
	return f
}

func (f *FieldSortOption) NestedFilter(nestedFilter Mappable) *FieldSortOption {
	f.nestedFilter = nestedFilter
	return f
}

func (f *FieldSortOption) Map() map[string]any {
	sortOptions := map[string]any{}

	if f.order != "" {
		sortOptions["order"] = f.order
	}

	if f.mode != "" {
		sortOptions["mode"] = f.mode
	}

	if f.missing != "" {
		sortOptions["missing"] = f.missing
	}

	if f.nestedPath != "" {
		sortOptions["nested_path"] = f.nestedPath

		if f.nestedFilter != nil {
			sortOptions["nested_filter"] = f.nestedFilter.Map()
		}
	}

	return map[string]any{
		f.field: sortOptions,
	}
}
