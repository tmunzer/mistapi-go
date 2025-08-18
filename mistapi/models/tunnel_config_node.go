// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// TunnelConfigNode represents a TunnelConfigNode struct.
// Only if `provider`==`zscaler-ipsec`, `provider`==`jse-ipsec` or `provider`==`custom-ipsec`
type TunnelConfigNode struct {
	Hosts []string `json:"hosts"`
	// Only if `provider`==`zscaler-gre`, `provider`==`jse-ipsec`, `provider`==`custom-ipsec` or `provider`==`custom-gre`
	InternalIps []string `json:"internal_ips,omitempty"`
	ProbeIps    []string `json:"probe_ips,omitempty"`
	// Only if  `provider`==`jse-ipsec` or `provider`==`custom-ipsec`
	RemoteIds            []string               `json:"remote_ids,omitempty"`
	WanNames             []string               `json:"wan_names"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for TunnelConfigNode,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (t TunnelConfigNode) String() string {
	return fmt.Sprintf(
		"TunnelConfigNode[Hosts=%v, InternalIps=%v, ProbeIps=%v, RemoteIds=%v, WanNames=%v, AdditionalProperties=%v]",
		t.Hosts, t.InternalIps, t.ProbeIps, t.RemoteIds, t.WanNames, t.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for TunnelConfigNode.
// It customizes the JSON marshaling process for TunnelConfigNode objects.
func (t TunnelConfigNode) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(t.AdditionalProperties,
		"hosts", "internal_ips", "probe_ips", "remote_ids", "wan_names"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(t.toMap())
}

// toMap converts the TunnelConfigNode object to a map representation for JSON marshaling.
func (t TunnelConfigNode) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, t.AdditionalProperties)
	structMap["hosts"] = t.Hosts
	if t.InternalIps != nil {
		structMap["internal_ips"] = t.InternalIps
	}
	if t.ProbeIps != nil {
		structMap["probe_ips"] = t.ProbeIps
	}
	if t.RemoteIds != nil {
		structMap["remote_ids"] = t.RemoteIds
	}
	structMap["wan_names"] = t.WanNames
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for TunnelConfigNode.
// It customizes the JSON unmarshaling process for TunnelConfigNode objects.
func (t *TunnelConfigNode) UnmarshalJSON(input []byte) error {
	var temp tempTunnelConfigNode
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "hosts", "internal_ips", "probe_ips", "remote_ids", "wan_names")
	if err != nil {
		return err
	}
	t.AdditionalProperties = additionalProperties

	t.Hosts = *temp.Hosts
	t.InternalIps = temp.InternalIps
	t.ProbeIps = temp.ProbeIps
	t.RemoteIds = temp.RemoteIds
	t.WanNames = *temp.WanNames
	return nil
}

// tempTunnelConfigNode is a temporary struct used for validating the fields of TunnelConfigNode.
type tempTunnelConfigNode struct {
	Hosts       *[]string `json:"hosts"`
	InternalIps []string  `json:"internal_ips,omitempty"`
	ProbeIps    []string  `json:"probe_ips,omitempty"`
	RemoteIds   []string  `json:"remote_ids,omitempty"`
	WanNames    *[]string `json:"wan_names"`
}

func (t *tempTunnelConfigNode) validate() error {
	var errs []string
	if t.Hosts == nil {
		errs = append(errs, "required field `hosts` is missing for type `tunnel_config_node`")
	}
	if t.WanNames == nil {
		errs = append(errs, "required field `wan_names` is missing for type `tunnel_config_node`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
