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

// Create a datafeed.
// Datafeeds retrieve data from Elasticsearch for analysis by an anomaly
// detection job.
// You can associate only one datafeed with each anomaly detection job.
// The datafeed contains a query that runs at a defined interval (`frequency`).
// If you are concerned about delayed data, you can add a delay (`query_delay')
// at each interval.
// By default, the datafeed uses the following query: `{"match_all": {"boost":
// 1}}`.
//
// When Elasticsearch security features are enabled, your datafeed remembers
// which roles the user who created it had
// at the time of creation and runs the query using those same roles. If you
// provide secondary authorization headers,
// those credentials are used instead.
// You must use Kibana, this API, or the create anomaly detection jobs API to
// create a datafeed. Do not add a datafeed
// directly to the `.ml-config` index. Do not give users `write` privileges on
// the `.ml-config` index.
package putdatafeed

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
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/expandwildcard"
)

const (
	datafeedidMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutDatafeed struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	datafeedid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewPutDatafeed type alias for index.
type NewPutDatafeed func(datafeedid string) *PutDatafeed

// NewPutDatafeedFunc returns a new instance of PutDatafeed with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewPutDatafeedFunc(tp elastictransport.Interface) NewPutDatafeed {
	return func(datafeedid string) *PutDatafeed {
		n := New(tp)

		n._datafeedid(datafeedid)

		return n
	}
}

// Create a datafeed.
// Datafeeds retrieve data from Elasticsearch for analysis by an anomaly
// detection job.
// You can associate only one datafeed with each anomaly detection job.
// The datafeed contains a query that runs at a defined interval (`frequency`).
// If you are concerned about delayed data, you can add a delay (`query_delay')
// at each interval.
// By default, the datafeed uses the following query: `{"match_all": {"boost":
// 1}}`.
//
// When Elasticsearch security features are enabled, your datafeed remembers
// which roles the user who created it had
// at the time of creation and runs the query using those same roles. If you
// provide secondary authorization headers,
// those credentials are used instead.
// You must use Kibana, this API, or the create anomaly detection jobs API to
// create a datafeed. Do not add a datafeed
// directly to the `.ml-config` index. Do not give users `write` privileges on
// the `.ml-config` index.
//
// https://www.elastic.co/docs/api/doc/elasticsearch/operation/operation-ml-put-datafeed
func New(tp elastictransport.Interface) *PutDatafeed {
	r := &PutDatafeed{
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
func (r *PutDatafeed) Raw(raw io.Reader) *PutDatafeed {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *PutDatafeed) Request(req *Request) *PutDatafeed {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *PutDatafeed) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for PutDatafeed: %w", err)
		}

		r.buf.Write(data)

	}

	if r.buf.Len() > 0 {
		r.raw = r.buf
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == datafeedidMask:
		path.WriteString("/")
		path.WriteString("_ml")
		path.WriteString("/")
		path.WriteString("datafeeds")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "datafeedid", r.datafeedid)
		}
		path.WriteString(r.datafeedid)

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
func (r PutDatafeed) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "ml.put_datafeed")
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
		instrument.BeforeRequest(req, "ml.put_datafeed")
		if reader := instrument.RecordRequestBody(ctx, "ml.put_datafeed", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "ml.put_datafeed")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the PutDatafeed query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a putdatafeed.Response
func (r PutDatafeed) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "ml.put_datafeed")
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

// Header set a key, value pair in the PutDatafeed headers map.
func (r *PutDatafeed) Header(key, value string) *PutDatafeed {
	r.headers.Set(key, value)

	return r
}

// DatafeedId A numerical character string that uniquely identifies the datafeed.
// This identifier can contain lowercase alphanumeric characters (a-z and 0-9),
// hyphens, and underscores.
// It must start and end with alphanumeric characters.
// API Name: datafeedid
func (r *PutDatafeed) _datafeedid(datafeedid string) *PutDatafeed {
	r.paramSet |= datafeedidMask
	r.datafeedid = datafeedid

	return r
}

