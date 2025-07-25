// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// CaptureSwitch represents a CaptureSwitch struct.
// Initiate a Switch (Junos) Packet Capture
type CaptureSwitch struct {
    // Duration of the capture, in seconds
    Duration             Optional[int]                                  `json:"duration"`
    // enum: `stream`
    Format               *CaptureSwitchFormatEnum                       `json:"format,omitempty"`
    MaxPktLen            Optional[int]                                  `json:"max_pkt_len"`
    // number of packets to capture, 0 for unlimited, default is 1024, maximum is 10000
    NumPackets           Optional[int]                                  `json:"num_packets"`
    // Property key is the port name. 6 ports max per switch supported, or 5 max with irb port auto-included into capture request
    Ports                map[string]CaptureSwitchPortsTcpdumpExpression `json:"ports,omitempty"`
    // Property key is the switch mac
    Switches             map[string]CaptureSwitchSwitches               `json:"switches"`
    // tcpdump expression, port specific if specified under ports dict, otherwise applicable across ports if specified at top level of payload. Port specific value overrides top level value when both exist.
    TcpdumpExpression    *string                                        `json:"tcpdump_expression,omitempty"`
    // enum: `switch`
    Type                 string                                         `json:"type"`
    AdditionalProperties map[string]interface{}                         `json:"_"`
}

// String implements the fmt.Stringer interface for CaptureSwitch,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CaptureSwitch) String() string {
    return fmt.Sprintf(
    	"CaptureSwitch[Duration=%v, Format=%v, MaxPktLen=%v, NumPackets=%v, Ports=%v, Switches=%v, TcpdumpExpression=%v, Type=%v, AdditionalProperties=%v]",
    	c.Duration, c.Format, c.MaxPktLen, c.NumPackets, c.Ports, c.Switches, c.TcpdumpExpression, c.Type, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CaptureSwitch.
// It customizes the JSON marshaling process for CaptureSwitch objects.
func (c CaptureSwitch) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "duration", "format", "max_pkt_len", "num_packets", "ports", "switches", "tcpdump_expression", "type"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CaptureSwitch object to a map representation for JSON marshaling.
func (c CaptureSwitch) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    if c.Duration.IsValueSet() {
        if c.Duration.Value() != nil {
            structMap["duration"] = c.Duration.Value()
        } else {
            structMap["duration"] = nil
        }
    }
    if c.Format != nil {
        structMap["format"] = c.Format
    }
    if c.MaxPktLen.IsValueSet() {
        if c.MaxPktLen.Value() != nil {
            structMap["max_pkt_len"] = c.MaxPktLen.Value()
        } else {
            structMap["max_pkt_len"] = nil
        }
    }
    if c.NumPackets.IsValueSet() {
        if c.NumPackets.Value() != nil {
            structMap["num_packets"] = c.NumPackets.Value()
        } else {
            structMap["num_packets"] = nil
        }
    }
    if c.Ports != nil {
        structMap["ports"] = c.Ports
    }
    structMap["switches"] = c.Switches
    if c.TcpdumpExpression != nil {
        structMap["tcpdump_expression"] = c.TcpdumpExpression
    }
    structMap["type"] = c.Type
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CaptureSwitch.
// It customizes the JSON unmarshaling process for CaptureSwitch objects.
func (c *CaptureSwitch) UnmarshalJSON(input []byte) error {
    var temp tempCaptureSwitch
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "duration", "format", "max_pkt_len", "num_packets", "ports", "switches", "tcpdump_expression", "type")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.Duration = temp.Duration
    c.Format = temp.Format
    c.MaxPktLen = temp.MaxPktLen
    c.NumPackets = temp.NumPackets
    c.Ports = temp.Ports
    c.Switches = *temp.Switches
    c.TcpdumpExpression = temp.TcpdumpExpression
    c.Type = *temp.Type
    return nil
}

// tempCaptureSwitch is a temporary struct used for validating the fields of CaptureSwitch.
type tempCaptureSwitch  struct {
    Duration          Optional[int]                                  `json:"duration"`
    Format            *CaptureSwitchFormatEnum                       `json:"format,omitempty"`
    MaxPktLen         Optional[int]                                  `json:"max_pkt_len"`
    NumPackets        Optional[int]                                  `json:"num_packets"`
    Ports             map[string]CaptureSwitchPortsTcpdumpExpression `json:"ports,omitempty"`
    Switches          *map[string]CaptureSwitchSwitches              `json:"switches"`
    TcpdumpExpression *string                                        `json:"tcpdump_expression,omitempty"`
    Type              *string                                        `json:"type"`
}

func (c *tempCaptureSwitch) validate() error {
    var errs []string
    if c.Switches == nil {
        errs = append(errs, "required field `switches` is missing for type `capture_switch`")
    }
    if c.Type == nil {
        errs = append(errs, "required field `type` is missing for type `capture_switch`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
