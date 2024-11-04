package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// Ticket represents a Ticket struct.
// Support Ticket
type Ticket struct {
    CaseNumber           *string           `json:"case_number,omitempty"`
    Comments             []TicketComment   `json:"comments,omitempty"`
    CreatedAt            *int              `json:"created_at,omitempty"`
    // Unique ID of the object instance in the Mist Organnization
    Id                   *uuid.UUID        `json:"id,omitempty"`
    Requester            *string           `json:"requester,omitempty"`
    // email of the requester
    RequesterEmail       *string           `json:"requester_email,omitempty"`
    // Ticket status. enum:
    // * open: ticket is open, Mist is working on it
    // * pending: ticket is open and Requester attention is needed (e.g. Mist is asking for some more information)
    // * solved: ticket is marked as solved / considered by Mist (requester can update it, causing it to re-open; or rate it)
    // * closed: ticket is archived and cannot be changed.
    Status               *TicketStatusEnum `json:"status,omitempty"`
    Subject              string            `json:"subject"`
    // question (default) / bug / critical
    Type                 string            `json:"type"`
    UpdatedAt            *int              `json:"updated_at,omitempty"`
    AdditionalProperties map[string]any    `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for Ticket.
// It customizes the JSON marshaling process for Ticket objects.
func (t Ticket) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(t.toMap())
}

// toMap converts the Ticket object to a map representation for JSON marshaling.
func (t Ticket) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, t.AdditionalProperties)
    if t.CaseNumber != nil {
        structMap["case_number"] = t.CaseNumber
    }
    if t.Comments != nil {
        structMap["comments"] = t.Comments
    }
    if t.CreatedAt != nil {
        structMap["created_at"] = t.CreatedAt
    }
    if t.Id != nil {
        structMap["id"] = t.Id
    }
    if t.Requester != nil {
        structMap["requester"] = t.Requester
    }
    if t.RequesterEmail != nil {
        structMap["requester_email"] = t.RequesterEmail
    }
    if t.Status != nil {
        structMap["status"] = t.Status
    }
    structMap["subject"] = t.Subject
    structMap["type"] = t.Type
    if t.UpdatedAt != nil {
        structMap["updated_at"] = t.UpdatedAt
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Ticket.
// It customizes the JSON unmarshaling process for Ticket objects.
func (t *Ticket) UnmarshalJSON(input []byte) error {
    var temp tempTicket
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "case_number", "comments", "created_at", "id", "requester", "requester_email", "status", "subject", "type", "updated_at")
    if err != nil {
    	return err
    }
    
    t.AdditionalProperties = additionalProperties
    t.CaseNumber = temp.CaseNumber
    t.Comments = temp.Comments
    t.CreatedAt = temp.CreatedAt
    t.Id = temp.Id
    t.Requester = temp.Requester
    t.RequesterEmail = temp.RequesterEmail
    t.Status = temp.Status
    t.Subject = *temp.Subject
    t.Type = *temp.Type
    t.UpdatedAt = temp.UpdatedAt
    return nil
}

// tempTicket is a temporary struct used for validating the fields of Ticket.
type tempTicket  struct {
    CaseNumber     *string           `json:"case_number,omitempty"`
    Comments       []TicketComment   `json:"comments,omitempty"`
    CreatedAt      *int              `json:"created_at,omitempty"`
    Id             *uuid.UUID        `json:"id,omitempty"`
    Requester      *string           `json:"requester,omitempty"`
    RequesterEmail *string           `json:"requester_email,omitempty"`
    Status         *TicketStatusEnum `json:"status,omitempty"`
    Subject        *string           `json:"subject"`
    Type           *string           `json:"type"`
    UpdatedAt      *int              `json:"updated_at,omitempty"`
}

func (t *tempTicket) validate() error {
    var errs []string
    if t.Subject == nil {
        errs = append(errs, "required field `subject` is missing for type `ticket`")
    }
    if t.Type == nil {
        errs = append(errs, "required field `type` is missing for type `ticket`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
