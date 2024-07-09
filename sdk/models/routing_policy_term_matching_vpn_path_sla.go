package models

import (
    "encoding/json"
)

// RoutingPolicyTermMatchingVpnPathSla represents a RoutingPolicyTermMatchingVpnPathSla struct.
type RoutingPolicyTermMatchingVpnPathSla struct {
    MaxJitter            Optional[int]  `json:"max_jitter"`
    MaxLatency           Optional[int]  `json:"max_latency"`
    MaxLoss              Optional[int]  `json:"max_loss"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for RoutingPolicyTermMatchingVpnPathSla.
// It customizes the JSON marshaling process for RoutingPolicyTermMatchingVpnPathSla objects.
func (r RoutingPolicyTermMatchingVpnPathSla) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the RoutingPolicyTermMatchingVpnPathSla object to a map representation for JSON marshaling.
func (r RoutingPolicyTermMatchingVpnPathSla) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    if r.MaxJitter.IsValueSet() {
        if r.MaxJitter.Value() != nil {
            structMap["max_jitter"] = r.MaxJitter.Value()
        } else {
            structMap["max_jitter"] = nil
        }
    }
    if r.MaxLatency.IsValueSet() {
        if r.MaxLatency.Value() != nil {
            structMap["max_latency"] = r.MaxLatency.Value()
        } else {
            structMap["max_latency"] = nil
        }
    }
    if r.MaxLoss.IsValueSet() {
        if r.MaxLoss.Value() != nil {
            structMap["max_loss"] = r.MaxLoss.Value()
        } else {
            structMap["max_loss"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RoutingPolicyTermMatchingVpnPathSla.
// It customizes the JSON unmarshaling process for RoutingPolicyTermMatchingVpnPathSla objects.
func (r *RoutingPolicyTermMatchingVpnPathSla) UnmarshalJSON(input []byte) error {
    var temp routingPolicyTermMatchingVpnPathSla
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "max_jitter", "max_latency", "max_loss")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.MaxJitter = temp.MaxJitter
    r.MaxLatency = temp.MaxLatency
    r.MaxLoss = temp.MaxLoss
    return nil
}

// routingPolicyTermMatchingVpnPathSla is a temporary struct used for validating the fields of RoutingPolicyTermMatchingVpnPathSla.
type routingPolicyTermMatchingVpnPathSla  struct {
    MaxJitter  Optional[int] `json:"max_jitter"`
    MaxLatency Optional[int] `json:"max_latency"`
    MaxLoss    Optional[int] `json:"max_loss"`
}
