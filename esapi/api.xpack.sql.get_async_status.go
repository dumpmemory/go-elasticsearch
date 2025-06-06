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
	"strings"
)

func newSQLGetAsyncStatusFunc(t Transport) SQLGetAsyncStatus {
	return func(id string, o ...func(*SQLGetAsyncStatusRequest)) (*Response, error) {
		var r = SQLGetAsyncStatusRequest{DocumentID: id}
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

// SQLGetAsyncStatus - Returns the current status of an async SQL search or a stored synchronous SQL search
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/master/get-async-sql-search-status-api.html.
type SQLGetAsyncStatus func(id string, o ...func(*SQLGetAsyncStatusRequest)) (*Response, error)

// SQLGetAsyncStatusRequest configures the SQL Get Async Status API request.
type SQLGetAsyncStatusRequest struct {
	DocumentID string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

// Do executes the request and returns response or error.
func (r SQLGetAsyncStatusRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "sql.get_async_status")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "GET"

	path.Grow(7 + 1 + len("_sql") + 1 + len("async") + 1 + len("status") + 1 + len(r.DocumentID))
	path.WriteString("http://")
	path.WriteString("/")
	path.WriteString("_sql")
	path.WriteString("/")
	path.WriteString("async")
	path.WriteString("/")
	path.WriteString("status")
	path.WriteString("/")
	path.WriteString(r.DocumentID)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.RecordPathPart(ctx, "id", r.DocumentID)
	}

	params = make(map[string]string)

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
		instrument.BeforeRequest(req, "sql.get_async_status")
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "sql.get_async_status")
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
func (f SQLGetAsyncStatus) WithContext(v context.Context) func(*SQLGetAsyncStatusRequest) {
	return func(r *SQLGetAsyncStatusRequest) {
		r.ctx = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f SQLGetAsyncStatus) WithPretty() func(*SQLGetAsyncStatusRequest) {
	return func(r *SQLGetAsyncStatusRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f SQLGetAsyncStatus) WithHuman() func(*SQLGetAsyncStatusRequest) {
	return func(r *SQLGetAsyncStatusRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f SQLGetAsyncStatus) WithErrorTrace() func(*SQLGetAsyncStatusRequest) {
	return func(r *SQLGetAsyncStatusRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f SQLGetAsyncStatus) WithFilterPath(v ...string) func(*SQLGetAsyncStatusRequest) {
	return func(r *SQLGetAsyncStatusRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f SQLGetAsyncStatus) WithHeader(h map[string]string) func(*SQLGetAsyncStatusRequest) {
	return func(r *SQLGetAsyncStatusRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f SQLGetAsyncStatus) WithOpaqueID(s string) func(*SQLGetAsyncStatusRequest) {
	return func(r *SQLGetAsyncStatusRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
