// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// ResponseAutoMapAssignmentDevice represents a ResponseAutoMapAssignmentDevice struct.
// Per-device validation result for auto map assignment
type ResponseAutoMapAssignmentDevice struct {
	// Provides the reason for the status if the AP is invalid
	Reason *string `json:"reason,omitempty"`
	// Indicates whether the device meets requirements for auto map assignment
	Valid                *bool                  `json:"valid,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseAutoMapAssignmentDevice,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseAutoMapAssignmentDevice) String() string {
	return fmt.Sprintf(
		"ResponseAutoMapAssignmentDevice[Reason=%v, Valid=%v, AdditionalProperties=%v]",
		r.Reason, r.Valid, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseAutoMapAssignmentDevice.
// It customizes the JSON marshaling process for ResponseAutoMapAssignmentDevice objects.
func (r ResponseAutoMapAssignmentDevice) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(r.AdditionalProperties,
		"reason", "valid"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(r.toMap())
}

// toMap converts the ResponseAutoMapAssignmentDevice object to a map representation for JSON marshaling.
func (r ResponseAutoMapAssignmentDevice) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, r.AdditionalProperties)
	if r.Reason != nil {
		structMap["reason"] = r.Reason
	}
	if r.Valid != nil {
		structMap["valid"] = r.Valid
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseAutoMapAssignmentDevice.
// It customizes the JSON unmarshaling process for ResponseAutoMapAssignmentDevice objects.
func (r *ResponseAutoMapAssignmentDevice) UnmarshalJSON(input []byte) error {
	var temp tempResponseAutoMapAssignmentDevice
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "reason", "valid")
	if err != nil {
		return err
	}
	r.AdditionalProperties = additionalProperties

	r.Reason = temp.Reason
	r.Valid = temp.Valid
	return nil
}

// tempResponseAutoMapAssignmentDevice is a temporary struct used for validating the fields of ResponseAutoMapAssignmentDevice.
type tempResponseAutoMapAssignmentDevice struct {
	Reason *string `json:"reason,omitempty"`
	Valid  *bool   `json:"valid,omitempty"`
}
