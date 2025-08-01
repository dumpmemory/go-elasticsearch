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

// KeyValueProcessor type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/ingest/_types/Processors.ts#L1177-L1229
type KeyValueProcessor struct {
	// Description Description of the processor.
	// Useful for describing the purpose of the processor or its configuration.
	Description *string `json:"description,omitempty"`
	// ExcludeKeys List of keys to exclude from document.
	ExcludeKeys []string `json:"exclude_keys,omitempty"`
	// Field The field to be parsed.
	// Supports template snippets.
	Field string `json:"field"`
	// FieldSplit Regex pattern to use for splitting key-value pairs.
	FieldSplit string `json:"field_split"`
	// If Conditionally execute the processor.
	If *Script `json:"if,omitempty"`
	// IgnoreFailure Ignore failures for the processor.
	IgnoreFailure *bool `json:"ignore_failure,omitempty"`
	// IgnoreMissing If `true` and `field` does not exist or is `null`, the processor quietly
	// exits without modifying the document.
	IgnoreMissing *bool `json:"ignore_missing,omitempty"`
	// IncludeKeys List of keys to filter and insert into document.
	// Defaults to including all keys.
	IncludeKeys []string `json:"include_keys,omitempty"`
	// OnFailure Handle failures for the processor.
	OnFailure []ProcessorContainer `json:"on_failure,omitempty"`
	// Prefix Prefix to be added to extracted keys.
	Prefix *string `json:"prefix,omitempty"`
	// StripBrackets If `true`. strip brackets `()`, `<>`, `[]` as well as quotes `'` and `"` from
	// extracted values.
	StripBrackets *bool `json:"strip_brackets,omitempty"`
	// Tag Identifier for the processor.
	// Useful for debugging and metrics.
	Tag *string `json:"tag,omitempty"`
	// TargetField The field to insert the extracted keys into.
	// Defaults to the root of the document.
	// Supports template snippets.
	TargetField *string `json:"target_field,omitempty"`
	// TrimKey String of characters to trim from extracted keys.
	TrimKey *string `json:"trim_key,omitempty"`
	// TrimValue String of characters to trim from extracted values.
	TrimValue *string `json:"trim_value,omitempty"`
	// ValueSplit Regex pattern to use for splitting the key from the value within a key-value
	// pair.
	ValueSplit string `json:"value_split"`
}

func (s *KeyValueProcessor) UnmarshalJSON(data []byte) error {

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

		case "exclude_keys":
			if err := dec.Decode(&s.ExcludeKeys); err != nil {
				return fmt.Errorf("%s | %w", "ExcludeKeys", err)
			}

		case "field":
			if err := dec.Decode(&s.Field); err != nil {
				return fmt.Errorf("%s | %w", "Field", err)
			}

		case "field_split":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "FieldSplit", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.FieldSplit = o

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

		case "include_keys":
			if err := dec.Decode(&s.IncludeKeys); err != nil {
				return fmt.Errorf("%s | %w", "IncludeKeys", err)
			}

		case "on_failure":
			if err := dec.Decode(&s.OnFailure); err != nil {
				return fmt.Errorf("%s | %w", "OnFailure", err)
			}

		case "prefix":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Prefix", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Prefix = &o

		case "strip_brackets":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "StripBrackets", err)
				}
				s.StripBrackets = &value
			case bool:
				s.StripBrackets = &v
			}

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

		case "trim_key":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "TrimKey", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.TrimKey = &o

		case "trim_value":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "TrimValue", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.TrimValue = &o

		case "value_split":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "ValueSplit", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ValueSplit = o

		}
	}
	return nil
}

// NewKeyValueProcessor returns a KeyValueProcessor.
func NewKeyValueProcessor() *KeyValueProcessor {
	r := &KeyValueProcessor{}

	return r
}

type KeyValueProcessorVariant interface {
	KeyValueProcessorCaster() *KeyValueProcessor
}

func (s *KeyValueProcessor) KeyValueProcessorCaster() *KeyValueProcessor {
	return s
}
