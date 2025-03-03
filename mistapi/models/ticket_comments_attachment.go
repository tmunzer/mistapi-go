package models

import (
    "encoding/json"
    "fmt"
    "github.com/google/uuid"
)

// TicketCommentsAttachment represents a TicketCommentsAttachment struct.
type TicketCommentsAttachment struct {
    ContentType          *string                `json:"content_type,omitempty"`
    ContentUrl           *string                `json:"content_url,omitempty"`
    CreatedAt            *int                   `json:"created_at,omitempty"`
    FileName             *string                `json:"file_name,omitempty"`
    // Unique ID of the object instance in the Mist Organization
    Id                   *uuid.UUID             `json:"id,omitempty"`
    SizeInBytes          *int                   `json:"size_in_bytes,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for TicketCommentsAttachment,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (t TicketCommentsAttachment) String() string {
    return fmt.Sprintf(
    	"TicketCommentsAttachment[ContentType=%v, ContentUrl=%v, CreatedAt=%v, FileName=%v, Id=%v, SizeInBytes=%v, AdditionalProperties=%v]",
    	t.ContentType, t.ContentUrl, t.CreatedAt, t.FileName, t.Id, t.SizeInBytes, t.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for TicketCommentsAttachment.
// It customizes the JSON marshaling process for TicketCommentsAttachment objects.
func (t TicketCommentsAttachment) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(t.AdditionalProperties,
        "content_type", "content_url", "created_at", "file_name", "id", "size_in_bytes"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(t.toMap())
}

// toMap converts the TicketCommentsAttachment object to a map representation for JSON marshaling.
func (t TicketCommentsAttachment) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, t.AdditionalProperties)
    if t.ContentType != nil {
        structMap["content_type"] = t.ContentType
    }
    if t.ContentUrl != nil {
        structMap["content_url"] = t.ContentUrl
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
    var temp tempTicketCommentsAttachment
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "content_type", "content_url", "created_at", "file_name", "id", "size_in_bytes")
    if err != nil {
    	return err
    }
    t.AdditionalProperties = additionalProperties
    
    t.ContentType = temp.ContentType
    t.ContentUrl = temp.ContentUrl
    t.CreatedAt = temp.CreatedAt
    t.FileName = temp.FileName
    t.Id = temp.Id
    t.SizeInBytes = temp.SizeInBytes
    return nil
}

// tempTicketCommentsAttachment is a temporary struct used for validating the fields of TicketCommentsAttachment.
type tempTicketCommentsAttachment  struct {
    ContentType *string    `json:"content_type,omitempty"`
    ContentUrl  *string    `json:"content_url,omitempty"`
    CreatedAt   *int       `json:"created_at,omitempty"`
    FileName    *string    `json:"file_name,omitempty"`
    Id          *uuid.UUID `json:"id,omitempty"`
    SizeInBytes *int       `json:"size_in_bytes,omitempty"`
}
