// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
    "github.com/google/uuid"
)

// StatsRssiZonesDevice represents a StatsRssiZonesDevice struct.
type StatsRssiZonesDevice struct {
    DeviceId             *uuid.UUID             `json:"device_id,omitempty"`
    Rssi                 *int                   `json:"rssi,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for StatsRssiZonesDevice,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s StatsRssiZonesDevice) String() string {
    return fmt.Sprintf(
    	"StatsRssiZonesDevice[DeviceId=%v, Rssi=%v, AdditionalProperties=%v]",
    	s.DeviceId, s.Rssi, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for StatsRssiZonesDevice.
// It customizes the JSON marshaling process for StatsRssiZonesDevice objects.
func (s StatsRssiZonesDevice) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "device_id", "rssi"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the StatsRssiZonesDevice object to a map representation for JSON marshaling.
func (s StatsRssiZonesDevice) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.DeviceId != nil {
        structMap["device_id"] = s.DeviceId
    }
    if s.Rssi != nil {
        structMap["rssi"] = s.Rssi
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsRssiZonesDevice.
// It customizes the JSON unmarshaling process for StatsRssiZonesDevice objects.
func (s *StatsRssiZonesDevice) UnmarshalJSON(input []byte) error {
    var temp tempStatsRssiZonesDevice
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "device_id", "rssi")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.DeviceId = temp.DeviceId
    s.Rssi = temp.Rssi
    return nil
}

// tempStatsRssiZonesDevice is a temporary struct used for validating the fields of StatsRssiZonesDevice.
type tempStatsRssiZonesDevice  struct {
    DeviceId *uuid.UUID `json:"device_id,omitempty"`
    Rssi     *int       `json:"rssi,omitempty"`
}
