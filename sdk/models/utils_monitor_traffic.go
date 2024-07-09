package models

import (
    "encoding/json"
)

// UtilsMonitorTraffic represents a UtilsMonitorTraffic struct.
type UtilsMonitorTraffic struct {
    // port name, if no port input is provided then all ports will be monitored
    Port                 *string        `json:"port,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for UtilsMonitorTraffic.
// It customizes the JSON marshaling process for UtilsMonitorTraffic objects.
func (u UtilsMonitorTraffic) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(u.toMap())
}

// toMap converts the UtilsMonitorTraffic object to a map representation for JSON marshaling.
func (u UtilsMonitorTraffic) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, u.AdditionalProperties)
    if u.Port != nil {
        structMap["port"] = u.Port
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UtilsMonitorTraffic.
// It customizes the JSON unmarshaling process for UtilsMonitorTraffic objects.
func (u *UtilsMonitorTraffic) UnmarshalJSON(input []byte) error {
    var temp utilsMonitorTraffic
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "port")
    if err != nil {
    	return err
    }
    
    u.AdditionalProperties = additionalProperties
    u.Port = temp.Port
    return nil
}

// utilsMonitorTraffic is a temporary struct used for validating the fields of UtilsMonitorTraffic.
type utilsMonitorTraffic  struct {
    Port *string `json:"port,omitempty"`
}
