// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// ResponseAutoOrientationDevice represents a ResponseAutoOrientationDevice struct.
type ResponseAutoOrientationDevice struct {
	// Provides the reason for the status if the AP is invalid.
	Reason *string `json:"reason,omitempty"`
	// Indicates whether the auto orient request is valid for the device.
	Valid                *bool                  `json:"valid,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseAutoOrientationDevice,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseAutoOrientationDevice) String() string {
	return fmt.Sprintf(
		"ResponseAutoOrientationDevice[Reason=%v, Valid=%v, AdditionalProperties=%v]",
		r.Reason, r.Valid, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseAutoOrientationDevice.
// It customizes the JSON marshaling process for ResponseAutoOrientationDevice objects.
func (r ResponseAutoOrientationDevice) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(r.AdditionalProperties,
		"reason", "valid"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(r.toMap())
}

// toMap converts the ResponseAutoOrientationDevice object to a map representation for JSON marshaling.
func (r ResponseAutoOrientationDevice) toMap() map[string]any {
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

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseAutoOrientationDevice.
// It customizes the JSON unmarshaling process for ResponseAutoOrientationDevice objects.
func (r *ResponseAutoOrientationDevice) UnmarshalJSON(input []byte) error {
	var temp tempResponseAutoOrientationDevice
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

// tempResponseAutoOrientationDevice is a temporary struct used for validating the fields of ResponseAutoOrientationDevice.
type tempResponseAutoOrientationDevice struct {
	Reason *string `json:"reason,omitempty"`
	Valid  *bool   `json:"valid,omitempty"`
}
