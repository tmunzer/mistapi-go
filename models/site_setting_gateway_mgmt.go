package models

import (
    "encoding/json"
)

// SiteSettingGatewayMgmt represents a SiteSettingGatewayMgmt struct.
// Gateway Site settings
type SiteSettingGatewayMgmt struct {
    // for SSR only, as direct root access is not allowed
    AdminSshkeys               []string                                   `json:"admin_sshkeys,omitempty"`
    AppProbing                 *AppProbing                                `json:"app_probing,omitempty"`
    // consumes uplink bandwidth, requires WA license
    AppUsage                   *bool                                      `json:"app_usage,omitempty"`
    AutoSignatureUpdate        *SiteSettingGatewayMgmtAutoSignatureUpdate `json:"auto_signature_update,omitempty"`
    // he rollback timer for commit confirmed
    ConfigRevertTimer          *float64                                   `json:"config_revert_timer,omitempty"`
    ProbeHosts                 []string                                   `json:"probe_hosts,omitempty"`
    // for SRX only
    RootPassword               *string                                    `json:"root_password,omitempty"`
    SecurityLogSourceAddress   *string                                    `json:"security_log_source_address,omitempty"`
    SecurityLogSourceInterface *string                                    `json:"security_log_source_interface,omitempty"`
    AdditionalProperties       map[string]any                             `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SiteSettingGatewayMgmt.
// It customizes the JSON marshaling process for SiteSettingGatewayMgmt objects.
func (s SiteSettingGatewayMgmt) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SiteSettingGatewayMgmt object to a map representation for JSON marshaling.
func (s SiteSettingGatewayMgmt) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.AdminSshkeys != nil {
        structMap["admin_sshkeys"] = s.AdminSshkeys
    }
    if s.AppProbing != nil {
        structMap["app_probing"] = s.AppProbing.toMap()
    }
    if s.AppUsage != nil {
        structMap["app_usage"] = s.AppUsage
    }
    if s.AutoSignatureUpdate != nil {
        structMap["auto_signature_update"] = s.AutoSignatureUpdate.toMap()
    }
    if s.ConfigRevertTimer != nil {
        structMap["config_revert_timer"] = s.ConfigRevertTimer
    }
    if s.ProbeHosts != nil {
        structMap["probe_hosts"] = s.ProbeHosts
    }
    if s.RootPassword != nil {
        structMap["root_password"] = s.RootPassword
    }
    if s.SecurityLogSourceAddress != nil {
        structMap["security_log_source_address"] = s.SecurityLogSourceAddress
    }
    if s.SecurityLogSourceInterface != nil {
        structMap["security_log_source_interface"] = s.SecurityLogSourceInterface
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SiteSettingGatewayMgmt.
// It customizes the JSON unmarshaling process for SiteSettingGatewayMgmt objects.
func (s *SiteSettingGatewayMgmt) UnmarshalJSON(input []byte) error {
    var temp siteSettingGatewayMgmt
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "admin_sshkeys", "app_probing", "app_usage", "auto_signature_update", "config_revert_timer", "probe_hosts", "root_password", "security_log_source_address", "security_log_source_interface")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.AdminSshkeys = temp.AdminSshkeys
    s.AppProbing = temp.AppProbing
    s.AppUsage = temp.AppUsage
    s.AutoSignatureUpdate = temp.AutoSignatureUpdate
    s.ConfigRevertTimer = temp.ConfigRevertTimer
    s.ProbeHosts = temp.ProbeHosts
    s.RootPassword = temp.RootPassword
    s.SecurityLogSourceAddress = temp.SecurityLogSourceAddress
    s.SecurityLogSourceInterface = temp.SecurityLogSourceInterface
    return nil
}

// siteSettingGatewayMgmt is a temporary struct used for validating the fields of SiteSettingGatewayMgmt.
type siteSettingGatewayMgmt  struct {
    AdminSshkeys               []string                                   `json:"admin_sshkeys,omitempty"`
    AppProbing                 *AppProbing                                `json:"app_probing,omitempty"`
    AppUsage                   *bool                                      `json:"app_usage,omitempty"`
    AutoSignatureUpdate        *SiteSettingGatewayMgmtAutoSignatureUpdate `json:"auto_signature_update,omitempty"`
    ConfigRevertTimer          *float64                                   `json:"config_revert_timer,omitempty"`
    ProbeHosts                 []string                                   `json:"probe_hosts,omitempty"`
    RootPassword               *string                                    `json:"root_password,omitempty"`
    SecurityLogSourceAddress   *string                                    `json:"security_log_source_address,omitempty"`
    SecurityLogSourceInterface *string                                    `json:"security_log_source_interface,omitempty"`
}
