// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// GatewayExtraRoute6 represents a GatewayExtraRoute6 struct.
// Gateway IPv6 extra route next-hop settings
type GatewayExtraRoute6 struct {
	// Next-hop IPv6 address for the gateway extra route
	Via                  *string                `json:"via,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for GatewayExtraRoute6,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (g GatewayExtraRoute6) String() string {
	return fmt.Sprintf(
		"GatewayExtraRoute6[Via=%v, AdditionalProperties=%v]",
		g.Via, g.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for GatewayExtraRoute6.
// It customizes the JSON marshaling process for GatewayExtraRoute6 objects.
func (g GatewayExtraRoute6) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(g.AdditionalProperties,
		"via"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(g.toMap())
}

// toMap converts the GatewayExtraRoute6 object to a map representation for JSON marshaling.
func (g GatewayExtraRoute6) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, g.AdditionalProperties)
	if g.Via != nil {
		structMap["via"] = g.Via
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GatewayExtraRoute6.
// It customizes the JSON unmarshaling process for GatewayExtraRoute6 objects.
func (g *GatewayExtraRoute6) UnmarshalJSON(input []byte) error {
	var temp tempGatewayExtraRoute6
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "via")
	if err != nil {
		return err
	}
	g.AdditionalProperties = additionalProperties

	g.Via = temp.Via
	return nil
}

// tempGatewayExtraRoute6 is a temporary struct used for validating the fields of GatewayExtraRoute6.
type tempGatewayExtraRoute6 struct {
	Via *string `json:"via,omitempty"`
}
