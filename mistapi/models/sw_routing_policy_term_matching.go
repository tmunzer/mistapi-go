// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// SwRoutingPolicyTermMatching represents a SwRoutingPolicyTermMatching struct.
// zero or more criteria/filter can be specified to match the term, all criteria have to be met
type SwRoutingPolicyTermMatching struct {
	AsPath    []BgpAs  `json:"as_path,omitempty"`
	Community []string `json:"community,omitempty"`
	// zero or more criteria/filter can be specified to match the term, all criteria have to be met
	Prefix               []string                                  `json:"prefix,omitempty"`
	Protocol             []SwRoutingPolicyTermMatchingProtocolEnum `json:"protocol,omitempty"`
	AdditionalProperties map[string]interface{}                    `json:"_"`
}

// String implements the fmt.Stringer interface for SwRoutingPolicyTermMatching,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SwRoutingPolicyTermMatching) String() string {
	return fmt.Sprintf(
		"SwRoutingPolicyTermMatching[AsPath=%v, Community=%v, Prefix=%v, Protocol=%v, AdditionalProperties=%v]",
		s.AsPath, s.Community, s.Prefix, s.Protocol, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SwRoutingPolicyTermMatching.
// It customizes the JSON marshaling process for SwRoutingPolicyTermMatching objects.
func (s SwRoutingPolicyTermMatching) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"as_path", "community", "prefix", "protocol"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SwRoutingPolicyTermMatching object to a map representation for JSON marshaling.
func (s SwRoutingPolicyTermMatching) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.AsPath != nil {
		structMap["as_path"] = s.AsPath
	}
	if s.Community != nil {
		structMap["community"] = s.Community
	}
	if s.Prefix != nil {
		structMap["prefix"] = s.Prefix
	}
	if s.Protocol != nil {
		structMap["protocol"] = s.Protocol
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SwRoutingPolicyTermMatching.
// It customizes the JSON unmarshaling process for SwRoutingPolicyTermMatching objects.
func (s *SwRoutingPolicyTermMatching) UnmarshalJSON(input []byte) error {
	var temp tempSwRoutingPolicyTermMatching
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "as_path", "community", "prefix", "protocol")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.AsPath = temp.AsPath
	s.Community = temp.Community
	s.Prefix = temp.Prefix
	s.Protocol = temp.Protocol
	return nil
}

// tempSwRoutingPolicyTermMatching is a temporary struct used for validating the fields of SwRoutingPolicyTermMatching.
type tempSwRoutingPolicyTermMatching struct {
	AsPath    []BgpAs                                   `json:"as_path,omitempty"`
	Community []string                                  `json:"community,omitempty"`
	Prefix    []string                                  `json:"prefix,omitempty"`
	Protocol  []SwRoutingPolicyTermMatchingProtocolEnum `json:"protocol,omitempty"`
}
