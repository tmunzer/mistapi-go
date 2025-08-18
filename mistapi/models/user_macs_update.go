// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
)

// UserMacsUpdate represents a UserMacsUpdate struct.
type UserMacsUpdate struct {
	Errors               []string               `json:"errors,omitempty"`
	Updated              []uuid.UUID            `json:"updated,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for UserMacsUpdate,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UserMacsUpdate) String() string {
	return fmt.Sprintf(
		"UserMacsUpdate[Errors=%v, Updated=%v, AdditionalProperties=%v]",
		u.Errors, u.Updated, u.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for UserMacsUpdate.
// It customizes the JSON marshaling process for UserMacsUpdate objects.
func (u UserMacsUpdate) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(u.AdditionalProperties,
		"errors", "updated"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(u.toMap())
}

// toMap converts the UserMacsUpdate object to a map representation for JSON marshaling.
func (u UserMacsUpdate) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, u.AdditionalProperties)
	if u.Errors != nil {
		structMap["errors"] = u.Errors
	}
	if u.Updated != nil {
		structMap["updated"] = u.Updated
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UserMacsUpdate.
// It customizes the JSON unmarshaling process for UserMacsUpdate objects.
func (u *UserMacsUpdate) UnmarshalJSON(input []byte) error {
	var temp tempUserMacsUpdate
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "errors", "updated")
	if err != nil {
		return err
	}
	u.AdditionalProperties = additionalProperties

	u.Errors = temp.Errors
	u.Updated = temp.Updated
	return nil
}

// tempUserMacsUpdate is a temporary struct used for validating the fields of UserMacsUpdate.
type tempUserMacsUpdate struct {
	Errors  []string    `json:"errors,omitempty"`
	Updated []uuid.UUID `json:"updated,omitempty"`
}
