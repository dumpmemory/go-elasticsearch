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

func newSnapshotCloneFunc(t Transport) SnapshotClone {
	return func(repository string, snapshot string, body io.Reader, target_snapshot string, o ...func(*SnapshotCloneRequest)) (*Response, error) {
		var r = SnapshotCloneRequest{Repository: repository, Snapshot: snapshot, Body: body, TargetSnapshot: target_snapshot}
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

// SnapshotClone clones indices from one snapshot into another snapshot in the same repository.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/master/modules-snapshots.html.
type SnapshotClone func(repository string, snapshot string, body io.Reader, target_snapshot string, o ...func(*SnapshotCloneRequest)) (*Response, error)

// SnapshotCloneRequest configures the Snapshot Clone API request.
type SnapshotCloneRequest struct {
	Body io.Reader

	Repository     string
	Snapshot       string
	TargetSnapshot string

	MasterTimeout time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

// Do executes the request and returns response or error.
func (r SnapshotCloneRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "snapshot.clone")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "PUT"

	path.Grow(7 + 1 + len("_snapshot") + 1 + len(r.Repository) + 1 + len(r.Snapshot) + 1 + len("_clone") + 1 + len(r.TargetSnapshot))
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
	path.WriteString("/")
	path.WriteString("_clone")
	path.WriteString("/")
	path.WriteString(r.TargetSnapshot)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.RecordPathPart(ctx, "target_snapshot", r.TargetSnapshot)
	}

	params = make(map[string]string)

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
		instrument.BeforeRequest(req, "snapshot.clone")
		if reader := instrument.RecordRequestBody(ctx, "snapshot.clone", r.Body); reader != nil {
			req.Body = reader
		}
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "snapshot.clone")
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
func (f SnapshotClone) WithContext(v context.Context) func(*SnapshotCloneRequest) {
	return func(r *SnapshotCloneRequest) {
		r.ctx = v
	}
}

// WithMasterTimeout - explicit operation timeout for connection to master node.
func (f SnapshotClone) WithMasterTimeout(v time.Duration) func(*SnapshotCloneRequest) {
	return func(r *SnapshotCloneRequest) {
		r.MasterTimeout = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f SnapshotClone) WithPretty() func(*SnapshotCloneRequest) {
	return func(r *SnapshotCloneRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f SnapshotClone) WithHuman() func(*SnapshotCloneRequest) {
	return func(r *SnapshotCloneRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f SnapshotClone) WithErrorTrace() func(*SnapshotCloneRequest) {
	return func(r *SnapshotCloneRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f SnapshotClone) WithFilterPath(v ...string) func(*SnapshotCloneRequest) {
	return func(r *SnapshotCloneRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f SnapshotClone) WithHeader(h map[string]string) func(*SnapshotCloneRequest) {
	return func(r *SnapshotCloneRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f SnapshotClone) WithOpaqueID(s string) func(*SnapshotCloneRequest) {
	return func(r *SnapshotCloneRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
