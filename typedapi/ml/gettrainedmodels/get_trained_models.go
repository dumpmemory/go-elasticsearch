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

// Get trained model configuration info.
package gettrainedmodels

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
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/include"
)

const (
	modelidMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetTrainedModels struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	modelid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewGetTrainedModels type alias for index.
type NewGetTrainedModels func() *GetTrainedModels

// NewGetTrainedModelsFunc returns a new instance of GetTrainedModels with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewGetTrainedModelsFunc(tp elastictransport.Interface) NewGetTrainedModels {
	return func() *GetTrainedModels {
		n := New(tp)

		return n
	}
}

// Get trained model configuration info.
//
// https://www.elastic.co/docs/api/doc/elasticsearch/operation/operation-ml-get-trained-models
func New(tp elastictransport.Interface) *GetTrainedModels {
	r := &GetTrainedModels{
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
func (r *GetTrainedModels) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == modelidMask:
		path.WriteString("/")
		path.WriteString("_ml")
		path.WriteString("/")
		path.WriteString("trained_models")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "modelid", r.modelid)
		}
		path.WriteString(r.modelid)

		method = http.MethodGet
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_ml")
		path.WriteString("/")
		path.WriteString("trained_models")

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
func (r GetTrainedModels) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "ml.get_trained_models")
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
		instrument.BeforeRequest(req, "ml.get_trained_models")
		if reader := instrument.RecordRequestBody(ctx, "ml.get_trained_models", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "ml.get_trained_models")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the GetTrainedModels query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a gettrainedmodels.Response
func (r GetTrainedModels) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "ml.get_trained_models")
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
func (r GetTrainedModels) IsSuccess(providedCtx context.Context) (bool, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "ml.get_trained_models")
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
		err := fmt.Errorf("an error happened during the GetTrainedModels query execution, status code: %d", res.StatusCode)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return false, err
	}

	return false, nil
}

// Header set a key, value pair in the GetTrainedModels headers map.
func (r *GetTrainedModels) Header(key, value string) *GetTrainedModels {
	r.headers.Set(key, value)

	return r
}

// ModelId The unique identifier of the trained model or a model alias.
//
// You can get information for multiple trained models in a single API
// request by using a comma-separated list of model IDs or a wildcard
// expression.
// API Name: modelid
func (r *GetTrainedModels) ModelId(modelid string) *GetTrainedModels {
	r.paramSet |= modelidMask
	r.modelid = modelid

	return r
}

// AllowNoMatch Specifies what to do when the request:
//
// - Contains wildcard expressions and there are no models that match.
// - Contains the _all string or no identifiers and there are no matches.
// - Contains wildcard expressions and there are only partial matches.
//
// If true, it returns an empty array when there are no matches and the
// subset of results when there are partial matches.
// API name: allow_no_match
func (r *GetTrainedModels) AllowNoMatch(allownomatch bool) *GetTrainedModels {
	r.values.Set("allow_no_match", strconv.FormatBool(allownomatch))

	return r
}

// DecompressDefinition Specifies whether the included model definition should be returned as a
// JSON map (true) or in a custom compressed format (false).
// API name: decompress_definition
func (r *GetTrainedModels) DecompressDefinition(decompressdefinition bool) *GetTrainedModels {
	r.values.Set("decompress_definition", strconv.FormatBool(decompressdefinition))

	return r
}

// ExcludeGenerated Indicates if certain fields should be removed from the configuration on
// retrieval. This allows the configuration to be in an acceptable format to
// be retrieved and then added to another cluster.
// API name: exclude_generated
func (r *GetTrainedModels) ExcludeGenerated(excludegenerated bool) *GetTrainedModels {
	r.values.Set("exclude_generated", strconv.FormatBool(excludegenerated))

	return r
}

// From Skips the specified number of models.
// API name: from
func (r *GetTrainedModels) From(from int) *GetTrainedModels {
	r.values.Set("from", strconv.Itoa(from))

	return r
}

// Include A comma delimited string of optional fields to include in the response
// body.
// API name: include
func (r *GetTrainedModels) Include(include include.Include) *GetTrainedModels {
	r.values.Set("include", include.String())

	return r
}

// Size Specifies the maximum number of models to obtain.
// API name: size
func (r *GetTrainedModels) Size(size int) *GetTrainedModels {
	r.values.Set("size", strconv.Itoa(size))

	return r
}

// Tags A comma delimited string of tags. A trained model can have many tags, or
// none. When supplied, only trained models that contain all the supplied
// tags are returned.
// API name: tags
func (r *GetTrainedModels) Tags(tags ...string) *GetTrainedModels {
	tmp := []string{}
	for _, item := range tags {
		tmp = append(tmp, fmt.Sprintf("%v", item))
	}
	r.values.Set("tags", strings.Join(tmp, ","))

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *GetTrainedModels) ErrorTrace(errortrace bool) *GetTrainedModels {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *GetTrainedModels) FilterPath(filterpaths ...string) *GetTrainedModels {
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
func (r *GetTrainedModels) Human(human bool) *GetTrainedModels {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *GetTrainedModels) Pretty(pretty bool) *GetTrainedModels {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}
