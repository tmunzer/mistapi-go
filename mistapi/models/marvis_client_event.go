// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
)

// MarvisClientEvent represents a MarvisClientEvent struct.
// A Marvis Client event record
type MarvisClientEvent struct {
	// Wi-Fi band at the time of the event
	Band *string `json:"band,omitempty"`
	// BSSID the client roamed to (for roam events)
	Bssid *string `json:"bssid,omitempty"`
	// Channel the client roamed to (for roam events)
	Channel *int `json:"channel,omitempty"`
	// UUID of the device the Marvis Client is installed on
	DeviceId *uuid.UUID `json:"device_id,omitempty"`
	// Device hostname
	Hostname *string `json:"hostname,omitempty"`
	// Last known location fix for a Marvis Client device
	Location *StatsMarvisClientLocation `json:"location,omitempty"`
	// List of neighboring APs observed at the time of the event
	NeighborApReport []MarvisClientEventNeighborAp `json:"neighbor_ap_report,omitempty"`
	// Organization UUID
	OrgId *uuid.UUID `json:"org_id,omitempty"`
	// Battery level percentage at the time of the event (for battery events)
	Percent *int `json:"percent,omitempty"`
	// BSSID the client roamed from (for roam events)
	PreBssid *string `json:"pre_bssid,omitempty"`
	// Channel the client roamed from (for roam events)
	PreChannel *int `json:"pre_channel,omitempty"`
	// RSSI before the roam event, in dBm
	PreRssi *int `json:"pre_rssi,omitempty"`
	// Wi-Fi RSSI at the time of the event, in dBm
	Rssi *int `json:"rssi,omitempty"`
	// SSID the client was connected to
	Ssid *string `json:"ssid,omitempty"`
	// Event timestamp, in epoch seconds
	Timestamp *int `json:"timestamp,omitempty"`
	// Event type
	Type *string `json:"type,omitempty"`
	// Device Wi-Fi IP address at the time of the event
	WifiIp *string `json:"wifi_ip,omitempty"`
	// Device Wi-Fi MAC address
	WifiMac              *string                `json:"wifi_mac,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for MarvisClientEvent,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m MarvisClientEvent) String() string {
	return fmt.Sprintf(
		"MarvisClientEvent[Band=%v, Bssid=%v, Channel=%v, DeviceId=%v, Hostname=%v, Location=%v, NeighborApReport=%v, OrgId=%v, Percent=%v, PreBssid=%v, PreChannel=%v, PreRssi=%v, Rssi=%v, Ssid=%v, Timestamp=%v, Type=%v, WifiIp=%v, WifiMac=%v, AdditionalProperties=%v]",
		m.Band, m.Bssid, m.Channel, m.DeviceId, m.Hostname, m.Location, m.NeighborApReport, m.OrgId, m.Percent, m.PreBssid, m.PreChannel, m.PreRssi, m.Rssi, m.Ssid, m.Timestamp, m.Type, m.WifiIp, m.WifiMac, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for MarvisClientEvent.
// It customizes the JSON marshaling process for MarvisClientEvent objects.
func (m MarvisClientEvent) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(m.AdditionalProperties,
		"band", "bssid", "channel", "device_id", "hostname", "location", "neighbor_ap_report", "org_id", "percent", "pre_bssid", "pre_channel", "pre_rssi", "rssi", "ssid", "timestamp", "type", "wifi_ip", "wifi_mac"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(m.toMap())
}

// toMap converts the MarvisClientEvent object to a map representation for JSON marshaling.
func (m MarvisClientEvent) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, m.AdditionalProperties)
	if m.Band != nil {
		structMap["band"] = m.Band
	}
	if m.Bssid != nil {
		structMap["bssid"] = m.Bssid
	}
	if m.Channel != nil {
		structMap["channel"] = m.Channel
	}
	if m.DeviceId != nil {
		structMap["device_id"] = m.DeviceId
	}
	if m.Hostname != nil {
		structMap["hostname"] = m.Hostname
	}
	if m.Location != nil {
		structMap["location"] = m.Location.toMap()
	}
	if m.NeighborApReport != nil {
		structMap["neighbor_ap_report"] = m.NeighborApReport
	}
	if m.OrgId != nil {
		structMap["org_id"] = m.OrgId
	}
	if m.Percent != nil {
		structMap["percent"] = m.Percent
	}
	if m.PreBssid != nil {
		structMap["pre_bssid"] = m.PreBssid
	}
	if m.PreChannel != nil {
		structMap["pre_channel"] = m.PreChannel
	}
	if m.PreRssi != nil {
		structMap["pre_rssi"] = m.PreRssi
	}
	if m.Rssi != nil {
		structMap["rssi"] = m.Rssi
	}
	if m.Ssid != nil {
		structMap["ssid"] = m.Ssid
	}
	if m.Timestamp != nil {
		structMap["timestamp"] = m.Timestamp
	}
	if m.Type != nil {
		structMap["type"] = m.Type
	}
	if m.WifiIp != nil {
		structMap["wifi_ip"] = m.WifiIp
	}
	if m.WifiMac != nil {
		structMap["wifi_mac"] = m.WifiMac
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MarvisClientEvent.
// It customizes the JSON unmarshaling process for MarvisClientEvent objects.
func (m *MarvisClientEvent) UnmarshalJSON(input []byte) error {
	var temp tempMarvisClientEvent
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "band", "bssid", "channel", "device_id", "hostname", "location", "neighbor_ap_report", "org_id", "percent", "pre_bssid", "pre_channel", "pre_rssi", "rssi", "ssid", "timestamp", "type", "wifi_ip", "wifi_mac")
	if err != nil {
		return err
	}
	m.AdditionalProperties = additionalProperties

	m.Band = temp.Band
	m.Bssid = temp.Bssid
	m.Channel = temp.Channel
	m.DeviceId = temp.DeviceId
	m.Hostname = temp.Hostname
	m.Location = temp.Location
	m.NeighborApReport = temp.NeighborApReport
	m.OrgId = temp.OrgId
	m.Percent = temp.Percent
	m.PreBssid = temp.PreBssid
	m.PreChannel = temp.PreChannel
	m.PreRssi = temp.PreRssi
	m.Rssi = temp.Rssi
	m.Ssid = temp.Ssid
	m.Timestamp = temp.Timestamp
	m.Type = temp.Type
	m.WifiIp = temp.WifiIp
	m.WifiMac = temp.WifiMac
	return nil
}

// tempMarvisClientEvent is a temporary struct used for validating the fields of MarvisClientEvent.
type tempMarvisClientEvent struct {
	Band             *string                       `json:"band,omitempty"`
	Bssid            *string                       `json:"bssid,omitempty"`
	Channel          *int                          `json:"channel,omitempty"`
	DeviceId         *uuid.UUID                    `json:"device_id,omitempty"`
	Hostname         *string                       `json:"hostname,omitempty"`
	Location         *StatsMarvisClientLocation    `json:"location,omitempty"`
	NeighborApReport []MarvisClientEventNeighborAp `json:"neighbor_ap_report,omitempty"`
	OrgId            *uuid.UUID                    `json:"org_id,omitempty"`
	Percent          *int                          `json:"percent,omitempty"`
	PreBssid         *string                       `json:"pre_bssid,omitempty"`
	PreChannel       *int                          `json:"pre_channel,omitempty"`
	PreRssi          *int                          `json:"pre_rssi,omitempty"`
	Rssi             *int                          `json:"rssi,omitempty"`
	Ssid             *string                       `json:"ssid,omitempty"`
	Timestamp        *int                          `json:"timestamp,omitempty"`
	Type             *string                       `json:"type,omitempty"`
	WifiIp           *string                       `json:"wifi_ip,omitempty"`
	WifiMac          *string                       `json:"wifi_mac,omitempty"`
}
