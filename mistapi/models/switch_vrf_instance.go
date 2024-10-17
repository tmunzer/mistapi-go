package models

import (
	"encoding/json"
)

// SwitchVrfInstance represents a SwitchVrfInstance struct.
type SwitchVrfInstance struct {
	// Property key is the destination CIDR (e.g. "10.0.0.0/8")
	ExtraRoutes          map[string]VrfExtraRoute `json:"extra_routes,omitempty"`
	Networks             []string                 `json:"networks,omitempty"`
	AdditionalProperties map[string]any           `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SwitchVrfInstance.
// It customizes the JSON marshaling process for SwitchVrfInstance objects.
func (s SwitchVrfInstance) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(s.toMap())
}

// toMap converts the SwitchVrfInstance object to a map representation for JSON marshaling.
func (s SwitchVrfInstance) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, s.AdditionalProperties)
	if s.ExtraRoutes != nil {
		structMap["extra_routes"] = s.ExtraRoutes
	}
	if s.Networks != nil {
		structMap["networks"] = s.Networks
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SwitchVrfInstance.
// It customizes the JSON unmarshaling process for SwitchVrfInstance objects.
func (s *SwitchVrfInstance) UnmarshalJSON(input []byte) error {
	var temp tempSwitchVrfInstance
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "extra_routes", "networks")
	if err != nil {
		return err
	}

	s.AdditionalProperties = additionalProperties
	s.ExtraRoutes = temp.ExtraRoutes
	s.Networks = temp.Networks
	return nil
}

// tempSwitchVrfInstance is a temporary struct used for validating the fields of SwitchVrfInstance.
type tempSwitchVrfInstance struct {
	ExtraRoutes map[string]VrfExtraRoute `json:"extra_routes,omitempty"`
	Networks    []string                 `json:"networks,omitempty"`
}
