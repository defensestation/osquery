// Modified by DefenseStation on 2024-06-06
// Changes: Updated ElasticSearch client to OpenSearch client, changed package name to 'osquery',
// updated references to OpenSearch documentation, and modified examples accordingly.

package osquery

import (
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
	client *opensearch.Client,
o ...func(*opensearchapi.SearchReq),
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

	// Apply any additional options to modify the SearchReq, such as context or index
	for _, option := range o {
		option(&searchReq)
	}

	// Get the HTTP request from the SearchReq object
	httpRequest, err := searchReq.GetRequest()
	if err != nil {
		return nil, fmt.Errorf("failed to get HTTP request: %w", err)
	}

	// Perform the search request (this will give you the count of matched documents)
	res, err := client.Perform(httpRequest)
	if err != nil {
		return nil, fmt.Errorf("failed to execute search request: %w", err)
	}

	// Parse the response into the SearchResp struct
	var searchResp opensearchapi.SearchResp
	if err := json.NewDecoder(res.Body).Decode(&searchResp); err != nil {
		return nil, fmt.Errorf("failed to parse search response: %w", err)
	}

	return &searchResp, nil
}
