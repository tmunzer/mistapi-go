// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// HaClusterNode represents a HaClusterNode struct.
type HaClusterNode struct {
	// only for HA. enum: `node0`, `node1`
	Node                 *HaClusterNodeEnum     `json:"node,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for HaClusterNode,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (h HaClusterNode) String() string {
	return fmt.Sprintf(
		"HaClusterNode[Node=%v, AdditionalProperties=%v]",
		h.Node, h.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for HaClusterNode.
// It customizes the JSON marshaling process for HaClusterNode objects.
func (h HaClusterNode) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(h.AdditionalProperties,
		"node"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(h.toMap())
}

// toMap converts the HaClusterNode object to a map representation for JSON marshaling.
func (h HaClusterNode) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, h.AdditionalProperties)
	if h.Node != nil {
		structMap["node"] = h.Node
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for HaClusterNode.
// It customizes the JSON unmarshaling process for HaClusterNode objects.
func (h *HaClusterNode) UnmarshalJSON(input []byte) error {
	var temp tempHaClusterNode
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "node")
	if err != nil {
		return err
	}
	h.AdditionalProperties = additionalProperties

	h.Node = temp.Node
	return nil
}

// tempHaClusterNode is a temporary struct used for validating the fields of HaClusterNode.
type tempHaClusterNode struct {
	Node *HaClusterNodeEnum `json:"node,omitempty"`
}
