// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// MxedgeVmParams represents a MxedgeVmParams struct.
// Mist Edge VM parameters
type MxedgeVmParams struct {
	// Mist Edge VM SKU or model to deploy
	Model *string `json:"model,omitempty"`
	// User given name (optional)
	Name *string `json:"name,omitempty"`
	// Base64 encoded user data
	UserData             *string                `json:"user_data,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for MxedgeVmParams,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m MxedgeVmParams) String() string {
	return fmt.Sprintf(
		"MxedgeVmParams[Model=%v, Name=%v, UserData=%v, AdditionalProperties=%v]",
		m.Model, m.Name, m.UserData, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for MxedgeVmParams.
// It customizes the JSON marshaling process for MxedgeVmParams objects.
func (m MxedgeVmParams) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(m.AdditionalProperties,
		"model", "name", "user_data"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(m.toMap())
}

// toMap converts the MxedgeVmParams object to a map representation for JSON marshaling.
func (m MxedgeVmParams) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, m.AdditionalProperties)
	if m.Model != nil {
		structMap["model"] = m.Model
	}
	if m.Name != nil {
		structMap["name"] = m.Name
	}
	if m.UserData != nil {
		structMap["user_data"] = m.UserData
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MxedgeVmParams.
// It customizes the JSON unmarshaling process for MxedgeVmParams objects.
func (m *MxedgeVmParams) UnmarshalJSON(input []byte) error {
	var temp tempMxedgeVmParams
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "model", "name", "user_data")
	if err != nil {
		return err
	}
	m.AdditionalProperties = additionalProperties

	m.Model = temp.Model
	m.Name = temp.Name
	m.UserData = temp.UserData
	return nil
}

// tempMxedgeVmParams is a temporary struct used for validating the fields of MxedgeVmParams.
type tempMxedgeVmParams struct {
	Model    *string `json:"model,omitempty"`
	Name     *string `json:"name,omitempty"`
	UserData *string `json:"user_data,omitempty"`
}
