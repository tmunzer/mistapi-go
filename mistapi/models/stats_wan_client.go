package models

import (
    "encoding/json"
    "github.com/google/uuid"
)

// StatsWanClient represents a StatsWanClient struct.
type StatsWanClient struct {
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

// MarshalJSON implements the json.Marshaler interface for StatsWanClient.
// It customizes the JSON marshaling process for StatsWanClient objects.
func (s StatsWanClient) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the StatsWanClient object to a map representation for JSON marshaling.
func (s StatsWanClient) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.DhcpExpireTime != nil {
        structMap["dhcp_expire_time"] = s.DhcpExpireTime
    }
    if s.DhcpStartTime != nil {
        structMap["dhcp_start_time"] = s.DhcpStartTime
    }
    if s.Hostname != nil {
        structMap["hostname"] = s.Hostname
    }
    if s.Ip != nil {
        structMap["ip"] = s.Ip
    }
    if s.IpSrc != nil {
        structMap["ip_src"] = s.IpSrc
    }
    if s.LastHostname != nil {
        structMap["last_hostname"] = s.LastHostname
    }
    if s.LastIp != nil {
        structMap["last_ip"] = s.LastIp
    }
    if s.Mfg != nil {
        structMap["mfg"] = s.Mfg
    }
    if s.Network != nil {
        structMap["network"] = s.Network
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
    if s.Wcid != nil {
        structMap["wcid"] = s.Wcid
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsWanClient.
// It customizes the JSON unmarshaling process for StatsWanClient objects.
func (s *StatsWanClient) UnmarshalJSON(input []byte) error {
    var temp tempStatsWanClient
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "dhcp_expire_time", "dhcp_start_time", "hostname", "ip", "ip_src", "last_hostname", "last_ip", "mfg", "network", "org_id", "site_id", "timestamp", "wcid")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.DhcpExpireTime = temp.DhcpExpireTime
    s.DhcpStartTime = temp.DhcpStartTime
    s.Hostname = temp.Hostname
    s.Ip = temp.Ip
    s.IpSrc = temp.IpSrc
    s.LastHostname = temp.LastHostname
    s.LastIp = temp.LastIp
    s.Mfg = temp.Mfg
    s.Network = temp.Network
    s.OrgId = temp.OrgId
    s.SiteId = temp.SiteId
    s.Timestamp = temp.Timestamp
    s.Wcid = temp.Wcid
    return nil
}

// tempStatsWanClient is a temporary struct used for validating the fields of StatsWanClient.
type tempStatsWanClient  struct {
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
