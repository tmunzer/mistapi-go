package models

import (
    "encoding/json"
)

// UtilsTraceroute represents a UtilsTraceroute struct.
type UtilsTraceroute struct {
    // host name
    Host                 *string                      `json:"host,omitempty"`
    // for SSR, optional, the source to initiate traceroute from
    Network              *string                      `json:"network,omitempty"`
    // only for HA. enum: `node0`, `node1`
    Node                 *HaClusterNodeEnum           `json:"node,omitempty"`
    // when `protocol`==`udp`, not supported in SSR. The udp port to use
    Port                 *int                         `json:"port,omitempty"`
    // enum: `icmp` (Only suported by AP/MxEdge), `udp`
    Protocol             *UtilsTracerouteProtocolEnum `json:"protocol,omitempty"`
    // not supported in SSR. Maximum time in seconds to wait for the response
    Timeout              *int                         `json:"timeout,omitempty"`
    // for SRX, optional, the source to initiate traceroute from. by default, master VRF/RI is assumed
    Vrf                  *string                      `json:"vrf,omitempty"`
    AdditionalProperties map[string]any               `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for UtilsTraceroute.
// It customizes the JSON marshaling process for UtilsTraceroute objects.
func (u UtilsTraceroute) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(u.toMap())
}

// toMap converts the UtilsTraceroute object to a map representation for JSON marshaling.
func (u UtilsTraceroute) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, u.AdditionalProperties)
    if u.Host != nil {
        structMap["host"] = u.Host
    }
    if u.Network != nil {
        structMap["network"] = u.Network
    }
    if u.Node != nil {
        structMap["node"] = u.Node
    }
    if u.Port != nil {
        structMap["port"] = u.Port
    }
    if u.Protocol != nil {
        structMap["protocol"] = u.Protocol
    }
    if u.Timeout != nil {
        structMap["timeout"] = u.Timeout
    }
    if u.Vrf != nil {
        structMap["vrf"] = u.Vrf
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UtilsTraceroute.
// It customizes the JSON unmarshaling process for UtilsTraceroute objects.
func (u *UtilsTraceroute) UnmarshalJSON(input []byte) error {
    var temp tempUtilsTraceroute
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "host", "network", "node", "port", "protocol", "timeout", "vrf")
    if err != nil {
    	return err
    }
    
    u.AdditionalProperties = additionalProperties
    u.Host = temp.Host
    u.Network = temp.Network
    u.Node = temp.Node
    u.Port = temp.Port
    u.Protocol = temp.Protocol
    u.Timeout = temp.Timeout
    u.Vrf = temp.Vrf
    return nil
}

// tempUtilsTraceroute is a temporary struct used for validating the fields of UtilsTraceroute.
type tempUtilsTraceroute  struct {
    Host     *string                      `json:"host,omitempty"`
    Network  *string                      `json:"network,omitempty"`
    Node     *HaClusterNodeEnum           `json:"node,omitempty"`
    Port     *int                         `json:"port,omitempty"`
    Protocol *UtilsTracerouteProtocolEnum `json:"protocol,omitempty"`
    Timeout  *int                         `json:"timeout,omitempty"`
    Vrf      *string                      `json:"vrf,omitempty"`
}
