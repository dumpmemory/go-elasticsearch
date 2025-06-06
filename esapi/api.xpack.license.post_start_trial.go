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
	"time"
)

func newLicensePostStartTrialFunc(t Transport) LicensePostStartTrial {
	return func(o ...func(*LicensePostStartTrialRequest)) (*Response, error) {
		var r = LicensePostStartTrialRequest{}
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

// LicensePostStartTrial - starts a limited time trial license.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/master/start-trial.html.
type LicensePostStartTrial func(o ...func(*LicensePostStartTrialRequest)) (*Response, error)

// LicensePostStartTrialRequest configures the License Post Start Trial API request.
type LicensePostStartTrialRequest struct {
	Acknowledge   *bool
	MasterTimeout time.Duration
	DocumentType  string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

// Do executes the request and returns response or error.
func (r LicensePostStartTrialRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "license.post_start_trial")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "POST"

	path.Grow(7 + len("/_license/start_trial"))
	path.WriteString("http://")
	path.WriteString("/_license/start_trial")

	params = make(map[string]string)

	if r.Acknowledge != nil {
		params["acknowledge"] = strconv.FormatBool(*r.Acknowledge)
	}

	if r.MasterTimeout != 0 {
		params["master_timeout"] = formatDuration(r.MasterTimeout)
	}

	if r.DocumentType != "" {
		params["type"] = r.DocumentType
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
		instrument.BeforeRequest(req, "license.post_start_trial")
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "license.post_start_trial")
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
func (f LicensePostStartTrial) WithContext(v context.Context) func(*LicensePostStartTrialRequest) {
	return func(r *LicensePostStartTrialRequest) {
		r.ctx = v
	}
}

// WithAcknowledge - whether the user has acknowledged acknowledge messages (default: false).
func (f LicensePostStartTrial) WithAcknowledge(v bool) func(*LicensePostStartTrialRequest) {
	return func(r *LicensePostStartTrialRequest) {
		r.Acknowledge = &v
	}
}

// WithMasterTimeout - timeout for processing on master node.
func (f LicensePostStartTrial) WithMasterTimeout(v time.Duration) func(*LicensePostStartTrialRequest) {
	return func(r *LicensePostStartTrialRequest) {
		r.MasterTimeout = v
	}
}

// WithDocumentType - the type of trial license to generate (default: "trial").
func (f LicensePostStartTrial) WithDocumentType(v string) func(*LicensePostStartTrialRequest) {
	return func(r *LicensePostStartTrialRequest) {
		r.DocumentType = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f LicensePostStartTrial) WithPretty() func(*LicensePostStartTrialRequest) {
	return func(r *LicensePostStartTrialRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f LicensePostStartTrial) WithHuman() func(*LicensePostStartTrialRequest) {
	return func(r *LicensePostStartTrialRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f LicensePostStartTrial) WithErrorTrace() func(*LicensePostStartTrialRequest) {
	return func(r *LicensePostStartTrialRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f LicensePostStartTrial) WithFilterPath(v ...string) func(*LicensePostStartTrialRequest) {
	return func(r *LicensePostStartTrialRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f LicensePostStartTrial) WithHeader(h map[string]string) func(*LicensePostStartTrialRequest) {
	return func(r *LicensePostStartTrialRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f LicensePostStartTrial) WithOpaqueID(s string) func(*LicensePostStartTrialRequest) {
	return func(r *LicensePostStartTrialRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
