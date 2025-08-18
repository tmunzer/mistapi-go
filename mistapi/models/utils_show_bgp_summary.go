// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// UtilsShowBgpSummary represents a UtilsShowBgpSummary struct.
type UtilsShowBgpSummary struct {
	// only for HA. enum: `node0`, `node1`
	Node                 *HaClusterNodeEnum     `json:"node,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for UtilsShowBgpSummary,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UtilsShowBgpSummary) String() string {
	return fmt.Sprintf(
		"UtilsShowBgpSummary[Node=%v, AdditionalProperties=%v]",
		u.Node, u.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for UtilsShowBgpSummary.
// It customizes the JSON marshaling process for UtilsShowBgpSummary objects.
func (u UtilsShowBgpSummary) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(u.AdditionalProperties,
		"node"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(u.toMap())
}

// toMap converts the UtilsShowBgpSummary object to a map representation for JSON marshaling.
func (u UtilsShowBgpSummary) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, u.AdditionalProperties)
	if u.Node != nil {
		structMap["node"] = u.Node
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UtilsShowBgpSummary.
// It customizes the JSON unmarshaling process for UtilsShowBgpSummary objects.
func (u *UtilsShowBgpSummary) UnmarshalJSON(input []byte) error {
	var temp tempUtilsShowBgpSummary
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "node")
	if err != nil {
		return err
	}
	u.AdditionalProperties = additionalProperties

	u.Node = temp.Node
	return nil
}

// tempUtilsShowBgpSummary is a temporary struct used for validating the fields of UtilsShowBgpSummary.
type tempUtilsShowBgpSummary struct {
	Node *HaClusterNodeEnum `json:"node,omitempty"`
}
