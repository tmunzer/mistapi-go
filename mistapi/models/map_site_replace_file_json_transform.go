package models

import (
    "encoding/json"
    "fmt"
)

// MapSiteReplaceFileJsonTransform represents a MapSiteReplaceFileJsonTransform struct.
// If `transform` is provided, all the locations of the objects on the map (AP, Zone, Vbeacon, Beacon) will be transformed as well (relative to the new Map)
type MapSiteReplaceFileJsonTransform struct {
    // Whether to rotate the replacing image, in degrees
    Rotation             *float64               `json:"rotation,omitempty"`
    // Whether to scale the replacing image
    Scale                *float64               `json:"scale,omitempty"`
    // Where the (0, 0) of the new image is relative to the original map
    X                    *float64               `json:"x,omitempty"`
    // Where the (0, 0) of the new image is relative to the original map
    Y                    *float64               `json:"y,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for MapSiteReplaceFileJsonTransform,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m MapSiteReplaceFileJsonTransform) String() string {
    return fmt.Sprintf(
    	"MapSiteReplaceFileJsonTransform[Rotation=%v, Scale=%v, X=%v, Y=%v, AdditionalProperties=%v]",
    	m.Rotation, m.Scale, m.X, m.Y, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for MapSiteReplaceFileJsonTransform.
// It customizes the JSON marshaling process for MapSiteReplaceFileJsonTransform objects.
func (m MapSiteReplaceFileJsonTransform) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(m.AdditionalProperties,
        "rotation", "scale", "x", "y"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(m.toMap())
}

// toMap converts the MapSiteReplaceFileJsonTransform object to a map representation for JSON marshaling.
func (m MapSiteReplaceFileJsonTransform) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, m.AdditionalProperties)
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
    var temp tempMapSiteReplaceFileJsonTransform
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "rotation", "scale", "x", "y")
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

// tempMapSiteReplaceFileJsonTransform is a temporary struct used for validating the fields of MapSiteReplaceFileJsonTransform.
type tempMapSiteReplaceFileJsonTransform  struct {
    Rotation *float64 `json:"rotation,omitempty"`
    Scale    *float64 `json:"scale,omitempty"`
    X        *float64 `json:"x,omitempty"`
    Y        *float64 `json:"y,omitempty"`
}
