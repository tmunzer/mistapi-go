package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/google/uuid"
    "strings"
)

// CaptureWireless represents a CaptureWireless struct.
// Initiate a Wireless Packet Capture
type CaptureWireless struct {
    ApMac                Optional[string]           `json:"ap_mac"`
    // enum: `24`, `24,5,6`, `5`, `6`
    Band                 *CaptureWirelessBandEnum   `json:"band,omitempty"`
    // Duration of the capture, in seconds
    Duration             *int                       `json:"duration,omitempty"`
    // pcap format. enum: `pcap`, `stream`
    Format               *CaptureWirelessFormatEnum `json:"format,omitempty"`
    // Max_len of each packet to capture
    MaxPktLen            *int                       `json:"max_pkt_len,omitempty"`
    // Number of packets to capture, 0 for unlimited
    NumPackets           *int                       `json:"num_packets,omitempty"`
    Ssid                 *string                    `json:"ssid,omitempty"`
    // enum: `wireless`
    Type                 string                     `json:"type"`
    // WLAN ID
    WlanId               *uuid.UUID                 `json:"wlan_id,omitempty"`
    AdditionalProperties map[string]interface{}     `json:"_"`
}

// String implements the fmt.Stringer interface for CaptureWireless,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CaptureWireless) String() string {
    return fmt.Sprintf(
    	"CaptureWireless[ApMac=%v, Band=%v, Duration=%v, Format=%v, MaxPktLen=%v, NumPackets=%v, Ssid=%v, Type=%v, WlanId=%v, AdditionalProperties=%v]",
    	c.ApMac, c.Band, c.Duration, c.Format, c.MaxPktLen, c.NumPackets, c.Ssid, c.Type, c.WlanId, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CaptureWireless.
// It customizes the JSON marshaling process for CaptureWireless objects.
func (c CaptureWireless) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "ap_mac", "band", "duration", "format", "max_pkt_len", "num_packets", "ssid", "type", "wlan_id"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CaptureWireless object to a map representation for JSON marshaling.
func (c CaptureWireless) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    if c.ApMac.IsValueSet() {
        if c.ApMac.Value() != nil {
            structMap["ap_mac"] = c.ApMac.Value()
        } else {
            structMap["ap_mac"] = nil
        }
    }
    if c.Band != nil {
        structMap["band"] = c.Band
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
    structMap["type"] = c.Type
    if c.WlanId != nil {
        structMap["wlan_id"] = c.WlanId
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CaptureWireless.
// It customizes the JSON unmarshaling process for CaptureWireless objects.
func (c *CaptureWireless) UnmarshalJSON(input []byte) error {
    var temp tempCaptureWireless
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "ap_mac", "band", "duration", "format", "max_pkt_len", "num_packets", "ssid", "type", "wlan_id")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.ApMac = temp.ApMac
    c.Band = temp.Band
    c.Duration = temp.Duration
    c.Format = temp.Format
    c.MaxPktLen = temp.MaxPktLen
    c.NumPackets = temp.NumPackets
    c.Ssid = temp.Ssid
    c.Type = *temp.Type
    c.WlanId = temp.WlanId
    return nil
}

// tempCaptureWireless is a temporary struct used for validating the fields of CaptureWireless.
type tempCaptureWireless  struct {
    ApMac      Optional[string]           `json:"ap_mac"`
    Band       *CaptureWirelessBandEnum   `json:"band,omitempty"`
    Duration   *int                       `json:"duration,omitempty"`
    Format     *CaptureWirelessFormatEnum `json:"format,omitempty"`
    MaxPktLen  *int                       `json:"max_pkt_len,omitempty"`
    NumPackets *int                       `json:"num_packets,omitempty"`
    Ssid       *string                    `json:"ssid,omitempty"`
    Type       *string                    `json:"type"`
    WlanId     *uuid.UUID                 `json:"wlan_id,omitempty"`
}

func (c *tempCaptureWireless) validate() error {
    var errs []string
    if c.Type == nil {
        errs = append(errs, "required field `type` is missing for type `capture_wireless`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
