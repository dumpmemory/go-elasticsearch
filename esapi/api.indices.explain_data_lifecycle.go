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
	"net/http"
	"strconv"
	"strings"
	"time"
)

func newIndicesExplainDataLifecycleFunc(t Transport) IndicesExplainDataLifecycle {
	return func(index string, o ...func(*IndicesExplainDataLifecycleRequest)) (*Response, error) {
		var r = IndicesExplainDataLifecycleRequest{Index: index}
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

// IndicesExplainDataLifecycle retrieves information about the index's current data stream lifecycle, such as any potential encountered error, time since creation etc.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/data-streams-explain-lifecycle.html.
type IndicesExplainDataLifecycle func(index string, o ...func(*IndicesExplainDataLifecycleRequest)) (*Response, error)

// IndicesExplainDataLifecycleRequest configures the Indices Explain Data Lifecycle API request.
type IndicesExplainDataLifecycleRequest struct {
	Index string

	IncludeDefaults *bool
	MasterTimeout   time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

// Do executes the request and returns response or error.
func (r IndicesExplainDataLifecycleRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "indices.explain_data_lifecycle")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "GET"

	path.Grow(7 + 1 + len(r.Index) + 1 + len("_lifecycle") + 1 + len("explain"))
	path.WriteString("http://")
	path.WriteString("/")
	path.WriteString(r.Index)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.RecordPathPart(ctx, "index", r.Index)
	}
	path.WriteString("/")
	path.WriteString("_lifecycle")
	path.WriteString("/")
	path.WriteString("explain")

	params = make(map[string]string)

	if r.IncludeDefaults != nil {
		params["include_defaults"] = strconv.FormatBool(*r.IncludeDefaults)
	}

	if r.MasterTimeout != 0 {
		params["master_timeout"] = formatDuration(r.MasterTimeout)
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

	req, err := newRequest(method, path.String(), nil)
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

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.BeforeRequest(req, "indices.explain_data_lifecycle")
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "indices.explain_data_lifecycle")
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
func (f IndicesExplainDataLifecycle) WithContext(v context.Context) func(*IndicesExplainDataLifecycleRequest) {
	return func(r *IndicesExplainDataLifecycleRequest) {
		r.ctx = v
	}
}

// WithIncludeDefaults - indicates if the api should return the default values the system uses for the index's lifecycle.
func (f IndicesExplainDataLifecycle) WithIncludeDefaults(v bool) func(*IndicesExplainDataLifecycleRequest) {
	return func(r *IndicesExplainDataLifecycleRequest) {
		r.IncludeDefaults = &v
	}
}

// WithMasterTimeout - specify timeout for connection to master.
func (f IndicesExplainDataLifecycle) WithMasterTimeout(v time.Duration) func(*IndicesExplainDataLifecycleRequest) {
	return func(r *IndicesExplainDataLifecycleRequest) {
		r.MasterTimeout = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f IndicesExplainDataLifecycle) WithPretty() func(*IndicesExplainDataLifecycleRequest) {
	return func(r *IndicesExplainDataLifecycleRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f IndicesExplainDataLifecycle) WithHuman() func(*IndicesExplainDataLifecycleRequest) {
	return func(r *IndicesExplainDataLifecycleRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f IndicesExplainDataLifecycle) WithErrorTrace() func(*IndicesExplainDataLifecycleRequest) {
	return func(r *IndicesExplainDataLifecycleRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f IndicesExplainDataLifecycle) WithFilterPath(v ...string) func(*IndicesExplainDataLifecycleRequest) {
	return func(r *IndicesExplainDataLifecycleRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f IndicesExplainDataLifecycle) WithHeader(h map[string]string) func(*IndicesExplainDataLifecycleRequest) {
	return func(r *IndicesExplainDataLifecycleRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f IndicesExplainDataLifecycle) WithOpaqueID(s string) func(*IndicesExplainDataLifecycleRequest) {
	return func(r *IndicesExplainDataLifecycleRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
