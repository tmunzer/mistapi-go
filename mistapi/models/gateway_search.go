package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/google/uuid"
    "strings"
)

// GatewaySearch represents a GatewaySearch struct.
type GatewaySearch struct {
    Clustered            *bool                  `json:"clustered,omitempty"`
    EvpnMissingLinks     *bool                  `json:"evpn_missing_links,omitempty"`
    EvpntopoId           *string                `json:"evpntopo_id,omitempty"`
    ExtIp                *string                `json:"ext_ip,omitempty"`
    Hostname             []string               `json:"hostname,omitempty"`
    Ip                   *string                `json:"ip,omitempty"`
    LastHostname         *string                `json:"last_hostname,omitempty"`
    LastTroubleCode      *string                `json:"last_trouble_code,omitempty"`
    LastTroubleTimestamp *int                   `json:"last_trouble_timestamp,omitempty"`
    Mac                  *string                `json:"mac,omitempty"`
    Managed              *bool                  `json:"managed,omitempty"`
    Model                *string                `json:"model,omitempty"`
    Node                 *string                `json:"node,omitempty"`
    Node0Mac             *string                `json:"node0_mac,omitempty"`
    Node1Mac             *string                `json:"node1_mac,omitempty"`
    NumMembers           *int                   `json:"num_members,omitempty"`
    OrgId                *uuid.UUID             `json:"org_id,omitempty"`
    Role                 *string                `json:"role,omitempty"`
    SiteId               *uuid.UUID             `json:"site_id,omitempty"`
    T128agentVersion     *string                `json:"t128agent_version,omitempty"`
    TimeDrifted          *bool                  `json:"time_drifted,omitempty"`
    Timestamp            *float64               `json:"timestamp,omitempty"`
    // Device Type. enum: `gateway`
    Type                 string                 `json:"type"`
    Uptime               *int                   `json:"uptime,omitempty"`
    Version              *string                `json:"version,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for GatewaySearch,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (g GatewaySearch) String() string {
    return fmt.Sprintf(
    	"GatewaySearch[Clustered=%v, EvpnMissingLinks=%v, EvpntopoId=%v, ExtIp=%v, Hostname=%v, Ip=%v, LastHostname=%v, LastTroubleCode=%v, LastTroubleTimestamp=%v, Mac=%v, Managed=%v, Model=%v, Node=%v, Node0Mac=%v, Node1Mac=%v, NumMembers=%v, OrgId=%v, Role=%v, SiteId=%v, T128agentVersion=%v, TimeDrifted=%v, Timestamp=%v, Type=%v, Uptime=%v, Version=%v, AdditionalProperties=%v]",
    	g.Clustered, g.EvpnMissingLinks, g.EvpntopoId, g.ExtIp, g.Hostname, g.Ip, g.LastHostname, g.LastTroubleCode, g.LastTroubleTimestamp, g.Mac, g.Managed, g.Model, g.Node, g.Node0Mac, g.Node1Mac, g.NumMembers, g.OrgId, g.Role, g.SiteId, g.T128agentVersion, g.TimeDrifted, g.Timestamp, g.Type, g.Uptime, g.Version, g.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for GatewaySearch.
// It customizes the JSON marshaling process for GatewaySearch objects.
func (g GatewaySearch) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(g.AdditionalProperties,
        "clustered", "evpn_missing_links", "evpntopo_id", "ext_ip", "hostname", "ip", "last_hostname", "last_trouble_code", "last_trouble_timestamp", "mac", "managed", "model", "node", "node0_mac", "node1_mac", "num_members", "org_id", "role", "site_id", "t128agent_version", "time_drifted", "timestamp", "type", "uptime", "version"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(g.toMap())
}

// toMap converts the GatewaySearch object to a map representation for JSON marshaling.
func (g GatewaySearch) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, g.AdditionalProperties)
    if g.Clustered != nil {
        structMap["clustered"] = g.Clustered
    }
    if g.EvpnMissingLinks != nil {
        structMap["evpn_missing_links"] = g.EvpnMissingLinks
    }
    if g.EvpntopoId != nil {
        structMap["evpntopo_id"] = g.EvpntopoId
    }
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
    if g.LastTroubleCode != nil {
        structMap["last_trouble_code"] = g.LastTroubleCode
    }
    if g.LastTroubleTimestamp != nil {
        structMap["last_trouble_timestamp"] = g.LastTroubleTimestamp
    }
    if g.Mac != nil {
        structMap["mac"] = g.Mac
    }
    if g.Managed != nil {
        structMap["managed"] = g.Managed
    }
    if g.Model != nil {
        structMap["model"] = g.Model
    }
    if g.Node != nil {
        structMap["node"] = g.Node
    }
    if g.Node0Mac != nil {
        structMap["node0_mac"] = g.Node0Mac
    }
    if g.Node1Mac != nil {
        structMap["node1_mac"] = g.Node1Mac
    }
    if g.NumMembers != nil {
        structMap["num_members"] = g.NumMembers
    }
    if g.OrgId != nil {
        structMap["org_id"] = g.OrgId
    }
    if g.Role != nil {
        structMap["role"] = g.Role
    }
    if g.SiteId != nil {
        structMap["site_id"] = g.SiteId
    }
    if g.T128agentVersion != nil {
        structMap["t128agent_version"] = g.T128agentVersion
    }
    if g.TimeDrifted != nil {
        structMap["time_drifted"] = g.TimeDrifted
    }
    if g.Timestamp != nil {
        structMap["timestamp"] = g.Timestamp
    }
    structMap["type"] = g.Type
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
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "clustered", "evpn_missing_links", "evpntopo_id", "ext_ip", "hostname", "ip", "last_hostname", "last_trouble_code", "last_trouble_timestamp", "mac", "managed", "model", "node", "node0_mac", "node1_mac", "num_members", "org_id", "role", "site_id", "t128agent_version", "time_drifted", "timestamp", "type", "uptime", "version")
    if err != nil {
    	return err
    }
    g.AdditionalProperties = additionalProperties
    
    g.Clustered = temp.Clustered
    g.EvpnMissingLinks = temp.EvpnMissingLinks
    g.EvpntopoId = temp.EvpntopoId
    g.ExtIp = temp.ExtIp
    g.Hostname = temp.Hostname
    g.Ip = temp.Ip
    g.LastHostname = temp.LastHostname
    g.LastTroubleCode = temp.LastTroubleCode
    g.LastTroubleTimestamp = temp.LastTroubleTimestamp
    g.Mac = temp.Mac
    g.Managed = temp.Managed
    g.Model = temp.Model
    g.Node = temp.Node
    g.Node0Mac = temp.Node0Mac
    g.Node1Mac = temp.Node1Mac
    g.NumMembers = temp.NumMembers
    g.OrgId = temp.OrgId
    g.Role = temp.Role
    g.SiteId = temp.SiteId
    g.T128agentVersion = temp.T128agentVersion
    g.TimeDrifted = temp.TimeDrifted
    g.Timestamp = temp.Timestamp
    g.Type = *temp.Type
    g.Uptime = temp.Uptime
    g.Version = temp.Version
    return nil
}

// tempGatewaySearch is a temporary struct used for validating the fields of GatewaySearch.
type tempGatewaySearch  struct {
    Clustered            *bool      `json:"clustered,omitempty"`
    EvpnMissingLinks     *bool      `json:"evpn_missing_links,omitempty"`
    EvpntopoId           *string    `json:"evpntopo_id,omitempty"`
    ExtIp                *string    `json:"ext_ip,omitempty"`
    Hostname             []string   `json:"hostname,omitempty"`
    Ip                   *string    `json:"ip,omitempty"`
    LastHostname         *string    `json:"last_hostname,omitempty"`
    LastTroubleCode      *string    `json:"last_trouble_code,omitempty"`
    LastTroubleTimestamp *int       `json:"last_trouble_timestamp,omitempty"`
    Mac                  *string    `json:"mac,omitempty"`
    Managed              *bool      `json:"managed,omitempty"`
    Model                *string    `json:"model,omitempty"`
    Node                 *string    `json:"node,omitempty"`
    Node0Mac             *string    `json:"node0_mac,omitempty"`
    Node1Mac             *string    `json:"node1_mac,omitempty"`
    NumMembers           *int       `json:"num_members,omitempty"`
    OrgId                *uuid.UUID `json:"org_id,omitempty"`
    Role                 *string    `json:"role,omitempty"`
    SiteId               *uuid.UUID `json:"site_id,omitempty"`
    T128agentVersion     *string    `json:"t128agent_version,omitempty"`
    TimeDrifted          *bool      `json:"time_drifted,omitempty"`
    Timestamp            *float64   `json:"timestamp,omitempty"`
    Type                 *string    `json:"type"`
    Uptime               *int       `json:"uptime,omitempty"`
    Version              *string    `json:"version,omitempty"`
}

func (g *tempGatewaySearch) validate() error {
    var errs []string
    if g.Type == nil {
        errs = append(errs, "required field `type` is missing for type `gateway_search`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
