package models

import (
    "encoding/json"
    "errors"
    "fmt"
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
    // Duration of the capture, in seconds
    Duration             Optional[int]              `json:"duration"`
    // enum: `pcap`, `stream`
    Format               *CaptureRadiotapFormatEnum `json:"format,omitempty"`
    MaxPktLen            Optional[int]              `json:"max_pkt_len"`
    // number of packets to capture, 0 for unlimited, default is 1024, maximum is 10000
    NumPackets           Optional[int]              `json:"num_packets"`
    Ssid                 *string                    `json:"ssid,omitempty"`
    // tcpdump expression
    TcpdumpExpression    Optional[string]           `json:"tcpdump_expression"`
    // enum: `radiotap`
    Type                 string                     `json:"type"`
    // WLAN id associated with the respective ssid.
    WlanId               *uuid.UUID                 `json:"wlan_id,omitempty"`
    AdditionalProperties map[string]interface{}     `json:"_"`
}

// String implements the fmt.Stringer interface for CaptureRadiotap,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CaptureRadiotap) String() string {
    return fmt.Sprintf(
    	"CaptureRadiotap[ApMac=%v, Band=%v, ClientMac=%v, Duration=%v, Format=%v, MaxPktLen=%v, NumPackets=%v, Ssid=%v, TcpdumpExpression=%v, Type=%v, WlanId=%v, AdditionalProperties=%v]",
    	c.ApMac, c.Band, c.ClientMac, c.Duration, c.Format, c.MaxPktLen, c.NumPackets, c.Ssid, c.TcpdumpExpression, c.Type, c.WlanId, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CaptureRadiotap.
// It customizes the JSON marshaling process for CaptureRadiotap objects.
func (c CaptureRadiotap) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "ap_mac", "band", "client_mac", "duration", "format", "max_pkt_len", "num_packets", "ssid", "tcpdump_expression", "type", "wlan_id"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CaptureRadiotap object to a map representation for JSON marshaling.
func (c CaptureRadiotap) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    if c.ApMac != nil {
        structMap["ap_mac"] = c.ApMac
    }
    if c.Band != nil {
        structMap["band"] = c.Band
    }
    if c.ClientMac != nil {
        structMap["client_mac"] = c.ClientMac
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
    if c.Ssid != nil {
        structMap["ssid"] = c.Ssid
    }
    if c.TcpdumpExpression.IsValueSet() {
        if c.TcpdumpExpression.Value() != nil {
            structMap["tcpdump_expression"] = c.TcpdumpExpression.Value()
        } else {
            structMap["tcpdump_expression"] = nil
        }
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
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "ap_mac", "band", "client_mac", "duration", "format", "max_pkt_len", "num_packets", "ssid", "tcpdump_expression", "type", "wlan_id")
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
    Duration          Optional[int]              `json:"duration"`
    Format            *CaptureRadiotapFormatEnum `json:"format,omitempty"`
    MaxPktLen         Optional[int]              `json:"max_pkt_len"`
    NumPackets        Optional[int]              `json:"num_packets"`
    Ssid              *string                    `json:"ssid,omitempty"`
    TcpdumpExpression Optional[string]           `json:"tcpdump_expression"`
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
