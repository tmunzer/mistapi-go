// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// OrgSettingMistNacFingerprinting represents a OrgSettingMistNacFingerprinting struct.
// Allows customer to enable client fingerprinting for policy enforcement
type OrgSettingMistNacFingerprinting struct {
	// enable/disable writes to NAC DDB fingerprint table
	Enabled *bool `json:"enabled,omitempty"`
	// enable/disable CoA triggers on fingerprint change for wired clients, always port-bounce
	GenerateCoa *bool `json:"generate_coa,omitempty"`
	// enable/disable CoA triggers on fingerprint change for wireless clients
	GenerateWirelessCoa *bool `json:"generate_wireless_coa,omitempty"`
	// enum: `reauth`, `disconnect`
	WirelessCoaType      *OrgSettingMistNacFingerprintingWirelessCoaEnum `json:"wireless_coa_type,omitempty"`
	AdditionalProperties map[string]interface{}                          `json:"_"`
}

// String implements the fmt.Stringer interface for OrgSettingMistNacFingerprinting,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (o OrgSettingMistNacFingerprinting) String() string {
	return fmt.Sprintf(
		"OrgSettingMistNacFingerprinting[Enabled=%v, GenerateCoa=%v, GenerateWirelessCoa=%v, WirelessCoaType=%v, AdditionalProperties=%v]",
		o.Enabled, o.GenerateCoa, o.GenerateWirelessCoa, o.WirelessCoaType, o.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for OrgSettingMistNacFingerprinting.
// It customizes the JSON marshaling process for OrgSettingMistNacFingerprinting objects.
func (o OrgSettingMistNacFingerprinting) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(o.AdditionalProperties,
		"enabled", "generate_coa", "generate_wireless_coa", "wireless_coa_type"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(o.toMap())
}

// toMap converts the OrgSettingMistNacFingerprinting object to a map representation for JSON marshaling.
func (o OrgSettingMistNacFingerprinting) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, o.AdditionalProperties)
	if o.Enabled != nil {
		structMap["enabled"] = o.Enabled
	}
	if o.GenerateCoa != nil {
		structMap["generate_coa"] = o.GenerateCoa
	}
	if o.GenerateWirelessCoa != nil {
		structMap["generate_wireless_coa"] = o.GenerateWirelessCoa
	}
	if o.WirelessCoaType != nil {
		structMap["wireless_coa_type"] = o.WirelessCoaType
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgSettingMistNacFingerprinting.
// It customizes the JSON unmarshaling process for OrgSettingMistNacFingerprinting objects.
func (o *OrgSettingMistNacFingerprinting) UnmarshalJSON(input []byte) error {
	var temp tempOrgSettingMistNacFingerprinting
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "enabled", "generate_coa", "generate_wireless_coa", "wireless_coa_type")
	if err != nil {
		return err
	}
	o.AdditionalProperties = additionalProperties

	o.Enabled = temp.Enabled
	o.GenerateCoa = temp.GenerateCoa
	o.GenerateWirelessCoa = temp.GenerateWirelessCoa
	o.WirelessCoaType = temp.WirelessCoaType
	return nil
}

// tempOrgSettingMistNacFingerprinting is a temporary struct used for validating the fields of OrgSettingMistNacFingerprinting.
type tempOrgSettingMistNacFingerprinting struct {
	Enabled             *bool                                           `json:"enabled,omitempty"`
	GenerateCoa         *bool                                           `json:"generate_coa,omitempty"`
	GenerateWirelessCoa *bool                                           `json:"generate_wireless_coa,omitempty"`
	WirelessCoaType     *OrgSettingMistNacFingerprintingWirelessCoaEnum `json:"wireless_coa_type,omitempty"`
}
