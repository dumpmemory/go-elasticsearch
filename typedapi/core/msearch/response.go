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

package msearch

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

// Response holds the response body struct for the package msearch
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/_global/msearch/MultiSearchResponse.ts#L25-L28
type Response struct {
	Responses []types.MsearchResponseItem `json:"responses"`
	Took      int64                       `json:"took"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{}
	return r
}

func (s *Response) UnmarshalJSON(data []byte) error {
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

		case "responses":
			messageArray := []json.RawMessage{}
			if err := dec.Decode(&messageArray); err != nil {
				return fmt.Errorf("%s | %w", "Responses", err)
			}
		responses_field:
			for _, message := range messageArray {
				keyDec := json.NewDecoder(bytes.NewReader(message))
				for {
					t, err := keyDec.Token()
					if err != nil {
						if errors.Is(err, io.EOF) {
							break
						}
						return fmt.Errorf("%s | %w", "Responses", err)
					}

					switch t {

					case "aggregations", "_clusters", "fields", "hits", "max_score", "num_reduce_phases", "pit_id", "profile", "_scroll_id", "_shards", "suggest", "terminated_early", "timed_out", "took":
						o := types.NewMultiSearchItem()
						localDec := json.NewDecoder(bytes.NewReader(message))
						if err := localDec.Decode(&o); err != nil {
							return fmt.Errorf("%s | %w", "Responses", err)
						}
						s.Responses = append(s.Responses, o)
						continue responses_field

					case "error":
						o := types.NewErrorResponseBase()
						localDec := json.NewDecoder(bytes.NewReader(message))
						if err := localDec.Decode(&o); err != nil {
							return fmt.Errorf("%s | %w", "Responses", err)
						}
						s.Responses = append(s.Responses, o)
						continue responses_field

					}
				}

				var o any
				localDec := json.NewDecoder(bytes.NewReader(message))
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "Responses", err)
				}
				s.Responses = append(s.Responses, o)
			}

		case "took":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Took", err)
				}
				s.Took = value
			case float64:
				f := int64(v)
				s.Took = f
			}

		}
	}
	return nil
}
