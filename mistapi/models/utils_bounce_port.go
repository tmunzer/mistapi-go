package models

import (
    "encoding/json"
)

// UtilsBouncePort represents a UtilsBouncePort struct.
type UtilsBouncePort struct {
    // the port to bounce
    Port                 *string        `json:"port,omitempty"`
    // list of ports to bounce
    Ports                []string       `json:"ports,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for UtilsBouncePort.
// It customizes the JSON marshaling process for UtilsBouncePort objects.
func (u UtilsBouncePort) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(u.toMap())
}

// toMap converts the UtilsBouncePort object to a map representation for JSON marshaling.
func (u UtilsBouncePort) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, u.AdditionalProperties)
    if u.Port != nil {
        structMap["port"] = u.Port
    }
    if u.Ports != nil {
        structMap["ports"] = u.Ports
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UtilsBouncePort.
// It customizes the JSON unmarshaling process for UtilsBouncePort objects.
func (u *UtilsBouncePort) UnmarshalJSON(input []byte) error {
    var temp tempUtilsBouncePort
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "port", "ports")
    if err != nil {
    	return err
    }
    
    u.AdditionalProperties = additionalProperties
    u.Port = temp.Port
    u.Ports = temp.Ports
    return nil
}

// tempUtilsBouncePort is a temporary struct used for validating the fields of UtilsBouncePort.
type tempUtilsBouncePort  struct {
    Port  *string  `json:"port,omitempty"`
    Ports []string `json:"ports,omitempty"`
}
