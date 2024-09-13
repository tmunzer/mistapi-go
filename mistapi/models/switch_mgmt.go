package models

import (
    "encoding/json"
)

// SwitchMgmt represents a SwitchMgmt struct.
// Switch settings
type SwitchMgmt struct {
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
    DisableOobDownAlarm  *bool                                    `json:"disable_oob_down_alarm,omitempty"`
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
    if s.ApAffinityThreshold != nil {
        structMap["ap_affinity_threshold"] = s.ApAffinityThreshold
    }
    if s.CliBanner != nil {
        structMap["cli_banner"] = s.CliBanner
    }
    if s.CliIdleTimeout != nil {
        structMap["cli_idle_timeout"] = s.CliIdleTimeout
    }
    if s.ConfigRevertTimer != nil {
        structMap["config_revert_timer"] = s.ConfigRevertTimer
    }
    if s.DhcpOptionFqdn != nil {
        structMap["dhcp_option_fqdn"] = s.DhcpOptionFqdn
    }
    if s.DisableOobDownAlarm != nil {
        structMap["disable_oob_down_alarm"] = s.DisableOobDownAlarm
    }
    if s.LocalAccounts != nil {
        structMap["local_accounts"] = s.LocalAccounts
    }
    if s.MxedgeProxyHost != nil {
        structMap["mxedge_proxy_host"] = s.MxedgeProxyHost
    }
    if s.MxedgeProxyPort != nil {
        structMap["mxedge_proxy_port"] = s.MxedgeProxyPort
    }
    if s.ProtectRe != nil {
        structMap["protect_re"] = s.ProtectRe.toMap()
    }
    if s.Radius != nil {
        structMap["radius"] = s.Radius.toMap()
    }
    if s.RootPassword != nil {
        structMap["root_password"] = s.RootPassword
    }
    if s.Tacacs != nil {
        structMap["tacacs"] = s.Tacacs.toMap()
    }
    if s.UseMxedgeProxy != nil {
        structMap["use_mxedge_proxy"] = s.UseMxedgeProxy
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SwitchMgmt.
// It customizes the JSON unmarshaling process for SwitchMgmt objects.
func (s *SwitchMgmt) UnmarshalJSON(input []byte) error {
    var temp tempSwitchMgmt
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "ap_affinity_threshold", "cli_banner", "cli_idle_timeout", "config_revert_timer", "dhcp_option_fqdn", "disable_oob_down_alarm", "local_accounts", "mxedge_proxy_host", "mxedge_proxy_port", "protect_re", "radius", "root_password", "tacacs", "use_mxedge_proxy")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.ApAffinityThreshold = temp.ApAffinityThreshold
    s.CliBanner = temp.CliBanner
    s.CliIdleTimeout = temp.CliIdleTimeout
    s.ConfigRevertTimer = temp.ConfigRevertTimer
    s.DhcpOptionFqdn = temp.DhcpOptionFqdn
    s.DisableOobDownAlarm = temp.DisableOobDownAlarm
    s.LocalAccounts = temp.LocalAccounts
    s.MxedgeProxyHost = temp.MxedgeProxyHost
    s.MxedgeProxyPort = temp.MxedgeProxyPort
    s.ProtectRe = temp.ProtectRe
    s.Radius = temp.Radius
    s.RootPassword = temp.RootPassword
    s.Tacacs = temp.Tacacs
    s.UseMxedgeProxy = temp.UseMxedgeProxy
    return nil
}

// tempSwitchMgmt is a temporary struct used for validating the fields of SwitchMgmt.
type tempSwitchMgmt  struct {
    ApAffinityThreshold *int                                     `json:"ap_affinity_threshold,omitempty"`
    CliBanner           *string                                  `json:"cli_banner,omitempty"`
    CliIdleTimeout      *int                                     `json:"cli_idle_timeout,omitempty"`
    ConfigRevertTimer   *int                                     `json:"config_revert_timer,omitempty"`
    DhcpOptionFqdn      *bool                                    `json:"dhcp_option_fqdn,omitempty"`
    DisableOobDownAlarm *bool                                    `json:"disable_oob_down_alarm,omitempty"`
    LocalAccounts       map[string]ConfigSwitchLocalAccountsUser `json:"local_accounts,omitempty"`
    MxedgeProxyHost     *string                                  `json:"mxedge_proxy_host,omitempty"`
    MxedgeProxyPort     *int                                     `json:"mxedge_proxy_port,omitempty"`
    ProtectRe           *ProtectRe                               `json:"protect_re,omitempty"`
    Radius              *ConfigSwitchRadius                      `json:"radius,omitempty"`
    RootPassword        *string                                  `json:"root_password,omitempty"`
    Tacacs              *Tacacs                                  `json:"tacacs,omitempty"`
    UseMxedgeProxy      *bool                                    `json:"use_mxedge_proxy,omitempty"`
}
