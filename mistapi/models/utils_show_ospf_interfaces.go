package models

import (
    "encoding/json"
    "fmt"
)

// UtilsShowOspfInterfaces represents a UtilsShowOspfInterfaces struct.
type UtilsShowOspfInterfaces struct {
    // only for HA. enum: `node0`, `node1`
    Node                 *HaClusterNodeEnum     `json:"node,omitempty"`
    // the network interface
    PortId               *string                `json:"port_id,omitempty"`
    // VRF name
    Vrf                  *string                `json:"vrf,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for UtilsShowOspfInterfaces,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UtilsShowOspfInterfaces) String() string {
    return fmt.Sprintf(
    	"UtilsShowOspfInterfaces[Node=%v, PortId=%v, Vrf=%v, AdditionalProperties=%v]",
    	u.Node, u.PortId, u.Vrf, u.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for UtilsShowOspfInterfaces.
// It customizes the JSON marshaling process for UtilsShowOspfInterfaces objects.
func (u UtilsShowOspfInterfaces) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(u.AdditionalProperties,
        "node", "port_id", "vrf"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UtilsShowOspfInterfaces object to a map representation for JSON marshaling.
func (u UtilsShowOspfInterfaces) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, u.AdditionalProperties)
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

// UnmarshalJSON implements the json.Unmarshaler interface for UtilsShowOspfInterfaces.
// It customizes the JSON unmarshaling process for UtilsShowOspfInterfaces objects.
func (u *UtilsShowOspfInterfaces) UnmarshalJSON(input []byte) error {
    var temp tempUtilsShowOspfInterfaces
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "node", "port_id", "vrf")
    if err != nil {
    	return err
    }
    u.AdditionalProperties = additionalProperties
    
    u.Node = temp.Node
    u.PortId = temp.PortId
    u.Vrf = temp.Vrf
    return nil
}

// tempUtilsShowOspfInterfaces is a temporary struct used for validating the fields of UtilsShowOspfInterfaces.
type tempUtilsShowOspfInterfaces  struct {
    Node   *HaClusterNodeEnum `json:"node,omitempty"`
    PortId *string            `json:"port_id,omitempty"`
    Vrf    *string            `json:"vrf,omitempty"`
}
