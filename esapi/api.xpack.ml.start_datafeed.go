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
	"io"
	"net/http"
	"strings"
	"time"
)

func newMLStartDatafeedFunc(t Transport) MLStartDatafeed {
	return func(datafeed_id string, o ...func(*MLStartDatafeedRequest)) (*Response, error) {
		var r = MLStartDatafeedRequest{DatafeedID: datafeed_id}
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

// MLStartDatafeed - Starts one or more datafeeds.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-start-datafeed.html.
type MLStartDatafeed func(datafeed_id string, o ...func(*MLStartDatafeedRequest)) (*Response, error)

// MLStartDatafeedRequest configures the ML Start Datafeed API request.
type MLStartDatafeedRequest struct {
	Body io.Reader

	DatafeedID string

	End     string
	Start   string
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
func (r MLStartDatafeedRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "ml.start_datafeed")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "POST"

	path.Grow(7 + 1 + len("_ml") + 1 + len("datafeeds") + 1 + len(r.DatafeedID) + 1 + len("_start"))
	path.WriteString("http://")
	path.WriteString("/")
	path.WriteString("_ml")
	path.WriteString("/")
	path.WriteString("datafeeds")
	path.WriteString("/")
	path.WriteString(r.DatafeedID)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.RecordPathPart(ctx, "datafeed_id", r.DatafeedID)
	}
	path.WriteString("/")
	path.WriteString("_start")

	params = make(map[string]string)

	if r.End != "" {
		params["end"] = r.End
	}

	if r.Start != "" {
		params["start"] = r.Start
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
		instrument.BeforeRequest(req, "ml.start_datafeed")
		if reader := instrument.RecordRequestBody(ctx, "ml.start_datafeed", r.Body); reader != nil {
			req.Body = reader
		}
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "ml.start_datafeed")
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
func (f MLStartDatafeed) WithContext(v context.Context) func(*MLStartDatafeedRequest) {
	return func(r *MLStartDatafeedRequest) {
		r.ctx = v
	}
}

// WithBody - The start datafeed parameters.
func (f MLStartDatafeed) WithBody(v io.Reader) func(*MLStartDatafeedRequest) {
	return func(r *MLStartDatafeedRequest) {
		r.Body = v
	}
}

// WithEnd - the end time when the datafeed should stop. when not set, the datafeed continues in real time.
func (f MLStartDatafeed) WithEnd(v string) func(*MLStartDatafeedRequest) {
	return func(r *MLStartDatafeedRequest) {
		r.End = v
	}
}

// WithStart - the start time from where the datafeed should begin.
func (f MLStartDatafeed) WithStart(v string) func(*MLStartDatafeedRequest) {
	return func(r *MLStartDatafeedRequest) {
		r.Start = v
	}
}

// WithTimeout - controls the time to wait until a datafeed has started. default to 20 seconds.
func (f MLStartDatafeed) WithTimeout(v time.Duration) func(*MLStartDatafeedRequest) {
	return func(r *MLStartDatafeedRequest) {
		r.Timeout = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f MLStartDatafeed) WithPretty() func(*MLStartDatafeedRequest) {
	return func(r *MLStartDatafeedRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f MLStartDatafeed) WithHuman() func(*MLStartDatafeedRequest) {
	return func(r *MLStartDatafeedRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f MLStartDatafeed) WithErrorTrace() func(*MLStartDatafeedRequest) {
	return func(r *MLStartDatafeedRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f MLStartDatafeed) WithFilterPath(v ...string) func(*MLStartDatafeedRequest) {
	return func(r *MLStartDatafeedRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f MLStartDatafeed) WithHeader(h map[string]string) func(*MLStartDatafeedRequest) {
	return func(r *MLStartDatafeedRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f MLStartDatafeed) WithOpaqueID(s string) func(*MLStartDatafeedRequest) {
	return func(r *MLStartDatafeedRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
