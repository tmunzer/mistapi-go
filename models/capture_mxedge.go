package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// CaptureMxedge represents a CaptureMxedge struct.
// Initiate a Wireless Packet Capture
type CaptureMxedge struct {
    // duration of the capture, in seconds
    Duration             *int                            `json:"duration,omitempty"`
    // pcap format
    Format               *CaptureMxedgeFormatEnum        `json:"format,omitempty"`
    // max_len of each packet to capture
    MaxPktLen            *int                            `json:"max_pkt_len,omitempty"`
    Mxedges              map[string]CaptureMxedgeMxedges `json:"mxedges,omitempty"`
    // number of packets to capture, 0 for unlimited
    NumPackets           *int                            `json:"num_packets,omitempty"`
    // mxedge
    Type                 string                          `json:"type"`
    AdditionalProperties map[string]any                  `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CaptureMxedge.
// It customizes the JSON marshaling process for CaptureMxedge objects.
func (c CaptureMxedge) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CaptureMxedge object to a map representation for JSON marshaling.
func (c CaptureMxedge) toMap() map[string]any {
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
    if c.Mxedges != nil {
        structMap["mxedges"] = c.Mxedges
    }
    if c.NumPackets != nil {
        structMap["num_packets"] = c.NumPackets
    }
    structMap["type"] = c.Type
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CaptureMxedge.
// It customizes the JSON unmarshaling process for CaptureMxedge objects.
func (c *CaptureMxedge) UnmarshalJSON(input []byte) error {
    var temp captureMxedge
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "duration", "format", "max_pkt_len", "mxedges", "num_packets", "type")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.Duration = temp.Duration
    c.Format = temp.Format
    c.MaxPktLen = temp.MaxPktLen
    c.Mxedges = temp.Mxedges
    c.NumPackets = temp.NumPackets
    c.Type = *temp.Type
    return nil
}

// captureMxedge is a temporary struct used for validating the fields of CaptureMxedge.
type captureMxedge  struct {
    Duration   *int                            `json:"duration,omitempty"`
    Format     *CaptureMxedgeFormatEnum        `json:"format,omitempty"`
    MaxPktLen  *int                            `json:"max_pkt_len,omitempty"`
    Mxedges    map[string]CaptureMxedgeMxedges `json:"mxedges,omitempty"`
    NumPackets *int                            `json:"num_packets,omitempty"`
    Type       *string                         `json:"type"`
}

func (c *captureMxedge) validate() error {
    var errs []string
    if c.Type == nil {
        errs = append(errs, "required field `type` is missing for type `Capture_Mxedge`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
