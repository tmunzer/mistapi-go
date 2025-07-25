// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// StatsClusterConfigRedundancyGroupInfoItem represents a StatsClusterConfigRedundancyGroupInfoItem struct.
type StatsClusterConfigRedundancyGroupInfoItem struct {
    Id                   *int                   `json:"Id,omitempty"`
    MonitoringFailure    *string                `json:"MonitoringFailure,omitempty"`
    Threshold            *int                   `json:"Threshold,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for StatsClusterConfigRedundancyGroupInfoItem,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s StatsClusterConfigRedundancyGroupInfoItem) String() string {
    return fmt.Sprintf(
    	"StatsClusterConfigRedundancyGroupInfoItem[Id=%v, MonitoringFailure=%v, Threshold=%v, AdditionalProperties=%v]",
    	s.Id, s.MonitoringFailure, s.Threshold, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for StatsClusterConfigRedundancyGroupInfoItem.
// It customizes the JSON marshaling process for StatsClusterConfigRedundancyGroupInfoItem objects.
func (s StatsClusterConfigRedundancyGroupInfoItem) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "Id", "MonitoringFailure", "Threshold"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the StatsClusterConfigRedundancyGroupInfoItem object to a map representation for JSON marshaling.
func (s StatsClusterConfigRedundancyGroupInfoItem) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
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
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "Id", "MonitoringFailure", "Threshold")
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
type tempStatsClusterConfigRedundancyGroupInfoItem  struct {
    Id                *int    `json:"Id,omitempty"`
    MonitoringFailure *string `json:"MonitoringFailure,omitempty"`
    Threshold         *int    `json:"Threshold,omitempty"`
}
