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
// https://github.com/elastic/elasticsearch-specification/tree/c6ef5fbc736f1dd6256c2babc92e07bf150cadb9

// Create an JinaAI inference endpoint.
//
// Create an inference endpoint to perform an inference task with the `jinaai`
// service.
//
// To review the available `rerank` models, refer to <https://jina.ai/reranker>.
// To review the available `text_embedding` models, refer to the
// <https://jina.ai/embeddings/>.
//
// When you create an inference endpoint, the associated machine learning model
// is automatically deployed if it is not already running.
// After creating the endpoint, wait for the model deployment to complete before
// using it.
// To verify the deployment status, use the get trained model statistics API.
// Look for `"state": "fully_allocated"` in the response and ensure that the
// `"allocation_count"` matches the `"target_allocation_count"`.
// Avoid creating multiple endpoints for the same model unless required, as each
// endpoint consumes significant resources.
package putjinaai

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
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/jinaaiservicetype"
)

const (
	tasktypeMask = iota + 1

	jinaaiinferenceidMask
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutJinaai struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	tasktype          string
	jinaaiinferenceid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewPutJinaai type alias for index.
type NewPutJinaai func(tasktype, jinaaiinferenceid string) *PutJinaai

// NewPutJinaaiFunc returns a new instance of PutJinaai with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewPutJinaaiFunc(tp elastictransport.Interface) NewPutJinaai {
	return func(tasktype, jinaaiinferenceid string) *PutJinaai {
		n := New(tp)

		n._tasktype(tasktype)

		n._jinaaiinferenceid(jinaaiinferenceid)

		return n
	}
}

// Create an JinaAI inference endpoint.
//
// Create an inference endpoint to perform an inference task with the `jinaai`
// service.
//
// To review the available `rerank` models, refer to <https://jina.ai/reranker>.
// To review the available `text_embedding` models, refer to the
// <https://jina.ai/embeddings/>.
//
// When you create an inference endpoint, the associated machine learning model
// is automatically deployed if it is not already running.
// After creating the endpoint, wait for the model deployment to complete before
// using it.
// To verify the deployment status, use the get trained model statistics API.
// Look for `"state": "fully_allocated"` in the response and ensure that the
// `"allocation_count"` matches the `"target_allocation_count"`.
// Avoid creating multiple endpoints for the same model unless required, as each
// endpoint consumes significant resources.
//
// https://www.elastic.co/docs/api/doc/elasticsearch/operation/operation-inference-put-jinaai
func New(tp elastictransport.Interface) *PutJinaai {
	r := &PutJinaai{
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
func (r *PutJinaai) Raw(raw io.Reader) *PutJinaai {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *PutJinaai) Request(req *Request) *PutJinaai {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *PutJinaai) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for PutJinaai: %w", err)
		}

		r.buf.Write(data)

	}

	if r.buf.Len() > 0 {
		r.raw = r.buf
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == tasktypeMask|jinaaiinferenceidMask:
		path.WriteString("/")
		path.WriteString("_inference")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "tasktype", r.tasktype)
		}
		path.WriteString(r.tasktype)
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "jinaaiinferenceid", r.jinaaiinferenceid)
		}
		path.WriteString(r.jinaaiinferenceid)

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
func (r PutJinaai) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "inference.put_jinaai")
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
		instrument.BeforeRequest(req, "inference.put_jinaai")
		if reader := instrument.RecordRequestBody(ctx, "inference.put_jinaai", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "inference.put_jinaai")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the PutJinaai query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a putjinaai.Response
func (r PutJinaai) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "inference.put_jinaai")
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

// Header set a key, value pair in the PutJinaai headers map.
func (r *PutJinaai) Header(key, value string) *PutJinaai {
	r.headers.Set(key, value)

	return r
}

// TaskType The type of the inference task that the model will perform.
// API Name: tasktype
func (r *PutJinaai) _tasktype(tasktype string) *PutJinaai {
	r.paramSet |= tasktypeMask
	r.tasktype = tasktype

	return r
}

// JinaaiInferenceId The unique identifier of the inference endpoint.
// API Name: jinaaiinferenceid
func (r *PutJinaai) _jinaaiinferenceid(jinaaiinferenceid string) *PutJinaai {
	r.paramSet |= jinaaiinferenceidMask
	r.jinaaiinferenceid = jinaaiinferenceid

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *PutJinaai) ErrorTrace(errortrace bool) *PutJinaai {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *PutJinaai) FilterPath(filterpaths ...string) *PutJinaai {
	tmp := []string{}
	for _, item := range filterpaths {
		tmp = append(tmp, fmt.Sprintf("%v", item))
	}
	r.values.Set("filter_path", strings.Join(tmp, ","))

	return r
}

// Human When set to `true` will return statistics in a format suitable for humans.
// For example `"exists_time": "1h"` for humans and
// `"eixsts_time_in_millis": 3600000` for computers. When disabled the human
// readable values will be omitted. This makes sense for responses being
// consumed
// only by machines.
// API name: human
func (r *PutJinaai) Human(human bool) *PutJinaai {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *PutJinaai) Pretty(pretty bool) *PutJinaai {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// The chunking configuration object.
// API name: chunking_settings
func (r *PutJinaai) ChunkingSettings(chunkingsettings types.InferenceChunkingSettingsVariant) *PutJinaai {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.ChunkingSettings = chunkingsettings.InferenceChunkingSettingsCaster()

	return r
}

// The type of service supported for the specified task type. In this case,
// `jinaai`.
// API name: service
func (r *PutJinaai) Service(service jinaaiservicetype.JinaAIServiceType) *PutJinaai {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Service = service
	return r
}

// Settings used to install the inference model. These settings are specific to
// the `jinaai` service.
// API name: service_settings
func (r *PutJinaai) ServiceSettings(servicesettings types.JinaAIServiceSettingsVariant) *PutJinaai {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.ServiceSettings = *servicesettings.JinaAIServiceSettingsCaster()

	return r
}

// Settings to configure the inference task.
// These settings are specific to the task type you specified.
// API name: task_settings
func (r *PutJinaai) TaskSettings(tasksettings types.JinaAITaskSettingsVariant) *PutJinaai {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.TaskSettings = tasksettings.JinaAITaskSettingsCaster()

	return r
}
