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
)

func newSecurityGetPrivilegesFunc(t Transport) SecurityGetPrivileges {
	return func(o ...func(*SecurityGetPrivilegesRequest)) (*Response, error) {
		var r = SecurityGetPrivilegesRequest{}
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

// SecurityGetPrivileges - Retrieves application privileges.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-get-privileges.html.
type SecurityGetPrivileges func(o ...func(*SecurityGetPrivilegesRequest)) (*Response, error)

// SecurityGetPrivilegesRequest configures the Security Get Privileges API request.
type SecurityGetPrivilegesRequest struct {
	Application string
	Name        string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

// Do executes the request and returns response or error.
func (r SecurityGetPrivilegesRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "security.get_privileges")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "GET"

	path.Grow(7 + 1 + len("_security") + 1 + len("privilege") + 1 + len(r.Application) + 1 + len(r.Name))
	path.WriteString("http://")
	path.WriteString("/")
	path.WriteString("_security")
	path.WriteString("/")
	path.WriteString("privilege")
	if r.Application != "" {
		path.WriteString("/")
		path.WriteString(r.Application)
		if instrument, ok := r.Instrument.(Instrumentation); ok {
			instrument.RecordPathPart(ctx, "application", r.Application)
		}
	}
	if r.Name != "" {
		path.WriteString("/")
		path.WriteString(r.Name)
		if instrument, ok := r.Instrument.(Instrumentation); ok {
			instrument.RecordPathPart(ctx, "name", r.Name)
		}
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
		instrument.BeforeRequest(req, "security.get_privileges")
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "security.get_privileges")
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
func (f SecurityGetPrivileges) WithContext(v context.Context) func(*SecurityGetPrivilegesRequest) {
	return func(r *SecurityGetPrivilegesRequest) {
		r.ctx = v
	}
}

// WithApplication - application name.
func (f SecurityGetPrivileges) WithApplication(v string) func(*SecurityGetPrivilegesRequest) {
	return func(r *SecurityGetPrivilegesRequest) {
		r.Application = v
	}
}

// WithName - privilege name.
func (f SecurityGetPrivileges) WithName(v string) func(*SecurityGetPrivilegesRequest) {
	return func(r *SecurityGetPrivilegesRequest) {
		r.Name = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f SecurityGetPrivileges) WithPretty() func(*SecurityGetPrivilegesRequest) {
	return func(r *SecurityGetPrivilegesRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f SecurityGetPrivileges) WithHuman() func(*SecurityGetPrivilegesRequest) {
	return func(r *SecurityGetPrivilegesRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f SecurityGetPrivileges) WithErrorTrace() func(*SecurityGetPrivilegesRequest) {
	return func(r *SecurityGetPrivilegesRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f SecurityGetPrivileges) WithFilterPath(v ...string) func(*SecurityGetPrivilegesRequest) {
	return func(r *SecurityGetPrivilegesRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f SecurityGetPrivileges) WithHeader(h map[string]string) func(*SecurityGetPrivilegesRequest) {
	return func(r *SecurityGetPrivilegesRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f SecurityGetPrivileges) WithOpaqueID(s string) func(*SecurityGetPrivilegesRequest) {
	return func(r *SecurityGetPrivilegesRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
