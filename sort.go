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

// SortOption is an interface for different types of sort options
type SortOption interface {
	Map() map[string]any
}

// ScriptSortParams represents a script-based sort option for elasticsearch
type ScriptSortParams struct {
	sortType string
	script   *ScriptField
	order    Order
}

// ScriptSort creates a new query of type "_script" with the provided
// type and script.
func ScriptSort(scriptField *ScriptField, sortType string) *ScriptSortParams {
	return &ScriptSortParams{
		script:   scriptField,
		sortType: sortType,
	}
}

func (s *ScriptSortParams) Order(order Order) *ScriptSortParams {
	s.order = order
	return s
}

func (s *ScriptSortParams) Map() map[string]any {
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

type SortParams struct {
	field        string
	order        Order
	mode         Mode
	nestedPath   string
	nestedFilter Mappable
}

func Field(field string) *SortParams {
	return &SortParams{
		field: field,
	}
}

func (s *SortParams) Order(order Order) *SortParams {
	s.order = order
	return s
}

func (s *SortParams) Mode(mode Mode) *SortParams {
	s.mode = mode
	return s
}

func (s *SortParams) NestedPath(nestedPath string) *SortParams {
	s.nestedPath = nestedPath
	return s
}

func (s *SortParams) NestedFilter(nestedFilter Mappable) *SortParams {
	s.nestedFilter = nestedFilter
	return s
}

func (s *SortParams) Map() map[string]any {
	sortOptions := map[string]any{}

	if s.order != "" {
		sortOptions["order"] = s.order
	}
	if s.mode != "" {
		sortOptions["mode"] = s.mode
	}
	if s.nestedPath != "" {
		sortOptions["nested_path"] = s.nestedPath

		if s.nestedFilter != nil {
			sortOptions["nested_filter"] = s.nestedFilter.Map()
		}
	}
	return map[string]any{
		s.field: sortOptions,
	}
}
