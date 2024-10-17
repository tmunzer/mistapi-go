package models

import (
	"encoding/json"
)

// StatsClusterConfigRedundancyGroupInfoItem represents a StatsClusterConfigRedundancyGroupInfoItem struct.
type StatsClusterConfigRedundancyGroupInfoItem struct {
	Id                   *int           `json:"Id,omitempty"`
	MonitoringFailure    *string        `json:"MonitoringFailure,omitempty"`
	Threshold            *int           `json:"Threshold,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for StatsClusterConfigRedundancyGroupInfoItem.
// It customizes the JSON marshaling process for StatsClusterConfigRedundancyGroupInfoItem objects.
func (s StatsClusterConfigRedundancyGroupInfoItem) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(s.toMap())
}

// toMap converts the StatsClusterConfigRedundancyGroupInfoItem object to a map representation for JSON marshaling.
func (s StatsClusterConfigRedundancyGroupInfoItem) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, s.AdditionalProperties)
	if s.Id != nil {
		structMap["Id"] = s.Id
	}
	if s.MonitoringFailure != nil {
		structMap["MonitoringFailure"] = s.MonitoringFailure
	}
	if s.Threshold != nil {
		structMap["Threshold"] = s.Threshold
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsClusterConfigRedundancyGroupInfoItem.
// It customizes the JSON unmarshaling process for StatsClusterConfigRedundancyGroupInfoItem objects.
func (s *StatsClusterConfigRedundancyGroupInfoItem) UnmarshalJSON(input []byte) error {
	var temp tempStatsClusterConfigRedundancyGroupInfoItem
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "Id", "MonitoringFailure", "Threshold")
	if err != nil {
		return err
	}

	s.AdditionalProperties = additionalProperties
	s.Id = temp.Id
	s.MonitoringFailure = temp.MonitoringFailure
	s.Threshold = temp.Threshold
	return nil
}

// tempStatsClusterConfigRedundancyGroupInfoItem is a temporary struct used for validating the fields of StatsClusterConfigRedundancyGroupInfoItem.
type tempStatsClusterConfigRedundancyGroupInfoItem struct {
	Id                *int    `json:"Id,omitempty"`
	MonitoringFailure *string `json:"MonitoringFailure,omitempty"`
	Threshold         *int    `json:"Threshold,omitempty"`
}
