// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// OrgSettingAutoDeviceprofileAssignment represents a OrgSettingAutoDeviceprofileAssignment struct.
type OrgSettingAutoDeviceprofileAssignment struct {
	Enable               *bool                                    `json:"enable,omitempty"`
	Rules                Optional[[]OrgSettingAutoAssignmentRule] `json:"rules"`
	AdditionalProperties map[string]interface{}                   `json:"_"`
}

// String implements the fmt.Stringer interface for OrgSettingAutoDeviceprofileAssignment,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (o OrgSettingAutoDeviceprofileAssignment) String() string {
	return fmt.Sprintf(
		"OrgSettingAutoDeviceprofileAssignment[Enable=%v, Rules=%v, AdditionalProperties=%v]",
		o.Enable, o.Rules, o.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for OrgSettingAutoDeviceprofileAssignment.
// It customizes the JSON marshaling process for OrgSettingAutoDeviceprofileAssignment objects.
func (o OrgSettingAutoDeviceprofileAssignment) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(o.AdditionalProperties,
		"enable", "rules"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(o.toMap())
}

// toMap converts the OrgSettingAutoDeviceprofileAssignment object to a map representation for JSON marshaling.
func (o OrgSettingAutoDeviceprofileAssignment) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, o.AdditionalProperties)
	if o.Enable != nil {
		structMap["enable"] = o.Enable
	}
	if o.Rules.IsValueSet() {
		if o.Rules.Value() != nil {
			structMap["rules"] = o.Rules.Value()
		} else {
			structMap["rules"] = nil
		}
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgSettingAutoDeviceprofileAssignment.
// It customizes the JSON unmarshaling process for OrgSettingAutoDeviceprofileAssignment objects.
func (o *OrgSettingAutoDeviceprofileAssignment) UnmarshalJSON(input []byte) error {
	var temp tempOrgSettingAutoDeviceprofileAssignment
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "enable", "rules")
	if err != nil {
		return err
	}
	o.AdditionalProperties = additionalProperties

	o.Enable = temp.Enable
	o.Rules = temp.Rules
	return nil
}

// tempOrgSettingAutoDeviceprofileAssignment is a temporary struct used for validating the fields of OrgSettingAutoDeviceprofileAssignment.
type tempOrgSettingAutoDeviceprofileAssignment struct {
	Enable *bool                                    `json:"enable,omitempty"`
	Rules  Optional[[]OrgSettingAutoAssignmentRule] `json:"rules"`
}
