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

func newSecurityGetBuiltinPrivilegesFunc(t Transport) SecurityGetBuiltinPrivileges {
	return func(o ...func(*SecurityGetBuiltinPrivilegesRequest)) (*Response, error) {
		var r = SecurityGetBuiltinPrivilegesRequest{}
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

// SecurityGetBuiltinPrivileges - Retrieves the list of cluster privileges and index privileges that are available in this version of Elasticsearch.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-get-builtin-privileges.html.
type SecurityGetBuiltinPrivileges func(o ...func(*SecurityGetBuiltinPrivilegesRequest)) (*Response, error)

// SecurityGetBuiltinPrivilegesRequest configures the Security Get Builtin Privileges API request.
type SecurityGetBuiltinPrivilegesRequest struct {
	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

// Do executes the request and returns response or error.
func (r SecurityGetBuiltinPrivilegesRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "security.get_builtin_privileges")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "GET"

	path.Grow(7 + len("/_security/privilege/_builtin"))
	path.WriteString("http://")
	path.WriteString("/_security/privilege/_builtin")

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
		instrument.BeforeRequest(req, "security.get_builtin_privileges")
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "security.get_builtin_privileges")
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
func (f SecurityGetBuiltinPrivileges) WithContext(v context.Context) func(*SecurityGetBuiltinPrivilegesRequest) {
	return func(r *SecurityGetBuiltinPrivilegesRequest) {
		r.ctx = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f SecurityGetBuiltinPrivileges) WithPretty() func(*SecurityGetBuiltinPrivilegesRequest) {
	return func(r *SecurityGetBuiltinPrivilegesRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f SecurityGetBuiltinPrivileges) WithHuman() func(*SecurityGetBuiltinPrivilegesRequest) {
	return func(r *SecurityGetBuiltinPrivilegesRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f SecurityGetBuiltinPrivileges) WithErrorTrace() func(*SecurityGetBuiltinPrivilegesRequest) {
	return func(r *SecurityGetBuiltinPrivilegesRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f SecurityGetBuiltinPrivileges) WithFilterPath(v ...string) func(*SecurityGetBuiltinPrivilegesRequest) {
	return func(r *SecurityGetBuiltinPrivilegesRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f SecurityGetBuiltinPrivileges) WithHeader(h map[string]string) func(*SecurityGetBuiltinPrivilegesRequest) {
	return func(r *SecurityGetBuiltinPrivilegesRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f SecurityGetBuiltinPrivileges) WithOpaqueID(s string) func(*SecurityGetBuiltinPrivilegesRequest) {
	return func(r *SecurityGetBuiltinPrivilegesRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
