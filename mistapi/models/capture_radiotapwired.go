package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// CaptureRadiotapwired represents a CaptureRadiotapwired struct.
// Initiate a Radiotap Packet Capture and Wired Packet Capture
type CaptureRadiotapwired struct {
    ApMac                     Optional[string]                `json:"ap_mac"`
    // only used for radiotap. enum: `24`, `24,5,6`, `5`, `6`
    Band                      *CaptureRadiotapwiredBandEnum   `json:"band,omitempty"`
    ClientMac                 Optional[string]                `json:"client_mac"`
    // Duration of the capture, in seconds
    Duration                  Optional[int]                   `json:"duration"`
    // enum: `pcap`, `stream`
    Format                    *CaptureRadiotapwiredFormatEnum `json:"format,omitempty"`
    MaxPktLen                 Optional[int]                   `json:"max_pkt_len"`
    // number of packets to capture, 0 for unlimited, default is 1024, maximum is 10000
    NumPackets                Optional[int]                   `json:"num_packets"`
    // tcpdump expression for radiotap interface (802.11 + radio headers)
    RadiotapTcpdumpExpression *string                         `json:"radiotap_tcpdump_expression,omitempty"`
    Ssid                      Optional[string]                `json:"ssid"`
    // tcpdump expression
    TcpdumpExpression         Optional[string]                `json:"tcpdump_expression"`
    // enum: `radiotap,wired`
    Type                      string                          `json:"type"`
    // tcpdump expression
    WiredTcpdumpExpression    Optional[string]                `json:"wired_tcpdump_expression"`
    // tcpdump expression for radiotap interface (802.11)
    WirelessTcpdumpExpression *string                         `json:"wireless_tcpdump_expression,omitempty"`
    // WLAN id associated with the respective ssid.
    WlanId                    Optional[string]                `json:"wlan_id"`
    AdditionalProperties      map[string]interface{}          `json:"_"`
}

// String implements the fmt.Stringer interface for CaptureRadiotapwired,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CaptureRadiotapwired) String() string {
    return fmt.Sprintf(
    	"CaptureRadiotapwired[ApMac=%v, Band=%v, ClientMac=%v, Duration=%v, Format=%v, MaxPktLen=%v, NumPackets=%v, RadiotapTcpdumpExpression=%v, Ssid=%v, TcpdumpExpression=%v, Type=%v, WiredTcpdumpExpression=%v, WirelessTcpdumpExpression=%v, WlanId=%v, AdditionalProperties=%v]",
    	c.ApMac, c.Band, c.ClientMac, c.Duration, c.Format, c.MaxPktLen, c.NumPackets, c.RadiotapTcpdumpExpression, c.Ssid, c.TcpdumpExpression, c.Type, c.WiredTcpdumpExpression, c.WirelessTcpdumpExpression, c.WlanId, c.AdditionalProperties)
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
    if c.TcpdumpExpression.IsValueSet() {
        if c.TcpdumpExpression.Value() != nil {
            structMap["tcpdump_expression"] = c.TcpdumpExpression.Value()
        } else {
            structMap["tcpdump_expression"] = nil
        }
    }
    structMap["type"] = c.Type
    if c.WiredTcpdumpExpression.IsValueSet() {
        if c.WiredTcpdumpExpression.Value() != nil {
            structMap["wired_tcpdump_expression"] = c.WiredTcpdumpExpression.Value()
        } else {
            structMap["wired_tcpdump_expression"] = nil
        }
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
    Duration                  Optional[int]                   `json:"duration"`
    Format                    *CaptureRadiotapwiredFormatEnum `json:"format,omitempty"`
    MaxPktLen                 Optional[int]                   `json:"max_pkt_len"`
    NumPackets                Optional[int]                   `json:"num_packets"`
    RadiotapTcpdumpExpression *string                         `json:"radiotap_tcpdump_expression,omitempty"`
    Ssid                      Optional[string]                `json:"ssid"`
    TcpdumpExpression         Optional[string]                `json:"tcpdump_expression"`
    Type                      *string                         `json:"type"`
    WiredTcpdumpExpression    Optional[string]                `json:"wired_tcpdump_expression"`
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
