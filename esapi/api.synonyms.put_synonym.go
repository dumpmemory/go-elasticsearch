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
)

func newSynonymsPutSynonymFunc(t Transport) SynonymsPutSynonym {
	return func(id string, body io.Reader, o ...func(*SynonymsPutSynonymRequest)) (*Response, error) {
		var r = SynonymsPutSynonymRequest{DocumentID: id, Body: body}
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

// SynonymsPutSynonym creates or updates a synonyms set
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/master/put-synonyms-set.html.
type SynonymsPutSynonym func(id string, body io.Reader, o ...func(*SynonymsPutSynonymRequest)) (*Response, error)

// SynonymsPutSynonymRequest configures the Synonyms Put Synonym API request.
type SynonymsPutSynonymRequest struct {
	DocumentID string

	Body io.Reader

	Refresh *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

// Do executes the request and returns response or error.
func (r SynonymsPutSynonymRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "synonyms.put_synonym")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "PUT"

	path.Grow(7 + 1 + len("_synonyms") + 1 + len(r.DocumentID))
	path.WriteString("http://")
	path.WriteString("/")
	path.WriteString("_synonyms")
	path.WriteString("/")
	path.WriteString(r.DocumentID)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.RecordPathPart(ctx, "id", r.DocumentID)
	}

	params = make(map[string]string)

	if r.Refresh != nil {
		params["refresh"] = strconv.FormatBool(*r.Refresh)
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
		instrument.BeforeRequest(req, "synonyms.put_synonym")
		if reader := instrument.RecordRequestBody(ctx, "synonyms.put_synonym", r.Body); reader != nil {
			req.Body = reader
		}
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "synonyms.put_synonym")
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
func (f SynonymsPutSynonym) WithContext(v context.Context) func(*SynonymsPutSynonymRequest) {
	return func(r *SynonymsPutSynonymRequest) {
		r.ctx = v
	}
}

// WithRefresh - refresh search analyzers to update synonyms.
func (f SynonymsPutSynonym) WithRefresh(v bool) func(*SynonymsPutSynonymRequest) {
	return func(r *SynonymsPutSynonymRequest) {
		r.Refresh = &v
	}
}

// WithPretty makes the response body pretty-printed.
func (f SynonymsPutSynonym) WithPretty() func(*SynonymsPutSynonymRequest) {
	return func(r *SynonymsPutSynonymRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f SynonymsPutSynonym) WithHuman() func(*SynonymsPutSynonymRequest) {
	return func(r *SynonymsPutSynonymRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f SynonymsPutSynonym) WithErrorTrace() func(*SynonymsPutSynonymRequest) {
	return func(r *SynonymsPutSynonymRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f SynonymsPutSynonym) WithFilterPath(v ...string) func(*SynonymsPutSynonymRequest) {
	return func(r *SynonymsPutSynonymRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f SynonymsPutSynonym) WithHeader(h map[string]string) func(*SynonymsPutSynonymRequest) {
	return func(r *SynonymsPutSynonymRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f SynonymsPutSynonym) WithOpaqueID(s string) func(*SynonymsPutSynonymRequest) {
	return func(r *SynonymsPutSynonymRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
