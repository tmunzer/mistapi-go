package models

import (
    "encoding/json"
)

// GatewayTemplateTunnelProbe represents a GatewayTemplateTunnelProbe struct.
// Only if:
// * `provider`== `custom-ipsec`
type GatewayTemplateTunnelProbe struct {
    // how often to trigger the probe
    Interval             *int                          `json:"interval,omitempty"`
    // number of consecutive misses before declaring the tunnel down
    Threshold            *int                          `json:"threshold,omitempty"`
    // time within which to complete the connectivity check
    Timeout              *int                          `json:"timeout,omitempty"`
    Type                 *GatewayTemplateProbeTypeEnum `json:"type,omitempty"`
    AdditionalProperties map[string]any                `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for GatewayTemplateTunnelProbe.
// It customizes the JSON marshaling process for GatewayTemplateTunnelProbe objects.
func (g GatewayTemplateTunnelProbe) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the GatewayTemplateTunnelProbe object to a map representation for JSON marshaling.
func (g GatewayTemplateTunnelProbe) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, g.AdditionalProperties)
    if g.Interval != nil {
        structMap["interval"] = g.Interval
    }
    if g.Threshold != nil {
        structMap["threshold"] = g.Threshold
    }
    if g.Timeout != nil {
        structMap["timeout"] = g.Timeout
    }
    if g.Type != nil {
        structMap["type"] = g.Type
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GatewayTemplateTunnelProbe.
// It customizes the JSON unmarshaling process for GatewayTemplateTunnelProbe objects.
func (g *GatewayTemplateTunnelProbe) UnmarshalJSON(input []byte) error {
    var temp gatewayTemplateTunnelProbe
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "interval", "threshold", "timeout", "type")
    if err != nil {
    	return err
    }
    
    g.AdditionalProperties = additionalProperties
    g.Interval = temp.Interval
    g.Threshold = temp.Threshold
    g.Timeout = temp.Timeout
    g.Type = temp.Type
    return nil
}

// gatewayTemplateTunnelProbe is a temporary struct used for validating the fields of GatewayTemplateTunnelProbe.
type gatewayTemplateTunnelProbe  struct {
    Interval  *int                          `json:"interval,omitempty"`
    Threshold *int                          `json:"threshold,omitempty"`
    Timeout   *int                          `json:"timeout,omitempty"`
    Type      *GatewayTemplateProbeTypeEnum `json:"type,omitempty"`
}
