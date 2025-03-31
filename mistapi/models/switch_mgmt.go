package models

import (
    "encoding/json"
    "fmt"
)

// SwitchMgmt represents a SwitchMgmt struct.
// Switch settings
type SwitchMgmt struct {
    // AP_affinity_threshold ap_affinity_threshold can be added as a field under site/setting. By default, this value is set to 12. If the field is set in both site/setting and org/setting, the value from site/setting will be used.
    ApAffinityThreshold  *int                                     `json:"ap_affinity_threshold,omitempty"`
    // Set Banners for switches. Allows markup formatting
    CliBanner            *string                                  `json:"cli_banner,omitempty"`
    // Sets timeout for switches
    CliIdleTimeout       *int                                     `json:"cli_idle_timeout,omitempty"`
    // Rollback timer for commit confirmed
    ConfigRevertTimer    *int                                     `json:"config_revert_timer,omitempty"`
    // Enable to provide the FQDN with DHCP option 81
    DhcpOptionFqdn       *bool                                    `json:"dhcp_option_fqdn,omitempty"`
    DisableOobDownAlarm  *bool                                    `json:"disable_oob_down_alarm,omitempty"`
    FipsEnabled          *bool                                    `json:"fips_enabled,omitempty"`
    // Property key is the user name. For Local user authentication
    LocalAccounts        map[string]ConfigSwitchLocalAccountsUser `json:"local_accounts,omitempty"`
    // IP Address or FQDN of the Mist Edge used to proxy the switch management traffic to the Mist Cloud
    MxedgeProxyHost      *string                                  `json:"mxedge_proxy_host,omitempty"`
    // Mist Edge port used to proxy the switch management traffic to the Mist Cloud. Value in range 1-65535
    MxedgeProxyPort      *SwitchMgmtMxedgeProxyPort               `json:"mxedge_proxy_port,omitempty"`
    // Restrict inbound-traffic to host
    // when enabled, all traffic that is not essential to our operation will be dropped
    // e.g. ntp / dns / traffic to mist will be allowed by default, if dhcpd is enabled, we'll make sure it works
    ProtectRe            *ProtectRe                               `json:"protect_re,omitempty"`
    // By default, `radius_config` will be used. if a different one has to be used set `use_different_radius
    Radius               *SwitchRadius                            `json:"radius,omitempty"`
    RootPassword         *string                                  `json:"root_password,omitempty"`
    Tacacs               *Tacacs                                  `json:"tacacs,omitempty"`
    // To use mxedge as proxy
    UseMxedgeProxy       *bool                                    `json:"use_mxedge_proxy,omitempty"`
    AdditionalProperties map[string]interface{}                   `json:"_"`
}

// String implements the fmt.Stringer interface for SwitchMgmt,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SwitchMgmt) String() string {
    return fmt.Sprintf(
    	"SwitchMgmt[ApAffinityThreshold=%v, CliBanner=%v, CliIdleTimeout=%v, ConfigRevertTimer=%v, DhcpOptionFqdn=%v, DisableOobDownAlarm=%v, FipsEnabled=%v, LocalAccounts=%v, MxedgeProxyHost=%v, MxedgeProxyPort=%v, ProtectRe=%v, Radius=%v, RootPassword=%v, Tacacs=%v, UseMxedgeProxy=%v, AdditionalProperties=%v]",
    	s.ApAffinityThreshold, s.CliBanner, s.CliIdleTimeout, s.ConfigRevertTimer, s.DhcpOptionFqdn, s.DisableOobDownAlarm, s.FipsEnabled, s.LocalAccounts, s.MxedgeProxyHost, s.MxedgeProxyPort, s.ProtectRe, s.Radius, s.RootPassword, s.Tacacs, s.UseMxedgeProxy, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SwitchMgmt.
// It customizes the JSON marshaling process for SwitchMgmt objects.
func (s SwitchMgmt) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "ap_affinity_threshold", "cli_banner", "cli_idle_timeout", "config_revert_timer", "dhcp_option_fqdn", "disable_oob_down_alarm", "fips_enabled", "local_accounts", "mxedge_proxy_host", "mxedge_proxy_port", "protect_re", "radius", "root_password", "tacacs", "use_mxedge_proxy"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SwitchMgmt object to a map representation for JSON marshaling.
func (s SwitchMgmt) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
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
    if s.FipsEnabled != nil {
        structMap["fips_enabled"] = s.FipsEnabled
    }
    if s.LocalAccounts != nil {
        structMap["local_accounts"] = s.LocalAccounts
    }
    if s.MxedgeProxyHost != nil {
        structMap["mxedge_proxy_host"] = s.MxedgeProxyHost
    }
    if s.MxedgeProxyPort != nil {
        structMap["mxedge_proxy_port"] = s.MxedgeProxyPort.toMap()
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
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "ap_affinity_threshold", "cli_banner", "cli_idle_timeout", "config_revert_timer", "dhcp_option_fqdn", "disable_oob_down_alarm", "fips_enabled", "local_accounts", "mxedge_proxy_host", "mxedge_proxy_port", "protect_re", "radius", "root_password", "tacacs", "use_mxedge_proxy")
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
    s.FipsEnabled = temp.FipsEnabled
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
    FipsEnabled         *bool                                    `json:"fips_enabled,omitempty"`
    LocalAccounts       map[string]ConfigSwitchLocalAccountsUser `json:"local_accounts,omitempty"`
    MxedgeProxyHost     *string                                  `json:"mxedge_proxy_host,omitempty"`
    MxedgeProxyPort     *SwitchMgmtMxedgeProxyPort               `json:"mxedge_proxy_port,omitempty"`
    ProtectRe           *ProtectRe                               `json:"protect_re,omitempty"`
    Radius              *SwitchRadius                            `json:"radius,omitempty"`
    RootPassword        *string                                  `json:"root_password,omitempty"`
    Tacacs              *Tacacs                                  `json:"tacacs,omitempty"`
    UseMxedgeProxy      *bool                                    `json:"use_mxedge_proxy,omitempty"`
}
