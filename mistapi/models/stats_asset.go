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

// StatsAsset represents a StatsAsset struct.
// Asset statistics
type StatsAsset struct {
	// Time-to-live in seconds; how long this asset data is valid in cache
	Ttl *int `json:"_ttl,omitempty"`
	// Battery voltage, in mV
	BatteryVoltage *float64 `json:"battery_voltage,omitempty"`
	Beam           *int     `json:"beam,omitempty"`
	// Source type
	By *string `json:"by,omitempty"`
	// Device ID of the loudest AP
	DeviceId              *uuid.UUID `json:"device_id,omitempty"`
	DeviceName            *string    `json:"device_name,omitempty"`
	Duration              *int       `json:"duration,omitempty"`
	EddystoneUidInstance  *string    `json:"eddystone_uid_instance,omitempty"`
	EddystoneUidNamespace *string    `json:"eddystone_uid_namespace,omitempty"`
	EddystoneUrlUrl       *string    `json:"eddystone_url_url,omitempty"`
	// Major number for iBeacon
	IbeaconMajor Optional[int] `json:"ibeacon_major"`
	// Minor number for iBeacon
	IbeaconMinor Optional[int]       `json:"ibeacon_minor"`
	IbeaconUuid  Optional[uuid.UUID] `json:"ibeacon_uuid"`
	// Unique ID of the object instance in the Mist Organization
	Id *uuid.UUID `json:"id,omitempty"`
	// Last seen timestamp
	LastSeen Optional[float64] `json:"last_seen"`
	// Bluetooth MAC
	Mac string `json:"mac"`
	// Manufacturer name resolved from company ID
	Manufacture *string `json:"manufacture,omitempty"`
	// Map where the device belongs to
	MapId *uuid.UUID `json:"map_id,omitempty"`
	// BLE manufacturer company ID from advertisement
	MfgCompanyId *int `json:"mfg_company_id,omitempty"`
	// Manufacturer-specific data (hex encoded)
	MfgData *string `json:"mfg_data,omitempty"`
	// Name / label of the device
	Name *string `json:"name,omitempty"`
	// Signal strength (RSSI) of the loudest AP in dBm
	Rssi *int `json:"rssi,omitempty"`
	// Only send this for individual asset stat
	Rssizones []AssetRssiZone `json:"rssizones,omitempty"`
	// List of all service data advertisements (maximum length of 10)
	ServicePackets []StatsAssetServicePacket `json:"service_packets,omitempty"`
	Temperature    *float64                  `json:"temperature,omitempty"`
	// X in pixel
	X *float64 `json:"x,omitempty"`
	// Y in pixel
	Y *float64 `json:"y,omitempty"`
	// Only send this for individual asset stat
	Zones                []AssetZone            `json:"zones,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for StatsAsset,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s StatsAsset) String() string {
	return fmt.Sprintf(
		"StatsAsset[Ttl=%v, BatteryVoltage=%v, Beam=%v, By=%v, DeviceId=%v, DeviceName=%v, Duration=%v, EddystoneUidInstance=%v, EddystoneUidNamespace=%v, EddystoneUrlUrl=%v, IbeaconMajor=%v, IbeaconMinor=%v, IbeaconUuid=%v, Id=%v, LastSeen=%v, Mac=%v, Manufacture=%v, MapId=%v, MfgCompanyId=%v, MfgData=%v, Name=%v, Rssi=%v, Rssizones=%v, ServicePackets=%v, Temperature=%v, X=%v, Y=%v, Zones=%v, AdditionalProperties=%v]",
		s.Ttl, s.BatteryVoltage, s.Beam, s.By, s.DeviceId, s.DeviceName, s.Duration, s.EddystoneUidInstance, s.EddystoneUidNamespace, s.EddystoneUrlUrl, s.IbeaconMajor, s.IbeaconMinor, s.IbeaconUuid, s.Id, s.LastSeen, s.Mac, s.Manufacture, s.MapId, s.MfgCompanyId, s.MfgData, s.Name, s.Rssi, s.Rssizones, s.ServicePackets, s.Temperature, s.X, s.Y, s.Zones, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for StatsAsset.
// It customizes the JSON marshaling process for StatsAsset objects.
func (s StatsAsset) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"_ttl", "battery_voltage", "beam", "by", "device_id", "device_name", "duration", "eddystone_uid_instance", "eddystone_uid_namespace", "eddystone_url_url", "ibeacon_major", "ibeacon_minor", "ibeacon_uuid", "id", "last_seen", "mac", "manufacture", "map_id", "mfg_company_id", "mfg_data", "name", "rssi", "rssizones", "service_packets", "temperature", "x", "y", "zones"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the StatsAsset object to a map representation for JSON marshaling.
func (s StatsAsset) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.Ttl != nil {
		structMap["_ttl"] = s.Ttl
	}
	if s.BatteryVoltage != nil {
		structMap["battery_voltage"] = s.BatteryVoltage
	}
	if s.Beam != nil {
		structMap["beam"] = s.Beam
	}
	if s.By != nil {
		structMap["by"] = s.By
	}
	if s.DeviceId != nil {
		structMap["device_id"] = s.DeviceId
	}
	if s.DeviceName != nil {
		structMap["device_name"] = s.DeviceName
	}
	if s.Duration != nil {
		structMap["duration"] = s.Duration
	}
	if s.EddystoneUidInstance != nil {
		structMap["eddystone_uid_instance"] = s.EddystoneUidInstance
	}
	if s.EddystoneUidNamespace != nil {
		structMap["eddystone_uid_namespace"] = s.EddystoneUidNamespace
	}
	if s.EddystoneUrlUrl != nil {
		structMap["eddystone_url_url"] = s.EddystoneUrlUrl
	}
	if s.IbeaconMajor.IsValueSet() {
		if s.IbeaconMajor.Value() != nil {
			structMap["ibeacon_major"] = s.IbeaconMajor.Value()
		} else {
			structMap["ibeacon_major"] = nil
		}
	}
	if s.IbeaconMinor.IsValueSet() {
		if s.IbeaconMinor.Value() != nil {
			structMap["ibeacon_minor"] = s.IbeaconMinor.Value()
		} else {
			structMap["ibeacon_minor"] = nil
		}
	}
	if s.IbeaconUuid.IsValueSet() {
		if s.IbeaconUuid.Value() != nil {
			structMap["ibeacon_uuid"] = s.IbeaconUuid.Value()
		} else {
			structMap["ibeacon_uuid"] = nil
		}
	}
	if s.Id != nil {
		structMap["id"] = s.Id
	}
	if s.LastSeen.IsValueSet() {
		if s.LastSeen.Value() != nil {
			structMap["last_seen"] = s.LastSeen.Value()
		} else {
			structMap["last_seen"] = nil
		}
	}
	structMap["mac"] = s.Mac
	if s.Manufacture != nil {
		structMap["manufacture"] = s.Manufacture
	}
	if s.MapId != nil {
		structMap["map_id"] = s.MapId
	}
	if s.MfgCompanyId != nil {
		structMap["mfg_company_id"] = s.MfgCompanyId
	}
	if s.MfgData != nil {
		structMap["mfg_data"] = s.MfgData
	}
	if s.Name != nil {
		structMap["name"] = s.Name
	}
	if s.Rssi != nil {
		structMap["rssi"] = s.Rssi
	}
	if s.Rssizones != nil {
		structMap["rssizones"] = s.Rssizones
	}
	if s.ServicePackets != nil {
		structMap["service_packets"] = s.ServicePackets
	}
	if s.Temperature != nil {
		structMap["temperature"] = s.Temperature
	}
	if s.X != nil {
		structMap["x"] = s.X
	}
	if s.Y != nil {
		structMap["y"] = s.Y
	}
	if s.Zones != nil {
		structMap["zones"] = s.Zones
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsAsset.
// It customizes the JSON unmarshaling process for StatsAsset objects.
func (s *StatsAsset) UnmarshalJSON(input []byte) error {
	var temp tempStatsAsset
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "_ttl", "battery_voltage", "beam", "by", "device_id", "device_name", "duration", "eddystone_uid_instance", "eddystone_uid_namespace", "eddystone_url_url", "ibeacon_major", "ibeacon_minor", "ibeacon_uuid", "id", "last_seen", "mac", "manufacture", "map_id", "mfg_company_id", "mfg_data", "name", "rssi", "rssizones", "service_packets", "temperature", "x", "y", "zones")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.Ttl = temp.Ttl
	s.BatteryVoltage = temp.BatteryVoltage
	s.Beam = temp.Beam
	s.By = temp.By
	s.DeviceId = temp.DeviceId
	s.DeviceName = temp.DeviceName
	s.Duration = temp.Duration
	s.EddystoneUidInstance = temp.EddystoneUidInstance
	s.EddystoneUidNamespace = temp.EddystoneUidNamespace
	s.EddystoneUrlUrl = temp.EddystoneUrlUrl
	s.IbeaconMajor = temp.IbeaconMajor
	s.IbeaconMinor = temp.IbeaconMinor
	s.IbeaconUuid = temp.IbeaconUuid
	s.Id = temp.Id
	s.LastSeen = temp.LastSeen
	s.Mac = *temp.Mac
	s.Manufacture = temp.Manufacture
	s.MapId = temp.MapId
	s.MfgCompanyId = temp.MfgCompanyId
	s.MfgData = temp.MfgData
	s.Name = temp.Name
	s.Rssi = temp.Rssi
	s.Rssizones = temp.Rssizones
	s.ServicePackets = temp.ServicePackets
	s.Temperature = temp.Temperature
	s.X = temp.X
	s.Y = temp.Y
	s.Zones = temp.Zones
	return nil
}

// tempStatsAsset is a temporary struct used for validating the fields of StatsAsset.
type tempStatsAsset struct {
	Ttl                   *int                      `json:"_ttl,omitempty"`
	BatteryVoltage        *float64                  `json:"battery_voltage,omitempty"`
	Beam                  *int                      `json:"beam,omitempty"`
	By                    *string                   `json:"by,omitempty"`
	DeviceId              *uuid.UUID                `json:"device_id,omitempty"`
	DeviceName            *string                   `json:"device_name,omitempty"`
	Duration              *int                      `json:"duration,omitempty"`
	EddystoneUidInstance  *string                   `json:"eddystone_uid_instance,omitempty"`
	EddystoneUidNamespace *string                   `json:"eddystone_uid_namespace,omitempty"`
	EddystoneUrlUrl       *string                   `json:"eddystone_url_url,omitempty"`
	IbeaconMajor          Optional[int]             `json:"ibeacon_major"`
	IbeaconMinor          Optional[int]             `json:"ibeacon_minor"`
	IbeaconUuid           Optional[uuid.UUID]       `json:"ibeacon_uuid"`
	Id                    *uuid.UUID                `json:"id,omitempty"`
	LastSeen              Optional[float64]         `json:"last_seen"`
	Mac                   *string                   `json:"mac"`
	Manufacture           *string                   `json:"manufacture,omitempty"`
	MapId                 *uuid.UUID                `json:"map_id,omitempty"`
	MfgCompanyId          *int                      `json:"mfg_company_id,omitempty"`
	MfgData               *string                   `json:"mfg_data,omitempty"`
	Name                  *string                   `json:"name,omitempty"`
	Rssi                  *int                      `json:"rssi,omitempty"`
	Rssizones             []AssetRssiZone           `json:"rssizones,omitempty"`
	ServicePackets        []StatsAssetServicePacket `json:"service_packets,omitempty"`
	Temperature           *float64                  `json:"temperature,omitempty"`
	X                     *float64                  `json:"x,omitempty"`
	Y                     *float64                  `json:"y,omitempty"`
	Zones                 []AssetZone               `json:"zones,omitempty"`
}

func (s *tempStatsAsset) validate() error {
	var errs []string
	if s.Mac == nil {
		errs = append(errs, "required field `mac` is missing for type `stats_asset`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
