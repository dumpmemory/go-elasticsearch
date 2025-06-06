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
//
// Code generated from specification version 9.1.0: DO NOT EDIT

package esapi

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
	"time"
)

func newGraphExploreFunc(t Transport) GraphExplore {
	return func(index []string, o ...func(*GraphExploreRequest)) (*Response, error) {
		var r = GraphExploreRequest{Index: index}
		for _, f := range o {
			f(&r)
		}

		if transport, ok := t.(Instrumented); ok {
			r.Instrument = transport.InstrumentationEnabled()
		}

		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

// GraphExplore - Explore extracted and summarized information about the documents and terms in an index.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/graph-explore-api.html.
type GraphExplore func(index []string, o ...func(*GraphExploreRequest)) (*Response, error)

// GraphExploreRequest configures the Graph Explore API request.
type GraphExploreRequest struct {
	Index []string

	Body io.Reader

	Routing string
	Timeout time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

// Do executes the request and returns response or error.
func (r GraphExploreRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "graph.explore")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "POST"

	if len(r.Index) == 0 {
		return nil, errors.New("index is required and cannot be nil or empty")
	}

	path.Grow(7 + 1 + len(strings.Join(r.Index, ",")) + 1 + len("_graph") + 1 + len("explore"))
	path.WriteString("http://")
	path.WriteString("/")
	path.WriteString(strings.Join(r.Index, ","))
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.RecordPathPart(ctx, "index", strings.Join(r.Index, ","))
	}
	path.WriteString("/")
	path.WriteString("_graph")
	path.WriteString("/")
	path.WriteString("explore")

	params = make(map[string]string)

	if r.Routing != "" {
		params["routing"] = r.Routing
	}

	if r.Timeout != 0 {
		params["timeout"] = formatDuration(r.Timeout)
	}

	if r.Pretty {
		params["pretty"] = "true"
	}

	if r.Human {
		params["human"] = "true"
	}

	if r.ErrorTrace {
		params["error_trace"] = "true"
	}

	if len(r.FilterPath) > 0 {
		params["filter_path"] = strings.Join(r.FilterPath, ",")
	}

	req, err := newRequest(method, path.String(), r.Body)
	if err != nil {
		if instrument, ok := r.Instrument.(Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return nil, err
	}

	if len(params) > 0 {
		q := req.URL.Query()
		for k, v := range params {
			q.Set(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	if len(r.Header) > 0 {
		if len(req.Header) == 0 {
			req.Header = r.Header
		} else {
			for k, vv := range r.Header {
				for _, v := range vv {
					req.Header.Add(k, v)
				}
			}
		}
	}

	if r.Body != nil && req.Header.Get(headerContentType) == "" {
		req.Header[headerContentType] = headerContentTypeJSON
	}

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.BeforeRequest(req, "graph.explore")
		if reader := instrument.RecordRequestBody(ctx, "graph.explore", r.Body); reader != nil {
			req.Body = reader
		}
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "graph.explore")
	}
	if err != nil {
		if instrument, ok := r.Instrument.(Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return nil, err
	}

	response := Response{
		StatusCode: res.StatusCode,
		Body:       res.Body,
		Header:     res.Header,
	}

	return &response, nil
}

// WithContext sets the request context.
func (f GraphExplore) WithContext(v context.Context) func(*GraphExploreRequest) {
	return func(r *GraphExploreRequest) {
		r.ctx = v
	}
}

// WithBody - Graph Query DSL.
func (f GraphExplore) WithBody(v io.Reader) func(*GraphExploreRequest) {
	return func(r *GraphExploreRequest) {
		r.Body = v
	}
}

// WithRouting - specific routing value.
func (f GraphExplore) WithRouting(v string) func(*GraphExploreRequest) {
	return func(r *GraphExploreRequest) {
		r.Routing = v
	}
}

// WithTimeout - explicit operation timeout.
func (f GraphExplore) WithTimeout(v time.Duration) func(*GraphExploreRequest) {
	return func(r *GraphExploreRequest) {
		r.Timeout = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f GraphExplore) WithPretty() func(*GraphExploreRequest) {
	return func(r *GraphExploreRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f GraphExplore) WithHuman() func(*GraphExploreRequest) {
	return func(r *GraphExploreRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f GraphExplore) WithErrorTrace() func(*GraphExploreRequest) {
	return func(r *GraphExploreRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f GraphExplore) WithFilterPath(v ...string) func(*GraphExploreRequest) {
	return func(r *GraphExploreRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f GraphExplore) WithHeader(h map[string]string) func(*GraphExploreRequest) {
	return func(r *GraphExploreRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f GraphExplore) WithOpaqueID(s string) func(*GraphExploreRequest) {
	return func(r *GraphExploreRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
