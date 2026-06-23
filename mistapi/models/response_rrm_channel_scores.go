// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// ResponseRrmChannelScores represents a ResponseRrmChannelScores struct.
// Response containing RRM channel score records
type ResponseRrmChannelScores struct {
	// RRM channel score records returned for a band
	Results              []RrmChannelScore      `json:"results"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseRrmChannelScores,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseRrmChannelScores) String() string {
	return fmt.Sprintf(
		"ResponseRrmChannelScores[Results=%v, AdditionalProperties=%v]",
		r.Results, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseRrmChannelScores.
// It customizes the JSON marshaling process for ResponseRrmChannelScores objects.
func (r ResponseRrmChannelScores) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(r.AdditionalProperties,
		"results"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(r.toMap())
}

// toMap converts the ResponseRrmChannelScores object to a map representation for JSON marshaling.
func (r ResponseRrmChannelScores) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, r.AdditionalProperties)
	structMap["results"] = r.Results
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseRrmChannelScores.
// It customizes the JSON unmarshaling process for ResponseRrmChannelScores objects.
func (r *ResponseRrmChannelScores) UnmarshalJSON(input []byte) error {
	var temp tempResponseRrmChannelScores
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "results")
	if err != nil {
		return err
	}
	r.AdditionalProperties = additionalProperties

	r.Results = *temp.Results
	return nil
}

// tempResponseRrmChannelScores is a temporary struct used for validating the fields of ResponseRrmChannelScores.
type tempResponseRrmChannelScores struct {
	Results *[]RrmChannelScore `json:"results"`
}

func (r *tempResponseRrmChannelScores) validate() error {
	var errs []string
	if r.Results == nil {
		errs = append(errs, "required field `results` is missing for type `response_rrm_channel_scores`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
