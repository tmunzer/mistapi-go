package models

import (
    "encoding/json"
)

// LatlngBr represents a LatlngBr struct.
// when type=google, latitude / longitude of the bottom-right corner
type LatlngBr struct {
    Lat                  *string                `json:"lat,omitempty"`
    Lng                  *string                `json:"lng,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for LatlngBr.
// It customizes the JSON marshaling process for LatlngBr objects.
func (l LatlngBr) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(l.AdditionalProperties,
        "lat", "lng"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(l.toMap())
}

// toMap converts the LatlngBr object to a map representation for JSON marshaling.
func (l LatlngBr) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, l.AdditionalProperties)
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
    var temp tempLatlngBr
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "lat", "lng")
    if err != nil {
    	return err
    }
    l.AdditionalProperties = additionalProperties
    
    l.Lat = temp.Lat
    l.Lng = temp.Lng
    return nil
}

// tempLatlngBr is a temporary struct used for validating the fields of LatlngBr.
type tempLatlngBr  struct {
    Lat *string `json:"lat,omitempty"`
    Lng *string `json:"lng,omitempty"`
}
