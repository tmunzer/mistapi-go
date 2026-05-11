// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// SsoDeleteAdmins represents a SsoDeleteAdmins struct.
type SsoDeleteAdmins struct {
	// List of admin email addresses to delete
	Emails               []string               `json:"emails"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SsoDeleteAdmins,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SsoDeleteAdmins) String() string {
	return fmt.Sprintf(
		"SsoDeleteAdmins[Emails=%v, AdditionalProperties=%v]",
		s.Emails, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SsoDeleteAdmins.
// It customizes the JSON marshaling process for SsoDeleteAdmins objects.
func (s SsoDeleteAdmins) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"emails"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SsoDeleteAdmins object to a map representation for JSON marshaling.
func (s SsoDeleteAdmins) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	structMap["emails"] = s.Emails
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SsoDeleteAdmins.
// It customizes the JSON unmarshaling process for SsoDeleteAdmins objects.
func (s *SsoDeleteAdmins) UnmarshalJSON(input []byte) error {
	var temp tempSsoDeleteAdmins
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "emails")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.Emails = *temp.Emails
	return nil
}

// tempSsoDeleteAdmins is a temporary struct used for validating the fields of SsoDeleteAdmins.
type tempSsoDeleteAdmins struct {
	Emails *[]string `json:"emails"`
}

func (s *tempSsoDeleteAdmins) validate() error {
	var errs []string
	if s.Emails == nil {
		errs = append(errs, "required field `emails` is missing for type `sso_delete_admins`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
