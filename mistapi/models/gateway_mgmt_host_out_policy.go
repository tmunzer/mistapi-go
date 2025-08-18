// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// GatewayMgmtHostOutPolicy represents a GatewayMgmtHostOutPolicy struct.
type GatewayMgmtHostOutPolicy struct {
	PathPreference       *string                `json:"path_preference,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for GatewayMgmtHostOutPolicy,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (g GatewayMgmtHostOutPolicy) String() string {
	return fmt.Sprintf(
		"GatewayMgmtHostOutPolicy[PathPreference=%v, AdditionalProperties=%v]",
		g.PathPreference, g.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for GatewayMgmtHostOutPolicy.
// It customizes the JSON marshaling process for GatewayMgmtHostOutPolicy objects.
func (g GatewayMgmtHostOutPolicy) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(g.AdditionalProperties,
		"path_preference"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(g.toMap())
}

// toMap converts the GatewayMgmtHostOutPolicy object to a map representation for JSON marshaling.
func (g GatewayMgmtHostOutPolicy) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, g.AdditionalProperties)
	if g.PathPreference != nil {
		structMap["path_preference"] = g.PathPreference
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GatewayMgmtHostOutPolicy.
// It customizes the JSON unmarshaling process for GatewayMgmtHostOutPolicy objects.
func (g *GatewayMgmtHostOutPolicy) UnmarshalJSON(input []byte) error {
	var temp tempGatewayMgmtHostOutPolicy
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "path_preference")
	if err != nil {
		return err
	}
	g.AdditionalProperties = additionalProperties

	g.PathPreference = temp.PathPreference
	return nil
}

// tempGatewayMgmtHostOutPolicy is a temporary struct used for validating the fields of GatewayMgmtHostOutPolicy.
type tempGatewayMgmtHostOutPolicy struct {
	PathPreference *string `json:"path_preference,omitempty"`
}
