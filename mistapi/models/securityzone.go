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

// Securityzone represents a Securityzone struct.
// Org-level security zone used by SRX gateways. The zone `name` is used as the security zone name on the device, and Networks reference a zone through their `zone_id`.
type Securityzone struct {
	// When the object has been created, in epoch
	CreatedTime *float64 `json:"created_time,omitempty"`
	// Unique ID of the object instance in the Mist Organization
	Id *uuid.UUID `json:"id,omitempty"`
	// When the object has been modified for the last time, in epoch
	ModifiedTime *float64 `json:"modified_time,omitempty"`
	// Security zone name used on the device. Must start with a letter or a digit, followed by letters, digits, hyphens or underscores
	Name string `json:"name"`
	// Unique identifier of a Mist organization
	OrgId                *uuid.UUID             `json:"org_id,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for Securityzone,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s Securityzone) String() string {
	return fmt.Sprintf(
		"Securityzone[CreatedTime=%v, Id=%v, ModifiedTime=%v, Name=%v, OrgId=%v, AdditionalProperties=%v]",
		s.CreatedTime, s.Id, s.ModifiedTime, s.Name, s.OrgId, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for Securityzone.
// It customizes the JSON marshaling process for Securityzone objects.
func (s Securityzone) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"created_time", "id", "modified_time", "name", "org_id"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the Securityzone object to a map representation for JSON marshaling.
func (s Securityzone) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.CreatedTime != nil {
		structMap["created_time"] = s.CreatedTime
	}
	if s.Id != nil {
		structMap["id"] = s.Id
	}
	if s.ModifiedTime != nil {
		structMap["modified_time"] = s.ModifiedTime
	}
	structMap["name"] = s.Name
	if s.OrgId != nil {
		structMap["org_id"] = s.OrgId
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Securityzone.
// It customizes the JSON unmarshaling process for Securityzone objects.
func (s *Securityzone) UnmarshalJSON(input []byte) error {
	var temp tempSecurityzone
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "created_time", "id", "modified_time", "name", "org_id")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.CreatedTime = temp.CreatedTime
	s.Id = temp.Id
	s.ModifiedTime = temp.ModifiedTime
	s.Name = *temp.Name
	s.OrgId = temp.OrgId
	return nil
}

// tempSecurityzone is a temporary struct used for validating the fields of Securityzone.
type tempSecurityzone struct {
	CreatedTime  *float64   `json:"created_time,omitempty"`
	Id           *uuid.UUID `json:"id,omitempty"`
	ModifiedTime *float64   `json:"modified_time,omitempty"`
	Name         *string    `json:"name"`
	OrgId        *uuid.UUID `json:"org_id,omitempty"`
}

func (s *tempSecurityzone) validate() error {
	var errs []string
	if s.Name == nil {
		errs = append(errs, "required field `name` is missing for type `securityzone`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
