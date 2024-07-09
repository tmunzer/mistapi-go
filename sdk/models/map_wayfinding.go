package models

import (
    "encoding/json"
)

// MapWayfinding represents a MapWayfinding struct.
// properties related to wayfinding
type MapWayfinding struct {
    Micello              *MapWayfindingMicello `json:"micello,omitempty"`
    SnapToPath           *bool                 `json:"snap_to_path,omitempty"`
    AdditionalProperties map[string]any        `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for MapWayfinding.
// It customizes the JSON marshaling process for MapWayfinding objects.
func (m MapWayfinding) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(m.toMap())
}

// toMap converts the MapWayfinding object to a map representation for JSON marshaling.
func (m MapWayfinding) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, m.AdditionalProperties)
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
    var temp mapWayfinding
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "micello", "snap_to_path")
    if err != nil {
    	return err
    }
    
    m.AdditionalProperties = additionalProperties
    m.Micello = temp.Micello
    m.SnapToPath = temp.SnapToPath
    return nil
}

// mapWayfinding is a temporary struct used for validating the fields of MapWayfinding.
type mapWayfinding  struct {
    Micello    *MapWayfindingMicello `json:"micello,omitempty"`
    SnapToPath *bool                 `json:"snap_to_path,omitempty"`
}
