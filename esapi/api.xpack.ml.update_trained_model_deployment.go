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

func newMLUpdateTrainedModelDeploymentFunc(t Transport) MLUpdateTrainedModelDeployment {
	return func(model_id string, o ...func(*MLUpdateTrainedModelDeploymentRequest)) (*Response, error) {
		var r = MLUpdateTrainedModelDeploymentRequest{ModelID: model_id}
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

// MLUpdateTrainedModelDeployment - Updates certain properties of trained model deployment.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/update-trained-model-deployment.html.
type MLUpdateTrainedModelDeployment func(model_id string, o ...func(*MLUpdateTrainedModelDeploymentRequest)) (*Response, error)

// MLUpdateTrainedModelDeploymentRequest configures the ML Update Trained Model Deployment API request.
type MLUpdateTrainedModelDeploymentRequest struct {
	Body io.Reader

	ModelID string

	NumberOfAllocations *int

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

// Do executes the request and returns response or error.
func (r MLUpdateTrainedModelDeploymentRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "ml.update_trained_model_deployment")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "POST"

	path.Grow(7 + 1 + len("_ml") + 1 + len("trained_models") + 1 + len(r.ModelID) + 1 + len("deployment") + 1 + len("_update"))
	path.WriteString("http://")
	path.WriteString("/")
	path.WriteString("_ml")
	path.WriteString("/")
	path.WriteString("trained_models")
	path.WriteString("/")
	path.WriteString(r.ModelID)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.RecordPathPart(ctx, "model_id", r.ModelID)
	}
	path.WriteString("/")
	path.WriteString("deployment")
	path.WriteString("/")
	path.WriteString("_update")

	params = make(map[string]string)

	if r.NumberOfAllocations != nil {
		params["number_of_allocations"] = strconv.FormatInt(int64(*r.NumberOfAllocations), 10)
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
		instrument.BeforeRequest(req, "ml.update_trained_model_deployment")
		if reader := instrument.RecordRequestBody(ctx, "ml.update_trained_model_deployment", r.Body); reader != nil {
			req.Body = reader
		}
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "ml.update_trained_model_deployment")
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
func (f MLUpdateTrainedModelDeployment) WithContext(v context.Context) func(*MLUpdateTrainedModelDeploymentRequest) {
	return func(r *MLUpdateTrainedModelDeploymentRequest) {
		r.ctx = v
	}
}

// WithBody - The updated trained model deployment settings.
func (f MLUpdateTrainedModelDeployment) WithBody(v io.Reader) func(*MLUpdateTrainedModelDeploymentRequest) {
	return func(r *MLUpdateTrainedModelDeploymentRequest) {
		r.Body = v
	}
}

// WithNumberOfAllocations - update the model deployment to this number of allocations..
func (f MLUpdateTrainedModelDeployment) WithNumberOfAllocations(v int) func(*MLUpdateTrainedModelDeploymentRequest) {
	return func(r *MLUpdateTrainedModelDeploymentRequest) {
		r.NumberOfAllocations = &v
	}
}

// WithPretty makes the response body pretty-printed.
func (f MLUpdateTrainedModelDeployment) WithPretty() func(*MLUpdateTrainedModelDeploymentRequest) {
	return func(r *MLUpdateTrainedModelDeploymentRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f MLUpdateTrainedModelDeployment) WithHuman() func(*MLUpdateTrainedModelDeploymentRequest) {
	return func(r *MLUpdateTrainedModelDeploymentRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f MLUpdateTrainedModelDeployment) WithErrorTrace() func(*MLUpdateTrainedModelDeploymentRequest) {
	return func(r *MLUpdateTrainedModelDeploymentRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f MLUpdateTrainedModelDeployment) WithFilterPath(v ...string) func(*MLUpdateTrainedModelDeploymentRequest) {
	return func(r *MLUpdateTrainedModelDeploymentRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f MLUpdateTrainedModelDeployment) WithHeader(h map[string]string) func(*MLUpdateTrainedModelDeploymentRequest) {
	return func(r *MLUpdateTrainedModelDeploymentRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f MLUpdateTrainedModelDeployment) WithOpaqueID(s string) func(*MLUpdateTrainedModelDeploymentRequest) {
	return func(r *MLUpdateTrainedModelDeploymentRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
