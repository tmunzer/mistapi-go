// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// CaptureScanAps represents a CaptureScanAps struct.
// Property key is the AP MAC address (e.g. "5c5b35000001"). All optionals, parent parameters will be used if not defined
type CaptureScanAps struct {
    // Only Single value allowed. enum: `24`, `5`, `6`
    Band                 *CaptureScanApsBandEnum `json:"band,omitempty"`
    // Specify the channel value where scan PCAP has to be started
    Channel              *string                 `json:"channel,omitempty"`
    // tcpdump expression, port specific if specified under ports dict, otherwise applicable across ports if specified at top level of payload. Port specific value overrides top level value when both exist.
    TcpdumpExpression    *string                 `json:"tcpdump_expression,omitempty"`
    // Specify the bandwidth value with respect to the channel.
    Width                *string                 `json:"width,omitempty"`
    AdditionalProperties map[string]interface{}  `json:"_"`
}

// String implements the fmt.Stringer interface for CaptureScanAps,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CaptureScanAps) String() string {
    return fmt.Sprintf(
    	"CaptureScanAps[Band=%v, Channel=%v, TcpdumpExpression=%v, Width=%v, AdditionalProperties=%v]",
    	c.Band, c.Channel, c.TcpdumpExpression, c.Width, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CaptureScanAps.
// It customizes the JSON marshaling process for CaptureScanAps objects.
func (c CaptureScanAps) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "band", "channel", "tcpdump_expression", "width"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CaptureScanAps object to a map representation for JSON marshaling.
func (c CaptureScanAps) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    if c.Band != nil {
        structMap["band"] = c.Band
    }
    if c.Channel != nil {
        structMap["channel"] = c.Channel
    }
    if c.TcpdumpExpression != nil {
        structMap["tcpdump_expression"] = c.TcpdumpExpression
    }
    if c.Width != nil {
        structMap["width"] = c.Width
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CaptureScanAps.
// It customizes the JSON unmarshaling process for CaptureScanAps objects.
func (c *CaptureScanAps) UnmarshalJSON(input []byte) error {
    var temp tempCaptureScanAps
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "band", "channel", "tcpdump_expression", "width")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.Band = temp.Band
    c.Channel = temp.Channel
    c.TcpdumpExpression = temp.TcpdumpExpression
    c.Width = temp.Width
    return nil
}

// tempCaptureScanAps is a temporary struct used for validating the fields of CaptureScanAps.
type tempCaptureScanAps  struct {
    Band              *CaptureScanApsBandEnum `json:"band,omitempty"`
    Channel           *string                 `json:"channel,omitempty"`
    TcpdumpExpression *string                 `json:"tcpdump_expression,omitempty"`
    Width             *string                 `json:"width,omitempty"`
}
