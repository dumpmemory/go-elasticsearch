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

// UserProfileUser type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/security/_types/UserProfile.ts#L32-L39
type UserProfileUser struct {
	Email       *string  `json:"email,omitempty"`
	FullName    *string  `json:"full_name,omitempty"`
	RealmDomain *string  `json:"realm_domain,omitempty"`
	RealmName   string   `json:"realm_name"`
	Roles       []string `json:"roles"`
	Username    string   `json:"username"`
}

func (s *UserProfileUser) UnmarshalJSON(data []byte) error {

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

		case "email":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Email", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Email = &o

		case "full_name":
			if err := dec.Decode(&s.FullName); err != nil {
				return fmt.Errorf("%s | %w", "FullName", err)
			}

		case "realm_domain":
			if err := dec.Decode(&s.RealmDomain); err != nil {
				return fmt.Errorf("%s | %w", "RealmDomain", err)
			}

		case "realm_name":
			if err := dec.Decode(&s.RealmName); err != nil {
				return fmt.Errorf("%s | %w", "RealmName", err)
			}

		case "roles":
			if err := dec.Decode(&s.Roles); err != nil {
				return fmt.Errorf("%s | %w", "Roles", err)
			}

		case "username":
			if err := dec.Decode(&s.Username); err != nil {
				return fmt.Errorf("%s | %w", "Username", err)
			}

		}
	}
	return nil
}

// NewUserProfileUser returns a UserProfileUser.
func NewUserProfileUser() *UserProfileUser {
	r := &UserProfileUser{}

	return r
}
