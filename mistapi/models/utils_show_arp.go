package models

import (
    "encoding/json"
)

// UtilsShowArp represents a UtilsShowArp struct.
type UtilsShowArp struct {
    // duration in sec for which refresh is enabled. Should be set only if interval is configured to non-zero value.
    Duration             *int           `json:"duration,omitempty"`
    // rate at which output will refresh
    Interval             *int           `json:"interval,omitempty"`
    // IP Address
    Ip                   *string        `json:"ip,omitempty"`
    // device Port ID
    PortId               *string        `json:"port_id,omitempty"`
    // VRF Name
    Vrf                  *string        `json:"vrf,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for UtilsShowArp.
// It customizes the JSON marshaling process for UtilsShowArp objects.
func (u UtilsShowArp) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(u.toMap())
}

// toMap converts the UtilsShowArp object to a map representation for JSON marshaling.
func (u UtilsShowArp) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, u.AdditionalProperties)
    if u.Duration != nil {
        structMap["duration"] = u.Duration
    }
    if u.Interval != nil {
        structMap["interval"] = u.Interval
    }
    if u.Ip != nil {
        structMap["ip"] = u.Ip
    }
    if u.PortId != nil {
        structMap["port_id"] = u.PortId
    }
    if u.Vrf != nil {
        structMap["vrf"] = u.Vrf
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UtilsShowArp.
// It customizes the JSON unmarshaling process for UtilsShowArp objects.
func (u *UtilsShowArp) UnmarshalJSON(input []byte) error {
    var temp tempUtilsShowArp
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "duration", "interval", "ip", "port_id", "vrf")
    if err != nil {
    	return err
    }
    
    u.AdditionalProperties = additionalProperties
    u.Duration = temp.Duration
    u.Interval = temp.Interval
    u.Ip = temp.Ip
    u.PortId = temp.PortId
    u.Vrf = temp.Vrf
    return nil
}

// tempUtilsShowArp is a temporary struct used for validating the fields of UtilsShowArp.
type tempUtilsShowArp  struct {
    Duration *int    `json:"duration,omitempty"`
    Interval *int    `json:"interval,omitempty"`
    Ip       *string `json:"ip,omitempty"`
    PortId   *string `json:"port_id,omitempty"`
    Vrf      *string `json:"vrf,omitempty"`
}
