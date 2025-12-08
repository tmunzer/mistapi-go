// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// MxedgeEventSysInfo represents a MxedgeEventSysInfo struct.
type MxedgeEventSysInfo struct {
	Resource             *string                `json:"resource,omitempty"`
	Severity             *EventSeverityEnum     `json:"severity,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for MxedgeEventSysInfo,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m MxedgeEventSysInfo) String() string {
	return fmt.Sprintf(
		"MxedgeEventSysInfo[Resource=%v, Severity=%v, AdditionalProperties=%v]",
		m.Resource, m.Severity, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for MxedgeEventSysInfo.
// It customizes the JSON marshaling process for MxedgeEventSysInfo objects.
func (m MxedgeEventSysInfo) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(m.AdditionalProperties,
		"resource", "severity"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(m.toMap())
}

// toMap converts the MxedgeEventSysInfo object to a map representation for JSON marshaling.
func (m MxedgeEventSysInfo) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, m.AdditionalProperties)
	if m.Resource != nil {
		structMap["resource"] = m.Resource
	}
	if m.Severity != nil {
		structMap["severity"] = m.Severity
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MxedgeEventSysInfo.
// It customizes the JSON unmarshaling process for MxedgeEventSysInfo objects.
func (m *MxedgeEventSysInfo) UnmarshalJSON(input []byte) error {
	var temp tempMxedgeEventSysInfo
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "resource", "severity")
	if err != nil {
		return err
	}
	m.AdditionalProperties = additionalProperties

	m.Resource = temp.Resource
	m.Severity = temp.Severity
	return nil
}

// tempMxedgeEventSysInfo is a temporary struct used for validating the fields of MxedgeEventSysInfo.
type tempMxedgeEventSysInfo struct {
	Resource *string            `json:"resource,omitempty"`
	Severity *EventSeverityEnum `json:"severity,omitempty"`
}
