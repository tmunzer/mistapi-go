package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// CaptureClient represents a CaptureClient struct.
// Initiate a Client Packet Capture
type CaptureClient struct {
    ApMac                Optional[string] `json:"ap_mac"`
    // client mac, required if `type`==`client`; optional otherwise
    ClientMac            Optional[string] `json:"client_mac"`
    // duration of the capture, in seconds
    Duration             Optional[int]    `json:"duration"`
    IncludesMcast        *bool            `json:"includes_mcast,omitempty"`
    MaxPktLen            Optional[int]    `json:"max_pkt_len"`
    // number of packets to capture, 0 for unlimited, default is 1024 for client-capture
    NumPackets           Optional[int]    `json:"num_packets"`
    // optional filter by ssid
    Ssid                 Optional[string] `json:"ssid"`
    // client
    Type                 string           `json:"type"`
    AdditionalProperties map[string]any   `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CaptureClient.
// It customizes the JSON marshaling process for CaptureClient objects.
func (c CaptureClient) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CaptureClient object to a map representation for JSON marshaling.
func (c CaptureClient) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    if c.ApMac.IsValueSet() {
        if c.ApMac.Value() != nil {
            structMap["ap_mac"] = c.ApMac.Value()
        } else {
            structMap["ap_mac"] = nil
        }
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
    if c.IncludesMcast != nil {
        structMap["includes_mcast"] = c.IncludesMcast
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
    if c.Ssid.IsValueSet() {
        if c.Ssid.Value() != nil {
            structMap["ssid"] = c.Ssid.Value()
        } else {
            structMap["ssid"] = nil
        }
    }
    structMap["type"] = c.Type
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CaptureClient.
// It customizes the JSON unmarshaling process for CaptureClient objects.
func (c *CaptureClient) UnmarshalJSON(input []byte) error {
    var temp captureClient
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "ap_mac", "client_mac", "duration", "includes_mcast", "max_pkt_len", "num_packets", "ssid", "type")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.ApMac = temp.ApMac
    c.ClientMac = temp.ClientMac
    c.Duration = temp.Duration
    c.IncludesMcast = temp.IncludesMcast
    c.MaxPktLen = temp.MaxPktLen
    c.NumPackets = temp.NumPackets
    c.Ssid = temp.Ssid
    c.Type = *temp.Type
    return nil
}

// captureClient is a temporary struct used for validating the fields of CaptureClient.
type captureClient  struct {
    ApMac         Optional[string] `json:"ap_mac"`
    ClientMac     Optional[string] `json:"client_mac"`
    Duration      Optional[int]    `json:"duration"`
    IncludesMcast *bool            `json:"includes_mcast,omitempty"`
    MaxPktLen     Optional[int]    `json:"max_pkt_len"`
    NumPackets    Optional[int]    `json:"num_packets"`
    Ssid          Optional[string] `json:"ssid"`
    Type          *string          `json:"type"`
}

func (c *captureClient) validate() error {
    var errs []string
    if c.Type == nil {
        errs = append(errs, "required field `type` is missing for type `Capture_Client`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}