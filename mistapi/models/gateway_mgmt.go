// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// GatewayMgmt represents a GatewayMgmt struct.
// Gateway settings
type GatewayMgmt struct {
	// Rollback timer for commit confirmed
	ConfigRevertTimer    *int                   `json:"config_revert_timer,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for GatewayMgmt,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (g GatewayMgmt) String() string {
	return fmt.Sprintf(
		"GatewayMgmt[ConfigRevertTimer=%v, AdditionalProperties=%v]",
		g.ConfigRevertTimer, g.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for GatewayMgmt.
// It customizes the JSON marshaling process for GatewayMgmt objects.
func (g GatewayMgmt) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(g.AdditionalProperties,
		"config_revert_timer"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(g.toMap())
}

// toMap converts the GatewayMgmt object to a map representation for JSON marshaling.
func (g GatewayMgmt) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, g.AdditionalProperties)
	if g.ConfigRevertTimer != nil {
		structMap["config_revert_timer"] = g.ConfigRevertTimer
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GatewayMgmt.
// It customizes the JSON unmarshaling process for GatewayMgmt objects.
func (g *GatewayMgmt) UnmarshalJSON(input []byte) error {
	var temp tempGatewayMgmt
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "config_revert_timer")
	if err != nil {
		return err
	}
	g.AdditionalProperties = additionalProperties

	g.ConfigRevertTimer = temp.ConfigRevertTimer
	return nil
}

// tempGatewayMgmt is a temporary struct used for validating the fields of GatewayMgmt.
type tempGatewayMgmt struct {
	ConfigRevertTimer *int `json:"config_revert_timer,omitempty"`
}
