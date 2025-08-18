// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// SwitchPortUsageStormControl represents a SwitchPortUsageStormControl struct.
// Switch storm control. Only if `mode`!=`dynamic`
type SwitchPortUsageStormControl struct {
	// Whether to disable storm control on broadcast traffic
	NoBroadcast *bool `json:"no_broadcast,omitempty"`
	// Whether to disable storm control on multicast traffic
	NoMulticast *bool `json:"no_multicast,omitempty"`
	// Whether to disable storm control on registered multicast traffic
	NoRegisteredMulticast *bool `json:"no_registered_multicast,omitempty"`
	// Whether to disable storm control on unknown unicast traffic
	NoUnknownUnicast *bool `json:"no_unknown_unicast,omitempty"`
	// Bandwidth-percentage, configures the storm control level as a percentage of the available bandwidth
	Percentage           *int                   `json:"percentage,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SwitchPortUsageStormControl,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SwitchPortUsageStormControl) String() string {
	return fmt.Sprintf(
		"SwitchPortUsageStormControl[NoBroadcast=%v, NoMulticast=%v, NoRegisteredMulticast=%v, NoUnknownUnicast=%v, Percentage=%v, AdditionalProperties=%v]",
		s.NoBroadcast, s.NoMulticast, s.NoRegisteredMulticast, s.NoUnknownUnicast, s.Percentage, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SwitchPortUsageStormControl.
// It customizes the JSON marshaling process for SwitchPortUsageStormControl objects.
func (s SwitchPortUsageStormControl) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"no_broadcast", "no_multicast", "no_registered_multicast", "no_unknown_unicast", "percentage"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SwitchPortUsageStormControl object to a map representation for JSON marshaling.
func (s SwitchPortUsageStormControl) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.NoBroadcast != nil {
		structMap["no_broadcast"] = s.NoBroadcast
	}
	if s.NoMulticast != nil {
		structMap["no_multicast"] = s.NoMulticast
	}
	if s.NoRegisteredMulticast != nil {
		structMap["no_registered_multicast"] = s.NoRegisteredMulticast
	}
	if s.NoUnknownUnicast != nil {
		structMap["no_unknown_unicast"] = s.NoUnknownUnicast
	}
	if s.Percentage != nil {
		structMap["percentage"] = s.Percentage
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SwitchPortUsageStormControl.
// It customizes the JSON unmarshaling process for SwitchPortUsageStormControl objects.
func (s *SwitchPortUsageStormControl) UnmarshalJSON(input []byte) error {
	var temp tempSwitchPortUsageStormControl
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "no_broadcast", "no_multicast", "no_registered_multicast", "no_unknown_unicast", "percentage")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.NoBroadcast = temp.NoBroadcast
	s.NoMulticast = temp.NoMulticast
	s.NoRegisteredMulticast = temp.NoRegisteredMulticast
	s.NoUnknownUnicast = temp.NoUnknownUnicast
	s.Percentage = temp.Percentage
	return nil
}

// tempSwitchPortUsageStormControl is a temporary struct used for validating the fields of SwitchPortUsageStormControl.
type tempSwitchPortUsageStormControl struct {
	NoBroadcast           *bool `json:"no_broadcast,omitempty"`
	NoMulticast           *bool `json:"no_multicast,omitempty"`
	NoRegisteredMulticast *bool `json:"no_registered_multicast,omitempty"`
	NoUnknownUnicast      *bool `json:"no_unknown_unicast,omitempty"`
	Percentage            *int  `json:"percentage,omitempty"`
}
