// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// SsoOpenroaming represents a SsoOpenroaming struct.
// Deprecated. OpenRoaming configuration is now expressed as top-level fields on the SSO object: `openroaming_ssids`, `openroaming_wba_client_cert`, and `openroaming_wba_client_key`.
type SsoOpenroaming struct {
	// Network SSID names enabled for OpenRoaming SSO
	Ssids []string `json:"ssids,omitempty"`
	// Deprecated. Use `openroaming_wba_client_cert` instead.
	WbaCert              *string                `json:"wba_cert,omitempty"` // Deprecated
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SsoOpenroaming,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SsoOpenroaming) String() string {
	return fmt.Sprintf(
		"SsoOpenroaming[Ssids=%v, WbaCert=%v, AdditionalProperties=%v]",
		s.Ssids, s.WbaCert, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SsoOpenroaming.
// It customizes the JSON marshaling process for SsoOpenroaming objects.
func (s SsoOpenroaming) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"ssids", "wba_cert"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SsoOpenroaming object to a map representation for JSON marshaling.
func (s SsoOpenroaming) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.Ssids != nil {
		structMap["ssids"] = s.Ssids
	}
	if s.WbaCert != nil {
		structMap["wba_cert"] = s.WbaCert
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SsoOpenroaming.
// It customizes the JSON unmarshaling process for SsoOpenroaming objects.
func (s *SsoOpenroaming) UnmarshalJSON(input []byte) error {
	var temp tempSsoOpenroaming
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "ssids", "wba_cert")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.Ssids = temp.Ssids
	s.WbaCert = temp.WbaCert
	return nil
}

// tempSsoOpenroaming is a temporary struct used for validating the fields of SsoOpenroaming.
type tempSsoOpenroaming struct {
	Ssids   []string `json:"ssids,omitempty"`
	WbaCert *string  `json:"wba_cert,omitempty"`
}
