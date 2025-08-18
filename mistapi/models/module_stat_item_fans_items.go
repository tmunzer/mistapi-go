// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// ModuleStatItemFansItems represents a ModuleStatItemFansItems struct.
type ModuleStatItemFansItems struct {
	Airflow              *string                `json:"airflow,omitempty"`
	Name                 *string                `json:"name,omitempty"`
	Status               *string                `json:"status,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ModuleStatItemFansItems,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m ModuleStatItemFansItems) String() string {
	return fmt.Sprintf(
		"ModuleStatItemFansItems[Airflow=%v, Name=%v, Status=%v, AdditionalProperties=%v]",
		m.Airflow, m.Name, m.Status, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ModuleStatItemFansItems.
// It customizes the JSON marshaling process for ModuleStatItemFansItems objects.
func (m ModuleStatItemFansItems) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(m.AdditionalProperties,
		"airflow", "name", "status"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(m.toMap())
}

// toMap converts the ModuleStatItemFansItems object to a map representation for JSON marshaling.
func (m ModuleStatItemFansItems) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, m.AdditionalProperties)
	if m.Airflow != nil {
		structMap["airflow"] = m.Airflow
	}
	if m.Name != nil {
		structMap["name"] = m.Name
	}
	if m.Status != nil {
		structMap["status"] = m.Status
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ModuleStatItemFansItems.
// It customizes the JSON unmarshaling process for ModuleStatItemFansItems objects.
func (m *ModuleStatItemFansItems) UnmarshalJSON(input []byte) error {
	var temp tempModuleStatItemFansItems
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "airflow", "name", "status")
	if err != nil {
		return err
	}
	m.AdditionalProperties = additionalProperties

	m.Airflow = temp.Airflow
	m.Name = temp.Name
	m.Status = temp.Status
	return nil
}

// tempModuleStatItemFansItems is a temporary struct used for validating the fields of ModuleStatItemFansItems.
type tempModuleStatItemFansItems struct {
	Airflow *string `json:"airflow,omitempty"`
	Name    *string `json:"name,omitempty"`
	Status  *string `json:"status,omitempty"`
}
