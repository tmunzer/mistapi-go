package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// TicketComment represents a TicketComment struct.
type TicketComment struct {
    AttachmentIds        []uuid.UUID                `json:"attachment_ids,omitempty"`
    Attachments          []TicketCommentsAttachment `json:"attachments,omitempty"`
    Author               string                     `json:"author"`
    Comment              string                     `json:"comment"`
    CreatedAt            int                        `json:"created_at"`
    AdditionalProperties map[string]any             `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for TicketComment.
// It customizes the JSON marshaling process for TicketComment objects.
func (t TicketComment) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(t.toMap())
}

// toMap converts the TicketComment object to a map representation for JSON marshaling.
func (t TicketComment) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, t.AdditionalProperties)
    if t.AttachmentIds != nil {
        structMap["attachment_ids"] = t.AttachmentIds
    }
    if t.Attachments != nil {
        structMap["attachments"] = t.Attachments
    }
    structMap["author"] = t.Author
    structMap["comment"] = t.Comment
    structMap["created_at"] = t.CreatedAt
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for TicketComment.
// It customizes the JSON unmarshaling process for TicketComment objects.
func (t *TicketComment) UnmarshalJSON(input []byte) error {
    var temp ticketComment
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "attachment_ids", "attachments", "author", "comment", "created_at")
    if err != nil {
    	return err
    }
    
    t.AdditionalProperties = additionalProperties
    t.AttachmentIds = temp.AttachmentIds
    t.Attachments = temp.Attachments
    t.Author = *temp.Author
    t.Comment = *temp.Comment
    t.CreatedAt = *temp.CreatedAt
    return nil
}

// ticketComment is a temporary struct used for validating the fields of TicketComment.
type ticketComment  struct {
    AttachmentIds []uuid.UUID                `json:"attachment_ids,omitempty"`
    Attachments   []TicketCommentsAttachment `json:"attachments,omitempty"`
    Author        *string                    `json:"author"`
    Comment       *string                    `json:"comment"`
    CreatedAt     *int                       `json:"created_at"`
}

func (t *ticketComment) validate() error {
    var errs []string
    if t.Author == nil {
        errs = append(errs, "required field `author` is missing for type `Ticket_Comment`")
    }
    if t.Comment == nil {
        errs = append(errs, "required field `comment` is missing for type `Ticket_Comment`")
    }
    if t.CreatedAt == nil {
        errs = append(errs, "required field `created_at` is missing for type `Ticket_Comment`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
