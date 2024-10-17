package models

import (
	"encoding/json"
	"github.com/google/uuid"
)

// MxedgeEvent represents a MxedgeEvent struct.
type MxedgeEvent struct {
	// component like PS1, PS2
	Component            *string        `json:"component,omitempty"`
	MxclusterId          *string        `json:"mxcluster_id,omitempty"`
	MxedgeId             *string        `json:"mxedge_id,omitempty"`
	OrgId                *uuid.UUID     `json:"org_id,omitempty"`
	Service              *string        `json:"service,omitempty"`
	Timestamp            *float64       `json:"timestamp,omitempty"`
	Type                 *string        `json:"type,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for MxedgeEvent.
// It customizes the JSON marshaling process for MxedgeEvent objects.
func (m MxedgeEvent) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(m.toMap())
}

// toMap converts the MxedgeEvent object to a map representation for JSON marshaling.
func (m MxedgeEvent) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, m.AdditionalProperties)
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
	additionalProperties, err := UnmarshalAdditionalProperties(input, "component", "mxcluster_id", "mxedge_id", "org_id", "service", "timestamp", "type")
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
