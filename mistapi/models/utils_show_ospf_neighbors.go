package models

import (
    "encoding/json"
    "fmt"
)

// UtilsShowOspfNeighbors represents a UtilsShowOspfNeighbors struct.
type UtilsShowOspfNeighbors struct {
    // Neighbor IP Address
    Neighbor             *string                `json:"neighbor,omitempty"`
    // only for HA. enum: `node0`, `node1`
    Node                 *HaClusterNodeEnum     `json:"node,omitempty"`
    // Network interface
    PortId               *string                `json:"port_id,omitempty"`
    // VRF name
    Vrf                  *string                `json:"vrf,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for UtilsShowOspfNeighbors,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UtilsShowOspfNeighbors) String() string {
    return fmt.Sprintf(
    	"UtilsShowOspfNeighbors[Neighbor=%v, Node=%v, PortId=%v, Vrf=%v, AdditionalProperties=%v]",
    	u.Neighbor, u.Node, u.PortId, u.Vrf, u.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for UtilsShowOspfNeighbors.
// It customizes the JSON marshaling process for UtilsShowOspfNeighbors objects.
func (u UtilsShowOspfNeighbors) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(u.AdditionalProperties,
        "neighbor", "node", "port_id", "vrf"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UtilsShowOspfNeighbors object to a map representation for JSON marshaling.
func (u UtilsShowOspfNeighbors) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, u.AdditionalProperties)
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
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "neighbor", "node", "port_id", "vrf")
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
