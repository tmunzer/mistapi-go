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
	// Can be URL (e.g. http://x.com, https://x.com:8080/path/to/resource), IP address, or IP:port combination
	Target *string `json:"target,omitempty"`
	// In milliseconds
	Threshold *int `json:"threshold,omitempty"`
	// enum: `application`, `curl`, `icmp`, `reachability`, `tcp`
	Type                 *SynthetictestConfigCustomProbeTypeEnum `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}                  `json:"_"`
}

// String implements the fmt.Stringer interface for SynthetictestConfigCustomProbe,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SynthetictestConfigCustomProbe) String() string {
	return fmt.Sprintf(
		"SynthetictestConfigCustomProbe[Aggressiveness=%v, Target=%v, Threshold=%v, Type=%v, AdditionalProperties=%v]",
		s.Aggressiveness, s.Target, s.Threshold, s.Type, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SynthetictestConfigCustomProbe.
// It customizes the JSON marshaling process for SynthetictestConfigCustomProbe objects.
func (s SynthetictestConfigCustomProbe) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"aggressiveness", "target", "threshold", "type"); err != nil {
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
	if s.Target != nil {
		structMap["target"] = s.Target
	}
	if s.Threshold != nil {
		structMap["threshold"] = s.Threshold
	}
	if s.Type != nil {
		structMap["type"] = s.Type
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
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "aggressiveness", "target", "threshold", "type")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.Aggressiveness = temp.Aggressiveness
	s.Target = temp.Target
	s.Threshold = temp.Threshold
	s.Type = temp.Type
	return nil
}

// tempSynthetictestConfigCustomProbe is a temporary struct used for validating the fields of SynthetictestConfigCustomProbe.
type tempSynthetictestConfigCustomProbe struct {
	Aggressiveness *SynthetictestConfigAggressivenessEnum  `json:"aggressiveness,omitempty"`
	Target         *string                                 `json:"target,omitempty"`
	Threshold      *int                                    `json:"threshold,omitempty"`
	Type           *SynthetictestConfigCustomProbeTypeEnum `json:"type,omitempty"`
}
