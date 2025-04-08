package models

import (
    "encoding/json"
    "fmt"
)

// SwitchRadiusConfig represents a SwitchRadiusConfig struct.
// Junos Radius config
type SwitchRadiusConfig struct {
    AcctImmediateUpdate  *bool                                      `json:"acct_immediate_update,omitempty"`
    // How frequently should interim accounting be reported, 60-65535. default is 0 (use one specified in Access-Accept request from RADIUS Server). Very frequent messages can affect the performance of the radius server, 600 and up is recommended when enabled
    AcctInterimInterval  *int                                       `json:"acct_interim_interval,omitempty"`
    AcctServers          []RadiusAcctServer                         `json:"acct_servers,omitempty"`
    // enum: `ordered`, `unordered`
    AuthServerSelection  *SwitchRadiusConfigAuthServerSelectionEnum `json:"auth_server_selection,omitempty"`
    AuthServers          []RadiusAuthServer                         `json:"auth_servers,omitempty"`
    // Radius auth session retries
    AuthServersRetries   *int                                       `json:"auth_servers_retries,omitempty"`
    // Radius auth session timeout
    AuthServersTimeout   *int                                       `json:"auth_servers_timeout,omitempty"`
    CoaEnabled           *bool                                      `json:"coa_enabled,omitempty"`
    // CoA Port, value from 1 to 65535, default is 3799
    CoaPort              *CoaPort                                   `json:"coa_port,omitempty"`
    FastDot1xTimers      *bool                                      `json:"fast_dot1x_timers,omitempty"`
    // Use `network`or `source_ip`. Which network the RADIUS server resides, if there's static IP for this network, we'd use it as source-ip
    Network              *string                                    `json:"network,omitempty"`
    // Use `network`or `source_ip`
    SourceIp             *string                                    `json:"source_ip,omitempty"`
    AdditionalProperties map[string]interface{}                     `json:"_"`
}

// String implements the fmt.Stringer interface for SwitchRadiusConfig,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SwitchRadiusConfig) String() string {
    return fmt.Sprintf(
    	"SwitchRadiusConfig[AcctImmediateUpdate=%v, AcctInterimInterval=%v, AcctServers=%v, AuthServerSelection=%v, AuthServers=%v, AuthServersRetries=%v, AuthServersTimeout=%v, CoaEnabled=%v, CoaPort=%v, FastDot1xTimers=%v, Network=%v, SourceIp=%v, AdditionalProperties=%v]",
    	s.AcctImmediateUpdate, s.AcctInterimInterval, s.AcctServers, s.AuthServerSelection, s.AuthServers, s.AuthServersRetries, s.AuthServersTimeout, s.CoaEnabled, s.CoaPort, s.FastDot1xTimers, s.Network, s.SourceIp, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SwitchRadiusConfig.
// It customizes the JSON marshaling process for SwitchRadiusConfig objects.
func (s SwitchRadiusConfig) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "acct_immediate_update", "acct_interim_interval", "acct_servers", "auth_server_selection", "auth_servers", "auth_servers_retries", "auth_servers_timeout", "coa_enabled", "coa_port", "fast_dot1x_timers", "network", "source_ip"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SwitchRadiusConfig object to a map representation for JSON marshaling.
func (s SwitchRadiusConfig) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.AcctImmediateUpdate != nil {
        structMap["acct_immediate_update"] = s.AcctImmediateUpdate
    }
    if s.AcctInterimInterval != nil {
        structMap["acct_interim_interval"] = s.AcctInterimInterval
    }
    if s.AcctServers != nil {
        structMap["acct_servers"] = s.AcctServers
    }
    if s.AuthServerSelection != nil {
        structMap["auth_server_selection"] = s.AuthServerSelection
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
    if s.CoaEnabled != nil {
        structMap["coa_enabled"] = s.CoaEnabled
    }
    if s.CoaPort != nil {
        structMap["coa_port"] = s.CoaPort.toMap()
    }
    if s.FastDot1xTimers != nil {
        structMap["fast_dot1x_timers"] = s.FastDot1xTimers
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
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "acct_immediate_update", "acct_interim_interval", "acct_servers", "auth_server_selection", "auth_servers", "auth_servers_retries", "auth_servers_timeout", "coa_enabled", "coa_port", "fast_dot1x_timers", "network", "source_ip")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.AcctImmediateUpdate = temp.AcctImmediateUpdate
    s.AcctInterimInterval = temp.AcctInterimInterval
    s.AcctServers = temp.AcctServers
    s.AuthServerSelection = temp.AuthServerSelection
    s.AuthServers = temp.AuthServers
    s.AuthServersRetries = temp.AuthServersRetries
    s.AuthServersTimeout = temp.AuthServersTimeout
    s.CoaEnabled = temp.CoaEnabled
    s.CoaPort = temp.CoaPort
    s.FastDot1xTimers = temp.FastDot1xTimers
    s.Network = temp.Network
    s.SourceIp = temp.SourceIp
    return nil
}

// tempSwitchRadiusConfig is a temporary struct used for validating the fields of SwitchRadiusConfig.
type tempSwitchRadiusConfig  struct {
    AcctImmediateUpdate *bool                                      `json:"acct_immediate_update,omitempty"`
    AcctInterimInterval *int                                       `json:"acct_interim_interval,omitempty"`
    AcctServers         []RadiusAcctServer                         `json:"acct_servers,omitempty"`
    AuthServerSelection *SwitchRadiusConfigAuthServerSelectionEnum `json:"auth_server_selection,omitempty"`
    AuthServers         []RadiusAuthServer                         `json:"auth_servers,omitempty"`
    AuthServersRetries  *int                                       `json:"auth_servers_retries,omitempty"`
    AuthServersTimeout  *int                                       `json:"auth_servers_timeout,omitempty"`
    CoaEnabled          *bool                                      `json:"coa_enabled,omitempty"`
    CoaPort             *CoaPort                                   `json:"coa_port,omitempty"`
    FastDot1xTimers     *bool                                      `json:"fast_dot1x_timers,omitempty"`
    Network             *string                                    `json:"network,omitempty"`
    SourceIp            *string                                    `json:"source_ip,omitempty"`
}
