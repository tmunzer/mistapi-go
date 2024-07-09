package models

import (
    "encoding/json"
    "github.com/google/uuid"
)

// EventOtherdevice represents a EventOtherdevice struct.
type EventOtherdevice struct {
    DeviceMac            *string        `json:"device_mac,omitempty"`
    Mac                  *string        `json:"mac,omitempty"`
    OrgId                *uuid.UUID     `json:"org_id,omitempty"`
    SiteId               *uuid.UUID     `json:"site_id,omitempty"`
    Text                 *string        `json:"text,omitempty"`
    Timestamp            *float64       `json:"timestamp,omitempty"`
    Type                 *string        `json:"type,omitempty"`
    Vendor               *string        `json:"vendor,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for EventOtherdevice.
// It customizes the JSON marshaling process for EventOtherdevice objects.
func (e EventOtherdevice) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(e.toMap())
}

// toMap converts the EventOtherdevice object to a map representation for JSON marshaling.
func (e EventOtherdevice) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, e.AdditionalProperties)
    if e.DeviceMac != nil {
        structMap["device_mac"] = e.DeviceMac
    }
    if e.Mac != nil {
        structMap["mac"] = e.Mac
    }
    if e.OrgId != nil {
        structMap["org_id"] = e.OrgId
    }
    if e.SiteId != nil {
        structMap["site_id"] = e.SiteId
    }
    if e.Text != nil {
        structMap["text"] = e.Text
    }
    if e.Timestamp != nil {
        structMap["timestamp"] = e.Timestamp
    }
    if e.Type != nil {
        structMap["type"] = e.Type
    }
    if e.Vendor != nil {
        structMap["vendor"] = e.Vendor
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for EventOtherdevice.
// It customizes the JSON unmarshaling process for EventOtherdevice objects.
func (e *EventOtherdevice) UnmarshalJSON(input []byte) error {
    var temp eventOtherdevice
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "device_mac", "mac", "org_id", "site_id", "text", "timestamp", "type", "vendor")
    if err != nil {
    	return err
    }
    
    e.AdditionalProperties = additionalProperties
    e.DeviceMac = temp.DeviceMac
    e.Mac = temp.Mac
    e.OrgId = temp.OrgId
    e.SiteId = temp.SiteId
    e.Text = temp.Text
    e.Timestamp = temp.Timestamp
    e.Type = temp.Type
    e.Vendor = temp.Vendor
    return nil
}

// eventOtherdevice is a temporary struct used for validating the fields of EventOtherdevice.
type eventOtherdevice  struct {
    DeviceMac *string    `json:"device_mac,omitempty"`
    Mac       *string    `json:"mac,omitempty"`
    OrgId     *uuid.UUID `json:"org_id,omitempty"`
    SiteId    *uuid.UUID `json:"site_id,omitempty"`
    Text      *string    `json:"text,omitempty"`
    Timestamp *float64   `json:"timestamp,omitempty"`
    Type      *string    `json:"type,omitempty"`
    Vendor    *string    `json:"vendor,omitempty"`
}