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

// Create an Anthropic inference endpoint.
//
// Create an inference endpoint to perform an inference task with the
// `anthropic` service.
package putanthropic

import (
	gobytes "bytes"
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
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/anthropicservicetype"
)

const (
	tasktypeMask = iota + 1

	anthropicinferenceidMask
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutAnthropic struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	tasktype             string
	anthropicinferenceid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewPutAnthropic type alias for index.
type NewPutAnthropic func(tasktype, anthropicinferenceid string) *PutAnthropic

// NewPutAnthropicFunc returns a new instance of PutAnthropic with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewPutAnthropicFunc(tp elastictransport.Interface) NewPutAnthropic {
	return func(tasktype, anthropicinferenceid string) *PutAnthropic {
		n := New(tp)

		n._tasktype(tasktype)

		n._anthropicinferenceid(anthropicinferenceid)

		return n
	}
}

// Create an Anthropic inference endpoint.
//
// Create an inference endpoint to perform an inference task with the
// `anthropic` service.
//
// https://www.elastic.co/docs/api/doc/elasticsearch/operation/operation-inference-put-anthropic
func New(tp elastictransport.Interface) *PutAnthropic {
	r := &PutAnthropic{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),

		buf: gobytes.NewBuffer(nil),
	}

	if instrumented, ok := r.transport.(elastictransport.Instrumented); ok {
		if instrument := instrumented.InstrumentationEnabled(); instrument != nil {
			r.instrument = instrument
		}
	}

	return r
}

// Raw takes a json payload as input which is then passed to the http.Request
// If specified Raw takes precedence on Request method.
func (r *PutAnthropic) Raw(raw io.Reader) *PutAnthropic {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *PutAnthropic) Request(req *Request) *PutAnthropic {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *PutAnthropic) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	if len(r.deferred) > 0 {
		for _, f := range r.deferred {
			deferredErr := f(r.req)
			if deferredErr != nil {
				return nil, deferredErr
			}
		}
	}

	if r.raw == nil && r.req != nil {

		data, err := json.Marshal(r.req)

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for PutAnthropic: %w", err)
		}

		r.buf.Write(data)

	}

	if r.buf.Len() > 0 {
		r.raw = r.buf
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == tasktypeMask|anthropicinferenceidMask:
		path.WriteString("/")
		path.WriteString("_inference")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "tasktype", r.tasktype)
		}
		path.WriteString(r.tasktype)
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "anthropicinferenceid", r.anthropicinferenceid)
		}
		path.WriteString(r.anthropicinferenceid)

		method = http.MethodPut
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

	if req.Header.Get("Content-Type") == "" {
		if r.raw != nil {
			req.Header.Set("Content-Type", "application/vnd.elasticsearch+json;compatible-with=9")
		}
	}

	if req.Header.Get("Accept") == "" {
		req.Header.Set("Accept", "application/vnd.elasticsearch+json;compatible-with=9")
	}

	if err != nil {
		return req, fmt.Errorf("could not build http.Request: %w", err)
	}

	return req, nil
}

// Perform runs the http.Request through the provided transport and returns an http.Response.
func (r PutAnthropic) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "inference.put_anthropic")
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
		instrument.BeforeRequest(req, "inference.put_anthropic")
		if reader := instrument.RecordRequestBody(ctx, "inference.put_anthropic", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "inference.put_anthropic")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the PutAnthropic query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a putanthropic.Response
func (r PutAnthropic) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "inference.put_anthropic")
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

// Header set a key, value pair in the PutAnthropic headers map.
func (r *PutAnthropic) Header(key, value string) *PutAnthropic {
	r.headers.Set(key, value)

	return r
}

// TaskType The task type.
// The only valid task type for the model to perform is `completion`.
// API Name: tasktype
func (r *PutAnthropic) _tasktype(tasktype string) *PutAnthropic {
	r.paramSet |= tasktypeMask
	r.tasktype = tasktype

	return r
}

// AnthropicInferenceId The unique identifier of the inference endpoint.
// API Name: anthropicinferenceid
func (r *PutAnthropic) _anthropicinferenceid(anthropicinferenceid string) *PutAnthropic {
	r.paramSet |= anthropicinferenceidMask
	r.anthropicinferenceid = anthropicinferenceid

	return r
}

// Timeout Specifies the amount of time to wait for the inference endpoint to be
// created.
// API name: timeout
func (r *PutAnthropic) Timeout(duration string) *PutAnthropic {
	r.values.Set("timeout", duration)

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *PutAnthropic) ErrorTrace(errortrace bool) *PutAnthropic {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *PutAnthropic) FilterPath(filterpaths ...string) *PutAnthropic {
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
func (r *PutAnthropic) Human(human bool) *PutAnthropic {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *PutAnthropic) Pretty(pretty bool) *PutAnthropic {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// The chunking configuration object.
// API name: chunking_settings
func (r *PutAnthropic) ChunkingSettings(chunkingsettings types.InferenceChunkingSettingsVariant) *PutAnthropic {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.ChunkingSettings = chunkingsettings.InferenceChunkingSettingsCaster()

	return r
}

// The type of service supported for the specified task type. In this case,
// `anthropic`.
// API name: service
func (r *PutAnthropic) Service(service anthropicservicetype.AnthropicServiceType) *PutAnthropic {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Service = service
	return r
}

// Settings used to install the inference model. These settings are specific to
// the `watsonxai` service.
// API name: service_settings
func (r *PutAnthropic) ServiceSettings(servicesettings types.AnthropicServiceSettingsVariant) *PutAnthropic {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.ServiceSettings = *servicesettings.AnthropicServiceSettingsCaster()

	return r
}

// Settings to configure the inference task.
// These settings are specific to the task type you specified.
// API name: task_settings
func (r *PutAnthropic) TaskSettings(tasksettings types.AnthropicTaskSettingsVariant) *PutAnthropic {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.TaskSettings = tasksettings.AnthropicTaskSettingsCaster()

	return r
}
