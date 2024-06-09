// Modified by DefenseStation on 2024-06-06
// Changes: Updated ElasticSearch client to OpenSearch client, changed package name to 'osquery',
// updated references to OpenSearch documentation, and modified examples accordingly.

package osquery

import (
	"bytes"
	"encoding/json"

	opensearch "github.com/opensearch-project/opensearch-go"
	opensearchapi "github.com/opensearch-project/opensearch-go/opensearchapi"
)

// CountRequest represents a request to get the number of matches for a search
// query, as described in:
// https://opensearch.org/docs/latest/api-reference/count/
type CountRequest struct {
	query Mappable
}

// Count creates a new count request with the provided query.
func Count(q Mappable) *CountRequest {
	return &CountRequest{
		query: q,
	}
}

// Map returns a map representation of the request, thus implementing the
// Mappable interface.
func (req *CountRequest) Map() map[string]interface{} {
	return map[string]interface{}{
		"query": req.query.Map(),
	}
}

// Run executes the request using the provided Count client. Zero or
// more search options can be provided as well. It returns the standard Response
// type of the official Go client.
func (req *CountRequest) Run(
	api *opensearch.Client,
	o ...func(*opensearchapi.CountRequest),
) (res *opensearchapi.Response, err error) {
	return req.RunCount(api.Count, o...)
}

// RunCount is the same as the Run method, except that it accepts a value of
// type opensearchapi.Count (usually this is the Count field of an opensearch.Client
// object). Since the OpenCount client does not provide an interface type
// for its API (which would allow implementation of mock clients), this provides
// a workaround. The Count function in the OS client is actually a field of a
// function type.
func (req *CountRequest) RunCount(
	count opensearchapi.Count,
	o ...func(*opensearchapi.CountRequest),
) (res *opensearchapi.Response, err error) {
	var b bytes.Buffer
	err = json.NewEncoder(&b).Encode(req.Map())
	if err != nil {
		return nil, err
	}

	opts := append([]func(*opensearchapi.CountRequest){count.WithBody(&b)}, o...)

	return count(opts...)
}
