package models

import (
    "encoding/json"
    "fmt"
)

// SiteWifi represents a SiteWifi struct.
// Wi-Fi site settings
type SiteWifi struct {
    CiscoEnabled                      *bool                          `json:"cisco_enabled,omitempty"`
    // Whether to disable 11k
    Disable11k                        *bool                          `json:"disable_11k,omitempty"`
    DisableRadiosWhenPowerConstrained *bool                          `json:"disable_radios_when_power_constrained,omitempty"`
    // When proxy_arp is enabled, check for arp spoofing.
    EnableArpSpoofCheck               *bool                          `json:"enable_arp_spoof_check,omitempty"`
    EnableSharedRadioScanning         *bool                          `json:"enable_shared_radio_scanning,omitempty"`
    // Enable WIFI feature (using SUB-MAN license)
    Enabled                           *bool                          `json:"enabled,omitempty"`
    // Whether to locate connected clients
    LocateConnected                   *bool                          `json:"locate_connected,omitempty"`
    // Whether to locate unconnected clients
    LocateUnconnected                 *bool                          `json:"locate_unconnected,omitempty"`
    // Whether to allow Mesh to use DFS channels. For DFS channels, Remote Mesh AP would have to do CAC when scanning for new Base AP, which is slow and will distrupt the connection. If roaming is desired, keep it disabled.
    MeshAllowDfs                      *bool                          `json:"mesh_allow_dfs,omitempty"`
    // Used to enable/disable CRM
    MeshEnableCrm                     *bool                          `json:"mesh_enable_crm,omitempty"`
    // Whether to enable Mesh feature for the site
    MeshEnabled                       *bool                          `json:"mesh_enabled,omitempty"`
    // Optional passphrase of mesh networking, default is generated randomly
    MeshPsk                           Optional[string]               `json:"mesh_psk"`
    // Optional ssid of mesh networking, default is based on site_id
    MeshSsid                          Optional[string]               `json:"mesh_ssid"`
    // enum: `default`, `disabled`, `enabled`
    ProxyArp                          Optional[SiteWifiProxyArpEnum] `json:"proxy_arp"`
    AdditionalProperties              map[string]interface{}         `json:"_"`
}

