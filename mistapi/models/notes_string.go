package models

import (
	"encoding/json"
)

// NotesString represents a NotesString struct.
type NotesString struct {
	Notes                *string        `json:"notes,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for NotesString.
// It customizes the JSON marshaling process for NotesString objects.
func (n NotesString) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(n.toMap())
}

// toMap converts the NotesString object to a map representation for JSON marshaling.
func (n NotesString) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, n.AdditionalProperties)
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
	additionalProperties, err := UnmarshalAdditionalProperties(input, "notes")
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
