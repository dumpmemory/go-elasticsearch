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

func newClusterRerouteFunc(t Transport) ClusterReroute {
	return func(o ...func(*ClusterRerouteRequest)) (*Response, error) {
		var r = ClusterRerouteRequest{}
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

// ClusterReroute allows to manually change the allocation of individual shards in the cluster.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/master/cluster-reroute.html.
type ClusterReroute func(o ...func(*ClusterRerouteRequest)) (*Response, error)

// ClusterRerouteRequest configures the Cluster Reroute API request.
type ClusterRerouteRequest struct {
	Body io.Reader

	DryRun        *bool
	Explain       *bool
	MasterTimeout time.Duration
	Metric        []string
	RetryFailed   *bool
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
func (r ClusterRerouteRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "cluster.reroute")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "POST"

	path.Grow(7 + len("/_cluster/reroute"))
	path.WriteString("http://")
	path.WriteString("/_cluster/reroute")

	params = make(map[string]string)

	if r.DryRun != nil {
		params["dry_run"] = strconv.FormatBool(*r.DryRun)
	}

	if r.Explain != nil {
		params["explain"] = strconv.FormatBool(*r.Explain)
	}

	if r.MasterTimeout != 0 {
		params["master_timeout"] = formatDuration(r.MasterTimeout)
	}

	if len(r.Metric) > 0 {
		params["metric"] = strings.Join(r.Metric, ",")
	}

	if r.RetryFailed != nil {
		params["retry_failed"] = strconv.FormatBool(*r.RetryFailed)
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
		instrument.BeforeRequest(req, "cluster.reroute")
		if reader := instrument.RecordRequestBody(ctx, "cluster.reroute", r.Body); reader != nil {
			req.Body = reader
		}
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "cluster.reroute")
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
func (f ClusterReroute) WithContext(v context.Context) func(*ClusterRerouteRequest) {
	return func(r *ClusterRerouteRequest) {
		r.ctx = v
	}
}

// WithBody - The definition of `commands` to perform (`move`, `cancel`, `allocate`).
func (f ClusterReroute) WithBody(v io.Reader) func(*ClusterRerouteRequest) {
	return func(r *ClusterRerouteRequest) {
		r.Body = v
	}
}

// WithDryRun - simulate the operation only and return the resulting state.
func (f ClusterReroute) WithDryRun(v bool) func(*ClusterRerouteRequest) {
	return func(r *ClusterRerouteRequest) {
		r.DryRun = &v
	}
}

// WithExplain - return an explanation of why the commands can or cannot be executed.
func (f ClusterReroute) WithExplain(v bool) func(*ClusterRerouteRequest) {
	return func(r *ClusterRerouteRequest) {
		r.Explain = &v
	}
}

// WithMasterTimeout - explicit operation timeout for connection to master node.
func (f ClusterReroute) WithMasterTimeout(v time.Duration) func(*ClusterRerouteRequest) {
	return func(r *ClusterRerouteRequest) {
		r.MasterTimeout = v
	}
}

// WithMetric - limit the information returned to the specified metrics. defaults to all but metadata.
func (f ClusterReroute) WithMetric(v ...string) func(*ClusterRerouteRequest) {
	return func(r *ClusterRerouteRequest) {
		r.Metric = v
	}
}

// WithRetryFailed - retries allocation of shards that are blocked due to too many subsequent allocation failures.
func (f ClusterReroute) WithRetryFailed(v bool) func(*ClusterRerouteRequest) {
	return func(r *ClusterRerouteRequest) {
		r.RetryFailed = &v
	}
}

// WithTimeout - explicit operation timeout.
func (f ClusterReroute) WithTimeout(v time.Duration) func(*ClusterRerouteRequest) {
	return func(r *ClusterRerouteRequest) {
		r.Timeout = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f ClusterReroute) WithPretty() func(*ClusterRerouteRequest) {
	return func(r *ClusterRerouteRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f ClusterReroute) WithHuman() func(*ClusterRerouteRequest) {
	return func(r *ClusterRerouteRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f ClusterReroute) WithErrorTrace() func(*ClusterRerouteRequest) {
	return func(r *ClusterRerouteRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f ClusterReroute) WithFilterPath(v ...string) func(*ClusterRerouteRequest) {
	return func(r *ClusterRerouteRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f ClusterReroute) WithHeader(h map[string]string) func(*ClusterRerouteRequest) {
	return func(r *ClusterRerouteRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f ClusterReroute) WithOpaqueID(s string) func(*ClusterRerouteRequest) {
	return func(r *ClusterRerouteRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
