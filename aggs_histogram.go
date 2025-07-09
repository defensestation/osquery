package osquery

type HistogramAggregation struct {
	name        string
	field       string
	interval    float64
	offset      *float64
	minDocCount *int
	aggs        []Aggregation
}

// HistogramAgg creates a new histogram aggregation.
func HistogramAgg(name string, field string, interval float64) *HistogramAggregation {
	return &HistogramAggregation{
		name:     name,
		field:    field,
		interval: interval,
	}
}

// Name returns the name of the aggregation.
func (agg *HistogramAggregation) Name() string {
	return agg.name
}

// Offset sets an optional offset value.
func (agg *HistogramAggregation) Offset(offset float64) *HistogramAggregation {
	agg.offset = &offset
	return agg
}

// MinDocCount sets the optional minimum document count for buckets.
func (agg *HistogramAggregation) MinDocCount(min int) *HistogramAggregation {
	agg.minDocCount = &min
	return agg
}

// Aggs sets sub-aggregations for the histogram buckets.
func (agg *HistogramAggregation) Aggs(aggs ...Aggregation) *HistogramAggregation {
	agg.aggs = aggs
	return agg
}

// Map builds the OpenSearch aggregation map.
func (agg *HistogramAggregation) Map() map[string]interface{} {
	histogramMap := map[string]interface{}{
		"field":    agg.field,
		"interval": agg.interval,
	}

	if agg.offset != nil {
		histogramMap["offset"] = *agg.offset
	}
	if agg.minDocCount != nil {
		histogramMap["min_doc_count"] = *agg.minDocCount
	}

	outerMap := map[string]interface{}{
		"histogram": histogramMap,
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
