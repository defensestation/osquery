package osquery

import (
	"fmt"
	"net/http"

	opensearchapi "github.com/opensearch-project/opensearch-go/v4/opensearchapi"
)

type Options struct {
	Indices []string
	Header  http.Header
	Params  interface{}
}

// ApplyOptions applies additional options to the request if provided.
// ApplyOptions applies additional options to the request if provided.
func ApplyOptions(req interface{}, options *Options) error {
	if options == nil {
		return nil
	}

	switch r := req.(type) {
	case *opensearchapi.SearchReq:
		if options.Indices != nil {
			r.Indices = options.Indices // Correctly assigning to the field
		}
		if options.Header != nil {
			r.Header = options.Header
		}
		if options.Params != nil {
			params, ok := options.Params.(*opensearchapi.SearchParams)
			if !ok {
				return fmt.Errorf("invalid type for SearchParams")
			}
			r.Params = *params
		}
	case *opensearchapi.DocumentDeleteByQueryReq:
		if options.Indices != nil {
			r.Indices = options.Indices
		}
		if options.Header != nil {
			r.Header = options.Header
		}
		if options.Params != nil {
			params, ok := options.Params.(*opensearchapi.DocumentDeleteByQueryParams)
			if !ok {
				return fmt.Errorf("invalid type for DocumentDeleteByQueryParams")
			}
			r.Params = *params
		}
	// Add more cases for other request types as needed
	default:
		return fmt.Errorf("unsupported request type: %T", req)
	}

	return nil
}
