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

func newSecurityHasPrivilegesFunc(t Transport) SecurityHasPrivileges {
	return func(body io.Reader, o ...func(*SecurityHasPrivilegesRequest)) (*Response, error) {
		var r = SecurityHasPrivilegesRequest{Body: body}
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

// SecurityHasPrivileges - Determines whether the specified user has a specified list of privileges.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-has-privileges.html.
type SecurityHasPrivileges func(body io.Reader, o ...func(*SecurityHasPrivilegesRequest)) (*Response, error)

// SecurityHasPrivilegesRequest configures the Security Has Privileges API request.
type SecurityHasPrivilegesRequest struct {
	Body io.Reader

	User string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

// Do executes the request and returns response or error.
func (r SecurityHasPrivilegesRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "security.has_privileges")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "POST"

	path.Grow(7 + 1 + len("_security") + 1 + len("user") + 1 + len(r.User) + 1 + len("_has_privileges"))
	path.WriteString("http://")
	path.WriteString("/")
	path.WriteString("_security")
	path.WriteString("/")
	path.WriteString("user")
	if r.User != "" {
		path.WriteString("/")
		path.WriteString(r.User)
		if instrument, ok := r.Instrument.(Instrumentation); ok {
			instrument.RecordPathPart(ctx, "user", r.User)
		}
	}
	path.WriteString("/")
	path.WriteString("_has_privileges")

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
		instrument.BeforeRequest(req, "security.has_privileges")
		if reader := instrument.RecordRequestBody(ctx, "security.has_privileges", r.Body); reader != nil {
			req.Body = reader
		}
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "security.has_privileges")
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
func (f SecurityHasPrivileges) WithContext(v context.Context) func(*SecurityHasPrivilegesRequest) {
	return func(r *SecurityHasPrivilegesRequest) {
		r.ctx = v
	}
}

// WithUser - username.
func (f SecurityHasPrivileges) WithUser(v string) func(*SecurityHasPrivilegesRequest) {
	return func(r *SecurityHasPrivilegesRequest) {
		r.User = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f SecurityHasPrivileges) WithPretty() func(*SecurityHasPrivilegesRequest) {
	return func(r *SecurityHasPrivilegesRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f SecurityHasPrivileges) WithHuman() func(*SecurityHasPrivilegesRequest) {
	return func(r *SecurityHasPrivilegesRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f SecurityHasPrivileges) WithErrorTrace() func(*SecurityHasPrivilegesRequest) {
	return func(r *SecurityHasPrivilegesRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f SecurityHasPrivileges) WithFilterPath(v ...string) func(*SecurityHasPrivilegesRequest) {
	return func(r *SecurityHasPrivilegesRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f SecurityHasPrivileges) WithHeader(h map[string]string) func(*SecurityHasPrivilegesRequest) {
	return func(r *SecurityHasPrivilegesRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f SecurityHasPrivileges) WithOpaqueID(s string) func(*SecurityHasPrivilegesRequest) {
	return func(r *SecurityHasPrivilegesRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
