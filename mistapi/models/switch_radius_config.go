package models

import (
    "encoding/json"
)

// SwitchRadiusConfig represents a SwitchRadiusConfig struct.
// Junos Radius config
type SwitchRadiusConfig struct {
    // how frequently should interim accounting be reported, 60-65535. default is 0 (use one specified in Access-Accept request from RADIUS Server). Very frequent messages can affect the performance of the radius server, 600 and up is recommended when enabled
    AcctInterimInterval  *int               `json:"acct_interim_interval,omitempty"`
    AcctServers          []RadiusAcctServer `json:"acct_servers,omitempty"`
    AuthServers          []RadiusAuthServer `json:"auth_servers,omitempty"`
    // radius auth session retries
    AuthServersRetries   *int               `json:"auth_servers_retries,omitempty"`
    // radius auth session timeout
    AuthServersTimeout   *int               `json:"auth_servers_timeout,omitempty"`
    // use `network`or `source_ip`
    // which network the RADIUS server resides, if there's static IP for this network, we'd use it as source-ip
    Network              *string            `json:"network,omitempty"`
    // use `network`or `source_ip`
    SourceIp             *string            `json:"source_ip,omitempty"`
    AdditionalProperties map[string]any     `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SwitchRadiusConfig.
// It customizes the JSON marshaling process for SwitchRadiusConfig objects.
func (s SwitchRadiusConfig) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SwitchRadiusConfig object to a map representation for JSON marshaling.
func (s SwitchRadiusConfig) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.AcctInterimInterval != nil {
        structMap["acct_interim_interval"] = s.AcctInterimInterval
    }
    if s.AcctServers != nil {
        structMap["acct_servers"] = s.AcctServers
    }
    if s.AuthServers != nil {
        structMap["auth_servers"] = s.AuthServers
    }
    if s.AuthServersRetries != nil {
        structMap["auth_servers_retries"] = s.AuthServersRetries
    }
    if s.AuthServersTimeout != nil {
        structMap["auth_servers_timeout"] = s.AuthServersTimeout
    }
    if s.Network != nil {
        structMap["network"] = s.Network
    }
    if s.SourceIp != nil {
        structMap["source_ip"] = s.SourceIp
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SwitchRadiusConfig.
// It customizes the JSON unmarshaling process for SwitchRadiusConfig objects.
func (s *SwitchRadiusConfig) UnmarshalJSON(input []byte) error {
    var temp tempSwitchRadiusConfig
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "acct_interim_interval", "acct_servers", "auth_servers", "auth_servers_retries", "auth_servers_timeout", "network", "source_ip")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.AcctInterimInterval = temp.AcctInterimInterval
    s.AcctServers = temp.AcctServers
    s.AuthServers = temp.AuthServers
    s.AuthServersRetries = temp.AuthServersRetries
    s.AuthServersTimeout = temp.AuthServersTimeout
    s.Network = temp.Network
    s.SourceIp = temp.SourceIp
    return nil
}

// tempSwitchRadiusConfig is a temporary struct used for validating the fields of SwitchRadiusConfig.
type tempSwitchRadiusConfig  struct {
    AcctInterimInterval *int               `json:"acct_interim_interval,omitempty"`
    AcctServers         []RadiusAcctServer `json:"acct_servers,omitempty"`
    AuthServers         []RadiusAuthServer `json:"auth_servers,omitempty"`
    AuthServersRetries  *int               `json:"auth_servers_retries,omitempty"`
    AuthServersTimeout  *int               `json:"auth_servers_timeout,omitempty"`
    Network             *string            `json:"network,omitempty"`
    SourceIp            *string            `json:"source_ip,omitempty"`
}
