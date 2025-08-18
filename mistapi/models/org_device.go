// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// OrgDevice represents a OrgDevice struct.
type OrgDevice struct {
	Mac                  string                 `json:"mac"`
	Name                 string                 `json:"name"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for OrgDevice,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (o OrgDevice) String() string {
	return fmt.Sprintf(
		"OrgDevice[Mac=%v, Name=%v, AdditionalProperties=%v]",
		o.Mac, o.Name, o.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for OrgDevice.
// It customizes the JSON marshaling process for OrgDevice objects.
func (o OrgDevice) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(o.AdditionalProperties,
		"mac", "name"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(o.toMap())
}

// toMap converts the OrgDevice object to a map representation for JSON marshaling.
func (o OrgDevice) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, o.AdditionalProperties)
	structMap["mac"] = o.Mac
	structMap["name"] = o.Name
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgDevice.
// It customizes the JSON unmarshaling process for OrgDevice objects.
func (o *OrgDevice) UnmarshalJSON(input []byte) error {
	var temp tempOrgDevice
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "mac", "name")
	if err != nil {
		return err
	}
	o.AdditionalProperties = additionalProperties

	o.Mac = *temp.Mac
	o.Name = *temp.Name
	return nil
}

// tempOrgDevice is a temporary struct used for validating the fields of OrgDevice.
type tempOrgDevice struct {
	Mac  *string `json:"mac"`
	Name *string `json:"name"`
}

func (o *tempOrgDevice) validate() error {
	var errs []string
	if o.Mac == nil {
		errs = append(errs, "required field `mac` is missing for type `org_device`")
	}
	if o.Name == nil {
		errs = append(errs, "required field `name` is missing for type `org_device`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
