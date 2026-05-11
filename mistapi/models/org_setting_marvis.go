// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// OrgSettingMarvis represents a OrgSettingMarvis struct.
type OrgSettingMarvis struct {
	// Self-driving network automation settings per domain
	SelfDriving          *MarvisSelfDriving     `json:"self_driving,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for OrgSettingMarvis,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (o OrgSettingMarvis) String() string {
	return fmt.Sprintf(
		"OrgSettingMarvis[SelfDriving=%v, AdditionalProperties=%v]",
		o.SelfDriving, o.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for OrgSettingMarvis.
// It customizes the JSON marshaling process for OrgSettingMarvis objects.
func (o OrgSettingMarvis) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(o.AdditionalProperties,
		"self_driving"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(o.toMap())
}

// toMap converts the OrgSettingMarvis object to a map representation for JSON marshaling.
func (o OrgSettingMarvis) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, o.AdditionalProperties)
	if o.SelfDriving != nil {
		structMap["self_driving"] = o.SelfDriving.toMap()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgSettingMarvis.
// It customizes the JSON unmarshaling process for OrgSettingMarvis objects.
func (o *OrgSettingMarvis) UnmarshalJSON(input []byte) error {
	var temp tempOrgSettingMarvis
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "self_driving")
	if err != nil {
		return err
	}
	o.AdditionalProperties = additionalProperties

	o.SelfDriving = temp.SelfDriving
	return nil
}

// tempOrgSettingMarvis is a temporary struct used for validating the fields of OrgSettingMarvis.
type tempOrgSettingMarvis struct {
	SelfDriving *MarvisSelfDriving `json:"self_driving,omitempty"`
}
