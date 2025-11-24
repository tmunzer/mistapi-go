// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// StatsMxedgeInactiveVlanStrs represents a StatsMxedgeInactiveVlanStrs struct.
// Inactive wired/L2TP VLANs. Entries can be individual VLANs or ranges.
type StatsMxedgeInactiveVlanStrs struct {
	// Inactive L2TP VLANs. Entries can be individual VLANs or ranges.
	L2tp []string `json:"l2tp,omitempty"`
	// Inactive wired VLANs. Entries can be individual VLANs or ranges.
	Wired                []string               `json:"wired,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for StatsMxedgeInactiveVlanStrs,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s StatsMxedgeInactiveVlanStrs) String() string {
	return fmt.Sprintf(
		"StatsMxedgeInactiveVlanStrs[L2tp=%v, Wired=%v, AdditionalProperties=%v]",
		s.L2tp, s.Wired, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for StatsMxedgeInactiveVlanStrs.
// It customizes the JSON marshaling process for StatsMxedgeInactiveVlanStrs objects.
func (s StatsMxedgeInactiveVlanStrs) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"l2tp", "wired"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the StatsMxedgeInactiveVlanStrs object to a map representation for JSON marshaling.
func (s StatsMxedgeInactiveVlanStrs) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.L2tp != nil {
		structMap["l2tp"] = s.L2tp
	}
	if s.Wired != nil {
		structMap["wired"] = s.Wired
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsMxedgeInactiveVlanStrs.
// It customizes the JSON unmarshaling process for StatsMxedgeInactiveVlanStrs objects.
func (s *StatsMxedgeInactiveVlanStrs) UnmarshalJSON(input []byte) error {
	var temp tempStatsMxedgeInactiveVlanStrs
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "l2tp", "wired")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.L2tp = temp.L2tp
	s.Wired = temp.Wired
	return nil
}

// tempStatsMxedgeInactiveVlanStrs is a temporary struct used for validating the fields of StatsMxedgeInactiveVlanStrs.
type tempStatsMxedgeInactiveVlanStrs struct {
	L2tp  []string `json:"l2tp,omitempty"`
	Wired []string `json:"wired,omitempty"`
}
