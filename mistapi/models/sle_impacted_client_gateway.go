// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// SleImpactedClientGateway represents a SleImpactedClientGateway struct.
type SleImpactedClientGateway struct {
	ChassisMac           *string                `json:"chassis_mac,omitempty"`
	GatewayMac           *string                `json:"gateway_mac,omitempty"`
	GatewayName          *string                `json:"gateway_name,omitempty"`
	Interfaces           []string               `json:"interfaces,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SleImpactedClientGateway,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SleImpactedClientGateway) String() string {
	return fmt.Sprintf(
		"SleImpactedClientGateway[ChassisMac=%v, GatewayMac=%v, GatewayName=%v, Interfaces=%v, AdditionalProperties=%v]",
		s.ChassisMac, s.GatewayMac, s.GatewayName, s.Interfaces, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SleImpactedClientGateway.
// It customizes the JSON marshaling process for SleImpactedClientGateway objects.
func (s SleImpactedClientGateway) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"chassis_mac", "gateway_mac", "gateway_name", "interfaces"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SleImpactedClientGateway object to a map representation for JSON marshaling.
func (s SleImpactedClientGateway) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.ChassisMac != nil {
		structMap["chassis_mac"] = s.ChassisMac
	}
	if s.GatewayMac != nil {
		structMap["gateway_mac"] = s.GatewayMac
	}
	if s.GatewayName != nil {
		structMap["gateway_name"] = s.GatewayName
	}
	if s.Interfaces != nil {
		structMap["interfaces"] = s.Interfaces
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SleImpactedClientGateway.
// It customizes the JSON unmarshaling process for SleImpactedClientGateway objects.
func (s *SleImpactedClientGateway) UnmarshalJSON(input []byte) error {
	var temp tempSleImpactedClientGateway
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "chassis_mac", "gateway_mac", "gateway_name", "interfaces")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.ChassisMac = temp.ChassisMac
	s.GatewayMac = temp.GatewayMac
	s.GatewayName = temp.GatewayName
	s.Interfaces = temp.Interfaces
	return nil
}

// tempSleImpactedClientGateway is a temporary struct used for validating the fields of SleImpactedClientGateway.
type tempSleImpactedClientGateway struct {
	ChassisMac  *string  `json:"chassis_mac,omitempty"`
	GatewayMac  *string  `json:"gateway_mac,omitempty"`
	GatewayName *string  `json:"gateway_name,omitempty"`
	Interfaces  []string `json:"interfaces,omitempty"`
}
