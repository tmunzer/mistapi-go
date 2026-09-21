// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
)

// ScepEvent represents a ScepEvent struct.
// Mist SCEP PKI operation event reported for an organization
type ScepEvent struct {
	// MDM or certificate provider that triggered the SCEP operation
	CertProvider *string `json:"cert_provider,omitempty"`
	// Common name presented in the SCEP certificate request
	CommonName *string `json:"common_name,omitempty"`
	// Device identifier associated with the SCEP operation. Empty when the SCEP request did not include a device ID
	DeviceId *string `json:"device_id,omitempty"`
	// Unique identifier of a Mist organization
	OrgId *uuid.UUID `json:"org_id,omitempty"`
	// Reason text describing the outcome of the SCEP operation
	Text *string `json:"text,omitempty"`
	// Epoch timestamp, in seconds
	Timestamp *float64 `json:"timestamp,omitempty"`
	// enum: `SCEP_PKI_OPERATION_FAILURE`, `SCEP_PKI_OPERATION_SUCCESS`
	Type                 *ScepEventTypeEnum     `json:"type,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ScepEvent,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s ScepEvent) String() string {
	return fmt.Sprintf(
		"ScepEvent[CertProvider=%v, CommonName=%v, DeviceId=%v, OrgId=%v, Text=%v, Timestamp=%v, Type=%v, AdditionalProperties=%v]",
		s.CertProvider, s.CommonName, s.DeviceId, s.OrgId, s.Text, s.Timestamp, s.Type, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ScepEvent.
// It customizes the JSON marshaling process for ScepEvent objects.
func (s ScepEvent) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"cert_provider", "common_name", "device_id", "org_id", "text", "timestamp", "type"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the ScepEvent object to a map representation for JSON marshaling.
func (s ScepEvent) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.CertProvider != nil {
		structMap["cert_provider"] = s.CertProvider
	}
	if s.CommonName != nil {
		structMap["common_name"] = s.CommonName
	}
	if s.DeviceId != nil {
		structMap["device_id"] = s.DeviceId
	}
	if s.OrgId != nil {
		structMap["org_id"] = s.OrgId
	}
	if s.Text != nil {
		structMap["text"] = s.Text
	}
	if s.Timestamp != nil {
		structMap["timestamp"] = s.Timestamp
	}
	if s.Type != nil {
		structMap["type"] = s.Type
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ScepEvent.
// It customizes the JSON unmarshaling process for ScepEvent objects.
func (s *ScepEvent) UnmarshalJSON(input []byte) error {
	var temp tempScepEvent
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "cert_provider", "common_name", "device_id", "org_id", "text", "timestamp", "type")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.CertProvider = temp.CertProvider
	s.CommonName = temp.CommonName
	s.DeviceId = temp.DeviceId
	s.OrgId = temp.OrgId
	s.Text = temp.Text
	s.Timestamp = temp.Timestamp
	s.Type = temp.Type
	return nil
}

// tempScepEvent is a temporary struct used for validating the fields of ScepEvent.
type tempScepEvent struct {
	CertProvider *string            `json:"cert_provider,omitempty"`
	CommonName   *string            `json:"common_name,omitempty"`
	DeviceId     *string            `json:"device_id,omitempty"`
	OrgId        *uuid.UUID         `json:"org_id,omitempty"`
	Text         *string            `json:"text,omitempty"`
	Timestamp    *float64           `json:"timestamp,omitempty"`
	Type         *ScepEventTypeEnum `json:"type,omitempty"`
}
