// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// IotendpointStats represents a IotendpointStats struct.
// IoT endpoint statistics
type IotendpointStats struct {
	// MAC address of the AP the endpoint was seen on
	ApMac *string `json:"ap_mac,omitempty"`
	// Unique identifier for the IoT endpoint
	Id *string `json:"id,omitempty"`
	// Link Quality Indicator (0–255)
	Lqi *int `json:"lqi,omitempty"`
	// MAC address of the IoT endpoint
	Mac *string `json:"mac,omitempty"`
	// Manufacturer name
	Mfg *string `json:"mfg,omitempty"`
	// Model name
	Model *string `json:"model,omitempty"`
	// Epoch timestamp of the last observation, in seconds
	Timestamp *float64 `json:"timestamp,omitempty"`
	// IoT endpoint type. enum: `zigbee`
	Type                 *string                `json:"type,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for IotendpointStats,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (i IotendpointStats) String() string {
	return fmt.Sprintf(
		"IotendpointStats[ApMac=%v, Id=%v, Lqi=%v, Mac=%v, Mfg=%v, Model=%v, Timestamp=%v, Type=%v, AdditionalProperties=%v]",
		i.ApMac, i.Id, i.Lqi, i.Mac, i.Mfg, i.Model, i.Timestamp, i.Type, i.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for IotendpointStats.
// It customizes the JSON marshaling process for IotendpointStats objects.
func (i IotendpointStats) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(i.AdditionalProperties,
		"ap_mac", "id", "lqi", "mac", "mfg", "model", "timestamp", "type"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(i.toMap())
}

// toMap converts the IotendpointStats object to a map representation for JSON marshaling.
func (i IotendpointStats) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, i.AdditionalProperties)
	if i.ApMac != nil {
		structMap["ap_mac"] = i.ApMac
	}
	if i.Id != nil {
		structMap["id"] = i.Id
	}
	if i.Lqi != nil {
		structMap["lqi"] = i.Lqi
	}
	if i.Mac != nil {
		structMap["mac"] = i.Mac
	}
	if i.Mfg != nil {
		structMap["mfg"] = i.Mfg
	}
	if i.Model != nil {
		structMap["model"] = i.Model
	}
	if i.Timestamp != nil {
		structMap["timestamp"] = i.Timestamp
	}
	if i.Type != nil {
		structMap["type"] = i.Type
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for IotendpointStats.
// It customizes the JSON unmarshaling process for IotendpointStats objects.
func (i *IotendpointStats) UnmarshalJSON(input []byte) error {
	var temp tempIotendpointStats
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "ap_mac", "id", "lqi", "mac", "mfg", "model", "timestamp", "type")
	if err != nil {
		return err
	}
	i.AdditionalProperties = additionalProperties

	i.ApMac = temp.ApMac
	i.Id = temp.Id
	i.Lqi = temp.Lqi
	i.Mac = temp.Mac
	i.Mfg = temp.Mfg
	i.Model = temp.Model
	i.Timestamp = temp.Timestamp
	i.Type = temp.Type
	return nil
}

// tempIotendpointStats is a temporary struct used for validating the fields of IotendpointStats.
type tempIotendpointStats struct {
	ApMac     *string  `json:"ap_mac,omitempty"`
	Id        *string  `json:"id,omitempty"`
	Lqi       *int     `json:"lqi,omitempty"`
	Mac       *string  `json:"mac,omitempty"`
	Mfg       *string  `json:"mfg,omitempty"`
	Model     *string  `json:"model,omitempty"`
	Timestamp *float64 `json:"timestamp,omitempty"`
	Type      *string  `json:"type,omitempty"`
}
