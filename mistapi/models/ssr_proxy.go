// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// SsrProxy represents a SsrProxy struct.
// SSR proxy configuration to talk to Mist
type SsrProxy struct {
	Disabled             *bool                  `json:"disabled,omitempty"`
	Url                  *string                `json:"url,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SsrProxy,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SsrProxy) String() string {
	return fmt.Sprintf(
		"SsrProxy[Disabled=%v, Url=%v, AdditionalProperties=%v]",
		s.Disabled, s.Url, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SsrProxy.
// It customizes the JSON marshaling process for SsrProxy objects.
func (s SsrProxy) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"disabled", "url"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SsrProxy object to a map representation for JSON marshaling.
func (s SsrProxy) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.Disabled != nil {
		structMap["disabled"] = s.Disabled
	}
	if s.Url != nil {
		structMap["url"] = s.Url
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SsrProxy.
// It customizes the JSON unmarshaling process for SsrProxy objects.
func (s *SsrProxy) UnmarshalJSON(input []byte) error {
	var temp tempSsrProxy
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "disabled", "url")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.Disabled = temp.Disabled
	s.Url = temp.Url
	return nil
}

// tempSsrProxy is a temporary struct used for validating the fields of SsrProxy.
type tempSsrProxy struct {
	Disabled *bool   `json:"disabled,omitempty"`
	Url      *string `json:"url,omitempty"`
}
