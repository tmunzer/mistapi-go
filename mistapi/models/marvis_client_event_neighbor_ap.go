// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// MarvisClientEventNeighborAp represents a MarvisClientEventNeighborAp struct.
// A neighboring AP observed in a Marvis Client event
type MarvisClientEventNeighborAp struct {
	// Wi-Fi band the AP is operating on
	Band *string `json:"band,omitempty"`
	// BSSID of the neighboring AP
	Bssid *string `json:"bssid,omitempty"`
	// Channel the neighboring AP is on
	Channel *int `json:"channel,omitempty"`
	// RSSI of the neighboring AP signal, in dBm
	Rssi                 *int                   `json:"rssi,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for MarvisClientEventNeighborAp,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m MarvisClientEventNeighborAp) String() string {
	return fmt.Sprintf(
		"MarvisClientEventNeighborAp[Band=%v, Bssid=%v, Channel=%v, Rssi=%v, AdditionalProperties=%v]",
		m.Band, m.Bssid, m.Channel, m.Rssi, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for MarvisClientEventNeighborAp.
// It customizes the JSON marshaling process for MarvisClientEventNeighborAp objects.
func (m MarvisClientEventNeighborAp) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(m.AdditionalProperties,
		"band", "bssid", "channel", "rssi"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(m.toMap())
}

// toMap converts the MarvisClientEventNeighborAp object to a map representation for JSON marshaling.
func (m MarvisClientEventNeighborAp) toMap() map[string]any {
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
	if m.Rssi != nil {
		structMap["rssi"] = m.Rssi
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MarvisClientEventNeighborAp.
// It customizes the JSON unmarshaling process for MarvisClientEventNeighborAp objects.
func (m *MarvisClientEventNeighborAp) UnmarshalJSON(input []byte) error {
	var temp tempMarvisClientEventNeighborAp
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "band", "bssid", "channel", "rssi")
	if err != nil {
		return err
	}
	m.AdditionalProperties = additionalProperties

	m.Band = temp.Band
	m.Bssid = temp.Bssid
	m.Channel = temp.Channel
	m.Rssi = temp.Rssi
	return nil
}

// tempMarvisClientEventNeighborAp is a temporary struct used for validating the fields of MarvisClientEventNeighborAp.
type tempMarvisClientEventNeighborAp struct {
	Band    *string `json:"band,omitempty"`
	Bssid   *string `json:"bssid,omitempty"`
	Channel *int    `json:"channel,omitempty"`
	Rssi    *int    `json:"rssi,omitempty"`
}
