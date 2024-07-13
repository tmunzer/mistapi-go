package models

import (
    "encoding/json"
)

// LatlngBr represents a LatlngBr struct.
// when type=google, latitude / longitude of the bottom-right corner
type LatlngBr struct {
    Lat                  *string        `json:"lat,omitempty"`
    Lng                  *string        `json:"lng,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for LatlngBr.
// It customizes the JSON marshaling process for LatlngBr objects.
func (l LatlngBr) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(l.toMap())
}

// toMap converts the LatlngBr object to a map representation for JSON marshaling.
func (l LatlngBr) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, l.AdditionalProperties)
    if l.Lat != nil {
        structMap["lat"] = l.Lat
    }
    if l.Lng != nil {
        structMap["lng"] = l.Lng
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for LatlngBr.
// It customizes the JSON unmarshaling process for LatlngBr objects.
func (l *LatlngBr) UnmarshalJSON(input []byte) error {
    var temp latlngBr
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "lat", "lng")
    if err != nil {
    	return err
    }
    
    l.AdditionalProperties = additionalProperties
    l.Lat = temp.Lat
    l.Lng = temp.Lng
    return nil
}

// latlngBr is a temporary struct used for validating the fields of LatlngBr.
type latlngBr  struct {
    Lat *string `json:"lat,omitempty"`
    Lng *string `json:"lng,omitempty"`
}
