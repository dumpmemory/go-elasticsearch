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
)

func newUpdateByQueryRethrottleFunc(t Transport) UpdateByQueryRethrottle {
	return func(task_id string, requests_per_second *int, o ...func(*UpdateByQueryRethrottleRequest)) (*Response, error) {
		var r = UpdateByQueryRethrottleRequest{TaskID: task_id, RequestsPerSecond: requests_per_second}
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

// UpdateByQueryRethrottle changes the number of requests per second for a particular Update By Query operation.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-update-by-query.html.
type UpdateByQueryRethrottle func(task_id string, requests_per_second *int, o ...func(*UpdateByQueryRethrottleRequest)) (*Response, error)

// UpdateByQueryRethrottleRequest configures the Update By Query Rethrottle API request.
type UpdateByQueryRethrottleRequest struct {
	TaskID string

	RequestsPerSecond *int

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

// Do executes the request and returns response or error.
func (r UpdateByQueryRethrottleRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "update_by_query_rethrottle")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "POST"

	path.Grow(7 + 1 + len("_update_by_query") + 1 + len(r.TaskID) + 1 + len("_rethrottle"))
	path.WriteString("http://")
	path.WriteString("/")
	path.WriteString("_update_by_query")
	path.WriteString("/")
	path.WriteString(r.TaskID)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.RecordPathPart(ctx, "task_id", r.TaskID)
	}
	path.WriteString("/")
	path.WriteString("_rethrottle")

	params = make(map[string]string)

	if r.RequestsPerSecond != nil {
		params["requests_per_second"] = strconv.FormatInt(int64(*r.RequestsPerSecond), 10)
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
		instrument.BeforeRequest(req, "update_by_query_rethrottle")
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "update_by_query_rethrottle")
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
func (f UpdateByQueryRethrottle) WithContext(v context.Context) func(*UpdateByQueryRethrottleRequest) {
	return func(r *UpdateByQueryRethrottleRequest) {
		r.ctx = v
	}
}

// WithRequestsPerSecond - the throttle to set on this request in floating sub-requests per second. -1 means set no throttle..
func (f UpdateByQueryRethrottle) WithRequestsPerSecond(v int) func(*UpdateByQueryRethrottleRequest) {
	return func(r *UpdateByQueryRethrottleRequest) {
		r.RequestsPerSecond = &v
	}
}

// WithPretty makes the response body pretty-printed.
func (f UpdateByQueryRethrottle) WithPretty() func(*UpdateByQueryRethrottleRequest) {
	return func(r *UpdateByQueryRethrottleRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f UpdateByQueryRethrottle) WithHuman() func(*UpdateByQueryRethrottleRequest) {
	return func(r *UpdateByQueryRethrottleRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f UpdateByQueryRethrottle) WithErrorTrace() func(*UpdateByQueryRethrottleRequest) {
	return func(r *UpdateByQueryRethrottleRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f UpdateByQueryRethrottle) WithFilterPath(v ...string) func(*UpdateByQueryRethrottleRequest) {
	return func(r *UpdateByQueryRethrottleRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f UpdateByQueryRethrottle) WithHeader(h map[string]string) func(*UpdateByQueryRethrottleRequest) {
	return func(r *UpdateByQueryRethrottleRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f UpdateByQueryRethrottle) WithOpaqueID(s string) func(*UpdateByQueryRethrottleRequest) {
	return func(r *UpdateByQueryRethrottleRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
