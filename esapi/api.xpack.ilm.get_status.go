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

func newILMGetStatusFunc(t Transport) ILMGetStatus {
	return func(o ...func(*ILMGetStatusRequest)) (*Response, error) {
		var r = ILMGetStatusRequest{}
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

// ILMGetStatus - Retrieves the current index lifecycle management (ILM) status.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/ilm-get-status.html.
type ILMGetStatus func(o ...func(*ILMGetStatusRequest)) (*Response, error)

// ILMGetStatusRequest configures the ILM Get Status API request.
type ILMGetStatusRequest struct {
	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

// Do executes the request and returns response or error.
func (r ILMGetStatusRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "ilm.get_status")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "GET"

	path.Grow(7 + len("/_ilm/status"))
	path.WriteString("http://")
	path.WriteString("/_ilm/status")

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
		instrument.BeforeRequest(req, "ilm.get_status")
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "ilm.get_status")
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
func (f ILMGetStatus) WithContext(v context.Context) func(*ILMGetStatusRequest) {
	return func(r *ILMGetStatusRequest) {
		r.ctx = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f ILMGetStatus) WithPretty() func(*ILMGetStatusRequest) {
	return func(r *ILMGetStatusRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f ILMGetStatus) WithHuman() func(*ILMGetStatusRequest) {
	return func(r *ILMGetStatusRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f ILMGetStatus) WithErrorTrace() func(*ILMGetStatusRequest) {
	return func(r *ILMGetStatusRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f ILMGetStatus) WithFilterPath(v ...string) func(*ILMGetStatusRequest) {
	return func(r *ILMGetStatusRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f ILMGetStatus) WithHeader(h map[string]string) func(*ILMGetStatusRequest) {
	return func(r *ILMGetStatusRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f ILMGetStatus) WithOpaqueID(s string) func(*ILMGetStatusRequest) {
	return func(r *ILMGetStatusRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
