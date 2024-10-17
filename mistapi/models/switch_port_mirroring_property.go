package models

import (
	"encoding/json"
)

// SwitchPortMirroringProperty represents a SwitchPortMirroringProperty struct.
type SwitchPortMirroringProperty struct {
	// at least one of the `input_port_ids_ingress`, `input_port_ids_egress` or `input_networks_ingress ` should be specified
	InputNetworksIngress []string `json:"input_networks_ingress,omitempty"`
	// at least one of the `input_port_ids_ingress`, `input_port_ids_egress` or `input_networks_ingress ` should be specified
	InputPortIdsEgress []string `json:"input_port_ids_egress,omitempty"`
	// at least one of the `input_port_ids_ingress`, `input_port_ids_egress` or `input_networks_ingress ` should be specified
	InputPortIdsIngress []string `json:"input_port_ids_ingress,omitempty"`
	// exaclty one of the `output_port_id` or `output_network` should be provided
	OutputNetwork *string `json:"output_network,omitempty"`
	// exaclty one of the `output_port_id` or `output_network` should be provided
	OutputPortId         *string        `json:"output_port_id,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SwitchPortMirroringProperty.
// It customizes the JSON marshaling process for SwitchPortMirroringProperty objects.
func (s SwitchPortMirroringProperty) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(s.toMap())
}

// toMap converts the SwitchPortMirroringProperty object to a map representation for JSON marshaling.
func (s SwitchPortMirroringProperty) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, s.AdditionalProperties)
	if s.InputNetworksIngress != nil {
		structMap["input_networks_ingress"] = s.InputNetworksIngress
	}
	if s.InputPortIdsEgress != nil {
		structMap["input_port_ids_egress"] = s.InputPortIdsEgress
	}
	if s.InputPortIdsIngress != nil {
		structMap["input_port_ids_ingress"] = s.InputPortIdsIngress
	}
	if s.OutputNetwork != nil {
		structMap["output_network"] = s.OutputNetwork
	}
	if s.OutputPortId != nil {
		structMap["output_port_id"] = s.OutputPortId
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SwitchPortMirroringProperty.
// It customizes the JSON unmarshaling process for SwitchPortMirroringProperty objects.
func (s *SwitchPortMirroringProperty) UnmarshalJSON(input []byte) error {
	var temp tempSwitchPortMirroringProperty
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "input_networks_ingress", "input_port_ids_egress", "input_port_ids_ingress", "output_network", "output_port_id")
	if err != nil {
		return err
	}

	s.AdditionalProperties = additionalProperties
	s.InputNetworksIngress = temp.InputNetworksIngress
	s.InputPortIdsEgress = temp.InputPortIdsEgress
	s.InputPortIdsIngress = temp.InputPortIdsIngress
	s.OutputNetwork = temp.OutputNetwork
	s.OutputPortId = temp.OutputPortId
	return nil
}

// tempSwitchPortMirroringProperty is a temporary struct used for validating the fields of SwitchPortMirroringProperty.
type tempSwitchPortMirroringProperty struct {
	InputNetworksIngress []string `json:"input_networks_ingress,omitempty"`
	InputPortIdsEgress   []string `json:"input_port_ids_egress,omitempty"`
	InputPortIdsIngress  []string `json:"input_port_ids_ingress,omitempty"`
	OutputNetwork        *string  `json:"output_network,omitempty"`
	OutputPortId         *string  `json:"output_port_id,omitempty"`
}
