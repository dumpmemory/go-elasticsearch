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

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// IngestPipeline type.
//
// https://github.com/elastic/elasticsearch-specification/blob/c6ef5fbc736f1dd6256c2babc92e07bf150cadb9/specification/ingest/_types/Pipeline.ts#L23-L51
type IngestPipeline struct {
	// Deprecated Marks this ingest pipeline as deprecated.
	// When a deprecated ingest pipeline is referenced as the default or final
	// pipeline when creating or updating a non-deprecated index template,
	// Elasticsearch will emit a deprecation warning.
	Deprecated *bool `json:"deprecated,omitempty"`
	// Description Description of the ingest pipeline.
	Description *string `json:"description,omitempty"`
	// Meta_ Arbitrary metadata about the ingest pipeline. This map is not automatically
	// generated by Elasticsearch.
	Meta_ Metadata `json:"_meta,omitempty"`
	// OnFailure Processors to run immediately after a processor failure.
	OnFailure []ProcessorContainer `json:"on_failure,omitempty"`
	// Processors Processors used to perform transformations on documents before indexing.
	// Processors run sequentially in the order specified.
	Processors []ProcessorContainer `json:"processors,omitempty"`
	// Version Version number used by external systems to track ingest pipelines.
	Version *int64 `json:"version,omitempty"`
}

func (s *IngestPipeline) UnmarshalJSON(data []byte) error {

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

// NewIngestPipeline returns a IngestPipeline.
func NewIngestPipeline() *IngestPipeline {
	r := &IngestPipeline{}

	return r
}

// true

type IngestPipelineVariant interface {
	IngestPipelineCaster() *IngestPipeline
}

func (s *IngestPipeline) IngestPipelineCaster() *IngestPipeline {
	return s
}
