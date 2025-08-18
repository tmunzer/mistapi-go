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

// Orggroup represents a Orggroup struct.
// Organizations Group
type Orggroup struct {
	// When the object has been created, in epoch
	CreatedTime *float64 `json:"created_time,omitempty"`
	// Unique ID of the object instance in the Mist Organization
	Id *uuid.UUID `json:"id,omitempty"`
	// When the object has been modified for the last time, in epoch
	ModifiedTime         *float64               `json:"modified_time,omitempty"`
	MspId                *uuid.UUID             `json:"msp_id,omitempty"`
	Name                 string                 `json:"name"`
	OrgIds               []uuid.UUID            `json:"org_ids,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for Orggroup,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (o Orggroup) String() string {
	return fmt.Sprintf(
		"Orggroup[CreatedTime=%v, Id=%v, ModifiedTime=%v, MspId=%v, Name=%v, OrgIds=%v, AdditionalProperties=%v]",
		o.CreatedTime, o.Id, o.ModifiedTime, o.MspId, o.Name, o.OrgIds, o.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for Orggroup.
// It customizes the JSON marshaling process for Orggroup objects.
func (o Orggroup) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(o.AdditionalProperties,
		"created_time", "id", "modified_time", "msp_id", "name", "org_ids"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(o.toMap())
}

// toMap converts the Orggroup object to a map representation for JSON marshaling.
func (o Orggroup) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, o.AdditionalProperties)
	if o.CreatedTime != nil {
		structMap["created_time"] = o.CreatedTime
	}
	if o.Id != nil {
		structMap["id"] = o.Id
	}
	if o.ModifiedTime != nil {
		structMap["modified_time"] = o.ModifiedTime
	}
	if o.MspId != nil {
		structMap["msp_id"] = o.MspId
	}
	structMap["name"] = o.Name
	if o.OrgIds != nil {
		structMap["org_ids"] = o.OrgIds
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Orggroup.
// It customizes the JSON unmarshaling process for Orggroup objects.
func (o *Orggroup) UnmarshalJSON(input []byte) error {
	var temp tempOrggroup
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "created_time", "id", "modified_time", "msp_id", "name", "org_ids")
	if err != nil {
		return err
	}
	o.AdditionalProperties = additionalProperties

	o.CreatedTime = temp.CreatedTime
	o.Id = temp.Id
	o.ModifiedTime = temp.ModifiedTime
	o.MspId = temp.MspId
	o.Name = *temp.Name
	o.OrgIds = temp.OrgIds
	return nil
}

// tempOrggroup is a temporary struct used for validating the fields of Orggroup.
type tempOrggroup struct {
	CreatedTime  *float64    `json:"created_time,omitempty"`
	Id           *uuid.UUID  `json:"id,omitempty"`
	ModifiedTime *float64    `json:"modified_time,omitempty"`
	MspId        *uuid.UUID  `json:"msp_id,omitempty"`
	Name         *string     `json:"name"`
	OrgIds       []uuid.UUID `json:"org_ids,omitempty"`
}

func (o *tempOrggroup) validate() error {
	var errs []string
	if o.Name == nil {
		errs = append(errs, "required field `name` is missing for type `orggroup`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
