// Modified by DefenseStation on 2024-06-06
// Changes: Updated ElasticSearch client to OpenSearch client, changed package name to 'osquery',
// updated references to OpenSearch documentation, and modified examples accordingly.

package osquery

import (
	context "context"

	opensearch "github.com/opensearch-project/opensearch-go/v4"
	opensearchapi "github.com/opensearch-project/opensearch-go/v4/opensearchapi"
)

// CustomQueryMap represents an arbitrary query map for custom queries.
type CustomQueryMap map[string]interface{}

// CustomQuery generates a custom request of type "query" from an arbitrary map
// provided by the user. It is useful for issuing a search request with a syntax
// that is not yet supported by the library. CustomQuery values are versatile,
// they can either be used as parameters for the library's Query function, or
// standalone by invoking their Run method.
func CustomQuery(m map[string]interface{}) *CustomQueryMap {
	q := CustomQueryMap(m)
	return &q
}

// Map returns the custom query as a map[string]interface{}, thus implementing
// the Mappable interface.
func (m *CustomQueryMap) Map() map[string]interface{} {
	return *m
}

// Run executes the custom query using the provided OpenSearch client. Zero
// or more search options can be provided as well. It returns the standard
// Response type of the official Go client.
func (m *CustomQueryMap) Run(
	ctx context.Context,
	api *opensearch.Client,
	options *Options,
) (res *opensearchapi.SearchResp, err error) {
	return Search().Query(m).Run(ctx, api, options)
}

//----------------------------------------------------------------------------//

// CustomAggMap represents an arbitrary aggregation map for custom aggregations.
type CustomAggMap struct {
	name string
	agg  map[string]interface{}
}

// CustomAgg generates a custom aggregation from an arbitrary map provided by
// the user.
func CustomAgg(name string, m map[string]interface{}) *CustomAggMap {
	return &CustomAggMap{
		name: name,
		agg:  m,
	}
}

// Name returns the name of the aggregation
func (agg *CustomAggMap) Name() string {
	return agg.name
}

// Map returns a map representation of the custom aggregation, thus implementing
// the Mappable interface
func (agg *CustomAggMap) Map() map[string]interface{} {
	return agg.agg
}
