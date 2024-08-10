package models

import (
    "encoding/json"
    "github.com/google/uuid"
)

// ClientsWanStats represents a ClientsWanStats struct.
type ClientsWanStats struct {
    DhcpExpireTime       *float64       `json:"dhcp_expire_time,omitempty"`
    DhcpStartTime        *float64       `json:"dhcp_start_time,omitempty"`
    Hostname             []string       `json:"hostname,omitempty"`
    Ip                   []string       `json:"ip,omitempty"`
    IpSrc                *string        `json:"ip_src,omitempty"`
    LastHostname         *string        `json:"last_hostname,omitempty"`
    LastIp               *string        `json:"last_ip,omitempty"`
    Mfg                  *string        `json:"mfg,omitempty"`
    Network              *string        `json:"network,omitempty"`
    OrgId                *uuid.UUID     `json:"org_id,omitempty"`
    SiteId               *uuid.UUID     `json:"site_id,omitempty"`
    Timestamp            *float64       `json:"timestamp,omitempty"`
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
    if c.DhcpExpireTime != nil {
        structMap["dhcp_expire_time"] = c.DhcpExpireTime
    }
    if c.DhcpStartTime != nil {
        structMap["dhcp_start_time"] = c.DhcpStartTime
    }
    if c.Hostname != nil {
        structMap["hostname"] = c.Hostname
    }
    if c.Ip != nil {
        structMap["ip"] = c.Ip
    }
    if c.IpSrc != nil {
        structMap["ip_src"] = c.IpSrc
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
    if c.Timestamp != nil {
        structMap["timestamp"] = c.Timestamp
    }
    if c.Wcid != nil {
        structMap["wcid"] = c.Wcid
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ClientsWanStats.
// It customizes the JSON unmarshaling process for ClientsWanStats objects.
func (c *ClientsWanStats) UnmarshalJSON(input []byte) error {
    var temp tempClientsWanStats
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "dhcp_expire_time", "dhcp_start_time", "hostname", "ip", "ip_src", "last_hostname", "last_ip", "mfg", "network", "org_id", "site_id", "timestamp", "wcid")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.DhcpExpireTime = temp.DhcpExpireTime
    c.DhcpStartTime = temp.DhcpStartTime
    c.Hostname = temp.Hostname
    c.Ip = temp.Ip
    c.IpSrc = temp.IpSrc
    c.LastHostname = temp.LastHostname
    c.LastIp = temp.LastIp
    c.Mfg = temp.Mfg
    c.Network = temp.Network
    c.OrgId = temp.OrgId
    c.SiteId = temp.SiteId
    c.Timestamp = temp.Timestamp
    c.Wcid = temp.Wcid
    return nil
}

// tempClientsWanStats is a temporary struct used for validating the fields of ClientsWanStats.
type tempClientsWanStats  struct {
    DhcpExpireTime *float64   `json:"dhcp_expire_time,omitempty"`
    DhcpStartTime  *float64   `json:"dhcp_start_time,omitempty"`
    Hostname       []string   `json:"hostname,omitempty"`
    Ip             []string   `json:"ip,omitempty"`
    IpSrc          *string    `json:"ip_src,omitempty"`
    LastHostname   *string    `json:"last_hostname,omitempty"`
    LastIp         *string    `json:"last_ip,omitempty"`
    Mfg            *string    `json:"mfg,omitempty"`
    Network        *string    `json:"network,omitempty"`
    OrgId          *uuid.UUID `json:"org_id,omitempty"`
    SiteId         *uuid.UUID `json:"site_id,omitempty"`
    Timestamp      *float64   `json:"timestamp,omitempty"`
    Wcid           *string    `json:"wcid,omitempty"`
}
