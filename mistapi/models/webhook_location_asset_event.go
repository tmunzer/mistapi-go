package models

import (
	"encoding/json"
	"github.com/google/uuid"
)

// WebhookLocationAssetEvent represents a WebhookLocationAssetEvent struct.
type WebhookLocationAssetEvent struct {
	BatteryVoltage        *int       `json:"battery_voltage,omitempty"`
	EddystoneUidInstance  *string    `json:"eddystone_uid_instance,omitempty"`
	EddystoneUidNamespace *string    `json:"eddystone_uid_namespace,omitempty"`
	EddystoneUrlUrl       *string    `json:"eddystone_url_url,omitempty"`
	IbeaconMajor          *int       `json:"ibeacon_major,omitempty"`
	IbeaconMinor          *int       `json:"ibeacon_minor,omitempty"`
	IbeaconUuid           *uuid.UUID `json:"ibeacon_uuid,omitempty"`
	Mac                   *string    `json:"mac,omitempty"`
	MapId                 *uuid.UUID `json:"map_id,omitempty"`
	// optional, BLE manufacturing company ID
	MfgCompanyId *int `json:"mfg_company_id,omitempty"`
	// optional, BLE manufacturing data in hex byte-string format (ie: “112233AABBCC”)
	MfgData   *string    `json:"mfg_data,omitempty"`
	SiteId    *uuid.UUID `json:"site_id,omitempty"`
	Timestamp *int       `json:"timestamp,omitempty"`
	Type      *string    `json:"type,omitempty"`
	// x, in meter
	X *float64 `json:"x,omitempty"`
	// y, in meter
	Y                    *float64       `json:"y,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for WebhookLocationAssetEvent.
// It customizes the JSON marshaling process for WebhookLocationAssetEvent objects.
func (w WebhookLocationAssetEvent) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(w.toMap())
}

// toMap converts the WebhookLocationAssetEvent object to a map representation for JSON marshaling.
func (w WebhookLocationAssetEvent) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, w.AdditionalProperties)
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
	if w.X != nil {
		structMap["x"] = w.X
	}
	if w.Y != nil {
		structMap["y"] = w.Y
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookLocationAssetEvent.
// It customizes the JSON unmarshaling process for WebhookLocationAssetEvent objects.
func (w *WebhookLocationAssetEvent) UnmarshalJSON(input []byte) error {
	var temp tempWebhookLocationAssetEvent
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "battery_voltage", "eddystone_uid_instance", "eddystone_uid_namespace", "eddystone_url_url", "ibeacon_major", "ibeacon_minor", "ibeacon_uuid", "mac", "map_id", "mfg_company_id", "mfg_data", "site_id", "timestamp", "type", "x", "y")
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
	w.Mac = temp.Mac
	w.MapId = temp.MapId
	w.MfgCompanyId = temp.MfgCompanyId
	w.MfgData = temp.MfgData
	w.SiteId = temp.SiteId
	w.Timestamp = temp.Timestamp
	w.Type = temp.Type
	w.X = temp.X
	w.Y = temp.Y
	return nil
}

// tempWebhookLocationAssetEvent is a temporary struct used for validating the fields of WebhookLocationAssetEvent.
type tempWebhookLocationAssetEvent struct {
	BatteryVoltage        *int       `json:"battery_voltage,omitempty"`
	EddystoneUidInstance  *string    `json:"eddystone_uid_instance,omitempty"`
	EddystoneUidNamespace *string    `json:"eddystone_uid_namespace,omitempty"`
	EddystoneUrlUrl       *string    `json:"eddystone_url_url,omitempty"`
	IbeaconMajor          *int       `json:"ibeacon_major,omitempty"`
	IbeaconMinor          *int       `json:"ibeacon_minor,omitempty"`
	IbeaconUuid           *uuid.UUID `json:"ibeacon_uuid,omitempty"`
	Mac                   *string    `json:"mac,omitempty"`
	MapId                 *uuid.UUID `json:"map_id,omitempty"`
	MfgCompanyId          *int       `json:"mfg_company_id,omitempty"`
	MfgData               *string    `json:"mfg_data,omitempty"`
	SiteId                *uuid.UUID `json:"site_id,omitempty"`
	Timestamp             *int       `json:"timestamp,omitempty"`
	Type                  *string    `json:"type,omitempty"`
	X                     *float64   `json:"x,omitempty"`
	Y                     *float64   `json:"y,omitempty"`
}
