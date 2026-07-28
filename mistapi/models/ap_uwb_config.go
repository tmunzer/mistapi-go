// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// ApUwbConfig represents a ApUwbConfig struct.
// Ultra-wideband (UWB) RTLS / OMLOX asset-visibility integration settings for an access point. The device-level value overrides the device profile value, which in turn overrides the site-level setting.
type ApUwbConfig struct {
	// Whether UWB RTLS integration is enabled
	Enabled *bool `json:"enabled,omitempty"`
	// RTLS server hostname or IP address
	Host *string `json:"host,omitempty"`
	// RTLS server port number
	Port *int `json:"port,omitempty"`
	// UWB time slot assigned to this AP, 0–15
	Slot *int `json:"slot,omitempty"`
	// UWB integration type. enum: `zigpos`
	Type                 *ApUwbConfigTypeEnum   `json:"type,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ApUwbConfig,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a ApUwbConfig) String() string {
	return fmt.Sprintf(
		"ApUwbConfig[Enabled=%v, Host=%v, Port=%v, Slot=%v, Type=%v, AdditionalProperties=%v]",
		a.Enabled, a.Host, a.Port, a.Slot, a.Type, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ApUwbConfig.
// It customizes the JSON marshaling process for ApUwbConfig objects.
func (a ApUwbConfig) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(a.AdditionalProperties,
		"enabled", "host", "port", "slot", "type"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(a.toMap())
}

// toMap converts the ApUwbConfig object to a map representation for JSON marshaling.
func (a ApUwbConfig) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, a.AdditionalProperties)
	if a.Enabled != nil {
		structMap["enabled"] = a.Enabled
	}
	if a.Host != nil {
		structMap["host"] = a.Host
	}
	if a.Port != nil {
		structMap["port"] = a.Port
	}
	if a.Slot != nil {
		structMap["slot"] = a.Slot
	}
	if a.Type != nil {
		structMap["type"] = a.Type
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ApUwbConfig.
// It customizes the JSON unmarshaling process for ApUwbConfig objects.
func (a *ApUwbConfig) UnmarshalJSON(input []byte) error {
	var temp tempApUwbConfig
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "enabled", "host", "port", "slot", "type")
	if err != nil {
		return err
	}
	a.AdditionalProperties = additionalProperties

	a.Enabled = temp.Enabled
	a.Host = temp.Host
	a.Port = temp.Port
	a.Slot = temp.Slot
	a.Type = temp.Type
	return nil
}

// tempApUwbConfig is a temporary struct used for validating the fields of ApUwbConfig.
type tempApUwbConfig struct {
	Enabled *bool                `json:"enabled,omitempty"`
	Host    *string              `json:"host,omitempty"`
	Port    *int                 `json:"port,omitempty"`
	Slot    *int                 `json:"slot,omitempty"`
	Type    *ApUwbConfigTypeEnum `json:"type,omitempty"`
}
