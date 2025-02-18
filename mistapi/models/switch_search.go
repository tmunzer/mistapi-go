package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/google/uuid"
    "strings"
)

// SwitchSearch represents a SwitchSearch struct.
type SwitchSearch struct {
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
    NumMembers           *int                   `json:"num_members,omitempty"`
    OrgId                *uuid.UUID             `json:"org_id,omitempty"`
    Role                 *string                `json:"role,omitempty"`
    SiteId               *uuid.UUID             `json:"site_id,omitempty"`
    TimeDrifted          *bool                  `json:"time_drifted,omitempty"`
    Timestamp            *float64               `json:"timestamp,omitempty"`
    // Device Type. enum: `switch`
    Type                 string                 `json:"type"`
    Uptime               *int                   `json:"uptime,omitempty"`
    Version              *string                `json:"version,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SwitchSearch,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SwitchSearch) String() string {
    return fmt.Sprintf(
    	"SwitchSearch[Clustered=%v, EvpnMissingLinks=%v, EvpntopoId=%v, ExtIp=%v, Hostname=%v, Ip=%v, LastHostname=%v, LastTroubleCode=%v, LastTroubleTimestamp=%v, Mac=%v, Managed=%v, Model=%v, NumMembers=%v, OrgId=%v, Role=%v, SiteId=%v, TimeDrifted=%v, Timestamp=%v, Type=%v, Uptime=%v, Version=%v, AdditionalProperties=%v]",
    	s.Clustered, s.EvpnMissingLinks, s.EvpntopoId, s.ExtIp, s.Hostname, s.Ip, s.LastHostname, s.LastTroubleCode, s.LastTroubleTimestamp, s.Mac, s.Managed, s.Model, s.NumMembers, s.OrgId, s.Role, s.SiteId, s.TimeDrifted, s.Timestamp, s.Type, s.Uptime, s.Version, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SwitchSearch.
// It customizes the JSON marshaling process for SwitchSearch objects.
func (s SwitchSearch) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "clustered", "evpn_missing_links", "evpntopo_id", "ext_ip", "hostname", "ip", "last_hostname", "last_trouble_code", "last_trouble_timestamp", "mac", "managed", "model", "num_members", "org_id", "role", "site_id", "time_drifted", "timestamp", "type", "uptime", "version"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SwitchSearch object to a map representation for JSON marshaling.
func (s SwitchSearch) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Clustered != nil {
        structMap["clustered"] = s.Clustered
    }
    if s.EvpnMissingLinks != nil {
        structMap["evpn_missing_links"] = s.EvpnMissingLinks
    }
    if s.EvpntopoId != nil {
        structMap["evpntopo_id"] = s.EvpntopoId
    }
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
    if s.LastTroubleCode != nil {
        structMap["last_trouble_code"] = s.LastTroubleCode
    }
    if s.LastTroubleTimestamp != nil {
        structMap["last_trouble_timestamp"] = s.LastTroubleTimestamp
    }
    if s.Mac != nil {
        structMap["mac"] = s.Mac
    }
    if s.Managed != nil {
        structMap["managed"] = s.Managed
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
    if s.Role != nil {
        structMap["role"] = s.Role
    }
    if s.SiteId != nil {
        structMap["site_id"] = s.SiteId
    }
    if s.TimeDrifted != nil {
        structMap["time_drifted"] = s.TimeDrifted
    }
    if s.Timestamp != nil {
        structMap["timestamp"] = s.Timestamp
    }
    structMap["type"] = s.Type
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
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "clustered", "evpn_missing_links", "evpntopo_id", "ext_ip", "hostname", "ip", "last_hostname", "last_trouble_code", "last_trouble_timestamp", "mac", "managed", "model", "num_members", "org_id", "role", "site_id", "time_drifted", "timestamp", "type", "uptime", "version")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Clustered = temp.Clustered
    s.EvpnMissingLinks = temp.EvpnMissingLinks
    s.EvpntopoId = temp.EvpntopoId
    s.ExtIp = temp.ExtIp
    s.Hostname = temp.Hostname
    s.Ip = temp.Ip
    s.LastHostname = temp.LastHostname
    s.LastTroubleCode = temp.LastTroubleCode
    s.LastTroubleTimestamp = temp.LastTroubleTimestamp
    s.Mac = temp.Mac
    s.Managed = temp.Managed
    s.Model = temp.Model
    s.NumMembers = temp.NumMembers
    s.OrgId = temp.OrgId
    s.Role = temp.Role
    s.SiteId = temp.SiteId
    s.TimeDrifted = temp.TimeDrifted
    s.Timestamp = temp.Timestamp
    s.Type = *temp.Type
    s.Uptime = temp.Uptime
    s.Version = temp.Version
    return nil
}

// tempSwitchSearch is a temporary struct used for validating the fields of SwitchSearch.
type tempSwitchSearch  struct {
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
    NumMembers           *int       `json:"num_members,omitempty"`
    OrgId                *uuid.UUID `json:"org_id,omitempty"`
    Role                 *string    `json:"role,omitempty"`
    SiteId               *uuid.UUID `json:"site_id,omitempty"`
    TimeDrifted          *bool      `json:"time_drifted,omitempty"`
    Timestamp            *float64   `json:"timestamp,omitempty"`
    Type                 *string    `json:"type"`
    Uptime               *int       `json:"uptime,omitempty"`
    Version              *string    `json:"version,omitempty"`
}

func (s *tempSwitchSearch) validate() error {
    var errs []string
    if s.Type == nil {
        errs = append(errs, "required field `type` is missing for type `switch_search`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
