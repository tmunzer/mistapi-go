// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// EdgeconnectRegisterCmd represents a EdgeconnectRegisterCmd struct.
// EdgeConnect device registration command response
type EdgeconnectRegisterCmd struct {
	// Registration code used to adopt an EdgeConnect device into Mist
	RegistrationCode     *string                `json:"registration_code,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for EdgeconnectRegisterCmd,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (e EdgeconnectRegisterCmd) String() string {
	return fmt.Sprintf(
		"EdgeconnectRegisterCmd[RegistrationCode=%v, AdditionalProperties=%v]",
		e.RegistrationCode, e.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for EdgeconnectRegisterCmd.
// It customizes the JSON marshaling process for EdgeconnectRegisterCmd objects.
func (e EdgeconnectRegisterCmd) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(e.AdditionalProperties,
		"registration_code"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(e.toMap())
}

// toMap converts the EdgeconnectRegisterCmd object to a map representation for JSON marshaling.
func (e EdgeconnectRegisterCmd) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, e.AdditionalProperties)
	if e.RegistrationCode != nil {
		structMap["registration_code"] = e.RegistrationCode
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for EdgeconnectRegisterCmd.
// It customizes the JSON unmarshaling process for EdgeconnectRegisterCmd objects.
func (e *EdgeconnectRegisterCmd) UnmarshalJSON(input []byte) error {
	var temp tempEdgeconnectRegisterCmd
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "registration_code")
	if err != nil {
		return err
	}
	e.AdditionalProperties = additionalProperties

	e.RegistrationCode = temp.RegistrationCode
	return nil
}

// tempEdgeconnectRegisterCmd is a temporary struct used for validating the fields of EdgeconnectRegisterCmd.
type tempEdgeconnectRegisterCmd struct {
	RegistrationCode *string `json:"registration_code,omitempty"`
}
