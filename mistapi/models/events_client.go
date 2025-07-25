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

// EventsClient represents a EventsClient struct.
// Client events
type EventsClient struct {
    Ap                   *string                `json:"ap,omitempty"`
    // enum: `24`, `5`, `6`
    Band                 Dot11BandEnum          `json:"band"`
    Bssid                *string                `json:"bssid,omitempty"`
    Channel              *int                   `json:"channel,omitempty"`
    // enum: `a`, `ac`, `ax`, `b`, `g`, `n`
    Proto                Dot11ProtoEnum         `json:"proto"`
    Ssid                 *string                `json:"ssid,omitempty"`
    Text                 *string                `json:"text,omitempty"`
    // Epoch (seconds)
    Timestamp            float64                `json:"timestamp"`
    // Event type, e.g. MARVIS_EVENT_CLIENT_FBT_FAILURE
    Type                 *string                `json:"type,omitempty"`
    // For assoc/disassoc events
    TypeCode             *int                   `json:"type_code,omitempty"`
    WlanId               *uuid.UUID             `json:"wlan_id,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for EventsClient,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (e EventsClient) String() string {
    return fmt.Sprintf(
    	"EventsClient[Ap=%v, Band=%v, Bssid=%v, Channel=%v, Proto=%v, Ssid=%v, Text=%v, Timestamp=%v, Type=%v, TypeCode=%v, WlanId=%v, AdditionalProperties=%v]",
    	e.Ap, e.Band, e.Bssid, e.Channel, e.Proto, e.Ssid, e.Text, e.Timestamp, e.Type, e.TypeCode, e.WlanId, e.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for EventsClient.
// It customizes the JSON marshaling process for EventsClient objects.
func (e EventsClient) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(e.AdditionalProperties,
        "ap", "band", "bssid", "channel", "proto", "ssid", "text", "timestamp", "type", "type_code", "wlan_id"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(e.toMap())
}

// toMap converts the EventsClient object to a map representation for JSON marshaling.
func (e EventsClient) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, e.AdditionalProperties)
    if e.Ap != nil {
        structMap["ap"] = e.Ap
    }
    structMap["band"] = e.Band
    if e.Bssid != nil {
        structMap["bssid"] = e.Bssid
    }
    if e.Channel != nil {
        structMap["channel"] = e.Channel
    }
    structMap["proto"] = e.Proto
    if e.Ssid != nil {
        structMap["ssid"] = e.Ssid
    }
    if e.Text != nil {
        structMap["text"] = e.Text
    }
    structMap["timestamp"] = e.Timestamp
    if e.Type != nil {
        structMap["type"] = e.Type
    }
    if e.TypeCode != nil {
        structMap["type_code"] = e.TypeCode
    }
    if e.WlanId != nil {
        structMap["wlan_id"] = e.WlanId
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for EventsClient.
// It customizes the JSON unmarshaling process for EventsClient objects.
func (e *EventsClient) UnmarshalJSON(input []byte) error {
    var temp tempEventsClient
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "ap", "band", "bssid", "channel", "proto", "ssid", "text", "timestamp", "type", "type_code", "wlan_id")
    if err != nil {
    	return err
    }
    e.AdditionalProperties = additionalProperties
    
    e.Ap = temp.Ap
    e.Band = *temp.Band
    e.Bssid = temp.Bssid
    e.Channel = temp.Channel
    e.Proto = *temp.Proto
    e.Ssid = temp.Ssid
    e.Text = temp.Text
    e.Timestamp = *temp.Timestamp
    e.Type = temp.Type
    e.TypeCode = temp.TypeCode
    e.WlanId = temp.WlanId
    return nil
}

// tempEventsClient is a temporary struct used for validating the fields of EventsClient.
type tempEventsClient  struct {
    Ap        *string         `json:"ap,omitempty"`
    Band      *Dot11BandEnum  `json:"band"`
    Bssid     *string         `json:"bssid,omitempty"`
    Channel   *int            `json:"channel,omitempty"`
    Proto     *Dot11ProtoEnum `json:"proto"`
    Ssid      *string         `json:"ssid,omitempty"`
    Text      *string         `json:"text,omitempty"`
    Timestamp *float64        `json:"timestamp"`
    Type      *string         `json:"type,omitempty"`
    TypeCode  *int            `json:"type_code,omitempty"`
    WlanId    *uuid.UUID      `json:"wlan_id,omitempty"`
}

func (e *tempEventsClient) validate() error {
    var errs []string
    if e.Band == nil {
        errs = append(errs, "required field `band` is missing for type `events_client`")
    }
    if e.Proto == nil {
        errs = append(errs, "required field `proto` is missing for type `events_client`")
    }
    if e.Timestamp == nil {
        errs = append(errs, "required field `timestamp` is missing for type `events_client`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
