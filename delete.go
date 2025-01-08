// Modified by DefenseStation on 2024-06-06
// Changes: Updated ElasticSearch client to OpenSearch client, changed package name to 'osquery',
// updated references to OpenSearch documentation, and modified examples accordingly.

package osquery

import (
	"bytes"
	"encoding/json"

	"github.com/opensearch-project/opensearch-go/v2"
	"github.com/opensearch-project/opensearch-go/v2/opensearchapi"
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
	o ...func(*opensearchapi.DeleteByQueryRequest),
) (res *opensearchapi.Response, err error) {
	return req.RunDelete(api.DeleteByQuery, o...)
}

// RunDelete is the same as the Run method, except that it accepts a value of
// type opensearchapi.DeleteByQuery (usually this is the DeleteByQuery field of an
// opensearch.Client object). Since the OpenSearch client does not provide
// an interface type for its API (which would allow implementation of mock
// clients), this provides a workaround. The Delete function in the OS client is
// actually a field of a function type.
func (req *DeleteRequest) RunDelete(
	del opensearchapi.DeleteByQuery,
	o ...func(*opensearchapi.DeleteByQueryRequest),
) (res *opensearchapi.Response, err error) {
	var b bytes.Buffer
	err = json.NewEncoder(&b).Encode(map[string]interface{}{
		"query": req.query.Map(),
	})
	if err != nil {
		return nil, err
	}

	return del(req.index, &b, o...)
}
