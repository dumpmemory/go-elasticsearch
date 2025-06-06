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

func newMLDeleteCalendarEventFunc(t Transport) MLDeleteCalendarEvent {
	return func(calendar_id string, event_id string, o ...func(*MLDeleteCalendarEventRequest)) (*Response, error) {
		var r = MLDeleteCalendarEventRequest{CalendarID: calendar_id, EventID: event_id}
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

// MLDeleteCalendarEvent - Deletes scheduled events from a calendar.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-delete-calendar-event.html.
type MLDeleteCalendarEvent func(calendar_id string, event_id string, o ...func(*MLDeleteCalendarEventRequest)) (*Response, error)

// MLDeleteCalendarEventRequest configures the ML Delete Calendar Event API request.
type MLDeleteCalendarEventRequest struct {
	CalendarID string
	EventID    string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

// Do executes the request and returns response or error.
func (r MLDeleteCalendarEventRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "ml.delete_calendar_event")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "DELETE"

	path.Grow(7 + 1 + len("_ml") + 1 + len("calendars") + 1 + len(r.CalendarID) + 1 + len("events") + 1 + len(r.EventID))
	path.WriteString("http://")
	path.WriteString("/")
	path.WriteString("_ml")
	path.WriteString("/")
	path.WriteString("calendars")
	path.WriteString("/")
	path.WriteString(r.CalendarID)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.RecordPathPart(ctx, "calendar_id", r.CalendarID)
	}
	path.WriteString("/")
	path.WriteString("events")
	path.WriteString("/")
	path.WriteString(r.EventID)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.RecordPathPart(ctx, "event_id", r.EventID)
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
		instrument.BeforeRequest(req, "ml.delete_calendar_event")
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "ml.delete_calendar_event")
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
func (f MLDeleteCalendarEvent) WithContext(v context.Context) func(*MLDeleteCalendarEventRequest) {
	return func(r *MLDeleteCalendarEventRequest) {
		r.ctx = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f MLDeleteCalendarEvent) WithPretty() func(*MLDeleteCalendarEventRequest) {
	return func(r *MLDeleteCalendarEventRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f MLDeleteCalendarEvent) WithHuman() func(*MLDeleteCalendarEventRequest) {
	return func(r *MLDeleteCalendarEventRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f MLDeleteCalendarEvent) WithErrorTrace() func(*MLDeleteCalendarEventRequest) {
	return func(r *MLDeleteCalendarEventRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f MLDeleteCalendarEvent) WithFilterPath(v ...string) func(*MLDeleteCalendarEventRequest) {
	return func(r *MLDeleteCalendarEventRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f MLDeleteCalendarEvent) WithHeader(h map[string]string) func(*MLDeleteCalendarEventRequest) {
	return func(r *MLDeleteCalendarEventRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f MLDeleteCalendarEvent) WithOpaqueID(s string) func(*MLDeleteCalendarEventRequest) {
	return func(r *MLDeleteCalendarEventRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
