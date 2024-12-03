package models

import (
    "encoding/json"
)

// ApTemplateWifi represents a ApTemplateWifi struct.
type ApTemplateWifi struct {
    CiscoEnabled                      *bool                  `json:"cisco_enabled,omitempty"`
    Disable11k                        *bool                  `json:"disable_11k,omitempty"`
    DisableRadiosWhenPowerConstrained *bool                  `json:"disable_radios_when_power_constrained,omitempty"`
    EnableArpSpoof                    *bool                  `json:"enable_arp_spoof,omitempty"`
    EnableSharedRadioScanning         *bool                  `json:"enable_shared_radio_scanning,omitempty"`
    Enabled                           *bool                  `json:"enabled,omitempty"`
    LocateConnected                   *bool                  `json:"locate_connected,omitempty"`
    LocateUnconnected                 *bool                  `json:"locate_unconnected,omitempty"`
    MeshAllowDfs                      *bool                  `json:"mesh_allow_dfs,omitempty"`
    MeshEnableCrm                     *bool                  `json:"mesh_enable_crm,omitempty"`
    MeshEnabled                       *bool                  `json:"mesh_enabled,omitempty"`
    ProxyArp                          *bool                  `json:"proxy_arp,omitempty"`
    AdditionalProperties              map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ApTemplateWifi.
// It customizes the JSON marshaling process for ApTemplateWifi objects.
func (a ApTemplateWifi) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(a.AdditionalProperties,
        "cisco_enabled", "disable_11k", "disable_radios_when_power_constrained", "enable_arp_spoof", "enable_shared_radio_scanning", "enabled", "locate_connected", "locate_unconnected", "mesh_allow_dfs", "mesh_enable_crm", "mesh_enabled", "proxy_arp"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(a.toMap())
}

// toMap converts the ApTemplateWifi object to a map representation for JSON marshaling.
func (a ApTemplateWifi) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, a.AdditionalProperties)
    if a.CiscoEnabled != nil {
        structMap["cisco_enabled"] = a.CiscoEnabled
    }
    if a.Disable11k != nil {
        structMap["disable_11k"] = a.Disable11k
    }
    if a.DisableRadiosWhenPowerConstrained != nil {
        structMap["disable_radios_when_power_constrained"] = a.DisableRadiosWhenPowerConstrained
    }
    if a.EnableArpSpoof != nil {
        structMap["enable_arp_spoof"] = a.EnableArpSpoof
    }
    if a.EnableSharedRadioScanning != nil {
        structMap["enable_shared_radio_scanning"] = a.EnableSharedRadioScanning
    }
    if a.Enabled != nil {
        structMap["enabled"] = a.Enabled
    }
    if a.LocateConnected != nil {
        structMap["locate_connected"] = a.LocateConnected
    }
    if a.LocateUnconnected != nil {
        structMap["locate_unconnected"] = a.LocateUnconnected
    }
    if a.MeshAllowDfs != nil {
        structMap["mesh_allow_dfs"] = a.MeshAllowDfs
    }
    if a.MeshEnableCrm != nil {
        structMap["mesh_enable_crm"] = a.MeshEnableCrm
    }
    if a.MeshEnabled != nil {
        structMap["mesh_enabled"] = a.MeshEnabled
    }
    if a.ProxyArp != nil {
        structMap["proxy_arp"] = a.ProxyArp
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ApTemplateWifi.
// It customizes the JSON unmarshaling process for ApTemplateWifi objects.
func (a *ApTemplateWifi) UnmarshalJSON(input []byte) error {
    var temp tempApTemplateWifi
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "cisco_enabled", "disable_11k", "disable_radios_when_power_constrained", "enable_arp_spoof", "enable_shared_radio_scanning", "enabled", "locate_connected", "locate_unconnected", "mesh_allow_dfs", "mesh_enable_crm", "mesh_enabled", "proxy_arp")
    if err != nil {
    	return err
    }
    a.AdditionalProperties = additionalProperties
    
    a.CiscoEnabled = temp.CiscoEnabled
    a.Disable11k = temp.Disable11k
    a.DisableRadiosWhenPowerConstrained = temp.DisableRadiosWhenPowerConstrained
    a.EnableArpSpoof = temp.EnableArpSpoof
    a.EnableSharedRadioScanning = temp.EnableSharedRadioScanning
    a.Enabled = temp.Enabled
    a.LocateConnected = temp.LocateConnected
    a.LocateUnconnected = temp.LocateUnconnected
    a.MeshAllowDfs = temp.MeshAllowDfs
    a.MeshEnableCrm = temp.MeshEnableCrm
    a.MeshEnabled = temp.MeshEnabled
    a.ProxyArp = temp.ProxyArp
    return nil
}

// tempApTemplateWifi is a temporary struct used for validating the fields of ApTemplateWifi.
type tempApTemplateWifi  struct {
    CiscoEnabled                      *bool `json:"cisco_enabled,omitempty"`
    Disable11k                        *bool `json:"disable_11k,omitempty"`
    DisableRadiosWhenPowerConstrained *bool `json:"disable_radios_when_power_constrained,omitempty"`
    EnableArpSpoof                    *bool `json:"enable_arp_spoof,omitempty"`
    EnableSharedRadioScanning         *bool `json:"enable_shared_radio_scanning,omitempty"`
    Enabled                           *bool `json:"enabled,omitempty"`
    LocateConnected                   *bool `json:"locate_connected,omitempty"`
    LocateUnconnected                 *bool `json:"locate_unconnected,omitempty"`
    MeshAllowDfs                      *bool `json:"mesh_allow_dfs,omitempty"`
    MeshEnableCrm                     *bool `json:"mesh_enable_crm,omitempty"`
    MeshEnabled                       *bool `json:"mesh_enabled,omitempty"`
    ProxyArp                          *bool `json:"proxy_arp,omitempty"`
}
