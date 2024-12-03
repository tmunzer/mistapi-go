package models

import (
    "encoding/json"
)

// TicketCommentImportFile represents a TicketCommentImportFile struct.
type TicketCommentImportFile struct {
    Comment              *string                `json:"comment,omitempty"`
    File                 *[]byte                `json:"file,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for TicketCommentImportFile.
// It customizes the JSON marshaling process for TicketCommentImportFile objects.
func (t TicketCommentImportFile) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(t.AdditionalProperties,
        "comment", "file"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(t.toMap())
}

// toMap converts the TicketCommentImportFile object to a map representation for JSON marshaling.
func (t TicketCommentImportFile) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, t.AdditionalProperties)
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
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "comment", "file")
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
