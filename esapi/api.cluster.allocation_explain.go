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
	"strconv"
	"strings"
	"time"
)

func newClusterAllocationExplainFunc(t Transport) ClusterAllocationExplain {
	return func(o ...func(*ClusterAllocationExplainRequest)) (*Response, error) {
		var r = ClusterAllocationExplainRequest{}
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

// ClusterAllocationExplain provides explanations for shard allocations in the cluster.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/master/cluster-allocation-explain.html.
type ClusterAllocationExplain func(o ...func(*ClusterAllocationExplainRequest)) (*Response, error)

// ClusterAllocationExplainRequest configures the Cluster Allocation Explain API request.
type ClusterAllocationExplainRequest struct {
	Body io.Reader

	IncludeDiskInfo     *bool
	IncludeYesDecisions *bool
	MasterTimeout       time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

// Do executes the request and returns response or error.
func (r ClusterAllocationExplainRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "cluster.allocation_explain")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "POST"

	path.Grow(7 + len("/_cluster/allocation/explain"))
	path.WriteString("http://")
	path.WriteString("/_cluster/allocation/explain")

	params = make(map[string]string)

	if r.IncludeDiskInfo != nil {
		params["include_disk_info"] = strconv.FormatBool(*r.IncludeDiskInfo)
	}

	if r.IncludeYesDecisions != nil {
		params["include_yes_decisions"] = strconv.FormatBool(*r.IncludeYesDecisions)
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
		instrument.BeforeRequest(req, "cluster.allocation_explain")
		if reader := instrument.RecordRequestBody(ctx, "cluster.allocation_explain", r.Body); reader != nil {
			req.Body = reader
		}
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "cluster.allocation_explain")
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
func (f ClusterAllocationExplain) WithContext(v context.Context) func(*ClusterAllocationExplainRequest) {
	return func(r *ClusterAllocationExplainRequest) {
		r.ctx = v
	}
}

// WithBody - The index, shard, and primary flag to explain. Empty means 'explain a randomly-chosen unassigned shard'.
func (f ClusterAllocationExplain) WithBody(v io.Reader) func(*ClusterAllocationExplainRequest) {
	return func(r *ClusterAllocationExplainRequest) {
		r.Body = v
	}
}

// WithIncludeDiskInfo - return information about disk usage and shard sizes (default: false).
func (f ClusterAllocationExplain) WithIncludeDiskInfo(v bool) func(*ClusterAllocationExplainRequest) {
	return func(r *ClusterAllocationExplainRequest) {
		r.IncludeDiskInfo = &v
	}
}

// WithIncludeYesDecisions - return 'yes' decisions in explanation (default: false).
func (f ClusterAllocationExplain) WithIncludeYesDecisions(v bool) func(*ClusterAllocationExplainRequest) {
	return func(r *ClusterAllocationExplainRequest) {
		r.IncludeYesDecisions = &v
	}
}

// WithMasterTimeout - timeout for connection to master node.
func (f ClusterAllocationExplain) WithMasterTimeout(v time.Duration) func(*ClusterAllocationExplainRequest) {
	return func(r *ClusterAllocationExplainRequest) {
		r.MasterTimeout = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f ClusterAllocationExplain) WithPretty() func(*ClusterAllocationExplainRequest) {
	return func(r *ClusterAllocationExplainRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f ClusterAllocationExplain) WithHuman() func(*ClusterAllocationExplainRequest) {
	return func(r *ClusterAllocationExplainRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f ClusterAllocationExplain) WithErrorTrace() func(*ClusterAllocationExplainRequest) {
	return func(r *ClusterAllocationExplainRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f ClusterAllocationExplain) WithFilterPath(v ...string) func(*ClusterAllocationExplainRequest) {
	return func(r *ClusterAllocationExplainRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f ClusterAllocationExplain) WithHeader(h map[string]string) func(*ClusterAllocationExplainRequest) {
	return func(r *ClusterAllocationExplainRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f ClusterAllocationExplain) WithOpaqueID(s string) func(*ClusterAllocationExplainRequest) {
	return func(r *ClusterAllocationExplainRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
