package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// Alarm represents a Alarm struct.
// additional information per alarm type
type Alarm struct {
    // UUID of the admin who acked the alarm
    AckAdminId           *uuid.UUID     `json:"ack_admin_id,omitempty"`
    // Name & Email ID of the admin who acked the alarm
    AckAdminName         *string        `json:"ack_admin_name,omitempty"`
    // Whether the alarm is acked or not
    Acked                *bool          `json:"acked,omitempty"`
    // Epoch (seconds) when the alarm was acked
    AckedTime            *int           `json:"acked_time,omitempty"`
    // additional information: List of MACs of the APs
    Aps                  []string       `json:"aps,omitempty"`
    // List of BSSIDs
    Bssids               []string       `json:"bssids,omitempty"`
    // Number of incident within an alarm window
    Count                int            `json:"count"`
    // additional information: List of MACs of the gateways
    Gateways             []string       `json:"gateways,omitempty"`
    // Group of the alarm
    Group                string         `json:"group"`
    // additional information: List of Hostnames of the devices (AP/Switch/Gateway)
    Hostnames            []string       `json:"hostnames,omitempty"`
    // UUID of the alarm
    Id                   uuid.UUID      `json:"id"`
    // Epoch (seconds) of the last incident/alarm within an alarm window
    LastSeen             int            `json:"last_seen"`
    // Text describing the alarm
    Note                 *string        `json:"note,omitempty"`
    OrgId                *uuid.UUID     `json:"org_id,omitempty"`
    // Severity of the alarm
    Severity             string         `json:"severity"`
    SiteId               *uuid.UUID     `json:"site_id,omitempty"`
    // List of SSIDs
    Ssids                []string       `json:"ssids,omitempty"`
    // additional information: List of MACs of the switches
    Switches             []string       `json:"switches,omitempty"`
    // Epoch (seconds) of the first incident/alarm
    Timestamp            int            `json:"timestamp"`
    // Key-name of the alarm type
    Type                 string         `json:"type"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for Alarm.
// It customizes the JSON marshaling process for Alarm objects.
func (a Alarm) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the Alarm object to a map representation for JSON marshaling.
func (a Alarm) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, a.AdditionalProperties)
    if a.AckAdminId != nil {
        structMap["ack_admin_id"] = a.AckAdminId
    }
    if a.AckAdminName != nil {
        structMap["ack_admin_name"] = a.AckAdminName
    }
    if a.Acked != nil {
        structMap["acked"] = a.Acked
    }
    if a.AckedTime != nil {
        structMap["acked_time"] = a.AckedTime
    }
    if a.Aps != nil {
        structMap["aps"] = a.Aps
    }
    if a.Bssids != nil {
        structMap["bssids"] = a.Bssids
    }
    structMap["count"] = a.Count
    if a.Gateways != nil {
        structMap["gateways"] = a.Gateways
    }
    structMap["group"] = a.Group
    if a.Hostnames != nil {
        structMap["hostnames"] = a.Hostnames
    }
    structMap["id"] = a.Id
    structMap["last_seen"] = a.LastSeen
    if a.Note != nil {
        structMap["note"] = a.Note
    }
    if a.OrgId != nil {
        structMap["org_id"] = a.OrgId
    }
    structMap["severity"] = a.Severity
    if a.SiteId != nil {
        structMap["site_id"] = a.SiteId
    }
    if a.Ssids != nil {
        structMap["ssids"] = a.Ssids
    }
    if a.Switches != nil {
        structMap["switches"] = a.Switches
    }
    structMap["timestamp"] = a.Timestamp
    structMap["type"] = a.Type
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Alarm.
// It customizes the JSON unmarshaling process for Alarm objects.
func (a *Alarm) UnmarshalJSON(input []byte) error {
    var temp alarm
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "ack_admin_id", "ack_admin_name", "acked", "acked_time", "aps", "bssids", "count", "gateways", "group", "hostnames", "id", "last_seen", "note", "org_id", "severity", "site_id", "ssids", "switches", "timestamp", "type")
    if err != nil {
    	return err
    }
    
    a.AdditionalProperties = additionalProperties
    a.AckAdminId = temp.AckAdminId
    a.AckAdminName = temp.AckAdminName
    a.Acked = temp.Acked
    a.AckedTime = temp.AckedTime
    a.Aps = temp.Aps
    a.Bssids = temp.Bssids
    a.Count = *temp.Count
    a.Gateways = temp.Gateways
    a.Group = *temp.Group
    a.Hostnames = temp.Hostnames
    a.Id = *temp.Id
    a.LastSeen = *temp.LastSeen
    a.Note = temp.Note
    a.OrgId = temp.OrgId
    a.Severity = *temp.Severity
    a.SiteId = temp.SiteId
    a.Ssids = temp.Ssids
    a.Switches = temp.Switches
    a.Timestamp = *temp.Timestamp
    a.Type = *temp.Type
    return nil
}

// alarm is a temporary struct used for validating the fields of Alarm.
type alarm  struct {
    AckAdminId   *uuid.UUID `json:"ack_admin_id,omitempty"`
    AckAdminName *string    `json:"ack_admin_name,omitempty"`
    Acked        *bool      `json:"acked,omitempty"`
    AckedTime    *int       `json:"acked_time,omitempty"`
    Aps          []string   `json:"aps,omitempty"`
    Bssids       []string   `json:"bssids,omitempty"`
    Count        *int       `json:"count"`
    Gateways     []string   `json:"gateways,omitempty"`
    Group        *string    `json:"group"`
    Hostnames    []string   `json:"hostnames,omitempty"`
    Id           *uuid.UUID `json:"id"`
    LastSeen     *int       `json:"last_seen"`
    Note         *string    `json:"note,omitempty"`
    OrgId        *uuid.UUID `json:"org_id,omitempty"`
    Severity     *string    `json:"severity"`
    SiteId       *uuid.UUID `json:"site_id,omitempty"`
    Ssids        []string   `json:"ssids,omitempty"`
    Switches     []string   `json:"switches,omitempty"`
    Timestamp    *int       `json:"timestamp"`
    Type         *string    `json:"type"`
}

func (a *alarm) validate() error {
    var errs []string
    if a.Count == nil {
        errs = append(errs, "required field `count` is missing for type `Alarm`")
    }
    if a.Group == nil {
        errs = append(errs, "required field `group` is missing for type `Alarm`")
    }
    if a.Id == nil {
        errs = append(errs, "required field `id` is missing for type `Alarm`")
    }
    if a.LastSeen == nil {
        errs = append(errs, "required field `last_seen` is missing for type `Alarm`")
    }
    if a.Severity == nil {
        errs = append(errs, "required field `severity` is missing for type `Alarm`")
    }
    if a.Timestamp == nil {
        errs = append(errs, "required field `timestamp` is missing for type `Alarm`")
    }
    if a.Type == nil {
        errs = append(errs, "required field `type` is missing for type `Alarm`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}