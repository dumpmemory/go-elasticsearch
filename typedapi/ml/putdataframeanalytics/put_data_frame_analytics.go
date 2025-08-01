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

// Create a data frame analytics job.
// This API creates a data frame analytics job that performs an analysis on the
// source indices and stores the outcome in a destination index.
// By default, the query used in the source configuration is `{"match_all":
// {}}`.
//
// If the destination index does not exist, it is created automatically when you
// start the job.
//
// If you supply only a subset of the regression or classification parameters,
// hyperparameter optimization occurs. It determines a value for each of the
// undefined parameters.
package putdataframeanalytics

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
)

const (
	idMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutDataFrameAnalytics struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	id string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewPutDataFrameAnalytics type alias for index.
type NewPutDataFrameAnalytics func(id string) *PutDataFrameAnalytics

// NewPutDataFrameAnalyticsFunc returns a new instance of PutDataFrameAnalytics with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewPutDataFrameAnalyticsFunc(tp elastictransport.Interface) NewPutDataFrameAnalytics {
	return func(id string) *PutDataFrameAnalytics {
		n := New(tp)

		n._id(id)

		return n
	}
}

// Create a data frame analytics job.
// This API creates a data frame analytics job that performs an analysis on the
// source indices and stores the outcome in a destination index.
// By default, the query used in the source configuration is `{"match_all":
// {}}`.
//
// If the destination index does not exist, it is created automatically when you
// start the job.
//
// If you supply only a subset of the regression or classification parameters,
// hyperparameter optimization occurs. It determines a value for each of the
// undefined parameters.
//
// https://www.elastic.co/docs/api/doc/elasticsearch/operation/operation-ml-put-data-frame-analytics
func New(tp elastictransport.Interface) *PutDataFrameAnalytics {
	r := &PutDataFrameAnalytics{
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
func (r *PutDataFrameAnalytics) Raw(raw io.Reader) *PutDataFrameAnalytics {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *PutDataFrameAnalytics) Request(req *Request) *PutDataFrameAnalytics {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *PutDataFrameAnalytics) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for PutDataFrameAnalytics: %w", err)
		}

		r.buf.Write(data)

	}

	if r.buf.Len() > 0 {
		r.raw = r.buf
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == idMask:
		path.WriteString("/")
		path.WriteString("_ml")
		path.WriteString("/")
		path.WriteString("data_frame")
		path.WriteString("/")
		path.WriteString("analytics")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "id", r.id)
		}
		path.WriteString(r.id)

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
func (r PutDataFrameAnalytics) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "ml.put_data_frame_analytics")
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
		instrument.BeforeRequest(req, "ml.put_data_frame_analytics")
		if reader := instrument.RecordRequestBody(ctx, "ml.put_data_frame_analytics", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "ml.put_data_frame_analytics")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the PutDataFrameAnalytics query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a putdataframeanalytics.Response
func (r PutDataFrameAnalytics) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "ml.put_data_frame_analytics")
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

// Header set a key, value pair in the PutDataFrameAnalytics headers map.
func (r *PutDataFrameAnalytics) Header(key, value string) *PutDataFrameAnalytics {
	r.headers.Set(key, value)

	return r
}

