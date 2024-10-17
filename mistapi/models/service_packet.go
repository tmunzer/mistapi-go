package models

import (
	"encoding/json"
)

// ServicePacket represents a ServicePacket struct.
type ServicePacket struct {
	// ata from service data
	ServiceData *string `json:"service_data,omitempty"`
	// UUID from service data
	ServiceUuid          *string        `json:"service_uuid,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ServicePacket.
// It customizes the JSON marshaling process for ServicePacket objects.
func (s ServicePacket) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(s.toMap())
}

// toMap converts the ServicePacket object to a map representation for JSON marshaling.
func (s ServicePacket) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, s.AdditionalProperties)
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
	additionalProperties, err := UnmarshalAdditionalProperties(input, "service_data", "service_uuid")
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
