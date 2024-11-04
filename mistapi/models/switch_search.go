package models

import (
    "encoding/json"
    "github.com/google/uuid"
)

// SwitchSearch represents a SwitchSearch struct.
type SwitchSearch struct {
    ExtIp                *string        `json:"ext_ip,omitempty"`
    Hostname             []string       `json:"hostname,omitempty"`
    Ip                   *string        `json:"ip,omitempty"`
    LastHostname         *string        `json:"last_hostname,omitempty"`
    Mac                  *string        `json:"mac,omitempty"`
    Model                *string        `json:"model,omitempty"`
    NumMembers           *int           `json:"num_members,omitempty"`
    OrgId                *uuid.UUID     `json:"org_id,omitempty"`
    SiteId               *uuid.UUID     `json:"site_id,omitempty"`
    Timestamp            *float64       `json:"timestamp,omitempty"`
    Type                 *string        `json:"type,omitempty"`
    Uptime               *int           `json:"uptime,omitempty"`
    Version              *string        `json:"version,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SwitchSearch.
// It customizes the JSON marshaling process for SwitchSearch objects.
func (s SwitchSearch) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SwitchSearch object to a map representation for JSON marshaling.
func (s SwitchSearch) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.ExtIp != nil {
        structMap["ext_ip"] = s.ExtIp
    }
    if s.Hostname != nil {
        structMap["hostname"] = s.Hostname
    }
    if s.Ip != nil {
        structMap["ip"] = s.Ip
    }
    if s.LastHostname != nil {
        structMap["last_hostname"] = s.LastHostname
    }
    if s.Mac != nil {
        structMap["mac"] = s.Mac
    }
    if s.Model != nil {
        structMap["model"] = s.Model
    }
    if s.NumMembers != nil {
        structMap["num_members"] = s.NumMembers
    }
    if s.OrgId != nil {
        structMap["org_id"] = s.OrgId
    }
    if s.SiteId != nil {
        structMap["site_id"] = s.SiteId
    }
    if s.Timestamp != nil {
        structMap["timestamp"] = s.Timestamp
    }
    if s.Type != nil {
        structMap["type"] = s.Type
    }
    if s.Uptime != nil {
        structMap["uptime"] = s.Uptime
    }
    if s.Version != nil {
        structMap["version"] = s.Version
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SwitchSearch.
// It customizes the JSON unmarshaling process for SwitchSearch objects.
func (s *SwitchSearch) UnmarshalJSON(input []byte) error {
    var temp tempSwitchSearch
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "ext_ip", "hostname", "ip", "last_hostname", "mac", "model", "num_members", "org_id", "site_id", "timestamp", "type", "uptime", "version")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.ExtIp = temp.ExtIp
    s.Hostname = temp.Hostname
    s.Ip = temp.Ip
    s.LastHostname = temp.LastHostname
    s.Mac = temp.Mac
    s.Model = temp.Model
    s.NumMembers = temp.NumMembers
    s.OrgId = temp.OrgId
    s.SiteId = temp.SiteId
    s.Timestamp = temp.Timestamp
    s.Type = temp.Type
    s.Uptime = temp.Uptime
    s.Version = temp.Version
    return nil
}

// tempSwitchSearch is a temporary struct used for validating the fields of SwitchSearch.
type tempSwitchSearch  struct {
    ExtIp        *string    `json:"ext_ip,omitempty"`
    Hostname     []string   `json:"hostname,omitempty"`
    Ip           *string    `json:"ip,omitempty"`
    LastHostname *string    `json:"last_hostname,omitempty"`
    Mac          *string    `json:"mac,omitempty"`
    Model        *string    `json:"model,omitempty"`
    NumMembers   *int       `json:"num_members,omitempty"`
    OrgId        *uuid.UUID `json:"org_id,omitempty"`
    SiteId       *uuid.UUID `json:"site_id,omitempty"`
    Timestamp    *float64   `json:"timestamp,omitempty"`
    Type         *string    `json:"type,omitempty"`
    Uptime       *int       `json:"uptime,omitempty"`
    Version      *string    `json:"version,omitempty"`
}
