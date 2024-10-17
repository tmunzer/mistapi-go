package models

import (
	"encoding/json"
)

// StatsApAutoUpgrade represents a StatsApAutoUpgrade struct.
type StatsApAutoUpgrade struct {
	Lastcheck            Optional[int64] `json:"lastcheck"`
	AdditionalProperties map[string]any  `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for StatsApAutoUpgrade.
// It customizes the JSON marshaling process for StatsApAutoUpgrade objects.
func (s StatsApAutoUpgrade) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(s.toMap())
}

// toMap converts the StatsApAutoUpgrade object to a map representation for JSON marshaling.
func (s StatsApAutoUpgrade) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, s.AdditionalProperties)
	if s.Lastcheck.IsValueSet() {
		if s.Lastcheck.Value() != nil {
			structMap["lastcheck"] = s.Lastcheck.Value()
		} else {
			structMap["lastcheck"] = nil
		}
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsApAutoUpgrade.
// It customizes the JSON unmarshaling process for StatsApAutoUpgrade objects.
func (s *StatsApAutoUpgrade) UnmarshalJSON(input []byte) error {
	var temp tempStatsApAutoUpgrade
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "lastcheck")
	if err != nil {
		return err
	}

	s.AdditionalProperties = additionalProperties
	s.Lastcheck = temp.Lastcheck
	return nil
}

// tempStatsApAutoUpgrade is a temporary struct used for validating the fields of StatsApAutoUpgrade.
type tempStatsApAutoUpgrade struct {
	Lastcheck Optional[int64] `json:"lastcheck"`
}
