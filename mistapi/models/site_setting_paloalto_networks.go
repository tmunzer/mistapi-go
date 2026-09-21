// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// SiteSettingPaloaltoNetworks represents a SiteSettingPaloaltoNetworks struct.
// Palo Alto Networks integration settings for the site
type SiteSettingPaloaltoNetworks struct {
	// Palo Alto Networks gateways integrated with a site
	Gateways []SiteSettingPaloaltoNetworkGateway `json:"gateways,omitempty"`
	// Source of the Mist NAC user role sent to firewall gateways. enum: `idp_role`, `radius_group`, `none`
	MistNacUserRoleSource *SiteSettingMistNacUserRoleSourceEnum `json:"mist_nac_user_role_source,omitempty"`
	// Whether Mist NAC user information is sent to Palo Alto Networks gateways
	SendMistNacUserInfo  *bool                  `json:"send_mist_nac_user_info,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SiteSettingPaloaltoNetworks,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SiteSettingPaloaltoNetworks) String() string {
	return fmt.Sprintf(
		"SiteSettingPaloaltoNetworks[Gateways=%v, MistNacUserRoleSource=%v, SendMistNacUserInfo=%v, AdditionalProperties=%v]",
		s.Gateways, s.MistNacUserRoleSource, s.SendMistNacUserInfo, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SiteSettingPaloaltoNetworks.
// It customizes the JSON marshaling process for SiteSettingPaloaltoNetworks objects.
func (s SiteSettingPaloaltoNetworks) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"gateways", "mist_nac_user_role_source", "send_mist_nac_user_info"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SiteSettingPaloaltoNetworks object to a map representation for JSON marshaling.
func (s SiteSettingPaloaltoNetworks) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.Gateways != nil {
		structMap["gateways"] = s.Gateways
	}
	if s.MistNacUserRoleSource != nil {
		structMap["mist_nac_user_role_source"] = s.MistNacUserRoleSource
	}
	if s.SendMistNacUserInfo != nil {
		structMap["send_mist_nac_user_info"] = s.SendMistNacUserInfo
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SiteSettingPaloaltoNetworks.
// It customizes the JSON unmarshaling process for SiteSettingPaloaltoNetworks objects.
func (s *SiteSettingPaloaltoNetworks) UnmarshalJSON(input []byte) error {
	var temp tempSiteSettingPaloaltoNetworks
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "gateways", "mist_nac_user_role_source", "send_mist_nac_user_info")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.Gateways = temp.Gateways
	s.MistNacUserRoleSource = temp.MistNacUserRoleSource
	s.SendMistNacUserInfo = temp.SendMistNacUserInfo
	return nil
}

// tempSiteSettingPaloaltoNetworks is a temporary struct used for validating the fields of SiteSettingPaloaltoNetworks.
type tempSiteSettingPaloaltoNetworks struct {
	Gateways              []SiteSettingPaloaltoNetworkGateway   `json:"gateways,omitempty"`
	MistNacUserRoleSource *SiteSettingMistNacUserRoleSourceEnum `json:"mist_nac_user_role_source,omitempty"`
	SendMistNacUserInfo   *bool                                 `json:"send_mist_nac_user_info,omitempty"`
}
