// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
)

// WebhookLocationCentrakEvent represents a WebhookLocationCentrakEvent struct.
type WebhookLocationCentrakEvent struct {
	// MAC address of the device
	Mac *string `json:"mac,omitempty"`
	// Map id
	MapId *string `json:"map_id,omitempty"`
	// Optional, BLE manufacturing company ID
	MfgCompanyId *int `json:"mfg_company_id,omitempty"`
	// Optional, BLE manufacturing data in hex byte-string format (i.e. "112233AABBCC")
	MfgData *string    `json:"mfg_data,omitempty"`
	SiteId  *uuid.UUID `json:"site_id,omitempty"`
	// Epoch (seconds)
	Timestamp *float64                             `json:"timestamp,omitempty"`
	Type      *WebhookLocationCentrakEventTypeEnum `json:"type,omitempty"`
	// Optional, list of extended beacon info packets heard from the client, frame and sequence control included with the payload
	WifiBeaconExtendedInfo []WifiBeaconExtendedInfoItems `json:"wifi_beacon_extended_info,omitempty"`
	// x, in meter
	X *float64 `json:"x,omitempty"`
	// y, in meter
	Y                    *float64               `json:"y,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for WebhookLocationCentrakEvent,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WebhookLocationCentrakEvent) String() string {
	return fmt.Sprintf(
		"WebhookLocationCentrakEvent[Mac=%v, MapId=%v, MfgCompanyId=%v, MfgData=%v, SiteId=%v, Timestamp=%v, Type=%v, WifiBeaconExtendedInfo=%v, X=%v, Y=%v, AdditionalProperties=%v]",
		w.Mac, w.MapId, w.MfgCompanyId, w.MfgData, w.SiteId, w.Timestamp, w.Type, w.WifiBeaconExtendedInfo, w.X, w.Y, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WebhookLocationCentrakEvent.
// It customizes the JSON marshaling process for WebhookLocationCentrakEvent objects.
func (w WebhookLocationCentrakEvent) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(w.AdditionalProperties,
		"mac", "map_id", "mfg_company_id", "mfg_data", "site_id", "timestamp", "type", "wifi_beacon_extended_info", "x", "y"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(w.toMap())
}

// toMap converts the WebhookLocationCentrakEvent object to a map representation for JSON marshaling.
func (w WebhookLocationCentrakEvent) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, w.AdditionalProperties)
	if w.Mac != nil {
		structMap["mac"] = w.Mac
	}
	if w.MapId != nil {
		structMap["map_id"] = w.MapId
	}
	if w.MfgCompanyId != nil {
		structMap["mfg_company_id"] = w.MfgCompanyId
	}
	if w.MfgData != nil {
		structMap["mfg_data"] = w.MfgData
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

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookLocationCentrakEvent.
// It customizes the JSON unmarshaling process for WebhookLocationCentrakEvent objects.
func (w *WebhookLocationCentrakEvent) UnmarshalJSON(input []byte) error {
	var temp tempWebhookLocationCentrakEvent
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "mac", "map_id", "mfg_company_id", "mfg_data", "site_id", "timestamp", "type", "wifi_beacon_extended_info", "x", "y")
	if err != nil {
		return err
	}
	w.AdditionalProperties = additionalProperties

	w.Mac = temp.Mac
	w.MapId = temp.MapId
	w.MfgCompanyId = temp.MfgCompanyId
	w.MfgData = temp.MfgData
	w.SiteId = temp.SiteId
	w.Timestamp = temp.Timestamp
	w.Type = temp.Type
	w.WifiBeaconExtendedInfo = temp.WifiBeaconExtendedInfo
	w.X = temp.X
	w.Y = temp.Y
	return nil
}

// tempWebhookLocationCentrakEvent is a temporary struct used for validating the fields of WebhookLocationCentrakEvent.
type tempWebhookLocationCentrakEvent struct {
	Mac                    *string                              `json:"mac,omitempty"`
	MapId                  *string                              `json:"map_id,omitempty"`
	MfgCompanyId           *int                                 `json:"mfg_company_id,omitempty"`
	MfgData                *string                              `json:"mfg_data,omitempty"`
	SiteId                 *uuid.UUID                           `json:"site_id,omitempty"`
	Timestamp              *float64                             `json:"timestamp,omitempty"`
	Type                   *WebhookLocationCentrakEventTypeEnum `json:"type,omitempty"`
	WifiBeaconExtendedInfo []WifiBeaconExtendedInfoItems        `json:"wifi_beacon_extended_info,omitempty"`
	X                      *float64                             `json:"x,omitempty"`
	Y                      *float64                             `json:"y,omitempty"`
}
