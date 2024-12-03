package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// UtilsReleaseDhcpLeases represents a UtilsReleaseDhcpLeases struct.
// Note: 
// * valid combinations for Junos: 
// * `port_id` 
// * `macs` + `network`
// * valid combinations for SSR: 
// * `port_id` 
// * `macs` + `network`
// * `port_id` + `network`
// * `network`
// * if network or port_id is specified and macs is empty, it means all clients under network or port_id
type UtilsReleaseDhcpLeases struct {
    // A list of client macs to be released
    Mac                  []string               `json:"mac,omitempty"`
    // The network for the leases IPs to be released
    Network              *string                `json:"network,omitempty"`
    // only for HA. enum: `node0`, `node1`
    Node                 *HaClusterNodeEnum     `json:"node,omitempty"`
    // The nework interface on which to release the current DHCP release
    PortId               string                 `json:"port_id"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for UtilsReleaseDhcpLeases.
// It customizes the JSON marshaling process for UtilsReleaseDhcpLeases objects.
func (u UtilsReleaseDhcpLeases) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(u.AdditionalProperties,
        "mac", "network", "node", "port_id"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UtilsReleaseDhcpLeases object to a map representation for JSON marshaling.
func (u UtilsReleaseDhcpLeases) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, u.AdditionalProperties)
    if u.Mac != nil {
        structMap["mac"] = u.Mac
    }
    if u.Network != nil {
        structMap["network"] = u.Network
    }
    if u.Node != nil {
        structMap["node"] = u.Node
    }
    structMap["port_id"] = u.PortId
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UtilsReleaseDhcpLeases.
// It customizes the JSON unmarshaling process for UtilsReleaseDhcpLeases objects.
func (u *UtilsReleaseDhcpLeases) UnmarshalJSON(input []byte) error {
    var temp tempUtilsReleaseDhcpLeases
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "mac", "network", "node", "port_id")
    if err != nil {
    	return err
    }
    u.AdditionalProperties = additionalProperties
    
    u.Mac = temp.Mac
    u.Network = temp.Network
    u.Node = temp.Node
    u.PortId = *temp.PortId
    return nil
}

// tempUtilsReleaseDhcpLeases is a temporary struct used for validating the fields of UtilsReleaseDhcpLeases.
type tempUtilsReleaseDhcpLeases  struct {
    Mac     []string           `json:"mac,omitempty"`
    Network *string            `json:"network,omitempty"`
    Node    *HaClusterNodeEnum `json:"node,omitempty"`
    PortId  *string            `json:"port_id"`
}

func (u *tempUtilsReleaseDhcpLeases) validate() error {
    var errs []string
    if u.PortId == nil {
        errs = append(errs, "required field `port_id` is missing for type `utils_release_dhcp_leases`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
