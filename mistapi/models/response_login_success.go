// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// ResponseLoginSuccess represents a ResponseLoginSuccess struct.
type ResponseLoginSuccess struct {
	Email                *string                `json:"email,omitempty"`
	TwoFactorPassed      *bool                  `json:"two_factor_passed,omitempty"`
	TwoFactorRequired    *bool                  `json:"two_factor_required,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseLoginSuccess,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseLoginSuccess) String() string {
	return fmt.Sprintf(
		"ResponseLoginSuccess[Email=%v, TwoFactorPassed=%v, TwoFactorRequired=%v, AdditionalProperties=%v]",
		r.Email, r.TwoFactorPassed, r.TwoFactorRequired, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseLoginSuccess.
// It customizes the JSON marshaling process for ResponseLoginSuccess objects.
func (r ResponseLoginSuccess) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(r.AdditionalProperties,
		"email", "two_factor_passed", "two_factor_required"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(r.toMap())
}

// toMap converts the ResponseLoginSuccess object to a map representation for JSON marshaling.
func (r ResponseLoginSuccess) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, r.AdditionalProperties)
	if r.Email != nil {
		structMap["email"] = r.Email
	}
	if r.TwoFactorPassed != nil {
		structMap["two_factor_passed"] = r.TwoFactorPassed
	}
	if r.TwoFactorRequired != nil {
		structMap["two_factor_required"] = r.TwoFactorRequired
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseLoginSuccess.
// It customizes the JSON unmarshaling process for ResponseLoginSuccess objects.
func (r *ResponseLoginSuccess) UnmarshalJSON(input []byte) error {
	var temp tempResponseLoginSuccess
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "email", "two_factor_passed", "two_factor_required")
	if err != nil {
		return err
	}
	r.AdditionalProperties = additionalProperties

	r.Email = temp.Email
	r.TwoFactorPassed = temp.TwoFactorPassed
	r.TwoFactorRequired = temp.TwoFactorRequired
	return nil
}

// tempResponseLoginSuccess is a temporary struct used for validating the fields of ResponseLoginSuccess.
type tempResponseLoginSuccess struct {
	Email             *string `json:"email,omitempty"`
	TwoFactorPassed   *bool   `json:"two_factor_passed,omitempty"`
	TwoFactorRequired *bool   `json:"two_factor_required,omitempty"`
}
