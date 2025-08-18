// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// ResponseCountMarvisActionsResult represents a ResponseCountMarvisActionsResult struct.
type ResponseCountMarvisActionsResult struct {
	Count                *int              `json:"count,omitempty"`
	AdditionalProperties map[string]string `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseCountMarvisActionsResult,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseCountMarvisActionsResult) String() string {
	return fmt.Sprintf(
		"ResponseCountMarvisActionsResult[Count=%v, AdditionalProperties=%v]",
		r.Count, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseCountMarvisActionsResult.
// It customizes the JSON marshaling process for ResponseCountMarvisActionsResult objects.
func (r ResponseCountMarvisActionsResult) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(r.AdditionalProperties,
		"count"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(r.toMap())
}

// toMap converts the ResponseCountMarvisActionsResult object to a map representation for JSON marshaling.
func (r ResponseCountMarvisActionsResult) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, r.AdditionalProperties)
	if r.Count != nil {
		structMap["count"] = r.Count
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseCountMarvisActionsResult.
// It customizes the JSON unmarshaling process for ResponseCountMarvisActionsResult objects.
func (r *ResponseCountMarvisActionsResult) UnmarshalJSON(input []byte) error {
	var temp tempResponseCountMarvisActionsResult
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[string](input, "count")
	if err != nil {
		return err
	}
	r.AdditionalProperties = additionalProperties

	r.Count = temp.Count
	return nil
}

// tempResponseCountMarvisActionsResult is a temporary struct used for validating the fields of ResponseCountMarvisActionsResult.
type tempResponseCountMarvisActionsResult struct {
	Count *int `json:"count,omitempty"`
}
