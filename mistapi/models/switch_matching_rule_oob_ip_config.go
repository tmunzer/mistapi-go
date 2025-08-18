// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// SwitchMatchingRuleOobIpConfig represents a SwitchMatchingRuleOobIpConfig struct.
// Out-of-Band Management interface configuration
type SwitchMatchingRuleOobIpConfig struct {
	// enum: `dhcp`, `static`
	Type *IpTypeEnum `json:"type,omitempty"`
	// If supported on the platform. If enabled, DNS will be using this routing-instance, too
	UseMgmtVrf *bool `json:"use_mgmt_vrf,omitempty"`
	// For host-out traffic (NTP/TACPLUS/RADIUS/SYSLOG/SNMP), if alternative source network/ip is desired
	UseMgmtVrfForHostOut *bool                  `json:"use_mgmt_vrf_for_host_out,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SwitchMatchingRuleOobIpConfig,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SwitchMatchingRuleOobIpConfig) String() string {
	return fmt.Sprintf(
		"SwitchMatchingRuleOobIpConfig[Type=%v, UseMgmtVrf=%v, UseMgmtVrfForHostOut=%v, AdditionalProperties=%v]",
		s.Type, s.UseMgmtVrf, s.UseMgmtVrfForHostOut, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SwitchMatchingRuleOobIpConfig.
// It customizes the JSON marshaling process for SwitchMatchingRuleOobIpConfig objects.
func (s SwitchMatchingRuleOobIpConfig) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"type", "use_mgmt_vrf", "use_mgmt_vrf_for_host_out"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SwitchMatchingRuleOobIpConfig object to a map representation for JSON marshaling.
func (s SwitchMatchingRuleOobIpConfig) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.Type != nil {
		structMap["type"] = s.Type
	}
	if s.UseMgmtVrf != nil {
		structMap["use_mgmt_vrf"] = s.UseMgmtVrf
	}
	if s.UseMgmtVrfForHostOut != nil {
		structMap["use_mgmt_vrf_for_host_out"] = s.UseMgmtVrfForHostOut
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SwitchMatchingRuleOobIpConfig.
// It customizes the JSON unmarshaling process for SwitchMatchingRuleOobIpConfig objects.
func (s *SwitchMatchingRuleOobIpConfig) UnmarshalJSON(input []byte) error {
	var temp tempSwitchMatchingRuleOobIpConfig
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "type", "use_mgmt_vrf", "use_mgmt_vrf_for_host_out")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.Type = temp.Type
	s.UseMgmtVrf = temp.UseMgmtVrf
	s.UseMgmtVrfForHostOut = temp.UseMgmtVrfForHostOut
	return nil
}

// tempSwitchMatchingRuleOobIpConfig is a temporary struct used for validating the fields of SwitchMatchingRuleOobIpConfig.
type tempSwitchMatchingRuleOobIpConfig struct {
	Type                 *IpTypeEnum `json:"type,omitempty"`
	UseMgmtVrf           *bool       `json:"use_mgmt_vrf,omitempty"`
	UseMgmtVrfForHostOut *bool       `json:"use_mgmt_vrf_for_host_out,omitempty"`
}
