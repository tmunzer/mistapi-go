package models

import (
    "encoding/json"
    "errors"
    "fmt"
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
    AdditionalProperties map[string]interface{}     `json:"_"`
}

// String implements the fmt.Stringer interface for TicketComment,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (t TicketComment) String() string {
    return fmt.Sprintf(
    	"TicketComment[AttachmentIds=%v, Attachments=%v, Author=%v, Comment=%v, CreatedAt=%v, AdditionalProperties=%v]",
    	t.AttachmentIds, t.Attachments, t.Author, t.Comment, t.CreatedAt, t.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for TicketComment.
// It customizes the JSON marshaling process for TicketComment objects.
func (t TicketComment) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(t.AdditionalProperties,
        "attachment_ids", "attachments", "author", "comment", "created_at"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(t.toMap())
}

// toMap converts the TicketComment object to a map representation for JSON marshaling.
func (t TicketComment) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, t.AdditionalProperties)
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
    var temp tempTicketComment
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "attachment_ids", "attachments", "author", "comment", "created_at")
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

// tempTicketComment is a temporary struct used for validating the fields of TicketComment.
type tempTicketComment  struct {
    AttachmentIds []uuid.UUID                `json:"attachment_ids,omitempty"`
    Attachments   []TicketCommentsAttachment `json:"attachments,omitempty"`
    Author        *string                    `json:"author"`
    Comment       *string                    `json:"comment"`
    CreatedAt     *int                       `json:"created_at"`
}

func (t *tempTicketComment) validate() error {
    var errs []string
    if t.Author == nil {
        errs = append(errs, "required field `author` is missing for type `ticket_comment`")
    }
    if t.Comment == nil {
        errs = append(errs, "required field `comment` is missing for type `ticket_comment`")
    }
    if t.CreatedAt == nil {
        errs = append(errs, "required field `created_at` is missing for type `ticket_comment`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
