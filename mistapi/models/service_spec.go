package models

import (
	"encoding/json"
)

// ServiceSpec represents a ServiceSpec struct.
type ServiceSpec struct {
	// port number, port range, or variable
	PortRange *string `json:"port_range,omitempty"`
	// `https`/ `tcp` / `udp` / `icmp` / `gre` / `any` / `:protocol_number`.
	// `protocol_number` is between 1-254
	Protocol             *string        `json:"protocol,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ServiceSpec.
// It customizes the JSON marshaling process for ServiceSpec objects.
func (s ServiceSpec) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(s.toMap())
}

// toMap converts the ServiceSpec object to a map representation for JSON marshaling.
func (s ServiceSpec) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, s.AdditionalProperties)
	if s.PortRange != nil {
		structMap["port_range"] = s.PortRange
	}
	if s.Protocol != nil {
		structMap["protocol"] = s.Protocol
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ServiceSpec.
// It customizes the JSON unmarshaling process for ServiceSpec objects.
func (s *ServiceSpec) UnmarshalJSON(input []byte) error {
	var temp tempServiceSpec
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "port_range", "protocol")
	if err != nil {
		return err
	}

	s.AdditionalProperties = additionalProperties
	s.PortRange = temp.PortRange
	s.Protocol = temp.Protocol
	return nil
}

// tempServiceSpec is a temporary struct used for validating the fields of ServiceSpec.
type tempServiceSpec struct {
	PortRange *string `json:"port_range,omitempty"`
	Protocol  *string `json:"protocol,omitempty"`
}
