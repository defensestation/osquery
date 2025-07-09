package osquery

type ReverseNestedAggregation struct {
	name string
	path *string       // Optional
	aggs []Aggregation // Optional sub-aggregations
}

// ReverseNestedAgg creates a new reverse_nested aggregation.
func ReverseNestedAgg(name string) *ReverseNestedAggregation {
	return &ReverseNestedAggregation{
		name: name,
	}
}

// Name returns the name of the aggregation.
func (agg *ReverseNestedAggregation) Name() string {
	return agg.name
}

// Path sets the optional reverse_nested path.
func (agg *ReverseNestedAggregation) Path(p string) *ReverseNestedAggregation {
	agg.path = &p
	return agg
}

// Aggs sets optional sub-aggregations.
func (agg *ReverseNestedAggregation) Aggs(aggs ...Aggregation) *ReverseNestedAggregation {
	agg.aggs = aggs
	return agg
}

// Map builds the OpenSearch aggregation map.
func (agg *ReverseNestedAggregation) Map() map[string]interface{} {
	reverseNestedBody := make(map[string]interface{})
	if agg.path != nil {
		reverseNestedBody["path"] = *agg.path
	}

	outerMap := map[string]interface{}{
		"reverse_nested": reverseNestedBody,
	}

	if len(agg.aggs) > 0 {
		subAggs := make(map[string]map[string]interface{})
		for _, sub := range agg.aggs {
			subAggs[sub.Name()] = sub.Map()
		}
		outerMap["aggs"] = subAggs
	}

	return outerMap
}
