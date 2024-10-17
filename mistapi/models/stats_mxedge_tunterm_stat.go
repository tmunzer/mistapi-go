package models

import (
	"encoding/json"
)

// StatsMxedgeTuntermStat represents a StatsMxedgeTuntermStat struct.
type StatsMxedgeTuntermStat struct {
	MonitoringFailed     *bool          `json:"monitoring_failed,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for StatsMxedgeTuntermStat.
// It customizes the JSON marshaling process for StatsMxedgeTuntermStat objects.
func (s StatsMxedgeTuntermStat) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(s.toMap())
}

// toMap converts the StatsMxedgeTuntermStat object to a map representation for JSON marshaling.
func (s StatsMxedgeTuntermStat) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, s.AdditionalProperties)
	if s.MonitoringFailed != nil {
		structMap["monitoring_failed"] = s.MonitoringFailed
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsMxedgeTuntermStat.
// It customizes the JSON unmarshaling process for StatsMxedgeTuntermStat objects.
func (s *StatsMxedgeTuntermStat) UnmarshalJSON(input []byte) error {
	var temp tempStatsMxedgeTuntermStat
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "monitoring_failed")
	if err != nil {
		return err
	}

	s.AdditionalProperties = additionalProperties
	s.MonitoringFailed = temp.MonitoringFailed
	return nil
}

// tempStatsMxedgeTuntermStat is a temporary struct used for validating the fields of StatsMxedgeTuntermStat.
type tempStatsMxedgeTuntermStat struct {
	MonitoringFailed *bool `json:"monitoring_failed,omitempty"`
}
