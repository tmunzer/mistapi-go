// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// ConstMarvisclientEvent represents a ConstMarvisclientEvent struct.
// A Marvis Client event type definition
type ConstMarvisclientEvent struct {
	// Human-readable name for this Marvis Client event type
	Display *string `json:"display,omitempty"`
	// Event type key used in Marvis Client event search and count APIs
	Key                  *string                `json:"key,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ConstMarvisclientEvent,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c ConstMarvisclientEvent) String() string {
	return fmt.Sprintf(
		"ConstMarvisclientEvent[Display=%v, Key=%v, AdditionalProperties=%v]",
		c.Display, c.Key, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ConstMarvisclientEvent.
// It customizes the JSON marshaling process for ConstMarvisclientEvent objects.
func (c ConstMarvisclientEvent) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(c.AdditionalProperties,
		"display", "key"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(c.toMap())
}

// toMap converts the ConstMarvisclientEvent object to a map representation for JSON marshaling.
func (c ConstMarvisclientEvent) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, c.AdditionalProperties)
	if c.Display != nil {
		structMap["display"] = c.Display
	}
	if c.Key != nil {
		structMap["key"] = c.Key
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ConstMarvisclientEvent.
// It customizes the JSON unmarshaling process for ConstMarvisclientEvent objects.
func (c *ConstMarvisclientEvent) UnmarshalJSON(input []byte) error {
	var temp tempConstMarvisclientEvent
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "display", "key")
	if err != nil {
		return err
	}
	c.AdditionalProperties = additionalProperties

	c.Display = temp.Display
	c.Key = temp.Key
	return nil
}

// tempConstMarvisclientEvent is a temporary struct used for validating the fields of ConstMarvisclientEvent.
type tempConstMarvisclientEvent struct {
	Display *string `json:"display,omitempty"`
	Key     *string `json:"key,omitempty"`
}
