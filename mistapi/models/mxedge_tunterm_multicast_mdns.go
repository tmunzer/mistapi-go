package models

import (
    "encoding/json"
    "fmt"
)

// MxedgeTuntermMulticastMdns represents a MxedgeTuntermMulticastMdns struct.
type MxedgeTuntermMulticastMdns struct {
    Enabled              *bool                  `json:"enabled,omitempty"`
    VlanIds              []string               `json:"vlan_ids,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for MxedgeTuntermMulticastMdns,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m MxedgeTuntermMulticastMdns) String() string {
    return fmt.Sprintf(
    	"MxedgeTuntermMulticastMdns[Enabled=%v, VlanIds=%v, AdditionalProperties=%v]",
    	m.Enabled, m.VlanIds, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for MxedgeTuntermMulticastMdns.
// It customizes the JSON marshaling process for MxedgeTuntermMulticastMdns objects.
func (m MxedgeTuntermMulticastMdns) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(m.AdditionalProperties,
        "enabled", "vlan_ids"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(m.toMap())
}

// toMap converts the MxedgeTuntermMulticastMdns object to a map representation for JSON marshaling.
func (m MxedgeTuntermMulticastMdns) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, m.AdditionalProperties)
    if m.Enabled != nil {
        structMap["enabled"] = m.Enabled
    }
    if m.VlanIds != nil {
        structMap["vlan_ids"] = m.VlanIds
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MxedgeTuntermMulticastMdns.
// It customizes the JSON unmarshaling process for MxedgeTuntermMulticastMdns objects.
func (m *MxedgeTuntermMulticastMdns) UnmarshalJSON(input []byte) error {
    var temp tempMxedgeTuntermMulticastMdns
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "enabled", "vlan_ids")
    if err != nil {
    	return err
    }
    m.AdditionalProperties = additionalProperties
    
    m.Enabled = temp.Enabled
    m.VlanIds = temp.VlanIds
    return nil
}

// tempMxedgeTuntermMulticastMdns is a temporary struct used for validating the fields of MxedgeTuntermMulticastMdns.
type tempMxedgeTuntermMulticastMdns  struct {
    Enabled *bool    `json:"enabled,omitempty"`
    VlanIds []string `json:"vlan_ids,omitempty"`
}
