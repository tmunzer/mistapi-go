package models

import (
	"encoding/json"
)

// StatsSwitchApRedundancyModule represents a StatsSwitchApRedundancyModule struct.
type StatsSwitchApRedundancyModule struct {
	NumAps                     *int           `json:"num_aps,omitempty"`
	NumApsWithSwitchRedundancy *int           `json:"num_aps_with_switch_redundancy,omitempty"`
	AdditionalProperties       map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for StatsSwitchApRedundancyModule.
// It customizes the JSON marshaling process for StatsSwitchApRedundancyModule objects.
func (s StatsSwitchApRedundancyModule) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(s.toMap())
}

// toMap converts the StatsSwitchApRedundancyModule object to a map representation for JSON marshaling.
func (s StatsSwitchApRedundancyModule) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, s.AdditionalProperties)
	if s.NumAps != nil {
		structMap["num_aps"] = s.NumAps
	}
	if s.NumApsWithSwitchRedundancy != nil {
		structMap["num_aps_with_switch_redundancy"] = s.NumApsWithSwitchRedundancy
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsSwitchApRedundancyModule.
// It customizes the JSON unmarshaling process for StatsSwitchApRedundancyModule objects.
func (s *StatsSwitchApRedundancyModule) UnmarshalJSON(input []byte) error {
	var temp tempStatsSwitchApRedundancyModule
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "num_aps", "num_aps_with_switch_redundancy")
	if err != nil {
		return err
	}

	s.AdditionalProperties = additionalProperties
	s.NumAps = temp.NumAps
	s.NumApsWithSwitchRedundancy = temp.NumApsWithSwitchRedundancy
	return nil
}

// tempStatsSwitchApRedundancyModule is a temporary struct used for validating the fields of StatsSwitchApRedundancyModule.
type tempStatsSwitchApRedundancyModule struct {
	NumAps                     *int `json:"num_aps,omitempty"`
	NumApsWithSwitchRedundancy *int `json:"num_aps_with_switch_redundancy,omitempty"`
}
