// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// CaptureGatewayGateways represents a CaptureGatewayGateways struct.
type CaptureGatewayGateways struct {
    // Property key is the port ID
    Ports                map[string]CaptureGatewayGatewaysPort `json:"ports,omitempty"`
    AdditionalProperties map[string]interface{}                `json:"_"`
}

// String implements the fmt.Stringer interface for CaptureGatewayGateways,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CaptureGatewayGateways) String() string {
    return fmt.Sprintf(
    	"CaptureGatewayGateways[Ports=%v, AdditionalProperties=%v]",
    	c.Ports, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CaptureGatewayGateways.
// It customizes the JSON marshaling process for CaptureGatewayGateways objects.
func (c CaptureGatewayGateways) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "ports"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CaptureGatewayGateways object to a map representation for JSON marshaling.
func (c CaptureGatewayGateways) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    if c.Ports != nil {
        structMap["ports"] = c.Ports
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CaptureGatewayGateways.
// It customizes the JSON unmarshaling process for CaptureGatewayGateways objects.
func (c *CaptureGatewayGateways) UnmarshalJSON(input []byte) error {
    var temp tempCaptureGatewayGateways
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "ports")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.Ports = temp.Ports
    return nil
}

// tempCaptureGatewayGateways is a temporary struct used for validating the fields of CaptureGatewayGateways.
type tempCaptureGatewayGateways  struct {
    Ports map[string]CaptureGatewayGatewaysPort `json:"ports,omitempty"`
}
