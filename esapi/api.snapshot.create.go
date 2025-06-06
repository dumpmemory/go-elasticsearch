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

func newSnapshotCreateFunc(t Transport) SnapshotCreate {
	return func(repository string, snapshot string, o ...func(*SnapshotCreateRequest)) (*Response, error) {
		var r = SnapshotCreateRequest{Repository: repository, Snapshot: snapshot}
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

// SnapshotCreate creates a snapshot in a repository.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/master/modules-snapshots.html.
type SnapshotCreate func(repository string, snapshot string, o ...func(*SnapshotCreateRequest)) (*Response, error)

// SnapshotCreateRequest configures the Snapshot Create API request.
type SnapshotCreateRequest struct {
	Body io.Reader

	Repository string
	Snapshot   string

	MasterTimeout     time.Duration
	WaitForCompletion *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

// Do executes the request and returns response or error.
func (r SnapshotCreateRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "snapshot.create")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "PUT"

	path.Grow(7 + 1 + len("_snapshot") + 1 + len(r.Repository) + 1 + len(r.Snapshot))
	path.WriteString("http://")
	path.WriteString("/")
	path.WriteString("_snapshot")
	path.WriteString("/")
	path.WriteString(r.Repository)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.RecordPathPart(ctx, "repository", r.Repository)
	}
	path.WriteString("/")
	path.WriteString(r.Snapshot)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.RecordPathPart(ctx, "snapshot", r.Snapshot)
	}

	params = make(map[string]string)

	if r.MasterTimeout != 0 {
		params["master_timeout"] = formatDuration(r.MasterTimeout)
	}

	if r.WaitForCompletion != nil {
		params["wait_for_completion"] = strconv.FormatBool(*r.WaitForCompletion)
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
		instrument.BeforeRequest(req, "snapshot.create")
		if reader := instrument.RecordRequestBody(ctx, "snapshot.create", r.Body); reader != nil {
			req.Body = reader
		}
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "snapshot.create")
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
func (f SnapshotCreate) WithContext(v context.Context) func(*SnapshotCreateRequest) {
	return func(r *SnapshotCreateRequest) {
		r.ctx = v
	}
}

// WithBody - The snapshot definition.
func (f SnapshotCreate) WithBody(v io.Reader) func(*SnapshotCreateRequest) {
	return func(r *SnapshotCreateRequest) {
		r.Body = v
	}
}

// WithMasterTimeout - explicit operation timeout for connection to master node.
func (f SnapshotCreate) WithMasterTimeout(v time.Duration) func(*SnapshotCreateRequest) {
	return func(r *SnapshotCreateRequest) {
		r.MasterTimeout = v
	}
}

// WithWaitForCompletion - should this request wait until the operation has completed before returning.
func (f SnapshotCreate) WithWaitForCompletion(v bool) func(*SnapshotCreateRequest) {
	return func(r *SnapshotCreateRequest) {
		r.WaitForCompletion = &v
	}
}

// WithPretty makes the response body pretty-printed.
func (f SnapshotCreate) WithPretty() func(*SnapshotCreateRequest) {
	return func(r *SnapshotCreateRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f SnapshotCreate) WithHuman() func(*SnapshotCreateRequest) {
	return func(r *SnapshotCreateRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f SnapshotCreate) WithErrorTrace() func(*SnapshotCreateRequest) {
	return func(r *SnapshotCreateRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f SnapshotCreate) WithFilterPath(v ...string) func(*SnapshotCreateRequest) {
	return func(r *SnapshotCreateRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f SnapshotCreate) WithHeader(h map[string]string) func(*SnapshotCreateRequest) {
	return func(r *SnapshotCreateRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f SnapshotCreate) WithOpaqueID(s string) func(*SnapshotCreateRequest) {
	return func(r *SnapshotCreateRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
