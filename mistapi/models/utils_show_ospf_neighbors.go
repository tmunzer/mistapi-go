package models

import (
    "encoding/json"
)

// UtilsShowOspfNeighbors represents a UtilsShowOspfNeighbors struct.
type UtilsShowOspfNeighbors struct {
    // Neighbor IP Address
    Neighbor             *string            `json:"neighbor,omitempty"`
    // only for HA. enum: `node0`, `node1`
    Node                 *HaClusterNodeEnum `json:"node,omitempty"`
    // the network interface
    PortId               *string            `json:"port_id,omitempty"`
    // VRF name
    Vrf                  *string            `json:"vrf,omitempty"`
    AdditionalProperties map[string]any     `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for UtilsShowOspfNeighbors.
// It customizes the JSON marshaling process for UtilsShowOspfNeighbors objects.
func (u UtilsShowOspfNeighbors) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(u.toMap())
}

// toMap converts the UtilsShowOspfNeighbors object to a map representation for JSON marshaling.
func (u UtilsShowOspfNeighbors) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, u.AdditionalProperties)
    if u.Neighbor != nil {
        structMap["neighbor"] = u.Neighbor
    }
    if u.Node != nil {
        structMap["node"] = u.Node
    }
    if u.PortId != nil {
        structMap["port_id"] = u.PortId
    }
    if u.Vrf != nil {
        structMap["vrf"] = u.Vrf
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UtilsShowOspfNeighbors.
// It customizes the JSON unmarshaling process for UtilsShowOspfNeighbors objects.
func (u *UtilsShowOspfNeighbors) UnmarshalJSON(input []byte) error {
    var temp tempUtilsShowOspfNeighbors
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "neighbor", "node", "port_id", "vrf")
    if err != nil {
    	return err
    }
    
    u.AdditionalProperties = additionalProperties
    u.Neighbor = temp.Neighbor
    u.Node = temp.Node
    u.PortId = temp.PortId
    u.Vrf = temp.Vrf
    return nil
}

// tempUtilsShowOspfNeighbors is a temporary struct used for validating the fields of UtilsShowOspfNeighbors.
type tempUtilsShowOspfNeighbors  struct {
    Neighbor *string            `json:"neighbor,omitempty"`
    Node     *HaClusterNodeEnum `json:"node,omitempty"`
    PortId   *string            `json:"port_id,omitempty"`
    Vrf      *string            `json:"vrf,omitempty"`
}