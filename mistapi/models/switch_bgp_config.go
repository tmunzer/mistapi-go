// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// SwitchBgpConfig represents a SwitchBgpConfig struct.
type SwitchBgpConfig struct {
	// enum: `external`, `internal`
	Type SwitchBgpConfigTypeEnum `json:"type"`
	// List of network names for BGP configuration. When a network is specified, a BGP group will be added to the VRF that network is part of.
	Networks []string `json:"networks,omitempty"`
	// Minimum interval in milliseconds for BFD hello packets. A neighbor is considered failed when the device stops receiving replies after the specified interval. Value must be between 1 and 255000.
	BfdMinimumInterval *int `json:"bfd_minimum_interval,omitempty"`
	// BGP AS, value in range 1-4294967294
	LocalAs BgpAs `json:"local_as"`
	// Hold time is three times the interval at which keepalive messages are sent. It indicates to the peer the length of time that it should consider the sender valid. Must be 0 or a number in the range 3-65535.
	HoldTime *SwitchBgpConfigHoldTime `json:"hold_time,omitempty"`
	AuthKey  *string                  `json:"auth_key,omitempty"`
	// Property key is the BGP Neighbor IP Address.
	Neighbors map[string]SwitchBgpConfigNeighbor `json:"neighbors,omitempty"`
	// Export policy must match one of the policy names defined in the `routing_policies` property.
	ExportPolicy *string `json:"export_policy,omitempty"`
	// Import policy must match one of the policy names defined in the `routing_policies` property.
	ImportPolicy         *string                `json:"import_policy,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SwitchBgpConfig,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SwitchBgpConfig) String() string {
	return fmt.Sprintf(
		"SwitchBgpConfig[Type=%v, Networks=%v, BfdMinimumInterval=%v, LocalAs=%v, HoldTime=%v, AuthKey=%v, Neighbors=%v, ExportPolicy=%v, ImportPolicy=%v, AdditionalProperties=%v]",
		s.Type, s.Networks, s.BfdMinimumInterval, s.LocalAs, s.HoldTime, s.AuthKey, s.Neighbors, s.ExportPolicy, s.ImportPolicy, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SwitchBgpConfig.
// It customizes the JSON marshaling process for SwitchBgpConfig objects.
func (s SwitchBgpConfig) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"type", "networks", "bfd_minimum_interval", "local_as", "hold_time", "auth_key", "neighbors", "export_policy", "import_policy"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SwitchBgpConfig object to a map representation for JSON marshaling.
func (s SwitchBgpConfig) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	structMap["type"] = s.Type
	if s.Networks != nil {
		structMap["networks"] = s.Networks
	}
	if s.BfdMinimumInterval != nil {
		structMap["bfd_minimum_interval"] = s.BfdMinimumInterval
	}
	structMap["local_as"] = s.LocalAs.toMap()
	if s.HoldTime != nil {
		structMap["hold_time"] = s.HoldTime.toMap()
	}
	if s.AuthKey != nil {
		structMap["auth_key"] = s.AuthKey
	}
	if s.Neighbors != nil {
		structMap["neighbors"] = s.Neighbors
	}
	if s.ExportPolicy != nil {
		structMap["export_policy"] = s.ExportPolicy
	}
	if s.ImportPolicy != nil {
		structMap["import_policy"] = s.ImportPolicy
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SwitchBgpConfig.
// It customizes the JSON unmarshaling process for SwitchBgpConfig objects.
func (s *SwitchBgpConfig) UnmarshalJSON(input []byte) error {
	var temp tempSwitchBgpConfig
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "type", "networks", "bfd_minimum_interval", "local_as", "hold_time", "auth_key", "neighbors", "export_policy", "import_policy")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.Type = *temp.Type
	s.Networks = temp.Networks
	s.BfdMinimumInterval = temp.BfdMinimumInterval
	s.LocalAs = *temp.LocalAs
	s.HoldTime = temp.HoldTime
	s.AuthKey = temp.AuthKey
	s.Neighbors = temp.Neighbors
	s.ExportPolicy = temp.ExportPolicy
	s.ImportPolicy = temp.ImportPolicy
	return nil
}

// tempSwitchBgpConfig is a temporary struct used for validating the fields of SwitchBgpConfig.
type tempSwitchBgpConfig struct {
	Type               *SwitchBgpConfigTypeEnum           `json:"type"`
	Networks           []string                           `json:"networks,omitempty"`
	BfdMinimumInterval *int                               `json:"bfd_minimum_interval,omitempty"`
	LocalAs            *BgpAs                             `json:"local_as"`
	HoldTime           *SwitchBgpConfigHoldTime           `json:"hold_time,omitempty"`
	AuthKey            *string                            `json:"auth_key,omitempty"`
	Neighbors          map[string]SwitchBgpConfigNeighbor `json:"neighbors,omitempty"`
	ExportPolicy       *string                            `json:"export_policy,omitempty"`
	ImportPolicy       *string                            `json:"import_policy,omitempty"`
}

func (s *tempSwitchBgpConfig) validate() error {
	var errs []string
	if s.Type == nil {
		errs = append(errs, "required field `type` is missing for type `switch_bgp_config`")
	}
	if s.LocalAs == nil {
		errs = append(errs, "required field `local_as` is missing for type `switch_bgp_config`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
