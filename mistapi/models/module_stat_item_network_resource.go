// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// ModuleStatItemNetworkResource represents a ModuleStatItemNetworkResource struct.
type ModuleStatItemNetworkResource struct {
	// current usage of the network resource
	Count *int `json:"count,omitempty"`
	// maximum usage of the network resource
	Limit *int `json:"limit,omitempty"`
	// type of the network resource (e.g. FIB, FLOW, ...)
	Type                 *string                `json:"type,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ModuleStatItemNetworkResource,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m ModuleStatItemNetworkResource) String() string {
	return fmt.Sprintf(
		"ModuleStatItemNetworkResource[Count=%v, Limit=%v, Type=%v, AdditionalProperties=%v]",
		m.Count, m.Limit, m.Type, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ModuleStatItemNetworkResource.
// It customizes the JSON marshaling process for ModuleStatItemNetworkResource objects.
func (m ModuleStatItemNetworkResource) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(m.AdditionalProperties,
		"count", "limit", "type"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(m.toMap())
}

// toMap converts the ModuleStatItemNetworkResource object to a map representation for JSON marshaling.
func (m ModuleStatItemNetworkResource) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, m.AdditionalProperties)
	if m.Count != nil {
		structMap["count"] = m.Count
	}
	if m.Limit != nil {
		structMap["limit"] = m.Limit
	}
	if m.Type != nil {
		structMap["type"] = m.Type
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ModuleStatItemNetworkResource.
// It customizes the JSON unmarshaling process for ModuleStatItemNetworkResource objects.
func (m *ModuleStatItemNetworkResource) UnmarshalJSON(input []byte) error {
	var temp tempModuleStatItemNetworkResource
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "count", "limit", "type")
	if err != nil {
		return err
	}
	m.AdditionalProperties = additionalProperties

	m.Count = temp.Count
	m.Limit = temp.Limit
	m.Type = temp.Type
	return nil
}

// tempModuleStatItemNetworkResource is a temporary struct used for validating the fields of ModuleStatItemNetworkResource.
type tempModuleStatItemNetworkResource struct {
	Count *int    `json:"count,omitempty"`
	Limit *int    `json:"limit,omitempty"`
	Type  *string `json:"type,omitempty"`
}
