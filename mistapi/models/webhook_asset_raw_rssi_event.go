// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
)

// WebhookAssetRawRssiEvent represents a WebhookAssetRawRssiEvent struct.
type WebhookAssetRawRssiEvent struct {
	// optional, coordinates (if any) of reporting AP (updated once in 60s per client)
	ApLoc []float64 `json:"ap_loc,omitempty"`
	// antenna index, clock-wise starting from the LED
	Beam *int `json:"beam,omitempty"`
	// Unique ID of the object instance in the Mist Organization
	DeviceId     *uuid.UUID          `json:"device_id,omitempty"`
	IbeaconMajor Optional[int]       `json:"ibeacon_major"`
	IbeaconMinor Optional[int]       `json:"ibeacon_minor"`
	IbeaconUuid  Optional[uuid.UUID] `json:"ibeacon_uuid"`
	IsAsset      *bool               `json:"is_asset,omitempty"`
	// MAC of the asset/ beacon
	Mac *string `json:"mac,omitempty"`
	// Unique ID of the object instance in the Mist Organization
	MapId *uuid.UUID `json:"map_id,omitempty"`
	// optional, BLE manufacturing company ID
	MfgCompanyId Optional[int] `json:"mfg_company_id"`
	// optional, BLE manufacturing data in hex byte-string format (ie “112233AABBCC”)
	MfgData        Optional[string]                        `json:"mfg_data"`
	OrgId          *uuid.UUID                              `json:"org_id,omitempty"`
	Rssi           *int                                    `json:"rssi,omitempty"`
	ServicePackets []WebhookAssetRawRssiEventServicePacket `json:"service_packets,omitempty"`
	SiteId         *uuid.UUID                              `json:"site_id,omitempty"`
	// Epoch (seconds)
	Timestamp            *float64               `json:"timestamp,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for WebhookAssetRawRssiEvent,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WebhookAssetRawRssiEvent) String() string {
	return fmt.Sprintf(
		"WebhookAssetRawRssiEvent[ApLoc=%v, Beam=%v, DeviceId=%v, IbeaconMajor=%v, IbeaconMinor=%v, IbeaconUuid=%v, IsAsset=%v, Mac=%v, MapId=%v, MfgCompanyId=%v, MfgData=%v, OrgId=%v, Rssi=%v, ServicePackets=%v, SiteId=%v, Timestamp=%v, AdditionalProperties=%v]",
		w.ApLoc, w.Beam, w.DeviceId, w.IbeaconMajor, w.IbeaconMinor, w.IbeaconUuid, w.IsAsset, w.Mac, w.MapId, w.MfgCompanyId, w.MfgData, w.OrgId, w.Rssi, w.ServicePackets, w.SiteId, w.Timestamp, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WebhookAssetRawRssiEvent.
// It customizes the JSON marshaling process for WebhookAssetRawRssiEvent objects.
func (w WebhookAssetRawRssiEvent) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(w.AdditionalProperties,
		"ap_loc", "beam", "device_id", "ibeacon_major", "ibeacon_minor", "ibeacon_uuid", "is_asset", "mac", "map_id", "mfg_company_id", "mfg_data", "org_id", "rssi", "service_packets", "site_id", "timestamp"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(w.toMap())
}

// toMap converts the WebhookAssetRawRssiEvent object to a map representation for JSON marshaling.
func (w WebhookAssetRawRssiEvent) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, w.AdditionalProperties)
	if w.ApLoc != nil {
		structMap["ap_loc"] = w.ApLoc
	}
	if w.Beam != nil {
		structMap["beam"] = w.Beam
	}
	if w.DeviceId != nil {
		structMap["device_id"] = w.DeviceId
	}
	if w.IbeaconMajor.IsValueSet() {
		if w.IbeaconMajor.Value() != nil {
			structMap["ibeacon_major"] = w.IbeaconMajor.Value()
		} else {
			structMap["ibeacon_major"] = nil
		}
	}
	if w.IbeaconMinor.IsValueSet() {
		if w.IbeaconMinor.Value() != nil {
			structMap["ibeacon_minor"] = w.IbeaconMinor.Value()
		} else {
			structMap["ibeacon_minor"] = nil
		}
	}
	if w.IbeaconUuid.IsValueSet() {
		if w.IbeaconUuid.Value() != nil {
			structMap["ibeacon_uuid"] = w.IbeaconUuid.Value()
		} else {
			structMap["ibeacon_uuid"] = nil
		}
	}
	if w.IsAsset != nil {
		structMap["is_asset"] = w.IsAsset
	}
	if w.Mac != nil {
		structMap["mac"] = w.Mac
	}
	if w.MapId != nil {
		structMap["map_id"] = w.MapId
	}
	if w.MfgCompanyId.IsValueSet() {
		if w.MfgCompanyId.Value() != nil {
			structMap["mfg_company_id"] = w.MfgCompanyId.Value()
		} else {
			structMap["mfg_company_id"] = nil
		}
	}
	if w.MfgData.IsValueSet() {
		if w.MfgData.Value() != nil {
			structMap["mfg_data"] = w.MfgData.Value()
		} else {
			structMap["mfg_data"] = nil
		}
	}
	if w.OrgId != nil {
		structMap["org_id"] = w.OrgId
	}
	if w.Rssi != nil {
		structMap["rssi"] = w.Rssi
	}
	if w.ServicePackets != nil {
		structMap["service_packets"] = w.ServicePackets
	}
	if w.SiteId != nil {
		structMap["site_id"] = w.SiteId
	}
	if w.Timestamp != nil {
		structMap["timestamp"] = w.Timestamp
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookAssetRawRssiEvent.
// It customizes the JSON unmarshaling process for WebhookAssetRawRssiEvent objects.
func (w *WebhookAssetRawRssiEvent) UnmarshalJSON(input []byte) error {
	var temp tempWebhookAssetRawRssiEvent
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "ap_loc", "beam", "device_id", "ibeacon_major", "ibeacon_minor", "ibeacon_uuid", "is_asset", "mac", "map_id", "mfg_company_id", "mfg_data", "org_id", "rssi", "service_packets", "site_id", "timestamp")
	if err != nil {
		return err
	}
	w.AdditionalProperties = additionalProperties

	w.ApLoc = temp.ApLoc
	w.Beam = temp.Beam
	w.DeviceId = temp.DeviceId
	w.IbeaconMajor = temp.IbeaconMajor
	w.IbeaconMinor = temp.IbeaconMinor
	w.IbeaconUuid = temp.IbeaconUuid
	w.IsAsset = temp.IsAsset
	w.Mac = temp.Mac
	w.MapId = temp.MapId
	w.MfgCompanyId = temp.MfgCompanyId
	w.MfgData = temp.MfgData
	w.OrgId = temp.OrgId
	w.Rssi = temp.Rssi
	w.ServicePackets = temp.ServicePackets
	w.SiteId = temp.SiteId
	w.Timestamp = temp.Timestamp
	return nil
}

// tempWebhookAssetRawRssiEvent is a temporary struct used for validating the fields of WebhookAssetRawRssiEvent.
type tempWebhookAssetRawRssiEvent struct {
	ApLoc          []float64                               `json:"ap_loc,omitempty"`
	Beam           *int                                    `json:"beam,omitempty"`
	DeviceId       *uuid.UUID                              `json:"device_id,omitempty"`
	IbeaconMajor   Optional[int]                           `json:"ibeacon_major"`
	IbeaconMinor   Optional[int]                           `json:"ibeacon_minor"`
	IbeaconUuid    Optional[uuid.UUID]                     `json:"ibeacon_uuid"`
	IsAsset        *bool                                   `json:"is_asset,omitempty"`
	Mac            *string                                 `json:"mac,omitempty"`
	MapId          *uuid.UUID                              `json:"map_id,omitempty"`
	MfgCompanyId   Optional[int]                           `json:"mfg_company_id"`
	MfgData        Optional[string]                        `json:"mfg_data"`
	OrgId          *uuid.UUID                              `json:"org_id,omitempty"`
	Rssi           *int                                    `json:"rssi,omitempty"`
	ServicePackets []WebhookAssetRawRssiEventServicePacket `json:"service_packets,omitempty"`
	SiteId         *uuid.UUID                              `json:"site_id,omitempty"`
	Timestamp      *float64                                `json:"timestamp,omitempty"`
}
