// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// MapWayfinding represents a MapWayfinding struct.
// Properties related to wayfinding
type MapWayfinding struct {
    Micello              *MapWayfindingMicello  `json:"micello,omitempty"`
    SnapToPath           *bool                  `json:"snap_to_path,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for MapWayfinding,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m MapWayfinding) String() string {
    return fmt.Sprintf(
    	"MapWayfinding[Micello=%v, SnapToPath=%v, AdditionalProperties=%v]",
    	m.Micello, m.SnapToPath, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for MapWayfinding.
// It customizes the JSON marshaling process for MapWayfinding objects.
func (m MapWayfinding) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(m.AdditionalProperties,
        "micello", "snap_to_path"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(m.toMap())
}

// toMap converts the MapWayfinding object to a map representation for JSON marshaling.
func (m MapWayfinding) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, m.AdditionalProperties)
    if m.Micello != nil {
        structMap["micello"] = m.Micello.toMap()
    }
    if m.SnapToPath != nil {
        structMap["snap_to_path"] = m.SnapToPath
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MapWayfinding.
// It customizes the JSON unmarshaling process for MapWayfinding objects.
func (m *MapWayfinding) UnmarshalJSON(input []byte) error {
    var temp tempMapWayfinding
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "micello", "snap_to_path")
    if err != nil {
    	return err
    }
    m.AdditionalProperties = additionalProperties
    
    m.Micello = temp.Micello
    m.SnapToPath = temp.SnapToPath
    return nil
}

// tempMapWayfinding is a temporary struct used for validating the fields of MapWayfinding.
type tempMapWayfinding  struct {
    Micello    *MapWayfindingMicello `json:"micello,omitempty"`
    SnapToPath *bool                 `json:"snap_to_path,omitempty"`
}
