package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// CaptureRadiotap represents a CaptureRadiotap struct.
// Initiate a Radiotap Packet Capture
type CaptureRadiotap struct {
    ApMac                *string                    `json:"ap_mac,omitempty"`
    // enum: `24`, `24,5,6`, `5`, `6`
    Band                 *CaptureRadiotapBandEnum   `json:"band,omitempty"`
    ClientMac            *string                    `json:"client_mac,omitempty"`
    // duration of the capture, in seconds
    Duration             *int                       `json:"duration,omitempty"`
    // enum: `pcap`, `stream`
    Format               *CaptureRadiotapFormatEnum `json:"format,omitempty"`
    // max_len of each packet to capture
    MaxPktLen            *int                       `json:"max_pkt_len,omitempty"`
    // number of packets to capture, 0 for unlimited
    NumPackets           *int                       `json:"num_packets,omitempty"`
    Ssid                 *string                    `json:"ssid,omitempty"`
    // tcpdump expression specific to radiotap
    TcpdumpExpression    *string                    `json:"tcpdump_expression,omitempty"`
    // enum: `radiotap`
    Type                 string                     `json:"type"`
    // wlan id associated with the respective ssid.
    WlanId               *uuid.UUID                 `json:"wlan_id,omitempty"`
    AdditionalProperties map[string]any             `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CaptureRadiotap.
// It customizes the JSON marshaling process for CaptureRadiotap objects.
func (c CaptureRadiotap) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CaptureRadiotap object to a map representation for JSON marshaling.
func (c CaptureRadiotap) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    if c.ApMac != nil {
        structMap["ap_mac"] = c.ApMac
    }
    if c.Band != nil {
        structMap["band"] = c.Band
    }
    if c.ClientMac != nil {
        structMap["client_mac"] = c.ClientMac
    }
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
    if c.Ssid != nil {
        structMap["ssid"] = c.Ssid
    }
    if c.TcpdumpExpression != nil {
        structMap["tcpdump_expression"] = c.TcpdumpExpression
    }
    structMap["type"] = c.Type
    if c.WlanId != nil {
        structMap["wlan_id"] = c.WlanId
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CaptureRadiotap.
// It customizes the JSON unmarshaling process for CaptureRadiotap objects.
func (c *CaptureRadiotap) UnmarshalJSON(input []byte) error {
    var temp tempCaptureRadiotap
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "ap_mac", "band", "client_mac", "duration", "format", "max_pkt_len", "num_packets", "ssid", "tcpdump_expression", "type", "wlan_id")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.ApMac = temp.ApMac
    c.Band = temp.Band
    c.ClientMac = temp.ClientMac
    c.Duration = temp.Duration
    c.Format = temp.Format
    c.MaxPktLen = temp.MaxPktLen
    c.NumPackets = temp.NumPackets
    c.Ssid = temp.Ssid
    c.TcpdumpExpression = temp.TcpdumpExpression
    c.Type = *temp.Type
    c.WlanId = temp.WlanId
    return nil
}

// tempCaptureRadiotap is a temporary struct used for validating the fields of CaptureRadiotap.
type tempCaptureRadiotap  struct {
    ApMac             *string                    `json:"ap_mac,omitempty"`
    Band              *CaptureRadiotapBandEnum   `json:"band,omitempty"`
    ClientMac         *string                    `json:"client_mac,omitempty"`
    Duration          *int                       `json:"duration,omitempty"`
    Format            *CaptureRadiotapFormatEnum `json:"format,omitempty"`
    MaxPktLen         *int                       `json:"max_pkt_len,omitempty"`
    NumPackets        *int                       `json:"num_packets,omitempty"`
    Ssid              *string                    `json:"ssid,omitempty"`
    TcpdumpExpression *string                    `json:"tcpdump_expression,omitempty"`
    Type              *string                    `json:"type"`
    WlanId            *uuid.UUID                 `json:"wlan_id,omitempty"`
}

func (c *tempCaptureRadiotap) validate() error {
    var errs []string
    if c.Type == nil {
        errs = append(errs, "required field `type` is missing for type `capture_radiotap`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
