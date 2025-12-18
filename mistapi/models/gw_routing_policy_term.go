// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// GwRoutingPolicyTerm represents a GwRoutingPolicyTerm struct.
type GwRoutingPolicyTerm struct {
	// When used as import policy
	Actions *GwRoutingPolicyTermAction `json:"actions,omitempty"`
	// zero or more criteria/filter can be specified to match the term, all criteria have to be met
	Matching             *GwRoutingPolicyTermMatching `json:"matching,omitempty"`
	AdditionalProperties map[string]interface{}       `json:"_"`
}

// String implements the fmt.Stringer interface for GwRoutingPolicyTerm,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (g GwRoutingPolicyTerm) String() string {
	return fmt.Sprintf(
		"GwRoutingPolicyTerm[Actions=%v, Matching=%v, AdditionalProperties=%v]",
		g.Actions, g.Matching, g.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for GwRoutingPolicyTerm.
// It customizes the JSON marshaling process for GwRoutingPolicyTerm objects.
func (g GwRoutingPolicyTerm) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(g.AdditionalProperties,
		"actions", "matching"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(g.toMap())
}

// toMap converts the GwRoutingPolicyTerm object to a map representation for JSON marshaling.
func (g GwRoutingPolicyTerm) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, g.AdditionalProperties)
	if g.Actions != nil {
		structMap["actions"] = g.Actions.toMap()
	}
	if g.Matching != nil {
		structMap["matching"] = g.Matching.toMap()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GwRoutingPolicyTerm.
// It customizes the JSON unmarshaling process for GwRoutingPolicyTerm objects.
func (g *GwRoutingPolicyTerm) UnmarshalJSON(input []byte) error {
	var temp tempGwRoutingPolicyTerm
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "actions", "matching")
	if err != nil {
		return err
	}
	g.AdditionalProperties = additionalProperties

	g.Actions = temp.Actions
	g.Matching = temp.Matching
	return nil
}

// tempGwRoutingPolicyTerm is a temporary struct used for validating the fields of GwRoutingPolicyTerm.
type tempGwRoutingPolicyTerm struct {
	Actions  *GwRoutingPolicyTermAction   `json:"actions,omitempty"`
	Matching *GwRoutingPolicyTermMatching `json:"matching,omitempty"`
}
