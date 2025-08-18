// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// ApClientBridge represents a ApClientBridge struct.
type ApClientBridge struct {
	Auth *ApClientBridgeAuth `json:"auth,omitempty"`
	// When acted as client bridge:
	// * only 5G radio can be used
	// * will not serve as AP on any radios
	Enabled              *bool                  `json:"enabled,omitempty"`
	Ssid                 *string                `json:"ssid,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ApClientBridge,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a ApClientBridge) String() string {
	return fmt.Sprintf(
		"ApClientBridge[Auth=%v, Enabled=%v, Ssid=%v, AdditionalProperties=%v]",
		a.Auth, a.Enabled, a.Ssid, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ApClientBridge.
// It customizes the JSON marshaling process for ApClientBridge objects.
func (a ApClientBridge) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(a.AdditionalProperties,
		"auth", "enabled", "ssid"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(a.toMap())
}

// toMap converts the ApClientBridge object to a map representation for JSON marshaling.
func (a ApClientBridge) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, a.AdditionalProperties)
	if a.Auth != nil {
		structMap["auth"] = a.Auth.toMap()
	}
	if a.Enabled != nil {
		structMap["enabled"] = a.Enabled
	}
	if a.Ssid != nil {
		structMap["ssid"] = a.Ssid
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ApClientBridge.
// It customizes the JSON unmarshaling process for ApClientBridge objects.
func (a *ApClientBridge) UnmarshalJSON(input []byte) error {
	var temp tempApClientBridge
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "auth", "enabled", "ssid")
	if err != nil {
		return err
	}
	a.AdditionalProperties = additionalProperties

	a.Auth = temp.Auth
	a.Enabled = temp.Enabled
	a.Ssid = temp.Ssid
	return nil
}

// tempApClientBridge is a temporary struct used for validating the fields of ApClientBridge.
type tempApClientBridge struct {
	Auth    *ApClientBridgeAuth `json:"auth,omitempty"`
	Enabled *bool               `json:"enabled,omitempty"`
	Ssid    *string             `json:"ssid,omitempty"`
}