// AllowNoIndices If true, wildcard indices expressions that resolve into no concrete indices
// are ignored. This includes the `_all`
// string or when no indices are specified.
// API name: allow_no_indices
func (r *PutDatafeed) AllowNoIndices(allownoindices bool) *PutDatafeed {
	r.values.Set("allow_no_indices", strconv.FormatBool(allownoindices))

	return r
}

// ExpandWildcards Type of index that wildcard patterns can match. If the request can target
// data streams, this argument determines
// whether wildcard expressions match hidden data streams. Supports
// comma-separated values.
// API name: expand_wildcards
func (r *PutDatafeed) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *PutDatafeed {
	tmp := []string{}
	for _, item := range expandwildcards {
		tmp = append(tmp, item.String())
	}
	r.values.Set("expand_wildcards", strings.Join(tmp, ","))

	return r
}

// IgnoreThrottled If true, concrete, expanded, or aliased indices are ignored when frozen.
// API name: ignore_throttled
func (r *PutDatafeed) IgnoreThrottled(ignorethrottled bool) *PutDatafeed {
	r.values.Set("ignore_throttled", strconv.FormatBool(ignorethrottled))

	return r
}

// IgnoreUnavailable If true, unavailable indices (missing or closed) are ignored.
// API name: ignore_unavailable
func (r *PutDatafeed) IgnoreUnavailable(ignoreunavailable bool) *PutDatafeed {
	r.values.Set("ignore_unavailable", strconv.FormatBool(ignoreunavailable))

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *PutDatafeed) ErrorTrace(errortrace bool) *PutDatafeed {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *PutDatafeed) FilterPath(filterpaths ...string) *PutDatafeed {
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
func (r *PutDatafeed) Human(human bool) *PutDatafeed {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *PutDatafeed) Pretty(pretty bool) *PutDatafeed {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// If set, the datafeed performs aggregation searches.
// Support for aggregations is limited and should be used only with low
// cardinality data.
// API name: aggregations
func (r *PutDatafeed) Aggregations(aggregations map[string]types.Aggregations) *PutDatafeed {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Aggregations = aggregations
	return r
}

func (r *PutDatafeed) AddAggregation(key string, value types.AggregationsVariant) *PutDatafeed {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	var tmp map[string]types.Aggregations
	if r.req.Aggregations == nil {
		r.req.Aggregations = make(map[string]types.Aggregations)
	} else {
		tmp = r.req.Aggregations
	}

	tmp[key] = *value.AggregationsCaster()

	r.req.Aggregations = tmp
	return r
}

// Datafeeds might be required to search over long time periods, for several
// months or years.
// This search is split into time chunks in order to ensure the load on
// Elasticsearch is managed.
// Chunking configuration controls how the size of these time chunks are
// calculated;
// it is an advanced configuration option.
// API name: chunking_config
func (r *PutDatafeed) ChunkingConfig(chunkingconfig types.ChunkingConfigVariant) *PutDatafeed {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.ChunkingConfig = chunkingconfig.ChunkingConfigCaster()

	return r
}

// Specifies whether the datafeed checks for missing data and the size of the
// window.
// The datafeed can optionally search over indices that have already been read
// in an effort to determine whether
// any data has subsequently been added to the index. If missing data is found,
// it is a good indication that the
// `query_delay` is set too low and the data is being indexed after the datafeed
// has passed that moment in time.
// This check runs only on real-time datafeeds.
// API name: delayed_data_check_config
func (r *PutDatafeed) DelayedDataCheckConfig(delayeddatacheckconfig types.DelayedDataCheckConfigVariant) *PutDatafeed {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.DelayedDataCheckConfig = delayeddatacheckconfig.DelayedDataCheckConfigCaster()

	return r
}

// The interval at which scheduled queries are made while the datafeed runs in
// real time.
// The default value is either the bucket span for short bucket spans, or, for
// longer bucket spans, a sensible
// fraction of the bucket span. When `frequency` is shorter than the bucket
// span, interim results for the last
// (partial) bucket are written then eventually overwritten by the full bucket
// results. If the datafeed uses
// aggregations, this value must be divisible by the interval of the date
// histogram aggregation.
// API name: frequency
func (r *PutDatafeed) Frequency(duration types.DurationVariant) *PutDatafeed {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Frequency = *duration.DurationCaster()

	return r
}

// API name: headers
func (r *PutDatafeed) Headers(httpheaders types.HttpHeadersVariant) *PutDatafeed {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Headers = *httpheaders.HttpHeadersCaster()

	return r
}

// An array of index names. Wildcards are supported. If any of the indices are
// in remote clusters, the master
// nodes and the machine learning nodes must have the `remote_cluster_client`
// role.
// API name: indices
func (r *PutDatafeed) Indices(indices ...string) *PutDatafeed {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Indices = indices

	return r
}

// Specifies index expansion options that are used during search
// API name: indices_options
func (r *PutDatafeed) IndicesOptions(indicesoptions types.IndicesOptionsVariant) *PutDatafeed {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.IndicesOptions = indicesoptions.IndicesOptionsCaster()

	return r
}

// Identifier for the anomaly detection job.
// API name: job_id
func (r *PutDatafeed) JobId(id string) *PutDatafeed {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.JobId = &id

	return r
}

// If a real-time datafeed has never seen any data (including during any initial
// training period), it automatically
// stops and closes the associated job after this many real-time searches return
// no documents. In other words,
// it stops after `frequency` times `max_empty_searches` of real-time operation.
// If not set, a datafeed with no
// end time that sees no data remains started until it is explicitly stopped. By
// default, it is not set.
// API name: max_empty_searches
func (r *PutDatafeed) MaxEmptySearches(maxemptysearches int) *PutDatafeed {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.MaxEmptySearches = &maxemptysearches

	return r
}

// The Elasticsearch query domain-specific language (DSL). This value
// corresponds to the query object in an
// Elasticsearch search POST body. All the options that are supported by
// Elasticsearch can be used, as this
// object is passed verbatim to Elasticsearch.
// API name: query
func (r *PutDatafeed) Query(query types.QueryVariant) *PutDatafeed {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Query = query.QueryCaster()

	return r
}

// The number of seconds behind real time that data is queried. For example, if
// data from 10:04 a.m. might
// not be searchable in Elasticsearch until 10:06 a.m., set this property to 120
// seconds. The default
// value is randomly selected between `60s` and `120s`. This randomness improves
// the query performance
// when there are multiple jobs running on the same node.
// API name: query_delay
func (r *PutDatafeed) QueryDelay(duration types.DurationVariant) *PutDatafeed {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.QueryDelay = *duration.DurationCaster()

	return r
}

// Specifies runtime fields for the datafeed search.
// API name: runtime_mappings
func (r *PutDatafeed) RuntimeMappings(runtimefields types.RuntimeFieldsVariant) *PutDatafeed {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.RuntimeMappings = *runtimefields.RuntimeFieldsCaster()

	return r
}

// Specifies scripts that evaluate custom expressions and returns script fields
// to the datafeed.
// The detector configuration objects in a job can contain functions that use
// these script fields.
// API name: script_fields
func (r *PutDatafeed) ScriptFields(scriptfields map[string]types.ScriptField) *PutDatafeed {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.ScriptFields = scriptfields
	return r
}

func (r *PutDatafeed) AddScriptField(key string, value types.ScriptFieldVariant) *PutDatafeed {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	var tmp map[string]types.ScriptField
	if r.req.ScriptFields == nil {
		r.req.ScriptFields = make(map[string]types.ScriptField)
	} else {
		tmp = r.req.ScriptFields
	}

	tmp[key] = *value.ScriptFieldCaster()

	r.req.ScriptFields = tmp
	return r
}

// The size parameter that is used in Elasticsearch searches when the datafeed
// does not use aggregations.
// The maximum value is the value of `index.max_result_window`, which is 10,000
// by default.
// API name: scroll_size
func (r *PutDatafeed) ScrollSize(scrollsize int) *PutDatafeed {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.ScrollSize = &scrollsize

	return r
}
