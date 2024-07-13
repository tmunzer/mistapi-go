package models

import (
    "encoding/json"
)

// TicketUploadAttachmentsFile represents a TicketUploadAttachmentsFile struct.
type TicketUploadAttachmentsFile struct {
    // ekahau or ibwave file
    File                 *[]byte        `json:"file,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for TicketUploadAttachmentsFile.
// It customizes the JSON marshaling process for TicketUploadAttachmentsFile objects.
func (t TicketUploadAttachmentsFile) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(t.toMap())
}

// toMap converts the TicketUploadAttachmentsFile object to a map representation for JSON marshaling.
func (t TicketUploadAttachmentsFile) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, t.AdditionalProperties)
    if t.File != nil {
        structMap["file"] = t.File
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for TicketUploadAttachmentsFile.
// It customizes the JSON unmarshaling process for TicketUploadAttachmentsFile objects.
func (t *TicketUploadAttachmentsFile) UnmarshalJSON(input []byte) error {
    var temp ticketUploadAttachmentsFile
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "file")
    if err != nil {
    	return err
    }
    
    t.AdditionalProperties = additionalProperties
    t.File = temp.File
    return nil
}

// ticketUploadAttachmentsFile is a temporary struct used for validating the fields of TicketUploadAttachmentsFile.
type ticketUploadAttachmentsFile  struct {
    File *[]byte `json:"file,omitempty"`
}
