package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// CaptureSwitch represents a CaptureSwitch struct.
// Initiate a Switch (Junos) Packet Capture
type CaptureSwitch struct {
    // duration of the capture, in seconds
    Duration             *int                                           `json:"duration,omitempty"`
    // enum: `stream`
    Format               *CaptureSwitchFormatEnum                       `json:"format,omitempty"`
    // max_len of each packet to capture
    MaxPktLen            *int                                           `json:"max_pkt_len,omitempty"`
    // number of packets to capture, 0 for unlimited
    NumPackets           *int                                           `json:"num_packets,omitempty"`
    // Property key is the port name. 6 ports max per switch supported, or 5 max with irb port auto-included into capture request
    Ports                map[string]CaptureSwitchPortsTcpdumpExpression `json:"ports,omitempty"`
    // Property key is the switch mac
    Switches             map[string]CaptureSwitchSwitches               `json:"switches,omitempty"`
    // tcpdump expression, port specific if specified under ports dict, otherwise applicable across ports if specified at top level of payload. Port specific value overrides top level value when both exist.
    TcpdumpExpression    *string                                        `json:"tcpdump_expression,omitempty"`
    // enum: `switch`
    Type                 string                                         `json:"type"`
    AdditionalProperties map[string]any                                 `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CaptureSwitch.
// It customizes the JSON marshaling process for CaptureSwitch objects.
func (c CaptureSwitch) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CaptureSwitch object to a map representation for JSON marshaling.
func (c CaptureSwitch) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    if c.Duration != nil {
        structMap["duration"] = c.Duration
    }
    if c.Format != nil {
        structMap["format"] = c.Format
    }
    if c.MaxPktLen != nil {
        structMap["max_pkt_len"] = c.MaxPktLen
    }
    if c.NumPackets != nil {
        structMap["num_packets"] = c.NumPackets
    }
    if c.Ports != nil {
        structMap["ports"] = c.Ports
    }
    if c.Switches != nil {
        structMap["switches"] = c.Switches
    }
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
    additionalProperties, err := UnmarshalAdditionalProperties(input, "duration", "format", "max_pkt_len", "num_packets", "ports", "switches", "tcpdump_expression", "type")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.Duration = temp.Duration
    c.Format = temp.Format
    c.MaxPktLen = temp.MaxPktLen
    c.NumPackets = temp.NumPackets
    c.Ports = temp.Ports
    c.Switches = temp.Switches
    c.TcpdumpExpression = temp.TcpdumpExpression
    c.Type = *temp.Type
    return nil
}

// tempCaptureSwitch is a temporary struct used for validating the fields of CaptureSwitch.
type tempCaptureSwitch  struct {
    Duration          *int                                           `json:"duration,omitempty"`
    Format            *CaptureSwitchFormatEnum                       `json:"format,omitempty"`
    MaxPktLen         *int                                           `json:"max_pkt_len,omitempty"`
    NumPackets        *int                                           `json:"num_packets,omitempty"`
    Ports             map[string]CaptureSwitchPortsTcpdumpExpression `json:"ports,omitempty"`
    Switches          map[string]CaptureSwitchSwitches               `json:"switches,omitempty"`
    TcpdumpExpression *string                                        `json:"tcpdump_expression,omitempty"`
    Type              *string                                        `json:"type"`
}

func (c *tempCaptureSwitch) validate() error {
    var errs []string
    if c.Type == nil {
        errs = append(errs, "required field `type` is missing for type `capture_switch`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
