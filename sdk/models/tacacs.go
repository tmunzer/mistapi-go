package models

import (
    "encoding/json"
)

// Tacacs represents a Tacacs struct.
type Tacacs struct {
    AcctServers          []TacacsAcctServer     `json:"acct_servers,omitempty"`
    DefaultRole          *TacacsDefaultRoleEnum `json:"default_role,omitempty"`
    Enabled              *bool                  `json:"enabled,omitempty"`
    // which network the TACACS server resides
    Network              *string                `json:"network,omitempty"`
    TacplusServers       []TacacsAuthServer     `json:"tacplus_servers,omitempty"`
    AdditionalProperties map[string]any         `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for Tacacs.
// It customizes the JSON marshaling process for Tacacs objects.
func (t Tacacs) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(t.toMap())
}

// toMap converts the Tacacs object to a map representation for JSON marshaling.
func (t Tacacs) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, t.AdditionalProperties)
    if t.AcctServers != nil {
        structMap["acct_servers"] = t.AcctServers
    }
    if t.DefaultRole != nil {
        structMap["default_role"] = t.DefaultRole
    }
    if t.Enabled != nil {
        structMap["enabled"] = t.Enabled
    }
    if t.Network != nil {
        structMap["network"] = t.Network
    }
    if t.TacplusServers != nil {
        structMap["tacplus_servers"] = t.TacplusServers
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Tacacs.
// It customizes the JSON unmarshaling process for Tacacs objects.
func (t *Tacacs) UnmarshalJSON(input []byte) error {
    var temp tacacs
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "acct_servers", "default_role", "enabled", "network", "tacplus_servers")
    if err != nil {
    	return err
    }
    
    t.AdditionalProperties = additionalProperties
    t.AcctServers = temp.AcctServers
    t.DefaultRole = temp.DefaultRole
    t.Enabled = temp.Enabled
    t.Network = temp.Network
    t.TacplusServers = temp.TacplusServers
    return nil
}

// tacacs is a temporary struct used for validating the fields of Tacacs.
type tacacs  struct {
    AcctServers    []TacacsAcctServer     `json:"acct_servers,omitempty"`
    DefaultRole    *TacacsDefaultRoleEnum `json:"default_role,omitempty"`
    Enabled        *bool                  `json:"enabled,omitempty"`
    Network        *string                `json:"network,omitempty"`
    TacplusServers []TacacsAuthServer     `json:"tacplus_servers,omitempty"`
}
