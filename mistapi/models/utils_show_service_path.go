// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// UtilsShowServicePath represents a UtilsShowServicePath struct.
// The exact service name for which to display the service path
type UtilsShowServicePath struct {
	// only for HA. enum: `node0`, `node1`
	Node                 *HaClusterNodeEnum     `json:"node,omitempty"`
	ServiceName          *string                `json:"service_name,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for UtilsShowServicePath,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UtilsShowServicePath) String() string {
	return fmt.Sprintf(
		"UtilsShowServicePath[Node=%v, ServiceName=%v, AdditionalProperties=%v]",
		u.Node, u.ServiceName, u.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for UtilsShowServicePath.
// It customizes the JSON marshaling process for UtilsShowServicePath objects.
func (u UtilsShowServicePath) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(u.AdditionalProperties,
		"node", "service_name"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(u.toMap())
}

// toMap converts the UtilsShowServicePath object to a map representation for JSON marshaling.
func (u UtilsShowServicePath) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, u.AdditionalProperties)
	if u.Node != nil {
		structMap["node"] = u.Node
	}
	if u.ServiceName != nil {
		structMap["service_name"] = u.ServiceName
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UtilsShowServicePath.
// It customizes the JSON unmarshaling process for UtilsShowServicePath objects.
func (u *UtilsShowServicePath) UnmarshalJSON(input []byte) error {
	var temp tempUtilsShowServicePath
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "node", "service_name")
	if err != nil {
		return err
	}
	u.AdditionalProperties = additionalProperties

	u.Node = temp.Node
	u.ServiceName = temp.ServiceName
	return nil
}

// tempUtilsShowServicePath is a temporary struct used for validating the fields of UtilsShowServicePath.
type tempUtilsShowServicePath struct {
	Node        *HaClusterNodeEnum `json:"node,omitempty"`
	ServiceName *string            `json:"service_name,omitempty"`
}
