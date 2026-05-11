// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
)

// WebhookMinisApplicationEvent represents a WebhookMinisApplicationEvent struct.
type WebhookMinisApplicationEvent struct {
	// MAC address of the device
	DeviceMac *string `json:"device_mac,omitempty"`
	// IP address test was performed to
	Ip *string `json:"ip,omitempty"`
	// latency in milliseconds
	Latency *int       `json:"latency,omitempty"`
	OrgId   *uuid.UUID `json:"org_id,omitempty"`
	// Name of the probe
	ProbeName *string `json:"probe_name,omitempty"`
	// Type of probe
	ProbeType *string    `json:"probe_type,omitempty"`
	SiteId    *uuid.UUID `json:"site_id,omitempty"`
	// Source IP address of the test
	SrcIp *string `json:"src_ip,omitempty"`
	// Whether the test was successful
	Success *bool `json:"success,omitempty"`
	// enum: `application`, `curl`, `icmp`, `reachability`, `tcp`
	TestType *SynthetictestConfigCustomProbeTypeEnum `json:"test_type,omitempty"`
	// Epoch (seconds)
	Timestamp *float64 `json:"timestamp,omitempty"`
	// VLAN ID used for the test
	Vlan                 *int                   `json:"vlan,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for WebhookMinisApplicationEvent,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WebhookMinisApplicationEvent) String() string {
	return fmt.Sprintf(
		"WebhookMinisApplicationEvent[DeviceMac=%v, Ip=%v, Latency=%v, OrgId=%v, ProbeName=%v, ProbeType=%v, SiteId=%v, SrcIp=%v, Success=%v, TestType=%v, Timestamp=%v, Vlan=%v, AdditionalProperties=%v]",
		w.DeviceMac, w.Ip, w.Latency, w.OrgId, w.ProbeName, w.ProbeType, w.SiteId, w.SrcIp, w.Success, w.TestType, w.Timestamp, w.Vlan, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WebhookMinisApplicationEvent.
// It customizes the JSON marshaling process for WebhookMinisApplicationEvent objects.
func (w WebhookMinisApplicationEvent) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(w.AdditionalProperties,
		"device_mac", "ip", "latency", "org_id", "probe_name", "probe_type", "site_id", "src_ip", "success", "test_type", "timestamp", "vlan"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(w.toMap())
}

// toMap converts the WebhookMinisApplicationEvent object to a map representation for JSON marshaling.
func (w WebhookMinisApplicationEvent) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, w.AdditionalProperties)
	if w.DeviceMac != nil {
		structMap["device_mac"] = w.DeviceMac
	}
	if w.Ip != nil {
		structMap["ip"] = w.Ip
	}
	if w.Latency != nil {
		structMap["latency"] = w.Latency
	}
	if w.OrgId != nil {
		structMap["org_id"] = w.OrgId
	}
	if w.ProbeName != nil {
		structMap["probe_name"] = w.ProbeName
	}
	if w.ProbeType != nil {
		structMap["probe_type"] = w.ProbeType
	}
	if w.SiteId != nil {
		structMap["site_id"] = w.SiteId
	}
	if w.SrcIp != nil {
		structMap["src_ip"] = w.SrcIp
	}
	if w.Success != nil {
		structMap["success"] = w.Success
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

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookMinisApplicationEvent.
// It customizes the JSON unmarshaling process for WebhookMinisApplicationEvent objects.
func (w *WebhookMinisApplicationEvent) UnmarshalJSON(input []byte) error {
	var temp tempWebhookMinisApplicationEvent
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "device_mac", "ip", "latency", "org_id", "probe_name", "probe_type", "site_id", "src_ip", "success", "test_type", "timestamp", "vlan")
	if err != nil {
		return err
	}
	w.AdditionalProperties = additionalProperties

	w.DeviceMac = temp.DeviceMac
	w.Ip = temp.Ip
	w.Latency = temp.Latency
	w.OrgId = temp.OrgId
	w.ProbeName = temp.ProbeName
	w.ProbeType = temp.ProbeType
	w.SiteId = temp.SiteId
	w.SrcIp = temp.SrcIp
	w.Success = temp.Success
	w.TestType = temp.TestType
	w.Timestamp = temp.Timestamp
	w.Vlan = temp.Vlan
	return nil
}

// tempWebhookMinisApplicationEvent is a temporary struct used for validating the fields of WebhookMinisApplicationEvent.
type tempWebhookMinisApplicationEvent struct {
	DeviceMac *string                                 `json:"device_mac,omitempty"`
	Ip        *string                                 `json:"ip,omitempty"`
	Latency   *int                                    `json:"latency,omitempty"`
	OrgId     *uuid.UUID                              `json:"org_id,omitempty"`
	ProbeName *string                                 `json:"probe_name,omitempty"`
	ProbeType *string                                 `json:"probe_type,omitempty"`
	SiteId    *uuid.UUID                              `json:"site_id,omitempty"`
	SrcIp     *string                                 `json:"src_ip,omitempty"`
	Success   *bool                                   `json:"success,omitempty"`
	TestType  *SynthetictestConfigCustomProbeTypeEnum `json:"test_type,omitempty"`
	Timestamp *float64                                `json:"timestamp,omitempty"`
	Vlan      *int                                    `json:"vlan,omitempty"`
}
