package models

import (
	"encoding/json"
)

// ResponseLoginSuccess represents a ResponseLoginSuccess struct.
type ResponseLoginSuccess struct {
	Email                *string        `json:"email,omitempty"`
	TwoFactorPassed      *bool          `json:"two_factor_passed,omitempty"`
	TwoFactorRequired    *bool          `json:"two_factor_required,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseLoginSuccess.
// It customizes the JSON marshaling process for ResponseLoginSuccess objects.
func (r ResponseLoginSuccess) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the ResponseLoginSuccess object to a map representation for JSON marshaling.
func (r ResponseLoginSuccess) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, r.AdditionalProperties)
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
	additionalProperties, err := UnmarshalAdditionalProperties(input, "email", "two_factor_passed", "two_factor_required")
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
