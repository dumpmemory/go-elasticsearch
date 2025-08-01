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

// HttpEmailAttachment type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/watcher/_types/Actions.ts#L218-L222
type HttpEmailAttachment struct {
	ContentType *string                     `json:"content_type,omitempty"`
	Inline      *bool                       `json:"inline,omitempty"`
	Request     *HttpInputRequestDefinition `json:"request,omitempty"`
}

func (s *HttpEmailAttachment) UnmarshalJSON(data []byte) error {

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

		case "content_type":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "ContentType", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ContentType = &o

		case "inline":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Inline", err)
				}
				s.Inline = &value
			case bool:
				s.Inline = &v
			}

		case "request":
			if err := dec.Decode(&s.Request); err != nil {
				return fmt.Errorf("%s | %w", "Request", err)
			}

		}
	}
	return nil
}

// NewHttpEmailAttachment returns a HttpEmailAttachment.
func NewHttpEmailAttachment() *HttpEmailAttachment {
	r := &HttpEmailAttachment{}

	return r
}

type HttpEmailAttachmentVariant interface {
	HttpEmailAttachmentCaster() *HttpEmailAttachment
}

func (s *HttpEmailAttachment) HttpEmailAttachmentCaster() *HttpEmailAttachment {
	return s
}
