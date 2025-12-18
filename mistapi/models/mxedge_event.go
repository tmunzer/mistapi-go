// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
)

// MxedgeEvent represents a MxedgeEvent struct.
type MxedgeEvent struct {
	// Unique ID of the object instance in the Mist Organization
	AuditId   *uuid.UUID       `json:"audit_id,omitempty"`
	Component Optional[string] `json:"component"`
	// Device id
	DeviceId     Optional[uuid.UUID] `json:"device_id"`
	DeviceType   *string             `json:"device_type,omitempty"`
	FromVersion  *string             `json:"from_version,omitempty"`
	Mac          *string             `json:"mac,omitempty"`
	MxclusterId  *string             `json:"mxcluster_id,omitempty"`
	MxedgeId     *string             `json:"mxedge_id,omitempty"`
	MxedgeName   *string             `json:"mxedge_name,omitempty"`
	OrgId        *uuid.UUID          `json:"org_id,omitempty"`
	Package      *string             `json:"package,omitempty"`
	Service      *string             `json:"service,omitempty"`
	Severity     *EventSeverityEnum  `json:"severity,omitempty"`
	SysInfoUsage *MxedgeEventSysInfo `json:"sys_info.usage,omitempty"`
	Text         *string             `json:"text,omitempty"`
	// Epoch (seconds)
	Timestamp            *float64               `json:"timestamp,omitempty"`
	ToVersion            *string                `json:"to_version,omitempty"`
	Type                 *string                `json:"type,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for MxedgeEvent,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m MxedgeEvent) String() string {
	return fmt.Sprintf(
		"MxedgeEvent[AuditId=%v, Component=%v, DeviceId=%v, DeviceType=%v, FromVersion=%v, Mac=%v, MxclusterId=%v, MxedgeId=%v, MxedgeName=%v, OrgId=%v, Package=%v, Service=%v, Severity=%v, SysInfoUsage=%v, Text=%v, Timestamp=%v, ToVersion=%v, Type=%v, AdditionalProperties=%v]",
		m.AuditId, m.Component, m.DeviceId, m.DeviceType, m.FromVersion, m.Mac, m.MxclusterId, m.MxedgeId, m.MxedgeName, m.OrgId, m.Package, m.Service, m.Severity, m.SysInfoUsage, m.Text, m.Timestamp, m.ToVersion, m.Type, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for MxedgeEvent.
// It customizes the JSON marshaling process for MxedgeEvent objects.
func (m MxedgeEvent) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(m.AdditionalProperties,
		"audit_id", "component", "device_id", "device_type", "from_version", "mac", "mxcluster_id", "mxedge_id", "mxedge_name", "org_id", "package", "service", "severity", "sys_info.usage", "text", "timestamp", "to_version", "type"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(m.toMap())
}

// toMap converts the MxedgeEvent object to a map representation for JSON marshaling.
func (m MxedgeEvent) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, m.AdditionalProperties)
	if m.AuditId != nil {
		structMap["audit_id"] = m.AuditId
	}
	if m.Component.IsValueSet() {
		if m.Component.Value() != nil {
			structMap["component"] = m.Component.Value()
		} else {
			structMap["component"] = nil
		}
	}
	if m.DeviceId.IsValueSet() {
		if m.DeviceId.Value() != nil {
			structMap["device_id"] = m.DeviceId.Value()
		} else {
			structMap["device_id"] = nil
		}
	}
	if m.DeviceType != nil {
		structMap["device_type"] = m.DeviceType
	}
	if m.FromVersion != nil {
		structMap["from_version"] = m.FromVersion
	}
	if m.Mac != nil {
		structMap["mac"] = m.Mac
	}
	if m.MxclusterId != nil {
		structMap["mxcluster_id"] = m.MxclusterId
	}
	if m.MxedgeId != nil {
		structMap["mxedge_id"] = m.MxedgeId
	}
	if m.MxedgeName != nil {
		structMap["mxedge_name"] = m.MxedgeName
	}
	if m.OrgId != nil {
		structMap["org_id"] = m.OrgId
	}
	if m.Package != nil {
		structMap["package"] = m.Package
	}
	if m.Service != nil {
		structMap["service"] = m.Service
	}
	if m.Severity != nil {
		structMap["severity"] = m.Severity
	}
	if m.SysInfoUsage != nil {
		structMap["sys_info.usage"] = m.SysInfoUsage.toMap()
	}
	if m.Text != nil {
		structMap["text"] = m.Text
	}
	if m.Timestamp != nil {
		structMap["timestamp"] = m.Timestamp
	}
	if m.ToVersion != nil {
		structMap["to_version"] = m.ToVersion
	}
	if m.Type != nil {
		structMap["type"] = m.Type
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MxedgeEvent.
// It customizes the JSON unmarshaling process for MxedgeEvent objects.
func (m *MxedgeEvent) UnmarshalJSON(input []byte) error {
	var temp tempMxedgeEvent
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "audit_id", "component", "device_id", "device_type", "from_version", "mac", "mxcluster_id", "mxedge_id", "mxedge_name", "org_id", "package", "service", "severity", "sys_info.usage", "text", "timestamp", "to_version", "type")
	if err != nil {
		return err
	}
	m.AdditionalProperties = additionalProperties

	m.AuditId = temp.AuditId
	m.Component = temp.Component
	m.DeviceId = temp.DeviceId
	m.DeviceType = temp.DeviceType
	m.FromVersion = temp.FromVersion
	m.Mac = temp.Mac
	m.MxclusterId = temp.MxclusterId
	m.MxedgeId = temp.MxedgeId
	m.MxedgeName = temp.MxedgeName
	m.OrgId = temp.OrgId
	m.Package = temp.Package
	m.Service = temp.Service
	m.Severity = temp.Severity
	m.SysInfoUsage = temp.SysInfoUsage
	m.Text = temp.Text
	m.Timestamp = temp.Timestamp
	m.ToVersion = temp.ToVersion
	m.Type = temp.Type
	return nil
}

// tempMxedgeEvent is a temporary struct used for validating the fields of MxedgeEvent.
type tempMxedgeEvent struct {
	AuditId      *uuid.UUID          `json:"audit_id,omitempty"`
	Component    Optional[string]    `json:"component"`
	DeviceId     Optional[uuid.UUID] `json:"device_id"`
	DeviceType   *string             `json:"device_type,omitempty"`
	FromVersion  *string             `json:"from_version,omitempty"`
	Mac          *string             `json:"mac,omitempty"`
	MxclusterId  *string             `json:"mxcluster_id,omitempty"`
	MxedgeId     *string             `json:"mxedge_id,omitempty"`
	MxedgeName   *string             `json:"mxedge_name,omitempty"`
	OrgId        *uuid.UUID          `json:"org_id,omitempty"`
	Package      *string             `json:"package,omitempty"`
	Service      *string             `json:"service,omitempty"`
	Severity     *EventSeverityEnum  `json:"severity,omitempty"`
	SysInfoUsage *MxedgeEventSysInfo `json:"sys_info.usage,omitempty"`
	Text         *string             `json:"text,omitempty"`
	Timestamp    *float64            `json:"timestamp,omitempty"`
	ToVersion    *string             `json:"to_version,omitempty"`
	Type         *string             `json:"type,omitempty"`
}
