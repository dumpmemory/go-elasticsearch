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

// JvmStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/ml/get_memory_stats/types.ts#L50-L63
type JvmStats struct {
	// HeapMax Maximum amount of memory available for use by the heap.
	HeapMax ByteSize `json:"heap_max,omitempty"`
	// HeapMaxInBytes Maximum amount of memory, in bytes, available for use by the heap.
	HeapMaxInBytes int `json:"heap_max_in_bytes"`
	// JavaInference Amount of Java heap currently being used for caching inference models.
	JavaInference ByteSize `json:"java_inference,omitempty"`
	// JavaInferenceInBytes Amount of Java heap, in bytes, currently being used for caching inference
	// models.
	JavaInferenceInBytes int `json:"java_inference_in_bytes"`
	// JavaInferenceMax Maximum amount of Java heap to be used for caching inference models.
	JavaInferenceMax ByteSize `json:"java_inference_max,omitempty"`
	// JavaInferenceMaxInBytes Maximum amount of Java heap, in bytes, to be used for caching inference
	// models.
	JavaInferenceMaxInBytes int `json:"java_inference_max_in_bytes"`
}

func (s *JvmStats) UnmarshalJSON(data []byte) error {

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

		case "heap_max":
			if err := dec.Decode(&s.HeapMax); err != nil {
				return fmt.Errorf("%s | %w", "HeapMax", err)
			}

		case "heap_max_in_bytes":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "HeapMaxInBytes", err)
				}
				s.HeapMaxInBytes = value
			case float64:
				f := int(v)
				s.HeapMaxInBytes = f
			}

		case "java_inference":
			if err := dec.Decode(&s.JavaInference); err != nil {
				return fmt.Errorf("%s | %w", "JavaInference", err)
			}

		case "java_inference_in_bytes":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "JavaInferenceInBytes", err)
				}
				s.JavaInferenceInBytes = value
			case float64:
				f := int(v)
				s.JavaInferenceInBytes = f
			}

		case "java_inference_max":
			if err := dec.Decode(&s.JavaInferenceMax); err != nil {
				return fmt.Errorf("%s | %w", "JavaInferenceMax", err)
			}

		case "java_inference_max_in_bytes":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "JavaInferenceMaxInBytes", err)
				}
				s.JavaInferenceMaxInBytes = value
			case float64:
				f := int(v)
				s.JavaInferenceMaxInBytes = f
			}

		}
	}
	return nil
}

// NewJvmStats returns a JvmStats.
func NewJvmStats() *JvmStats {
	r := &JvmStats{}

	return r
}
