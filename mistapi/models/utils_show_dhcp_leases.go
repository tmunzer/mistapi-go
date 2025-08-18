// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// UtilsShowDhcpLeases represents a UtilsShowDhcpLeases struct.
type UtilsShowDhcpLeases struct {
	// DHCP network for the leases, returns full table if not specified
	Network string `json:"network"`
	// only for HA. enum: `node0`, `node1`
	Node                 *HaClusterNodeEnum     `json:"node,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for UtilsShowDhcpLeases,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UtilsShowDhcpLeases) String() string {
	return fmt.Sprintf(
		"UtilsShowDhcpLeases[Network=%v, Node=%v, AdditionalProperties=%v]",
		u.Network, u.Node, u.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for UtilsShowDhcpLeases.
// It customizes the JSON marshaling process for UtilsShowDhcpLeases objects.
func (u UtilsShowDhcpLeases) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(u.AdditionalProperties,
		"network", "node"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(u.toMap())
}

// toMap converts the UtilsShowDhcpLeases object to a map representation for JSON marshaling.
func (u UtilsShowDhcpLeases) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, u.AdditionalProperties)
	structMap["network"] = u.Network
	if u.Node != nil {
		structMap["node"] = u.Node
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UtilsShowDhcpLeases.
// It customizes the JSON unmarshaling process for UtilsShowDhcpLeases objects.
func (u *UtilsShowDhcpLeases) UnmarshalJSON(input []byte) error {
	var temp tempUtilsShowDhcpLeases
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "network", "node")
	if err != nil {
		return err
	}
	u.AdditionalProperties = additionalProperties

	u.Network = *temp.Network
	u.Node = temp.Node
	return nil
}

// tempUtilsShowDhcpLeases is a temporary struct used for validating the fields of UtilsShowDhcpLeases.
type tempUtilsShowDhcpLeases struct {
	Network *string            `json:"network"`
	Node    *HaClusterNodeEnum `json:"node,omitempty"`
}

func (u *tempUtilsShowDhcpLeases) validate() error {
	var errs []string
	if u.Network == nil {
		errs = append(errs, "required field `network` is missing for type `utils_show_dhcp_leases`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
