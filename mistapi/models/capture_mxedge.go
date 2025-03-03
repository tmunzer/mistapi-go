package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// CaptureMxedge represents a CaptureMxedge struct.
// Initiate a Wireless Packet Capture
type CaptureMxedge struct {
    // Duration of the capture, in seconds
    Duration             *int                            `json:"duration,omitempty"`
    // PCAP format. enum:
    // * `stream`: to Mist cloud
    // * `tzsp`: stream packets (over UDP as TZSP packets) to a remote host (typically running Wireshark)
    Format               *CaptureMxedgeFormatEnum        `json:"format,omitempty"`
    // Max_len of each packet to capture
    MaxPktLen            *int                            `json:"max_pkt_len,omitempty"`
    Mxedges              map[string]CaptureMxedgeMxedges `json:"mxedges,omitempty"`
    // Number of packets to capture, 0 for unlimited
    NumPackets           *int                            `json:"num_packets,omitempty"`
    // enum: `mxedge`
    Type                 string                          `json:"type"`
    // Required if `format`==`tzsp`. Remote host accessible to mxedges over the network for receiving the captured packets.
    TzspHost             *string                         `json:"tzsp_host,omitempty"`
    // If `format`==`tzsp`. Port on remote host for receiving the captured packets
    TzspPort             *int                            `json:"tzsp_port,omitempty"`
    AdditionalProperties map[string]interface{}          `json:"_"`
}

// String implements the fmt.Stringer interface for CaptureMxedge,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CaptureMxedge) String() string {
    return fmt.Sprintf(
    	"CaptureMxedge[Duration=%v, Format=%v, MaxPktLen=%v, Mxedges=%v, NumPackets=%v, Type=%v, TzspHost=%v, TzspPort=%v, AdditionalProperties=%v]",
    	c.Duration, c.Format, c.MaxPktLen, c.Mxedges, c.NumPackets, c.Type, c.TzspHost, c.TzspPort, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CaptureMxedge.
// It customizes the JSON marshaling process for CaptureMxedge objects.
func (c CaptureMxedge) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "duration", "format", "max_pkt_len", "mxedges", "num_packets", "type", "tzsp_host", "tzsp_port"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CaptureMxedge object to a map representation for JSON marshaling.
func (c CaptureMxedge) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
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
    if c.TzspHost != nil {
        structMap["tzsp_host"] = c.TzspHost
    }
    if c.TzspPort != nil {
        structMap["tzsp_port"] = c.TzspPort
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CaptureMxedge.
// It customizes the JSON unmarshaling process for CaptureMxedge objects.
func (c *CaptureMxedge) UnmarshalJSON(input []byte) error {
    var temp tempCaptureMxedge
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "duration", "format", "max_pkt_len", "mxedges", "num_packets", "type", "tzsp_host", "tzsp_port")
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
    c.TzspHost = temp.TzspHost
    c.TzspPort = temp.TzspPort
    return nil
}

// tempCaptureMxedge is a temporary struct used for validating the fields of CaptureMxedge.
type tempCaptureMxedge  struct {
    Duration   *int                            `json:"duration,omitempty"`
    Format     *CaptureMxedgeFormatEnum        `json:"format,omitempty"`
    MaxPktLen  *int                            `json:"max_pkt_len,omitempty"`
    Mxedges    map[string]CaptureMxedgeMxedges `json:"mxedges,omitempty"`
    NumPackets *int                            `json:"num_packets,omitempty"`
    Type       *string                         `json:"type"`
    TzspHost   *string                         `json:"tzsp_host,omitempty"`
    TzspPort   *int                            `json:"tzsp_port,omitempty"`
}

func (c *tempCaptureMxedge) validate() error {
    var errs []string
    if c.Type == nil {
        errs = append(errs, "required field `type` is missing for type `capture_mxedge`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
