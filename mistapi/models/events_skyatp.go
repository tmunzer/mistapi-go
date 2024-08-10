package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// EventsSkyatp represents a EventsSkyatp struct.
// SkyATP events
type EventsSkyatp struct {
    DeviceMac            string         `json:"device_mac"`
    ForSite              *bool          `json:"for_site,omitempty"`
    Ip                   string         `json:"ip"`
    Mac                  string         `json:"mac"`
    OrgId                uuid.UUID      `json:"org_id"`
    SiteId               uuid.UUID      `json:"site_id"`
    ThreatLevel          int            `json:"threat_level"`
    Timestamp            float64        `json:"timestamp"`
    Type                 string         `json:"type"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for EventsSkyatp.
// It customizes the JSON marshaling process for EventsSkyatp objects.
func (e EventsSkyatp) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(e.toMap())
}

// toMap converts the EventsSkyatp object to a map representation for JSON marshaling.
func (e EventsSkyatp) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, e.AdditionalProperties)
    structMap["device_mac"] = e.DeviceMac
    if e.ForSite != nil {
        structMap["for_site"] = e.ForSite
    }
    structMap["ip"] = e.Ip
    structMap["mac"] = e.Mac
    structMap["org_id"] = e.OrgId
    structMap["site_id"] = e.SiteId
    structMap["threat_level"] = e.ThreatLevel
    structMap["timestamp"] = e.Timestamp
    structMap["type"] = e.Type
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for EventsSkyatp.
// It customizes the JSON unmarshaling process for EventsSkyatp objects.
func (e *EventsSkyatp) UnmarshalJSON(input []byte) error {
    var temp tempEventsSkyatp
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "device_mac", "for_site", "ip", "mac", "org_id", "site_id", "threat_level", "timestamp", "type")
    if err != nil {
    	return err
    }
    
    e.AdditionalProperties = additionalProperties
    e.DeviceMac = *temp.DeviceMac
    e.ForSite = temp.ForSite
    e.Ip = *temp.Ip
    e.Mac = *temp.Mac
    e.OrgId = *temp.OrgId
    e.SiteId = *temp.SiteId
    e.ThreatLevel = *temp.ThreatLevel
    e.Timestamp = *temp.Timestamp
    e.Type = *temp.Type
    return nil
}

// tempEventsSkyatp is a temporary struct used for validating the fields of EventsSkyatp.
type tempEventsSkyatp  struct {
    DeviceMac   *string    `json:"device_mac"`
    ForSite     *bool      `json:"for_site,omitempty"`
    Ip          *string    `json:"ip"`
    Mac         *string    `json:"mac"`
    OrgId       *uuid.UUID `json:"org_id"`
    SiteId      *uuid.UUID `json:"site_id"`
    ThreatLevel *int       `json:"threat_level"`
    Timestamp   *float64   `json:"timestamp"`
    Type        *string    `json:"type"`
}

func (e *tempEventsSkyatp) validate() error {
    var errs []string
    if e.DeviceMac == nil {
        errs = append(errs, "required field `device_mac` is missing for type `events_skyatp`")
    }
    if e.Ip == nil {
        errs = append(errs, "required field `ip` is missing for type `events_skyatp`")
    }
    if e.Mac == nil {
        errs = append(errs, "required field `mac` is missing for type `events_skyatp`")
    }
    if e.OrgId == nil {
        errs = append(errs, "required field `org_id` is missing for type `events_skyatp`")
    }
    if e.SiteId == nil {
        errs = append(errs, "required field `site_id` is missing for type `events_skyatp`")
    }
    if e.ThreatLevel == nil {
        errs = append(errs, "required field `threat_level` is missing for type `events_skyatp`")
    }
    if e.Timestamp == nil {
        errs = append(errs, "required field `timestamp` is missing for type `events_skyatp`")
    }
    if e.Type == nil {
        errs = append(errs, "required field `type` is missing for type `events_skyatp`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
