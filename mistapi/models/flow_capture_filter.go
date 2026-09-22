// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// FlowCaptureFilter represents a FlowCaptureFilter struct.
// Normalized filter applied to a flow capture session
type FlowCaptureFilter struct {
	DstIp   *string `json:"dst_ip,omitempty"`
	DstPort *int    `json:"dst_port,omitempty"`
	// Flow capture protocol filter. enum: `tcp`, `udp`, `icmp`, `icmp6`
	Protocol             *FlowCaptureProtocolEnum `json:"protocol,omitempty"`
	SrcIp                *string                  `json:"src_ip,omitempty"`
	SrcPort              *int                     `json:"src_port,omitempty"`
	AdditionalProperties map[string]interface{}   `json:"_"`
}

// String implements the fmt.Stringer interface for FlowCaptureFilter,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (f FlowCaptureFilter) String() string {
	return fmt.Sprintf(
		"FlowCaptureFilter[DstIp=%v, DstPort=%v, Protocol=%v, SrcIp=%v, SrcPort=%v, AdditionalProperties=%v]",
		f.DstIp, f.DstPort, f.Protocol, f.SrcIp, f.SrcPort, f.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for FlowCaptureFilter.
// It customizes the JSON marshaling process for FlowCaptureFilter objects.
func (f FlowCaptureFilter) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(f.AdditionalProperties,
		"dst_ip", "dst_port", "protocol", "src_ip", "src_port"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(f.toMap())
}

// toMap converts the FlowCaptureFilter object to a map representation for JSON marshaling.
func (f FlowCaptureFilter) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, f.AdditionalProperties)
	if f.DstIp != nil {
		structMap["dst_ip"] = f.DstIp
	}
	if f.DstPort != nil {
		structMap["dst_port"] = f.DstPort
	}
	if f.Protocol != nil {
		structMap["protocol"] = f.Protocol
	}
	if f.SrcIp != nil {
		structMap["src_ip"] = f.SrcIp
	}
	if f.SrcPort != nil {
		structMap["src_port"] = f.SrcPort
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for FlowCaptureFilter.
// It customizes the JSON unmarshaling process for FlowCaptureFilter objects.
func (f *FlowCaptureFilter) UnmarshalJSON(input []byte) error {
	var temp tempFlowCaptureFilter
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "dst_ip", "dst_port", "protocol", "src_ip", "src_port")
	if err != nil {
		return err
	}
	f.AdditionalProperties = additionalProperties

	f.DstIp = temp.DstIp
	f.DstPort = temp.DstPort
	f.Protocol = temp.Protocol
	f.SrcIp = temp.SrcIp
	f.SrcPort = temp.SrcPort
	return nil
}

// tempFlowCaptureFilter is a temporary struct used for validating the fields of FlowCaptureFilter.
type tempFlowCaptureFilter struct {
	DstIp    *string                  `json:"dst_ip,omitempty"`
	DstPort  *int                     `json:"dst_port,omitempty"`
	Protocol *FlowCaptureProtocolEnum `json:"protocol,omitempty"`
	SrcIp    *string                  `json:"src_ip,omitempty"`
	SrcPort  *int                     `json:"src_port,omitempty"`
}
