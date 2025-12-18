// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// SwRoutingPolicyTerm represents a SwRoutingPolicyTerm struct.
type SwRoutingPolicyTerm struct {
	// When used as import policy
	Actions *SwRoutingPolicyTermAction `json:"actions,omitempty"`
	// zero or more criteria/filter can be specified to match the term, all criteria have to be met
	Matching             *SwRoutingPolicyTermMatching `json:"matching,omitempty"`
	Name                 string                       `json:"name"`
	AdditionalProperties map[string]interface{}       `json:"_"`
}

// String implements the fmt.Stringer interface for SwRoutingPolicyTerm,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SwRoutingPolicyTerm) String() string {
	return fmt.Sprintf(
		"SwRoutingPolicyTerm[Actions=%v, Matching=%v, Name=%v, AdditionalProperties=%v]",
		s.Actions, s.Matching, s.Name, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SwRoutingPolicyTerm.
// It customizes the JSON marshaling process for SwRoutingPolicyTerm objects.
func (s SwRoutingPolicyTerm) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"actions", "matching", "name"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SwRoutingPolicyTerm object to a map representation for JSON marshaling.
func (s SwRoutingPolicyTerm) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.Actions != nil {
		structMap["actions"] = s.Actions.toMap()
	}
	if s.Matching != nil {
		structMap["matching"] = s.Matching.toMap()
	}
	structMap["name"] = s.Name
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SwRoutingPolicyTerm.
// It customizes the JSON unmarshaling process for SwRoutingPolicyTerm objects.
func (s *SwRoutingPolicyTerm) UnmarshalJSON(input []byte) error {
	var temp tempSwRoutingPolicyTerm
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "actions", "matching", "name")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.Actions = temp.Actions
	s.Matching = temp.Matching
	s.Name = *temp.Name
	return nil
}

// tempSwRoutingPolicyTerm is a temporary struct used for validating the fields of SwRoutingPolicyTerm.
type tempSwRoutingPolicyTerm struct {
	Actions  *SwRoutingPolicyTermAction   `json:"actions,omitempty"`
	Matching *SwRoutingPolicyTermMatching `json:"matching,omitempty"`
	Name     *string                      `json:"name"`
}

func (s *tempSwRoutingPolicyTerm) validate() error {
	var errs []string
	if s.Name == nil {
		errs = append(errs, "required field `name` is missing for type `sw_routing_policy_term`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
