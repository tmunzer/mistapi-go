package models

import (
    "encoding/json"
)

// SwitchMgmt represents a SwitchMgmt struct.
type SwitchMgmt struct {
    ConfigRevert         *int           `json:"config_revert,omitempty"`
    // restrict inbound-traffic to host
    // when enabled, all traffic that is not essential to our operation will be dropped 
    // e.g. ntp / dns / traffic to mist will be allowed by default, if dhcpd is enabled, we'll make sure it works
    ProtectRe            *ProtectRe     `json:"protect_re,omitempty"`
    RootPassword         *string        `json:"root_password,omitempty"`
    Tacacs               *Tacacs        `json:"tacacs,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SwitchMgmt.
// It customizes the JSON marshaling process for SwitchMgmt objects.
func (s SwitchMgmt) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SwitchMgmt object to a map representation for JSON marshaling.
func (s SwitchMgmt) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.ConfigRevert != nil {
        structMap["config_revert"] = s.ConfigRevert
    }
    if s.ProtectRe != nil {
        structMap["protect_re"] = s.ProtectRe.toMap()
    }
    if s.RootPassword != nil {
        structMap["root_password"] = s.RootPassword
    }
    if s.Tacacs != nil {
        structMap["tacacs"] = s.Tacacs.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SwitchMgmt.
// It customizes the JSON unmarshaling process for SwitchMgmt objects.
func (s *SwitchMgmt) UnmarshalJSON(input []byte) error {
    var temp switchMgmt
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "config_revert", "protect_re", "root_password", "tacacs")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.ConfigRevert = temp.ConfigRevert
    s.ProtectRe = temp.ProtectRe
    s.RootPassword = temp.RootPassword
    s.Tacacs = temp.Tacacs
    return nil
}

// switchMgmt is a temporary struct used for validating the fields of SwitchMgmt.
type switchMgmt  struct {
    ConfigRevert *int       `json:"config_revert,omitempty"`
    ProtectRe    *ProtectRe `json:"protect_re,omitempty"`
    RootPassword *string    `json:"root_password,omitempty"`
    Tacacs       *Tacacs    `json:"tacacs,omitempty"`
}
