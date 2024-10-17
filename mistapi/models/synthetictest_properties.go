package models

import (
	"encoding/json"
)

// SynthetictestProperties represents a SynthetictestProperties struct.
type SynthetictestProperties struct {
	CustomTestUrls []string `json:"custom_test_urls,omitempty"`
	// for some vlans where we don't want this to run
	Disabled             *bool                `json:"disabled,omitempty"`
	VlanIds              []VlanIdWithVariable `json:"vlan_ids,omitempty"`
	AdditionalProperties map[string]any       `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SynthetictestProperties.
// It customizes the JSON marshaling process for SynthetictestProperties objects.
func (s SynthetictestProperties) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(s.toMap())
}

// toMap converts the SynthetictestProperties object to a map representation for JSON marshaling.
func (s SynthetictestProperties) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, s.AdditionalProperties)
	if s.CustomTestUrls != nil {
		structMap["custom_test_urls"] = s.CustomTestUrls
	}
	if s.Disabled != nil {
		structMap["disabled"] = s.Disabled
	}
	if s.VlanIds != nil {
		structMap["vlan_ids"] = s.VlanIds
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SynthetictestProperties.
// It customizes the JSON unmarshaling process for SynthetictestProperties objects.
func (s *SynthetictestProperties) UnmarshalJSON(input []byte) error {
	var temp tempSynthetictestProperties
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "custom_test_urls", "disabled", "vlan_ids")
	if err != nil {
		return err
	}

	s.AdditionalProperties = additionalProperties
	s.CustomTestUrls = temp.CustomTestUrls
	s.Disabled = temp.Disabled
	s.VlanIds = temp.VlanIds
	return nil
}

// tempSynthetictestProperties is a temporary struct used for validating the fields of SynthetictestProperties.
type tempSynthetictestProperties struct {
	CustomTestUrls []string             `json:"custom_test_urls,omitempty"`
	Disabled       *bool                `json:"disabled,omitempty"`
	VlanIds        []VlanIdWithVariable `json:"vlan_ids,omitempty"`
}
