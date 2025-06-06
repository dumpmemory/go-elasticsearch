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

func newMLFlushJobFunc(t Transport) MLFlushJob {
	return func(job_id string, o ...func(*MLFlushJobRequest)) (*Response, error) {
		var r = MLFlushJobRequest{JobID: job_id}
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

// MLFlushJob - Forces any buffered data to be processed by the job.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-flush-job.html.
type MLFlushJob func(job_id string, o ...func(*MLFlushJobRequest)) (*Response, error)

// MLFlushJobRequest configures the ML Flush Job API request.
type MLFlushJobRequest struct {
	Body io.Reader

	JobID string

	AdvanceTime string
	CalcInterim *bool
	End         string
	SkipTime    string
	Start       string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

// Do executes the request and returns response or error.
func (r MLFlushJobRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "ml.flush_job")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "POST"

	path.Grow(7 + 1 + len("_ml") + 1 + len("anomaly_detectors") + 1 + len(r.JobID) + 1 + len("_flush"))
	path.WriteString("http://")
	path.WriteString("/")
	path.WriteString("_ml")
	path.WriteString("/")
	path.WriteString("anomaly_detectors")
	path.WriteString("/")
	path.WriteString(r.JobID)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.RecordPathPart(ctx, "job_id", r.JobID)
	}
	path.WriteString("/")
	path.WriteString("_flush")

	params = make(map[string]string)

	if r.AdvanceTime != "" {
		params["advance_time"] = r.AdvanceTime
	}

	if r.CalcInterim != nil {
		params["calc_interim"] = strconv.FormatBool(*r.CalcInterim)
	}

	if r.End != "" {
		params["end"] = r.End
	}

	if r.SkipTime != "" {
		params["skip_time"] = r.SkipTime
	}

	if r.Start != "" {
		params["start"] = r.Start
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
		instrument.BeforeRequest(req, "ml.flush_job")
		if reader := instrument.RecordRequestBody(ctx, "ml.flush_job", r.Body); reader != nil {
			req.Body = reader
		}
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "ml.flush_job")
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
func (f MLFlushJob) WithContext(v context.Context) func(*MLFlushJobRequest) {
	return func(r *MLFlushJobRequest) {
		r.ctx = v
	}
}

// WithBody - Flush parameters.
func (f MLFlushJob) WithBody(v io.Reader) func(*MLFlushJobRequest) {
	return func(r *MLFlushJobRequest) {
		r.Body = v
	}
}

// WithAdvanceTime - advances time to the given value generating results and updating the model for the advanced interval.
func (f MLFlushJob) WithAdvanceTime(v string) func(*MLFlushJobRequest) {
	return func(r *MLFlushJobRequest) {
		r.AdvanceTime = v
	}
}

// WithCalcInterim - calculates interim results for the most recent bucket or all buckets within the latency period.
func (f MLFlushJob) WithCalcInterim(v bool) func(*MLFlushJobRequest) {
	return func(r *MLFlushJobRequest) {
		r.CalcInterim = &v
	}
}

// WithEnd - when used in conjunction with calc_interim, specifies the range of buckets on which to calculate interim results.
func (f MLFlushJob) WithEnd(v string) func(*MLFlushJobRequest) {
	return func(r *MLFlushJobRequest) {
		r.End = v
	}
}

// WithSkipTime - skips time to the given value without generating results or updating the model for the skipped interval.
func (f MLFlushJob) WithSkipTime(v string) func(*MLFlushJobRequest) {
	return func(r *MLFlushJobRequest) {
		r.SkipTime = v
	}
}

// WithStart - when used in conjunction with calc_interim, specifies the range of buckets on which to calculate interim results.
func (f MLFlushJob) WithStart(v string) func(*MLFlushJobRequest) {
	return func(r *MLFlushJobRequest) {
		r.Start = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f MLFlushJob) WithPretty() func(*MLFlushJobRequest) {
	return func(r *MLFlushJobRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f MLFlushJob) WithHuman() func(*MLFlushJobRequest) {
	return func(r *MLFlushJobRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f MLFlushJob) WithErrorTrace() func(*MLFlushJobRequest) {
	return func(r *MLFlushJobRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f MLFlushJob) WithFilterPath(v ...string) func(*MLFlushJobRequest) {
	return func(r *MLFlushJobRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f MLFlushJob) WithHeader(h map[string]string) func(*MLFlushJobRequest) {
	return func(r *MLFlushJobRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f MLFlushJob) WithOpaqueID(s string) func(*MLFlushJobRequest) {
	return func(r *MLFlushJobRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
