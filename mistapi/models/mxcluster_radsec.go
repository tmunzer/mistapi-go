package models

import (
    "encoding/json"
)

// MxclusterRadsec represents a MxclusterRadsec struct.
// MxEdge Radsec Configuration
type MxclusterRadsec struct {
    // list of RADIUS accounting servers, optional, order matters where the first one is treated as primary
    AcctServers          []MxclusterRadsecAcctServer         `json:"acct_servers,omitempty"`
    // list of RADIUS authentication servers, order matters where the first one is treated as primary
    AuthServers          []MxclusterRadsecAuthServer         `json:"auth_servers,omitempty"`
    // whether to enable service on Mist Edge i.e. RADIUS proxy over TLS
    Enabled              *bool                               `json:"enabled,omitempty"`
    // whether to match ssid in request message to select from a subset of RADIUS servers
    MatchSsid            *bool                               `json:"match_ssid,omitempty"`
    // hostnames or IPs for Mist AP to use as the TLS Server (i.e. they are reachable from AP) in addition to `tunterm_hosts`
    ProxyHosts           []string                            `json:"proxy_hosts,omitempty"`
    // ordered (default) / unordered. When ordered, Mist Edge will prefer and go back to the first radius server if possible
    ServerSelection      *MxclusterRadsecServerSelectionEnum `json:"server_selection,omitempty"`
    // Specify source address to use when connecting to RADIUS servers
    Source               *MxclusterRadsecSourceEnum          `json:"source,omitempty"`
    AdditionalProperties map[string]any                      `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for MxclusterRadsec.
// It customizes the JSON marshaling process for MxclusterRadsec objects.
func (m MxclusterRadsec) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(m.toMap())
}

// toMap converts the MxclusterRadsec object to a map representation for JSON marshaling.
func (m MxclusterRadsec) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, m.AdditionalProperties)
    if m.AcctServers != nil {
        structMap["acct_servers"] = m.AcctServers
    }
    if m.AuthServers != nil {
        structMap["auth_servers"] = m.AuthServers
    }
    if m.Enabled != nil {
        structMap["enabled"] = m.Enabled
    }
    if m.MatchSsid != nil {
        structMap["match_ssid"] = m.MatchSsid
    }
    if m.ProxyHosts != nil {
        structMap["proxy_hosts"] = m.ProxyHosts
    }
    if m.ServerSelection != nil {
        structMap["server_selection"] = m.ServerSelection
    }
    if m.Source != nil {
        structMap["source"] = m.Source
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MxclusterRadsec.
// It customizes the JSON unmarshaling process for MxclusterRadsec objects.
func (m *MxclusterRadsec) UnmarshalJSON(input []byte) error {
    var temp mxclusterRadsec
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "acct_servers", "auth_servers", "enabled", "match_ssid", "proxy_hosts", "server_selection", "source")
    if err != nil {
    	return err
    }
    
    m.AdditionalProperties = additionalProperties
    m.AcctServers = temp.AcctServers
    m.AuthServers = temp.AuthServers
    m.Enabled = temp.Enabled
    m.MatchSsid = temp.MatchSsid
    m.ProxyHosts = temp.ProxyHosts
    m.ServerSelection = temp.ServerSelection
    m.Source = temp.Source
    return nil
}

// mxclusterRadsec is a temporary struct used for validating the fields of MxclusterRadsec.
type mxclusterRadsec  struct {
    AcctServers     []MxclusterRadsecAcctServer         `json:"acct_servers,omitempty"`
    AuthServers     []MxclusterRadsecAuthServer         `json:"auth_servers,omitempty"`
    Enabled         *bool                               `json:"enabled,omitempty"`
    MatchSsid       *bool                               `json:"match_ssid,omitempty"`
    ProxyHosts      []string                            `json:"proxy_hosts,omitempty"`
    ServerSelection *MxclusterRadsecServerSelectionEnum `json:"server_selection,omitempty"`
    Source          *MxclusterRadsecSourceEnum          `json:"source,omitempty"`
}
