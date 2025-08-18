// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// SiteSettingPaloaltoNetworkGateway represents a SiteSettingPaloaltoNetworkGateway struct.
type SiteSettingPaloaltoNetworkGateway struct {
	ApiKey               *string                `json:"api_key,omitempty"`
	ApiUrl               *string                `json:"api_url,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SiteSettingPaloaltoNetworkGateway,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SiteSettingPaloaltoNetworkGateway) String() string {
	return fmt.Sprintf(
		"SiteSettingPaloaltoNetworkGateway[ApiKey=%v, ApiUrl=%v, AdditionalProperties=%v]",
		s.ApiKey, s.ApiUrl, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SiteSettingPaloaltoNetworkGateway.
// It customizes the JSON marshaling process for SiteSettingPaloaltoNetworkGateway objects.
func (s SiteSettingPaloaltoNetworkGateway) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"api_key", "api_url"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SiteSettingPaloaltoNetworkGateway object to a map representation for JSON marshaling.
func (s SiteSettingPaloaltoNetworkGateway) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.ApiKey != nil {
		structMap["api_key"] = s.ApiKey
	}
	if s.ApiUrl != nil {
		structMap["api_url"] = s.ApiUrl
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SiteSettingPaloaltoNetworkGateway.
// It customizes the JSON unmarshaling process for SiteSettingPaloaltoNetworkGateway objects.
func (s *SiteSettingPaloaltoNetworkGateway) UnmarshalJSON(input []byte) error {
	var temp tempSiteSettingPaloaltoNetworkGateway
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "api_key", "api_url")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.ApiKey = temp.ApiKey
	s.ApiUrl = temp.ApiUrl
	return nil
}

// tempSiteSettingPaloaltoNetworkGateway is a temporary struct used for validating the fields of SiteSettingPaloaltoNetworkGateway.
type tempSiteSettingPaloaltoNetworkGateway struct {
	ApiKey *string `json:"api_key,omitempty"`
	ApiUrl *string `json:"api_url,omitempty"`
}
