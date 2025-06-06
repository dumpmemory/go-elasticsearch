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

//go:build logging_default
// +build logging_default

// This collection of examples demonstrates how to configure the default logger of the client.
//
// To enable logging, pass a logger implementation to the "Logger" option in the client configuration.
//
// You can use one of the bundled loggers, or a custom "elastictransport.Logger" interface implementation.

package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9"
)

func main() {
	log.SetFlags(0)

	var es *elasticsearch.Client

	// ==============================================================================================
	//
	// "TextLogger" writes basic information about the request and response as plain text to the output.
	//
	es, _ = elasticsearch.NewClient(elasticsearch.Config{
		Logger: &elastictransport.TextLogger{Output: os.Stdout},
	})
	run(es, "Text")

	// ==============================================================================================
	//
	// "ColorLogger" is optimized for displaying information in the terminal during development.
	//
	es, _ = elasticsearch.NewClient(elasticsearch.Config{
		Logger: &elastictransport.ColorLogger{Output: os.Stdout},
	})
	run(es, "Color")

	// ==============================================================================================
	//
	// To log the request and response bodies, use the respective configuration options.
	//
	es, _ = elasticsearch.NewClient(elasticsearch.Config{
		Logger: &elastictransport.ColorLogger{
			Output:             os.Stdout,
			EnableRequestBody:  true,
			EnableResponseBody: true,
		},
	})
	run(es, "Request/Response Body")

	// ==============================================================================================
	//
	// "CurlLogger" writes the information formatted as runnable curl commands,
	// pretty-printing the response body (when enabled), useful eg. for sharing.
	//
	es, _ = elasticsearch.NewClient(elasticsearch.Config{
		Logger: &elastictransport.CurlLogger{Output: os.Stdout, EnableRequestBody: true, EnableResponseBody: true},
	})
	run(es, "Curl")

	// ==============================================================================================
	//
	// "JSONLogger" writes the information as JSON and is suitable for production logging.
	//
	es, _ = elasticsearch.NewClient(elasticsearch.Config{
		Logger: &elastictransport.JSONLogger{Output: os.Stdout},
	})
	run(es, "JSON")
}

// ------------------------------------------------------------------------------------------------

func run(es *elasticsearch.Client, name string) {
	log.Println("███", fmt.Sprintf("\x1b[1m%s\x1b[0m", name), strings.Repeat("█", 75-len(name)))

	es.Delete("test", "1")
	es.Exists("test", "1")

	es.Index(
		"test",
		strings.NewReader(`{"title" : "logging"}`),
		es.Index.WithRefresh("true"),
		es.Index.WithPretty(),
		es.Index.WithFilterPath("result", "_id"),
	)

	es.Search(es.Search.WithQuery("{FAIL"))

	res, err := es.Search(
		es.Search.WithIndex("test"),
		es.Search.WithBody(strings.NewReader(`{"query" : {"match" : { "title" : "logging" } } }`)),
		es.Search.WithSize(1),
		es.Search.WithPretty(),
		es.Search.WithFilterPath("took", "hits.hits"),
	)

	s := res.String()
	// log.Println("\x1b[1mResponse:\x1b[0m", s)
	if len(s) <= len("[200 OK] ") {
		log.Fatal("Response body is empty")
	}

	if err != nil {
		log.Fatalf("Error:   %s", err)
	}

	log.Print("\n")
}
