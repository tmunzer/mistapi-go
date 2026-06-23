// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"strings"
)

// Ticket represents a Ticket struct.
// Support ticket record with status, comments, and metadata
type Ticket struct {
	// Support case number associated with this ticket
	CaseNumber *string `json:"case_number,omitempty"`
	// List of comments on a support ticket
	Comments []TicketComment `json:"comments,omitempty"`
	// Time when this support ticket was created, in epoch seconds
	CreatedAt *int `json:"created_at,omitempty"`
	// Unique ID of the object instance in the Mist Organization
	Id *uuid.UUID `json:"id,omitempty"`
	// User who opened or requested this support ticket
	Requester *string `json:"requester,omitempty"`
	// Email of the requester
	RequesterEmail *string `json:"requester_email,omitempty"`
	// Ticket status. enum: `closed`, `open`, `pending`, `solved`. `open` means Mist is working on it, `pending` means requester attention is needed, `solved` means Mist considers it resolved but it can still be updated or rated, and `closed` means it is archived
	Status *TicketStatusEnum `json:"status,omitempty"`
	// Short summary of the support request
	Subject string `json:"subject"`
	// Question (default) / bug / critical
	Type string `json:"type"`
	// Time when this support ticket was last updated, in epoch seconds
	UpdatedAt            *int                   `json:"updated_at,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for Ticket,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (t Ticket) String() string {
	return fmt.Sprintf(
		"Ticket[CaseNumber=%v, Comments=%v, CreatedAt=%v, Id=%v, Requester=%v, RequesterEmail=%v, Status=%v, Subject=%v, Type=%v, UpdatedAt=%v, AdditionalProperties=%v]",
		t.CaseNumber, t.Comments, t.CreatedAt, t.Id, t.Requester, t.RequesterEmail, t.Status, t.Subject, t.Type, t.UpdatedAt, t.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for Ticket.
// It customizes the JSON marshaling process for Ticket objects.
func (t Ticket) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(t.AdditionalProperties,
		"case_number", "comments", "created_at", "id", "requester", "requester_email", "status", "subject", "type", "updated_at"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(t.toMap())
}

// toMap converts the Ticket object to a map representation for JSON marshaling.
func (t Ticket) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, t.AdditionalProperties)
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
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "case_number", "comments", "created_at", "id", "requester", "requester_email", "status", "subject", "type", "updated_at")
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
type tempTicket struct {
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
	return errors.New(strings.Join(errs, "\n"))
}
