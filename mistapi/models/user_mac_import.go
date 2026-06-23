// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// UserMacImport represents a UserMacImport struct.
// Result of importing user MAC entries
type UserMacImport struct {
	// MAC addresses added during a user MAC import
	Added []string `json:"added,omitempty"`
	// Status message returned for asynchronous imports
	Detail *string `json:"detail,omitempty"`
	// Error messages returned during a user MAC import
	Errors []string `json:"errors,omitempty"`
	// MAC addresses updated during a user MAC import
	Updated              []string               `json:"updated,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for UserMacImport,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UserMacImport) String() string {
	return fmt.Sprintf(
		"UserMacImport[Added=%v, Detail=%v, Errors=%v, Updated=%v, AdditionalProperties=%v]",
		u.Added, u.Detail, u.Errors, u.Updated, u.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for UserMacImport.
// It customizes the JSON marshaling process for UserMacImport objects.
func (u UserMacImport) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(u.AdditionalProperties,
		"added", "detail", "errors", "updated"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(u.toMap())
}

// toMap converts the UserMacImport object to a map representation for JSON marshaling.
func (u UserMacImport) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, u.AdditionalProperties)
	if u.Added != nil {
		structMap["added"] = u.Added
	}
	if u.Detail != nil {
		structMap["detail"] = u.Detail
	}
	if u.Errors != nil {
		structMap["errors"] = u.Errors
	}
	if u.Updated != nil {
		structMap["updated"] = u.Updated
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UserMacImport.
// It customizes the JSON unmarshaling process for UserMacImport objects.
func (u *UserMacImport) UnmarshalJSON(input []byte) error {
	var temp tempUserMacImport
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "added", "detail", "errors", "updated")
	if err != nil {
		return err
	}
	u.AdditionalProperties = additionalProperties

	u.Added = temp.Added
	u.Detail = temp.Detail
	u.Errors = temp.Errors
	u.Updated = temp.Updated
	return nil
}

// tempUserMacImport is a temporary struct used for validating the fields of UserMacImport.
type tempUserMacImport struct {
	Added   []string `json:"added,omitempty"`
	Detail  *string  `json:"detail,omitempty"`
	Errors  []string `json:"errors,omitempty"`
	Updated []string `json:"updated,omitempty"`
}
