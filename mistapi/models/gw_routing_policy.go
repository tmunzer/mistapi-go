// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// GwRoutingPolicy represents a GwRoutingPolicy struct.
type GwRoutingPolicy struct {
	// zero or more criteria/filter can be specified to match the term, all criteria have to be met
	Terms                []GwRoutingPolicyTerm  `json:"terms,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for GwRoutingPolicy,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (g GwRoutingPolicy) String() string {
	return fmt.Sprintf(
		"GwRoutingPolicy[Terms=%v, AdditionalProperties=%v]",
		g.Terms, g.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for GwRoutingPolicy.
// It customizes the JSON marshaling process for GwRoutingPolicy objects.
func (g GwRoutingPolicy) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(g.AdditionalProperties,
		"terms"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(g.toMap())
}

// toMap converts the GwRoutingPolicy object to a map representation for JSON marshaling.
func (g GwRoutingPolicy) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, g.AdditionalProperties)
	if g.Terms != nil {
		structMap["terms"] = g.Terms
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GwRoutingPolicy.
// It customizes the JSON unmarshaling process for GwRoutingPolicy objects.
func (g *GwRoutingPolicy) UnmarshalJSON(input []byte) error {
	var temp tempGwRoutingPolicy
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "terms")
	if err != nil {
		return err
	}
	g.AdditionalProperties = additionalProperties

	g.Terms = temp.Terms
	return nil
}

// tempGwRoutingPolicy is a temporary struct used for validating the fields of GwRoutingPolicy.
type tempGwRoutingPolicy struct {
	Terms []GwRoutingPolicyTerm `json:"terms,omitempty"`
}
