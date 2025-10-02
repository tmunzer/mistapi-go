// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// OrgSettingJuniperSrx represents a OrgSettingJuniperSrx struct.
type OrgSettingJuniperSrx struct {
	// auto_upgrade device first time it is onboarded
	AutoUpgrade          *JuniperSrxAutoUpgrade `json:"auto_upgrade,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for OrgSettingJuniperSrx,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (o OrgSettingJuniperSrx) String() string {
	return fmt.Sprintf(
		"OrgSettingJuniperSrx[AutoUpgrade=%v, AdditionalProperties=%v]",
		o.AutoUpgrade, o.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for OrgSettingJuniperSrx.
// It customizes the JSON marshaling process for OrgSettingJuniperSrx objects.
func (o OrgSettingJuniperSrx) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(o.AdditionalProperties,
		"auto_upgrade"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(o.toMap())
}

// toMap converts the OrgSettingJuniperSrx object to a map representation for JSON marshaling.
func (o OrgSettingJuniperSrx) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, o.AdditionalProperties)
	if o.AutoUpgrade != nil {
		structMap["auto_upgrade"] = o.AutoUpgrade.toMap()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgSettingJuniperSrx.
// It customizes the JSON unmarshaling process for OrgSettingJuniperSrx objects.
func (o *OrgSettingJuniperSrx) UnmarshalJSON(input []byte) error {
	var temp tempOrgSettingJuniperSrx
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "auto_upgrade")
	if err != nil {
		return err
	}
	o.AdditionalProperties = additionalProperties

	o.AutoUpgrade = temp.AutoUpgrade
	return nil
}

// tempOrgSettingJuniperSrx is a temporary struct used for validating the fields of OrgSettingJuniperSrx.
type tempOrgSettingJuniperSrx struct {
	AutoUpgrade *JuniperSrxAutoUpgrade `json:"auto_upgrade,omitempty"`
}
