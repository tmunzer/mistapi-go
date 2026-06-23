// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// SsoDeleteAdminsResponse represents a SsoDeleteAdminsResponse struct.
// Result of deleting SSO admin accounts
type SsoDeleteAdminsResponse struct {
	// List of email addresses that were successfully deleted
	Deleted []string `json:"deleted,omitempty"`
	// List of error messages for emails that could not be deleted
	Errors               []string               `json:"errors,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SsoDeleteAdminsResponse,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SsoDeleteAdminsResponse) String() string {
	return fmt.Sprintf(
		"SsoDeleteAdminsResponse[Deleted=%v, Errors=%v, AdditionalProperties=%v]",
		s.Deleted, s.Errors, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SsoDeleteAdminsResponse.
// It customizes the JSON marshaling process for SsoDeleteAdminsResponse objects.
func (s SsoDeleteAdminsResponse) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"deleted", "errors"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SsoDeleteAdminsResponse object to a map representation for JSON marshaling.
func (s SsoDeleteAdminsResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.Deleted != nil {
		structMap["deleted"] = s.Deleted
	}
	if s.Errors != nil {
		structMap["errors"] = s.Errors
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SsoDeleteAdminsResponse.
// It customizes the JSON unmarshaling process for SsoDeleteAdminsResponse objects.
func (s *SsoDeleteAdminsResponse) UnmarshalJSON(input []byte) error {
	var temp tempSsoDeleteAdminsResponse
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "deleted", "errors")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.Deleted = temp.Deleted
	s.Errors = temp.Errors
	return nil
}

// tempSsoDeleteAdminsResponse is a temporary struct used for validating the fields of SsoDeleteAdminsResponse.
type tempSsoDeleteAdminsResponse struct {
	Deleted []string `json:"deleted,omitempty"`
	Errors  []string `json:"errors,omitempty"`
}
