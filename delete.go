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
	ctx context.Context,
	client *opensearch.Client,
	options *Options,
) (*opensearchapi.DocumentDeleteByQueryResp, error) {
	// Serialize the request body to JSON
	body, err := json.Marshal(req.query.Map())
	if err != nil {
		return nil, fmt.Errorf("failed to serialize request body: %w", err)
	}

	// Create a DeleteReq with the request body
	deleteReq := opensearchapi.DocumentDeleteByQueryReq{
		Body:    bytes.NewReader(body),                           // Pass the encoded request body
	}

	// Apply any additional options to modify the DeleteReq, such as context or index
	ApplyOptions(deleteReq, options)

	var deleteResp opensearchapi.DocumentDeleteByQueryResp

	// Execute the delete request using the OpenSearch client's Do method
    if _, err := client.Do(ctx, deleteReq, &deleteResp); err != nil {
        return nil, fmt.Errorf("delete request failed: %w", err)
    }

	return &deleteResp, nil
}
