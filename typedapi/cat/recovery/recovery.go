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

// Get shard recovery information.
//
// Get information about ongoing and completed shard recoveries.
// Shard recovery is the process of initializing a shard copy, such as restoring
// a primary shard from a snapshot or syncing a replica shard from a primary
// shard. When a shard recovery completes, the recovered shard is available for
// search and indexing.
// For data streams, the API returns information about the stream’s backing
// indices.
// IMPORTANT: cat APIs are only intended for human consumption using the command
// line or Kibana console. They are not intended for use by applications. For
// application consumption, use the index recovery API.
package recovery

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
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/bytes"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/catrecoverycolumn"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/timeunit"
)

const (
	indexMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Recovery struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	index string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewRecovery type alias for index.
type NewRecovery func() *Recovery

// NewRecoveryFunc returns a new instance of Recovery with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewRecoveryFunc(tp elastictransport.Interface) NewRecovery {
	return func() *Recovery {
		n := New(tp)

		return n
	}
}

// Get shard recovery information.
//
// Get information about ongoing and completed shard recoveries.
// Shard recovery is the process of initializing a shard copy, such as restoring
// a primary shard from a snapshot or syncing a replica shard from a primary
// shard. When a shard recovery completes, the recovered shard is available for
// search and indexing.
// For data streams, the API returns information about the stream’s backing
// indices.
// IMPORTANT: cat APIs are only intended for human consumption using the command
// line or Kibana console. They are not intended for use by applications. For
// application consumption, use the index recovery API.
//
// https://www.elastic.co/docs/api/doc/elasticsearch/operation/operation-cat-recovery
func New(tp elastictransport.Interface) *Recovery {
	r := &Recovery{
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
func (r *Recovery) HttpRequest(ctx context.Context) (*http.Request, error) {
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
		path.WriteString("recovery")

		method = http.MethodGet
	case r.paramSet == indexMask:
		path.WriteString("/")
		path.WriteString("_cat")
		path.WriteString("/")
		path.WriteString("recovery")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "index", r.index)
		}
		path.WriteString(r.index)

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
func (r Recovery) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "cat.recovery")
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
		instrument.BeforeRequest(req, "cat.recovery")
		if reader := instrument.RecordRequestBody(ctx, "cat.recovery", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "cat.recovery")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the Recovery query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a recovery.Response
func (r Recovery) Do(providedCtx context.Context) (Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "cat.recovery")
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
func (r Recovery) IsSuccess(providedCtx context.Context) (bool, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "cat.recovery")
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
		err := fmt.Errorf("an error happened during the Recovery query execution, status code: %d", res.StatusCode)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return false, err
	}

	return false, nil
}

// Header set a key, value pair in the Recovery headers map.
func (r *Recovery) Header(key, value string) *Recovery {
	r.headers.Set(key, value)

	return r
}

// Index A comma-separated list of data streams, indices, and aliases used to limit
// the request.
// Supports wildcards (`*`). To target all data streams and indices, omit this
// parameter or use `*` or `_all`.
// API Name: index
func (r *Recovery) Index(index string) *Recovery {
	r.paramSet |= indexMask
	r.index = index

	return r
}

// ActiveOnly If `true`, the response only includes ongoing shard recoveries.
// API name: active_only
func (r *Recovery) ActiveOnly(activeonly bool) *Recovery {
	r.values.Set("active_only", strconv.FormatBool(activeonly))

	return r
}

// Bytes The unit used to display byte values.
// API name: bytes
func (r *Recovery) Bytes(bytes bytes.Bytes) *Recovery {
	r.values.Set("bytes", bytes.String())

	return r
}

// Detailed If `true`, the response includes detailed information about shard recoveries.
// API name: detailed
func (r *Recovery) Detailed(detailed bool) *Recovery {
	r.values.Set("detailed", strconv.FormatBool(detailed))

	return r
}

// H A comma-separated list of columns names to display.
// It supports simple wildcards.
// API name: h
func (r *Recovery) H(catrecoverycolumns ...catrecoverycolumn.CatRecoveryColumn) *Recovery {
	tmp := []string{}
	for _, item := range catrecoverycolumns {
		tmp = append(tmp, item.String())
	}
	r.values.Set("expand_wildcards", strings.Join(tmp, ","))

	return r
}

// S A comma-separated list of column names or aliases that determines the sort
// order.
// Sorting defaults to ascending and can be changed by setting `:asc`
// or `:desc` as a suffix to the column name.
// API name: s
func (r *Recovery) S(names ...string) *Recovery {
	r.values.Set("s", strings.Join(names, ","))

	return r
}

// Time The unit used to display time values.
// API name: time
func (r *Recovery) Time(time timeunit.TimeUnit) *Recovery {
	r.values.Set("time", time.String())

	return r
}

// Format Specifies the format to return the columnar data in, can be set to
// `text`, `json`, `cbor`, `yaml`, or `smile`.
// API name: format
func (r *Recovery) Format(format string) *Recovery {
	r.values.Set("format", format)

	return r
}

// Help When set to `true` will output available columns. This option
// can't be combined with any other query string option.
// API name: help
func (r *Recovery) Help(help bool) *Recovery {
	r.values.Set("help", strconv.FormatBool(help))

	return r
}

// V When set to `true` will enable verbose output.
// API name: v
func (r *Recovery) V(v bool) *Recovery {
	r.values.Set("v", strconv.FormatBool(v))

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *Recovery) ErrorTrace(errortrace bool) *Recovery {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *Recovery) FilterPath(filterpaths ...string) *Recovery {
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
func (r *Recovery) Human(human bool) *Recovery {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *Recovery) Pretty(pretty bool) *Recovery {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}
