package models

import (
	"encoding/json"
)

// SiteSettingPaloaltoNetworks represents a SiteSettingPaloaltoNetworks struct.
type SiteSettingPaloaltoNetworks struct {
	Gateways             []SiteSettingPaloaltoNetworkGateway `json:"gateways,omitempty"`
	SendMistNacUserInfo  *bool                               `json:"send_mist_nac_user_info,omitempty"`
	AdditionalProperties map[string]any                      `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SiteSettingPaloaltoNetworks.
// It customizes the JSON marshaling process for SiteSettingPaloaltoNetworks objects.
func (s SiteSettingPaloaltoNetworks) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(s.toMap())
}

// toMap converts the SiteSettingPaloaltoNetworks object to a map representation for JSON marshaling.
func (s SiteSettingPaloaltoNetworks) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, s.AdditionalProperties)
	if s.Gateways != nil {
		structMap["gateways"] = s.Gateways
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
	additionalProperties, err := UnmarshalAdditionalProperties(input, "gateways", "send_mist_nac_user_info")
	if err != nil {
		return err
	}

	s.AdditionalProperties = additionalProperties
	s.Gateways = temp.Gateways
	s.SendMistNacUserInfo = temp.SendMistNacUserInfo
	return nil
}

// tempSiteSettingPaloaltoNetworks is a temporary struct used for validating the fields of SiteSettingPaloaltoNetworks.
type tempSiteSettingPaloaltoNetworks struct {
	Gateways            []SiteSettingPaloaltoNetworkGateway `json:"gateways,omitempty"`
	SendMistNacUserInfo *bool                               `json:"send_mist_nac_user_info,omitempty"`
}
