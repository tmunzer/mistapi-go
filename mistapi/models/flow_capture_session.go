// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
)

// FlowCaptureSession represents a FlowCaptureSession struct.
// Current flow capture session status for a site
type FlowCaptureSession struct {
	// Normalized filter applied to a flow capture session
	CaptureFilter *FlowCaptureFilter `json:"capture_filter,omitempty"`
	Duration      *int               `json:"duration,omitempty"`
	Enabled       *bool              `json:"enabled,omitempty"`
	Expiry        *int               `json:"expiry,omitempty"`
	Id            *uuid.UUID         `json:"id,omitempty"`
	// Switches that failed flow capture validation
	InvalidSwitches *interface{} `json:"invalid_switches,omitempty"`
	// Unique identifier of a Mist organization
	OrgId *uuid.UUID `json:"org_id,omitempty"`
	// Unique identifier of a Mist site
	SiteId               *uuid.UUID             `json:"site_id,omitempty"`
	SwitchCount          *int                   `json:"switch_count,omitempty"`
	Timestamp            *int                   `json:"timestamp,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for FlowCaptureSession,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (f FlowCaptureSession) String() string {
	return fmt.Sprintf(
		"FlowCaptureSession[CaptureFilter=%v, Duration=%v, Enabled=%v, Expiry=%v, Id=%v, InvalidSwitches=%v, OrgId=%v, SiteId=%v, SwitchCount=%v, Timestamp=%v, AdditionalProperties=%v]",
		f.CaptureFilter, f.Duration, f.Enabled, f.Expiry, f.Id, f.InvalidSwitches, f.OrgId, f.SiteId, f.SwitchCount, f.Timestamp, f.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for FlowCaptureSession.
// It customizes the JSON marshaling process for FlowCaptureSession objects.
func (f FlowCaptureSession) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(f.AdditionalProperties,
		"capture_filter", "duration", "enabled", "expiry", "id", "invalid_switches", "org_id", "site_id", "switch_count", "timestamp"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(f.toMap())
}

// toMap converts the FlowCaptureSession object to a map representation for JSON marshaling.
func (f FlowCaptureSession) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, f.AdditionalProperties)
	if f.CaptureFilter != nil {
		structMap["capture_filter"] = f.CaptureFilter.toMap()
	}
	if f.Duration != nil {
		structMap["duration"] = f.Duration
	}
	if f.Enabled != nil {
		structMap["enabled"] = f.Enabled
	}
	if f.Expiry != nil {
		structMap["expiry"] = f.Expiry
	}
	if f.Id != nil {
		structMap["id"] = f.Id
	}
	if f.InvalidSwitches != nil {
		structMap["invalid_switches"] = f.InvalidSwitches
	}
	if f.OrgId != nil {
		structMap["org_id"] = f.OrgId
	}
	if f.SiteId != nil {
		structMap["site_id"] = f.SiteId
	}
	if f.SwitchCount != nil {
		structMap["switch_count"] = f.SwitchCount
	}
	if f.Timestamp != nil {
		structMap["timestamp"] = f.Timestamp
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for FlowCaptureSession.
// It customizes the JSON unmarshaling process for FlowCaptureSession objects.
func (f *FlowCaptureSession) UnmarshalJSON(input []byte) error {
	var temp tempFlowCaptureSession
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "capture_filter", "duration", "enabled", "expiry", "id", "invalid_switches", "org_id", "site_id", "switch_count", "timestamp")
	if err != nil {
		return err
	}
	f.AdditionalProperties = additionalProperties

	f.CaptureFilter = temp.CaptureFilter
	f.Duration = temp.Duration
	f.Enabled = temp.Enabled
	f.Expiry = temp.Expiry
	f.Id = temp.Id
	f.InvalidSwitches = temp.InvalidSwitches
	f.OrgId = temp.OrgId
	f.SiteId = temp.SiteId
	f.SwitchCount = temp.SwitchCount
	f.Timestamp = temp.Timestamp
	return nil
}

// tempFlowCaptureSession is a temporary struct used for validating the fields of FlowCaptureSession.
type tempFlowCaptureSession struct {
	CaptureFilter   *FlowCaptureFilter `json:"capture_filter,omitempty"`
	Duration        *int               `json:"duration,omitempty"`
	Enabled         *bool              `json:"enabled,omitempty"`
	Expiry          *int               `json:"expiry,omitempty"`
	Id              *uuid.UUID         `json:"id,omitempty"`
	InvalidSwitches *interface{}       `json:"invalid_switches,omitempty"`
	OrgId           *uuid.UUID         `json:"org_id,omitempty"`
	SiteId          *uuid.UUID         `json:"site_id,omitempty"`
	SwitchCount     *int               `json:"switch_count,omitempty"`
	Timestamp       *int               `json:"timestamp,omitempty"`
}
