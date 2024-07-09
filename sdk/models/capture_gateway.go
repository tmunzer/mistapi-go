package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// CaptureGateway represents a CaptureGateway struct.
// Initiate a Gateway (SSR) Packet Capture
type CaptureGateway struct {
    // duration of the capture, in seconds
    Duration             *int                              `json:"duration,omitempty"`
    Format               *CaptureGatewayFormatEnum         `json:"format,omitempty"`
    // List of SSRs. Property key is the SSR MAC
    Gateways             map[string]CaptureGatewayGateways `json:"gateways,omitempty"`
    // max_len of each packet to capture
    MaxPktLen            *int                              `json:"max_pkt_len,omitempty"`
    // number of packets to capture, 0 for unlimited
    NumPackets           *int                              `json:"num_packets,omitempty"`
    // dict of port which uses port id as the key
    Ports                []string                          `json:"ports,omitempty"`
    Type                 string                            `json:"type"`
    AdditionalProperties map[string]any                    `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CaptureGateway.
// It customizes the JSON marshaling process for CaptureGateway objects.
func (c CaptureGateway) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CaptureGateway object to a map representation for JSON marshaling.
func (c CaptureGateway) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    if c.Duration != nil {
        structMap["duration"] = c.Duration
    }
    if c.Format != nil {
        structMap["format"] = c.Format
    }
    if c.Gateways != nil {
        structMap["gateways"] = c.Gateways
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
    structMap["type"] = c.Type
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CaptureGateway.
// It customizes the JSON unmarshaling process for CaptureGateway objects.
func (c *CaptureGateway) UnmarshalJSON(input []byte) error {
    var temp captureGateway
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "duration", "format", "gateways", "max_pkt_len", "num_packets", "ports", "type")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.Duration = temp.Duration
    c.Format = temp.Format
    c.Gateways = temp.Gateways
    c.MaxPktLen = temp.MaxPktLen
    c.NumPackets = temp.NumPackets
    c.Ports = temp.Ports
    c.Type = *temp.Type
    return nil
}

// captureGateway is a temporary struct used for validating the fields of CaptureGateway.
type captureGateway  struct {
    Duration   *int                              `json:"duration,omitempty"`
    Format     *CaptureGatewayFormatEnum         `json:"format,omitempty"`
    Gateways   map[string]CaptureGatewayGateways `json:"gateways,omitempty"`
    MaxPktLen  *int                              `json:"max_pkt_len,omitempty"`
    NumPackets *int                              `json:"num_packets,omitempty"`
    Ports      []string                          `json:"ports,omitempty"`
    Type       *string                           `json:"type"`
}

func (c *captureGateway) validate() error {
    var errs []string
    if c.Type == nil {
        errs = append(errs, "required field `type` is missing for type `Capture_Gateway`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
