// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// TunnelConfigProbe represents a TunnelConfigProbe struct.
// Only if `provider`==`custom-ipsec`
type TunnelConfigProbe struct {
    // How often to trigger the probe
    Interval             *int                       `json:"interval,omitempty"`
    // Number of consecutive misses before declaring the tunnel down
    Threshold            *int                       `json:"threshold,omitempty"`
    // Time within which to complete the connectivity check
    Timeout              *int                       `json:"timeout,omitempty"`
    // enum: `http`, `icmp`
    Type                 *TunnelConfigProbeTypeEnum `json:"type,omitempty"`
    AdditionalProperties map[string]interface{}     `json:"_"`
}

// String implements the fmt.Stringer interface for TunnelConfigProbe,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (t TunnelConfigProbe) String() string {
    return fmt.Sprintf(
    	"TunnelConfigProbe[Interval=%v, Threshold=%v, Timeout=%v, Type=%v, AdditionalProperties=%v]",
    	t.Interval, t.Threshold, t.Timeout, t.Type, t.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for TunnelConfigProbe.
// It customizes the JSON marshaling process for TunnelConfigProbe objects.
func (t TunnelConfigProbe) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(t.AdditionalProperties,
        "interval", "threshold", "timeout", "type"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(t.toMap())
}

// toMap converts the TunnelConfigProbe object to a map representation for JSON marshaling.
func (t TunnelConfigProbe) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, t.AdditionalProperties)
    if t.Interval != nil {
        structMap["interval"] = t.Interval
    }
    if t.Threshold != nil {
        structMap["threshold"] = t.Threshold
    }
    if t.Timeout != nil {
        structMap["timeout"] = t.Timeout
    }
    if t.Type != nil {
        structMap["type"] = t.Type
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for TunnelConfigProbe.
// It customizes the JSON unmarshaling process for TunnelConfigProbe objects.
func (t *TunnelConfigProbe) UnmarshalJSON(input []byte) error {
    var temp tempTunnelConfigProbe
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "interval", "threshold", "timeout", "type")
    if err != nil {
    	return err
    }
    t.AdditionalProperties = additionalProperties
    
    t.Interval = temp.Interval
    t.Threshold = temp.Threshold
    t.Timeout = temp.Timeout
    t.Type = temp.Type
    return nil
}

// tempTunnelConfigProbe is a temporary struct used for validating the fields of TunnelConfigProbe.
type tempTunnelConfigProbe  struct {
    Interval  *int                       `json:"interval,omitempty"`
    Threshold *int                       `json:"threshold,omitempty"`
    Timeout   *int                       `json:"timeout,omitempty"`
    Type      *TunnelConfigProbeTypeEnum `json:"type,omitempty"`
}
