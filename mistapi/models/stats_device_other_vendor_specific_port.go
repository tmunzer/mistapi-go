// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// StatsDeviceOtherVendorSpecificPort represents a StatsDeviceOtherVendorSpecificPort struct.
type StatsDeviceOtherVendorSpecificPort struct {
    BytesIn              *int64                 `json:"bytes_in,omitempty"`
    BytesOut             *int64                 `json:"bytes_out,omitempty"`
    Carrier              *string                `json:"carrier,omitempty"`
    Imei                 *string                `json:"imei,omitempty"`
    Imsi                 *string                `json:"imsi,omitempty"`
    Ip                   *string                `json:"ip,omitempty"`
    Link                 *bool                  `json:"link,omitempty"`
    Mode                 *string                `json:"mode,omitempty"`
    Rsrp                 *float64               `json:"rsrp,omitempty"`
    Rsrq                 *float64               `json:"rsrq,omitempty"`
    Rssi                 *int                   `json:"rssi,omitempty"`
    ServiceMode          *string                `json:"service_mode,omitempty"`
    Sinr                 *float64               `json:"sinr,omitempty"`
    State                *string                `json:"state,omitempty"`
    Type                 *string                `json:"type,omitempty"`
    Uptime               *int                   `json:"uptime,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for StatsDeviceOtherVendorSpecificPort,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s StatsDeviceOtherVendorSpecificPort) String() string {
    return fmt.Sprintf(
    	"StatsDeviceOtherVendorSpecificPort[BytesIn=%v, BytesOut=%v, Carrier=%v, Imei=%v, Imsi=%v, Ip=%v, Link=%v, Mode=%v, Rsrp=%v, Rsrq=%v, Rssi=%v, ServiceMode=%v, Sinr=%v, State=%v, Type=%v, Uptime=%v, AdditionalProperties=%v]",
    	s.BytesIn, s.BytesOut, s.Carrier, s.Imei, s.Imsi, s.Ip, s.Link, s.Mode, s.Rsrp, s.Rsrq, s.Rssi, s.ServiceMode, s.Sinr, s.State, s.Type, s.Uptime, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for StatsDeviceOtherVendorSpecificPort.
// It customizes the JSON marshaling process for StatsDeviceOtherVendorSpecificPort objects.
func (s StatsDeviceOtherVendorSpecificPort) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "bytes_in", "bytes_out", "carrier", "imei", "imsi", "ip", "link", "mode", "rsrp", "rsrq", "rssi", "service_mode", "sinr", "state", "type", "uptime"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the StatsDeviceOtherVendorSpecificPort object to a map representation for JSON marshaling.
func (s StatsDeviceOtherVendorSpecificPort) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.BytesIn != nil {
        structMap["bytes_in"] = s.BytesIn
    }
    if s.BytesOut != nil {
        structMap["bytes_out"] = s.BytesOut
    }
    if s.Carrier != nil {
        structMap["carrier"] = s.Carrier
    }
    if s.Imei != nil {
        structMap["imei"] = s.Imei
    }
    if s.Imsi != nil {
        structMap["imsi"] = s.Imsi
    }
    if s.Ip != nil {
        structMap["ip"] = s.Ip
    }
    if s.Link != nil {
        structMap["link"] = s.Link
    }
    if s.Mode != nil {
        structMap["mode"] = s.Mode
    }
    if s.Rsrp != nil {
        structMap["rsrp"] = s.Rsrp
    }
    if s.Rsrq != nil {
        structMap["rsrq"] = s.Rsrq
    }
    if s.Rssi != nil {
        structMap["rssi"] = s.Rssi
    }
    if s.ServiceMode != nil {
        structMap["service_mode"] = s.ServiceMode
    }
    if s.Sinr != nil {
        structMap["sinr"] = s.Sinr
    }
    if s.State != nil {
        structMap["state"] = s.State
    }
    if s.Type != nil {
        structMap["type"] = s.Type
    }
    if s.Uptime != nil {
        structMap["uptime"] = s.Uptime
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsDeviceOtherVendorSpecificPort.
// It customizes the JSON unmarshaling process for StatsDeviceOtherVendorSpecificPort objects.
func (s *StatsDeviceOtherVendorSpecificPort) UnmarshalJSON(input []byte) error {
    var temp tempStatsDeviceOtherVendorSpecificPort
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "bytes_in", "bytes_out", "carrier", "imei", "imsi", "ip", "link", "mode", "rsrp", "rsrq", "rssi", "service_mode", "sinr", "state", "type", "uptime")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.BytesIn = temp.BytesIn
    s.BytesOut = temp.BytesOut
    s.Carrier = temp.Carrier
    s.Imei = temp.Imei
    s.Imsi = temp.Imsi
    s.Ip = temp.Ip
    s.Link = temp.Link
    s.Mode = temp.Mode
    s.Rsrp = temp.Rsrp
    s.Rsrq = temp.Rsrq
    s.Rssi = temp.Rssi
    s.ServiceMode = temp.ServiceMode
    s.Sinr = temp.Sinr
    s.State = temp.State
    s.Type = temp.Type
    s.Uptime = temp.Uptime
    return nil
}

// tempStatsDeviceOtherVendorSpecificPort is a temporary struct used for validating the fields of StatsDeviceOtherVendorSpecificPort.
type tempStatsDeviceOtherVendorSpecificPort  struct {
    BytesIn     *int64   `json:"bytes_in,omitempty"`
    BytesOut    *int64   `json:"bytes_out,omitempty"`
    Carrier     *string  `json:"carrier,omitempty"`
    Imei        *string  `json:"imei,omitempty"`
    Imsi        *string  `json:"imsi,omitempty"`
    Ip          *string  `json:"ip,omitempty"`
    Link        *bool    `json:"link,omitempty"`
    Mode        *string  `json:"mode,omitempty"`
    Rsrp        *float64 `json:"rsrp,omitempty"`
    Rsrq        *float64 `json:"rsrq,omitempty"`
    Rssi        *int     `json:"rssi,omitempty"`
    ServiceMode *string  `json:"service_mode,omitempty"`
    Sinr        *float64 `json:"sinr,omitempty"`
    State       *string  `json:"state,omitempty"`
    Type        *string  `json:"type,omitempty"`
    Uptime      *int     `json:"uptime,omitempty"`
}
