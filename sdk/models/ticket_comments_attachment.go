package models

import (
    "encoding/json"
    "github.com/google/uuid"
)

// TicketCommentsAttachment represents a TicketCommentsAttachment struct.
type TicketCommentsAttachment struct {
    ContentType          *string        `json:"content_type,omitempty"`
    CreatedAt            *int           `json:"created_at,omitempty"`
    FileName             *string        `json:"file_name,omitempty"`
    Id                   *uuid.UUID     `json:"id,omitempty"`
    SizeInBytes          *int           `json:"size_in_bytes,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for TicketCommentsAttachment.
// It customizes the JSON marshaling process for TicketCommentsAttachment objects.
func (t TicketCommentsAttachment) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(t.toMap())
}

// toMap converts the TicketCommentsAttachment object to a map representation for JSON marshaling.
func (t TicketCommentsAttachment) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, t.AdditionalProperties)
    if t.ContentType != nil {
        structMap["content_type"] = t.ContentType
    }
    if t.CreatedAt != nil {
        structMap["created_at"] = t.CreatedAt
    }
    if t.FileName != nil {
        structMap["file_name"] = t.FileName
    }
    if t.Id != nil {
        structMap["id"] = t.Id
    }
    if t.SizeInBytes != nil {
        structMap["size_in_bytes"] = t.SizeInBytes
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for TicketCommentsAttachment.
// It customizes the JSON unmarshaling process for TicketCommentsAttachment objects.
func (t *TicketCommentsAttachment) UnmarshalJSON(input []byte) error {
    var temp ticketCommentsAttachment
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "content_type", "created_at", "file_name", "id", "size_in_bytes")
    if err != nil {
    	return err
    }
    
    t.AdditionalProperties = additionalProperties
    t.ContentType = temp.ContentType
    t.CreatedAt = temp.CreatedAt
    t.FileName = temp.FileName
    t.Id = temp.Id
    t.SizeInBytes = temp.SizeInBytes
    return nil
}

// ticketCommentsAttachment is a temporary struct used for validating the fields of TicketCommentsAttachment.
type ticketCommentsAttachment  struct {
    ContentType *string    `json:"content_type,omitempty"`
    CreatedAt   *int       `json:"created_at,omitempty"`
    FileName    *string    `json:"file_name,omitempty"`
    Id          *uuid.UUID `json:"id,omitempty"`
    SizeInBytes *int       `json:"size_in_bytes,omitempty"`
}
