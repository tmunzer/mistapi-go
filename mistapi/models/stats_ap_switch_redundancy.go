package models

import (
	"encoding/json"
)

// StatsApSwitchRedundancy represents a StatsApSwitchRedundancy struct.
type StatsApSwitchRedundancy struct {
	NumRedundantAps      Optional[int]  `json:"num_redundant_aps"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for StatsApSwitchRedundancy.
// It customizes the JSON marshaling process for StatsApSwitchRedundancy objects.
func (s StatsApSwitchRedundancy) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(s.toMap())
}

// toMap converts the StatsApSwitchRedundancy object to a map representation for JSON marshaling.
func (s StatsApSwitchRedundancy) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, s.AdditionalProperties)
	if s.NumRedundantAps.IsValueSet() {
		if s.NumRedundantAps.Value() != nil {
			structMap["num_redundant_aps"] = s.NumRedundantAps.Value()
		} else {
			structMap["num_redundant_aps"] = nil
		}
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsApSwitchRedundancy.
// It customizes the JSON unmarshaling process for StatsApSwitchRedundancy objects.
func (s *StatsApSwitchRedundancy) UnmarshalJSON(input []byte) error {
	var temp tempStatsApSwitchRedundancy
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "num_redundant_aps")
	if err != nil {
		return err
	}

	s.AdditionalProperties = additionalProperties
	s.NumRedundantAps = temp.NumRedundantAps
	return nil
}

// tempStatsApSwitchRedundancy is a temporary struct used for validating the fields of StatsApSwitchRedundancy.
type tempStatsApSwitchRedundancy struct {
	NumRedundantAps Optional[int] `json:"num_redundant_aps"`
}
