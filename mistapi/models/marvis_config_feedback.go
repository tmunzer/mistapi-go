// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// MarvisConfigFeedback represents a MarvisConfigFeedback struct.
// Feedback submission for a Marvis config action
type MarvisConfigFeedback struct {
	// Free-text note about the feedback
	Note *string `json:"note,omitempty"`
	// Feedback type. enum: `invalid`
	Type                 *TypeEnum              `json:"type,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for MarvisConfigFeedback,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m MarvisConfigFeedback) String() string {
	return fmt.Sprintf(
		"MarvisConfigFeedback[Note=%v, Type=%v, AdditionalProperties=%v]",
		m.Note, m.Type, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for MarvisConfigFeedback.
// It customizes the JSON marshaling process for MarvisConfigFeedback objects.
func (m MarvisConfigFeedback) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(m.AdditionalProperties,
		"note", "type"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(m.toMap())
}

// toMap converts the MarvisConfigFeedback object to a map representation for JSON marshaling.
func (m MarvisConfigFeedback) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, m.AdditionalProperties)
	if m.Note != nil {
		structMap["note"] = m.Note
	}
	if m.Type != nil {
		structMap["type"] = m.Type
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MarvisConfigFeedback.
// It customizes the JSON unmarshaling process for MarvisConfigFeedback objects.
func (m *MarvisConfigFeedback) UnmarshalJSON(input []byte) error {
	var temp tempMarvisConfigFeedback
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "note", "type")
	if err != nil {
		return err
	}
	m.AdditionalProperties = additionalProperties

	m.Note = temp.Note
	m.Type = temp.Type
	return nil
}

// tempMarvisConfigFeedback is a temporary struct used for validating the fields of MarvisConfigFeedback.
type tempMarvisConfigFeedback struct {
	Note *string   `json:"note,omitempty"`
	Type *TypeEnum `json:"type,omitempty"`
}
