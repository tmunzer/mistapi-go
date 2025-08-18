// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"strings"
)

// EventsSkyatp represents a EventsSkyatp struct.
// SkyATP events
type EventsSkyatp struct {
	DeviceMac   string    `json:"device_mac"`
	ForSite     *bool     `json:"for_site,omitempty"`
	Ip          string    `json:"ip"`
	Mac         string    `json:"mac"`
	OrgId       uuid.UUID `json:"org_id"`
	SiteId      uuid.UUID `json:"site_id"`
	ThreatLevel int       `json:"threat_level"`
	// Epoch (seconds)
	Timestamp            float64                `json:"timestamp"`
	Type                 string                 `json:"type"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for EventsSkyatp,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (e EventsSkyatp) String() string {
	return fmt.Sprintf(
		"EventsSkyatp[DeviceMac=%v, ForSite=%v, Ip=%v, Mac=%v, OrgId=%v, SiteId=%v, ThreatLevel=%v, Timestamp=%v, Type=%v, AdditionalProperties=%v]",
		e.DeviceMac, e.ForSite, e.Ip, e.Mac, e.OrgId, e.SiteId, e.ThreatLevel, e.Timestamp, e.Type, e.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for EventsSkyatp.
// It customizes the JSON marshaling process for EventsSkyatp objects.
func (e EventsSkyatp) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(e.AdditionalProperties,
		"device_mac", "for_site", "ip", "mac", "org_id", "site_id", "threat_level", "timestamp", "type"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(e.toMap())
}

// toMap converts the EventsSkyatp object to a map representation for JSON marshaling.
func (e EventsSkyatp) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, e.AdditionalProperties)
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
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "device_mac", "for_site", "ip", "mac", "org_id", "site_id", "threat_level", "timestamp", "type")
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
type tempEventsSkyatp struct {
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
