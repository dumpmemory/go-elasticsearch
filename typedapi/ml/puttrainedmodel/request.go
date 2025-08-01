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

package puttrainedmodel

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/trainedmodeltype"
)

// Request holds the request body struct for the package puttrainedmodel
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/ml/put_trained_model/MlPutTrainedModelRequest.ts#L31-L135
type Request struct {

	// CompressedDefinition The compressed (GZipped and Base64 encoded) inference definition of the
	// model. If compressed_definition is specified, then definition cannot be
	// specified.
	CompressedDefinition *string `json:"compressed_definition,omitempty"`
	// Definition The inference definition for the model. If definition is specified, then
	// compressed_definition cannot be specified.
	Definition *types.Definition `json:"definition,omitempty"`
	// Description A human-readable description of the inference trained model.
	Description *string `json:"description,omitempty"`
	// InferenceConfig The default configuration for inference. This can be either a regression
	// or classification configuration. It must match the underlying
	// definition.trained_model's target_type. For pre-packaged models such as
	// ELSER the config is not required.
	InferenceConfig *types.InferenceConfigCreateContainer `json:"inference_config,omitempty"`
	// Input The input field names for the model definition.
	Input *types.Input `json:"input,omitempty"`
	// Metadata An object map that contains metadata about the model.
	Metadata json.RawMessage `json:"metadata,omitempty"`
	// ModelSizeBytes The estimated memory usage in bytes to keep the trained model in memory.
	// This property is supported only if defer_definition_decompression is true
	// or the model definition is not supplied.
	ModelSizeBytes *int64 `json:"model_size_bytes,omitempty"`
	// ModelType The model type.
	ModelType *trainedmodeltype.TrainedModelType `json:"model_type,omitempty"`
	// PlatformArchitecture The platform architecture (if applicable) of the trained mode. If the model
	// only works on one platform, because it is heavily optimized for a particular
	// processor architecture and OS combination, then this field specifies which.
	// The format of the string must match the platform identifiers used by
	// Elasticsearch,
	// so one of, `linux-x86_64`, `linux-aarch64`, `darwin-x86_64`,
	// `darwin-aarch64`,
	// or `windows-x86_64`. For portable models (those that work independent of
	// processor
	// architecture or OS features), leave this field unset.
	PlatformArchitecture *string `json:"platform_architecture,omitempty"`
	// PrefixStrings Optional prefix strings applied at inference
	PrefixStrings *types.TrainedModelPrefixStrings `json:"prefix_strings,omitempty"`
	// Tags An array of tags to organize the model.
	Tags []string `json:"tags,omitempty"`
}

// NewRequest returns a Request
func NewRequest() *Request {
	r := &Request{}

	return r
}

// FromJSON allows to load an arbitrary json into the request structure
func (r *Request) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Puttrainedmodel request: %w", err)
	}

	return &req, nil
}
