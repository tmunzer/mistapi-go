package models

import (
    "encoding/json"
)

// NoteString represents a NoteString struct.
type NoteString struct {
    // Some text note describing the intent
    Note                 *string        `json:"note,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for NoteString.
// It customizes the JSON marshaling process for NoteString objects.
func (n NoteString) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(n.toMap())
}

// toMap converts the NoteString object to a map representation for JSON marshaling.
func (n NoteString) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, n.AdditionalProperties)
    if n.Note != nil {
        structMap["note"] = n.Note
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for NoteString.
// It customizes the JSON unmarshaling process for NoteString objects.
func (n *NoteString) UnmarshalJSON(input []byte) error {
    var temp noteString
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "note")
    if err != nil {
    	return err
    }
    
    n.AdditionalProperties = additionalProperties
    n.Note = temp.Note
    return nil
}

// noteString is a temporary struct used for validating the fields of NoteString.
type noteString  struct {
    Note *string `json:"note,omitempty"`
}
