package models

import (
	"encoding/json"
)

// MxclusterRadsec represents a MxclusterRadsec struct.
// MxEdge Radsec Configuration
type MxclusterRadsec struct {
	// list of RADIUS accounting servers, optional, order matters where the first one is treated as primary
	AcctServers []MxclusterRadsecAcctServer `json:"acct_servers,omitempty"`
	// list of RADIUS authentication servers, order matters where the first one is treated as primary
	AuthServers []MxclusterRadsecAuthServer `json:"auth_servers,omitempty"`
	// whether to enable service on Mist Edge i.e. RADIUS proxy over TLS
	Enabled *bool `json:"enabled,omitempty"`
	// whether to match ssid in request message to select from a subset of RADIUS servers
	MatchSsid *bool `json:"match_ssid,omitempty"`
	// SSpecify NAS-IP-ADDRESS, NAS-IPv6-ADDRESS to use with auth_servers. enum: `any`, `oob`, `oob6`, `tunnel`, `tunnel6`
	NasIpSource *MxclusterRadsecNasIpSourceEnum `json:"nas_ip_source,omitempty"`
	// hostnames or IPs for Mist AP to use as the TLS Server (i.e. they are reachable from AP) in addition to `tunterm_hosts`
	ProxyHosts []string `json:"proxy_hosts,omitempty"`
	// When ordered, Mist Edge will prefer and go back to the first radius server if possible. enum: `ordered`, `unordered`
	ServerSelection *MxclusterRadsecServerSelectionEnum `json:"server_selection,omitempty"`
	// Specify IP address to connect to auth_servers and acct_servers. enum: `any`, `oob`, `oob6`, `tunnel`, `tunnel6`
	SrcIpSource          *MxclusterRadsecSrcIpSourceEnum `json:"src_ip_source,omitempty"`
	AdditionalProperties map[string]any                  `json:"_"`
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
	if m.NasIpSource != nil {
		structMap["nas_ip_source"] = m.NasIpSource
	}
	if m.ProxyHosts != nil {
		structMap["proxy_hosts"] = m.ProxyHosts
	}
	if m.ServerSelection != nil {
		structMap["server_selection"] = m.ServerSelection
	}
	if m.SrcIpSource != nil {
		structMap["src_ip_source"] = m.SrcIpSource
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MxclusterRadsec.
// It customizes the JSON unmarshaling process for MxclusterRadsec objects.
func (m *MxclusterRadsec) UnmarshalJSON(input []byte) error {
	var temp tempMxclusterRadsec
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "acct_servers", "auth_servers", "enabled", "match_ssid", "nas_ip_source", "proxy_hosts", "server_selection", "src_ip_source")
	if err != nil {
		return err
	}

	m.AdditionalProperties = additionalProperties
	m.AcctServers = temp.AcctServers
	m.AuthServers = temp.AuthServers
	m.Enabled = temp.Enabled
	m.MatchSsid = temp.MatchSsid
	m.NasIpSource = temp.NasIpSource
	m.ProxyHosts = temp.ProxyHosts
	m.ServerSelection = temp.ServerSelection
	m.SrcIpSource = temp.SrcIpSource
	return nil
}

// tempMxclusterRadsec is a temporary struct used for validating the fields of MxclusterRadsec.
type tempMxclusterRadsec struct {
	AcctServers     []MxclusterRadsecAcctServer         `json:"acct_servers,omitempty"`
	AuthServers     []MxclusterRadsecAuthServer         `json:"auth_servers,omitempty"`
	Enabled         *bool                               `json:"enabled,omitempty"`
	MatchSsid       *bool                               `json:"match_ssid,omitempty"`
	NasIpSource     *MxclusterRadsecNasIpSourceEnum     `json:"nas_ip_source,omitempty"`
	ProxyHosts      []string                            `json:"proxy_hosts,omitempty"`
	ServerSelection *MxclusterRadsecServerSelectionEnum `json:"server_selection,omitempty"`
	SrcIpSource     *MxclusterRadsecSrcIpSourceEnum     `json:"src_ip_source,omitempty"`
}
