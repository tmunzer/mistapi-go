package models

import (
	"encoding/json"
)

// SiteMxtunnelRadsec represents a SiteMxtunnelRadsec struct.
type SiteMxtunnelRadsec struct {
	AcctServers          []RadiusAcctServer `json:"acct_servers,omitempty"`
	AuthServers          []RadiusAuthServer `json:"auth_servers,omitempty"`
	Enabled              *bool              `json:"enabled,omitempty"`
	UseMxedge            *bool              `json:"use_mxedge,omitempty"`
	AdditionalProperties map[string]any     `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SiteMxtunnelRadsec.
// It customizes the JSON marshaling process for SiteMxtunnelRadsec objects.
func (s SiteMxtunnelRadsec) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(s.toMap())
}

// toMap converts the SiteMxtunnelRadsec object to a map representation for JSON marshaling.
func (s SiteMxtunnelRadsec) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, s.AdditionalProperties)
	if s.AcctServers != nil {
		structMap["acct_servers"] = s.AcctServers
	}
	if s.AuthServers != nil {
		structMap["auth_servers"] = s.AuthServers
	}
	if s.Enabled != nil {
		structMap["enabled"] = s.Enabled
	}
	if s.UseMxedge != nil {
		structMap["use_mxedge"] = s.UseMxedge
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SiteMxtunnelRadsec.
// It customizes the JSON unmarshaling process for SiteMxtunnelRadsec objects.
func (s *SiteMxtunnelRadsec) UnmarshalJSON(input []byte) error {
	var temp tempSiteMxtunnelRadsec
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "acct_servers", "auth_servers", "enabled", "use_mxedge")
	if err != nil {
		return err
	}

	s.AdditionalProperties = additionalProperties
	s.AcctServers = temp.AcctServers
	s.AuthServers = temp.AuthServers
	s.Enabled = temp.Enabled
	s.UseMxedge = temp.UseMxedge
	return nil
}

// tempSiteMxtunnelRadsec is a temporary struct used for validating the fields of SiteMxtunnelRadsec.
type tempSiteMxtunnelRadsec struct {
	AcctServers []RadiusAcctServer `json:"acct_servers,omitempty"`
	AuthServers []RadiusAuthServer `json:"auth_servers,omitempty"`
	Enabled     *bool              `json:"enabled,omitempty"`
	UseMxedge   *bool              `json:"use_mxedge,omitempty"`
}
