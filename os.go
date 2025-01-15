// // Modified by DefenseStation on 2024-06-06
// Changes: Updated ElasticSearch client to OpenSearch client, changed package name to 'osquery',
// updated references to OpenSearch documentation, and modified examples accordingly.

// package osquery provides a non-obtrusive, idiomatic and easy-to-use query
// and aggregation builder for the official Go client
// (https://github.com/opensearch-project/opensearch-go/v4) for the OpenSearch
// database (https://opensearch.org/).
//
// osquery alleviates the need to use extremely nested maps
// (map[string]interface{}) and serializing queries to JSON manually. It also
// helps eliminating common mistakes such as misspelling query types, as
// everything is statically typed.
//
// Using `osquery` can make your code much easier to write, read and maintain,
// and significantly reduce the amount of code you write.
//
//
//
// Usage
//
//
//
// osquery provides a method chaining-style API for building and executing
// queries and aggregations. It does not wrap the official Go client nor does it
// require you to change your existing code in order to integrate the library.
// Queries can be directly built with `osquery`, and executed by passing an
// `*opensearch.Client` instance (with optional search parameters). Results
// are returned as-is from the official client (e.g. `*opensearchapi.Response` objects).
//
// Getting started is extremely simple:
//
//     package main
//
//     import (
//         "context"
//         "log"
//
//         "github.com/defensestatuib/osquery"
//         opensearch "github.com/opensearch-project/opensearch-go/v4"
//     )
//
//     func main() {
//         // connect to an OpenSearch instance
//         osclient, err := opensearch.NewDefaultClient()
//         if err != nil {
//             log.Fatalf("Failed creating client: %s", err)
//         }
//
//         // run a boolean search query
//         qRes, err := osquery.Query(
//             osquery.
//                 Bool().
//                 Must(osquery.Term("title", "Go and Stuff")).
//                 Filter(osquery.Term("tag", "tech")),
//             ).Run(
//                 osclient,
//                 osclient.Search.WithContext(context.TODO()),
//                 osclient.Search.WithIndex("test"),
//             )
//         if err != nil {
//             log.Fatalf("Failed searching for stuff: %s", err)
//         }
//
//         defer qRes.Body.Close()
//
//         // run an aggregation
//         aRes, err := osquery.Aggregate(
//             osquery.Avg("average_score", "score"),
//             osquery.Max("max_score", "score"),
//         ).Run(
//             osclient,
//             osclient.Search.WithContext(context.TODO()),
//             osclient.Search.WithIndex("test"),
//         )
//         if err != nil {
//             log.Fatalf("Failed searching for stuff: %s", err)
//         }
//
//         defer aRes.Body.Close()
//
//         // ...
//     }
//
//
//
// Notes
//
//
//
// * osquery currently supports version 7 of the OpenSearch Go client.
// * The library cannot currently generate "short queries". For example,
//  whereas OpenSearch can accept this:
//
//     { "query": { "term": { "user": "Kimchy" } } }
//
// The library will always generate this:
//
//     { "query": { "term": { "user": { "value": "Kimchy" } } } }
//
// This is also true for queries such as "bool", where fields like "must" can
// either receive one query object, or an array of query objects. `osquery` will
// generate an array even if there's only one query object.
// Modified by DefenseStation on 2024-06-06
// Changes: Updated ElasticSearch client to OpenSearch client, changed package name to 'osquery',
// updated references to OpenSearch documentation, and modified examples accordingly.

package osquery

// Mappable is the interface implemented by the various query and aggregation
// types provided by the package. It allows the library to easily transform the
// different queries to "generic" maps that can be easily encoded to JSON.
type Mappable interface {
	Map() map[string]interface{}
}

// Aggregation is an interface that each aggregation type must implement. It
// is simply an extension of the Mappable interface to include a Named function,
// which returns the name of the aggregation.
type Aggregation interface {
	Mappable
	Name() string
}
