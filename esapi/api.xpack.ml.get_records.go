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
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func newMLGetRecordsFunc(t Transport) MLGetRecords {
	return func(job_id string, o ...func(*MLGetRecordsRequest)) (*Response, error) {
		var r = MLGetRecordsRequest{JobID: job_id}
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

// MLGetRecords - Retrieves anomaly records for an anomaly detection job.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-get-record.html.
type MLGetRecords func(job_id string, o ...func(*MLGetRecordsRequest)) (*Response, error)

// MLGetRecordsRequest configures the ML Get Records API request.
type MLGetRecordsRequest struct {
	Body io.Reader

	JobID string

	Desc           *bool
	End            string
	ExcludeInterim *bool
	From           *int
	RecordScore    interface{}
	Size           *int
	Sort           string
	Start          string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

// Do executes the request and returns response or error.
func (r MLGetRecordsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "ml.get_records")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "POST"

	path.Grow(7 + 1 + len("_ml") + 1 + len("anomaly_detectors") + 1 + len(r.JobID) + 1 + len("results") + 1 + len("records"))
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
	path.WriteString("results")
	path.WriteString("/")
	path.WriteString("records")

	params = make(map[string]string)

	if r.Desc != nil {
		params["desc"] = strconv.FormatBool(*r.Desc)
	}

	if r.End != "" {
		params["end"] = r.End
	}

	if r.ExcludeInterim != nil {
		params["exclude_interim"] = strconv.FormatBool(*r.ExcludeInterim)
	}

	if r.From != nil {
		params["from"] = strconv.FormatInt(int64(*r.From), 10)
	}

	if r.RecordScore != nil {
		params["record_score"] = fmt.Sprintf("%v", r.RecordScore)
	}

	if r.Size != nil {
		params["size"] = strconv.FormatInt(int64(*r.Size), 10)
	}

	if r.Sort != "" {
		params["sort"] = r.Sort
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
		instrument.BeforeRequest(req, "ml.get_records")
		if reader := instrument.RecordRequestBody(ctx, "ml.get_records", r.Body); reader != nil {
			req.Body = reader
		}
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "ml.get_records")
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
func (f MLGetRecords) WithContext(v context.Context) func(*MLGetRecordsRequest) {
	return func(r *MLGetRecordsRequest) {
		r.ctx = v
	}
}

// WithBody - Record selection criteria.
func (f MLGetRecords) WithBody(v io.Reader) func(*MLGetRecordsRequest) {
	return func(r *MLGetRecordsRequest) {
		r.Body = v
	}
}

// WithDesc - set the sort direction.
func (f MLGetRecords) WithDesc(v bool) func(*MLGetRecordsRequest) {
	return func(r *MLGetRecordsRequest) {
		r.Desc = &v
	}
}

// WithEnd - end time filter for records.
func (f MLGetRecords) WithEnd(v string) func(*MLGetRecordsRequest) {
	return func(r *MLGetRecordsRequest) {
		r.End = v
	}
}

// WithExcludeInterim - exclude interim results.
func (f MLGetRecords) WithExcludeInterim(v bool) func(*MLGetRecordsRequest) {
	return func(r *MLGetRecordsRequest) {
		r.ExcludeInterim = &v
	}
}

// WithFrom - skips a number of records.
func (f MLGetRecords) WithFrom(v int) func(*MLGetRecordsRequest) {
	return func(r *MLGetRecordsRequest) {
		r.From = &v
	}
}

// WithRecordScore - returns records with anomaly scores greater or equal than this value.
func (f MLGetRecords) WithRecordScore(v interface{}) func(*MLGetRecordsRequest) {
	return func(r *MLGetRecordsRequest) {
		r.RecordScore = v
	}
}

// WithSize - specifies a max number of records to get.
func (f MLGetRecords) WithSize(v int) func(*MLGetRecordsRequest) {
	return func(r *MLGetRecordsRequest) {
		r.Size = &v
	}
}

// WithSort - sort records by a particular field.
func (f MLGetRecords) WithSort(v string) func(*MLGetRecordsRequest) {
	return func(r *MLGetRecordsRequest) {
		r.Sort = v
	}
}

// WithStart - start time filter for records.
func (f MLGetRecords) WithStart(v string) func(*MLGetRecordsRequest) {
	return func(r *MLGetRecordsRequest) {
		r.Start = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f MLGetRecords) WithPretty() func(*MLGetRecordsRequest) {
	return func(r *MLGetRecordsRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f MLGetRecords) WithHuman() func(*MLGetRecordsRequest) {
	return func(r *MLGetRecordsRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f MLGetRecords) WithErrorTrace() func(*MLGetRecordsRequest) {
	return func(r *MLGetRecordsRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f MLGetRecords) WithFilterPath(v ...string) func(*MLGetRecordsRequest) {
	return func(r *MLGetRecordsRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f MLGetRecords) WithHeader(h map[string]string) func(*MLGetRecordsRequest) {
	return func(r *MLGetRecordsRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f MLGetRecords) WithOpaqueID(s string) func(*MLGetRecordsRequest) {
	return func(r *MLGetRecordsRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
