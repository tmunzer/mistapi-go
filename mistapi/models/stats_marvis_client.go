// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
)

// StatsMarvisClient represents a StatsMarvisClient struct.
// Marvis Client stats record returned by search
type StatsMarvisClient struct {
	// Whether the device battery is currently charging
	BatteryCharging *bool `json:"battery_charging,omitempty"`
	// Battery level percentage (0–100)
	BatteryLevel *int `json:"battery_level,omitempty"`
	// Background CPU utilization (0–100)
	CpuBackground *float64 `json:"cpu_background,omitempty"`
	// Idle CPU percentage (0–100)
	CpuIdle *float64 `json:"cpu_idle,omitempty"`
	// System CPU utilization (0–100)
	CpuSystem *float64 `json:"cpu_system,omitempty"`
	// User-space CPU utilization (0–100)
	CpuUser *float64 `json:"cpu_user,omitempty"`
	// UUID of the device the Marvis Client is installed on
	DeviceId *uuid.UUID `json:"device_id,omitempty"`
	// Device hostname
	Hostname *string `json:"hostname,omitempty"`
	// Last known location fix for a Marvis Client device
	Location *StatsMarvisClientLocation `json:"location,omitempty"`
	// Total device memory, in bytes
	MemoryTotal *int `json:"memory_total,omitempty"`
	// Memory in use, in bytes
	MemoryUsage *int `json:"memory_usage,omitempty"`
	// Device manufacturer
	Mfg *string `json:"mfg,omitempty"`
	// Device model name
	Model *string `json:"model,omitempty"`
	// Organization UUID
	OrgId *uuid.UUID `json:"org_id,omitempty"`
	// OS type or platform (e.g. Android, iOS)
	OsType *string `json:"os_type,omitempty"`
	// OS version string
	OsVersion *string `json:"os_version,omitempty"`
	// Device serial number
	Serial *string `json:"serial,omitempty"`
	// Total device storage, in bytes
	StorageTotal *int `json:"storage_total,omitempty"`
	// Storage in use, in bytes
	StorageUsage *int `json:"storage_usage,omitempty"`
	// Timestamp of the stats record, in epoch seconds
	Timestamp *int `json:"timestamp,omitempty"`
	// Wi-Fi band the device is connected on
	WifiBand *string `json:"wifi_band,omitempty"`
	// BSSID the device is connected to
	WifiBssid *string `json:"wifi_bssid,omitempty"`
	// Wi-Fi channel the device is on
	WifiChannel *int `json:"wifi_channel,omitempty"`
	// Device Wi-Fi IP address
	WifiIp *string `json:"wifi_ip,omitempty"`
	// Device Wi-Fi MAC address
	WifiMac *string `json:"wifi_mac,omitempty"`
	// Wi-Fi RSSI, in dBm
	WifiRssi *int `json:"wifi_rssi,omitempty"`
	// SSID the device is connected to
	WifiSsid             *string                `json:"wifi_ssid,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for StatsMarvisClient,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s StatsMarvisClient) String() string {
	return fmt.Sprintf(
		"StatsMarvisClient[BatteryCharging=%v, BatteryLevel=%v, CpuBackground=%v, CpuIdle=%v, CpuSystem=%v, CpuUser=%v, DeviceId=%v, Hostname=%v, Location=%v, MemoryTotal=%v, MemoryUsage=%v, Mfg=%v, Model=%v, OrgId=%v, OsType=%v, OsVersion=%v, Serial=%v, StorageTotal=%v, StorageUsage=%v, Timestamp=%v, WifiBand=%v, WifiBssid=%v, WifiChannel=%v, WifiIp=%v, WifiMac=%v, WifiRssi=%v, WifiSsid=%v, AdditionalProperties=%v]",
		s.BatteryCharging, s.BatteryLevel, s.CpuBackground, s.CpuIdle, s.CpuSystem, s.CpuUser, s.DeviceId, s.Hostname, s.Location, s.MemoryTotal, s.MemoryUsage, s.Mfg, s.Model, s.OrgId, s.OsType, s.OsVersion, s.Serial, s.StorageTotal, s.StorageUsage, s.Timestamp, s.WifiBand, s.WifiBssid, s.WifiChannel, s.WifiIp, s.WifiMac, s.WifiRssi, s.WifiSsid, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for StatsMarvisClient.
// It customizes the JSON marshaling process for StatsMarvisClient objects.
func (s StatsMarvisClient) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"battery_charging", "battery_level", "cpu_background", "cpu_idle", "cpu_system", "cpu_user", "device_id", "hostname", "location", "memory_total", "memory_usage", "mfg", "model", "org_id", "os_type", "os_version", "serial", "storage_total", "storage_usage", "timestamp", "wifi_band", "wifi_bssid", "wifi_channel", "wifi_ip", "wifi_mac", "wifi_rssi", "wifi_ssid"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the StatsMarvisClient object to a map representation for JSON marshaling.
func (s StatsMarvisClient) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.BatteryCharging != nil {
		structMap["battery_charging"] = s.BatteryCharging
	}
	if s.BatteryLevel != nil {
		structMap["battery_level"] = s.BatteryLevel
	}
	if s.CpuBackground != nil {
		structMap["cpu_background"] = s.CpuBackground
	}
	if s.CpuIdle != nil {
		structMap["cpu_idle"] = s.CpuIdle
	}
	if s.CpuSystem != nil {
		structMap["cpu_system"] = s.CpuSystem
	}
	if s.CpuUser != nil {
		structMap["cpu_user"] = s.CpuUser
	}
	if s.DeviceId != nil {
		structMap["device_id"] = s.DeviceId
	}
	if s.Hostname != nil {
		structMap["hostname"] = s.Hostname
	}
	if s.Location != nil {
		structMap["location"] = s.Location.toMap()
	}
	if s.MemoryTotal != nil {
		structMap["memory_total"] = s.MemoryTotal
	}
	if s.MemoryUsage != nil {
		structMap["memory_usage"] = s.MemoryUsage
	}
	if s.Mfg != nil {
		structMap["mfg"] = s.Mfg
	}
	if s.Model != nil {
		structMap["model"] = s.Model
	}
	if s.OrgId != nil {
		structMap["org_id"] = s.OrgId
	}
	if s.OsType != nil {
		structMap["os_type"] = s.OsType
	}
	if s.OsVersion != nil {
		structMap["os_version"] = s.OsVersion
	}
	if s.Serial != nil {
		structMap["serial"] = s.Serial
	}
	if s.StorageTotal != nil {
		structMap["storage_total"] = s.StorageTotal
	}
	if s.StorageUsage != nil {
		structMap["storage_usage"] = s.StorageUsage
	}
	if s.Timestamp != nil {
		structMap["timestamp"] = s.Timestamp
	}
	if s.WifiBand != nil {
		structMap["wifi_band"] = s.WifiBand
	}
	if s.WifiBssid != nil {
		structMap["wifi_bssid"] = s.WifiBssid
	}
	if s.WifiChannel != nil {
		structMap["wifi_channel"] = s.WifiChannel
	}
	if s.WifiIp != nil {
		structMap["wifi_ip"] = s.WifiIp
	}
	if s.WifiMac != nil {
		structMap["wifi_mac"] = s.WifiMac
	}
	if s.WifiRssi != nil {
		structMap["wifi_rssi"] = s.WifiRssi
	}
	if s.WifiSsid != nil {
		structMap["wifi_ssid"] = s.WifiSsid
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsMarvisClient.
// It customizes the JSON unmarshaling process for StatsMarvisClient objects.
func (s *StatsMarvisClient) UnmarshalJSON(input []byte) error {
	var temp tempStatsMarvisClient
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "battery_charging", "battery_level", "cpu_background", "cpu_idle", "cpu_system", "cpu_user", "device_id", "hostname", "location", "memory_total", "memory_usage", "mfg", "model", "org_id", "os_type", "os_version", "serial", "storage_total", "storage_usage", "timestamp", "wifi_band", "wifi_bssid", "wifi_channel", "wifi_ip", "wifi_mac", "wifi_rssi", "wifi_ssid")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.BatteryCharging = temp.BatteryCharging
	s.BatteryLevel = temp.BatteryLevel
	s.CpuBackground = temp.CpuBackground
	s.CpuIdle = temp.CpuIdle
	s.CpuSystem = temp.CpuSystem
	s.CpuUser = temp.CpuUser
	s.DeviceId = temp.DeviceId
	s.Hostname = temp.Hostname
	s.Location = temp.Location
	s.MemoryTotal = temp.MemoryTotal
	s.MemoryUsage = temp.MemoryUsage
	s.Mfg = temp.Mfg
	s.Model = temp.Model
	s.OrgId = temp.OrgId
	s.OsType = temp.OsType
	s.OsVersion = temp.OsVersion
	s.Serial = temp.Serial
	s.StorageTotal = temp.StorageTotal
	s.StorageUsage = temp.StorageUsage
	s.Timestamp = temp.Timestamp
	s.WifiBand = temp.WifiBand
	s.WifiBssid = temp.WifiBssid
	s.WifiChannel = temp.WifiChannel
	s.WifiIp = temp.WifiIp
	s.WifiMac = temp.WifiMac
	s.WifiRssi = temp.WifiRssi
	s.WifiSsid = temp.WifiSsid
	return nil
}

// tempStatsMarvisClient is a temporary struct used for validating the fields of StatsMarvisClient.
type tempStatsMarvisClient struct {
	BatteryCharging *bool                      `json:"battery_charging,omitempty"`
	BatteryLevel    *int                       `json:"battery_level,omitempty"`
	CpuBackground   *float64                   `json:"cpu_background,omitempty"`
	CpuIdle         *float64                   `json:"cpu_idle,omitempty"`
	CpuSystem       *float64                   `json:"cpu_system,omitempty"`
	CpuUser         *float64                   `json:"cpu_user,omitempty"`
	DeviceId        *uuid.UUID                 `json:"device_id,omitempty"`
	Hostname        *string                    `json:"hostname,omitempty"`
	Location        *StatsMarvisClientLocation `json:"location,omitempty"`
	MemoryTotal     *int                       `json:"memory_total,omitempty"`
	MemoryUsage     *int                       `json:"memory_usage,omitempty"`
	Mfg             *string                    `json:"mfg,omitempty"`
	Model           *string                    `json:"model,omitempty"`
	OrgId           *uuid.UUID                 `json:"org_id,omitempty"`
	OsType          *string                    `json:"os_type,omitempty"`
	OsVersion       *string                    `json:"os_version,omitempty"`
	Serial          *string                    `json:"serial,omitempty"`
	StorageTotal    *int                       `json:"storage_total,omitempty"`
	StorageUsage    *int                       `json:"storage_usage,omitempty"`
	Timestamp       *int                       `json:"timestamp,omitempty"`
	WifiBand        *string                    `json:"wifi_band,omitempty"`
	WifiBssid       *string                    `json:"wifi_bssid,omitempty"`
	WifiChannel     *int                       `json:"wifi_channel,omitempty"`
	WifiIp          *string                    `json:"wifi_ip,omitempty"`
	WifiMac         *string                    `json:"wifi_mac,omitempty"`
	WifiRssi        *int                       `json:"wifi_rssi,omitempty"`
	WifiSsid        *string                    `json:"wifi_ssid,omitempty"`
}
