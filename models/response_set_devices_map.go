package models

import (
    "encoding/json"
)

// ResponseSetDevicesMap represents a ResponseSetDevicesMap struct.
type ResponseSetDevicesMap struct {
    Locked               []string       `json:"locked,omitempty"`
    Moved                []string       `json:"moved,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseSetDevicesMap.
// It customizes the JSON marshaling process for ResponseSetDevicesMap objects.
func (r ResponseSetDevicesMap) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseSetDevicesMap object to a map representation for JSON marshaling.
func (r ResponseSetDevicesMap) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    if r.Locked != nil {
        structMap["locked"] = r.Locked
    }
    if r.Moved != nil {
        structMap["moved"] = r.Moved
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseSetDevicesMap.
// It customizes the JSON unmarshaling process for ResponseSetDevicesMap objects.
func (r *ResponseSetDevicesMap) UnmarshalJSON(input []byte) error {
    var temp responseSetDevicesMap
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "locked", "moved")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.Locked = temp.Locked
    r.Moved = temp.Moved
    return nil
}

// responseSetDevicesMap is a temporary struct used for validating the fields of ResponseSetDevicesMap.
type responseSetDevicesMap  struct {
    Locked []string `json:"locked,omitempty"`
    Moved  []string `json:"moved,omitempty"`
}
