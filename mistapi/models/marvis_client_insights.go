// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// MarvisClientInsights represents a MarvisClientInsights struct.
// Time-series performance metrics for a Marvis Client device
type MarvisClientInsights struct {
	// Average battery level per interval bucket
	AvgBattery []float64 `json:"avg_battery,omitempty"`
	// Average cellular RSSI per interval bucket, in dBm
	AvgCellularRssi []float64 `json:"avg_cellular_rssi,omitempty"`
	// Average CPU utilization per interval bucket (0–100)
	AvgCpu []float64 `json:"avg_cpu,omitempty"`
	// Average memory utilization per interval bucket (0–100)
	AvgMemory []float64 `json:"avg_memory,omitempty"`
	// Average Wi-Fi RSSI per interval bucket, in dBm
	AvgWifiRssi []float64 `json:"avg_wifi_rssi,omitempty"`
	// End of the reporting window, in epoch seconds
	End *int `json:"end,omitempty"`
	// Duration of each interval bucket, in seconds
	Interval *int `json:"interval,omitempty"`
	// Maximum number of results requested
	Limit *int `json:"limit,omitempty"`
	// Current page number
	Page *int `json:"page,omitempty"`
	// List of ISO 8601 timestamp strings for each interval bucket
	Rt []string `json:"rt,omitempty"`
	// Start of the reporting window, in epoch seconds
	Start                *int                   `json:"start,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for MarvisClientInsights,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m MarvisClientInsights) String() string {
	return fmt.Sprintf(
		"MarvisClientInsights[AvgBattery=%v, AvgCellularRssi=%v, AvgCpu=%v, AvgMemory=%v, AvgWifiRssi=%v, End=%v, Interval=%v, Limit=%v, Page=%v, Rt=%v, Start=%v, AdditionalProperties=%v]",
		m.AvgBattery, m.AvgCellularRssi, m.AvgCpu, m.AvgMemory, m.AvgWifiRssi, m.End, m.Interval, m.Limit, m.Page, m.Rt, m.Start, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for MarvisClientInsights.
// It customizes the JSON marshaling process for MarvisClientInsights objects.
func (m MarvisClientInsights) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(m.AdditionalProperties,
		"avg_battery", "avg_cellular_rssi", "avg_cpu", "avg_memory", "avg_wifi_rssi", "end", "interval", "limit", "page", "rt", "start"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(m.toMap())
}

// toMap converts the MarvisClientInsights object to a map representation for JSON marshaling.
func (m MarvisClientInsights) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, m.AdditionalProperties)
	if m.AvgBattery != nil {
		structMap["avg_battery"] = m.AvgBattery
	}
	if m.AvgCellularRssi != nil {
		structMap["avg_cellular_rssi"] = m.AvgCellularRssi
	}
	if m.AvgCpu != nil {
		structMap["avg_cpu"] = m.AvgCpu
	}
	if m.AvgMemory != nil {
		structMap["avg_memory"] = m.AvgMemory
	}
	if m.AvgWifiRssi != nil {
		structMap["avg_wifi_rssi"] = m.AvgWifiRssi
	}
	if m.End != nil {
		structMap["end"] = m.End
	}
	if m.Interval != nil {
		structMap["interval"] = m.Interval
	}
	if m.Limit != nil {
		structMap["limit"] = m.Limit
	}
	if m.Page != nil {
		structMap["page"] = m.Page
	}
	if m.Rt != nil {
		structMap["rt"] = m.Rt
	}
	if m.Start != nil {
		structMap["start"] = m.Start
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MarvisClientInsights.
// It customizes the JSON unmarshaling process for MarvisClientInsights objects.
func (m *MarvisClientInsights) UnmarshalJSON(input []byte) error {
	var temp tempMarvisClientInsights
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "avg_battery", "avg_cellular_rssi", "avg_cpu", "avg_memory", "avg_wifi_rssi", "end", "interval", "limit", "page", "rt", "start")
	if err != nil {
		return err
	}
	m.AdditionalProperties = additionalProperties

	m.AvgBattery = temp.AvgBattery
	m.AvgCellularRssi = temp.AvgCellularRssi
	m.AvgCpu = temp.AvgCpu
	m.AvgMemory = temp.AvgMemory
	m.AvgWifiRssi = temp.AvgWifiRssi
	m.End = temp.End
	m.Interval = temp.Interval
	m.Limit = temp.Limit
	m.Page = temp.Page
	m.Rt = temp.Rt
	m.Start = temp.Start
	return nil
}

// tempMarvisClientInsights is a temporary struct used for validating the fields of MarvisClientInsights.
type tempMarvisClientInsights struct {
	AvgBattery      []float64 `json:"avg_battery,omitempty"`
	AvgCellularRssi []float64 `json:"avg_cellular_rssi,omitempty"`
	AvgCpu          []float64 `json:"avg_cpu,omitempty"`
	AvgMemory       []float64 `json:"avg_memory,omitempty"`
	AvgWifiRssi     []float64 `json:"avg_wifi_rssi,omitempty"`
	End             *int      `json:"end,omitempty"`
	Interval        *int      `json:"interval,omitempty"`
	Limit           *int      `json:"limit,omitempty"`
	Page            *int      `json:"page,omitempty"`
	Rt              []string  `json:"rt,omitempty"`
	Start           *int      `json:"start,omitempty"`
}
