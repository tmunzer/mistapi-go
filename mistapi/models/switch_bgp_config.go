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
	AuthKey *string `json:"auth_key,omitempty"`
	// Minimum interval in milliseconds for BFD hello packets. A neighbor is considered failed when the device stops receiving replies after the specified interval. Value must be between 1 and 255000.
	BfdMinimumInterval *int `json:"bfd_minimum_interval,omitempty"`
	// Export policy must match one of the policy names defined in the `routing_policies` property.
	ExportPolicy *string `json:"export_policy,omitempty"`
	// Hold time is three times the interval at which keepalive messages are sent. It indicates to the peer the length of time that it should consider the sender valid. Must be 0 or a number in the range 3-65535.
	HoldTime *int `json:"hold_time,omitempty"`
	// Import policy must match one of the policy names defined in the `routing_policies` property.
	ImportPolicy *string `json:"import_policy,omitempty"`
	// BGP AS, value in range 1-4294967294. Can be a Variable (e.g. `{{bgp_as}}` )
	LocalAs BgpAs `json:"local_as"`
	// Property key is the BGP Neighbor IP Address.
	Neighbors map[string]SwitchBgpConfigNeighbor `json:"neighbors,omitempty"`
	// List of network names for BGP configuration. When a network is specified, a BGP group will be added to the VRF that network is part of.
	Networks []string `json:"networks,omitempty"`
	// enum: `external`, `internal`
	Type                 SwitchBgpConfigTypeEnum `json:"type"`
	AdditionalProperties map[string]interface{}  `json:"_"`
}

// String implements the fmt.Stringer interface for SwitchBgpConfig,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SwitchBgpConfig) String() string {
	return fmt.Sprintf(
		"SwitchBgpConfig[AuthKey=%v, BfdMinimumInterval=%v, ExportPolicy=%v, HoldTime=%v, ImportPolicy=%v, LocalAs=%v, Neighbors=%v, Networks=%v, Type=%v, AdditionalProperties=%v]",
		s.AuthKey, s.BfdMinimumInterval, s.ExportPolicy, s.HoldTime, s.ImportPolicy, s.LocalAs, s.Neighbors, s.Networks, s.Type, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SwitchBgpConfig.
// It customizes the JSON marshaling process for SwitchBgpConfig objects.
func (s SwitchBgpConfig) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"auth_key", "bfd_minimum_interval", "export_policy", "hold_time", "import_policy", "local_as", "neighbors", "networks", "type"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SwitchBgpConfig object to a map representation for JSON marshaling.
func (s SwitchBgpConfig) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.AuthKey != nil {
		structMap["auth_key"] = s.AuthKey
	}
	if s.BfdMinimumInterval != nil {
		structMap["bfd_minimum_interval"] = s.BfdMinimumInterval
	}
	if s.ExportPolicy != nil {
		structMap["export_policy"] = s.ExportPolicy
	}
	if s.HoldTime != nil {
		structMap["hold_time"] = s.HoldTime
	}
	if s.ImportPolicy != nil {
		structMap["import_policy"] = s.ImportPolicy
	}
	structMap["local_as"] = s.LocalAs.toMap()
	if s.Neighbors != nil {
		structMap["neighbors"] = s.Neighbors
	}
	if s.Networks != nil {
		structMap["networks"] = s.Networks
	}
	structMap["type"] = s.Type
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
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "auth_key", "bfd_minimum_interval", "export_policy", "hold_time", "import_policy", "local_as", "neighbors", "networks", "type")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.AuthKey = temp.AuthKey
	s.BfdMinimumInterval = temp.BfdMinimumInterval
	s.ExportPolicy = temp.ExportPolicy
	s.HoldTime = temp.HoldTime
	s.ImportPolicy = temp.ImportPolicy
	s.LocalAs = *temp.LocalAs
	s.Neighbors = temp.Neighbors
	s.Networks = temp.Networks
	s.Type = *temp.Type
	return nil
}

// tempSwitchBgpConfig is a temporary struct used for validating the fields of SwitchBgpConfig.
type tempSwitchBgpConfig struct {
	AuthKey            *string                            `json:"auth_key,omitempty"`
	BfdMinimumInterval *int                               `json:"bfd_minimum_interval,omitempty"`
	ExportPolicy       *string                            `json:"export_policy,omitempty"`
	HoldTime           *int                               `json:"hold_time,omitempty"`
	ImportPolicy       *string                            `json:"import_policy,omitempty"`
	LocalAs            *BgpAs                             `json:"local_as"`
	Neighbors          map[string]SwitchBgpConfigNeighbor `json:"neighbors,omitempty"`
	Networks           []string                           `json:"networks,omitempty"`
	Type               *SwitchBgpConfigTypeEnum           `json:"type"`
}

func (s *tempSwitchBgpConfig) validate() error {
	var errs []string
	if s.LocalAs == nil {
		errs = append(errs, "required field `local_as` is missing for type `switch_bgp_config`")
	}
	if s.Type == nil {
		errs = append(errs, "required field `type` is missing for type `switch_bgp_config`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
