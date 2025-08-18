// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// ShellNode represents a ShellNode struct.
type ShellNode struct {
	// only for HA. enum: `node0`, `node1`
	Node                 *HaClusterNodeEnum     `json:"node,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ShellNode,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s ShellNode) String() string {
	return fmt.Sprintf(
		"ShellNode[Node=%v, AdditionalProperties=%v]",
		s.Node, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ShellNode.
// It customizes the JSON marshaling process for ShellNode objects.
func (s ShellNode) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"node"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the ShellNode object to a map representation for JSON marshaling.
func (s ShellNode) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.Node != nil {
		structMap["node"] = s.Node
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ShellNode.
// It customizes the JSON unmarshaling process for ShellNode objects.
func (s *ShellNode) UnmarshalJSON(input []byte) error {
	var temp tempShellNode
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "node")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.Node = temp.Node
	return nil
}

// tempShellNode is a temporary struct used for validating the fields of ShellNode.
type tempShellNode struct {
	Node *HaClusterNodeEnum `json:"node,omitempty"`
}
