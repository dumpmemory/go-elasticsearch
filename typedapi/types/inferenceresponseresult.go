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

// InferenceResponseResult type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/ml/_types/inference.ts#L478-L526
type InferenceResponseResult struct {
	// Entities If the model is trained for named entity recognition (NER) tasks, the
	// response contains the recognized entities.
	Entities []TrainedModelEntities `json:"entities,omitempty"`
	// FeatureImportance The feature importance for the inference results. Relevant only for
	// classification or regression models
	FeatureImportance []TrainedModelInferenceFeatureImportance `json:"feature_importance,omitempty"`
	// IsTruncated Indicates whether the input text was truncated to meet the model's maximum
	// sequence length limit. This property
	// is present only when it is true.
	IsTruncated *bool `json:"is_truncated,omitempty"`
	// PredictedValue If the model is trained for a text classification or zero shot classification
	// task, the response is the
	// predicted class.
	// For named entity recognition (NER) tasks, it contains the annotated text
	// output.
	// For fill mask tasks, it contains the top prediction for replacing the mask
	// token.
	// For text embedding tasks, it contains the raw numerical text embedding
	// values.
	// For regression models, its a numerical value
	// For classification models, it may be an integer, double, boolean or string
	// depending on prediction type
	PredictedValue [][]ScalarValue `json:"predicted_value,omitempty"`
	// PredictedValueSequence For fill mask tasks, the response contains the input text sequence with the
	// mask token replaced by the predicted
	// value.
	// Additionally
	PredictedValueSequence *string `json:"predicted_value_sequence,omitempty"`
	// PredictionProbability Specifies a probability for the predicted value.
	PredictionProbability *Float64 `json:"prediction_probability,omitempty"`
	// PredictionScore Specifies a confidence score for the predicted value.
	PredictionScore *Float64 `json:"prediction_score,omitempty"`
	// TopClasses For fill mask, text classification, and zero shot classification tasks, the
	// response contains a list of top
	// class entries.
	TopClasses []TopClassEntry `json:"top_classes,omitempty"`
	// Warning If the request failed, the response contains the reason for the failure.
	Warning *string `json:"warning,omitempty"`
}

func (s *InferenceResponseResult) UnmarshalJSON(data []byte) error {

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

		case "entities":
			if err := dec.Decode(&s.Entities); err != nil {
				return fmt.Errorf("%s | %w", "Entities", err)
			}

		case "feature_importance":
			if err := dec.Decode(&s.FeatureImportance); err != nil {
				return fmt.Errorf("%s | %w", "FeatureImportance", err)
			}

		case "is_truncated":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "IsTruncated", err)
				}
				s.IsTruncated = &value
			case bool:
				s.IsTruncated = &v
			}

		case "predicted_value":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new([]ScalarValue)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "PredictedValue", err)
				}

				s.PredictedValue = append(s.PredictedValue, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.PredictedValue); err != nil {
					return fmt.Errorf("%s | %w", "PredictedValue", err)
				}
			}

		case "predicted_value_sequence":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "PredictedValueSequence", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PredictedValueSequence = &o

		case "prediction_probability":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "PredictionProbability", err)
				}
				f := Float64(value)
				s.PredictionProbability = &f
			case float64:
				f := Float64(v)
				s.PredictionProbability = &f
			}

		case "prediction_score":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "PredictionScore", err)
				}
				f := Float64(value)
				s.PredictionScore = &f
			case float64:
				f := Float64(v)
				s.PredictionScore = &f
			}

		case "top_classes":
			if err := dec.Decode(&s.TopClasses); err != nil {
				return fmt.Errorf("%s | %w", "TopClasses", err)
			}

		case "warning":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Warning", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Warning = &o

		}
	}
	return nil
}

// NewInferenceResponseResult returns a InferenceResponseResult.
func NewInferenceResponseResult() *InferenceResponseResult {
	r := &InferenceResponseResult{}

	return r
}
