// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// TuntermPortConfig represents a TuntermPortConfig struct.
// Ethernet port configurations
type TuntermPortConfig struct {
	// List of ports to be used for downstream (to AP) purpose
	DownstreamPorts []string `json:"downstream_ports,omitempty"`
	// Whether to separate upstream / downstream ports. default is false where all ports will be used.
	SeparateUpstreamDownstream *bool `json:"separate_upstream_downstream,omitempty"`
	// Native VLAN id for upstream ports
	UpstreamPortVlanId *int `json:"upstream_port_vlan_id,omitempty"`
	// List of ports to be used for upstream purpose (to LAN)
	UpstreamPorts        []string               `json:"upstream_ports,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for TuntermPortConfig,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (t TuntermPortConfig) String() string {
	return fmt.Sprintf(
		"TuntermPortConfig[DownstreamPorts=%v, SeparateUpstreamDownstream=%v, UpstreamPortVlanId=%v, UpstreamPorts=%v, AdditionalProperties=%v]",
		t.DownstreamPorts, t.SeparateUpstreamDownstream, t.UpstreamPortVlanId, t.UpstreamPorts, t.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for TuntermPortConfig.
// It customizes the JSON marshaling process for TuntermPortConfig objects.
func (t TuntermPortConfig) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(t.AdditionalProperties,
		"downstream_ports", "separate_upstream_downstream", "upstream_port_vlan_id", "upstream_ports"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(t.toMap())
}

// toMap converts the TuntermPortConfig object to a map representation for JSON marshaling.
func (t TuntermPortConfig) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, t.AdditionalProperties)
	if t.DownstreamPorts != nil {
		structMap["downstream_ports"] = t.DownstreamPorts
	}
	if t.SeparateUpstreamDownstream != nil {
		structMap["separate_upstream_downstream"] = t.SeparateUpstreamDownstream
	}
	if t.UpstreamPortVlanId != nil {
		structMap["upstream_port_vlan_id"] = t.UpstreamPortVlanId
	}
	if t.UpstreamPorts != nil {
		structMap["upstream_ports"] = t.UpstreamPorts
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for TuntermPortConfig.
// It customizes the JSON unmarshaling process for TuntermPortConfig objects.
func (t *TuntermPortConfig) UnmarshalJSON(input []byte) error {
	var temp tempTuntermPortConfig
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "downstream_ports", "separate_upstream_downstream", "upstream_port_vlan_id", "upstream_ports")
	if err != nil {
		return err
	}
	t.AdditionalProperties = additionalProperties

	t.DownstreamPorts = temp.DownstreamPorts
	t.SeparateUpstreamDownstream = temp.SeparateUpstreamDownstream
	t.UpstreamPortVlanId = temp.UpstreamPortVlanId
	t.UpstreamPorts = temp.UpstreamPorts
	return nil
}

// tempTuntermPortConfig is a temporary struct used for validating the fields of TuntermPortConfig.
type tempTuntermPortConfig struct {
	DownstreamPorts            []string `json:"downstream_ports,omitempty"`
	SeparateUpstreamDownstream *bool    `json:"separate_upstream_downstream,omitempty"`
	UpstreamPortVlanId         *int     `json:"upstream_port_vlan_id,omitempty"`
	UpstreamPorts              []string `json:"upstream_ports,omitempty"`
}
