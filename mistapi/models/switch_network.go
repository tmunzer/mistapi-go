package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// SwitchNetwork represents a SwitchNetwork struct.
// A network represents a network segment. It can either represent a VLAN (then usually ties to a L3 subnet), optionally associate it with a subnet which can later be used to create addition routes. Used for ports doing `family ethernet-switching`. It can also be a pure L3-subnet that can then be used against a port that with `family inet`.
type SwitchNetwork struct {
    // Only required for EVPN-VXLAN networks, IPv4 Virtual Gateway
    Gateway              *string                `json:"gateway,omitempty"`
    // Only required for EVPN-VXLAN networks, IPv6 Virtual Gateway
    Gateway6             *string                `json:"gateway6,omitempty"`
    // whether to stop clients to talk to each other, default is false (when enabled, a unique isolation_vlan_id is required). NOTE: this features requires uplink device to also a be Juniper device and `inter_switch_link` to be set
    Isolation            *bool                  `json:"isolation,omitempty"`
    IsolationVlanId      *string                `json:"isolation_vlan_id,omitempty"`
    // Optional for pure switching, required when L3 / routing features are used
    Subnet               *string                `json:"subnet,omitempty"`
    // Optional for pure switching, required when L3 / routing features are used
    Subnet6              *string                `json:"subnet6,omitempty"`
    VlanId               VlanIdWithVariable     `json:"vlan_id"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SwitchNetwork,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SwitchNetwork) String() string {
    return fmt.Sprintf(
    	"SwitchNetwork[Gateway=%v, Gateway6=%v, Isolation=%v, IsolationVlanId=%v, Subnet=%v, Subnet6=%v, VlanId=%v, AdditionalProperties=%v]",
    	s.Gateway, s.Gateway6, s.Isolation, s.IsolationVlanId, s.Subnet, s.Subnet6, s.VlanId, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SwitchNetwork.
// It customizes the JSON marshaling process for SwitchNetwork objects.
func (s SwitchNetwork) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "gateway", "gateway6", "isolation", "isolation_vlan_id", "subnet", "subnet6", "vlan_id"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SwitchNetwork object to a map representation for JSON marshaling.
func (s SwitchNetwork) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Gateway != nil {
        structMap["gateway"] = s.Gateway
    }
    if s.Gateway6 != nil {
        structMap["gateway6"] = s.Gateway6
    }
    if s.Isolation != nil {
        structMap["isolation"] = s.Isolation
    }
    if s.IsolationVlanId != nil {
        structMap["isolation_vlan_id"] = s.IsolationVlanId
    }
    if s.Subnet != nil {
        structMap["subnet"] = s.Subnet
    }
    if s.Subnet6 != nil {
        structMap["subnet6"] = s.Subnet6
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
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "gateway", "gateway6", "isolation", "isolation_vlan_id", "subnet", "subnet6", "vlan_id")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Gateway = temp.Gateway
    s.Gateway6 = temp.Gateway6
    s.Isolation = temp.Isolation
    s.IsolationVlanId = temp.IsolationVlanId
    s.Subnet = temp.Subnet
    s.Subnet6 = temp.Subnet6
    s.VlanId = *temp.VlanId
    return nil
}

// tempSwitchNetwork is a temporary struct used for validating the fields of SwitchNetwork.
type tempSwitchNetwork  struct {
    Gateway         *string             `json:"gateway,omitempty"`
    Gateway6        *string             `json:"gateway6,omitempty"`
    Isolation       *bool               `json:"isolation,omitempty"`
    IsolationVlanId *string             `json:"isolation_vlan_id,omitempty"`
    Subnet          *string             `json:"subnet,omitempty"`
    Subnet6         *string             `json:"subnet6,omitempty"`
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
    return errors.New(strings.Join (errs, "\n"))
}
