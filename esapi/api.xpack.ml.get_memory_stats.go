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
	"time"
)

func newMLGetMemoryStatsFunc(t Transport) MLGetMemoryStats {
	return func(o ...func(*MLGetMemoryStatsRequest)) (*Response, error) {
		var r = MLGetMemoryStatsRequest{}
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

// MLGetMemoryStats - Returns information on how ML is using memory.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/get-ml-memory.html.
type MLGetMemoryStats func(o ...func(*MLGetMemoryStatsRequest)) (*Response, error)

// MLGetMemoryStatsRequest configures the ML Get Memory Stats API request.
type MLGetMemoryStatsRequest struct {
	NodeID string

	MasterTimeout time.Duration
	Timeout       time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

// Do executes the request and returns response or error.
func (r MLGetMemoryStatsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "ml.get_memory_stats")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "GET"

	path.Grow(7 + 1 + len("_ml") + 1 + len("memory") + 1 + len(r.NodeID) + 1 + len("_stats"))
	path.WriteString("http://")
	path.WriteString("/")
	path.WriteString("_ml")
	path.WriteString("/")
	path.WriteString("memory")
	if r.NodeID != "" {
		path.WriteString("/")
		path.WriteString(r.NodeID)
		if instrument, ok := r.Instrument.(Instrumentation); ok {
			instrument.RecordPathPart(ctx, "node_id", r.NodeID)
		}
	}
	path.WriteString("/")
	path.WriteString("_stats")

	params = make(map[string]string)

	if r.MasterTimeout != 0 {
		params["master_timeout"] = formatDuration(r.MasterTimeout)
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
		instrument.BeforeRequest(req, "ml.get_memory_stats")
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "ml.get_memory_stats")
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
func (f MLGetMemoryStats) WithContext(v context.Context) func(*MLGetMemoryStatsRequest) {
	return func(r *MLGetMemoryStatsRequest) {
		r.ctx = v
	}
}

// WithNodeID - specifies the node or nodes to retrieve stats for..
func (f MLGetMemoryStats) WithNodeID(v string) func(*MLGetMemoryStatsRequest) {
	return func(r *MLGetMemoryStatsRequest) {
		r.NodeID = v
	}
}

// WithMasterTimeout - explicit operation timeout for connection to master node.
func (f MLGetMemoryStats) WithMasterTimeout(v time.Duration) func(*MLGetMemoryStatsRequest) {
	return func(r *MLGetMemoryStatsRequest) {
		r.MasterTimeout = v
	}
}

// WithTimeout - explicit operation timeout.
func (f MLGetMemoryStats) WithTimeout(v time.Duration) func(*MLGetMemoryStatsRequest) {
	return func(r *MLGetMemoryStatsRequest) {
		r.Timeout = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f MLGetMemoryStats) WithPretty() func(*MLGetMemoryStatsRequest) {
	return func(r *MLGetMemoryStatsRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f MLGetMemoryStats) WithHuman() func(*MLGetMemoryStatsRequest) {
	return func(r *MLGetMemoryStatsRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f MLGetMemoryStats) WithErrorTrace() func(*MLGetMemoryStatsRequest) {
	return func(r *MLGetMemoryStatsRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f MLGetMemoryStats) WithFilterPath(v ...string) func(*MLGetMemoryStatsRequest) {
	return func(r *MLGetMemoryStatsRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f MLGetMemoryStats) WithHeader(h map[string]string) func(*MLGetMemoryStatsRequest) {
	return func(r *MLGetMemoryStatsRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f MLGetMemoryStats) WithOpaqueID(s string) func(*MLGetMemoryStatsRequest) {
	return func(r *MLGetMemoryStatsRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
