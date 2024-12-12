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
    Duration             *int                                  `json:"duration,omitempty"`
    // enum: `stream`
    Format               *CaptureGatewayFormatEnum             `json:"format,omitempty"`
    // List of SSRs. Property key is the SSR MAC
    Gateways             map[string]CaptureGatewayGateways     `json:"gateways"`
    // max_len of each packet to capture
    MaxPktLen            *int                                  `json:"max_pkt_len,omitempty"`
    // number of packets to capture, 0 for unlimited
    NumPackets           *int                                  `json:"num_packets,omitempty"`
    // Property key is the port ID
    Ports                map[string]CaptureGatewayGatewaysPort `json:"ports,omitempty"`
    // enum: `gateway`
    Type                 string                                `json:"type"`
    AdditionalProperties map[string]interface{}                `json:"_"`
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
    if c.Duration != nil {
        structMap["duration"] = c.Duration
    }
    if c.Format != nil {
        structMap["format"] = c.Format
    }
    structMap["gateways"] = c.Gateways
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
    Duration   *int                                  `json:"duration,omitempty"`
    Format     *CaptureGatewayFormatEnum             `json:"format,omitempty"`
    Gateways   *map[string]CaptureGatewayGateways    `json:"gateways"`
    MaxPktLen  *int                                  `json:"max_pkt_len,omitempty"`
    NumPackets *int                                  `json:"num_packets,omitempty"`
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
