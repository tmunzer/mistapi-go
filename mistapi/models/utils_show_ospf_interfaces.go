package models

import (
	"encoding/json"
)

// UtilsShowOspfInterfaces represents a UtilsShowOspfInterfaces struct.
type UtilsShowOspfInterfaces struct {
	// only for HA. enum: `node0`, `node1`
	Node *HaClusterNodeEnum `json:"node,omitempty"`
	// the network interface
	PortId *string `json:"port_id,omitempty"`
	// VRF name
	Vrf                  *string        `json:"vrf,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for UtilsShowOspfInterfaces.
// It customizes the JSON marshaling process for UtilsShowOspfInterfaces objects.
func (u UtilsShowOspfInterfaces) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(u.toMap())
}

// toMap converts the UtilsShowOspfInterfaces object to a map representation for JSON marshaling.
func (u UtilsShowOspfInterfaces) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, u.AdditionalProperties)
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
	additionalProperties, err := UnmarshalAdditionalProperties(input, "node", "port_id", "vrf")
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
type tempUtilsShowOspfInterfaces struct {
	Node   *HaClusterNodeEnum `json:"node,omitempty"`
	PortId *string            `json:"port_id,omitempty"`
	Vrf    *string            `json:"vrf,omitempty"`
}
