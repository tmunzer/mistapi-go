// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// OrgSettingMistNacMdm represents a OrgSettingMistNacMdm struct.
// MDM (Mobile Device Management) CoA configuration
type OrgSettingMistNacMdm struct {
	// CoA type to send. enum: `reauth`, `disconnect`
	CoaType              *NacCoaTypeEnum        `json:"coa_type,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for OrgSettingMistNacMdm,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (o OrgSettingMistNacMdm) String() string {
	return fmt.Sprintf(
		"OrgSettingMistNacMdm[CoaType=%v, AdditionalProperties=%v]",
		o.CoaType, o.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for OrgSettingMistNacMdm.
// It customizes the JSON marshaling process for OrgSettingMistNacMdm objects.
func (o OrgSettingMistNacMdm) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(o.AdditionalProperties,
		"coa_type"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(o.toMap())
}

// toMap converts the OrgSettingMistNacMdm object to a map representation for JSON marshaling.
func (o OrgSettingMistNacMdm) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, o.AdditionalProperties)
	if o.CoaType != nil {
		structMap["coa_type"] = o.CoaType
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgSettingMistNacMdm.
// It customizes the JSON unmarshaling process for OrgSettingMistNacMdm objects.
func (o *OrgSettingMistNacMdm) UnmarshalJSON(input []byte) error {
	var temp tempOrgSettingMistNacMdm
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "coa_type")
	if err != nil {
		return err
	}
	o.AdditionalProperties = additionalProperties

	o.CoaType = temp.CoaType
	return nil
}

// tempOrgSettingMistNacMdm is a temporary struct used for validating the fields of OrgSettingMistNacMdm.
type tempOrgSettingMistNacMdm struct {
	CoaType *NacCoaTypeEnum `json:"coa_type,omitempty"`
}
