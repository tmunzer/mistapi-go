// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// CaptureGatewayGatewaysPort represents a CaptureGatewayGatewaysPort struct.
type CaptureGatewayGatewaysPort struct {
	// tcpdump expression per port
	TcpdumpExpression    *string                `json:"tcpdump_expression,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for CaptureGatewayGatewaysPort,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CaptureGatewayGatewaysPort) String() string {
	return fmt.Sprintf(
		"CaptureGatewayGatewaysPort[TcpdumpExpression=%v, AdditionalProperties=%v]",
		c.TcpdumpExpression, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CaptureGatewayGatewaysPort.
// It customizes the JSON marshaling process for CaptureGatewayGatewaysPort objects.
func (c CaptureGatewayGatewaysPort) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(c.AdditionalProperties,
		"tcpdump_expression"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(c.toMap())
}

// toMap converts the CaptureGatewayGatewaysPort object to a map representation for JSON marshaling.
func (c CaptureGatewayGatewaysPort) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, c.AdditionalProperties)
	if c.TcpdumpExpression != nil {
		structMap["tcpdump_expression"] = c.TcpdumpExpression
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CaptureGatewayGatewaysPort.
// It customizes the JSON unmarshaling process for CaptureGatewayGatewaysPort objects.
func (c *CaptureGatewayGatewaysPort) UnmarshalJSON(input []byte) error {
	var temp tempCaptureGatewayGatewaysPort
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "tcpdump_expression")
	if err != nil {
		return err
	}
	c.AdditionalProperties = additionalProperties

	c.TcpdumpExpression = temp.TcpdumpExpression
	return nil
}

// tempCaptureGatewayGatewaysPort is a temporary struct used for validating the fields of CaptureGatewayGatewaysPort.
type tempCaptureGatewayGatewaysPort struct {
	TcpdumpExpression *string `json:"tcpdump_expression,omitempty"`
}
