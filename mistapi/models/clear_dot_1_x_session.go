// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// ClearDot1xSession represents a ClearDot1xSession struct.
type ClearDot1xSession struct {
	// List of port IDs where the dot1x session must be cleared. Use `all` to clear sessions on all ports.
	Ports                []string               `json:"ports,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ClearDot1xSession,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c ClearDot1xSession) String() string {
	return fmt.Sprintf(
		"ClearDot1xSession[Ports=%v, AdditionalProperties=%v]",
		c.Ports, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ClearDot1xSession.
// It customizes the JSON marshaling process for ClearDot1xSession objects.
func (c ClearDot1xSession) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(c.AdditionalProperties,
		"ports"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(c.toMap())
}

// toMap converts the ClearDot1xSession object to a map representation for JSON marshaling.
func (c ClearDot1xSession) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, c.AdditionalProperties)
	if c.Ports != nil {
		structMap["ports"] = c.Ports
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ClearDot1xSession.
// It customizes the JSON unmarshaling process for ClearDot1xSession objects.
func (c *ClearDot1xSession) UnmarshalJSON(input []byte) error {
	var temp tempClearDot1xSession
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

// tempClearDot1xSession is a temporary struct used for validating the fields of ClearDot1xSession.
type tempClearDot1xSession struct {
	Ports []string `json:"ports,omitempty"`
}
