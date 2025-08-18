// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// SwitchPortConfigOverwrite represents a SwitchPortConfigOverwrite struct.
// Switch port config
type SwitchPortConfigOverwrite struct {
	Description *string `json:"description,omitempty"`
	// Whether the port is disabled
	Disabled *bool `json:"disabled,omitempty"`
	// Link connection mode. enum: `auto`, `full`, `half`
	Duplex *SwitchPortUsageDuplexOverwriteEnum `json:"duplex,omitempty"`
	// Max number of mac addresses, default is 0 for unlimited, otherwise range is 1 to 16383 (upper bound constrained by platform)
	MacLimit *SwitchPortUsageMacLimitOverwrite `json:"mac_limit,omitempty"`
	// Whether PoE capabilities are disabled for a port
	PoeDisabled *bool `json:"poe_disabled,omitempty"`
	// Native network/vlan for untagged traffic
	PortNetwork *string `json:"port_network,omitempty"`
	// Port Speed, default is auto to automatically negotiate speed enum: `100m`, `10m`, `1g`, `2.5g`, `5g`, `10g`, `25g`, `40g`, `100g`,`auto`
	Speed                *SwitchPortUsageSpeedOverwriteEnum `json:"speed,omitempty"`
	AdditionalProperties map[string]interface{}             `json:"_"`
}

// String implements the fmt.Stringer interface for SwitchPortConfigOverwrite,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SwitchPortConfigOverwrite) String() string {
	return fmt.Sprintf(
		"SwitchPortConfigOverwrite[Description=%v, Disabled=%v, Duplex=%v, MacLimit=%v, PoeDisabled=%v, PortNetwork=%v, Speed=%v, AdditionalProperties=%v]",
		s.Description, s.Disabled, s.Duplex, s.MacLimit, s.PoeDisabled, s.PortNetwork, s.Speed, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SwitchPortConfigOverwrite.
// It customizes the JSON marshaling process for SwitchPortConfigOverwrite objects.
func (s SwitchPortConfigOverwrite) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"description", "disabled", "duplex", "mac_limit", "poe_disabled", "port_network", "speed"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SwitchPortConfigOverwrite object to a map representation for JSON marshaling.
func (s SwitchPortConfigOverwrite) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.Description != nil {
		structMap["description"] = s.Description
	}
	if s.Disabled != nil {
		structMap["disabled"] = s.Disabled
	}
	if s.Duplex != nil {
		structMap["duplex"] = s.Duplex
	}
	if s.MacLimit != nil {
		structMap["mac_limit"] = s.MacLimit.toMap()
	}
	if s.PoeDisabled != nil {
		structMap["poe_disabled"] = s.PoeDisabled
	}
	if s.PortNetwork != nil {
		structMap["port_network"] = s.PortNetwork
	}
	if s.Speed != nil {
		structMap["speed"] = s.Speed
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SwitchPortConfigOverwrite.
// It customizes the JSON unmarshaling process for SwitchPortConfigOverwrite objects.
func (s *SwitchPortConfigOverwrite) UnmarshalJSON(input []byte) error {
	var temp tempSwitchPortConfigOverwrite
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "description", "disabled", "duplex", "mac_limit", "poe_disabled", "port_network", "speed")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.Description = temp.Description
	s.Disabled = temp.Disabled
	s.Duplex = temp.Duplex
	s.MacLimit = temp.MacLimit
	s.PoeDisabled = temp.PoeDisabled
	s.PortNetwork = temp.PortNetwork
	s.Speed = temp.Speed
	return nil
}

// tempSwitchPortConfigOverwrite is a temporary struct used for validating the fields of SwitchPortConfigOverwrite.
type tempSwitchPortConfigOverwrite struct {
	Description *string                             `json:"description,omitempty"`
	Disabled    *bool                               `json:"disabled,omitempty"`
	Duplex      *SwitchPortUsageDuplexOverwriteEnum `json:"duplex,omitempty"`
	MacLimit    *SwitchPortUsageMacLimitOverwrite   `json:"mac_limit,omitempty"`
	PoeDisabled *bool                               `json:"poe_disabled,omitempty"`
	PortNetwork *string                             `json:"port_network,omitempty"`
	Speed       *SwitchPortUsageSpeedOverwriteEnum  `json:"speed,omitempty"`
}
