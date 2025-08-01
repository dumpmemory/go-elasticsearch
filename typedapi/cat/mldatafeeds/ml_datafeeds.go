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

// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

// Get datafeeds.
//
// Get configuration and usage information about datafeeds.
// This API returns a maximum of 10,000 datafeeds.
// If the Elasticsearch security features are enabled, you must have
// `monitor_ml`, `monitor`, `manage_ml`, or `manage`
// cluster privileges to use this API.
//
// IMPORTANT: CAT APIs are only intended for human consumption using the Kibana
// console or command line. They are not intended for use by applications. For
// application consumption, use the get datafeed statistics API.
package mldatafeeds

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/catdatafeedcolumn"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/timeunit"
)

const (
	datafeedidMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type MlDatafeeds struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	datafeedid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewMlDatafeeds type alias for index.
type NewMlDatafeeds func() *MlDatafeeds

// NewMlDatafeedsFunc returns a new instance of MlDatafeeds with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewMlDatafeedsFunc(tp elastictransport.Interface) NewMlDatafeeds {
	return func() *MlDatafeeds {
		n := New(tp)

		return n
	}
}

// Get datafeeds.
//
// Get configuration and usage information about datafeeds.
// This API returns a maximum of 10,000 datafeeds.
// If the Elasticsearch security features are enabled, you must have
// `monitor_ml`, `monitor`, `manage_ml`, or `manage`
// cluster privileges to use this API.
//
// IMPORTANT: CAT APIs are only intended for human consumption using the Kibana
// console or command line. They are not intended for use by applications. For
// application consumption, use the get datafeed statistics API.
//
// https://www.elastic.co/docs/api/doc/elasticsearch/operation/operation-cat-ml-datafeeds
func New(tp elastictransport.Interface) *MlDatafeeds {
	r := &MlDatafeeds{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
	}

	if instrumented, ok := r.transport.(elastictransport.Instrumented); ok {
		if instrument := instrumented.InstrumentationEnabled(); instrument != nil {
			r.instrument = instrument
		}
	}

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *MlDatafeeds) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_cat")
		path.WriteString("/")
		path.WriteString("ml")
		path.WriteString("/")
		path.WriteString("datafeeds")

		method = http.MethodGet
	case r.paramSet == datafeedidMask:
		path.WriteString("/")
		path.WriteString("_cat")
		path.WriteString("/")
		path.WriteString("ml")
		path.WriteString("/")
		path.WriteString("datafeeds")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "datafeedid", r.datafeedid)
		}
		path.WriteString(r.datafeedid)

		method = http.MethodGet
	}

	r.path.Path = path.String()
	r.path.RawQuery = r.values.Encode()

	if r.path.Path == "" {
		return nil, ErrBuildPath
	}

	if ctx != nil {
		req, err = http.NewRequestWithContext(ctx, method, r.path.String(), r.raw)
	} else {
		req, err = http.NewRequest(method, r.path.String(), r.raw)
	}

	req.Header = r.headers.Clone()

	if req.Header.Get("Accept") == "" {
		req.Header.Set("Accept", "application/vnd.elasticsearch+json;compatible-with=9")
	}

	if err != nil {
		return req, fmt.Errorf("could not build http.Request: %w", err)
	}

	return req, nil
}

// Perform runs the http.Request through the provided transport and returns an http.Response.
func (r MlDatafeeds) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "cat.ml_datafeeds")
			defer instrument.Close(ctx)
		}
	}
	if ctx == nil {
		ctx = providedCtx
	}

	req, err := r.HttpRequest(ctx)
	if err != nil {
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return nil, err
	}

	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.BeforeRequest(req, "cat.ml_datafeeds")
		if reader := instrument.RecordRequestBody(ctx, "cat.ml_datafeeds", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "cat.ml_datafeeds")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the MlDatafeeds query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a mldatafeeds.Response
func (r MlDatafeeds) Do(providedCtx context.Context) (Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "cat.ml_datafeeds")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	response := NewResponse()

	res, err := r.Perform(ctx)
	if err != nil {
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode < 299 {
		err = json.NewDecoder(res.Body).Decode(&response)
		if err != nil {
			if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
				instrument.RecordError(ctx, err)
			}
			return nil, err
		}

		return response, nil
	}

	errorResponse := types.NewElasticsearchError()
	err = json.NewDecoder(res.Body).Decode(errorResponse)
	if err != nil {
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return nil, err
	}

	if errorResponse.Status == 0 {
		errorResponse.Status = res.StatusCode
	}

	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.RecordError(ctx, errorResponse)
	}
	return nil, errorResponse
}

