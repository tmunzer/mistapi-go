// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// MapWayfindingMicello represents a MapWayfindingMicello struct.
type MapWayfindingMicello struct {
	AccountKey           *string                `json:"account_key,omitempty"`
	DefaultLevelId       *int                   `json:"default_level_id,omitempty"`
	MapId                *string                `json:"map_id,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for MapWayfindingMicello,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m MapWayfindingMicello) String() string {
	return fmt.Sprintf(
		"MapWayfindingMicello[AccountKey=%v, DefaultLevelId=%v, MapId=%v, AdditionalProperties=%v]",
		m.AccountKey, m.DefaultLevelId, m.MapId, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for MapWayfindingMicello.
// It customizes the JSON marshaling process for MapWayfindingMicello objects.
func (m MapWayfindingMicello) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(m.AdditionalProperties,
		"account_key", "default_level_id", "map_id"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(m.toMap())
}

// toMap converts the MapWayfindingMicello object to a map representation for JSON marshaling.
func (m MapWayfindingMicello) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, m.AdditionalProperties)
	if m.AccountKey != nil {
		structMap["account_key"] = m.AccountKey
	}
	if m.DefaultLevelId != nil {
		structMap["default_level_id"] = m.DefaultLevelId
	}
	if m.MapId != nil {
		structMap["map_id"] = m.MapId
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MapWayfindingMicello.
// It customizes the JSON unmarshaling process for MapWayfindingMicello objects.
func (m *MapWayfindingMicello) UnmarshalJSON(input []byte) error {
	var temp tempMapWayfindingMicello
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "account_key", "default_level_id", "map_id")
	if err != nil {
		return err
	}
	m.AdditionalProperties = additionalProperties

	m.AccountKey = temp.AccountKey
	m.DefaultLevelId = temp.DefaultLevelId
	m.MapId = temp.MapId
	return nil
}

// tempMapWayfindingMicello is a temporary struct used for validating the fields of MapWayfindingMicello.
type tempMapWayfindingMicello struct {
	AccountKey     *string `json:"account_key,omitempty"`
	DefaultLevelId *int    `json:"default_level_id,omitempty"`
	MapId          *string `json:"map_id,omitempty"`
}
