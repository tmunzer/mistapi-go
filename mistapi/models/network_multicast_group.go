// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// NetworkMulticastGroup represents a NetworkMulticastGroup struct.
type NetworkMulticastGroup struct {
	// RP (rendezvous point) IP Address
	RpIp                 *string                `json:"rp_ip,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for NetworkMulticastGroup,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (n NetworkMulticastGroup) String() string {
	return fmt.Sprintf(
		"NetworkMulticastGroup[RpIp=%v, AdditionalProperties=%v]",
		n.RpIp, n.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for NetworkMulticastGroup.
// It customizes the JSON marshaling process for NetworkMulticastGroup objects.
func (n NetworkMulticastGroup) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(n.AdditionalProperties,
		"rp_ip"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(n.toMap())
}

// toMap converts the NetworkMulticastGroup object to a map representation for JSON marshaling.
func (n NetworkMulticastGroup) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, n.AdditionalProperties)
	if n.RpIp != nil {
		structMap["rp_ip"] = n.RpIp
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for NetworkMulticastGroup.
// It customizes the JSON unmarshaling process for NetworkMulticastGroup objects.
func (n *NetworkMulticastGroup) UnmarshalJSON(input []byte) error {
	var temp tempNetworkMulticastGroup
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "rp_ip")
	if err != nil {
		return err
	}
	n.AdditionalProperties = additionalProperties

	n.RpIp = temp.RpIp
	return nil
}

// tempNetworkMulticastGroup is a temporary struct used for validating the fields of NetworkMulticastGroup.
type tempNetworkMulticastGroup struct {
	RpIp *string `json:"rp_ip,omitempty"`
}
