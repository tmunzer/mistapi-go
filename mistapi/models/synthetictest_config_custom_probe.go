// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// SynthetictestConfigCustomProbe represents a SynthetictestConfigCustomProbe struct.
type SynthetictestConfigCustomProbe struct {
	// enum: `auto`, `high`, `low`
	Aggressiveness *SynthetictestConfigAggressivenessEnum `json:"aggressiveness,omitempty"`
	// If `type`==`icmp` or `type`==`tcp`, Host to be used for the custom probe
	Host *string `json:"host,omitempty"`
	// If `type`==`tcp`, Port to be used for the custom probe
	Port *int `json:"port,omitempty"`
	// In milliseconds
	Threshold *int `json:"threshold,omitempty"`
	// enum: `curl`, `icmp`, `tcp`
	Type *SynthetictestConfigCustomProbeTypeEnum `json:"type,omitempty"`
	// If `type`==`curl`, URL to be used for the custom probe, can be url or IP
	Url                  *string                `json:"url,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SynthetictestConfigCustomProbe,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SynthetictestConfigCustomProbe) String() string {
	return fmt.Sprintf(
		"SynthetictestConfigCustomProbe[Aggressiveness=%v, Host=%v, Port=%v, Threshold=%v, Type=%v, Url=%v, AdditionalProperties=%v]",
		s.Aggressiveness, s.Host, s.Port, s.Threshold, s.Type, s.Url, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SynthetictestConfigCustomProbe.
// It customizes the JSON marshaling process for SynthetictestConfigCustomProbe objects.
func (s SynthetictestConfigCustomProbe) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"aggressiveness", "host", "port", "threshold", "type", "url"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SynthetictestConfigCustomProbe object to a map representation for JSON marshaling.
func (s SynthetictestConfigCustomProbe) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.Aggressiveness != nil {
		structMap["aggressiveness"] = s.Aggressiveness
	}
	if s.Host != nil {
		structMap["host"] = s.Host
	}
	if s.Port != nil {
		structMap["port"] = s.Port
	}
	if s.Threshold != nil {
		structMap["threshold"] = s.Threshold
	}
	if s.Type != nil {
		structMap["type"] = s.Type
	}
	if s.Url != nil {
		structMap["url"] = s.Url
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SynthetictestConfigCustomProbe.
// It customizes the JSON unmarshaling process for SynthetictestConfigCustomProbe objects.
func (s *SynthetictestConfigCustomProbe) UnmarshalJSON(input []byte) error {
	var temp tempSynthetictestConfigCustomProbe
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "aggressiveness", "host", "port", "threshold", "type", "url")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.Aggressiveness = temp.Aggressiveness
	s.Host = temp.Host
	s.Port = temp.Port
	s.Threshold = temp.Threshold
	s.Type = temp.Type
	s.Url = temp.Url
	return nil
}

// tempSynthetictestConfigCustomProbe is a temporary struct used for validating the fields of SynthetictestConfigCustomProbe.
type tempSynthetictestConfigCustomProbe struct {
	Aggressiveness *SynthetictestConfigAggressivenessEnum  `json:"aggressiveness,omitempty"`
	Host           *string                                 `json:"host,omitempty"`
	Port           *int                                    `json:"port,omitempty"`
	Threshold      *int                                    `json:"threshold,omitempty"`
	Type           *SynthetictestConfigCustomProbeTypeEnum `json:"type,omitempty"`
	Url            *string                                 `json:"url,omitempty"`
}
