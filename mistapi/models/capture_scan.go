package models

import (
	"encoding/json"
	"errors"
	"strings"
)

// CaptureScan represents a CaptureScan struct.
// Initiate a Scan Radio Packet Capture
type CaptureScan struct {
	// filter by ap_mac
	ApMac Optional[string] `json:"ap_mac"`
	// dictionary key is AP mac and value is a dictionary which contains key “band”, “bandwidth”, “channel” and “tcpdump_expression”. In case keys are missed we will take parent value if parent values are not set we will use default value
	Aps map[string]CaptureScanAps `json:"aps,omitempty"`
	// Only Single value allowed, default value gets applied when user provides wrong values. enum: `24`, `5`, `6`
	Band Optional[CaptureScanBandEnum] `json:"band"`
	// channel width for the band.enum: `20`, `40`, `80` (only applicable for band_5 and band_6), `160` (only for band_6)
	Bandwidth *Dot11BandwidthEnum `json:"bandwidth,omitempty"`
	// specify the channel value where scan PCAP has to be started, default value gets applied when user provides wrong values
	Channel *int `json:"channel,omitempty"`
	// filter by client mac
	ClientMac Optional[string] `json:"client_mac"`
	// duration of the capture, in seconds
	Duration *int `json:"duration,omitempty"`
	// enum: `pcap`, `stream`
	Format *CaptureScanFormatEnum `json:"format,omitempty"`
	// max_len of each packet to capture
	MaxPktLen *int `json:"max_pkt_len,omitempty"`
	// number of packets to capture, 0 for unlimited
	NumPackets *int `json:"num_packets,omitempty"`
	// tcpdump expression, port specific if specified under ports dict, otherwise applicable across ports if specified at top level of payload. Port specific value overrides top level value when both exist.
	TcpdumpExpression *string `json:"tcpdump_expression,omitempty"`
	// enum: `scan`
	Type string `json:"type"`
	// specify the bandwidth value with respect to the channel.
	Width                *string        `json:"width,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CaptureScan.
// It customizes the JSON marshaling process for CaptureScan objects.
func (c CaptureScan) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CaptureScan object to a map representation for JSON marshaling.
func (c CaptureScan) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, c.AdditionalProperties)
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
	additionalProperties, err := UnmarshalAdditionalProperties(input, "ap_mac", "aps", "band", "bandwidth", "channel", "client_mac", "duration", "format", "max_pkt_len", "num_packets", "tcpdump_expression", "type", "width")
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
type tempCaptureScan struct {
	ApMac             Optional[string]              `json:"ap_mac"`
	Aps               map[string]CaptureScanAps     `json:"aps,omitempty"`
	Band              Optional[CaptureScanBandEnum] `json:"band"`
	Bandwidth         *Dot11BandwidthEnum           `json:"bandwidth,omitempty"`
	Channel           *int                          `json:"channel,omitempty"`
	ClientMac         Optional[string]              `json:"client_mac"`
	Duration          *int                          `json:"duration,omitempty"`
	Format            *CaptureScanFormatEnum        `json:"format,omitempty"`
	MaxPktLen         *int                          `json:"max_pkt_len,omitempty"`
	NumPackets        *int                          `json:"num_packets,omitempty"`
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
	return errors.New(strings.Join(errs, "\n"))
}
