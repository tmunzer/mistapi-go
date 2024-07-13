package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// EventsClient represents a EventsClient struct.
// client events
type EventsClient struct {
    Ap                   *string        `json:"ap,omitempty"`
    Band                 Dot11BandEnum  `json:"band"`
    Bssid                *string        `json:"bssid,omitempty"`
    Channel              *int           `json:"channel,omitempty"`
    Proto                Dot11ProtoEnum `json:"proto"`
    Ssid                 *string        `json:"ssid,omitempty"`
    Text                 *string        `json:"text,omitempty"`
    Timestamp            float64        `json:"timestamp"`
    // event type, e.g. MARVIS_EVENT_CLIENT_FBT_FAILURE
    Type                 *string        `json:"type,omitempty"`
    // for assoc/disassoc events
    TypeCode             *int           `json:"type_code,omitempty"`
    WlanId               *uuid.UUID     `json:"wlan_id,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for EventsClient.
// It customizes the JSON marshaling process for EventsClient objects.
func (e EventsClient) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(e.toMap())
}

// toMap converts the EventsClient object to a map representation for JSON marshaling.
func (e EventsClient) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, e.AdditionalProperties)
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
    var temp eventsClient
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "ap", "band", "bssid", "channel", "proto", "ssid", "text", "timestamp", "type", "type_code", "wlan_id")
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

// eventsClient is a temporary struct used for validating the fields of EventsClient.
type eventsClient  struct {
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

func (e *eventsClient) validate() error {
    var errs []string
    if e.Band == nil {
        errs = append(errs, "required field `band` is missing for type `Events_Client`")
    }
    if e.Proto == nil {
        errs = append(errs, "required field `proto` is missing for type `Events_Client`")
    }
    if e.Timestamp == nil {
        errs = append(errs, "required field `timestamp` is missing for type `Events_Client`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