// IsSuccess allows to run a query with a context and retrieve the result as a boolean.
// This only exists for endpoints without a request payload and allows for quick control flow.
func (r MlDatafeeds) IsSuccess(providedCtx context.Context) (bool, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "cat.ml_datafeeds")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	res, err := r.Perform(ctx)

	if err != nil {
		return false, err
	}
	io.Copy(io.Discard, res.Body)
	err = res.Body.Close()
	if err != nil {
		return false, err
	}

	if res.StatusCode >= 200 && res.StatusCode < 300 {
		return true, nil
	}

	if res.StatusCode != 404 {
		err := fmt.Errorf("an error happened during the MlDatafeeds query execution, status code: %d", res.StatusCode)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return false, err
	}

	return false, nil
}

// Header set a key, value pair in the MlDatafeeds headers map.
func (r *MlDatafeeds) Header(key, value string) *MlDatafeeds {
	r.headers.Set(key, value)

	return r
}

// DatafeedId A numerical character string that uniquely identifies the datafeed.
// API Name: datafeedid
func (r *MlDatafeeds) DatafeedId(datafeedid string) *MlDatafeeds {
	r.paramSet |= datafeedidMask
	r.datafeedid = datafeedid

	return r
}

// AllowNoMatch Specifies what to do when the request:
//
// * Contains wildcard expressions and there are no datafeeds that match.
// * Contains the `_all` string or no identifiers and there are no matches.
// * Contains wildcard expressions and there are only partial matches.
//
// If `true`, the API returns an empty datafeeds array when there are no matches
// and the subset of results when
// there are partial matches. If `false`, the API returns a 404 status code when
// there are no matches or only
// partial matches.
// API name: allow_no_match
func (r *MlDatafeeds) AllowNoMatch(allownomatch bool) *MlDatafeeds {
	r.values.Set("allow_no_match", strconv.FormatBool(allownomatch))

	return r
}

// H Comma-separated list of column names to display.
// API name: h
func (r *MlDatafeeds) H(catdatafeedcolumns ...catdatafeedcolumn.CatDatafeedColumn) *MlDatafeeds {
	tmp := []string{}
	for _, item := range catdatafeedcolumns {
		tmp = append(tmp, item.String())
	}
	r.values.Set("expand_wildcards", strings.Join(tmp, ","))

	return r
}

// S Comma-separated list of column names or column aliases used to sort the
// response.
// API name: s
func (r *MlDatafeeds) S(catdatafeedcolumns ...catdatafeedcolumn.CatDatafeedColumn) *MlDatafeeds {
	tmp := []string{}
	for _, item := range catdatafeedcolumns {
		tmp = append(tmp, item.String())
	}
	r.values.Set("expand_wildcards", strings.Join(tmp, ","))

	return r
}

// Time The unit used to display time values.
// API name: time
func (r *MlDatafeeds) Time(time timeunit.TimeUnit) *MlDatafeeds {
	r.values.Set("time", time.String())

	return r
}

// Format Specifies the format to return the columnar data in, can be set to
// `text`, `json`, `cbor`, `yaml`, or `smile`.
// API name: format
func (r *MlDatafeeds) Format(format string) *MlDatafeeds {
	r.values.Set("format", format)

	return r
}

// Help When set to `true` will output available columns. This option
// can't be combined with any other query string option.
// API name: help
func (r *MlDatafeeds) Help(help bool) *MlDatafeeds {
	r.values.Set("help", strconv.FormatBool(help))

	return r
}

// V When set to `true` will enable verbose output.
// API name: v
func (r *MlDatafeeds) V(v bool) *MlDatafeeds {
	r.values.Set("v", strconv.FormatBool(v))

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *MlDatafeeds) ErrorTrace(errortrace bool) *MlDatafeeds {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *MlDatafeeds) FilterPath(filterpaths ...string) *MlDatafeeds {
	tmp := []string{}
	for _, item := range filterpaths {
		tmp = append(tmp, fmt.Sprintf("%v", item))
	}
	r.values.Set("filter_path", strings.Join(tmp, ","))

	return r
}

// Human When set to `true` will return statistics in a format suitable for humans.
// For example `"exists_time": "1h"` for humans and
// `"exists_time_in_millis": 3600000` for computers. When disabled the human
// readable values will be omitted. This makes sense for responses being
// consumed
// only by machines.
// API name: human
func (r *MlDatafeeds) Human(human bool) *MlDatafeeds {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *MlDatafeeds) Pretty(pretty bool) *MlDatafeeds {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}
