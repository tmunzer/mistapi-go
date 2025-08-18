// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// CaptureMxedgeMxedgesInterfaces represents a CaptureMxedgeMxedgesInterfaces struct.
// Property key is the Port name (e.g. "port1", "kni0", "lacp0", "ipsec", "drop", "oobm"), currently limited to specifying one interface per mxedge
type CaptureMxedgeMxedgesInterfaces struct {
	// tcpdump expression common for wired,radiotap
	TcpdumpExpression    *string                `json:"tcpdump_expression,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for CaptureMxedgeMxedgesInterfaces,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CaptureMxedgeMxedgesInterfaces) String() string {
	return fmt.Sprintf(
		"CaptureMxedgeMxedgesInterfaces[TcpdumpExpression=%v, AdditionalProperties=%v]",
		c.TcpdumpExpression, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CaptureMxedgeMxedgesInterfaces.
// It customizes the JSON marshaling process for CaptureMxedgeMxedgesInterfaces objects.
func (c CaptureMxedgeMxedgesInterfaces) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(c.AdditionalProperties,
		"tcpdump_expression"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(c.toMap())
}

// toMap converts the CaptureMxedgeMxedgesInterfaces object to a map representation for JSON marshaling.
func (c CaptureMxedgeMxedgesInterfaces) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, c.AdditionalProperties)
	if c.TcpdumpExpression != nil {
		structMap["tcpdump_expression"] = c.TcpdumpExpression
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CaptureMxedgeMxedgesInterfaces.
// It customizes the JSON unmarshaling process for CaptureMxedgeMxedgesInterfaces objects.
func (c *CaptureMxedgeMxedgesInterfaces) UnmarshalJSON(input []byte) error {
	var temp tempCaptureMxedgeMxedgesInterfaces
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

// tempCaptureMxedgeMxedgesInterfaces is a temporary struct used for validating the fields of CaptureMxedgeMxedgesInterfaces.
type tempCaptureMxedgeMxedgesInterfaces struct {
	TcpdumpExpression *string `json:"tcpdump_expression,omitempty"`
}
