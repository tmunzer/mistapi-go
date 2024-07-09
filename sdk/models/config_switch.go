package models

import (
    "encoding/json"
)

// ConfigSwitch represents a ConfigSwitch struct.
// Switch settings
type ConfigSwitch struct {
    // ap_affinity_threshold ap_affinity_threshold can be added as a field under site/setting. By default this value is set to 12. If the field is set in both site/setting and org/setting, the value from site/setting will be used.
    ApAffinityThreshold  *int                                     `json:"ap_affinity_threshold,omitempty"`
    // Set Banners for switches. Allows markup formatting
    CliBanner            *string                                  `json:"cli_banner,omitempty"`
    // Sets timeout for switches
    CliIdleTimeout       *int                                     `json:"cli_idle_timeout,omitempty"`
    // the rollback timer for commit confirmed
    ConfigRevertTimer    *int                                     `json:"config_revert_timer,omitempty"`
    // Enable to provide the FQDN with DHCP option 81
    DhcpOptionFqdn       *bool                                    `json:"dhcp_option_fqdn,omitempty"`
    // Property key is the user name. For Local user authentication
    LocalAccounts        map[string]ConfigSwitchLocalAccountsUser `json:"local_accounts,omitempty"`
    MxedgeProxyHost      *string                                  `json:"mxedge_proxy_host,omitempty"`
    MxedgeProxyPort      *int                                     `json:"mxedge_proxy_port,omitempty"`
    // restrict inbound-traffic to host
    // when enabled, all traffic that is not essential to our operation will be dropped 
    // e.g. ntp / dns / traffic to mist will be allowed by default, if dhcpd is enabled, we'll make sure it works
    ProtectRe            *ProtectRe                               `json:"protect_re,omitempty"`
    // by default, `radius_config` will be used. if a different one has to be used set `use_different_radius
    Radius               *ConfigSwitchRadius                      `json:"radius,omitempty"`
    RootPassword         *string                                  `json:"root_password,omitempty"`
    Tacacs               *Tacacs                                  `json:"tacacs,omitempty"`
    // to use mxedge as proxy
    UseMxedgeProxy       *bool                                    `json:"use_mxedge_proxy,omitempty"`
    AdditionalProperties map[string]any                           `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ConfigSwitch.
// It customizes the JSON marshaling process for ConfigSwitch objects.
func (c ConfigSwitch) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the ConfigSwitch object to a map representation for JSON marshaling.
func (c ConfigSwitch) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    if c.ApAffinityThreshold != nil {
        structMap["ap_affinity_threshold"] = c.ApAffinityThreshold
    }
    if c.CliBanner != nil {
        structMap["cli_banner"] = c.CliBanner
    }
    if c.CliIdleTimeout != nil {
        structMap["cli_idle_timeout"] = c.CliIdleTimeout
    }
    if c.ConfigRevertTimer != nil {
        structMap["config_revert_timer"] = c.ConfigRevertTimer
    }
    if c.DhcpOptionFqdn != nil {
        structMap["dhcp_option_fqdn"] = c.DhcpOptionFqdn
    }
    if c.LocalAccounts != nil {
        structMap["local_accounts"] = c.LocalAccounts
    }
    if c.MxedgeProxyHost != nil {
        structMap["mxedge_proxy_host"] = c.MxedgeProxyHost
    }
    if c.MxedgeProxyPort != nil {
        structMap["mxedge_proxy_port"] = c.MxedgeProxyPort
    }
    if c.ProtectRe != nil {
        structMap["protect_re"] = c.ProtectRe.toMap()
    }
    if c.Radius != nil {
        structMap["radius"] = c.Radius.toMap()
    }
    if c.RootPassword != nil {
        structMap["root_password"] = c.RootPassword
    }
    if c.Tacacs != nil {
        structMap["tacacs"] = c.Tacacs.toMap()
    }
    if c.UseMxedgeProxy != nil {
        structMap["use_mxedge_proxy"] = c.UseMxedgeProxy
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ConfigSwitch.
// It customizes the JSON unmarshaling process for ConfigSwitch objects.
func (c *ConfigSwitch) UnmarshalJSON(input []byte) error {
    var temp configSwitch
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "ap_affinity_threshold", "cli_banner", "cli_idle_timeout", "config_revert_timer", "dhcp_option_fqdn", "local_accounts", "mxedge_proxy_host", "mxedge_proxy_port", "protect_re", "radius", "root_password", "tacacs", "use_mxedge_proxy")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.ApAffinityThreshold = temp.ApAffinityThreshold
    c.CliBanner = temp.CliBanner
    c.CliIdleTimeout = temp.CliIdleTimeout
    c.ConfigRevertTimer = temp.ConfigRevertTimer
    c.DhcpOptionFqdn = temp.DhcpOptionFqdn
    c.LocalAccounts = temp.LocalAccounts
    c.MxedgeProxyHost = temp.MxedgeProxyHost
    c.MxedgeProxyPort = temp.MxedgeProxyPort
    c.ProtectRe = temp.ProtectRe
    c.Radius = temp.Radius
    c.RootPassword = temp.RootPassword
    c.Tacacs = temp.Tacacs
    c.UseMxedgeProxy = temp.UseMxedgeProxy
    return nil
}

// configSwitch is a temporary struct used for validating the fields of ConfigSwitch.
type configSwitch  struct {
    ApAffinityThreshold *int                                     `json:"ap_affinity_threshold,omitempty"`
    CliBanner           *string                                  `json:"cli_banner,omitempty"`
    CliIdleTimeout      *int                                     `json:"cli_idle_timeout,omitempty"`
    ConfigRevertTimer   *int                                     `json:"config_revert_timer,omitempty"`
    DhcpOptionFqdn      *bool                                    `json:"dhcp_option_fqdn,omitempty"`
    LocalAccounts       map[string]ConfigSwitchLocalAccountsUser `json:"local_accounts,omitempty"`
    MxedgeProxyHost     *string                                  `json:"mxedge_proxy_host,omitempty"`
    MxedgeProxyPort     *int                                     `json:"mxedge_proxy_port,omitempty"`
    ProtectRe           *ProtectRe                               `json:"protect_re,omitempty"`
    Radius              *ConfigSwitchRadius                      `json:"radius,omitempty"`
    RootPassword        *string                                  `json:"root_password,omitempty"`
    Tacacs              *Tacacs                                  `json:"tacacs,omitempty"`
    UseMxedgeProxy      *bool                                    `json:"use_mxedge_proxy,omitempty"`
}
