package models

import (
	"encoding/json"
	"github.com/google/uuid"
)

// ServicePathEvent represents a ServicePathEvent struct.
type ServicePathEvent struct {
	Mac                  *string        `json:"mac,omitempty"`
	Model                *string        `json:"model,omitempty"`
	OrgId                *uuid.UUID     `json:"org_id,omitempty"`
	Policy               *string        `json:"policy,omitempty"`
	PortId               *string        `json:"port_id,omitempty"`
	SiteId               *uuid.UUID     `json:"site_id,omitempty"`
	Text                 *string        `json:"text,omitempty"`
	Timestamp            *float64       `json:"timestamp,omitempty"`
	Type                 *string        `json:"type,omitempty"`
	Version              *string        `json:"version,omitempty"`
	VpnName              *string        `json:"vpn_name,omitempty"`
	VpnPath              *string        `json:"vpn_path,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ServicePathEvent.
// It customizes the JSON marshaling process for ServicePathEvent objects.
func (s ServicePathEvent) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(s.toMap())
}

// toMap converts the ServicePathEvent object to a map representation for JSON marshaling.
func (s ServicePathEvent) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, s.AdditionalProperties)
	if s.Mac != nil {
		structMap["mac"] = s.Mac
	}
	if s.Model != nil {
		structMap["model"] = s.Model
	}
	if s.OrgId != nil {
		structMap["org_id"] = s.OrgId
	}
	if s.Policy != nil {
		structMap["policy"] = s.Policy
	}
	if s.PortId != nil {
		structMap["port_id"] = s.PortId
	}
	if s.SiteId != nil {
		structMap["site_id"] = s.SiteId
	}
	if s.Text != nil {
		structMap["text"] = s.Text
	}
	if s.Timestamp != nil {
		structMap["timestamp"] = s.Timestamp
	}
	if s.Type != nil {
		structMap["type"] = s.Type
	}
	if s.Version != nil {
		structMap["version"] = s.Version
	}
	if s.VpnName != nil {
		structMap["vpn_name"] = s.VpnName
	}
	if s.VpnPath != nil {
		structMap["vpn_path"] = s.VpnPath
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ServicePathEvent.
// It customizes the JSON unmarshaling process for ServicePathEvent objects.
func (s *ServicePathEvent) UnmarshalJSON(input []byte) error {
	var temp tempServicePathEvent
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "mac", "model", "org_id", "policy", "port_id", "site_id", "text", "timestamp", "type", "version", "vpn_name", "vpn_path")
	if err != nil {
		return err
	}

	s.AdditionalProperties = additionalProperties
	s.Mac = temp.Mac
	s.Model = temp.Model
	s.OrgId = temp.OrgId
	s.Policy = temp.Policy
	s.PortId = temp.PortId
	s.SiteId = temp.SiteId
	s.Text = temp.Text
	s.Timestamp = temp.Timestamp
	s.Type = temp.Type
	s.Version = temp.Version
	s.VpnName = temp.VpnName
	s.VpnPath = temp.VpnPath
	return nil
}

// tempServicePathEvent is a temporary struct used for validating the fields of ServicePathEvent.
type tempServicePathEvent struct {
	Mac       *string    `json:"mac,omitempty"`
	Model     *string    `json:"model,omitempty"`
	OrgId     *uuid.UUID `json:"org_id,omitempty"`
	Policy    *string    `json:"policy,omitempty"`
	PortId    *string    `json:"port_id,omitempty"`
	SiteId    *uuid.UUID `json:"site_id,omitempty"`
	Text      *string    `json:"text,omitempty"`
	Timestamp *float64   `json:"timestamp,omitempty"`
	Type      *string    `json:"type,omitempty"`
	Version   *string    `json:"version,omitempty"`
	VpnName   *string    `json:"vpn_name,omitempty"`
	VpnPath   *string    `json:"vpn_path,omitempty"`
}
