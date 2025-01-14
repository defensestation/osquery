// Modified by DefenseStation on 2024-06-06
// Changes: Updated ElasticSearch client to OpenSearch client, changed package name to 'osquery',
// updated references to OpenSearch documentation, and modified examples accordingly.

package osquery

import (
	"context"
	"bytes"
	"encoding/json"
	"fmt"

	opensearch "github.com/opensearch-project/opensearch-go/v4"
	opensearchapi "github.com/opensearch-project/opensearch-go/v4/opensearchapi"
)

// CountRequest represents a request to get the number of matches for a search
// query, as described in:
// https://opensearch.org/docs/latest/api-reference/count/
type CountRequest struct {
	Query Mappable
}

// Count creates a new count request with the provided query.
func Count(q Mappable) *CountRequest {
	return &CountRequest{
		Query: q,
	}
}

// Map returns a map representation of the request, thus implementing the
// Mappable interface.
func (req *CountRequest) Map() map[string]interface{} {
	return map[string]interface{}{
		"query": req.Query.Map(),
	}
}

// Run executes the request using the provided OpenSearch client. It returns
// the HTTP response directly for further processing.
func (req *CountRequest) Run(
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
    ApplyOptions(&searchReq, options)

    // Create a variable to hold the response
    var searchResp opensearchapi.SearchResp

    // Execute the search request using the OpenSearch client's Do method
    if _, err := client.Do(ctx, searchReq, &searchResp); err != nil {
        return nil, fmt.Errorf("search request failed: %w", err)
    }

    // Return the parsed response
    return &searchResp, nil
}
