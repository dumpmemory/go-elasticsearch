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

// Delete a rollup job.
//
// A job must be stopped before it can be deleted.
// If you attempt to delete a started job, an error occurs.
// Similarly, if you attempt to delete a nonexistent job, an exception occurs.
//
// IMPORTANT: When you delete a job, you remove only the process that is
// actively monitoring and rolling up data.
// The API does not delete any previously rolled up data.
// This is by design; a user may wish to roll up a static data set.
// Because the data set is static, after it has been fully rolled up there is no
// need to keep the indexing rollup job around (as there will be no new data).
// Thus the job can be deleted, leaving behind the rolled up data for analysis.
// If you wish to also remove the rollup data and the rollup index contains the
// data for only a single job, you can delete the whole rollup index.
// If the rollup index stores data from several jobs, you must issue a
// delete-by-query that targets the rollup job's identifier in the rollup index.
// For example:
//
// ```
// POST my_rollup_index/_delete_by_query
//
//	{
//	  "query": {
//	    "term": {
//	      "_rollup.id": "the_rollup_job_id"
//	    }
//	  }
//	}
//
// ```
package deletejob

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
)

const (
	idMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type DeleteJob struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	id string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewDeleteJob type alias for index.
type NewDeleteJob func(id string) *DeleteJob

// NewDeleteJobFunc returns a new instance of DeleteJob with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewDeleteJobFunc(tp elastictransport.Interface) NewDeleteJob {
	return func(id string) *DeleteJob {
		n := New(tp)

		n._id(id)

		return n
	}
}

// Delete a rollup job.
//
// A job must be stopped before it can be deleted.
// If you attempt to delete a started job, an error occurs.
// Similarly, if you attempt to delete a nonexistent job, an exception occurs.
//
// IMPORTANT: When you delete a job, you remove only the process that is
// actively monitoring and rolling up data.
// The API does not delete any previously rolled up data.
// This is by design; a user may wish to roll up a static data set.
// Because the data set is static, after it has been fully rolled up there is no
// need to keep the indexing rollup job around (as there will be no new data).
// Thus the job can be deleted, leaving behind the rolled up data for analysis.
// If you wish to also remove the rollup data and the rollup index contains the
// data for only a single job, you can delete the whole rollup index.
// If the rollup index stores data from several jobs, you must issue a
// delete-by-query that targets the rollup job's identifier in the rollup index.
// For example:
//
// ```
// POST my_rollup_index/_delete_by_query
//
//	{
//	  "query": {
//	    "term": {
//	      "_rollup.id": "the_rollup_job_id"
//	    }
//	  }
//	}
//
// ```
//
// https://www.elastic.co/docs/api/doc/elasticsearch/operation/operation-rollup-delete-job
func New(tp elastictransport.Interface) *DeleteJob {
	r := &DeleteJob{
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
func (r *DeleteJob) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == idMask:
		path.WriteString("/")
		path.WriteString("_rollup")
		path.WriteString("/")
		path.WriteString("job")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "id", r.id)
		}
		path.WriteString(r.id)

		method = http.MethodDelete
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
func (r DeleteJob) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "rollup.delete_job")
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
		instrument.BeforeRequest(req, "rollup.delete_job")
		if reader := instrument.RecordRequestBody(ctx, "rollup.delete_job", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "rollup.delete_job")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the DeleteJob query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a deletejob.Response
func (r DeleteJob) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "rollup.delete_job")
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
		err = json.NewDecoder(res.Body).Decode(response)
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
func (r DeleteJob) IsSuccess(providedCtx context.Context) (bool, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "rollup.delete_job")
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
		err := fmt.Errorf("an error happened during the DeleteJob query execution, status code: %d", res.StatusCode)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return false, err
	}

	return false, nil
}

// Header set a key, value pair in the DeleteJob headers map.
func (r *DeleteJob) Header(key, value string) *DeleteJob {
	r.headers.Set(key, value)

	return r
}

// Id Identifier for the job.
// API Name: id
func (r *DeleteJob) _id(id string) *DeleteJob {
	r.paramSet |= idMask
	r.id = id

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *DeleteJob) ErrorTrace(errortrace bool) *DeleteJob {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *DeleteJob) FilterPath(filterpaths ...string) *DeleteJob {
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
func (r *DeleteJob) Human(human bool) *DeleteJob {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *DeleteJob) Pretty(pretty bool) *DeleteJob {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}
