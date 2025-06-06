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

func newMLGetModelSnapshotsFunc(t Transport) MLGetModelSnapshots {
	return func(job_id string, o ...func(*MLGetModelSnapshotsRequest)) (*Response, error) {
		var r = MLGetModelSnapshotsRequest{JobID: job_id}
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

// MLGetModelSnapshots - Retrieves information about model snapshots.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-get-snapshot.html.
type MLGetModelSnapshots func(job_id string, o ...func(*MLGetModelSnapshotsRequest)) (*Response, error)

// MLGetModelSnapshotsRequest configures the ML Get Model Snapshots API request.
type MLGetModelSnapshotsRequest struct {
	Body io.Reader

	JobID      string
	SnapshotID string

	Desc  *bool
	End   interface{}
	From  *int
	Size  *int
	Sort  string
	Start interface{}

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

// Do executes the request and returns response or error.
func (r MLGetModelSnapshotsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "ml.get_model_snapshots")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "POST"

	path.Grow(7 + 1 + len("_ml") + 1 + len("anomaly_detectors") + 1 + len(r.JobID) + 1 + len("model_snapshots") + 1 + len(r.SnapshotID))
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
	path.WriteString("model_snapshots")
	if r.SnapshotID != "" {
		path.WriteString("/")
		path.WriteString(r.SnapshotID)
		if instrument, ok := r.Instrument.(Instrumentation); ok {
			instrument.RecordPathPart(ctx, "snapshot_id", r.SnapshotID)
		}
	}

	params = make(map[string]string)

	if r.Desc != nil {
		params["desc"] = strconv.FormatBool(*r.Desc)
	}

	if r.End != nil {
		params["end"] = fmt.Sprintf("%v", r.End)
	}

	if r.From != nil {
		params["from"] = strconv.FormatInt(int64(*r.From), 10)
	}

	if r.Size != nil {
		params["size"] = strconv.FormatInt(int64(*r.Size), 10)
	}

	if r.Sort != "" {
		params["sort"] = r.Sort
	}

	if r.Start != nil {
		params["start"] = fmt.Sprintf("%v", r.Start)
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
		instrument.BeforeRequest(req, "ml.get_model_snapshots")
		if reader := instrument.RecordRequestBody(ctx, "ml.get_model_snapshots", r.Body); reader != nil {
			req.Body = reader
		}
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "ml.get_model_snapshots")
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
func (f MLGetModelSnapshots) WithContext(v context.Context) func(*MLGetModelSnapshotsRequest) {
	return func(r *MLGetModelSnapshotsRequest) {
		r.ctx = v
	}
}

// WithBody - Model snapshot selection criteria.
func (f MLGetModelSnapshots) WithBody(v io.Reader) func(*MLGetModelSnapshotsRequest) {
	return func(r *MLGetModelSnapshotsRequest) {
		r.Body = v
	}
}

// WithSnapshotID - the ID of the snapshot to fetch.
func (f MLGetModelSnapshots) WithSnapshotID(v string) func(*MLGetModelSnapshotsRequest) {
	return func(r *MLGetModelSnapshotsRequest) {
		r.SnapshotID = v
	}
}

// WithDesc - true if the results should be sorted in descending order.
func (f MLGetModelSnapshots) WithDesc(v bool) func(*MLGetModelSnapshotsRequest) {
	return func(r *MLGetModelSnapshotsRequest) {
		r.Desc = &v
	}
}

// WithEnd - the filter 'end' query parameter.
func (f MLGetModelSnapshots) WithEnd(v interface{}) func(*MLGetModelSnapshotsRequest) {
	return func(r *MLGetModelSnapshotsRequest) {
		r.End = v
	}
}

// WithFrom - skips a number of documents.
func (f MLGetModelSnapshots) WithFrom(v int) func(*MLGetModelSnapshotsRequest) {
	return func(r *MLGetModelSnapshotsRequest) {
		r.From = &v
	}
}

// WithSize - the default number of documents returned in queries as a string..
func (f MLGetModelSnapshots) WithSize(v int) func(*MLGetModelSnapshotsRequest) {
	return func(r *MLGetModelSnapshotsRequest) {
		r.Size = &v
	}
}

// WithSort - name of the field to sort on.
func (f MLGetModelSnapshots) WithSort(v string) func(*MLGetModelSnapshotsRequest) {
	return func(r *MLGetModelSnapshotsRequest) {
		r.Sort = v
	}
}

// WithStart - the filter 'start' query parameter.
func (f MLGetModelSnapshots) WithStart(v interface{}) func(*MLGetModelSnapshotsRequest) {
	return func(r *MLGetModelSnapshotsRequest) {
		r.Start = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f MLGetModelSnapshots) WithPretty() func(*MLGetModelSnapshotsRequest) {
	return func(r *MLGetModelSnapshotsRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f MLGetModelSnapshots) WithHuman() func(*MLGetModelSnapshotsRequest) {
	return func(r *MLGetModelSnapshotsRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f MLGetModelSnapshots) WithErrorTrace() func(*MLGetModelSnapshotsRequest) {
	return func(r *MLGetModelSnapshotsRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f MLGetModelSnapshots) WithFilterPath(v ...string) func(*MLGetModelSnapshotsRequest) {
	return func(r *MLGetModelSnapshotsRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f MLGetModelSnapshots) WithHeader(h map[string]string) func(*MLGetModelSnapshotsRequest) {
	return func(r *MLGetModelSnapshotsRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f MLGetModelSnapshots) WithOpaqueID(s string) func(*MLGetModelSnapshotsRequest) {
	return func(r *MLGetModelSnapshotsRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
