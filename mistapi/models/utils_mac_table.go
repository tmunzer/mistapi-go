package models

import (
    "encoding/json"
)

// UtilsMacTable represents a UtilsMacTable struct.
type UtilsMacTable struct {
    MacAddress           *string        `json:"mac_address,omitempty"`
    PortId               *string        `json:"port_id,omitempty"`
    VlanId               *string        `json:"vlan_id,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for UtilsMacTable.
// It customizes the JSON marshaling process for UtilsMacTable objects.
func (u UtilsMacTable) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(u.toMap())
}

// toMap converts the UtilsMacTable object to a map representation for JSON marshaling.
func (u UtilsMacTable) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, u.AdditionalProperties)
    if u.MacAddress != nil {
        structMap["mac_address"] = u.MacAddress
    }
    if u.PortId != nil {
        structMap["port_id"] = u.PortId
    }
    if u.VlanId != nil {
        structMap["vlan_id"] = u.VlanId
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UtilsMacTable.
// It customizes the JSON unmarshaling process for UtilsMacTable objects.
func (u *UtilsMacTable) UnmarshalJSON(input []byte) error {
    var temp tempUtilsMacTable
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "mac_address", "port_id", "vlan_id")
    if err != nil {
    	return err
    }
    
    u.AdditionalProperties = additionalProperties
    u.MacAddress = temp.MacAddress
    u.PortId = temp.PortId
    u.VlanId = temp.VlanId
    return nil
}

// tempUtilsMacTable is a temporary struct used for validating the fields of UtilsMacTable.
type tempUtilsMacTable  struct {
    MacAddress *string `json:"mac_address,omitempty"`
    PortId     *string `json:"port_id,omitempty"`
    VlanId     *string `json:"vlan_id,omitempty"`
}