// String implements the fmt.Stringer interface for SiteWifi,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SiteWifi) String() string {
    return fmt.Sprintf(
    	"SiteWifi[CiscoEnabled=%v, Disable11k=%v, DisableRadiosWhenPowerConstrained=%v, EnableArpSpoofCheck=%v, EnableSharedRadioScanning=%v, Enabled=%v, LocateConnected=%v, LocateUnconnected=%v, MeshAllowDfs=%v, MeshEnableCrm=%v, MeshEnabled=%v, MeshPsk=%v, MeshSsid=%v, ProxyArp=%v, AdditionalProperties=%v]",
    	s.CiscoEnabled, s.Disable11k, s.DisableRadiosWhenPowerConstrained, s.EnableArpSpoofCheck, s.EnableSharedRadioScanning, s.Enabled, s.LocateConnected, s.LocateUnconnected, s.MeshAllowDfs, s.MeshEnableCrm, s.MeshEnabled, s.MeshPsk, s.MeshSsid, s.ProxyArp, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SiteWifi.
// It customizes the JSON marshaling process for SiteWifi objects.
func (s SiteWifi) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "cisco_enabled", "disable_11k", "disable_radios_when_power_constrained", "enable_arp_spoof_check", "enable_shared_radio_scanning", "enabled", "locate_connected", "locate_unconnected", "mesh_allow_dfs", "mesh_enable_crm", "mesh_enabled", "mesh_psk", "mesh_ssid", "proxy_arp"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SiteWifi object to a map representation for JSON marshaling.
func (s SiteWifi) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.CiscoEnabled != nil {
        structMap["cisco_enabled"] = s.CiscoEnabled
    }
    if s.Disable11k != nil {
        structMap["disable_11k"] = s.Disable11k
    }
    if s.DisableRadiosWhenPowerConstrained != nil {
        structMap["disable_radios_when_power_constrained"] = s.DisableRadiosWhenPowerConstrained
    }
    if s.EnableArpSpoofCheck != nil {
        structMap["enable_arp_spoof_check"] = s.EnableArpSpoofCheck
    }
    if s.EnableSharedRadioScanning != nil {
        structMap["enable_shared_radio_scanning"] = s.EnableSharedRadioScanning
    }
    if s.Enabled != nil {
        structMap["enabled"] = s.Enabled
    }
    if s.LocateConnected != nil {
        structMap["locate_connected"] = s.LocateConnected
    }
    if s.LocateUnconnected != nil {
        structMap["locate_unconnected"] = s.LocateUnconnected
    }
    if s.MeshAllowDfs != nil {
        structMap["mesh_allow_dfs"] = s.MeshAllowDfs
    }
    if s.MeshEnableCrm != nil {
        structMap["mesh_enable_crm"] = s.MeshEnableCrm
    }
    if s.MeshEnabled != nil {
        structMap["mesh_enabled"] = s.MeshEnabled
    }
    if s.MeshPsk.IsValueSet() {
        if s.MeshPsk.Value() != nil {
            structMap["mesh_psk"] = s.MeshPsk.Value()
        } else {
            structMap["mesh_psk"] = nil
        }
    }
    if s.MeshSsid.IsValueSet() {
        if s.MeshSsid.Value() != nil {
            structMap["mesh_ssid"] = s.MeshSsid.Value()
        } else {
            structMap["mesh_ssid"] = nil
        }
    }
    if s.ProxyArp.IsValueSet() {
        if s.ProxyArp.Value() != nil {
            structMap["proxy_arp"] = s.ProxyArp.Value()
        } else {
            structMap["proxy_arp"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SiteWifi.
// It customizes the JSON unmarshaling process for SiteWifi objects.
func (s *SiteWifi) UnmarshalJSON(input []byte) error {
    var temp tempSiteWifi
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "cisco_enabled", "disable_11k", "disable_radios_when_power_constrained", "enable_arp_spoof_check", "enable_shared_radio_scanning", "enabled", "locate_connected", "locate_unconnected", "mesh_allow_dfs", "mesh_enable_crm", "mesh_enabled", "mesh_psk", "mesh_ssid", "proxy_arp")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.CiscoEnabled = temp.CiscoEnabled
    s.Disable11k = temp.Disable11k
    s.DisableRadiosWhenPowerConstrained = temp.DisableRadiosWhenPowerConstrained
    s.EnableArpSpoofCheck = temp.EnableArpSpoofCheck
    s.EnableSharedRadioScanning = temp.EnableSharedRadioScanning
    s.Enabled = temp.Enabled
    s.LocateConnected = temp.LocateConnected
    s.LocateUnconnected = temp.LocateUnconnected
    s.MeshAllowDfs = temp.MeshAllowDfs
    s.MeshEnableCrm = temp.MeshEnableCrm
    s.MeshEnabled = temp.MeshEnabled
    s.MeshPsk = temp.MeshPsk
    s.MeshSsid = temp.MeshSsid
    s.ProxyArp = temp.ProxyArp
    return nil
}

// tempSiteWifi is a temporary struct used for validating the fields of SiteWifi.
type tempSiteWifi  struct {
    CiscoEnabled                      *bool                          `json:"cisco_enabled,omitempty"`
    Disable11k                        *bool                          `json:"disable_11k,omitempty"`
    DisableRadiosWhenPowerConstrained *bool                          `json:"disable_radios_when_power_constrained,omitempty"`
    EnableArpSpoofCheck               *bool                          `json:"enable_arp_spoof_check,omitempty"`
    EnableSharedRadioScanning         *bool                          `json:"enable_shared_radio_scanning,omitempty"`
    Enabled                           *bool                          `json:"enabled,omitempty"`
    LocateConnected                   *bool                          `json:"locate_connected,omitempty"`
    LocateUnconnected                 *bool                          `json:"locate_unconnected,omitempty"`
    MeshAllowDfs                      *bool                          `json:"mesh_allow_dfs,omitempty"`
    MeshEnableCrm                     *bool                          `json:"mesh_enable_crm,omitempty"`
    MeshEnabled                       *bool                          `json:"mesh_enabled,omitempty"`
    MeshPsk                           Optional[string]               `json:"mesh_psk"`
    MeshSsid                          Optional[string]               `json:"mesh_ssid"`
    ProxyArp                          Optional[SiteWifiProxyArpEnum] `json:"proxy_arp"`
}
