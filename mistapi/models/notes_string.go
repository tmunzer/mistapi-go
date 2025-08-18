// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// NotesString represents a NotesString struct.
type NotesString struct {
	Notes                *string                `json:"notes,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for NotesString,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (n NotesString) String() string {
	return fmt.Sprintf(
		"NotesString[Notes=%v, AdditionalProperties=%v]",
		n.Notes, n.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for NotesString.
// It customizes the JSON marshaling process for NotesString objects.
func (n NotesString) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(n.AdditionalProperties,
		"notes"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(n.toMap())
}

// toMap converts the NotesString object to a map representation for JSON marshaling.
func (n NotesString) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, n.AdditionalProperties)
	if n.Notes != nil {
		structMap["notes"] = n.Notes
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for NotesString.
// It customizes the JSON unmarshaling process for NotesString objects.
func (n *NotesString) UnmarshalJSON(input []byte) error {
	var temp tempNotesString
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "notes")
	if err != nil {
		return err
	}
	n.AdditionalProperties = additionalProperties

	n.Notes = temp.Notes
	return nil
}

// tempNotesString is a temporary struct used for validating the fields of NotesString.
type tempNotesString struct {
	Notes *string `json:"notes,omitempty"`
}
