package models

import (
    "encoding/json"
)

// UtilsShowForwardingTable represents a UtilsShowForwardingTable struct.
type UtilsShowForwardingTable struct {
    // only for HA
    Node                 *HaClusterNodeEnum `json:"node,omitempty"`
    // IP Prefix
    Prefix               *string            `json:"prefix,omitempty"`
    // only supported with SSR
    ServiceIp            *string            `json:"service_ip,omitempty"`
    // only supported with SSR
    ServiceName          *string            `json:"service_name,omitempty"`
    // only supported with SSR
    ServicePort          *int               `json:"service_port,omitempty"`
    // only supported with SSR
    ServiceProtocol      *string            `json:"service_protocol,omitempty"`
    // only supported with SSR
    ServiceTenant        *string            `json:"service_tenant,omitempty"`
    // VRF Name
    Vrf                  *string            `json:"vrf,omitempty"`
    AdditionalProperties map[string]any     `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for UtilsShowForwardingTable.
// It customizes the JSON marshaling process for UtilsShowForwardingTable objects.
func (u UtilsShowForwardingTable) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(u.toMap())
}

// toMap converts the UtilsShowForwardingTable object to a map representation for JSON marshaling.
func (u UtilsShowForwardingTable) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, u.AdditionalProperties)
    if u.Node != nil {
        structMap["node"] = u.Node
    }
    if u.Prefix != nil {
        structMap["prefix"] = u.Prefix
    }
    if u.ServiceIp != nil {
        structMap["service_ip"] = u.ServiceIp
    }
    if u.ServiceName != nil {
        structMap["service_name"] = u.ServiceName
    }
    if u.ServicePort != nil {
        structMap["service_port"] = u.ServicePort
    }
    if u.ServiceProtocol != nil {
        structMap["service_protocol"] = u.ServiceProtocol
    }
    if u.ServiceTenant != nil {
        structMap["service_tenant"] = u.ServiceTenant
    }
    if u.Vrf != nil {
        structMap["vrf"] = u.Vrf
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UtilsShowForwardingTable.
// It customizes the JSON unmarshaling process for UtilsShowForwardingTable objects.
func (u *UtilsShowForwardingTable) UnmarshalJSON(input []byte) error {
    var temp utilsShowForwardingTable
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "node", "prefix", "service_ip", "service_name", "service_port", "service_protocol", "service_tenant", "vrf")
    if err != nil {
    	return err
    }
    
    u.AdditionalProperties = additionalProperties
    u.Node = temp.Node
    u.Prefix = temp.Prefix
    u.ServiceIp = temp.ServiceIp
    u.ServiceName = temp.ServiceName
    u.ServicePort = temp.ServicePort
    u.ServiceProtocol = temp.ServiceProtocol
    u.ServiceTenant = temp.ServiceTenant
    u.Vrf = temp.Vrf
    return nil
}

// utilsShowForwardingTable is a temporary struct used for validating the fields of UtilsShowForwardingTable.
type utilsShowForwardingTable  struct {
    Node            *HaClusterNodeEnum `json:"node,omitempty"`
    Prefix          *string            `json:"prefix,omitempty"`
    ServiceIp       *string            `json:"service_ip,omitempty"`
    ServiceName     *string            `json:"service_name,omitempty"`
    ServicePort     *int               `json:"service_port,omitempty"`
    ServiceProtocol *string            `json:"service_protocol,omitempty"`
    ServiceTenant   *string            `json:"service_tenant,omitempty"`
    Vrf             *string            `json:"vrf,omitempty"`
}
