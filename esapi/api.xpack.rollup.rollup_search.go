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
	"strconv"
	"strings"
)

func newRollupRollupSearchFunc(t Transport) RollupRollupSearch {
	return func(index []string, body io.Reader, o ...func(*RollupRollupSearchRequest)) (*Response, error) {
		var r = RollupRollupSearchRequest{Index: index, Body: body}
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

// RollupRollupSearch - Enables searching rolled-up data using the standard query DSL.
//
// This API is experimental.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/master/rollup-search.html.
type RollupRollupSearch func(index []string, body io.Reader, o ...func(*RollupRollupSearchRequest)) (*Response, error)

// RollupRollupSearchRequest configures the Rollup Rollup Search API request.
type RollupRollupSearchRequest struct {
	Index []string

	Body io.Reader

	RestTotalHitsAsInt *bool
	TypedKeys          *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

// Do executes the request and returns response or error.
func (r RollupRollupSearchRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "rollup.rollup_search")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "POST"

	if len(r.Index) == 0 {
		return nil, errors.New("index is required and cannot be nil or empty")
	}

	path.Grow(7 + 1 + len(strings.Join(r.Index, ",")) + 1 + len("_rollup_search"))
	path.WriteString("http://")
	path.WriteString("/")
	path.WriteString(strings.Join(r.Index, ","))
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.RecordPathPart(ctx, "index", strings.Join(r.Index, ","))
	}
	path.WriteString("/")
	path.WriteString("_rollup_search")

	params = make(map[string]string)

	if r.RestTotalHitsAsInt != nil {
		params["rest_total_hits_as_int"] = strconv.FormatBool(*r.RestTotalHitsAsInt)
	}

	if r.TypedKeys != nil {
		params["typed_keys"] = strconv.FormatBool(*r.TypedKeys)
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
		instrument.BeforeRequest(req, "rollup.rollup_search")
		if reader := instrument.RecordRequestBody(ctx, "rollup.rollup_search", r.Body); reader != nil {
			req.Body = reader
		}
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "rollup.rollup_search")
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
func (f RollupRollupSearch) WithContext(v context.Context) func(*RollupRollupSearchRequest) {
	return func(r *RollupRollupSearchRequest) {
		r.ctx = v
	}
}

// WithRestTotalHitsAsInt - indicates whether hits.total should be rendered as an integer or an object in the rest search response.
func (f RollupRollupSearch) WithRestTotalHitsAsInt(v bool) func(*RollupRollupSearchRequest) {
	return func(r *RollupRollupSearchRequest) {
		r.RestTotalHitsAsInt = &v
	}
}

// WithTypedKeys - specify whether aggregation and suggester names should be prefixed by their respective types in the response.
func (f RollupRollupSearch) WithTypedKeys(v bool) func(*RollupRollupSearchRequest) {
	return func(r *RollupRollupSearchRequest) {
		r.TypedKeys = &v
	}
}

// WithPretty makes the response body pretty-printed.
func (f RollupRollupSearch) WithPretty() func(*RollupRollupSearchRequest) {
	return func(r *RollupRollupSearchRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f RollupRollupSearch) WithHuman() func(*RollupRollupSearchRequest) {
	return func(r *RollupRollupSearchRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f RollupRollupSearch) WithErrorTrace() func(*RollupRollupSearchRequest) {
	return func(r *RollupRollupSearchRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f RollupRollupSearch) WithFilterPath(v ...string) func(*RollupRollupSearchRequest) {
	return func(r *RollupRollupSearchRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f RollupRollupSearch) WithHeader(h map[string]string) func(*RollupRollupSearchRequest) {
	return func(r *RollupRollupSearchRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f RollupRollupSearch) WithOpaqueID(s string) func(*RollupRollupSearchRequest) {
	return func(r *RollupRollupSearchRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
