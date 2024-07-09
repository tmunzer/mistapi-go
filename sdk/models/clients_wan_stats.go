package models

import (
    "encoding/json"
    "github.com/google/uuid"
)

// ClientsWanStats represents a ClientsWanStats struct.
type ClientsWanStats struct {
    When                 *string        `json:"When,omitempty"`
    Hostname             []string       `json:"hostname,omitempty"`
    Ip                   []string       `json:"ip,omitempty"`
    LastHostname         *string        `json:"last_hostname,omitempty"`
    LastIp               *string        `json:"last_ip,omitempty"`
    Mfg                  *string        `json:"mfg,omitempty"`
    Network              *string        `json:"network,omitempty"`
    OrgId                *uuid.UUID     `json:"org_id,omitempty"`
    SiteId               *uuid.UUID     `json:"site_id,omitempty"`
    Wcid                 *string        `json:"wcid,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ClientsWanStats.
// It customizes the JSON marshaling process for ClientsWanStats objects.
func (c ClientsWanStats) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the ClientsWanStats object to a map representation for JSON marshaling.
func (c ClientsWanStats) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    if c.When != nil {
        structMap["When"] = c.When
    }
    if c.Hostname != nil {
        structMap["hostname"] = c.Hostname
    }
    if c.Ip != nil {
        structMap["ip"] = c.Ip
    }
    if c.LastHostname != nil {
        structMap["last_hostname"] = c.LastHostname
    }
    if c.LastIp != nil {
        structMap["last_ip"] = c.LastIp
    }
    if c.Mfg != nil {
        structMap["mfg"] = c.Mfg
    }
    if c.Network != nil {
        structMap["network"] = c.Network
    }
    if c.OrgId != nil {
        structMap["org_id"] = c.OrgId
    }
    if c.SiteId != nil {
        structMap["site_id"] = c.SiteId
    }
    if c.Wcid != nil {
        structMap["wcid"] = c.Wcid
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ClientsWanStats.
// It customizes the JSON unmarshaling process for ClientsWanStats objects.
func (c *ClientsWanStats) UnmarshalJSON(input []byte) error {
    var temp clientsWanStats
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "When", "hostname", "ip", "last_hostname", "last_ip", "mfg", "network", "org_id", "site_id", "wcid")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.When = temp.When
    c.Hostname = temp.Hostname
    c.Ip = temp.Ip
    c.LastHostname = temp.LastHostname
    c.LastIp = temp.LastIp
    c.Mfg = temp.Mfg
    c.Network = temp.Network
    c.OrgId = temp.OrgId
    c.SiteId = temp.SiteId
    c.Wcid = temp.Wcid
    return nil
}

// clientsWanStats is a temporary struct used for validating the fields of ClientsWanStats.
type clientsWanStats  struct {
    When         *string    `json:"When,omitempty"`
    Hostname     []string   `json:"hostname,omitempty"`
    Ip           []string   `json:"ip,omitempty"`
    LastHostname *string    `json:"last_hostname,omitempty"`
    LastIp       *string    `json:"last_ip,omitempty"`
    Mfg          *string    `json:"mfg,omitempty"`
    Network      *string    `json:"network,omitempty"`
    OrgId        *uuid.UUID `json:"org_id,omitempty"`
    SiteId       *uuid.UUID `json:"site_id,omitempty"`
    Wcid         *string    `json:"wcid,omitempty"`
}
