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

package authenticate

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

// Response holds the response body struct for the package authenticate
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/security/authenticate/SecurityAuthenticateResponse.ts#L24-L42
type Response struct {
	ApiKey              *types.AuthenticateApiKey `json:"api_key,omitempty"`
	AuthenticationRealm types.RealmInfo           `json:"authentication_realm"`
	AuthenticationType  string                    `json:"authentication_type"`
	Email               *string                   `json:"email,omitempty"`
	Enabled             bool                      `json:"enabled"`
	FullName            *string                   `json:"full_name,omitempty"`
	LookupRealm         types.RealmInfo           `json:"lookup_realm"`
	Metadata            types.Metadata            `json:"metadata"`
	Roles               []string                  `json:"roles"`
	Token               *types.AuthenticateToken  `json:"token,omitempty"`
	Username            string                    `json:"username"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{}
	return r
}
