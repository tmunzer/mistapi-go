// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// ResponseOrgSuppressAlarm represents a ResponseOrgSuppressAlarm struct.
type ResponseOrgSuppressAlarm struct {
	Results              []ResponseOrgSuppressAlarmItem `json:"results,omitempty"`
	AdditionalProperties map[string]interface{}         `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseOrgSuppressAlarm,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseOrgSuppressAlarm) String() string {
	return fmt.Sprintf(
		"ResponseOrgSuppressAlarm[Results=%v, AdditionalProperties=%v]",
		r.Results, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseOrgSuppressAlarm.
// It customizes the JSON marshaling process for ResponseOrgSuppressAlarm objects.
func (r ResponseOrgSuppressAlarm) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(r.AdditionalProperties,
		"results"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(r.toMap())
}

// toMap converts the ResponseOrgSuppressAlarm object to a map representation for JSON marshaling.
func (r ResponseOrgSuppressAlarm) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, r.AdditionalProperties)
	if r.Results != nil {
		structMap["results"] = r.Results
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseOrgSuppressAlarm.
// It customizes the JSON unmarshaling process for ResponseOrgSuppressAlarm objects.
func (r *ResponseOrgSuppressAlarm) UnmarshalJSON(input []byte) error {
	var temp tempResponseOrgSuppressAlarm
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "results")
	if err != nil {
		return err
	}
	r.AdditionalProperties = additionalProperties

	r.Results = temp.Results
	return nil
}

// tempResponseOrgSuppressAlarm is a temporary struct used for validating the fields of ResponseOrgSuppressAlarm.
type tempResponseOrgSuppressAlarm struct {
	Results []ResponseOrgSuppressAlarmItem `json:"results,omitempty"`
}
