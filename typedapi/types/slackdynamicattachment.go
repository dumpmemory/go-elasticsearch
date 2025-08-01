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

// SlackDynamicAttachment type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/watcher/_types/Actions.ts#L125-L128
type SlackDynamicAttachment struct {
	AttachmentTemplate SlackAttachment `json:"attachment_template"`
	ListPath           string          `json:"list_path"`
}

func (s *SlackDynamicAttachment) UnmarshalJSON(data []byte) error {

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

		case "attachment_template":
			if err := dec.Decode(&s.AttachmentTemplate); err != nil {
				return fmt.Errorf("%s | %w", "AttachmentTemplate", err)
			}

		case "list_path":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "ListPath", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ListPath = o

		}
	}
	return nil
}

// NewSlackDynamicAttachment returns a SlackDynamicAttachment.
func NewSlackDynamicAttachment() *SlackDynamicAttachment {
	r := &SlackDynamicAttachment{}

	return r
}

type SlackDynamicAttachmentVariant interface {
	SlackDynamicAttachmentCaster() *SlackDynamicAttachment
}

func (s *SlackDynamicAttachment) SlackDynamicAttachmentCaster() *SlackDynamicAttachment {
	return s
}
