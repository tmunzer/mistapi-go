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

// UserMac represents a UserMac struct.
type UserMac struct {
	// Unique ID of the object instance in the Mist Organization
	Id     *uuid.UUID `json:"id,omitempty"`
	Labels []string   `json:"labels,omitempty"`
	// Only non-local-admin MAC is accepted
	Mac                  string                 `json:"mac"`
	Name                 *string                `json:"name,omitempty"`
	Notes                *string                `json:"notes,omitempty"`
	RadiusGroup          *string                `json:"radius_group,omitempty"`
	Vlan                 *string                `json:"vlan,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for UserMac,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UserMac) String() string {
	return fmt.Sprintf(
		"UserMac[Id=%v, Labels=%v, Mac=%v, Name=%v, Notes=%v, RadiusGroup=%v, Vlan=%v, AdditionalProperties=%v]",
		u.Id, u.Labels, u.Mac, u.Name, u.Notes, u.RadiusGroup, u.Vlan, u.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for UserMac.
// It customizes the JSON marshaling process for UserMac objects.
func (u UserMac) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(u.AdditionalProperties,
		"id", "labels", "mac", "name", "notes", "radius_group", "vlan"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(u.toMap())
}

// toMap converts the UserMac object to a map representation for JSON marshaling.
func (u UserMac) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, u.AdditionalProperties)
	if u.Id != nil {
		structMap["id"] = u.Id
	}
	if u.Labels != nil {
		structMap["labels"] = u.Labels
	}
	structMap["mac"] = u.Mac
	if u.Name != nil {
		structMap["name"] = u.Name
	}
	if u.Notes != nil {
		structMap["notes"] = u.Notes
	}
	if u.RadiusGroup != nil {
		structMap["radius_group"] = u.RadiusGroup
	}
	if u.Vlan != nil {
		structMap["vlan"] = u.Vlan
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UserMac.
// It customizes the JSON unmarshaling process for UserMac objects.
func (u *UserMac) UnmarshalJSON(input []byte) error {
	var temp tempUserMac
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "id", "labels", "mac", "name", "notes", "radius_group", "vlan")
	if err != nil {
		return err
	}
	u.AdditionalProperties = additionalProperties

	u.Id = temp.Id
	u.Labels = temp.Labels
	u.Mac = *temp.Mac
	u.Name = temp.Name
	u.Notes = temp.Notes
	u.RadiusGroup = temp.RadiusGroup
	u.Vlan = temp.Vlan
	return nil
}

// tempUserMac is a temporary struct used for validating the fields of UserMac.
type tempUserMac struct {
	Id          *uuid.UUID `json:"id,omitempty"`
	Labels      []string   `json:"labels,omitempty"`
	Mac         *string    `json:"mac"`
	Name        *string    `json:"name,omitempty"`
	Notes       *string    `json:"notes,omitempty"`
	RadiusGroup *string    `json:"radius_group,omitempty"`
	Vlan        *string    `json:"vlan,omitempty"`
}

func (u *tempUserMac) validate() error {
	var errs []string
	if u.Mac == nil {
		errs = append(errs, "required field `mac` is missing for type `user_mac`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
