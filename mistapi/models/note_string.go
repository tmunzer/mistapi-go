// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// NoteString represents a NoteString struct.
type NoteString struct {
	// Some text note describing the intent
	Note                 *string                `json:"note,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for NoteString,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (n NoteString) String() string {
	return fmt.Sprintf(
		"NoteString[Note=%v, AdditionalProperties=%v]",
		n.Note, n.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for NoteString.
// It customizes the JSON marshaling process for NoteString objects.
func (n NoteString) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(n.AdditionalProperties,
		"note"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(n.toMap())
}

// toMap converts the NoteString object to a map representation for JSON marshaling.
func (n NoteString) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, n.AdditionalProperties)
	if n.Note != nil {
		structMap["note"] = n.Note
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for NoteString.
// It customizes the JSON unmarshaling process for NoteString objects.
func (n *NoteString) UnmarshalJSON(input []byte) error {
	var temp tempNoteString
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "note")
	if err != nil {
		return err
	}
	n.AdditionalProperties = additionalProperties

	n.Note = temp.Note
	return nil
}

// tempNoteString is a temporary struct used for validating the fields of NoteString.
type tempNoteString struct {
	Note *string `json:"note,omitempty"`
}
