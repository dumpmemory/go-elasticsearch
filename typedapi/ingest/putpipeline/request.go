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

package putpipeline

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

// Request holds the request body struct for the package putpipeline
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/ingest/put_pipeline/PutPipelineRequest.ts#L25-L90
type Request struct {

	// Deprecated Marks this ingest pipeline as deprecated.
	// When a deprecated ingest pipeline is referenced as the default or final
	// pipeline when creating or updating a non-deprecated index template,
	// Elasticsearch will emit a deprecation warning.
	Deprecated *bool `json:"deprecated,omitempty"`
	// Description Description of the ingest pipeline.
	Description *string `json:"description,omitempty"`
	// Meta_ Optional metadata about the ingest pipeline. May have any contents. This map
	// is not automatically generated by Elasticsearch.
	Meta_ types.Metadata `json:"_meta,omitempty"`
	// OnFailure Processors to run immediately after a processor failure. Each processor
	// supports a processor-level `on_failure` value. If a processor without an
	// `on_failure` value fails, Elasticsearch uses this pipeline-level parameter as
	// a fallback. The processors in this parameter run sequentially in the order
	// specified. Elasticsearch will not attempt to run the pipeline's remaining
	// processors.
	OnFailure []types.ProcessorContainer `json:"on_failure,omitempty"`
	// Processors Processors used to perform transformations on documents before indexing.
	// Processors run sequentially in the order specified.
	Processors []types.ProcessorContainer `json:"processors,omitempty"`
	// Version Version number used by external systems to track ingest pipelines. This
	// parameter is intended for external systems only. Elasticsearch does not use
	// or validate pipeline version numbers.
	Version *int64 `json:"version,omitempty"`
}

// NewRequest returns a Request
func NewRequest() *Request {
	r := &Request{}

	return r
}

// FromJSON allows to load an arbitrary json into the request structure
func (r *Request) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Putpipeline request: %w", err)
	}

	return &req, nil
}

func (s *Request) UnmarshalJSON(data []byte) error {
	dec := json.NewDecoder(bytes.NewReader(data))

	for {
		t, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		switch t {

		case "deprecated":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Deprecated", err)
				}
				s.Deprecated = &value
			case bool:
				s.Deprecated = &v
			}

		case "description":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Description", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Description = &o

		case "_meta":
			if err := dec.Decode(&s.Meta_); err != nil {
				return fmt.Errorf("%s | %w", "Meta_", err)
			}

		case "on_failure":
			if err := dec.Decode(&s.OnFailure); err != nil {
				return fmt.Errorf("%s | %w", "OnFailure", err)
			}

		case "processors":
			if err := dec.Decode(&s.Processors); err != nil {
				return fmt.Errorf("%s | %w", "Processors", err)
			}

		case "version":
			if err := dec.Decode(&s.Version); err != nil {
				return fmt.Errorf("%s | %w", "Version", err)
			}

		}
	}
	return nil
}
