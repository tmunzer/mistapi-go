// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// ClearDot1xSession represents a ClearDot1xSession struct.
type ClearDot1xSession struct {
	// ID of the port where the dot1x session must be cleared. If not provided, the sessions on all the port will be cleared.
	PortId               *string                `json:"port_id,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ClearDot1xSession,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c ClearDot1xSession) String() string {
	return fmt.Sprintf(
		"ClearDot1xSession[PortId=%v, AdditionalProperties=%v]",
		c.PortId, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ClearDot1xSession.
// It customizes the JSON marshaling process for ClearDot1xSession objects.
func (c ClearDot1xSession) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(c.AdditionalProperties,
		"port_id"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(c.toMap())
}

// toMap converts the ClearDot1xSession object to a map representation for JSON marshaling.
func (c ClearDot1xSession) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, c.AdditionalProperties)
	if c.PortId != nil {
		structMap["port_id"] = c.PortId
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
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "port_id")
	if err != nil {
		return err
	}
	c.AdditionalProperties = additionalProperties

	c.PortId = temp.PortId
	return nil
}

// tempClearDot1xSession is a temporary struct used for validating the fields of ClearDot1xSession.
type tempClearDot1xSession struct {
	PortId *string `json:"port_id,omitempty"`
}
