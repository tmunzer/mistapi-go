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

// EventsDeviceAp represents a EventsDeviceAp struct.
// AP events
type EventsDeviceAp struct {
    Ap                   *string                `json:"ap,omitempty"`
    Apfw                 *string                `json:"apfw,omitempty"`
    Count                *int                   `json:"count,omitempty"`
    DeviceType           *string                `json:"device_type,omitempty"`
    Mac                  *string                `json:"mac,omitempty"`
    OrgId                *uuid.UUID             `json:"org_id,omitempty"`
    PortId               *string                `json:"port_id,omitempty"`
    SiteId               *uuid.UUID             `json:"site_id,omitempty"`
    Text                 *string                `json:"text,omitempty"`
    // Epoch (seconds)
    Timestamp            float64                `json:"timestamp"`
    Type                 *string                `json:"type,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for EventsDeviceAp,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (e EventsDeviceAp) String() string {
    return fmt.Sprintf(
    	"EventsDeviceAp[Ap=%v, Apfw=%v, Count=%v, DeviceType=%v, Mac=%v, OrgId=%v, PortId=%v, SiteId=%v, Text=%v, Timestamp=%v, Type=%v, AdditionalProperties=%v]",
    	e.Ap, e.Apfw, e.Count, e.DeviceType, e.Mac, e.OrgId, e.PortId, e.SiteId, e.Text, e.Timestamp, e.Type, e.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for EventsDeviceAp.
// It customizes the JSON marshaling process for EventsDeviceAp objects.
func (e EventsDeviceAp) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(e.AdditionalProperties,
        "ap", "apfw", "count", "device_type", "mac", "org_id", "port_id", "site_id", "text", "timestamp", "type"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(e.toMap())
}

// toMap converts the EventsDeviceAp object to a map representation for JSON marshaling.
func (e EventsDeviceAp) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, e.AdditionalProperties)
    if e.Ap != nil {
        structMap["ap"] = e.Ap
    }
    if e.Apfw != nil {
        structMap["apfw"] = e.Apfw
    }
    if e.Count != nil {
        structMap["count"] = e.Count
    }
    if e.DeviceType != nil {
        structMap["device_type"] = e.DeviceType
    }
    if e.Mac != nil {
        structMap["mac"] = e.Mac
    }
    if e.OrgId != nil {
        structMap["org_id"] = e.OrgId
    }
    if e.PortId != nil {
        structMap["port_id"] = e.PortId
    }
    if e.SiteId != nil {
        structMap["site_id"] = e.SiteId
    }
    if e.Text != nil {
        structMap["text"] = e.Text
    }
    structMap["timestamp"] = e.Timestamp
    if e.Type != nil {
        structMap["type"] = e.Type
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for EventsDeviceAp.
// It customizes the JSON unmarshaling process for EventsDeviceAp objects.
func (e *EventsDeviceAp) UnmarshalJSON(input []byte) error {
    var temp tempEventsDeviceAp
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "ap", "apfw", "count", "device_type", "mac", "org_id", "port_id", "site_id", "text", "timestamp", "type")
    if err != nil {
    	return err
    }
    e.AdditionalProperties = additionalProperties
    
    e.Ap = temp.Ap
    e.Apfw = temp.Apfw
    e.Count = temp.Count
    e.DeviceType = temp.DeviceType
    e.Mac = temp.Mac
    e.OrgId = temp.OrgId
    e.PortId = temp.PortId
    e.SiteId = temp.SiteId
    e.Text = temp.Text
    e.Timestamp = *temp.Timestamp
    e.Type = temp.Type
    return nil
}

// tempEventsDeviceAp is a temporary struct used for validating the fields of EventsDeviceAp.
type tempEventsDeviceAp  struct {
    Ap         *string    `json:"ap,omitempty"`
    Apfw       *string    `json:"apfw,omitempty"`
    Count      *int       `json:"count,omitempty"`
    DeviceType *string    `json:"device_type,omitempty"`
    Mac        *string    `json:"mac,omitempty"`
    OrgId      *uuid.UUID `json:"org_id,omitempty"`
    PortId     *string    `json:"port_id,omitempty"`
    SiteId     *uuid.UUID `json:"site_id,omitempty"`
    Text       *string    `json:"text,omitempty"`
    Timestamp  *float64   `json:"timestamp"`
    Type       *string    `json:"type,omitempty"`
}

func (e *tempEventsDeviceAp) validate() error {
    var errs []string
    if e.Timestamp == nil {
        errs = append(errs, "required field `timestamp` is missing for type `events_device_ap`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
