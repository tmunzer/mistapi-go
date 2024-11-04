package models

import (
    "encoding/json"
)

// UtilsShowEvpnDatabase represents a UtilsShowEvpnDatabase struct.
type UtilsShowEvpnDatabase struct {
    // duration in sec for which refresh is enabled. Should be set only if interval is configured to non-zero value.
    Duration             *int           `json:"duration,omitempty"`
    // rate at which output will refresh
    Interval             *int           `json:"interval,omitempty"`
    // client mac filter
    Mac                  *string        `json:"mac,omitempty"`
    // interface name
    PortId               *string        `json:"port_id,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for UtilsShowEvpnDatabase.
// It customizes the JSON marshaling process for UtilsShowEvpnDatabase objects.
func (u UtilsShowEvpnDatabase) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(u.toMap())
}

// toMap converts the UtilsShowEvpnDatabase object to a map representation for JSON marshaling.
func (u UtilsShowEvpnDatabase) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, u.AdditionalProperties)
    if u.Duration != nil {
        structMap["duration"] = u.Duration
    }
    if u.Interval != nil {
        structMap["interval"] = u.Interval
    }
    if u.Mac != nil {
        structMap["mac"] = u.Mac
    }
    if u.PortId != nil {
        structMap["port_id"] = u.PortId
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UtilsShowEvpnDatabase.
// It customizes the JSON unmarshaling process for UtilsShowEvpnDatabase objects.
func (u *UtilsShowEvpnDatabase) UnmarshalJSON(input []byte) error {
    var temp tempUtilsShowEvpnDatabase
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "duration", "interval", "mac", "port_id")
    if err != nil {
    	return err
    }
    
    u.AdditionalProperties = additionalProperties
    u.Duration = temp.Duration
    u.Interval = temp.Interval
    u.Mac = temp.Mac
    u.PortId = temp.PortId
    return nil
}

// tempUtilsShowEvpnDatabase is a temporary struct used for validating the fields of UtilsShowEvpnDatabase.
type tempUtilsShowEvpnDatabase  struct {
    Duration *int    `json:"duration,omitempty"`
    Interval *int    `json:"interval,omitempty"`
    Mac      *string `json:"mac,omitempty"`
    PortId   *string `json:"port_id,omitempty"`
}
