package models

import (
    "encoding/json"
)

// RadiusConfig represents a RadiusConfig struct.
// Junos Radius config
type RadiusConfig struct {
    // how frequently should interim accounting be reported, 60-65535. default is 0 (use one specified in Access-Accept request from RADIUS Server). Very frequent messages can affect the performance of the radius server, 600 and up is recommended when enabled
    AcctInterimInterval  *int                   `json:"acct_interim_interval,omitempty"`
    AcctServers          []RadiusAcctServer     `json:"acct_servers,omitempty"`
    AuthServers          []RadiusAuthServer     `json:"auth_servers,omitempty"`
    // radius auth session retries
    AuthServersRetries   *int                   `json:"auth_servers_retries,omitempty"`
    // radius auth session timeout
    AuthServersTimeout   *int                   `json:"auth_servers_timeout,omitempty"`
    CoaEnabled           *bool                  `json:"coa_enabled,omitempty"`
    CoaPort              *int                   `json:"coa_port,omitempty"`
    // use `network`or `source_ip`
    // which network the RADIUS server resides, if there's static IP for this network, we'd use it as source-ip
    Network              *string                `json:"network,omitempty"`
    // use `network`or `source_ip`
    SourceIp             *string                `json:"source_ip,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for RadiusConfig.
// It customizes the JSON marshaling process for RadiusConfig objects.
func (r RadiusConfig) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "acct_interim_interval", "acct_servers", "auth_servers", "auth_servers_retries", "auth_servers_timeout", "coa_enabled", "coa_port", "network", "source_ip"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the RadiusConfig object to a map representation for JSON marshaling.
func (r RadiusConfig) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    if r.AcctInterimInterval != nil {
        structMap["acct_interim_interval"] = r.AcctInterimInterval
    }
    if r.AcctServers != nil {
        structMap["acct_servers"] = r.AcctServers
    }
    if r.AuthServers != nil {
        structMap["auth_servers"] = r.AuthServers
    }
    if r.AuthServersRetries != nil {
        structMap["auth_servers_retries"] = r.AuthServersRetries
    }
    if r.AuthServersTimeout != nil {
        structMap["auth_servers_timeout"] = r.AuthServersTimeout
    }
    if r.CoaEnabled != nil {
        structMap["coa_enabled"] = r.CoaEnabled
    }
    if r.CoaPort != nil {
        structMap["coa_port"] = r.CoaPort
    }
    if r.Network != nil {
        structMap["network"] = r.Network
    }
    if r.SourceIp != nil {
        structMap["source_ip"] = r.SourceIp
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RadiusConfig.
// It customizes the JSON unmarshaling process for RadiusConfig objects.
func (r *RadiusConfig) UnmarshalJSON(input []byte) error {
    var temp tempRadiusConfig
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "acct_interim_interval", "acct_servers", "auth_servers", "auth_servers_retries", "auth_servers_timeout", "coa_enabled", "coa_port", "network", "source_ip")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.AcctInterimInterval = temp.AcctInterimInterval
    r.AcctServers = temp.AcctServers
    r.AuthServers = temp.AuthServers
    r.AuthServersRetries = temp.AuthServersRetries
    r.AuthServersTimeout = temp.AuthServersTimeout
    r.CoaEnabled = temp.CoaEnabled
    r.CoaPort = temp.CoaPort
    r.Network = temp.Network
    r.SourceIp = temp.SourceIp
    return nil
}

// tempRadiusConfig is a temporary struct used for validating the fields of RadiusConfig.
type tempRadiusConfig  struct {
    AcctInterimInterval *int               `json:"acct_interim_interval,omitempty"`
    AcctServers         []RadiusAcctServer `json:"acct_servers,omitempty"`
    AuthServers         []RadiusAuthServer `json:"auth_servers,omitempty"`
    AuthServersRetries  *int               `json:"auth_servers_retries,omitempty"`
    AuthServersTimeout  *int               `json:"auth_servers_timeout,omitempty"`
    CoaEnabled          *bool              `json:"coa_enabled,omitempty"`
    CoaPort             *int               `json:"coa_port,omitempty"`
    Network             *string            `json:"network,omitempty"`
    SourceIp            *string            `json:"source_ip,omitempty"`
}