// Id Identifier for the data frame analytics job. This identifier can contain
// lowercase alphanumeric characters (a-z and 0-9), hyphens, and
// underscores. It must start and end with alphanumeric characters.
// API Name: id
func (r *PutDataFrameAnalytics) _id(id string) *PutDataFrameAnalytics {
	r.paramSet |= idMask
	r.id = id

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *PutDataFrameAnalytics) ErrorTrace(errortrace bool) *PutDataFrameAnalytics {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *PutDataFrameAnalytics) FilterPath(filterpaths ...string) *PutDataFrameAnalytics {
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
func (r *PutDataFrameAnalytics) Human(human bool) *PutDataFrameAnalytics {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *PutDataFrameAnalytics) Pretty(pretty bool) *PutDataFrameAnalytics {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// Specifies whether this job can start when there is insufficient machine
// learning node capacity for it to be immediately assigned to a node. If
// set to `false` and a machine learning node with capacity to run the job
// cannot be immediately found, the API returns an error. If set to `true`,
// the API does not return an error; the job waits in the `starting` state
// until sufficient machine learning node capacity is available. This
// behavior is also affected by the cluster-wide
// `xpack.ml.max_lazy_ml_nodes` setting.
// API name: allow_lazy_start
func (r *PutDataFrameAnalytics) AllowLazyStart(allowlazystart bool) *PutDataFrameAnalytics {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.AllowLazyStart = &allowlazystart

	return r
}

// The analysis configuration, which contains the information necessary to
// perform one of the following types of analysis: classification, outlier
// detection, or regression.
// API name: analysis
func (r *PutDataFrameAnalytics) Analysis(analysis types.DataframeAnalysisContainerVariant) *PutDataFrameAnalytics {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Analysis = *analysis.DataframeAnalysisContainerCaster()

	return r
}

// Specifies `includes` and/or `excludes` patterns to select which fields
// will be included in the analysis. The patterns specified in `excludes`
// are applied last, therefore `excludes` takes precedence. In other words,
// if the same field is specified in both `includes` and `excludes`, then
// the field will not be included in the analysis. If `analyzed_fields` is
// not set, only the relevant fields will be included. For example, all the
// numeric fields for outlier detection.
// The supported fields vary for each type of analysis. Outlier detection
// requires numeric or `boolean` data to analyze. The algorithms don’t
// support missing values therefore fields that have data types other than
// numeric or boolean are ignored. Documents where included fields contain
// missing values, null values, or an array are also ignored. Therefore the
// `dest` index may contain documents that don’t have an outlier score.
// Regression supports fields that are numeric, `boolean`, `text`,
// `keyword`, and `ip` data types. It is also tolerant of missing values.
// Fields that are supported are included in the analysis, other fields are
// ignored. Documents where included fields contain an array with two or
// more values are also ignored. Documents in the `dest` index that don’t
// contain a results field are not included in the regression analysis.
// Classification supports fields that are numeric, `boolean`, `text`,
// `keyword`, and `ip` data types. It is also tolerant of missing values.
// Fields that are supported are included in the analysis, other fields are
// ignored. Documents where included fields contain an array with two or
// more values are also ignored. Documents in the `dest` index that don’t
// contain a results field are not included in the classification analysis.
// Classification analysis can be improved by mapping ordinal variable
// values to a single number. For example, in case of age ranges, you can
// model the values as `0-14 = 0`, `15-24 = 1`, `25-34 = 2`, and so on.
// API name: analyzed_fields
func (r *PutDataFrameAnalytics) AnalyzedFields(analyzedfields types.DataframeAnalysisAnalyzedFieldsVariant) *PutDataFrameAnalytics {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.AnalyzedFields = analyzedfields.DataframeAnalysisAnalyzedFieldsCaster()

	return r
}

// A description of the job.
// API name: description
func (r *PutDataFrameAnalytics) Description(description string) *PutDataFrameAnalytics {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Description = &description

	return r
}

// The destination configuration.
// API name: dest
func (r *PutDataFrameAnalytics) Dest(dest types.DataframeAnalyticsDestinationVariant) *PutDataFrameAnalytics {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Dest = *dest.DataframeAnalyticsDestinationCaster()

	return r
}

// API name: headers
func (r *PutDataFrameAnalytics) Headers(httpheaders types.HttpHeadersVariant) *PutDataFrameAnalytics {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Headers = *httpheaders.HttpHeadersCaster()

	return r
}

// The maximum number of threads to be used by the analysis. Using more
// threads may decrease the time necessary to complete the analysis at the
// cost of using more CPU. Note that the process may use additional threads
// for operational functionality other than the analysis itself.
// API name: max_num_threads
func (r *PutDataFrameAnalytics) MaxNumThreads(maxnumthreads int) *PutDataFrameAnalytics {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.MaxNumThreads = &maxnumthreads

	return r
}

// API name: _meta
func (r *PutDataFrameAnalytics) Meta_(metadata types.MetadataVariant) *PutDataFrameAnalytics {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Meta_ = *metadata.MetadataCaster()

	return r
}

// The approximate maximum amount of memory resources that are permitted for
// analytical processing. If your `elasticsearch.yml` file contains an
// `xpack.ml.max_model_memory_limit` setting, an error occurs when you try
// to create data frame analytics jobs that have `model_memory_limit` values
// greater than that setting.
// API name: model_memory_limit
func (r *PutDataFrameAnalytics) ModelMemoryLimit(modelmemorylimit string) *PutDataFrameAnalytics {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.ModelMemoryLimit = &modelmemorylimit

	return r
}

// The configuration of how to source the analysis data.
// API name: source
func (r *PutDataFrameAnalytics) Source(source types.DataframeAnalyticsSourceVariant) *PutDataFrameAnalytics {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Source = *source.DataframeAnalyticsSourceCaster()

	return r
}

// API name: version
func (r *PutDataFrameAnalytics) Version(versionstring string) *PutDataFrameAnalytics {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Version = &versionstring

	return r
}
