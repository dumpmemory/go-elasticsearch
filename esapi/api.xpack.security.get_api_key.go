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
	"strconv"
	"strings"
)

func newSecurityGetAPIKeyFunc(t Transport) SecurityGetAPIKey {
	return func(o ...func(*SecurityGetAPIKeyRequest)) (*Response, error) {
		var r = SecurityGetAPIKeyRequest{}
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

// SecurityGetAPIKey - Retrieves information for one or more API keys.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-get-api-key.html.
type SecurityGetAPIKey func(o ...func(*SecurityGetAPIKeyRequest)) (*Response, error)

// SecurityGetAPIKeyRequest configures the Security GetAPI Key API request.
type SecurityGetAPIKeyRequest struct {
	ActiveOnly     *bool
	ID             string
	Name           string
	Owner          *bool
	RealmName      string
	Username       string
	WithLimitedBy  *bool
	WithProfileUID *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

// Do executes the request and returns response or error.
func (r SecurityGetAPIKeyRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "security.get_api_key")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "GET"

	path.Grow(7 + len("/_security/api_key"))
	path.WriteString("http://")
	path.WriteString("/_security/api_key")

	params = make(map[string]string)

	if r.ActiveOnly != nil {
		params["active_only"] = strconv.FormatBool(*r.ActiveOnly)
	}

	if r.ID != "" {
		params["id"] = r.ID
	}

	if r.Name != "" {
		params["name"] = r.Name
	}

	if r.Owner != nil {
		params["owner"] = strconv.FormatBool(*r.Owner)
	}

	if r.RealmName != "" {
		params["realm_name"] = r.RealmName
	}

	if r.Username != "" {
		params["username"] = r.Username
	}

	if r.WithLimitedBy != nil {
		params["with_limited_by"] = strconv.FormatBool(*r.WithLimitedBy)
	}

	if r.WithProfileUID != nil {
		params["with_profile_uid"] = strconv.FormatBool(*r.WithProfileUID)
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
		instrument.BeforeRequest(req, "security.get_api_key")
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "security.get_api_key")
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
func (f SecurityGetAPIKey) WithContext(v context.Context) func(*SecurityGetAPIKeyRequest) {
	return func(r *SecurityGetAPIKeyRequest) {
		r.ctx = v
	}
}

// WithActiveOnly - flag to limit response to only active (not invalidated or expired) api keys.
func (f SecurityGetAPIKey) WithActiveOnly(v bool) func(*SecurityGetAPIKeyRequest) {
	return func(r *SecurityGetAPIKeyRequest) {
		r.ActiveOnly = &v
	}
}

// WithID - api key ID of the api key to be retrieved.
func (f SecurityGetAPIKey) WithID(v string) func(*SecurityGetAPIKeyRequest) {
	return func(r *SecurityGetAPIKeyRequest) {
		r.ID = v
	}
}

// WithName - api key name of the api key to be retrieved.
func (f SecurityGetAPIKey) WithName(v string) func(*SecurityGetAPIKeyRequest) {
	return func(r *SecurityGetAPIKeyRequest) {
		r.Name = v
	}
}

// WithOwner - flag to query api keys owned by the currently authenticated user.
func (f SecurityGetAPIKey) WithOwner(v bool) func(*SecurityGetAPIKeyRequest) {
	return func(r *SecurityGetAPIKeyRequest) {
		r.Owner = &v
	}
}

// WithRealmName - realm name of the user who created this api key to be retrieved.
func (f SecurityGetAPIKey) WithRealmName(v string) func(*SecurityGetAPIKeyRequest) {
	return func(r *SecurityGetAPIKeyRequest) {
		r.RealmName = v
	}
}

// WithUsername - user name of the user who created this api key to be retrieved.
func (f SecurityGetAPIKey) WithUsername(v string) func(*SecurityGetAPIKeyRequest) {
	return func(r *SecurityGetAPIKeyRequest) {
		r.Username = v
	}
}

// WithWithLimitedBy - flag to show the limited-by role descriptors of api keys.
func (f SecurityGetAPIKey) WithWithLimitedBy(v bool) func(*SecurityGetAPIKeyRequest) {
	return func(r *SecurityGetAPIKeyRequest) {
		r.WithLimitedBy = &v
	}
}

// WithWithProfileUID - flag to also retrieve the api key's owner profile uid, if it exists.
func (f SecurityGetAPIKey) WithWithProfileUID(v bool) func(*SecurityGetAPIKeyRequest) {
	return func(r *SecurityGetAPIKeyRequest) {
		r.WithProfileUID = &v
	}
}

// WithPretty makes the response body pretty-printed.
func (f SecurityGetAPIKey) WithPretty() func(*SecurityGetAPIKeyRequest) {
	return func(r *SecurityGetAPIKeyRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f SecurityGetAPIKey) WithHuman() func(*SecurityGetAPIKeyRequest) {
	return func(r *SecurityGetAPIKeyRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f SecurityGetAPIKey) WithErrorTrace() func(*SecurityGetAPIKeyRequest) {
	return func(r *SecurityGetAPIKeyRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f SecurityGetAPIKey) WithFilterPath(v ...string) func(*SecurityGetAPIKeyRequest) {
	return func(r *SecurityGetAPIKeyRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f SecurityGetAPIKey) WithHeader(h map[string]string) func(*SecurityGetAPIKeyRequest) {
	return func(r *SecurityGetAPIKeyRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f SecurityGetAPIKey) WithOpaqueID(s string) func(*SecurityGetAPIKeyRequest) {
	return func(r *SecurityGetAPIKeyRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
