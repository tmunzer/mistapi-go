// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// SiteSettingJuniperSrx represents a SiteSettingJuniperSrx struct.
type SiteSettingJuniperSrx struct {
	// auto_upgrade device first time it is onboarded
	AutoUpgrade          *JuniperSrxAutoUpgrade         `json:"auto_upgrade,omitempty"`
	Gateways             []SiteSettingJuniperSrxGateway `json:"gateways,omitempty"`
	SendMistNacUserInfo  *bool                          `json:"send_mist_nac_user_info,omitempty"`
	AdditionalProperties map[string]interface{}         `json:"_"`
}

// String implements the fmt.Stringer interface for SiteSettingJuniperSrx,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SiteSettingJuniperSrx) String() string {
	return fmt.Sprintf(
		"SiteSettingJuniperSrx[AutoUpgrade=%v, Gateways=%v, SendMistNacUserInfo=%v, AdditionalProperties=%v]",
		s.AutoUpgrade, s.Gateways, s.SendMistNacUserInfo, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SiteSettingJuniperSrx.
// It customizes the JSON marshaling process for SiteSettingJuniperSrx objects.
func (s SiteSettingJuniperSrx) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"auto_upgrade", "gateways", "send_mist_nac_user_info"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SiteSettingJuniperSrx object to a map representation for JSON marshaling.
func (s SiteSettingJuniperSrx) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.AutoUpgrade != nil {
		structMap["auto_upgrade"] = s.AutoUpgrade.toMap()
	}
	if s.Gateways != nil {
		structMap["gateways"] = s.Gateways
	}
	if s.SendMistNacUserInfo != nil {
		structMap["send_mist_nac_user_info"] = s.SendMistNacUserInfo
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SiteSettingJuniperSrx.
// It customizes the JSON unmarshaling process for SiteSettingJuniperSrx objects.
func (s *SiteSettingJuniperSrx) UnmarshalJSON(input []byte) error {
	var temp tempSiteSettingJuniperSrx
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "auto_upgrade", "gateways", "send_mist_nac_user_info")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.AutoUpgrade = temp.AutoUpgrade
	s.Gateways = temp.Gateways
	s.SendMistNacUserInfo = temp.SendMistNacUserInfo
	return nil
}

// tempSiteSettingJuniperSrx is a temporary struct used for validating the fields of SiteSettingJuniperSrx.
type tempSiteSettingJuniperSrx struct {
	AutoUpgrade         *JuniperSrxAutoUpgrade         `json:"auto_upgrade,omitempty"`
	Gateways            []SiteSettingJuniperSrxGateway `json:"gateways,omitempty"`
	SendMistNacUserInfo *bool                          `json:"send_mist_nac_user_info,omitempty"`
}
