package models

import (
    "encoding/json"
)

// ResponseAutoZoneZoneVertex represents a ResponseAutoZoneZoneVertex struct.
type ResponseAutoZoneZoneVertex struct {
    X                    *int                   `json:"x,omitempty"`
    Y                    *int                   `json:"y,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseAutoZoneZoneVertex.
// It customizes the JSON marshaling process for ResponseAutoZoneZoneVertex objects.
func (r ResponseAutoZoneZoneVertex) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "x", "y"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseAutoZoneZoneVertex object to a map representation for JSON marshaling.
func (r ResponseAutoZoneZoneVertex) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    if r.X != nil {
        structMap["x"] = r.X
    }
    if r.Y != nil {
        structMap["y"] = r.Y
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseAutoZoneZoneVertex.
// It customizes the JSON unmarshaling process for ResponseAutoZoneZoneVertex objects.
func (r *ResponseAutoZoneZoneVertex) UnmarshalJSON(input []byte) error {
    var temp tempResponseAutoZoneZoneVertex
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "x", "y")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.X = temp.X
    r.Y = temp.Y
    return nil
}

// tempResponseAutoZoneZoneVertex is a temporary struct used for validating the fields of ResponseAutoZoneZoneVertex.
type tempResponseAutoZoneZoneVertex  struct {
    X *int `json:"x,omitempty"`
    Y *int `json:"y,omitempty"`
}
