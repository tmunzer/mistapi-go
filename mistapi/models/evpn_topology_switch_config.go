package models

import (
    "encoding/json"
    "fmt"
)

// EvpnTopologySwitchConfig represents a EvpnTopologySwitchConfig struct.
type EvpnTopologySwitchConfig struct {
    DhcpdConfig          *EvpnTopologySwitchConfigDhcpdConfig `json:"dhcpd_config,omitempty"`
    // Property key is network name
    Networks             map[string]SwitchNetwork             `json:"networks,omitempty"`
    // Additional IP Addresses configured on the switch. Property key is the port network name
    OtherIpConfigs       map[string]JunosOtherIpConfig        `json:"other_ip_configs,omitempty"`
    // Property key is the port name or range (e.g. "ge-0/0/0-10")
    PortConfig           map[string]JunosPortConfig           `json:"port_config,omitempty"`
    // Property key is the port usage name. Defines the profiles of port configuration configured on the switch
    PortUsages           map[string]SwitchPortUsage           `json:"port_usages,omitempty"`
    // Used for OSPF / BGP / EVPN
    RouterId             *string                              `json:"router_id,omitempty"`
    VrfConfig            *EvpnTopologySwitchConfigVrfConfig   `json:"vrf_config,omitempty"`
    AdditionalProperties map[string]interface{}               `json:"_"`
}

// String implements the fmt.Stringer interface for EvpnTopologySwitchConfig,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (e EvpnTopologySwitchConfig) String() string {
    return fmt.Sprintf(
    	"EvpnTopologySwitchConfig[DhcpdConfig=%v, Networks=%v, OtherIpConfigs=%v, PortConfig=%v, PortUsages=%v, RouterId=%v, VrfConfig=%v, AdditionalProperties=%v]",
    	e.DhcpdConfig, e.Networks, e.OtherIpConfigs, e.PortConfig, e.PortUsages, e.RouterId, e.VrfConfig, e.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for EvpnTopologySwitchConfig.
// It customizes the JSON marshaling process for EvpnTopologySwitchConfig objects.
func (e EvpnTopologySwitchConfig) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(e.AdditionalProperties,
        "dhcpd_config", "networks", "other_ip_configs", "port_config", "port_usages", "router_id", "vrf_config"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(e.toMap())
}

// toMap converts the EvpnTopologySwitchConfig object to a map representation for JSON marshaling.
func (e EvpnTopologySwitchConfig) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, e.AdditionalProperties)
    if e.DhcpdConfig != nil {
        structMap["dhcpd_config"] = e.DhcpdConfig.toMap()
    }
    if e.Networks != nil {
        structMap["networks"] = e.Networks
    }
    if e.OtherIpConfigs != nil {
        structMap["other_ip_configs"] = e.OtherIpConfigs
    }
    if e.PortConfig != nil {
        structMap["port_config"] = e.PortConfig
    }
    if e.PortUsages != nil {
        structMap["port_usages"] = e.PortUsages
    }
    if e.RouterId != nil {
        structMap["router_id"] = e.RouterId
    }
    if e.VrfConfig != nil {
        structMap["vrf_config"] = e.VrfConfig.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for EvpnTopologySwitchConfig.
// It customizes the JSON unmarshaling process for EvpnTopologySwitchConfig objects.
func (e *EvpnTopologySwitchConfig) UnmarshalJSON(input []byte) error {
    var temp tempEvpnTopologySwitchConfig
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "dhcpd_config", "networks", "other_ip_configs", "port_config", "port_usages", "router_id", "vrf_config")
    if err != nil {
    	return err
    }
    e.AdditionalProperties = additionalProperties
    
    e.DhcpdConfig = temp.DhcpdConfig
    e.Networks = temp.Networks
    e.OtherIpConfigs = temp.OtherIpConfigs
    e.PortConfig = temp.PortConfig
    e.PortUsages = temp.PortUsages
    e.RouterId = temp.RouterId
    e.VrfConfig = temp.VrfConfig
    return nil
}

// tempEvpnTopologySwitchConfig is a temporary struct used for validating the fields of EvpnTopologySwitchConfig.
type tempEvpnTopologySwitchConfig  struct {
    DhcpdConfig    *EvpnTopologySwitchConfigDhcpdConfig `json:"dhcpd_config,omitempty"`
    Networks       map[string]SwitchNetwork             `json:"networks,omitempty"`
    OtherIpConfigs map[string]JunosOtherIpConfig        `json:"other_ip_configs,omitempty"`
    PortConfig     map[string]JunosPortConfig           `json:"port_config,omitempty"`
    PortUsages     map[string]SwitchPortUsage           `json:"port_usages,omitempty"`
    RouterId       *string                              `json:"router_id,omitempty"`
    VrfConfig      *EvpnTopologySwitchConfigVrfConfig   `json:"vrf_config,omitempty"`
}
