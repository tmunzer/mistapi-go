// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// OrgSettingAutoDeviceNamingRule represents a OrgSettingAutoDeviceNamingRule struct.
type OrgSettingAutoDeviceNamingRule struct {
	// "[0:3]"            // "abcdef" -> "abc"
	// "split(.)[1]"      // "a.b.c" -> "b"
	// "split(-)[1][0:3]" // "a1234-b5678-c90" -> "b56"'
	Expression *string `json:"expression,omitempty"`
	// enum: `ap`, `gateway`, `switch`
	MatchDevice *DeviceTypeDefaultApEnum `json:"match_device,omitempty"`
	// Prefix to append to the device name
	Prefix *string `json:"prefix,omitempty"`
	// enum: `lldp_port_desc`, `mac`
	Src *OrgSettingAutoDeviceNamingRuleSrcEnum `json:"src,omitempty"`
	// Suffix to append to the device name
	Suffix               *string                `json:"suffix,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for OrgSettingAutoDeviceNamingRule,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (o OrgSettingAutoDeviceNamingRule) String() string {
	return fmt.Sprintf(
		"OrgSettingAutoDeviceNamingRule[Expression=%v, MatchDevice=%v, Prefix=%v, Src=%v, Suffix=%v, AdditionalProperties=%v]",
		o.Expression, o.MatchDevice, o.Prefix, o.Src, o.Suffix, o.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for OrgSettingAutoDeviceNamingRule.
// It customizes the JSON marshaling process for OrgSettingAutoDeviceNamingRule objects.
func (o OrgSettingAutoDeviceNamingRule) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(o.AdditionalProperties,
		"expression", "match_device", "prefix", "src", "suffix"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(o.toMap())
}

// toMap converts the OrgSettingAutoDeviceNamingRule object to a map representation for JSON marshaling.
func (o OrgSettingAutoDeviceNamingRule) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, o.AdditionalProperties)
	if o.Expression != nil {
		structMap["expression"] = o.Expression
	}
	if o.MatchDevice != nil {
		structMap["match_device"] = o.MatchDevice
	}
	if o.Prefix != nil {
		structMap["prefix"] = o.Prefix
	}
	if o.Src != nil {
		structMap["src"] = o.Src
	}
	if o.Suffix != nil {
		structMap["suffix"] = o.Suffix
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgSettingAutoDeviceNamingRule.
// It customizes the JSON unmarshaling process for OrgSettingAutoDeviceNamingRule objects.
func (o *OrgSettingAutoDeviceNamingRule) UnmarshalJSON(input []byte) error {
	var temp tempOrgSettingAutoDeviceNamingRule
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "expression", "match_device", "prefix", "src", "suffix")
	if err != nil {
		return err
	}
	o.AdditionalProperties = additionalProperties

	o.Expression = temp.Expression
	o.MatchDevice = temp.MatchDevice
	o.Prefix = temp.Prefix
	o.Src = temp.Src
	o.Suffix = temp.Suffix
	return nil
}

// tempOrgSettingAutoDeviceNamingRule is a temporary struct used for validating the fields of OrgSettingAutoDeviceNamingRule.
type tempOrgSettingAutoDeviceNamingRule struct {
	Expression  *string                                `json:"expression,omitempty"`
	MatchDevice *DeviceTypeDefaultApEnum               `json:"match_device,omitempty"`
	Prefix      *string                                `json:"prefix,omitempty"`
	Src         *OrgSettingAutoDeviceNamingRuleSrcEnum `json:"src,omitempty"`
	Suffix      *string                                `json:"suffix,omitempty"`
}
