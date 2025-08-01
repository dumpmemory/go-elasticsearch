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

func newCreateFunc(t Transport) Create {
	return func(index string, id string, body io.Reader, o ...func(*CreateRequest)) (*Response, error) {
		var r = CreateRequest{Index: index, DocumentID: id, Body: body}
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

// Create creates a new document in the index.
//
// Returns a 409 response when a document with a same ID already exists in the index.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/master/docs-index_.html.
type Create func(index string, id string, body io.Reader, o ...func(*CreateRequest)) (*Response, error)

// CreateRequest configures the Create API request.
type CreateRequest struct {
	Index      string
	DocumentID string

	Body io.Reader

	IncludeSourceOnError *bool
	Pipeline             string
	Refresh              string
	RequireAlias         *bool
	RequireDataStream    *bool
	Routing              string
	Timeout              time.Duration
	Version              *int
	VersionType          string
	WaitForActiveShards  string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

// Do executes the request and returns response or error.
func (r CreateRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "create")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "PUT"

	path.Grow(7 + 1 + len(r.Index) + 1 + len("_create") + 1 + len(r.DocumentID))
	path.WriteString("http://")
	path.WriteString("/")
	path.WriteString(r.Index)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.RecordPathPart(ctx, "index", r.Index)
	}
	path.WriteString("/")
	path.WriteString("_create")
	path.WriteString("/")
	path.WriteString(r.DocumentID)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.RecordPathPart(ctx, "id", r.DocumentID)
	}

	params = make(map[string]string)

	if r.IncludeSourceOnError != nil {
		params["include_source_on_error"] = strconv.FormatBool(*r.IncludeSourceOnError)
	}

	if r.Pipeline != "" {
		params["pipeline"] = r.Pipeline
	}

	if r.Refresh != "" {
		params["refresh"] = r.Refresh
	}

	if r.RequireAlias != nil {
		params["require_alias"] = strconv.FormatBool(*r.RequireAlias)
	}

	if r.RequireDataStream != nil {
		params["require_data_stream"] = strconv.FormatBool(*r.RequireDataStream)
	}

	if r.Routing != "" {
		params["routing"] = r.Routing
	}

	if r.Timeout != 0 {
		params["timeout"] = formatDuration(r.Timeout)
	}

	if r.Version != nil {
		params["version"] = strconv.FormatInt(int64(*r.Version), 10)
	}

	if r.VersionType != "" {
		params["version_type"] = r.VersionType
	}

	if r.WaitForActiveShards != "" {
		params["wait_for_active_shards"] = r.WaitForActiveShards
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
		instrument.BeforeRequest(req, "create")
		if reader := instrument.RecordRequestBody(ctx, "create", r.Body); reader != nil {
			req.Body = reader
		}
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "create")
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
func (f Create) WithContext(v context.Context) func(*CreateRequest) {
	return func(r *CreateRequest) {
		r.ctx = v
	}
}

// WithIncludeSourceOnError - true or false if to include the document source in the error message in case of parsing errors. defaults to true..
func (f Create) WithIncludeSourceOnError(v bool) func(*CreateRequest) {
	return func(r *CreateRequest) {
		r.IncludeSourceOnError = &v
	}
}

// WithPipeline - the pipeline ID to preprocess incoming documents with.
func (f Create) WithPipeline(v string) func(*CreateRequest) {
	return func(r *CreateRequest) {
		r.Pipeline = v
	}
}

// WithRefresh - if `true` then refresh the affected shards to make this operation visible to search, if `wait_for` then wait for a refresh to make this operation visible to search, if `false` (the default) then do nothing with refreshes..
func (f Create) WithRefresh(v string) func(*CreateRequest) {
	return func(r *CreateRequest) {
		r.Refresh = v
	}
}

// WithRequireAlias - when true, requires destination to be an alias. default is false.
func (f Create) WithRequireAlias(v bool) func(*CreateRequest) {
	return func(r *CreateRequest) {
		r.RequireAlias = &v
	}
}

// WithRequireDataStream - when true, requires destination to be a data stream (existing or to be created). default is false.
func (f Create) WithRequireDataStream(v bool) func(*CreateRequest) {
	return func(r *CreateRequest) {
		r.RequireDataStream = &v
	}
}

// WithRouting - specific routing value.
func (f Create) WithRouting(v string) func(*CreateRequest) {
	return func(r *CreateRequest) {
		r.Routing = v
	}
}

// WithTimeout - explicit operation timeout.
func (f Create) WithTimeout(v time.Duration) func(*CreateRequest) {
	return func(r *CreateRequest) {
		r.Timeout = v
	}
}

// WithVersion - explicit version number for concurrency control.
func (f Create) WithVersion(v int) func(*CreateRequest) {
	return func(r *CreateRequest) {
		r.Version = &v
	}
}

// WithVersionType - specific version type.
func (f Create) WithVersionType(v string) func(*CreateRequest) {
	return func(r *CreateRequest) {
		r.VersionType = v
	}
}

// WithWaitForActiveShards - sets the number of shard copies that must be active before proceeding with the index operation. defaults to 1, meaning the primary shard only. set to `all` for all shard copies, otherwise set to any non-negative value less than or equal to the total number of copies for the shard (number of replicas + 1).
func (f Create) WithWaitForActiveShards(v string) func(*CreateRequest) {
	return func(r *CreateRequest) {
		r.WaitForActiveShards = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f Create) WithPretty() func(*CreateRequest) {
	return func(r *CreateRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f Create) WithHuman() func(*CreateRequest) {
	return func(r *CreateRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f Create) WithErrorTrace() func(*CreateRequest) {
	return func(r *CreateRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f Create) WithFilterPath(v ...string) func(*CreateRequest) {
	return func(r *CreateRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f Create) WithHeader(h map[string]string) func(*CreateRequest) {
	return func(r *CreateRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f Create) WithOpaqueID(s string) func(*CreateRequest) {
	return func(r *CreateRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
