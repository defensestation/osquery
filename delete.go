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

// DeleteRequest represents a request to OpenSearch's Delete By Query API,
// described in
// https://opensearch.org/docs/latest/search-plugins/sql/sql/delete/
type DeleteRequest struct {
	index []string
	query Mappable
}

// Delete creates a new DeleteRequest object, to be filled via method chaining.
func Delete() *DeleteRequest {
	return &DeleteRequest{}
}

// Index sets the index names for the request
func (req *DeleteRequest) Index(index ...string) *DeleteRequest {
	req.index = index
	return req
}

// Query sets a query for the request.
func (req *DeleteRequest) Query(q Mappable) *DeleteRequest {
	req.query = q
	return req
}

// Run executes the request using the provided OpenSearch client.
func (req *DeleteRequest) Run(
	api *opensearch.Client,
	o ...func(*opensearchapi.DocumentDeleteByQueryReq),
) (*opensearchapi.DocumentDeleteByQueryResp, error) {
	var b bytes.Buffer
	// Convert the DeleteReq to a JSON-encoded body
	if err := json.NewEncoder(&b).Encode(req.query.Map()); err != nil {
		return nil, fmt.Errorf("failed to serialize request body: %w", err)
	}

	// Create a DeleteReq with the request body
	deleteReq := opensearchapi.DocumentDeleteByQueryReq{
		Body:    &b,                           // Pass the encoded request body
		Header:  nil,                          // Optional headers (you can set them if needed)
		Params: opensearchapi.DocumentDeleteByQueryParams{},   // Optional search parameters
	}

	// Apply any additional options to modify the DeleteReq, such as context or index
	for _, option := range o {
		option(&deleteReq)
	}

	// Get the HTTP request from the DeleteReq object
	httpRequest, err := deleteReq.GetRequest()
	if err != nil {
		return nil, fmt.Errorf("failed to get HTTP request: %w", err)
	}

	// Perform the search request using the `Perform` method
	res, err := api.Perform(httpRequest)
	if err != nil {
		return nil, fmt.Errorf("search request failed: %w", err)
	}

	// Parse the response into the DocumentDeleteByQueryResp struct
	var deleteResp opensearchapi.DocumentDeleteByQueryResp
	if err := json.NewDecoder(res.Body).Decode(&deleteResp); err != nil {
		return nil, fmt.Errorf("failed to parse delete response: %w", err)
	}

	return &deleteResp, nil
}
