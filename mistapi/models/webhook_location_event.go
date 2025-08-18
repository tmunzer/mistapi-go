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

// WebhookLocationEvent represents a WebhookLocationEvent struct.
type WebhookLocationEvent struct {
	BatteryVoltage        *int       `json:"battery_voltage,omitempty"`
	EddystoneUidInstance  *string    `json:"eddystone_uid_instance,omitempty"`
	EddystoneUidNamespace *string    `json:"eddystone_uid_namespace,omitempty"`
	EddystoneUrlUrl       *string    `json:"eddystone_url_url,omitempty"`
	IbeaconMajor          *int       `json:"ibeacon_major,omitempty"`
	IbeaconMinor          *int       `json:"ibeacon_minor,omitempty"`
	IbeaconUuid           *uuid.UUID `json:"ibeacon_uuid,omitempty"`
	// Unique ID of the object instance in the Mist Organization
	Id  uuid.UUID `json:"id"`
	Mac *string   `json:"mac,omitempty"`
	// Map id
	MapId uuid.UUID `json:"map_id"`
	// Optional, BLE manufacturing company ID
	MfgCompanyId *int `json:"mfg_company_id,omitempty"`
	// Optional, BLE manufacturing data in hex byte-string format (ie "112233AABBCC")
	MfgData *string `json:"mfg_data,omitempty"`
	// Name of the client, may be empty
	Name   *string   `json:"name,omitempty"`
	SiteId uuid.UUID `json:"site_id"`
	// Epoch (seconds)
	Timestamp float64 `json:"timestamp"`
	Type      string  `json:"type"`
	// Optional, list of extended beacon info packets heard from the client, frame and sequence control included with the payload
	WifiBeaconExtendedInfo []WifiBeaconExtendedInfoItems `json:"wifi_beacon_extended_info,omitempty"`
	// x, in meter
	X float64 `json:"x"`
	// y, in meter
	Y                    float64                `json:"y"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for WebhookLocationEvent,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WebhookLocationEvent) String() string {
	return fmt.Sprintf(
		"WebhookLocationEvent[BatteryVoltage=%v, EddystoneUidInstance=%v, EddystoneUidNamespace=%v, EddystoneUrlUrl=%v, IbeaconMajor=%v, IbeaconMinor=%v, IbeaconUuid=%v, Id=%v, Mac=%v, MapId=%v, MfgCompanyId=%v, MfgData=%v, Name=%v, SiteId=%v, Timestamp=%v, Type=%v, WifiBeaconExtendedInfo=%v, X=%v, Y=%v, AdditionalProperties=%v]",
		w.BatteryVoltage, w.EddystoneUidInstance, w.EddystoneUidNamespace, w.EddystoneUrlUrl, w.IbeaconMajor, w.IbeaconMinor, w.IbeaconUuid, w.Id, w.Mac, w.MapId, w.MfgCompanyId, w.MfgData, w.Name, w.SiteId, w.Timestamp, w.Type, w.WifiBeaconExtendedInfo, w.X, w.Y, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WebhookLocationEvent.
// It customizes the JSON marshaling process for WebhookLocationEvent objects.
func (w WebhookLocationEvent) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(w.AdditionalProperties,
		"battery_voltage", "eddystone_uid_instance", "eddystone_uid_namespace", "eddystone_url_url", "ibeacon_major", "ibeacon_minor", "ibeacon_uuid", "id", "mac", "map_id", "mfg_company_id", "mfg_data", "name", "site_id", "timestamp", "type", "wifi_beacon_extended_info", "x", "y"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(w.toMap())
}

// toMap converts the WebhookLocationEvent object to a map representation for JSON marshaling.
func (w WebhookLocationEvent) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, w.AdditionalProperties)
	if w.BatteryVoltage != nil {
		structMap["battery_voltage"] = w.BatteryVoltage
	}
	if w.EddystoneUidInstance != nil {
		structMap["eddystone_uid_instance"] = w.EddystoneUidInstance
	}
	if w.EddystoneUidNamespace != nil {
		structMap["eddystone_uid_namespace"] = w.EddystoneUidNamespace
	}
	if w.EddystoneUrlUrl != nil {
		structMap["eddystone_url_url"] = w.EddystoneUrlUrl
	}
	if w.IbeaconMajor != nil {
		structMap["ibeacon_major"] = w.IbeaconMajor
	}
	if w.IbeaconMinor != nil {
		structMap["ibeacon_minor"] = w.IbeaconMinor
	}
	if w.IbeaconUuid != nil {
		structMap["ibeacon_uuid"] = w.IbeaconUuid
	}
	structMap["id"] = w.Id
	if w.Mac != nil {
		structMap["mac"] = w.Mac
	}
	structMap["map_id"] = w.MapId
	if w.MfgCompanyId != nil {
		structMap["mfg_company_id"] = w.MfgCompanyId
	}
	if w.MfgData != nil {
		structMap["mfg_data"] = w.MfgData
	}
	if w.Name != nil {
		structMap["name"] = w.Name
	}
	structMap["site_id"] = w.SiteId
	structMap["timestamp"] = w.Timestamp
	structMap["type"] = w.Type
	if w.WifiBeaconExtendedInfo != nil {
		structMap["wifi_beacon_extended_info"] = w.WifiBeaconExtendedInfo
	}
	structMap["x"] = w.X
	structMap["y"] = w.Y
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookLocationEvent.
// It customizes the JSON unmarshaling process for WebhookLocationEvent objects.
func (w *WebhookLocationEvent) UnmarshalJSON(input []byte) error {
	var temp tempWebhookLocationEvent
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "battery_voltage", "eddystone_uid_instance", "eddystone_uid_namespace", "eddystone_url_url", "ibeacon_major", "ibeacon_minor", "ibeacon_uuid", "id", "mac", "map_id", "mfg_company_id", "mfg_data", "name", "site_id", "timestamp", "type", "wifi_beacon_extended_info", "x", "y")
	if err != nil {
		return err
	}
	w.AdditionalProperties = additionalProperties

	w.BatteryVoltage = temp.BatteryVoltage
	w.EddystoneUidInstance = temp.EddystoneUidInstance
	w.EddystoneUidNamespace = temp.EddystoneUidNamespace
	w.EddystoneUrlUrl = temp.EddystoneUrlUrl
	w.IbeaconMajor = temp.IbeaconMajor
	w.IbeaconMinor = temp.IbeaconMinor
	w.IbeaconUuid = temp.IbeaconUuid
	w.Id = *temp.Id
	w.Mac = temp.Mac
	w.MapId = *temp.MapId
	w.MfgCompanyId = temp.MfgCompanyId
	w.MfgData = temp.MfgData
	w.Name = temp.Name
	w.SiteId = *temp.SiteId
	w.Timestamp = *temp.Timestamp
	w.Type = *temp.Type
	w.WifiBeaconExtendedInfo = temp.WifiBeaconExtendedInfo
	w.X = *temp.X
	w.Y = *temp.Y
	return nil
}

// tempWebhookLocationEvent is a temporary struct used for validating the fields of WebhookLocationEvent.
type tempWebhookLocationEvent struct {
	BatteryVoltage         *int                          `json:"battery_voltage,omitempty"`
	EddystoneUidInstance   *string                       `json:"eddystone_uid_instance,omitempty"`
	EddystoneUidNamespace  *string                       `json:"eddystone_uid_namespace,omitempty"`
	EddystoneUrlUrl        *string                       `json:"eddystone_url_url,omitempty"`
	IbeaconMajor           *int                          `json:"ibeacon_major,omitempty"`
	IbeaconMinor           *int                          `json:"ibeacon_minor,omitempty"`
	IbeaconUuid            *uuid.UUID                    `json:"ibeacon_uuid,omitempty"`
	Id                     *uuid.UUID                    `json:"id"`
	Mac                    *string                       `json:"mac,omitempty"`
	MapId                  *uuid.UUID                    `json:"map_id"`
	MfgCompanyId           *int                          `json:"mfg_company_id,omitempty"`
	MfgData                *string                       `json:"mfg_data,omitempty"`
	Name                   *string                       `json:"name,omitempty"`
	SiteId                 *uuid.UUID                    `json:"site_id"`
	Timestamp              *float64                      `json:"timestamp"`
	Type                   *string                       `json:"type"`
	WifiBeaconExtendedInfo []WifiBeaconExtendedInfoItems `json:"wifi_beacon_extended_info,omitempty"`
	X                      *float64                      `json:"x"`
	Y                      *float64                      `json:"y"`
}

func (w *tempWebhookLocationEvent) validate() error {
	var errs []string
	if w.Id == nil {
		errs = append(errs, "required field `id` is missing for type `webhook_location_event`")
	}
	if w.MapId == nil {
		errs = append(errs, "required field `map_id` is missing for type `webhook_location_event`")
	}
	if w.SiteId == nil {
		errs = append(errs, "required field `site_id` is missing for type `webhook_location_event`")
	}
	if w.Timestamp == nil {
		errs = append(errs, "required field `timestamp` is missing for type `webhook_location_event`")
	}
	if w.Type == nil {
		errs = append(errs, "required field `type` is missing for type `webhook_location_event`")
	}
	if w.X == nil {
		errs = append(errs, "required field `x` is missing for type `webhook_location_event`")
	}
	if w.Y == nil {
		errs = append(errs, "required field `y` is missing for type `webhook_location_event`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
