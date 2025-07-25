// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
    "github.com/google/uuid"
)

// WebhookLocationClientEvent represents a WebhookLocationClientEvent struct.
type WebhookLocationClientEvent struct {
    Mac                    *string                       `json:"mac,omitempty"`
    MapId                  *uuid.UUID                    `json:"map_id,omitempty"`
    SiteId                 *uuid.UUID                    `json:"site_id,omitempty"`
    // Epoch (seconds)
    Timestamp              *float64                      `json:"timestamp,omitempty"`
    Type                   *string                       `json:"type,omitempty"`
    // Optional, list of extended beacon info packets heard from the client, frame and sequence control included with the payload
    WifiBeaconExtendedInfo []WifiBeaconExtendedInfoItems `json:"wifi_beacon_extended_info,omitempty"`
    // x, in meter
    X                      *float64                      `json:"x,omitempty"`
    // y, in meter
    Y                      *float64                      `json:"y,omitempty"`
    AdditionalProperties   map[string]interface{}        `json:"_"`
}

// String implements the fmt.Stringer interface for WebhookLocationClientEvent,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WebhookLocationClientEvent) String() string {
    return fmt.Sprintf(
    	"WebhookLocationClientEvent[Mac=%v, MapId=%v, SiteId=%v, Timestamp=%v, Type=%v, WifiBeaconExtendedInfo=%v, X=%v, Y=%v, AdditionalProperties=%v]",
    	w.Mac, w.MapId, w.SiteId, w.Timestamp, w.Type, w.WifiBeaconExtendedInfo, w.X, w.Y, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WebhookLocationClientEvent.
// It customizes the JSON marshaling process for WebhookLocationClientEvent objects.
func (w WebhookLocationClientEvent) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(w.AdditionalProperties,
        "mac", "map_id", "site_id", "timestamp", "type", "wifi_beacon_extended_info", "x", "y"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WebhookLocationClientEvent object to a map representation for JSON marshaling.
func (w WebhookLocationClientEvent) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, w.AdditionalProperties)
    if w.Mac != nil {
        structMap["mac"] = w.Mac
    }
    if w.MapId != nil {
        structMap["map_id"] = w.MapId
    }
    if w.SiteId != nil {
        structMap["site_id"] = w.SiteId
    }
    if w.Timestamp != nil {
        structMap["timestamp"] = w.Timestamp
    }
    if w.Type != nil {
        structMap["type"] = w.Type
    }
    if w.WifiBeaconExtendedInfo != nil {
        structMap["wifi_beacon_extended_info"] = w.WifiBeaconExtendedInfo
    }
    if w.X != nil {
        structMap["x"] = w.X
    }
    if w.Y != nil {
        structMap["y"] = w.Y
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookLocationClientEvent.
// It customizes the JSON unmarshaling process for WebhookLocationClientEvent objects.
func (w *WebhookLocationClientEvent) UnmarshalJSON(input []byte) error {
    var temp tempWebhookLocationClientEvent
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "mac", "map_id", "site_id", "timestamp", "type", "wifi_beacon_extended_info", "x", "y")
    if err != nil {
    	return err
    }
    w.AdditionalProperties = additionalProperties
    
    w.Mac = temp.Mac
    w.MapId = temp.MapId
    w.SiteId = temp.SiteId
    w.Timestamp = temp.Timestamp
    w.Type = temp.Type
    w.WifiBeaconExtendedInfo = temp.WifiBeaconExtendedInfo
    w.X = temp.X
    w.Y = temp.Y
    return nil
}

// tempWebhookLocationClientEvent is a temporary struct used for validating the fields of WebhookLocationClientEvent.
type tempWebhookLocationClientEvent  struct {
    Mac                    *string                       `json:"mac,omitempty"`
    MapId                  *uuid.UUID                    `json:"map_id,omitempty"`
    SiteId                 *uuid.UUID                    `json:"site_id,omitempty"`
    Timestamp              *float64                      `json:"timestamp,omitempty"`
    Type                   *string                       `json:"type,omitempty"`
    WifiBeaconExtendedInfo []WifiBeaconExtendedInfoItems `json:"wifi_beacon_extended_info,omitempty"`
    X                      *float64                      `json:"x,omitempty"`
    Y                      *float64                      `json:"y,omitempty"`
}
