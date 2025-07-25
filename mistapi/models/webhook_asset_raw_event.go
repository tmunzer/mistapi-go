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

// WebhookAssetRawEvent represents a WebhookAssetRawEvent struct.
type WebhookAssetRawEvent struct {
    // Asset id
    AssetId               uuid.UUID              `json:"asset_id"`
    // Antenna index, from 1-8, clock-wise starting from the LED
    Beam                  int                    `json:"beam"`
    // Device where the asset reading is from
    DeviceId              uuid.UUID              `json:"device_id"`
    // iBeacon major
    IbeaconMajor          *int                   `json:"ibeacon_major,omitempty"`
    // iBeacon minor
    IbeaconMinor          *int                   `json:"ibeacon_minor,omitempty"`
    // iBeacon UUID
    IbeaconUuid           *uuid.UUID             `json:"ibeacon_uuid,omitempty"`
    // MAC of the beacon
    Mac                   string                 `json:"mac"`
    // Map id
    MapId                 uuid.UUID              `json:"map_id"`
    // Optional, BLE manufacturing company ID
    MfgCompanyId          float64                `json:"mfg_company_id"`
    // Optional, BLE manufacturing data in hex byte-string format (ie: "112233AABBCC")
    MfgData               string                 `json:"mfg_data"`
    // Signal strength
    Rssi                  float64                `json:"rssi"`
    // Optional, data from service data
    ServiceDataData       *string                `json:"service_data_data,omitempty"`
    // Optional, last data transmit time from service data
    ServiceDataLastRxTime *int                   `json:"service_data_last_rx_time,omitempty"`
    // Optional, data transmit count from service data
    ServiceDataRxCnt      *int                   `json:"service_data_rx_cnt,omitempty"`
    // Optional, UUID from service data
    ServiceDataUuid       *uuid.UUID             `json:"service_data_uuid,omitempty"`
    // List of service data packets heard from the asset/ beacon
    ServicePackets        []ServicePacket        `json:"service_packets,omitempty"`
    SiteId                uuid.UUID              `json:"site_id"`
    // Epoch (seconds)
    Timestamp             float64                `json:"timestamp"`
    AdditionalProperties  map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for WebhookAssetRawEvent,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WebhookAssetRawEvent) String() string {
    return fmt.Sprintf(
    	"WebhookAssetRawEvent[AssetId=%v, Beam=%v, DeviceId=%v, IbeaconMajor=%v, IbeaconMinor=%v, IbeaconUuid=%v, Mac=%v, MapId=%v, MfgCompanyId=%v, MfgData=%v, Rssi=%v, ServiceDataData=%v, ServiceDataLastRxTime=%v, ServiceDataRxCnt=%v, ServiceDataUuid=%v, ServicePackets=%v, SiteId=%v, Timestamp=%v, AdditionalProperties=%v]",
    	w.AssetId, w.Beam, w.DeviceId, w.IbeaconMajor, w.IbeaconMinor, w.IbeaconUuid, w.Mac, w.MapId, w.MfgCompanyId, w.MfgData, w.Rssi, w.ServiceDataData, w.ServiceDataLastRxTime, w.ServiceDataRxCnt, w.ServiceDataUuid, w.ServicePackets, w.SiteId, w.Timestamp, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WebhookAssetRawEvent.
// It customizes the JSON marshaling process for WebhookAssetRawEvent objects.
func (w WebhookAssetRawEvent) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(w.AdditionalProperties,
        "asset_id", "beam", "device_id", "ibeacon_major", "ibeacon_minor", "ibeacon_uuid", "mac", "map_id", "mfg_company_id", "mfg_data", "rssi", "service_data_data", "service_data_last_rx_time", "service_data_rx_cnt", "service_data_uuid", "service_packets", "site_id", "timestamp"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WebhookAssetRawEvent object to a map representation for JSON marshaling.
func (w WebhookAssetRawEvent) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, w.AdditionalProperties)
    structMap["asset_id"] = w.AssetId
    structMap["beam"] = w.Beam
    structMap["device_id"] = w.DeviceId
    if w.IbeaconMajor != nil {
        structMap["ibeacon_major"] = w.IbeaconMajor
    }
    if w.IbeaconMinor != nil {
        structMap["ibeacon_minor"] = w.IbeaconMinor
    }
    if w.IbeaconUuid != nil {
        structMap["ibeacon_uuid"] = w.IbeaconUuid
    }
    structMap["mac"] = w.Mac
    structMap["map_id"] = w.MapId
    structMap["mfg_company_id"] = w.MfgCompanyId
    structMap["mfg_data"] = w.MfgData
    structMap["rssi"] = w.Rssi
    if w.ServiceDataData != nil {
        structMap["service_data_data"] = w.ServiceDataData
    }
    if w.ServiceDataLastRxTime != nil {
        structMap["service_data_last_rx_time"] = w.ServiceDataLastRxTime
    }
    if w.ServiceDataRxCnt != nil {
        structMap["service_data_rx_cnt"] = w.ServiceDataRxCnt
    }
    if w.ServiceDataUuid != nil {
        structMap["service_data_uuid"] = w.ServiceDataUuid
    }
    if w.ServicePackets != nil {
        structMap["service_packets"] = w.ServicePackets
    }
    structMap["site_id"] = w.SiteId
    structMap["timestamp"] = w.Timestamp
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookAssetRawEvent.
// It customizes the JSON unmarshaling process for WebhookAssetRawEvent objects.
func (w *WebhookAssetRawEvent) UnmarshalJSON(input []byte) error {
    var temp tempWebhookAssetRawEvent
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "asset_id", "beam", "device_id", "ibeacon_major", "ibeacon_minor", "ibeacon_uuid", "mac", "map_id", "mfg_company_id", "mfg_data", "rssi", "service_data_data", "service_data_last_rx_time", "service_data_rx_cnt", "service_data_uuid", "service_packets", "site_id", "timestamp")
    if err != nil {
    	return err
    }
    w.AdditionalProperties = additionalProperties
    
    w.AssetId = *temp.AssetId
    w.Beam = *temp.Beam
    w.DeviceId = *temp.DeviceId
    w.IbeaconMajor = temp.IbeaconMajor
    w.IbeaconMinor = temp.IbeaconMinor
    w.IbeaconUuid = temp.IbeaconUuid
    w.Mac = *temp.Mac
    w.MapId = *temp.MapId
    w.MfgCompanyId = *temp.MfgCompanyId
    w.MfgData = *temp.MfgData
    w.Rssi = *temp.Rssi
    w.ServiceDataData = temp.ServiceDataData
    w.ServiceDataLastRxTime = temp.ServiceDataLastRxTime
    w.ServiceDataRxCnt = temp.ServiceDataRxCnt
    w.ServiceDataUuid = temp.ServiceDataUuid
    w.ServicePackets = temp.ServicePackets
    w.SiteId = *temp.SiteId
    w.Timestamp = *temp.Timestamp
    return nil
}

// tempWebhookAssetRawEvent is a temporary struct used for validating the fields of WebhookAssetRawEvent.
type tempWebhookAssetRawEvent  struct {
    AssetId               *uuid.UUID      `json:"asset_id"`
    Beam                  *int            `json:"beam"`
    DeviceId              *uuid.UUID      `json:"device_id"`
    IbeaconMajor          *int            `json:"ibeacon_major,omitempty"`
    IbeaconMinor          *int            `json:"ibeacon_minor,omitempty"`
    IbeaconUuid           *uuid.UUID      `json:"ibeacon_uuid,omitempty"`
    Mac                   *string         `json:"mac"`
    MapId                 *uuid.UUID      `json:"map_id"`
    MfgCompanyId          *float64        `json:"mfg_company_id"`
    MfgData               *string         `json:"mfg_data"`
    Rssi                  *float64        `json:"rssi"`
    ServiceDataData       *string         `json:"service_data_data,omitempty"`
    ServiceDataLastRxTime *int            `json:"service_data_last_rx_time,omitempty"`
    ServiceDataRxCnt      *int            `json:"service_data_rx_cnt,omitempty"`
    ServiceDataUuid       *uuid.UUID      `json:"service_data_uuid,omitempty"`
    ServicePackets        []ServicePacket `json:"service_packets,omitempty"`
    SiteId                *uuid.UUID      `json:"site_id"`
    Timestamp             *float64        `json:"timestamp"`
}

func (w *tempWebhookAssetRawEvent) validate() error {
    var errs []string
    if w.AssetId == nil {
        errs = append(errs, "required field `asset_id` is missing for type `webhook_asset_raw_event`")
    }
    if w.Beam == nil {
        errs = append(errs, "required field `beam` is missing for type `webhook_asset_raw_event`")
    }
    if w.DeviceId == nil {
        errs = append(errs, "required field `device_id` is missing for type `webhook_asset_raw_event`")
    }
    if w.Mac == nil {
        errs = append(errs, "required field `mac` is missing for type `webhook_asset_raw_event`")
    }
    if w.MapId == nil {
        errs = append(errs, "required field `map_id` is missing for type `webhook_asset_raw_event`")
    }
    if w.MfgCompanyId == nil {
        errs = append(errs, "required field `mfg_company_id` is missing for type `webhook_asset_raw_event`")
    }
    if w.MfgData == nil {
        errs = append(errs, "required field `mfg_data` is missing for type `webhook_asset_raw_event`")
    }
    if w.Rssi == nil {
        errs = append(errs, "required field `rssi` is missing for type `webhook_asset_raw_event`")
    }
    if w.SiteId == nil {
        errs = append(errs, "required field `site_id` is missing for type `webhook_asset_raw_event`")
    }
    if w.Timestamp == nil {
        errs = append(errs, "required field `timestamp` is missing for type `webhook_asset_raw_event`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
