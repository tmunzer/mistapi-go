// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
)

// WebhookMinisReachabilityEvent represents a WebhookMinisReachabilityEvent struct.
type WebhookMinisReachabilityEvent struct {
	// Average latency in milliseconds
	AvgLatency *float64 `json:"avg_latency,omitempty"`
	// MAC address of the device performing the test
	DeviceMac *string `json:"device_mac,omitempty"`
	// Packet loss percentage
	LossPercentage *float64 `json:"loss_percentage,omitempty"`
	// Maximum latency in milliseconds
	MaxLatency *float64 `json:"max_latency,omitempty"`
	// Minimum latency in milliseconds
	MinLatency *float64   `json:"min_latency,omitempty"`
	OrgId      *uuid.UUID `json:"org_id,omitempty"`
	// Name of the probe
	ProbeName *string `json:"probe_name,omitempty"`
	// Target host or IP for the probe
	ProbeTarget *string `json:"probe_target,omitempty"`
	// Type of probe
	ProbeType *string `json:"probe_type,omitempty"`
	// Protocol used for the test
	Protocol *string    `json:"protocol,omitempty"`
	SiteId   *uuid.UUID `json:"site_id,omitempty"`
	// Type of test performed
	TestType *string `json:"test_type,omitempty"`
	// Epoch (seconds)
	Timestamp *float64 `json:"timestamp,omitempty"`
	// VLAN ID used for the test
	Vlan                 *int                   `json:"vlan,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for WebhookMinisReachabilityEvent,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WebhookMinisReachabilityEvent) String() string {
	return fmt.Sprintf(
		"WebhookMinisReachabilityEvent[AvgLatency=%v, DeviceMac=%v, LossPercentage=%v, MaxLatency=%v, MinLatency=%v, OrgId=%v, ProbeName=%v, ProbeTarget=%v, ProbeType=%v, Protocol=%v, SiteId=%v, TestType=%v, Timestamp=%v, Vlan=%v, AdditionalProperties=%v]",
		w.AvgLatency, w.DeviceMac, w.LossPercentage, w.MaxLatency, w.MinLatency, w.OrgId, w.ProbeName, w.ProbeTarget, w.ProbeType, w.Protocol, w.SiteId, w.TestType, w.Timestamp, w.Vlan, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WebhookMinisReachabilityEvent.
// It customizes the JSON marshaling process for WebhookMinisReachabilityEvent objects.
func (w WebhookMinisReachabilityEvent) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(w.AdditionalProperties,
		"avg_latency", "device_mac", "loss_percentage", "max_latency", "min_latency", "org_id", "probe_name", "probe_target", "probe_type", "protocol", "site_id", "test_type", "timestamp", "vlan"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(w.toMap())
}

// toMap converts the WebhookMinisReachabilityEvent object to a map representation for JSON marshaling.
func (w WebhookMinisReachabilityEvent) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, w.AdditionalProperties)
	if w.AvgLatency != nil {
		structMap["avg_latency"] = w.AvgLatency
	}
	if w.DeviceMac != nil {
		structMap["device_mac"] = w.DeviceMac
	}
	if w.LossPercentage != nil {
		structMap["loss_percentage"] = w.LossPercentage
	}
	if w.MaxLatency != nil {
		structMap["max_latency"] = w.MaxLatency
	}
	if w.MinLatency != nil {
		structMap["min_latency"] = w.MinLatency
	}
	if w.OrgId != nil {
		structMap["org_id"] = w.OrgId
	}
	if w.ProbeName != nil {
		structMap["probe_name"] = w.ProbeName
	}
	if w.ProbeTarget != nil {
		structMap["probe_target"] = w.ProbeTarget
	}
	if w.ProbeType != nil {
		structMap["probe_type"] = w.ProbeType
	}
	if w.Protocol != nil {
		structMap["protocol"] = w.Protocol
	}
	if w.SiteId != nil {
		structMap["site_id"] = w.SiteId
	}
	if w.TestType != nil {
		structMap["test_type"] = w.TestType
	}
	if w.Timestamp != nil {
		structMap["timestamp"] = w.Timestamp
	}
	if w.Vlan != nil {
		structMap["vlan"] = w.Vlan
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookMinisReachabilityEvent.
// It customizes the JSON unmarshaling process for WebhookMinisReachabilityEvent objects.
func (w *WebhookMinisReachabilityEvent) UnmarshalJSON(input []byte) error {
	var temp tempWebhookMinisReachabilityEvent
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "avg_latency", "device_mac", "loss_percentage", "max_latency", "min_latency", "org_id", "probe_name", "probe_target", "probe_type", "protocol", "site_id", "test_type", "timestamp", "vlan")
	if err != nil {
		return err
	}
	w.AdditionalProperties = additionalProperties

	w.AvgLatency = temp.AvgLatency
	w.DeviceMac = temp.DeviceMac
	w.LossPercentage = temp.LossPercentage
	w.MaxLatency = temp.MaxLatency
	w.MinLatency = temp.MinLatency
	w.OrgId = temp.OrgId
	w.ProbeName = temp.ProbeName
	w.ProbeTarget = temp.ProbeTarget
	w.ProbeType = temp.ProbeType
	w.Protocol = temp.Protocol
	w.SiteId = temp.SiteId
	w.TestType = temp.TestType
	w.Timestamp = temp.Timestamp
	w.Vlan = temp.Vlan
	return nil
}

// tempWebhookMinisReachabilityEvent is a temporary struct used for validating the fields of WebhookMinisReachabilityEvent.
type tempWebhookMinisReachabilityEvent struct {
	AvgLatency     *float64   `json:"avg_latency,omitempty"`
	DeviceMac      *string    `json:"device_mac,omitempty"`
	LossPercentage *float64   `json:"loss_percentage,omitempty"`
	MaxLatency     *float64   `json:"max_latency,omitempty"`
	MinLatency     *float64   `json:"min_latency,omitempty"`
	OrgId          *uuid.UUID `json:"org_id,omitempty"`
	ProbeName      *string    `json:"probe_name,omitempty"`
	ProbeTarget    *string    `json:"probe_target,omitempty"`
	ProbeType      *string    `json:"probe_type,omitempty"`
	Protocol       *string    `json:"protocol,omitempty"`
	SiteId         *uuid.UUID `json:"site_id,omitempty"`
	TestType       *string    `json:"test_type,omitempty"`
	Timestamp      *float64   `json:"timestamp,omitempty"`
	Vlan           *int       `json:"vlan,omitempty"`
}
