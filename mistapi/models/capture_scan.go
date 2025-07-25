// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// CaptureScan represents a CaptureScan struct.
// Initiate a Scan Radio Packet Capture
type CaptureScan struct {
    // Filter by ap_mac
    ApMac                Optional[string]              `json:"ap_mac"`
    // Dictionary key is AP mac and value is a dictionary which contains key "band", "bandwidth", "channel" and "tcpdump_expression". In case keys are missed we will take parent value if parent values are not set we will use default value
    Aps                  map[string]CaptureScanAps     `json:"aps,omitempty"`
    // Only Single value allowed, default value gets applied when user provides wrong values. enum: `24`, `5`, `6`
    Band                 Optional[CaptureScanBandEnum] `json:"band"`
    // channel width for the band.enum: `20`, `40`, `80` (only applicable for band_5 and band_6), `160` (only for band_6)
    Bandwidth            *Dot11BandwidthEnum           `json:"bandwidth,omitempty"`
    // Specify the channel value where scan PCAP has to be started, default value gets applied when user provides wrong values
    Channel              *int                          `json:"channel,omitempty"`
    // Filter by client mac
    ClientMac            Optional[string]              `json:"client_mac"`
    // Duration of the capture, in seconds
    Duration             Optional[int]                 `json:"duration"`
    // enum: `pcap`, `stream`
    Format               *CaptureScanFormatEnum        `json:"format,omitempty"`
    MaxPktLen            Optional[int]                 `json:"max_pkt_len"`
    // number of packets to capture, 0 for unlimited, default is 1024, maximum is 10000
    NumPackets           Optional[int]                 `json:"num_packets"`
    // tcpdump expression, port specific if specified under ports dict, otherwise applicable across ports if specified at top level of payload. Port specific value overrides top level value when both exist.
    TcpdumpExpression    *string                       `json:"tcpdump_expression,omitempty"`
    // enum: `scan`
    Type                 string                        `json:"type"`
    // Specify the bandwidth value with respect to the channel.
    Width                *string                       `json:"width,omitempty"`
    AdditionalProperties map[string]interface{}        `json:"_"`
}

// String implements the fmt.Stringer interface for CaptureScan,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CaptureScan) String() string {
    return fmt.Sprintf(
    	"CaptureScan[ApMac=%v, Aps=%v, Band=%v, Bandwidth=%v, Channel=%v, ClientMac=%v, Duration=%v, Format=%v, MaxPktLen=%v, NumPackets=%v, TcpdumpExpression=%v, Type=%v, Width=%v, AdditionalProperties=%v]",
    	c.ApMac, c.Aps, c.Band, c.Bandwidth, c.Channel, c.ClientMac, c.Duration, c.Format, c.MaxPktLen, c.NumPackets, c.TcpdumpExpression, c.Type, c.Width, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CaptureScan.
// It customizes the JSON marshaling process for CaptureScan objects.
func (c CaptureScan) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "ap_mac", "aps", "band", "bandwidth", "channel", "client_mac", "duration", "format", "max_pkt_len", "num_packets", "tcpdump_expression", "type", "width"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CaptureScan object to a map representation for JSON marshaling.
func (c CaptureScan) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    if c.ApMac.IsValueSet() {
        if c.ApMac.Value() != nil {
            structMap["ap_mac"] = c.ApMac.Value()
        } else {
            structMap["ap_mac"] = nil
        }
    }
    if c.Aps != nil {
        structMap["aps"] = c.Aps
    }
    if c.Band.IsValueSet() {
        if c.Band.Value() != nil {
            structMap["band"] = c.Band.Value()
        } else {
            structMap["band"] = nil
        }
    }
    if c.Bandwidth != nil {
        structMap["bandwidth"] = c.Bandwidth
    }
    if c.Channel != nil {
        structMap["channel"] = c.Channel
    }
    if c.ClientMac.IsValueSet() {
        if c.ClientMac.Value() != nil {
            structMap["client_mac"] = c.ClientMac.Value()
        } else {
            structMap["client_mac"] = nil
        }
    }
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
    if c.TcpdumpExpression != nil {
        structMap["tcpdump_expression"] = c.TcpdumpExpression
    }
    structMap["type"] = c.Type
    if c.Width != nil {
        structMap["width"] = c.Width
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CaptureScan.
// It customizes the JSON unmarshaling process for CaptureScan objects.
func (c *CaptureScan) UnmarshalJSON(input []byte) error {
    var temp tempCaptureScan
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "ap_mac", "aps", "band", "bandwidth", "channel", "client_mac", "duration", "format", "max_pkt_len", "num_packets", "tcpdump_expression", "type", "width")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.ApMac = temp.ApMac
    c.Aps = temp.Aps
    c.Band = temp.Band
    c.Bandwidth = temp.Bandwidth
    c.Channel = temp.Channel
    c.ClientMac = temp.ClientMac
    c.Duration = temp.Duration
    c.Format = temp.Format
    c.MaxPktLen = temp.MaxPktLen
    c.NumPackets = temp.NumPackets
    c.TcpdumpExpression = temp.TcpdumpExpression
    c.Type = *temp.Type
    c.Width = temp.Width
    return nil
}

// tempCaptureScan is a temporary struct used for validating the fields of CaptureScan.
type tempCaptureScan  struct {
    ApMac             Optional[string]              `json:"ap_mac"`
    Aps               map[string]CaptureScanAps     `json:"aps,omitempty"`
    Band              Optional[CaptureScanBandEnum] `json:"band"`
    Bandwidth         *Dot11BandwidthEnum           `json:"bandwidth,omitempty"`
    Channel           *int                          `json:"channel,omitempty"`
    ClientMac         Optional[string]              `json:"client_mac"`
    Duration          Optional[int]                 `json:"duration"`
    Format            *CaptureScanFormatEnum        `json:"format,omitempty"`
    MaxPktLen         Optional[int]                 `json:"max_pkt_len"`
    NumPackets        Optional[int]                 `json:"num_packets"`
    TcpdumpExpression *string                       `json:"tcpdump_expression,omitempty"`
    Type              *string                       `json:"type"`
    Width             *string                       `json:"width,omitempty"`
}

func (c *tempCaptureScan) validate() error {
    var errs []string
    if c.Type == nil {
        errs = append(errs, "required field `type` is missing for type `capture_scan`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
