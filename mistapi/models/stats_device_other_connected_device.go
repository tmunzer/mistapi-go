package models

import (
    "encoding/json"
    "fmt"
)

// StatsDeviceOtherConnectedDevice represents a StatsDeviceOtherConnectedDevice struct.
type StatsDeviceOtherConnectedDevice struct {
    Mac                  *string                `json:"mac,omitempty"`
    Name                 *string                `json:"name,omitempty"`
    PortId               *string                `json:"port_id,omitempty"`
    Type                 *string                `json:"type,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for StatsDeviceOtherConnectedDevice,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s StatsDeviceOtherConnectedDevice) String() string {
    return fmt.Sprintf(
    	"StatsDeviceOtherConnectedDevice[Mac=%v, Name=%v, PortId=%v, Type=%v, AdditionalProperties=%v]",
    	s.Mac, s.Name, s.PortId, s.Type, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for StatsDeviceOtherConnectedDevice.
// It customizes the JSON marshaling process for StatsDeviceOtherConnectedDevice objects.
func (s StatsDeviceOtherConnectedDevice) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "mac", "name", "port_id", "type"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the StatsDeviceOtherConnectedDevice object to a map representation for JSON marshaling.
func (s StatsDeviceOtherConnectedDevice) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Mac != nil {
        structMap["mac"] = s.Mac
    }
    if s.Name != nil {
        structMap["name"] = s.Name
    }
    if s.PortId != nil {
        structMap["port_id"] = s.PortId
    }
    if s.Type != nil {
        structMap["type"] = s.Type
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsDeviceOtherConnectedDevice.
// It customizes the JSON unmarshaling process for StatsDeviceOtherConnectedDevice objects.
func (s *StatsDeviceOtherConnectedDevice) UnmarshalJSON(input []byte) error {
    var temp tempStatsDeviceOtherConnectedDevice
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "mac", "name", "port_id", "type")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Mac = temp.Mac
    s.Name = temp.Name
    s.PortId = temp.PortId
    s.Type = temp.Type
    return nil
}

// tempStatsDeviceOtherConnectedDevice is a temporary struct used for validating the fields of StatsDeviceOtherConnectedDevice.
type tempStatsDeviceOtherConnectedDevice  struct {
    Mac    *string `json:"mac,omitempty"`
    Name   *string `json:"name,omitempty"`
    PortId *string `json:"port_id,omitempty"`
    Type   *string `json:"type,omitempty"`
}
