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

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// AttachmentProcessor type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/ingest/_types/Processors.ts#L346-L387
type AttachmentProcessor struct {
	// Description Description of the processor.
	// Useful for describing the purpose of the processor or its configuration.
	Description *string `json:"description,omitempty"`
	// Field The field to get the base64 encoded field from.
	Field string `json:"field"`
	// If Conditionally execute the processor.
	If *Script `json:"if,omitempty"`
	// IgnoreFailure Ignore failures for the processor.
	IgnoreFailure *bool `json:"ignore_failure,omitempty"`
	// IgnoreMissing If `true` and field does not exist, the processor quietly exits without
	// modifying the document.
	IgnoreMissing *bool `json:"ignore_missing,omitempty"`
	// IndexedChars The number of chars being used for extraction to prevent huge fields.
	// Use `-1` for no limit.
	IndexedChars *int64 `json:"indexed_chars,omitempty"`
	// IndexedCharsField Field name from which you can overwrite the number of chars being used for
	// extraction.
	IndexedCharsField *string `json:"indexed_chars_field,omitempty"`
	// OnFailure Handle failures for the processor.
	OnFailure []ProcessorContainer `json:"on_failure,omitempty"`
	// Properties Array of properties to select to be stored.
	// Can be `content`, `title`, `name`, `author`, `keywords`, `date`,
	// `content_type`, `content_length`, `language`.
	Properties []string `json:"properties,omitempty"`
	// RemoveBinary If true, the binary field will be removed from the document
	RemoveBinary *bool `json:"remove_binary,omitempty"`
	// ResourceName Field containing the name of the resource to decode.
	// If specified, the processor passes this resource name to the underlying Tika
	// library to enable Resource Name Based Detection.
	ResourceName *string `json:"resource_name,omitempty"`
	// Tag Identifier for the processor.
	// Useful for debugging and metrics.
	Tag *string `json:"tag,omitempty"`
	// TargetField The field that will hold the attachment information.
	TargetField *string `json:"target_field,omitempty"`
}

func (s *AttachmentProcessor) UnmarshalJSON(data []byte) error {

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

		case "field":
			if err := dec.Decode(&s.Field); err != nil {
				return fmt.Errorf("%s | %w", "Field", err)
			}

		case "if":
			if err := dec.Decode(&s.If); err != nil {
				return fmt.Errorf("%s | %w", "If", err)
			}

		case "ignore_failure":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "IgnoreFailure", err)
				}
				s.IgnoreFailure = &value
			case bool:
				s.IgnoreFailure = &v
			}

		case "ignore_missing":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "IgnoreMissing", err)
				}
				s.IgnoreMissing = &value
			case bool:
				s.IgnoreMissing = &v
			}

		case "indexed_chars":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "IndexedChars", err)
				}
				s.IndexedChars = &value
			case float64:
				f := int64(v)
				s.IndexedChars = &f
			}

		case "indexed_chars_field":
			if err := dec.Decode(&s.IndexedCharsField); err != nil {
				return fmt.Errorf("%s | %w", "IndexedCharsField", err)
			}

		case "on_failure":
			if err := dec.Decode(&s.OnFailure); err != nil {
				return fmt.Errorf("%s | %w", "OnFailure", err)
			}

		case "properties":
			if err := dec.Decode(&s.Properties); err != nil {
				return fmt.Errorf("%s | %w", "Properties", err)
			}

		case "remove_binary":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "RemoveBinary", err)
				}
				s.RemoveBinary = &value
			case bool:
				s.RemoveBinary = &v
			}

		case "resource_name":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "ResourceName", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ResourceName = &o

		case "tag":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Tag", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Tag = &o

		case "target_field":
			if err := dec.Decode(&s.TargetField); err != nil {
				return fmt.Errorf("%s | %w", "TargetField", err)
			}

		}
	}
	return nil
}

// NewAttachmentProcessor returns a AttachmentProcessor.
func NewAttachmentProcessor() *AttachmentProcessor {
	r := &AttachmentProcessor{}

	return r
}

type AttachmentProcessorVariant interface {
	AttachmentProcessorCaster() *AttachmentProcessor
}

func (s *AttachmentProcessor) AttachmentProcessorCaster() *AttachmentProcessor {
	return s
}
