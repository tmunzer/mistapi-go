// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// UtilsReleaseDhcp represents a UtilsReleaseDhcp struct.
type UtilsReleaseDhcp struct {
    // only for HA. enum: `node0`, `node1`
    Node                 *HaClusterNodeEnum     `json:"node,omitempty"`
    // The network interface on which to release the current DHCP release
    PortId               string                 `json:"port_id"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for UtilsReleaseDhcp,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UtilsReleaseDhcp) String() string {
    return fmt.Sprintf(
    	"UtilsReleaseDhcp[Node=%v, PortId=%v, AdditionalProperties=%v]",
    	u.Node, u.PortId, u.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for UtilsReleaseDhcp.
// It customizes the JSON marshaling process for UtilsReleaseDhcp objects.
func (u UtilsReleaseDhcp) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(u.AdditionalProperties,
        "node", "port_id"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UtilsReleaseDhcp object to a map representation for JSON marshaling.
func (u UtilsReleaseDhcp) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, u.AdditionalProperties)
    if u.Node != nil {
        structMap["node"] = u.Node
    }
    structMap["port_id"] = u.PortId
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UtilsReleaseDhcp.
// It customizes the JSON unmarshaling process for UtilsReleaseDhcp objects.
func (u *UtilsReleaseDhcp) UnmarshalJSON(input []byte) error {
    var temp tempUtilsReleaseDhcp
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "node", "port_id")
    if err != nil {
    	return err
    }
    u.AdditionalProperties = additionalProperties
    
    u.Node = temp.Node
    u.PortId = *temp.PortId
    return nil
}

// tempUtilsReleaseDhcp is a temporary struct used for validating the fields of UtilsReleaseDhcp.
type tempUtilsReleaseDhcp  struct {
    Node   *HaClusterNodeEnum `json:"node,omitempty"`
    PortId *string            `json:"port_id"`
}

func (u *tempUtilsReleaseDhcp) validate() error {
    var errs []string
    if u.PortId == nil {
        errs = append(errs, "required field `port_id` is missing for type `utils_release_dhcp`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
