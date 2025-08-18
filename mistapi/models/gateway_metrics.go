// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// GatewayMetrics represents a GatewayMetrics struct.
type GatewayMetrics struct {
	// Config success score
	ConfigSuccess *float64 `json:"config_success,omitempty"`
	// Version compliance score, major version for gateway, type
	VersionCompliance    *GatewayComplianceVersion `json:"version_compliance,omitempty"`
	AdditionalProperties map[string]interface{}    `json:"_"`
}

// String implements the fmt.Stringer interface for GatewayMetrics,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (g GatewayMetrics) String() string {
	return fmt.Sprintf(
		"GatewayMetrics[ConfigSuccess=%v, VersionCompliance=%v, AdditionalProperties=%v]",
		g.ConfigSuccess, g.VersionCompliance, g.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for GatewayMetrics.
// It customizes the JSON marshaling process for GatewayMetrics objects.
func (g GatewayMetrics) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(g.AdditionalProperties,
		"config_success", "version_compliance"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(g.toMap())
}

// toMap converts the GatewayMetrics object to a map representation for JSON marshaling.
func (g GatewayMetrics) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, g.AdditionalProperties)
	if g.ConfigSuccess != nil {
		structMap["config_success"] = g.ConfigSuccess
	}
	if g.VersionCompliance != nil {
		structMap["version_compliance"] = g.VersionCompliance.toMap()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GatewayMetrics.
// It customizes the JSON unmarshaling process for GatewayMetrics objects.
func (g *GatewayMetrics) UnmarshalJSON(input []byte) error {
	var temp tempGatewayMetrics
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "config_success", "version_compliance")
	if err != nil {
		return err
	}
	g.AdditionalProperties = additionalProperties

	g.ConfigSuccess = temp.ConfigSuccess
	g.VersionCompliance = temp.VersionCompliance
	return nil
}

// tempGatewayMetrics is a temporary struct used for validating the fields of GatewayMetrics.
type tempGatewayMetrics struct {
	ConfigSuccess     *float64                  `json:"config_success,omitempty"`
	VersionCompliance *GatewayComplianceVersion `json:"version_compliance,omitempty"`
}
