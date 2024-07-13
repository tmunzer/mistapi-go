package models

import (
    "encoding/json"
)

// LocateSwitch represents a LocateSwitch struct.
type LocateSwitch struct {
    // minutes the leds should keep flashing
    Duration             *int           `json:"duration,omitempty"`
    // for virtual chassis, the MAC of the member
    Mac                  *string        `json:"mac,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for LocateSwitch.
// It customizes the JSON marshaling process for LocateSwitch objects.
func (l LocateSwitch) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(l.toMap())
}

// toMap converts the LocateSwitch object to a map representation for JSON marshaling.
func (l LocateSwitch) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, l.AdditionalProperties)
    if l.Duration != nil {
        structMap["duration"] = l.Duration
    }
    if l.Mac != nil {
        structMap["mac"] = l.Mac
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for LocateSwitch.
// It customizes the JSON unmarshaling process for LocateSwitch objects.
func (l *LocateSwitch) UnmarshalJSON(input []byte) error {
    var temp locateSwitch
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "duration", "mac")
    if err != nil {
    	return err
    }
    
    l.AdditionalProperties = additionalProperties
    l.Duration = temp.Duration
    l.Mac = temp.Mac
    return nil
}

// locateSwitch is a temporary struct used for validating the fields of LocateSwitch.
type locateSwitch  struct {
    Duration *int    `json:"duration,omitempty"`
    Mac      *string `json:"mac,omitempty"`
}
