package models

import (
    "encoding/json"
)

// GatewayTemplateTunnelNode represents a GatewayTemplateTunnelNode struct.
type GatewayTemplateTunnelNode struct {
    Hosts                []string               `json:"hosts,omitempty"`
    // Only if:
    // * `provider`== `zscaler-gre`
    // * `provider`== `custom-gre`
    InternalIps          []string               `json:"internal_ips,omitempty"`
    ProbeIps             []string               `json:"probe_ips,omitempty"`
    // Only if `provider`== `custom-ipsec`
    RemoteIds            []string               `json:"remote_ids,omitempty"`
    WanNames             []string               `json:"wan_names,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for GatewayTemplateTunnelNode.
// It customizes the JSON marshaling process for GatewayTemplateTunnelNode objects.
func (g GatewayTemplateTunnelNode) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(g.AdditionalProperties,
        "hosts", "internal_ips", "probe_ips", "remote_ids", "wan_names"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(g.toMap())
}

// toMap converts the GatewayTemplateTunnelNode object to a map representation for JSON marshaling.
func (g GatewayTemplateTunnelNode) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, g.AdditionalProperties)
    if g.Hosts != nil {
        structMap["hosts"] = g.Hosts
    }
    if g.InternalIps != nil {
        structMap["internal_ips"] = g.InternalIps
    }
    if g.ProbeIps != nil {
        structMap["probe_ips"] = g.ProbeIps
    }
    if g.RemoteIds != nil {
        structMap["remote_ids"] = g.RemoteIds
    }
    if g.WanNames != nil {
        structMap["wan_names"] = g.WanNames
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GatewayTemplateTunnelNode.
// It customizes the JSON unmarshaling process for GatewayTemplateTunnelNode objects.
func (g *GatewayTemplateTunnelNode) UnmarshalJSON(input []byte) error {
    var temp tempGatewayTemplateTunnelNode
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "hosts", "internal_ips", "probe_ips", "remote_ids", "wan_names")
    if err != nil {
    	return err
    }
    g.AdditionalProperties = additionalProperties
    
    g.Hosts = temp.Hosts
    g.InternalIps = temp.InternalIps
    g.ProbeIps = temp.ProbeIps
    g.RemoteIds = temp.RemoteIds
    g.WanNames = temp.WanNames
    return nil
}

// tempGatewayTemplateTunnelNode is a temporary struct used for validating the fields of GatewayTemplateTunnelNode.
type tempGatewayTemplateTunnelNode  struct {
    Hosts       []string `json:"hosts,omitempty"`
    InternalIps []string `json:"internal_ips,omitempty"`
    ProbeIps    []string `json:"probe_ips,omitempty"`
    RemoteIds   []string `json:"remote_ids,omitempty"`
    WanNames    []string `json:"wan_names,omitempty"`
}
