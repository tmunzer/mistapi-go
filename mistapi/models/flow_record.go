// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
)

// FlowRecord represents a FlowRecord struct.
// Network flow record reported by a switch device
type FlowRecord struct {
	// MAC address of the device
	DeviceMac *string `json:"device_mac,omitempty"`
	// Flow direction. enum: `egress`, `ingress`
	Direction *FlowRecordDirectionEnum `json:"direction,omitempty"`
	// Destination IP address
	DstIp *string `json:"dst_ip,omitempty"`
	// Destination port number
	DstPort *int `json:"dst_port,omitempty"`
	// Flow duration in seconds
	Duration *int64 `json:"duration,omitempty"`
	// Flow end time in epoch seconds
	EndTime *int64 `json:"end_time,omitempty"`
	// Unique flow identifier
	FlowId *int64 `json:"flow_id,omitempty"`
	// Unique identifier of a Mist organization
	OrgId *uuid.UUID `json:"org_id,omitempty"`
	// Protocol (e.g. `tcp`, `udp`, `icmp`)
	Protocol *string `json:"protocol,omitempty"`
	// Percentage of packets sampled (e.g. `0.1` means 0.1% of packets are captured via sFlow; `100.0` means all packets are captured via FBT)
	SamplingPercentage *float64 `json:"sampling_percentage,omitempty"`
	// Unique identifier of a Mist site
	SiteId *uuid.UUID `json:"site_id,omitempty"`
	// Source IP address
	SrcIp *string `json:"src_ip,omitempty"`
	// Source port number
	SrcPort *int `json:"src_port,omitempty"`
	// Flow start time in epoch seconds
	StartTime *int64 `json:"start_time,omitempty"`
	// Flow state. enum: `active`, `aged-out`
	State *FlowRecordStateEnum `json:"state,omitempty"`
	// Epoch time (in seconds) when the flow record was last updated or completed
	Timestamp *int64 `json:"timestamp,omitempty"`
	// Total number of bytes in the flow
	TotalBytes *int64 `json:"total_bytes,omitempty"`
	// Total number of packets in the flow
	TotalPkts            *int64                 `json:"total_pkts,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for FlowRecord,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (f FlowRecord) String() string {
	return fmt.Sprintf(
		"FlowRecord[DeviceMac=%v, Direction=%v, DstIp=%v, DstPort=%v, Duration=%v, EndTime=%v, FlowId=%v, OrgId=%v, Protocol=%v, SamplingPercentage=%v, SiteId=%v, SrcIp=%v, SrcPort=%v, StartTime=%v, State=%v, Timestamp=%v, TotalBytes=%v, TotalPkts=%v, AdditionalProperties=%v]",
		f.DeviceMac, f.Direction, f.DstIp, f.DstPort, f.Duration, f.EndTime, f.FlowId, f.OrgId, f.Protocol, f.SamplingPercentage, f.SiteId, f.SrcIp, f.SrcPort, f.StartTime, f.State, f.Timestamp, f.TotalBytes, f.TotalPkts, f.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for FlowRecord.
// It customizes the JSON marshaling process for FlowRecord objects.
func (f FlowRecord) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(f.AdditionalProperties,
		"device_mac", "direction", "dst_ip", "dst_port", "duration", "end_time", "flow_id", "org_id", "protocol", "sampling_percentage", "site_id", "src_ip", "src_port", "start_time", "state", "timestamp", "total_bytes", "total_pkts"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(f.toMap())
}

// toMap converts the FlowRecord object to a map representation for JSON marshaling.
func (f FlowRecord) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, f.AdditionalProperties)
	if f.DeviceMac != nil {
		structMap["device_mac"] = f.DeviceMac
	}
	if f.Direction != nil {
		structMap["direction"] = f.Direction
	}
	if f.DstIp != nil {
		structMap["dst_ip"] = f.DstIp
	}
	if f.DstPort != nil {
		structMap["dst_port"] = f.DstPort
	}
	if f.Duration != nil {
		structMap["duration"] = f.Duration
	}
	if f.EndTime != nil {
		structMap["end_time"] = f.EndTime
	}
	if f.FlowId != nil {
		structMap["flow_id"] = f.FlowId
	}
	if f.OrgId != nil {
		structMap["org_id"] = f.OrgId
	}
	if f.Protocol != nil {
		structMap["protocol"] = f.Protocol
	}
	if f.SamplingPercentage != nil {
		structMap["sampling_percentage"] = f.SamplingPercentage
	}
	if f.SiteId != nil {
		structMap["site_id"] = f.SiteId
	}
	if f.SrcIp != nil {
		structMap["src_ip"] = f.SrcIp
	}
	if f.SrcPort != nil {
		structMap["src_port"] = f.SrcPort
	}
	if f.StartTime != nil {
		structMap["start_time"] = f.StartTime
	}
	if f.State != nil {
		structMap["state"] = f.State
	}
	if f.Timestamp != nil {
		structMap["timestamp"] = f.Timestamp
	}
	if f.TotalBytes != nil {
		structMap["total_bytes"] = f.TotalBytes
	}
	if f.TotalPkts != nil {
		structMap["total_pkts"] = f.TotalPkts
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for FlowRecord.
// It customizes the JSON unmarshaling process for FlowRecord objects.
func (f *FlowRecord) UnmarshalJSON(input []byte) error {
	var temp tempFlowRecord
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "device_mac", "direction", "dst_ip", "dst_port", "duration", "end_time", "flow_id", "org_id", "protocol", "sampling_percentage", "site_id", "src_ip", "src_port", "start_time", "state", "timestamp", "total_bytes", "total_pkts")
	if err != nil {
		return err
	}
	f.AdditionalProperties = additionalProperties

	f.DeviceMac = temp.DeviceMac
	f.Direction = temp.Direction
	f.DstIp = temp.DstIp
	f.DstPort = temp.DstPort
	f.Duration = temp.Duration
	f.EndTime = temp.EndTime
	f.FlowId = temp.FlowId
	f.OrgId = temp.OrgId
	f.Protocol = temp.Protocol
	f.SamplingPercentage = temp.SamplingPercentage
	f.SiteId = temp.SiteId
	f.SrcIp = temp.SrcIp
	f.SrcPort = temp.SrcPort
	f.StartTime = temp.StartTime
	f.State = temp.State
	f.Timestamp = temp.Timestamp
	f.TotalBytes = temp.TotalBytes
	f.TotalPkts = temp.TotalPkts
	return nil
}

// tempFlowRecord is a temporary struct used for validating the fields of FlowRecord.
type tempFlowRecord struct {
	DeviceMac          *string                  `json:"device_mac,omitempty"`
	Direction          *FlowRecordDirectionEnum `json:"direction,omitempty"`
	DstIp              *string                  `json:"dst_ip,omitempty"`
	DstPort            *int                     `json:"dst_port,omitempty"`
	Duration           *int64                   `json:"duration,omitempty"`
	EndTime            *int64                   `json:"end_time,omitempty"`
	FlowId             *int64                   `json:"flow_id,omitempty"`
	OrgId              *uuid.UUID               `json:"org_id,omitempty"`
	Protocol           *string                  `json:"protocol,omitempty"`
	SamplingPercentage *float64                 `json:"sampling_percentage,omitempty"`
	SiteId             *uuid.UUID               `json:"site_id,omitempty"`
	SrcIp              *string                  `json:"src_ip,omitempty"`
	SrcPort            *int                     `json:"src_port,omitempty"`
	StartTime          *int64                   `json:"start_time,omitempty"`
	State              *FlowRecordStateEnum     `json:"state,omitempty"`
	Timestamp          *int64                   `json:"timestamp,omitempty"`
	TotalBytes         *int64                   `json:"total_bytes,omitempty"`
	TotalPkts          *int64                   `json:"total_pkts,omitempty"`
}
