// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// CaptureGateway represents a CaptureGateway struct.
// Initiate a Gateway (SSR/SRX) Packet Capture
type CaptureGateway struct {
    // Duration of the capture, in seconds
    Duration             Optional[int]                         `json:"duration"`
    // enum: `stream`
    Format               *CaptureGatewayFormatEnum             `json:"format,omitempty"`
    // List of SSRs. Property key is the SSR MAC
    Gateways             map[string]CaptureGatewayGateways     `json:"gateways"`
    // minimum is 64 (SSR) / 68 (SRX) maximum is 10240 (SSR) / 1520 (SRX)
    MaxPktLen            Optional[int]                         `json:"max_pkt_len"`
    // number of packets to capture, 0 for unlimited, default is 1024, maximum is 10000
    NumPackets           Optional[int]                         `json:"num_packets"`
    // Property key is the port ID
    Ports                map[string]CaptureGatewayGatewaysPort `json:"ports,omitempty"`
    // enum: `gateway`
    Type                 string                                `json:"type"`
    AdditionalProperties map[string]interface{}                `json:"_"`
}

// String implements the fmt.Stringer interface for CaptureGateway,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CaptureGateway) String() string {
    return fmt.Sprintf(
    	"CaptureGateway[Duration=%v, Format=%v, Gateways=%v, MaxPktLen=%v, NumPackets=%v, Ports=%v, Type=%v, AdditionalProperties=%v]",
    	c.Duration, c.Format, c.Gateways, c.MaxPktLen, c.NumPackets, c.Ports, c.Type, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CaptureGateway.
// It customizes the JSON marshaling process for CaptureGateway objects.
func (c CaptureGateway) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "duration", "format", "gateways", "max_pkt_len", "num_packets", "ports", "type"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CaptureGateway object to a map representation for JSON marshaling.
func (c CaptureGateway) toMap() map[string]any {
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
    structMap["gateways"] = c.Gateways
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
    structMap["type"] = c.Type
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CaptureGateway.
// It customizes the JSON unmarshaling process for CaptureGateway objects.
func (c *CaptureGateway) UnmarshalJSON(input []byte) error {
    var temp tempCaptureGateway
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "duration", "format", "gateways", "max_pkt_len", "num_packets", "ports", "type")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.Duration = temp.Duration
    c.Format = temp.Format
    c.Gateways = *temp.Gateways
    c.MaxPktLen = temp.MaxPktLen
    c.NumPackets = temp.NumPackets
    c.Ports = temp.Ports
    c.Type = *temp.Type
    return nil
}

// tempCaptureGateway is a temporary struct used for validating the fields of CaptureGateway.
type tempCaptureGateway  struct {
    Duration   Optional[int]                         `json:"duration"`
    Format     *CaptureGatewayFormatEnum             `json:"format,omitempty"`
    Gateways   *map[string]CaptureGatewayGateways    `json:"gateways"`
    MaxPktLen  Optional[int]                         `json:"max_pkt_len"`
    NumPackets Optional[int]                         `json:"num_packets"`
    Ports      map[string]CaptureGatewayGatewaysPort `json:"ports,omitempty"`
    Type       *string                               `json:"type"`
}

func (c *tempCaptureGateway) validate() error {
    var errs []string
    if c.Gateways == nil {
        errs = append(errs, "required field `gateways` is missing for type `capture_gateway`")
    }
    if c.Type == nil {
        errs = append(errs, "required field `type` is missing for type `capture_gateway`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
