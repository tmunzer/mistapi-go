package models

import (
    "encoding/json"
)

// SwitchStatsClientItem represents a SwitchStatsClientItem struct.
type SwitchStatsClientItem struct {
    DeviceMac            *string        `json:"device_mac,omitempty"`
    Hostname             *string        `json:"hostname,omitempty"`
    Mac                  *string        `json:"mac,omitempty"`
    PortId               *string        `json:"port_id,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SwitchStatsClientItem.
// It customizes the JSON marshaling process for SwitchStatsClientItem objects.
func (s SwitchStatsClientItem) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SwitchStatsClientItem object to a map representation for JSON marshaling.
func (s SwitchStatsClientItem) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
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

// UnmarshalJSON implements the json.Unmarshaler interface for SwitchStatsClientItem.
// It customizes the JSON unmarshaling process for SwitchStatsClientItem objects.
func (s *SwitchStatsClientItem) UnmarshalJSON(input []byte) error {
    var temp tempSwitchStatsClientItem
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "device_mac", "hostname", "mac", "port_id")
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

// tempSwitchStatsClientItem is a temporary struct used for validating the fields of SwitchStatsClientItem.
type tempSwitchStatsClientItem  struct {
    DeviceMac *string `json:"device_mac,omitempty"`
    Hostname  *string `json:"hostname,omitempty"`
    Mac       *string `json:"mac,omitempty"`
    PortId    *string `json:"port_id,omitempty"`
}
