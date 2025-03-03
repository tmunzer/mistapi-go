package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// CaptureWired represents a CaptureWired struct.
// Initiate a Wired Packet Capture
type CaptureWired struct {
    ApMac                Optional[string]        `json:"ap_mac"`
    // Duration of the capture, in seconds
    Duration             Optional[int]           `json:"duration"`
    // pcap format. enum: `pcap`, `stream`
    Format               *CaptureWiredFormatEnum `json:"format,omitempty"`
    MaxPktLen            Optional[int]           `json:"max_pkt_len"`
    // number of packets to capture, 0 for unlimited, default is 1024, maximum is 10000
    NumPackets           Optional[int]           `json:"num_packets"`
    // tcpdump expression
    TcpdumpExpression    Optional[string]        `json:"tcpdump_expression"`
    // enum: `wired`
    Type                 string                  `json:"type"`
    AdditionalProperties map[string]interface{}  `json:"_"`
}

// String implements the fmt.Stringer interface for CaptureWired,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CaptureWired) String() string {
    return fmt.Sprintf(
    	"CaptureWired[ApMac=%v, Duration=%v, Format=%v, MaxPktLen=%v, NumPackets=%v, TcpdumpExpression=%v, Type=%v, AdditionalProperties=%v]",
    	c.ApMac, c.Duration, c.Format, c.MaxPktLen, c.NumPackets, c.TcpdumpExpression, c.Type, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CaptureWired.
// It customizes the JSON marshaling process for CaptureWired objects.
func (c CaptureWired) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "ap_mac", "duration", "format", "max_pkt_len", "num_packets", "tcpdump_expression", "type"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CaptureWired object to a map representation for JSON marshaling.
func (c CaptureWired) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    if c.ApMac.IsValueSet() {
        if c.ApMac.Value() != nil {
            structMap["ap_mac"] = c.ApMac.Value()
        } else {
            structMap["ap_mac"] = nil
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
    if c.TcpdumpExpression.IsValueSet() {
        if c.TcpdumpExpression.Value() != nil {
            structMap["tcpdump_expression"] = c.TcpdumpExpression.Value()
        } else {
            structMap["tcpdump_expression"] = nil
        }
    }
    structMap["type"] = c.Type
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CaptureWired.
// It customizes the JSON unmarshaling process for CaptureWired objects.
func (c *CaptureWired) UnmarshalJSON(input []byte) error {
    var temp tempCaptureWired
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "ap_mac", "duration", "format", "max_pkt_len", "num_packets", "tcpdump_expression", "type")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.ApMac = temp.ApMac
    c.Duration = temp.Duration
    c.Format = temp.Format
    c.MaxPktLen = temp.MaxPktLen
    c.NumPackets = temp.NumPackets
    c.TcpdumpExpression = temp.TcpdumpExpression
    c.Type = *temp.Type
    return nil
}

// tempCaptureWired is a temporary struct used for validating the fields of CaptureWired.
type tempCaptureWired  struct {
    ApMac             Optional[string]        `json:"ap_mac"`
    Duration          Optional[int]           `json:"duration"`
    Format            *CaptureWiredFormatEnum `json:"format,omitempty"`
    MaxPktLen         Optional[int]           `json:"max_pkt_len"`
    NumPackets        Optional[int]           `json:"num_packets"`
    TcpdumpExpression Optional[string]        `json:"tcpdump_expression"`
    Type              *string                 `json:"type"`
}

func (c *tempCaptureWired) validate() error {
    var errs []string
    if c.Type == nil {
        errs = append(errs, "required field `type` is missing for type `capture_wired`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
