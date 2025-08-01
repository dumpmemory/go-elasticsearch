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

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/segmentsortmissing"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/segmentsortmode"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/segmentsortorder"
)

// IndexSegmentSort type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/indices/_types/IndexSegmentSort.ts#L22-L27
type IndexSegmentSort struct {
	Field   []string                                `json:"field,omitempty"`
	Missing []segmentsortmissing.SegmentSortMissing `json:"missing,omitempty"`
	Mode    []segmentsortmode.SegmentSortMode       `json:"mode,omitempty"`
	Order   []segmentsortorder.SegmentSortOrder     `json:"order,omitempty"`
}

func (s *IndexSegmentSort) UnmarshalJSON(data []byte) error {

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

		case "field":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new(string)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "Field", err)
				}

				s.Field = append(s.Field, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.Field); err != nil {
					return fmt.Errorf("%s | %w", "Field", err)
				}
			}

		case "missing":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := &segmentsortmissing.SegmentSortMissing{}
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "Missing", err)
				}

				s.Missing = append(s.Missing, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.Missing); err != nil {
					return fmt.Errorf("%s | %w", "Missing", err)
				}
			}

		case "mode":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := &segmentsortmode.SegmentSortMode{}
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "Mode", err)
				}

				s.Mode = append(s.Mode, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.Mode); err != nil {
					return fmt.Errorf("%s | %w", "Mode", err)
				}
			}

		case "order":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := &segmentsortorder.SegmentSortOrder{}
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "Order", err)
				}

				s.Order = append(s.Order, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.Order); err != nil {
					return fmt.Errorf("%s | %w", "Order", err)
				}
			}

		}
	}
	return nil
}

// NewIndexSegmentSort returns a IndexSegmentSort.
func NewIndexSegmentSort() *IndexSegmentSort {
	r := &IndexSegmentSort{}

	return r
}

type IndexSegmentSortVariant interface {
	IndexSegmentSortCaster() *IndexSegmentSort
}

func (s *IndexSegmentSort) IndexSegmentSortCaster() *IndexSegmentSort {
	return s
}
