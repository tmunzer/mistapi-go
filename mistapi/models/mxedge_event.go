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
	// Component like PS1, PS2
	Component   *string    `json:"component,omitempty"`
	MxclusterId *string    `json:"mxcluster_id,omitempty"`
	MxedgeId    *string    `json:"mxedge_id,omitempty"`
	OrgId       *uuid.UUID `json:"org_id,omitempty"`
	Service     *string    `json:"service,omitempty"`
	// Epoch (seconds)
	Timestamp            *float64               `json:"timestamp,omitempty"`
	Type                 *string                `json:"type,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for MxedgeEvent,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m MxedgeEvent) String() string {
	return fmt.Sprintf(
		"MxedgeEvent[Component=%v, MxclusterId=%v, MxedgeId=%v, OrgId=%v, Service=%v, Timestamp=%v, Type=%v, AdditionalProperties=%v]",
		m.Component, m.MxclusterId, m.MxedgeId, m.OrgId, m.Service, m.Timestamp, m.Type, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for MxedgeEvent.
// It customizes the JSON marshaling process for MxedgeEvent objects.
func (m MxedgeEvent) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(m.AdditionalProperties,
		"component", "mxcluster_id", "mxedge_id", "org_id", "service", "timestamp", "type"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(m.toMap())
}

// toMap converts the MxedgeEvent object to a map representation for JSON marshaling.
func (m MxedgeEvent) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, m.AdditionalProperties)
	if m.Component != nil {
		structMap["component"] = m.Component
	}
	if m.MxclusterId != nil {
		structMap["mxcluster_id"] = m.MxclusterId
	}
	if m.MxedgeId != nil {
		structMap["mxedge_id"] = m.MxedgeId
	}
	if m.OrgId != nil {
		structMap["org_id"] = m.OrgId
	}
	if m.Service != nil {
		structMap["service"] = m.Service
	}
	if m.Timestamp != nil {
		structMap["timestamp"] = m.Timestamp
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
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "component", "mxcluster_id", "mxedge_id", "org_id", "service", "timestamp", "type")
	if err != nil {
		return err
	}
	m.AdditionalProperties = additionalProperties

	m.Component = temp.Component
	m.MxclusterId = temp.MxclusterId
	m.MxedgeId = temp.MxedgeId
	m.OrgId = temp.OrgId
	m.Service = temp.Service
	m.Timestamp = temp.Timestamp
	m.Type = temp.Type
	return nil
}

// tempMxedgeEvent is a temporary struct used for validating the fields of MxedgeEvent.
type tempMxedgeEvent struct {
	Component   *string    `json:"component,omitempty"`
	MxclusterId *string    `json:"mxcluster_id,omitempty"`
	MxedgeId    *string    `json:"mxedge_id,omitempty"`
	OrgId       *uuid.UUID `json:"org_id,omitempty"`
	Service     *string    `json:"service,omitempty"`
	Timestamp   *float64   `json:"timestamp,omitempty"`
	Type        *string    `json:"type,omitempty"`
}
