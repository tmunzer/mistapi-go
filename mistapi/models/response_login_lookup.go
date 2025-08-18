// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// ResponseLoginLookup represents a ResponseLoginLookup struct.
type ResponseLoginLookup struct {
	SsoUrl               *string                `json:"sso_url,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseLoginLookup,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseLoginLookup) String() string {
	return fmt.Sprintf(
		"ResponseLoginLookup[SsoUrl=%v, AdditionalProperties=%v]",
		r.SsoUrl, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseLoginLookup.
// It customizes the JSON marshaling process for ResponseLoginLookup objects.
func (r ResponseLoginLookup) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(r.AdditionalProperties,
		"sso_url"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(r.toMap())
}

// toMap converts the ResponseLoginLookup object to a map representation for JSON marshaling.
func (r ResponseLoginLookup) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, r.AdditionalProperties)
	if r.SsoUrl != nil {
		structMap["sso_url"] = r.SsoUrl
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseLoginLookup.
// It customizes the JSON unmarshaling process for ResponseLoginLookup objects.
func (r *ResponseLoginLookup) UnmarshalJSON(input []byte) error {
	var temp tempResponseLoginLookup
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "sso_url")
	if err != nil {
		return err
	}
	r.AdditionalProperties = additionalProperties

	r.SsoUrl = temp.SsoUrl
	return nil
}

// tempResponseLoginLookup is a temporary struct used for validating the fields of ResponseLoginLookup.
type tempResponseLoginLookup struct {
	SsoUrl *string `json:"sso_url,omitempty"`
}
