package models

import (
	"encoding/json"
)

// CaptureGatewayGateways represents a CaptureGatewayGateways struct.
type CaptureGatewayGateways struct {
	// Property key is the port ID
	Ports                map[string]CaptureGatewayGatewaysPort `json:"ports,omitempty"`
	AdditionalProperties map[string]any                        `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CaptureGatewayGateways.
// It customizes the JSON marshaling process for CaptureGatewayGateways objects.
func (c CaptureGatewayGateways) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CaptureGatewayGateways object to a map representation for JSON marshaling.
func (c CaptureGatewayGateways) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, c.AdditionalProperties)
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
	additionalProperties, err := UnmarshalAdditionalProperties(input, "ports")
	if err != nil {
		return err
	}

	c.AdditionalProperties = additionalProperties
	c.Ports = temp.Ports
	return nil
}

// tempCaptureGatewayGateways is a temporary struct used for validating the fields of CaptureGatewayGateways.
type tempCaptureGatewayGateways struct {
	Ports map[string]CaptureGatewayGatewaysPort `json:"ports,omitempty"`
}
