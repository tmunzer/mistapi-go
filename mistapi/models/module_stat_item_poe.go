// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// ModuleStatItemPoe represents a ModuleStatItemPoe struct.
type ModuleStatItemPoe struct {
	MaxPower             *float64               `json:"max_power,omitempty"`
	PowerDraw            *float64               `json:"power_draw,omitempty"`
	Status               *string                `json:"status,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ModuleStatItemPoe,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m ModuleStatItemPoe) String() string {
	return fmt.Sprintf(
		"ModuleStatItemPoe[MaxPower=%v, PowerDraw=%v, Status=%v, AdditionalProperties=%v]",
		m.MaxPower, m.PowerDraw, m.Status, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ModuleStatItemPoe.
// It customizes the JSON marshaling process for ModuleStatItemPoe objects.
func (m ModuleStatItemPoe) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(m.AdditionalProperties,
		"max_power", "power_draw", "status"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(m.toMap())
}

// toMap converts the ModuleStatItemPoe object to a map representation for JSON marshaling.
func (m ModuleStatItemPoe) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, m.AdditionalProperties)
	if m.MaxPower != nil {
		structMap["max_power"] = m.MaxPower
	}
	if m.PowerDraw != nil {
		structMap["power_draw"] = m.PowerDraw
	}
	if m.Status != nil {
		structMap["status"] = m.Status
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ModuleStatItemPoe.
// It customizes the JSON unmarshaling process for ModuleStatItemPoe objects.
func (m *ModuleStatItemPoe) UnmarshalJSON(input []byte) error {
	var temp tempModuleStatItemPoe
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "max_power", "power_draw", "status")
	if err != nil {
		return err
	}
	m.AdditionalProperties = additionalProperties

	m.MaxPower = temp.MaxPower
	m.PowerDraw = temp.PowerDraw
	m.Status = temp.Status
	return nil
}

// tempModuleStatItemPoe is a temporary struct used for validating the fields of ModuleStatItemPoe.
type tempModuleStatItemPoe struct {
	MaxPower  *float64 `json:"max_power,omitempty"`
	PowerDraw *float64 `json:"power_draw,omitempty"`
	Status    *string  `json:"status,omitempty"`
}
