// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

/*
Package esapi provides the Go API for Elasticsearch.

It is automatically included in the client provided by the
github.com/elastic/go-elasticsearch package:

	es, _ := elasticsearch.NewDefaultClient()
	res, _ := es.Info()
	log.Println(res)

For each Elasticsearch API, such as "Index", the package exports two corresponding types:
a function and a struct.

The function type allows you to call the Elasticsearch API as a method on the client,
passing the parameters as arguments:

	res, err := es.Index(
		"test",                                  // Index name
		strings.NewReader(`{"title" : "Test"}`), // Document body
		es.Index.WithDocumentID("1"),            // Document ID
		es.Index.WithRefresh("true"),            // Refresh
	)
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}
	defer res.Body.Close()

	log.Println(res)

	// => [201 Created] {"_index":"test","_type":"_doc","_id":"1" ...

The struct type allows for a more hands-on approach, where you create a new struct, with the
request configuration as fields, and call the Do() method
with a context and the client as arguments:

	req := esapi.IndexRequest{
		Index:      "test",                                  // Index name
		Body:       strings.NewReader(`{"title" : "Test"}`), // Document body
		DocumentID: "1",                                     // Document ID
		Refresh:    "true",                                  // Refresh
	}

	res, err := req.Do(context.Background(), es)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	log.Println(res)

	// => [200 OK] {"_index":"test","_type":"_doc","_id":"1","_version":2 ...

The function type is a wrapper around the struct, and allows
to configure and perform the request in a more expressive way.
It has a minor overhead compared to using a struct directly;
refer to the esapi_benchmark_test.go suite for concrete numbers.

See the documentation for each API function or struct at
https://pkg.go.dev/github.com/elastic/go-elasticsearch,
or locally by:

	go doc github.com/elastic/go-elasticsearch/v9/esapi Index
	go doc github.com/elastic/go-elasticsearch/v9/esapi IndexRequest

# Response

The esapi.Response type is a lightweight wrapper around http.Response.

The res.Body field is an io.ReadCloser, leaving the JSON parsing to the
calling code, in the interest of performance and flexibility
(eg. to allow using a custom JSON parser).

It is imperative to close the response body for a non-nil response.

The Response type implements a couple of convenience methods for accessing
the status, checking an error status code or printing
the response body for debugging purposes.

# Additional Information

See the Elasticsearch documentation at
https://www.elastic.co/guide/en/elasticsearch/reference/master/api-conventions.html for detailed information
about the API endpoints and parameters.

The Go API is generated from the Elasticsearch JSON specification at
https://github.com/elastic/elasticsearch/tree/master/rest-api-spec/src/main/resources/rest-api-spec/api
by the internal package available at
https://github.com/elastic/go-elasticsearch/tree/master/internal/cmd/generate/commands/gensource.

The API is tested by integration tests common to all Elasticsearch official clients, generated from the
source at https://github.com/elastic/elasticsearch/tree/master/rest-api-spec/src/main/resources/rest-api-spec/test.
The generator is provided by the internal package available at internal/cmd/generate/commands/gentests.
*/
package esapi
