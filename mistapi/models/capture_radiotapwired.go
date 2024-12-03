package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// CaptureRadiotapwired represents a CaptureRadiotapwired struct.
// Initiate a Radiotap Packet Capture and Wired Packet Capture
type CaptureRadiotapwired struct {
    ApMac                     Optional[string]                `json:"ap_mac"`
    // only used for radiotap. enum: `24`, `24,5,6`, `5`, `6`
    Band                      *CaptureRadiotapwiredBandEnum   `json:"band,omitempty"`
    ClientMac                 Optional[string]                `json:"client_mac"`
    // duration of the capture, in seconds
    Duration                  *int                            `json:"duration,omitempty"`
    // enum: `pcap`, `stream`
    Format                    *CaptureRadiotapwiredFormatEnum `json:"format,omitempty"`
    // max_len of each packet to capture
    MaxPktLen                 *int                            `json:"max_pkt_len,omitempty"`
    // number of packets to capture, 0 for unlimited
    NumPackets                *int                            `json:"num_packets,omitempty"`
    // tcpdump expression for radiotap interface (802.11 + radio headers)
    RadiotapTcpdumpExpression *string                         `json:"radiotap_tcpdump_expression,omitempty"`
    Ssid                      Optional[string]                `json:"ssid"`
    // tcpdump expression common for wired,radiotap
    TcpdumpExpression         *string                         `json:"tcpdump_expression,omitempty"`
    // enum: `radiotap,wired`
    Type                      string                          `json:"type"`
    // tcpdump expression for wired
    WiredTcpdumpExpression    *string                         `json:"wired_tcpdump_expression,omitempty"`
    // tcpdump expression for radiotap interface (802.11)
    WirelessTcpdumpExpression *string                         `json:"wireless_tcpdump_expression,omitempty"`
    // wlan id associated with the respective ssid.
    WlanId                    Optional[string]                `json:"wlan_id"`
    AdditionalProperties      map[string]interface{}          `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CaptureRadiotapwired.
// It customizes the JSON marshaling process for CaptureRadiotapwired objects.
func (c CaptureRadiotapwired) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "ap_mac", "band", "client_mac", "duration", "format", "max_pkt_len", "num_packets", "radiotap_tcpdump_expression", "ssid", "tcpdump_expression", "type", "wired_tcpdump_expression", "wireless_tcpdump_expression", "wlan_id"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CaptureRadiotapwired object to a map representation for JSON marshaling.
func (c CaptureRadiotapwired) toMap() map[string]any {
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
    if c.ClientMac.IsValueSet() {
        if c.ClientMac.Value() != nil {
            structMap["client_mac"] = c.ClientMac.Value()
        } else {
            structMap["client_mac"] = nil
        }
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
    if c.RadiotapTcpdumpExpression != nil {
        structMap["radiotap_tcpdump_expression"] = c.RadiotapTcpdumpExpression
    }
    if c.Ssid.IsValueSet() {
        if c.Ssid.Value() != nil {
            structMap["ssid"] = c.Ssid.Value()
        } else {
            structMap["ssid"] = nil
        }
    }
    if c.TcpdumpExpression != nil {
        structMap["tcpdump_expression"] = c.TcpdumpExpression
    }
    structMap["type"] = c.Type
    if c.WiredTcpdumpExpression != nil {
        structMap["wired_tcpdump_expression"] = c.WiredTcpdumpExpression
    }
    if c.WirelessTcpdumpExpression != nil {
        structMap["wireless_tcpdump_expression"] = c.WirelessTcpdumpExpression
    }
    if c.WlanId.IsValueSet() {
        if c.WlanId.Value() != nil {
            structMap["wlan_id"] = c.WlanId.Value()
        } else {
            structMap["wlan_id"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CaptureRadiotapwired.
// It customizes the JSON unmarshaling process for CaptureRadiotapwired objects.
func (c *CaptureRadiotapwired) UnmarshalJSON(input []byte) error {
    var temp tempCaptureRadiotapwired
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "ap_mac", "band", "client_mac", "duration", "format", "max_pkt_len", "num_packets", "radiotap_tcpdump_expression", "ssid", "tcpdump_expression", "type", "wired_tcpdump_expression", "wireless_tcpdump_expression", "wlan_id")
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
    c.RadiotapTcpdumpExpression = temp.RadiotapTcpdumpExpression
    c.Ssid = temp.Ssid
    c.TcpdumpExpression = temp.TcpdumpExpression
    c.Type = *temp.Type
    c.WiredTcpdumpExpression = temp.WiredTcpdumpExpression
    c.WirelessTcpdumpExpression = temp.WirelessTcpdumpExpression
    c.WlanId = temp.WlanId
    return nil
}

// tempCaptureRadiotapwired is a temporary struct used for validating the fields of CaptureRadiotapwired.
type tempCaptureRadiotapwired  struct {
    ApMac                     Optional[string]                `json:"ap_mac"`
    Band                      *CaptureRadiotapwiredBandEnum   `json:"band,omitempty"`
    ClientMac                 Optional[string]                `json:"client_mac"`
    Duration                  *int                            `json:"duration,omitempty"`
    Format                    *CaptureRadiotapwiredFormatEnum `json:"format,omitempty"`
    MaxPktLen                 *int                            `json:"max_pkt_len,omitempty"`
    NumPackets                *int                            `json:"num_packets,omitempty"`
    RadiotapTcpdumpExpression *string                         `json:"radiotap_tcpdump_expression,omitempty"`
    Ssid                      Optional[string]                `json:"ssid"`
    TcpdumpExpression         *string                         `json:"tcpdump_expression,omitempty"`
    Type                      *string                         `json:"type"`
    WiredTcpdumpExpression    *string                         `json:"wired_tcpdump_expression,omitempty"`
    WirelessTcpdumpExpression *string                         `json:"wireless_tcpdump_expression,omitempty"`
    WlanId                    Optional[string]                `json:"wlan_id"`
}

func (c *tempCaptureRadiotapwired) validate() error {
    var errs []string
    if c.Type == nil {
        errs = append(errs, "required field `type` is missing for type `capture_radiotapwired`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
