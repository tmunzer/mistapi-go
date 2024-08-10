package models

import (
    "encoding/json"
)

// TicketCommentImportFile represents a TicketCommentImportFile struct.
type TicketCommentImportFile struct {
    Comment              *string        `json:"comment,omitempty"`
    File                 *[]byte        `json:"file,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for TicketCommentImportFile.
// It customizes the JSON marshaling process for TicketCommentImportFile objects.
func (t TicketCommentImportFile) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(t.toMap())
}

// toMap converts the TicketCommentImportFile object to a map representation for JSON marshaling.
func (t TicketCommentImportFile) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, t.AdditionalProperties)
    if t.Comment != nil {
        structMap["comment"] = t.Comment
    }
    if t.File != nil {
        structMap["file"] = t.File
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for TicketCommentImportFile.
// It customizes the JSON unmarshaling process for TicketCommentImportFile objects.
func (t *TicketCommentImportFile) UnmarshalJSON(input []byte) error {
    var temp tempTicketCommentImportFile
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "comment", "file")
    if err != nil {
    	return err
    }
    
    t.AdditionalProperties = additionalProperties
    t.Comment = temp.Comment
    t.File = temp.File
    return nil
}

// tempTicketCommentImportFile is a temporary struct used for validating the fields of TicketCommentImportFile.
type tempTicketCommentImportFile  struct {
    Comment *string `json:"comment,omitempty"`
    File    *[]byte `json:"file,omitempty"`
}
