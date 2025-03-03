package models

import (
    "encoding/json"
    "fmt"
)

// GatewayPortVpnPathTrafficShaping represents a GatewayPortVpnPathTrafficShaping struct.
// Only if the VPN `type`==`hub_spoke`
type GatewayPortVpnPathTrafficShaping struct {
    // percentages for different class of traffic: high / medium / low / best-effort. Sum must be equal to 100
    ClassPercentages     []int                  `json:"class_percentages,omitempty"`
    Enabled              *bool                  `json:"enabled,omitempty"`
    // Interface Transmit Cap in kbps
    MaxTxKbps            *int                   `json:"max_tx_kbps,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for GatewayPortVpnPathTrafficShaping,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (g GatewayPortVpnPathTrafficShaping) String() string {
    return fmt.Sprintf(
    	"GatewayPortVpnPathTrafficShaping[ClassPercentages=%v, Enabled=%v, MaxTxKbps=%v, AdditionalProperties=%v]",
    	g.ClassPercentages, g.Enabled, g.MaxTxKbps, g.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for GatewayPortVpnPathTrafficShaping.
// It customizes the JSON marshaling process for GatewayPortVpnPathTrafficShaping objects.
func (g GatewayPortVpnPathTrafficShaping) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(g.AdditionalProperties,
        "class_percentages", "enabled", "max_tx_kbps"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(g.toMap())
}

// toMap converts the GatewayPortVpnPathTrafficShaping object to a map representation for JSON marshaling.
func (g GatewayPortVpnPathTrafficShaping) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, g.AdditionalProperties)
    if g.ClassPercentages != nil {
        structMap["class_percentages"] = g.ClassPercentages
    }
    if g.Enabled != nil {
        structMap["enabled"] = g.Enabled
    }
    if g.MaxTxKbps != nil {
        structMap["max_tx_kbps"] = g.MaxTxKbps
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GatewayPortVpnPathTrafficShaping.
// It customizes the JSON unmarshaling process for GatewayPortVpnPathTrafficShaping objects.
func (g *GatewayPortVpnPathTrafficShaping) UnmarshalJSON(input []byte) error {
    var temp tempGatewayPortVpnPathTrafficShaping
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "class_percentages", "enabled", "max_tx_kbps")
    if err != nil {
    	return err
    }
    g.AdditionalProperties = additionalProperties
    
    g.ClassPercentages = temp.ClassPercentages
    g.Enabled = temp.Enabled
    g.MaxTxKbps = temp.MaxTxKbps
    return nil
}

// tempGatewayPortVpnPathTrafficShaping is a temporary struct used for validating the fields of GatewayPortVpnPathTrafficShaping.
type tempGatewayPortVpnPathTrafficShaping  struct {
    ClassPercentages []int `json:"class_percentages,omitempty"`
    Enabled          *bool `json:"enabled,omitempty"`
    MaxTxKbps        *int  `json:"max_tx_kbps,omitempty"`
}
