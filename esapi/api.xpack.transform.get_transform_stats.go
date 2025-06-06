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

func newTransformGetTransformStatsFunc(t Transport) TransformGetTransformStats {
	return func(transform_id string, o ...func(*TransformGetTransformStatsRequest)) (*Response, error) {
		var r = TransformGetTransformStatsRequest{TransformID: transform_id}
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

// TransformGetTransformStats - Retrieves usage information for transforms.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/get-transform-stats.html.
type TransformGetTransformStats func(transform_id string, o ...func(*TransformGetTransformStatsRequest)) (*Response, error)

// TransformGetTransformStatsRequest configures the Transform Get Transform Stats API request.
type TransformGetTransformStatsRequest struct {
	TransformID string

	AllowNoMatch *bool
	From         *int
	Size         *int
	Timeout      time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

// Do executes the request and returns response or error.
func (r TransformGetTransformStatsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "transform.get_transform_stats")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "GET"

	path.Grow(7 + 1 + len("_transform") + 1 + len(r.TransformID) + 1 + len("_stats"))
	path.WriteString("http://")
	path.WriteString("/")
	path.WriteString("_transform")
	path.WriteString("/")
	path.WriteString(r.TransformID)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.RecordPathPart(ctx, "transform_id", r.TransformID)
	}
	path.WriteString("/")
	path.WriteString("_stats")

	params = make(map[string]string)

	if r.AllowNoMatch != nil {
		params["allow_no_match"] = strconv.FormatBool(*r.AllowNoMatch)
	}

	if r.From != nil {
		params["from"] = strconv.FormatInt(int64(*r.From), 10)
	}

	if r.Size != nil {
		params["size"] = strconv.FormatInt(int64(*r.Size), 10)
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
		instrument.BeforeRequest(req, "transform.get_transform_stats")
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "transform.get_transform_stats")
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
func (f TransformGetTransformStats) WithContext(v context.Context) func(*TransformGetTransformStatsRequest) {
	return func(r *TransformGetTransformStatsRequest) {
		r.ctx = v
	}
}

// WithAllowNoMatch - whether to ignore if a wildcard expression matches no transforms. (this includes `_all` string or when no transforms have been specified).
func (f TransformGetTransformStats) WithAllowNoMatch(v bool) func(*TransformGetTransformStatsRequest) {
	return func(r *TransformGetTransformStatsRequest) {
		r.AllowNoMatch = &v
	}
}

// WithFrom - skips a number of transform stats, defaults to 0.
func (f TransformGetTransformStats) WithFrom(v int) func(*TransformGetTransformStatsRequest) {
	return func(r *TransformGetTransformStatsRequest) {
		r.From = &v
	}
}

// WithSize - specifies a max number of transform stats to get, defaults to 100.
func (f TransformGetTransformStats) WithSize(v int) func(*TransformGetTransformStatsRequest) {
	return func(r *TransformGetTransformStatsRequest) {
		r.Size = &v
	}
}

// WithTimeout - controls the time to wait for the stats.
func (f TransformGetTransformStats) WithTimeout(v time.Duration) func(*TransformGetTransformStatsRequest) {
	return func(r *TransformGetTransformStatsRequest) {
		r.Timeout = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f TransformGetTransformStats) WithPretty() func(*TransformGetTransformStatsRequest) {
	return func(r *TransformGetTransformStatsRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f TransformGetTransformStats) WithHuman() func(*TransformGetTransformStatsRequest) {
	return func(r *TransformGetTransformStatsRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f TransformGetTransformStats) WithErrorTrace() func(*TransformGetTransformStatsRequest) {
	return func(r *TransformGetTransformStatsRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f TransformGetTransformStats) WithFilterPath(v ...string) func(*TransformGetTransformStatsRequest) {
	return func(r *TransformGetTransformStatsRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f TransformGetTransformStats) WithHeader(h map[string]string) func(*TransformGetTransformStatsRequest) {
	return func(r *TransformGetTransformStatsRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f TransformGetTransformStats) WithOpaqueID(s string) func(*TransformGetTransformStatsRequest) {
	return func(r *TransformGetTransformStatsRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
