package models

import (
    "encoding/json"
)

// MapSiteReplaceFileJsonTransform represents a MapSiteReplaceFileJsonTransform struct.
// If `transform` is provided, all the locations of the objects on the map (AP, Zone, Vbeacon, Beacon) will be transformed as well (relative to the new Map)
type MapSiteReplaceFileJsonTransform struct {
    // whether to rotate the replacing image, in degrees
    Rotation             *float64       `json:"rotation,omitempty"`
    // whether to scale the replacing image
    Scale                *float64       `json:"scale,omitempty"`
    // where the (0, 0) of the new image is relative to the original map
    X                    *float64       `json:"x,omitempty"`
    // where the (0, 0) of the new image is relative to the original map
    Y                    *float64       `json:"y,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for MapSiteReplaceFileJsonTransform.
// It customizes the JSON marshaling process for MapSiteReplaceFileJsonTransform objects.
func (m MapSiteReplaceFileJsonTransform) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(m.toMap())
}

// toMap converts the MapSiteReplaceFileJsonTransform object to a map representation for JSON marshaling.
func (m MapSiteReplaceFileJsonTransform) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, m.AdditionalProperties)
    if m.Rotation != nil {
        structMap["rotation"] = m.Rotation
    }
    if m.Scale != nil {
        structMap["scale"] = m.Scale
    }
    if m.X != nil {
        structMap["x"] = m.X
    }
    if m.Y != nil {
        structMap["y"] = m.Y
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MapSiteReplaceFileJsonTransform.
// It customizes the JSON unmarshaling process for MapSiteReplaceFileJsonTransform objects.
func (m *MapSiteReplaceFileJsonTransform) UnmarshalJSON(input []byte) error {
    var temp mapSiteReplaceFileJsonTransform
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "rotation", "scale", "x", "y")
    if err != nil {
    	return err
    }
    
    m.AdditionalProperties = additionalProperties
    m.Rotation = temp.Rotation
    m.Scale = temp.Scale
    m.X = temp.X
    m.Y = temp.Y
    return nil
}

// mapSiteReplaceFileJsonTransform is a temporary struct used for validating the fields of MapSiteReplaceFileJsonTransform.
type mapSiteReplaceFileJsonTransform  struct {
    Rotation *float64 `json:"rotation,omitempty"`
    Scale    *float64 `json:"scale,omitempty"`
    X        *float64 `json:"x,omitempty"`
    Y        *float64 `json:"y,omitempty"`
}
