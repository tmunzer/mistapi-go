package models

import (
    "encoding/json"
    "github.com/google/uuid"
)

// GatewaySearch represents a GatewaySearch struct.
type GatewaySearch struct {
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

// MarshalJSON implements the json.Marshaler interface for GatewaySearch.
// It customizes the JSON marshaling process for GatewaySearch objects.
func (g GatewaySearch) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the GatewaySearch object to a map representation for JSON marshaling.
func (g GatewaySearch) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, g.AdditionalProperties)
    if g.ExtIp != nil {
        structMap["ext_ip"] = g.ExtIp
    }
    if g.Hostname != nil {
        structMap["hostname"] = g.Hostname
    }
    if g.Ip != nil {
        structMap["ip"] = g.Ip
    }
    if g.LastHostname != nil {
        structMap["last_hostname"] = g.LastHostname
    }
    if g.Mac != nil {
        structMap["mac"] = g.Mac
    }
    if g.Model != nil {
        structMap["model"] = g.Model
    }
    if g.NumMembers != nil {
        structMap["num_members"] = g.NumMembers
    }
    if g.OrgId != nil {
        structMap["org_id"] = g.OrgId
    }
    if g.SiteId != nil {
        structMap["site_id"] = g.SiteId
    }
    if g.Timestamp != nil {
        structMap["timestamp"] = g.Timestamp
    }
    if g.Type != nil {
        structMap["type"] = g.Type
    }
    if g.Uptime != nil {
        structMap["uptime"] = g.Uptime
    }
    if g.Version != nil {
        structMap["version"] = g.Version
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GatewaySearch.
// It customizes the JSON unmarshaling process for GatewaySearch objects.
func (g *GatewaySearch) UnmarshalJSON(input []byte) error {
    var temp tempGatewaySearch
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "ext_ip", "hostname", "ip", "last_hostname", "mac", "model", "num_members", "org_id", "site_id", "timestamp", "type", "uptime", "version")
    if err != nil {
    	return err
    }
    
    g.AdditionalProperties = additionalProperties
    g.ExtIp = temp.ExtIp
    g.Hostname = temp.Hostname
    g.Ip = temp.Ip
    g.LastHostname = temp.LastHostname
    g.Mac = temp.Mac
    g.Model = temp.Model
    g.NumMembers = temp.NumMembers
    g.OrgId = temp.OrgId
    g.SiteId = temp.SiteId
    g.Timestamp = temp.Timestamp
    g.Type = temp.Type
    g.Uptime = temp.Uptime
    g.Version = temp.Version
    return nil
}

// tempGatewaySearch is a temporary struct used for validating the fields of GatewaySearch.
type tempGatewaySearch  struct {
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
