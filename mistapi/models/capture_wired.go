package models

import (
	"encoding/json"
	"errors"
	"strings"
)

// CaptureWired represents a CaptureWired struct.
// Initiate a Wired Packet Capture
type CaptureWired struct {
	ApMac Optional[string] `json:"ap_mac"`
	// duration of the capture, in seconds
	Duration *int `json:"duration,omitempty"`
	// pcap format. enum: `pcap`, `stream`
	Format *CaptureWiredFormatEnum `json:"format,omitempty"`
	// max_len of each packet to capture
	MaxPktLen *int `json:"max_pkt_len,omitempty"`
	// number of packets to capture, 0 for unlimited
	NumPackets *int `json:"num_packets,omitempty"`
	// tcpdump expression
	TcpdumpExpression Optional[string] `json:"tcpdump_expression"`
	// enum: `wired`
	Type                 string         `json:"type"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CaptureWired.
// It customizes the JSON marshaling process for CaptureWired objects.
func (c CaptureWired) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CaptureWired object to a map representation for JSON marshaling.
func (c CaptureWired) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, c.AdditionalProperties)
	if c.ApMac.IsValueSet() {
		if c.ApMac.Value() != nil {
			structMap["ap_mac"] = c.ApMac.Value()
		} else {
			structMap["ap_mac"] = nil
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
	additionalProperties, err := UnmarshalAdditionalProperties(input, "ap_mac", "duration", "format", "max_pkt_len", "num_packets", "tcpdump_expression", "type")
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
type tempCaptureWired struct {
	ApMac             Optional[string]        `json:"ap_mac"`
	Duration          *int                    `json:"duration,omitempty"`
	Format            *CaptureWiredFormatEnum `json:"format,omitempty"`
	MaxPktLen         *int                    `json:"max_pkt_len,omitempty"`
	NumPackets        *int                    `json:"num_packets,omitempty"`
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
	return errors.New(strings.Join(errs, "\n"))
}
