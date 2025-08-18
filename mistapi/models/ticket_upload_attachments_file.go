// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// TicketUploadAttachmentsFile represents a TicketUploadAttachmentsFile struct.
type TicketUploadAttachmentsFile struct {
	// Ekahau or ibwave file
	File                 *[]byte                `json:"file,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for TicketUploadAttachmentsFile,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (t TicketUploadAttachmentsFile) String() string {
	return fmt.Sprintf(
		"TicketUploadAttachmentsFile[File=%v, AdditionalProperties=%v]",
		t.File, t.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for TicketUploadAttachmentsFile.
// It customizes the JSON marshaling process for TicketUploadAttachmentsFile objects.
func (t TicketUploadAttachmentsFile) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(t.AdditionalProperties,
		"file"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(t.toMap())
}

// toMap converts the TicketUploadAttachmentsFile object to a map representation for JSON marshaling.
func (t TicketUploadAttachmentsFile) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, t.AdditionalProperties)
	if t.File != nil {
		structMap["file"] = t.File
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for TicketUploadAttachmentsFile.
// It customizes the JSON unmarshaling process for TicketUploadAttachmentsFile objects.
func (t *TicketUploadAttachmentsFile) UnmarshalJSON(input []byte) error {
	var temp tempTicketUploadAttachmentsFile
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "file")
	if err != nil {
		return err
	}
	t.AdditionalProperties = additionalProperties

	t.File = temp.File
	return nil
}

// tempTicketUploadAttachmentsFile is a temporary struct used for validating the fields of TicketUploadAttachmentsFile.
type tempTicketUploadAttachmentsFile struct {
	File *[]byte `json:"file,omitempty"`
}
