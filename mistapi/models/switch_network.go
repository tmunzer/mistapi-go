package models

import (
	"encoding/json"
	"errors"
	"strings"
)

// SwitchNetwork represents a SwitchNetwork struct.
// A network represents a network segment. It can either represent a VLAN (then usually ties to a L3 subnet), optionally associate it with a subnet which can later be used to create addition routes. Used for ports doing `family ethernet-switching`. It can also be a pure L3-subnet that can then be used against a port that with `family inet`.
type SwitchNetwork struct {
	// whether to stop clients to talk to each other, default is false (when enabled, a unique isolation_vlan_id is required)
	// NOTE: this features requires uplink device to also a be Juniper device and `inter_switch_link` to be set
	Isolation       *bool   `json:"isolation,omitempty"`
	IsolationVlanId *string `json:"isolation_vlan_id,omitempty"`
	// optional for pure switching, required when L3 / routing features are used
	Subnet               *string            `json:"subnet,omitempty"`
	VlanId               VlanIdWithVariable `json:"vlan_id"`
	AdditionalProperties map[string]any     `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SwitchNetwork.
// It customizes the JSON marshaling process for SwitchNetwork objects.
func (s SwitchNetwork) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(s.toMap())
}

// toMap converts the SwitchNetwork object to a map representation for JSON marshaling.
func (s SwitchNetwork) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, s.AdditionalProperties)
	if s.Isolation != nil {
		structMap["isolation"] = s.Isolation
	}
	if s.IsolationVlanId != nil {
		structMap["isolation_vlan_id"] = s.IsolationVlanId
	}
	if s.Subnet != nil {
		structMap["subnet"] = s.Subnet
	}
	structMap["vlan_id"] = s.VlanId.toMap()
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SwitchNetwork.
// It customizes the JSON unmarshaling process for SwitchNetwork objects.
func (s *SwitchNetwork) UnmarshalJSON(input []byte) error {
	var temp tempSwitchNetwork
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "isolation", "isolation_vlan_id", "subnet", "vlan_id")
	if err != nil {
		return err
	}

	s.AdditionalProperties = additionalProperties
	s.Isolation = temp.Isolation
	s.IsolationVlanId = temp.IsolationVlanId
	s.Subnet = temp.Subnet
	s.VlanId = *temp.VlanId
	return nil
}

// tempSwitchNetwork is a temporary struct used for validating the fields of SwitchNetwork.
type tempSwitchNetwork struct {
	Isolation       *bool               `json:"isolation,omitempty"`
	IsolationVlanId *string             `json:"isolation_vlan_id,omitempty"`
	Subnet          *string             `json:"subnet,omitempty"`
	VlanId          *VlanIdWithVariable `json:"vlan_id"`
}

func (s *tempSwitchNetwork) validate() error {
	var errs []string
	if s.VlanId == nil {
		errs = append(errs, "required field `vlan_id` is missing for type `switch_network`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
