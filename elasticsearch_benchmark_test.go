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

//go:build !integration
// +build !integration

package elasticsearch_test

import (
	"context"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"testing"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

var defaultResponse = http.Response{
	Status:        "200 OK",
	StatusCode:    200,
	ContentLength: 2,
	Header: http.Header(map[string][]string{
		"Content-Type":      {"application/json"},
		"X-Elastic-Product": {"Elasticsearch"},
	}),
	Body: ioutil.NopCloser(strings.NewReader(`{}`)),
}

type FakeTransport struct {
	FakeResponse *http.Response
}

func (t *FakeTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return t.FakeResponse, nil
}

func newFakeTransport() *FakeTransport {
	return &FakeTransport{FakeResponse: &defaultResponse}
}

func BenchmarkClient(b *testing.B) {
	b.ReportAllocs()

	b.Run("Create client with defaults", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := elasticsearch.NewDefaultClient()

			if err != nil {
				b.Fatalf("Unexpected error when creating a client: %s", err)
			}
		}
	})
}

func BenchmarkClientAPI(b *testing.B) {
	b.ReportAllocs()

	ctx := context.Background()

	client, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
		Transport: newFakeTransport(),
	})
	if err != nil {
		b.Fatalf("ERROR: %s", err)
	}

	b.Run("InfoRequest{}.Do()", func(b *testing.B) {
		b.ResetTimer()

		req := esapi.InfoRequest{}

		for i := 0; i < b.N; i++ {
			if _, err := req.Do(ctx, client); err != nil {
				b.Errorf("Unexpected error when getting a response: %s", err)
			}
		}
	})

	b.Run("IndexRequest{...}.Do()", func(b *testing.B) {
		b.ResetTimer()
		var body strings.Builder

		for i := 0; i < b.N; i++ {
			docID := strconv.FormatInt(int64(i), 10)

			body.Reset()
			body.WriteString(`{"foo" : "bar `)
			body.WriteString(docID)
			body.WriteString(`	" }`)

			req := esapi.IndexRequest{
				Index:      "test",
				DocumentID: docID,
				Body:       strings.NewReader(body.String()),
				Refresh:    "true",
				Pretty:     true,
				Timeout:    100,
			}
			if _, err := req.Do(ctx, client); err != nil {
				b.Errorf("Unexpected error when getting a response: %s", err)
			}
		}
	})

	b.Run("Index(...)", func(b *testing.B) {
		b.ResetTimer()
		var body strings.Builder

		for i := 0; i < b.N; i++ {
			docID := strconv.FormatInt(int64(i), 10)

			body.Reset()
			body.WriteString(`{"foo" : "bar `)
			body.WriteString(docID)
			body.WriteString(`	" }`)

			_, err := client.Index(
				"test",
				strings.NewReader(body.String()),
				client.Index.WithDocumentID(docID),
				client.Index.WithRefresh("true"),
				client.Index.WithPretty(),
				client.Index.WithTimeout(100),
				client.Index.WithContext(ctx),
			)
			if err != nil {
				b.Errorf("Unexpected error when getting a response: %s", err)
			}
		}
	})

	b.Run("SearchRequest{...}.Do()", func(b *testing.B) {
		b.ResetTimer()

		body := `{"foo" : "bar"}`
		indx := []string{"test"}

		for i := 0; i < b.N; i++ {
			req := esapi.SearchRequest{
				Index:   indx,
				Body:    strings.NewReader(body),
				Size:    esapi.IntPtr(25),
				Pretty:  true,
				Timeout: 100,
			}
			if _, err := req.Do(ctx, client); err != nil {
				b.Errorf("Unexpected error when getting a response: %s", err)
			}
		}
	})

	b.Run("Search(...)", func(b *testing.B) {
		b.ResetTimer()

		body := `{"foo" : "bar"}`
		indx := "test"

		for i := 0; i < b.N; i++ {
			_, err := client.Search(
				client.Search.WithIndex(indx),
				client.Search.WithBody(strings.NewReader(body)),
				client.Search.WithSize(25),
				client.Search.WithPretty(),
				client.Search.WithTimeout(100),
				client.Search.WithContext(ctx),
			)
			if err != nil {
				b.Errorf("Unexpected error when getting a response: %s", err)
			}
		}
	})

	b.Run("BulkRequest{...}.Do()", func(b *testing.B) {
		b.ResetTimer()
		var body strings.Builder

		for i := 0; i < b.N; i++ {
			docID := strconv.FormatInt(int64(i), 10)

			body.Reset()
			body.WriteString(`{"index" : { "_index" : "test", "_type" : "_doc", "_id" : "` + docID + `" }}`)
			body.WriteString(`{"foo" : "bar `)
			body.WriteString(docID)
			body.WriteString(`	" }`)

			req := esapi.BulkRequest{
				Body:    strings.NewReader(body.String()),
				Refresh: "true",
				Pretty:  true,
				Timeout: 100,
			}
			if _, err := req.Do(ctx, client); err != nil {
				b.Errorf("Unexpected error when getting a response: %s", err)
			}
		}
	})

	b.Run("Bulk(...)", func(b *testing.B) {
		b.ResetTimer()
		var body strings.Builder

		for i := 0; i < b.N; i++ {
			docID := strconv.FormatInt(int64(i), 10)

			body.Reset()
			body.WriteString(`{"index" : { "_index" : "test", "_type" : "_doc", "_id" : "` + docID + `" }}`)
			body.WriteString(`{"foo" : "bar `)
			body.WriteString(docID)
			body.WriteString(`	" }`)

			_, err := client.Bulk(
				strings.NewReader(body.String()),
				client.Bulk.WithRefresh("true"),
				client.Bulk.WithPretty(),
				client.Bulk.WithTimeout(100),
				client.Bulk.WithContext(ctx),
			)
			if err != nil {
				b.Errorf("Unexpected error when getting a response: %s", err)
			}
		}
	})
}
