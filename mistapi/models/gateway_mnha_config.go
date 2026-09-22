// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// GatewayMnhaConfig represents a GatewayMnhaConfig struct.
// Multi-Node High Availability (MNHA) configuration, supported on SRX devices only. When enabled, the device operates in MNHA mode instead of chassis-cluster mode.
type GatewayMnhaConfig struct {
	// Whether MNHA mode is enabled
	Enabled              *bool                  `json:"enabled,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for GatewayMnhaConfig,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (g GatewayMnhaConfig) String() string {
	return fmt.Sprintf(
		"GatewayMnhaConfig[Enabled=%v, AdditionalProperties=%v]",
		g.Enabled, g.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for GatewayMnhaConfig.
// It customizes the JSON marshaling process for GatewayMnhaConfig objects.
func (g GatewayMnhaConfig) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(g.AdditionalProperties,
		"enabled"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(g.toMap())
}

// toMap converts the GatewayMnhaConfig object to a map representation for JSON marshaling.
func (g GatewayMnhaConfig) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, g.AdditionalProperties)
	if g.Enabled != nil {
		structMap["enabled"] = g.Enabled
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GatewayMnhaConfig.
// It customizes the JSON unmarshaling process for GatewayMnhaConfig objects.
func (g *GatewayMnhaConfig) UnmarshalJSON(input []byte) error {
	var temp tempGatewayMnhaConfig
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "enabled")
	if err != nil {
		return err
	}
	g.AdditionalProperties = additionalProperties

	g.Enabled = temp.Enabled
	return nil
}

// tempGatewayMnhaConfig is a temporary struct used for validating the fields of GatewayMnhaConfig.
type tempGatewayMnhaConfig struct {
	Enabled *bool `json:"enabled,omitempty"`
}
