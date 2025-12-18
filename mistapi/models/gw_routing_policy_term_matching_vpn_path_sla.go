// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// GwRoutingPolicyTermMatchingVpnPathSla represents a GwRoutingPolicyTermMatchingVpnPathSla struct.
type GwRoutingPolicyTermMatchingVpnPathSla struct {
	MaxJitter            Optional[int]          `json:"max_jitter"`
	MaxLatency           Optional[int]          `json:"max_latency"`
	MaxLoss              Optional[int]          `json:"max_loss"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for GwRoutingPolicyTermMatchingVpnPathSla,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (g GwRoutingPolicyTermMatchingVpnPathSla) String() string {
	return fmt.Sprintf(
		"GwRoutingPolicyTermMatchingVpnPathSla[MaxJitter=%v, MaxLatency=%v, MaxLoss=%v, AdditionalProperties=%v]",
		g.MaxJitter, g.MaxLatency, g.MaxLoss, g.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for GwRoutingPolicyTermMatchingVpnPathSla.
// It customizes the JSON marshaling process for GwRoutingPolicyTermMatchingVpnPathSla objects.
func (g GwRoutingPolicyTermMatchingVpnPathSla) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(g.AdditionalProperties,
		"max_jitter", "max_latency", "max_loss"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(g.toMap())
}

// toMap converts the GwRoutingPolicyTermMatchingVpnPathSla object to a map representation for JSON marshaling.
func (g GwRoutingPolicyTermMatchingVpnPathSla) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, g.AdditionalProperties)
	if g.MaxJitter.IsValueSet() {
		if g.MaxJitter.Value() != nil {
			structMap["max_jitter"] = g.MaxJitter.Value()
		} else {
			structMap["max_jitter"] = nil
		}
	}
	if g.MaxLatency.IsValueSet() {
		if g.MaxLatency.Value() != nil {
			structMap["max_latency"] = g.MaxLatency.Value()
		} else {
			structMap["max_latency"] = nil
		}
	}
	if g.MaxLoss.IsValueSet() {
		if g.MaxLoss.Value() != nil {
			structMap["max_loss"] = g.MaxLoss.Value()
		} else {
			structMap["max_loss"] = nil
		}
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GwRoutingPolicyTermMatchingVpnPathSla.
// It customizes the JSON unmarshaling process for GwRoutingPolicyTermMatchingVpnPathSla objects.
func (g *GwRoutingPolicyTermMatchingVpnPathSla) UnmarshalJSON(input []byte) error {
	var temp tempGwRoutingPolicyTermMatchingVpnPathSla
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "max_jitter", "max_latency", "max_loss")
	if err != nil {
		return err
	}
	g.AdditionalProperties = additionalProperties

	g.MaxJitter = temp.MaxJitter
	g.MaxLatency = temp.MaxLatency
	g.MaxLoss = temp.MaxLoss
	return nil
}

// tempGwRoutingPolicyTermMatchingVpnPathSla is a temporary struct used for validating the fields of GwRoutingPolicyTermMatchingVpnPathSla.
type tempGwRoutingPolicyTermMatchingVpnPathSla struct {
	MaxJitter  Optional[int] `json:"max_jitter"`
	MaxLatency Optional[int] `json:"max_latency"`
	MaxLoss    Optional[int] `json:"max_loss"`
}
