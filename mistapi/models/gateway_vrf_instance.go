// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// GatewayVrfInstance represents a GatewayVrfInstance struct.
type GatewayVrfInstance struct {
    Networks             []string               `json:"networks,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for GatewayVrfInstance,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (g GatewayVrfInstance) String() string {
    return fmt.Sprintf(
    	"GatewayVrfInstance[Networks=%v, AdditionalProperties=%v]",
    	g.Networks, g.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for GatewayVrfInstance.
// It customizes the JSON marshaling process for GatewayVrfInstance objects.
func (g GatewayVrfInstance) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(g.AdditionalProperties,
        "networks"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(g.toMap())
}

// toMap converts the GatewayVrfInstance object to a map representation for JSON marshaling.
func (g GatewayVrfInstance) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, g.AdditionalProperties)
    if g.Networks != nil {
        structMap["networks"] = g.Networks
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GatewayVrfInstance.
// It customizes the JSON unmarshaling process for GatewayVrfInstance objects.
func (g *GatewayVrfInstance) UnmarshalJSON(input []byte) error {
    var temp tempGatewayVrfInstance
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "networks")
    if err != nil {
    	return err
    }
    g.AdditionalProperties = additionalProperties
    
    g.Networks = temp.Networks
    return nil
}

// tempGatewayVrfInstance is a temporary struct used for validating the fields of GatewayVrfInstance.
type tempGatewayVrfInstance  struct {
    Networks []string `json:"networks,omitempty"`
}
