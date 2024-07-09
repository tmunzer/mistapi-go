package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// WebhookDiscoveredRawRssiEvent represents a WebhookDiscoveredRawRssiEvent struct.
type WebhookDiscoveredRawRssiEvent struct {
    // coordinates (if any) of reporting AP (updated once in 60s per client)
    ApLoc                []float64       `json:"ap_loc,omitempty"`
    // antenna index, from 1-8, clock-wise starting from the LED
    Beam                 int             `json:"beam"`
    // device id of the reporting AP
    DeviceId             uuid.UUID       `json:"device_id"`
    IbeaconMajor         *int            `json:"ibeacon_major,omitempty"`
    IbeaconMinor         *int            `json:"ibeacon_minor,omitempty"`
    IbeaconUuid          *uuid.UUID      `json:"ibeacon_uuid,omitempty"`
    IsAsset              *bool           `json:"is_asset,omitempty"`
    // MAC of the asset/ beacon
    Mac                  string          `json:"mac"`
    MapId                uuid.UUID       `json:"map_id"`
    // BLE manufacturing company ID
    MfgCompanyId         *string         `json:"mfg_company_id,omitempty"`
    // BLE manufacturing data in hex byte-string format (ie: “112233AABBCC”)
    MfgData              *string         `json:"mfg_data,omitempty"`
    OrgId                uuid.UUID       `json:"org_id"`
    // signal strength
    Rssi                 float64         `json:"rssi"`
    // list of service data packets heard from the asset/ beacon
    ServicePackets       []ServicePacket `json:"service_packets,omitempty"`
    SiteId               uuid.UUID       `json:"site_id"`
    Timestamp            *int            `json:"timestamp,omitempty"`
    AdditionalProperties map[string]any  `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for WebhookDiscoveredRawRssiEvent.
// It customizes the JSON marshaling process for WebhookDiscoveredRawRssiEvent objects.
func (w WebhookDiscoveredRawRssiEvent) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(w.toMap())
}

// toMap converts the WebhookDiscoveredRawRssiEvent object to a map representation for JSON marshaling.
func (w WebhookDiscoveredRawRssiEvent) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, w.AdditionalProperties)
    if w.ApLoc != nil {
        structMap["ap_loc"] = w.ApLoc
    }
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
    if w.IsAsset != nil {
        structMap["is_asset"] = w.IsAsset
    }
    structMap["mac"] = w.Mac
    structMap["map_id"] = w.MapId
    if w.MfgCompanyId != nil {
        structMap["mfg_company_id"] = w.MfgCompanyId
    }
    if w.MfgData != nil {
        structMap["mfg_data"] = w.MfgData
    }
    structMap["org_id"] = w.OrgId
    structMap["rssi"] = w.Rssi
    if w.ServicePackets != nil {
        structMap["service_packets"] = w.ServicePackets
    }
    structMap["site_id"] = w.SiteId
    if w.Timestamp != nil {
        structMap["timestamp"] = w.Timestamp
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookDiscoveredRawRssiEvent.
// It customizes the JSON unmarshaling process for WebhookDiscoveredRawRssiEvent objects.
func (w *WebhookDiscoveredRawRssiEvent) UnmarshalJSON(input []byte) error {
    var temp webhookDiscoveredRawRssiEvent
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "ap_loc", "beam", "device_id", "ibeacon_major", "ibeacon_minor", "ibeacon_uuid", "is_asset", "mac", "map_id", "mfg_company_id", "mfg_data", "org_id", "rssi", "service_packets", "site_id", "timestamp")
    if err != nil {
    	return err
    }
    
    w.AdditionalProperties = additionalProperties
    w.ApLoc = temp.ApLoc
    w.Beam = *temp.Beam
    w.DeviceId = *temp.DeviceId
    w.IbeaconMajor = temp.IbeaconMajor
    w.IbeaconMinor = temp.IbeaconMinor
    w.IbeaconUuid = temp.IbeaconUuid
    w.IsAsset = temp.IsAsset
    w.Mac = *temp.Mac
    w.MapId = *temp.MapId
    w.MfgCompanyId = temp.MfgCompanyId
    w.MfgData = temp.MfgData
    w.OrgId = *temp.OrgId
    w.Rssi = *temp.Rssi
    w.ServicePackets = temp.ServicePackets
    w.SiteId = *temp.SiteId
    w.Timestamp = temp.Timestamp
    return nil
}

// webhookDiscoveredRawRssiEvent is a temporary struct used for validating the fields of WebhookDiscoveredRawRssiEvent.
type webhookDiscoveredRawRssiEvent  struct {
    ApLoc          []float64       `json:"ap_loc,omitempty"`
    Beam           *int            `json:"beam"`
    DeviceId       *uuid.UUID      `json:"device_id"`
    IbeaconMajor   *int            `json:"ibeacon_major,omitempty"`
    IbeaconMinor   *int            `json:"ibeacon_minor,omitempty"`
    IbeaconUuid    *uuid.UUID      `json:"ibeacon_uuid,omitempty"`
    IsAsset        *bool           `json:"is_asset,omitempty"`
    Mac            *string         `json:"mac"`
    MapId          *uuid.UUID      `json:"map_id"`
    MfgCompanyId   *string         `json:"mfg_company_id,omitempty"`
    MfgData        *string         `json:"mfg_data,omitempty"`
    OrgId          *uuid.UUID      `json:"org_id"`
    Rssi           *float64        `json:"rssi"`
    ServicePackets []ServicePacket `json:"service_packets,omitempty"`
    SiteId         *uuid.UUID      `json:"site_id"`
    Timestamp      *int            `json:"timestamp,omitempty"`
}

func (w *webhookDiscoveredRawRssiEvent) validate() error {
    var errs []string
    if w.Beam == nil {
        errs = append(errs, "required field `beam` is missing for type `Webhook_Discovered_Raw_Rssi_Event`")
    }
    if w.DeviceId == nil {
        errs = append(errs, "required field `device_id` is missing for type `Webhook_Discovered_Raw_Rssi_Event`")
    }
    if w.Mac == nil {
        errs = append(errs, "required field `mac` is missing for type `Webhook_Discovered_Raw_Rssi_Event`")
    }
    if w.MapId == nil {
        errs = append(errs, "required field `map_id` is missing for type `Webhook_Discovered_Raw_Rssi_Event`")
    }
    if w.OrgId == nil {
        errs = append(errs, "required field `org_id` is missing for type `Webhook_Discovered_Raw_Rssi_Event`")
    }
    if w.Rssi == nil {
        errs = append(errs, "required field `rssi` is missing for type `Webhook_Discovered_Raw_Rssi_Event`")
    }
    if w.SiteId == nil {
        errs = append(errs, "required field `site_id` is missing for type `Webhook_Discovered_Raw_Rssi_Event`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}