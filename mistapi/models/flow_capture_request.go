// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// FlowCaptureRequest represents a FlowCaptureRequest struct.
// Flow capture request for one or more switches
type FlowCaptureRequest struct {
	DstIp    *string `json:"dst_ip,omitempty"`
	DstPort  *int    `json:"dst_port,omitempty"`
	Duration *int    `json:"duration,omitempty"`
	// Flow capture protocol filter. enum: `tcp`, `udp`, `icmp`, `icmp6`
	Protocol             *FlowCaptureProtocolEnum `json:"protocol,omitempty"`
	SrcIp                *string                  `json:"src_ip,omitempty"`
	SrcPort              *int                     `json:"src_port,omitempty"`
	Switches             []string                 `json:"switches"`
	AdditionalProperties map[string]interface{}   `json:"_"`
}

// String implements the fmt.Stringer interface for FlowCaptureRequest,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (f FlowCaptureRequest) String() string {
	return fmt.Sprintf(
		"FlowCaptureRequest[DstIp=%v, DstPort=%v, Duration=%v, Protocol=%v, SrcIp=%v, SrcPort=%v, Switches=%v, AdditionalProperties=%v]",
		f.DstIp, f.DstPort, f.Duration, f.Protocol, f.SrcIp, f.SrcPort, f.Switches, f.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for FlowCaptureRequest.
// It customizes the JSON marshaling process for FlowCaptureRequest objects.
func (f FlowCaptureRequest) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(f.AdditionalProperties,
		"dst_ip", "dst_port", "duration", "protocol", "src_ip", "src_port", "switches"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(f.toMap())
}

// toMap converts the FlowCaptureRequest object to a map representation for JSON marshaling.
func (f FlowCaptureRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, f.AdditionalProperties)
	if f.DstIp != nil {
		structMap["dst_ip"] = f.DstIp
	}
	if f.DstPort != nil {
		structMap["dst_port"] = f.DstPort
	}
	if f.Duration != nil {
		structMap["duration"] = f.Duration
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
	structMap["switches"] = f.Switches
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for FlowCaptureRequest.
// It customizes the JSON unmarshaling process for FlowCaptureRequest objects.
func (f *FlowCaptureRequest) UnmarshalJSON(input []byte) error {
	var temp tempFlowCaptureRequest
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "dst_ip", "dst_port", "duration", "protocol", "src_ip", "src_port", "switches")
	if err != nil {
		return err
	}
	f.AdditionalProperties = additionalProperties

	f.DstIp = temp.DstIp
	f.DstPort = temp.DstPort
	f.Duration = temp.Duration
	f.Protocol = temp.Protocol
	f.SrcIp = temp.SrcIp
	f.SrcPort = temp.SrcPort
	f.Switches = *temp.Switches
	return nil
}

// tempFlowCaptureRequest is a temporary struct used for validating the fields of FlowCaptureRequest.
type tempFlowCaptureRequest struct {
	DstIp    *string                  `json:"dst_ip,omitempty"`
	DstPort  *int                     `json:"dst_port,omitempty"`
	Duration *int                     `json:"duration,omitempty"`
	Protocol *FlowCaptureProtocolEnum `json:"protocol,omitempty"`
	SrcIp    *string                  `json:"src_ip,omitempty"`
	SrcPort  *int                     `json:"src_port,omitempty"`
	Switches *[]string                `json:"switches"`
}

func (f *tempFlowCaptureRequest) validate() error {
	var errs []string
	if f.Switches == nil {
		errs = append(errs, "required field `switches` is missing for type `flow_capture_request`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
