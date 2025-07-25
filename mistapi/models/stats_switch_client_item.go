// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// StatsSwitchClientItem represents a StatsSwitchClientItem struct.
type StatsSwitchClientItem struct {
    DeviceMac            *string                `json:"device_mac,omitempty"`
    Hostname             *string                `json:"hostname,omitempty"`
    Mac                  *string                `json:"mac,omitempty"`
    PortId               *string                `json:"port_id,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for StatsSwitchClientItem,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s StatsSwitchClientItem) String() string {
    return fmt.Sprintf(
    	"StatsSwitchClientItem[DeviceMac=%v, Hostname=%v, Mac=%v, PortId=%v, AdditionalProperties=%v]",
    	s.DeviceMac, s.Hostname, s.Mac, s.PortId, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for StatsSwitchClientItem.
// It customizes the JSON marshaling process for StatsSwitchClientItem objects.
func (s StatsSwitchClientItem) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "device_mac", "hostname", "mac", "port_id"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the StatsSwitchClientItem object to a map representation for JSON marshaling.
func (s StatsSwitchClientItem) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.DeviceMac != nil {
        structMap["device_mac"] = s.DeviceMac
    }
    if s.Hostname != nil {
        structMap["hostname"] = s.Hostname
    }
    if s.Mac != nil {
        structMap["mac"] = s.Mac
    }
    if s.PortId != nil {
        structMap["port_id"] = s.PortId
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsSwitchClientItem.
// It customizes the JSON unmarshaling process for StatsSwitchClientItem objects.
func (s *StatsSwitchClientItem) UnmarshalJSON(input []byte) error {
    var temp tempStatsSwitchClientItem
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "device_mac", "hostname", "mac", "port_id")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.DeviceMac = temp.DeviceMac
    s.Hostname = temp.Hostname
    s.Mac = temp.Mac
    s.PortId = temp.PortId
    return nil
}

// tempStatsSwitchClientItem is a temporary struct used for validating the fields of StatsSwitchClientItem.
type tempStatsSwitchClientItem  struct {
    DeviceMac *string `json:"device_mac,omitempty"`
    Hostname  *string `json:"hostname,omitempty"`
    Mac       *string `json:"mac,omitempty"`
    PortId    *string `json:"port_id,omitempty"`
}
