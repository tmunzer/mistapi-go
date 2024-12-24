package models

import (
    "encoding/json"
    "fmt"
)

// TicketAttachment represents a TicketAttachment struct.
type TicketAttachment struct {
    ContentUrl           *string                `json:"content_url,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for TicketAttachment,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (t TicketAttachment) String() string {
    return fmt.Sprintf(
    	"TicketAttachment[ContentUrl=%v, AdditionalProperties=%v]",
    	t.ContentUrl, t.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for TicketAttachment.
// It customizes the JSON marshaling process for TicketAttachment objects.
func (t TicketAttachment) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(t.AdditionalProperties,
        "content_url"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(t.toMap())
}

// toMap converts the TicketAttachment object to a map representation for JSON marshaling.
func (t TicketAttachment) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, t.AdditionalProperties)
    if t.ContentUrl != nil {
        structMap["content_url"] = t.ContentUrl
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for TicketAttachment.
// It customizes the JSON unmarshaling process for TicketAttachment objects.
func (t *TicketAttachment) UnmarshalJSON(input []byte) error {
    var temp tempTicketAttachment
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "content_url")
    if err != nil {
    	return err
    }
    t.AdditionalProperties = additionalProperties
    
    t.ContentUrl = temp.ContentUrl
    return nil
}

// tempTicketAttachment is a temporary struct used for validating the fields of TicketAttachment.
type tempTicketAttachment  struct {
    ContentUrl *string `json:"content_url,omitempty"`
}
