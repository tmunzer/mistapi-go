// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// SleImpactedUsersClient represents a SleImpactedUsersClient struct.
type SleImpactedUsersClient struct {
	Degraded             *float64                   `json:"degraded,omitempty"`
	Duration             *float64                   `json:"duration,omitempty"`
	Gateways             []SleImpactedClientGateway `json:"gateways,omitempty"`
	Mac                  *string                    `json:"mac,omitempty"`
	Name                 *string                    `json:"name,omitempty"`
	SrcIp                *string                    `json:"src_ip,omitempty"`
	Total                *float64                   `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}     `json:"_"`
}

// String implements the fmt.Stringer interface for SleImpactedUsersClient,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SleImpactedUsersClient) String() string {
	return fmt.Sprintf(
		"SleImpactedUsersClient[Degraded=%v, Duration=%v, Gateways=%v, Mac=%v, Name=%v, SrcIp=%v, Total=%v, AdditionalProperties=%v]",
		s.Degraded, s.Duration, s.Gateways, s.Mac, s.Name, s.SrcIp, s.Total, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SleImpactedUsersClient.
// It customizes the JSON marshaling process for SleImpactedUsersClient objects.
func (s SleImpactedUsersClient) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"degraded", "duration", "gateways", "mac", "name", "src_ip", "total"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SleImpactedUsersClient object to a map representation for JSON marshaling.
func (s SleImpactedUsersClient) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.Degraded != nil {
		structMap["degraded"] = s.Degraded
	}
	if s.Duration != nil {
		structMap["duration"] = s.Duration
	}
	if s.Gateways != nil {
		structMap["gateways"] = s.Gateways
	}
	if s.Mac != nil {
		structMap["mac"] = s.Mac
	}
	if s.Name != nil {
		structMap["name"] = s.Name
	}
	if s.SrcIp != nil {
		structMap["src_ip"] = s.SrcIp
	}
	if s.Total != nil {
		structMap["total"] = s.Total
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SleImpactedUsersClient.
// It customizes the JSON unmarshaling process for SleImpactedUsersClient objects.
func (s *SleImpactedUsersClient) UnmarshalJSON(input []byte) error {
	var temp tempSleImpactedUsersClient
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "degraded", "duration", "gateways", "mac", "name", "src_ip", "total")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.Degraded = temp.Degraded
	s.Duration = temp.Duration
	s.Gateways = temp.Gateways
	s.Mac = temp.Mac
	s.Name = temp.Name
	s.SrcIp = temp.SrcIp
	s.Total = temp.Total
	return nil
}

// tempSleImpactedUsersClient is a temporary struct used for validating the fields of SleImpactedUsersClient.
type tempSleImpactedUsersClient struct {
	Degraded *float64                   `json:"degraded,omitempty"`
	Duration *float64                   `json:"duration,omitempty"`
	Gateways []SleImpactedClientGateway `json:"gateways,omitempty"`
	Mac      *string                    `json:"mac,omitempty"`
	Name     *string                    `json:"name,omitempty"`
	SrcIp    *string                    `json:"src_ip,omitempty"`
	Total    *float64                   `json:"total,omitempty"`
}
