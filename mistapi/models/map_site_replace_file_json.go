package models

import (
    "encoding/json"
)

// MapSiteReplaceFileJson represents a MapSiteReplaceFileJson struct.
type MapSiteReplaceFileJson struct {
    // If `transform` is provided, all the locations of the objects on the map (AP, Zone, Vbeacon, Beacon) will be transformed as well (relative to the new Map)
    Transform            *MapSiteReplaceFileJsonTransform `json:"transform,omitempty"`
    AdditionalProperties map[string]interface{}           `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for MapSiteReplaceFileJson.
// It customizes the JSON marshaling process for MapSiteReplaceFileJson objects.
func (m MapSiteReplaceFileJson) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(m.AdditionalProperties,
        "transform"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(m.toMap())
}

// toMap converts the MapSiteReplaceFileJson object to a map representation for JSON marshaling.
func (m MapSiteReplaceFileJson) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, m.AdditionalProperties)
    if m.Transform != nil {
        structMap["transform"] = m.Transform.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MapSiteReplaceFileJson.
// It customizes the JSON unmarshaling process for MapSiteReplaceFileJson objects.
func (m *MapSiteReplaceFileJson) UnmarshalJSON(input []byte) error {
    var temp tempMapSiteReplaceFileJson
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "transform")
    if err != nil {
    	return err
    }
    m.AdditionalProperties = additionalProperties
    
    m.Transform = temp.Transform
    return nil
}

// tempMapSiteReplaceFileJson is a temporary struct used for validating the fields of MapSiteReplaceFileJson.
type tempMapSiteReplaceFileJson  struct {
    Transform *MapSiteReplaceFileJsonTransform `json:"transform,omitempty"`
}
