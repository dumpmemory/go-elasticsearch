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

func newSynonymsGetSynonymRuleFunc(t Transport) SynonymsGetSynonymRule {
	return func(rule_id string, set_id string, o ...func(*SynonymsGetSynonymRuleRequest)) (*Response, error) {
		var r = SynonymsGetSynonymRuleRequest{RuleID: rule_id, SetID: set_id}
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

// SynonymsGetSynonymRule retrieves a synonym rule from a synonym set
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/master/get-synonym-rule.html.
type SynonymsGetSynonymRule func(rule_id string, set_id string, o ...func(*SynonymsGetSynonymRuleRequest)) (*Response, error)

// SynonymsGetSynonymRuleRequest configures the Synonyms Get Synonym Rule API request.
type SynonymsGetSynonymRuleRequest struct {
	RuleID string
	SetID  string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

// Do executes the request and returns response or error.
func (r SynonymsGetSynonymRuleRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "synonyms.get_synonym_rule")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "GET"

	path.Grow(7 + 1 + len("_synonyms") + 1 + len(r.SetID) + 1 + len(r.RuleID))
	path.WriteString("http://")
	path.WriteString("/")
	path.WriteString("_synonyms")
	path.WriteString("/")
	path.WriteString(r.SetID)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.RecordPathPart(ctx, "set_id", r.SetID)
	}
	path.WriteString("/")
	path.WriteString(r.RuleID)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.RecordPathPart(ctx, "rule_id", r.RuleID)
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
		instrument.BeforeRequest(req, "synonyms.get_synonym_rule")
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "synonyms.get_synonym_rule")
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
func (f SynonymsGetSynonymRule) WithContext(v context.Context) func(*SynonymsGetSynonymRuleRequest) {
	return func(r *SynonymsGetSynonymRuleRequest) {
		r.ctx = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f SynonymsGetSynonymRule) WithPretty() func(*SynonymsGetSynonymRuleRequest) {
	return func(r *SynonymsGetSynonymRuleRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f SynonymsGetSynonymRule) WithHuman() func(*SynonymsGetSynonymRuleRequest) {
	return func(r *SynonymsGetSynonymRuleRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f SynonymsGetSynonymRule) WithErrorTrace() func(*SynonymsGetSynonymRuleRequest) {
	return func(r *SynonymsGetSynonymRuleRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f SynonymsGetSynonymRule) WithFilterPath(v ...string) func(*SynonymsGetSynonymRuleRequest) {
	return func(r *SynonymsGetSynonymRuleRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f SynonymsGetSynonymRule) WithHeader(h map[string]string) func(*SynonymsGetSynonymRuleRequest) {
	return func(r *SynonymsGetSynonymRuleRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f SynonymsGetSynonymRule) WithOpaqueID(s string) func(*SynonymsGetSynonymRuleRequest) {
	return func(r *SynonymsGetSynonymRuleRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
