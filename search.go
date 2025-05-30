package osquery

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/opensearch-project/opensearch-go/v4"
	"github.com/opensearch-project/opensearch-go/v4/opensearchapi"
)

// SearchRequest represents the parameters for an OpenSearch query.
type SearchRequest struct {
	aggs         []Aggregation
	explain      *bool
	from         *uint64
	highlight    Mappable
	searchAfter  []interface{}
	postFilter   Mappable
	query        Mappable
	size         *uint64
	sort         []SortOption
	source       Source
	timeout      *time.Duration
	scriptFields []*ScriptField
}

// Search creates a new SearchRequest object, to be filled via method chaining.
func Search() *SearchRequest {
	return &SearchRequest{}
}

// Query sets a query for the request.
func (req *SearchRequest) Query(q Mappable) *SearchRequest {
	req.query = q
	return req
}

// Aggs sets one or more aggregations for the request.
func (req *SearchRequest) Aggs(aggs ...Aggregation) *SearchRequest {
	req.aggs = append(req.aggs, aggs...)
	return req
}

// PostFilter sets a post_filter for the request.
func (req *SearchRequest) PostFilter(filter Mappable) *SearchRequest {
	req.postFilter = filter
	return req
}

// From sets a document offset to start from.
func (req *SearchRequest) From(offset uint64) *SearchRequest {
	req.from = &offset
	return req
}

// Size sets the number of hits to return. The default - according to the OS
// documentation - is 10.
func (req *SearchRequest) Size(size uint64) *SearchRequest {
	req.size = &size
	return req
}

// Sort appends one or more sort options.
// Accepts any type that implements SortOption (field, script, raw)
func (req *SearchRequest) Sort(opts ...SortOption) *SearchRequest {
	if opts != nil {
		req.sort = append(req.sort, opts...)
	}
	return req
}

// SearchAfter retrieve the sorted result
func (req *SearchRequest) SearchAfter(s ...interface{}) *SearchRequest {
	req.searchAfter = append(req.searchAfter, s...)
	return req
}

// Explain sets whether the OpenSearch API should return an explanation for
// how each hit's score was calculated.
func (req *SearchRequest) Explain(b bool) *SearchRequest {
	req.explain = &b
	return req
}

// Timeout sets a timeout for the request.
func (req *SearchRequest) Timeout(dur time.Duration) *SearchRequest {
	req.timeout = &dur
	return req
}

// SourceIncludes sets the keys to return from the matching documents.
func (req *SearchRequest) SourceIncludes(keys ...string) *SearchRequest {
	req.source.includes = keys
	return req
}

// SourceExcludes sets the keys to not return from the matching documents.
func (req *SearchRequest) SourceExcludes(keys ...string) *SearchRequest {
	req.source.excludes = keys
	return req
}

// Highlight sets a highlight for the request.
func (req *SearchRequest) Highlight(highlight Mappable) *SearchRequest {
	req.highlight = highlight
	return req
}

func (req *SearchRequest) ScriptFields(fields ...*ScriptField) *SearchRequest {
	req.scriptFields = append(req.scriptFields, fields...)
	return req
}

// Map converts the SearchRequest to a map for the body.
func (req *SearchRequest) Map() map[string]interface{} {
	m := make(map[string]interface{})
	if req.query != nil {
		m["query"] = req.query.Map()
	}
	if len(req.aggs) > 0 {
		aggs := make(map[string]interface{})
		for _, agg := range req.aggs {
			aggs[agg.Name()] = agg.Map()
		}
		m["aggs"] = aggs
	}
	if req.postFilter != nil {
		m["post_filter"] = req.postFilter.Map()
	}
	if req.size != nil {
		m["size"] = *req.size
	}
	if len(req.sort) > 0 {
		sortSlice := make([]any, 0, len(req.sort))
		for _, params := range req.sort {
			sortSlice = append(sortSlice, params.Map())
		}
		m["sort"] = sortSlice
	}
	if req.from != nil {
		m["from"] = *req.from
	}
	if req.explain != nil {
		m["explain"] = *req.explain
	}
	if req.timeout != nil {
		m["timeout"] = fmt.Sprintf("%.0fs", req.timeout.Seconds())
	}
	if req.highlight != nil {
		m["highlight"] = req.highlight.Map()
	}
	if req.searchAfter != nil {
		m["search_after"] = req.searchAfter
	}

	if len(req.scriptFields) > 0 {
		scripts := make(map[string]interface{})
		for _, script := range req.scriptFields {
			scripts[script.Name()] = script.Map()
		}
		m["script_fields"] = scripts
	}
	source := req.source.Map()
	if len(source) > 0 {
		m["_source"] = source
	}

	return m
}

// MarshalJSON implements the json.Marshaler interface.
func (req *SearchRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(req.Map())
}

// Run executes the search using the OpenSearch client, applying additional options.
func (req *SearchRequest) Run(
	ctx context.Context,
	client *opensearch.Client,
	options *Options,
) (*opensearchapi.SearchResp, error) {
	// Serialize the request body to JSON
	body, err := json.Marshal(req.Map())
	if err != nil {
		return nil, fmt.Errorf("failed to serialize request body: %w", err)
	}

	// Create a Search request, setting size to 0 to avoid fetching documents
	searchReq := opensearchapi.SearchReq{
		Body: bytes.NewReader(body),
	}

	// Apply additional options if provided
	err = ApplyOptions(&searchReq, options)
	if err != nil {
		return nil, err
	}

	// fmt.Printf("%+v", searchReq)

	// Create a variable to hold the response
	var searchResp opensearchapi.SearchResp

	// Execute the search request using the OpenSearch client's Do method
	if _, err := client.Do(ctx, searchReq, &searchResp); err != nil {
		return nil, fmt.Errorf("search request failed: %w", err)
	}

	// Return the parsed response
	return &searchResp, nil
}

// Query is a shortcut for creating a SearchRequest with only a query. It is
// mostly included to maintain the API provided by osquery in early releases.
func Query(q Mappable) *SearchRequest {
	return Search().Query(q)
}

// Aggregate is a shortcut for creating a SearchRequest with aggregations. It is
// mostly included to maintain the API provided by osquery in early releases.
func Aggregate(aggs ...Aggregation) *SearchRequest {
	return Search().Aggs(aggs...)
}
