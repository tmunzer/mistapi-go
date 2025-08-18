// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// SwitchMatchingRuleIpConfig represents a SwitchMatchingRuleIpConfig struct.
// In-Band Management interface configuration
type SwitchMatchingRuleIpConfig struct {
	// VLAN Name for the management interface
	Network *string `json:"network,omitempty"`
	// enum: `dhcp`, `static`
	Type                 *IpTypeEnum            `json:"type,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SwitchMatchingRuleIpConfig,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SwitchMatchingRuleIpConfig) String() string {
	return fmt.Sprintf(
		"SwitchMatchingRuleIpConfig[Network=%v, Type=%v, AdditionalProperties=%v]",
		s.Network, s.Type, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SwitchMatchingRuleIpConfig.
// It customizes the JSON marshaling process for SwitchMatchingRuleIpConfig objects.
func (s SwitchMatchingRuleIpConfig) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"network", "type"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SwitchMatchingRuleIpConfig object to a map representation for JSON marshaling.
func (s SwitchMatchingRuleIpConfig) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.Network != nil {
		structMap["network"] = s.Network
	}
	if s.Type != nil {
		structMap["type"] = s.Type
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SwitchMatchingRuleIpConfig.
// It customizes the JSON unmarshaling process for SwitchMatchingRuleIpConfig objects.
func (s *SwitchMatchingRuleIpConfig) UnmarshalJSON(input []byte) error {
	var temp tempSwitchMatchingRuleIpConfig
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "network", "type")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.Network = temp.Network
	s.Type = temp.Type
	return nil
}

// tempSwitchMatchingRuleIpConfig is a temporary struct used for validating the fields of SwitchMatchingRuleIpConfig.
type tempSwitchMatchingRuleIpConfig struct {
	Network *string     `json:"network,omitempty"`
	Type    *IpTypeEnum `json:"type,omitempty"`
}
