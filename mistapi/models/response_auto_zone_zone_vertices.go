package models

import (
    "encoding/json"
)

// ResponseAutoZoneZoneVertices represents a ResponseAutoZoneZoneVertices struct.
type ResponseAutoZoneZoneVertices struct {
    X                    *int           `json:"x,omitempty"`
    Y                    *int           `json:"y,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseAutoZoneZoneVertices.
// It customizes the JSON marshaling process for ResponseAutoZoneZoneVertices objects.
func (r ResponseAutoZoneZoneVertices) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseAutoZoneZoneVertices object to a map representation for JSON marshaling.
func (r ResponseAutoZoneZoneVertices) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    if r.X != nil {
        structMap["x"] = r.X
    }
    if r.Y != nil {
        structMap["y"] = r.Y
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseAutoZoneZoneVertices.
// It customizes the JSON unmarshaling process for ResponseAutoZoneZoneVertices objects.
func (r *ResponseAutoZoneZoneVertices) UnmarshalJSON(input []byte) error {
    var temp tempResponseAutoZoneZoneVertices
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "x", "y")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.X = temp.X
    r.Y = temp.Y
    return nil
}

// tempResponseAutoZoneZoneVertices is a temporary struct used for validating the fields of ResponseAutoZoneZoneVertices.
type tempResponseAutoZoneZoneVertices  struct {
    X *int `json:"x,omitempty"`
    Y *int `json:"y,omitempty"`
}
