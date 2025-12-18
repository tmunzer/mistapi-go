// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// SwRoutingPolicy represents a SwRoutingPolicy struct.
type SwRoutingPolicy struct {
	// at least criteria/filter must be specified to match the term, all criteria have to be met
	Terms                []SwRoutingPolicyTerm  `json:"terms,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SwRoutingPolicy,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SwRoutingPolicy) String() string {
	return fmt.Sprintf(
		"SwRoutingPolicy[Terms=%v, AdditionalProperties=%v]",
		s.Terms, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SwRoutingPolicy.
// It customizes the JSON marshaling process for SwRoutingPolicy objects.
func (s SwRoutingPolicy) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"terms"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SwRoutingPolicy object to a map representation for JSON marshaling.
func (s SwRoutingPolicy) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.Terms != nil {
		structMap["terms"] = s.Terms
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SwRoutingPolicy.
// It customizes the JSON unmarshaling process for SwRoutingPolicy objects.
func (s *SwRoutingPolicy) UnmarshalJSON(input []byte) error {
	var temp tempSwRoutingPolicy
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "terms")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.Terms = temp.Terms
	return nil
}

// tempSwRoutingPolicy is a temporary struct used for validating the fields of SwRoutingPolicy.
type tempSwRoutingPolicy struct {
	Terms []SwRoutingPolicyTerm `json:"terms,omitempty"`
}
