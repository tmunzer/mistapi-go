package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// CaptureNewAssoc represents a CaptureNewAssoc struct.
// Initiate a packet Capture for New Wireless Client Associations
type CaptureNewAssoc struct {
    ApMac                *string                `json:"ap_mac,omitempty"`
    // client mac, required if `type`==`client`; optional otherwise
    ClientMac            *string                `json:"client_mac,omitempty"`
    // duration of the capture, in seconds
    Duration             *int                   `json:"duration,omitempty"`
    IncludesMcast        *bool                  `json:"includes_mcast,omitempty"`
    MaxPktLen            *int                   `json:"max_pkt_len,omitempty"`
    // number of packets to capture, 0 for unlimited
    NumPackets           *int                   `json:"num_packets,omitempty"`
    // optional filter by ssid
    Ssid                 *string                `json:"ssid,omitempty"`
    // enum: `new_assoc`
    Type                 string                 `json:"type"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for CaptureNewAssoc,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CaptureNewAssoc) String() string {
    return fmt.Sprintf(
    	"CaptureNewAssoc[ApMac=%v, ClientMac=%v, Duration=%v, IncludesMcast=%v, MaxPktLen=%v, NumPackets=%v, Ssid=%v, Type=%v, AdditionalProperties=%v]",
    	c.ApMac, c.ClientMac, c.Duration, c.IncludesMcast, c.MaxPktLen, c.NumPackets, c.Ssid, c.Type, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CaptureNewAssoc.
// It customizes the JSON marshaling process for CaptureNewAssoc objects.
func (c CaptureNewAssoc) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "ap_mac", "client_mac", "duration", "includes_mcast", "max_pkt_len", "num_packets", "ssid", "type"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CaptureNewAssoc object to a map representation for JSON marshaling.
func (c CaptureNewAssoc) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    if c.ApMac != nil {
        structMap["ap_mac"] = c.ApMac
    }
    if c.ClientMac != nil {
        structMap["client_mac"] = c.ClientMac
    }
    if c.Duration != nil {
        structMap["duration"] = c.Duration
    }
    if c.IncludesMcast != nil {
        structMap["includes_mcast"] = c.IncludesMcast
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
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CaptureNewAssoc.
// It customizes the JSON unmarshaling process for CaptureNewAssoc objects.
func (c *CaptureNewAssoc) UnmarshalJSON(input []byte) error {
    var temp tempCaptureNewAssoc
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "ap_mac", "client_mac", "duration", "includes_mcast", "max_pkt_len", "num_packets", "ssid", "type")
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

// tempCaptureNewAssoc is a temporary struct used for validating the fields of CaptureNewAssoc.
type tempCaptureNewAssoc  struct {
    ApMac         *string `json:"ap_mac,omitempty"`
    ClientMac     *string `json:"client_mac,omitempty"`
    Duration      *int    `json:"duration,omitempty"`
    IncludesMcast *bool   `json:"includes_mcast,omitempty"`
    MaxPktLen     *int    `json:"max_pkt_len,omitempty"`
    NumPackets    *int    `json:"num_packets,omitempty"`
    Ssid          *string `json:"ssid,omitempty"`
    Type          *string `json:"type"`
}

func (c *tempCaptureNewAssoc) validate() error {
    var errs []string
    if c.Type == nil {
        errs = append(errs, "required field `type` is missing for type `capture_new_assoc`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
