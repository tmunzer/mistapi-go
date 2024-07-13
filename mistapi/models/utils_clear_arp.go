package models

import (
    "encoding/json"
)

// UtilsClearArp represents a UtilsClearArp struct.
type UtilsClearArp struct {
    // The IP address for which to clear an ARP entry. port_id must be specified. Both vlan and ip cannot be specified.
    Ip                   *string            `json:"ip,omitempty"`
    // only for HA
    Node                 *HaClusterNodeEnum `json:"node,omitempty"`
    // The device interface on which to clear the ARP cache.
    PortId               *string            `json:"port_id,omitempty"`
    // The VLAN on which to clear the ARP cache. port_id must be specified. Both vlan and ip cannot be specified.
    Vlan                 *int               `json:"vlan,omitempty"`
    // The vrf for which to clear an ARP entry. applicable for switch.
    Vrf                  *string            `json:"vrf,omitempty"`
    AdditionalProperties map[string]any     `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for UtilsClearArp.
// It customizes the JSON marshaling process for UtilsClearArp objects.
func (u UtilsClearArp) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(u.toMap())
}

// toMap converts the UtilsClearArp object to a map representation for JSON marshaling.
func (u UtilsClearArp) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, u.AdditionalProperties)
    if u.Ip != nil {
        structMap["ip"] = u.Ip
    }
    if u.Node != nil {
        structMap["node"] = u.Node
    }
    if u.PortId != nil {
        structMap["port_id"] = u.PortId
    }
    if u.Vlan != nil {
        structMap["vlan"] = u.Vlan
    }
    if u.Vrf != nil {
        structMap["vrf"] = u.Vrf
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UtilsClearArp.
// It customizes the JSON unmarshaling process for UtilsClearArp objects.
func (u *UtilsClearArp) UnmarshalJSON(input []byte) error {
    var temp utilsClearArp
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "ip", "node", "port_id", "vlan", "vrf")
    if err != nil {
    	return err
    }
    
    u.AdditionalProperties = additionalProperties
    u.Ip = temp.Ip
    u.Node = temp.Node
    u.PortId = temp.PortId
    u.Vlan = temp.Vlan
    u.Vrf = temp.Vrf
    return nil
}

// utilsClearArp is a temporary struct used for validating the fields of UtilsClearArp.
type utilsClearArp  struct {
    Ip     *string            `json:"ip,omitempty"`
    Node   *HaClusterNodeEnum `json:"node,omitempty"`
    PortId *string            `json:"port_id,omitempty"`
    Vlan   *int               `json:"vlan,omitempty"`
    Vrf    *string            `json:"vrf,omitempty"`
}
