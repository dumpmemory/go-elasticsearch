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
)

func newMLPutDataFrameAnalyticsFunc(t Transport) MLPutDataFrameAnalytics {
	return func(id string, body io.Reader, o ...func(*MLPutDataFrameAnalyticsRequest)) (*Response, error) {
		var r = MLPutDataFrameAnalyticsRequest{ID: id, Body: body}
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

// MLPutDataFrameAnalytics - Instantiates a data frame analytics job.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/put-dfanalytics.html.
type MLPutDataFrameAnalytics func(id string, body io.Reader, o ...func(*MLPutDataFrameAnalyticsRequest)) (*Response, error)

// MLPutDataFrameAnalyticsRequest configures the ML Put Data Frame Analytics API request.
type MLPutDataFrameAnalyticsRequest struct {
	ID string

	Body io.Reader

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

// Do executes the request and returns response or error.
func (r MLPutDataFrameAnalyticsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "ml.put_data_frame_analytics")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "PUT"

	path.Grow(7 + 1 + len("_ml") + 1 + len("data_frame") + 1 + len("analytics") + 1 + len(r.ID))
	path.WriteString("http://")
	path.WriteString("/")
	path.WriteString("_ml")
	path.WriteString("/")
	path.WriteString("data_frame")
	path.WriteString("/")
	path.WriteString("analytics")
	path.WriteString("/")
	path.WriteString(r.ID)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.RecordPathPart(ctx, "id", r.ID)
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
		instrument.BeforeRequest(req, "ml.put_data_frame_analytics")
		if reader := instrument.RecordRequestBody(ctx, "ml.put_data_frame_analytics", r.Body); reader != nil {
			req.Body = reader
		}
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "ml.put_data_frame_analytics")
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
func (f MLPutDataFrameAnalytics) WithContext(v context.Context) func(*MLPutDataFrameAnalyticsRequest) {
	return func(r *MLPutDataFrameAnalyticsRequest) {
		r.ctx = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f MLPutDataFrameAnalytics) WithPretty() func(*MLPutDataFrameAnalyticsRequest) {
	return func(r *MLPutDataFrameAnalyticsRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f MLPutDataFrameAnalytics) WithHuman() func(*MLPutDataFrameAnalyticsRequest) {
	return func(r *MLPutDataFrameAnalyticsRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f MLPutDataFrameAnalytics) WithErrorTrace() func(*MLPutDataFrameAnalyticsRequest) {
	return func(r *MLPutDataFrameAnalyticsRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f MLPutDataFrameAnalytics) WithFilterPath(v ...string) func(*MLPutDataFrameAnalyticsRequest) {
	return func(r *MLPutDataFrameAnalyticsRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f MLPutDataFrameAnalytics) WithHeader(h map[string]string) func(*MLPutDataFrameAnalyticsRequest) {
	return func(r *MLPutDataFrameAnalyticsRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f MLPutDataFrameAnalytics) WithOpaqueID(s string) func(*MLPutDataFrameAnalyticsRequest) {
	return func(r *MLPutDataFrameAnalyticsRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
