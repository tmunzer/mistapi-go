// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// SwitchIotPort represents a SwitchIotPort struct.
// Switch IOT port configuration
type SwitchIotPort struct {
    // Alarm class for the switch iot port in. enum: `minor`, `major`
    AlarmClass           *SwitchIotPortAlarmClassEnum `json:"alarm_class,omitempty"`
    Enabled              *bool                        `json:"enabled,omitempty"`
    // Only for "OUT" ports, input source for the switch iot port out. enum: `IN0`, `IN1`
    InputSrc             *SwitchIotPortInputSrcEnum   `json:"input_src,omitempty"`
    Name                 *string                      `json:"name,omitempty"`
    AdditionalProperties map[string]interface{}       `json:"_"`
}

// String implements the fmt.Stringer interface for SwitchIotPort,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SwitchIotPort) String() string {
    return fmt.Sprintf(
    	"SwitchIotPort[AlarmClass=%v, Enabled=%v, InputSrc=%v, Name=%v, AdditionalProperties=%v]",
    	s.AlarmClass, s.Enabled, s.InputSrc, s.Name, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SwitchIotPort.
// It customizes the JSON marshaling process for SwitchIotPort objects.
func (s SwitchIotPort) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "alarm_class", "enabled", "input_src", "name"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SwitchIotPort object to a map representation for JSON marshaling.
func (s SwitchIotPort) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.AlarmClass != nil {
        structMap["alarm_class"] = s.AlarmClass
    }
    if s.Enabled != nil {
        structMap["enabled"] = s.Enabled
    }
    if s.InputSrc != nil {
        structMap["input_src"] = s.InputSrc
    }
    if s.Name != nil {
        structMap["name"] = s.Name
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SwitchIotPort.
// It customizes the JSON unmarshaling process for SwitchIotPort objects.
func (s *SwitchIotPort) UnmarshalJSON(input []byte) error {
    var temp tempSwitchIotPort
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "alarm_class", "enabled", "input_src", "name")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.AlarmClass = temp.AlarmClass
    s.Enabled = temp.Enabled
    s.InputSrc = temp.InputSrc
    s.Name = temp.Name
    return nil
}

// tempSwitchIotPort is a temporary struct used for validating the fields of SwitchIotPort.
type tempSwitchIotPort  struct {
    AlarmClass *SwitchIotPortAlarmClassEnum `json:"alarm_class,omitempty"`
    Enabled    *bool                        `json:"enabled,omitempty"`
    InputSrc   *SwitchIotPortInputSrcEnum   `json:"input_src,omitempty"`
    Name       *string                      `json:"name,omitempty"`
}
