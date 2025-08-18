// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// ServicePacket represents a ServicePacket struct.
type ServicePacket struct {
	// ata from service data
	ServiceData *string `json:"service_data,omitempty"`
	// UUID from service data
	ServiceUuid          *string                `json:"service_uuid,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ServicePacket,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s ServicePacket) String() string {
	return fmt.Sprintf(
		"ServicePacket[ServiceData=%v, ServiceUuid=%v, AdditionalProperties=%v]",
		s.ServiceData, s.ServiceUuid, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ServicePacket.
// It customizes the JSON marshaling process for ServicePacket objects.
func (s ServicePacket) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"service_data", "service_uuid"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the ServicePacket object to a map representation for JSON marshaling.
func (s ServicePacket) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.ServiceData != nil {
		structMap["service_data"] = s.ServiceData
	}
	if s.ServiceUuid != nil {
		structMap["service_uuid"] = s.ServiceUuid
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ServicePacket.
// It customizes the JSON unmarshaling process for ServicePacket objects.
func (s *ServicePacket) UnmarshalJSON(input []byte) error {
	var temp tempServicePacket
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "service_data", "service_uuid")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.ServiceData = temp.ServiceData
	s.ServiceUuid = temp.ServiceUuid
	return nil
}

// tempServicePacket is a temporary struct used for validating the fields of ServicePacket.
type tempServicePacket struct {
	ServiceData *string `json:"service_data,omitempty"`
	ServiceUuid *string `json:"service_uuid,omitempty"`
}
