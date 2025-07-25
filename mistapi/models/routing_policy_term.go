// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// RoutingPolicyTerm represents a RoutingPolicyTerm struct.
type RoutingPolicyTerm struct {
    // When used as import policy
    Action               *RoutingPolicyTermAction   `json:"action,omitempty"`
    // zero or more criteria/filter can be specified to match the term, all criteria have to be met
    Matching             *RoutingPolicyTermMatching `json:"matching,omitempty"`
    AdditionalProperties map[string]interface{}     `json:"_"`
}

// String implements the fmt.Stringer interface for RoutingPolicyTerm,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r RoutingPolicyTerm) String() string {
    return fmt.Sprintf(
    	"RoutingPolicyTerm[Action=%v, Matching=%v, AdditionalProperties=%v]",
    	r.Action, r.Matching, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for RoutingPolicyTerm.
// It customizes the JSON marshaling process for RoutingPolicyTerm objects.
func (r RoutingPolicyTerm) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "action", "matching"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the RoutingPolicyTerm object to a map representation for JSON marshaling.
func (r RoutingPolicyTerm) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    if r.Action != nil {
        structMap["action"] = r.Action.toMap()
    }
    if r.Matching != nil {
        structMap["matching"] = r.Matching.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RoutingPolicyTerm.
// It customizes the JSON unmarshaling process for RoutingPolicyTerm objects.
func (r *RoutingPolicyTerm) UnmarshalJSON(input []byte) error {
    var temp tempRoutingPolicyTerm
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "action", "matching")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.Action = temp.Action
    r.Matching = temp.Matching
    return nil
}

// tempRoutingPolicyTerm is a temporary struct used for validating the fields of RoutingPolicyTerm.
type tempRoutingPolicyTerm  struct {
    Action   *RoutingPolicyTermAction   `json:"action,omitempty"`
    Matching *RoutingPolicyTermMatching `json:"matching,omitempty"`
}
