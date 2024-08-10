package models

import (
    "encoding/json"
)

// ClusterConfigStatsRedundancyGroupInfoItem represents a ClusterConfigStatsRedundancyGroupInfoItem struct.
type ClusterConfigStatsRedundancyGroupInfoItem struct {
    Id                   *int           `json:"Id,omitempty"`
    MonitoringFailure    *string        `json:"MonitoringFailure,omitempty"`
    Threshold            *int           `json:"Threshold,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ClusterConfigStatsRedundancyGroupInfoItem.
// It customizes the JSON marshaling process for ClusterConfigStatsRedundancyGroupInfoItem objects.
func (c ClusterConfigStatsRedundancyGroupInfoItem) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the ClusterConfigStatsRedundancyGroupInfoItem object to a map representation for JSON marshaling.
func (c ClusterConfigStatsRedundancyGroupInfoItem) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    if c.Id != nil {
        structMap["Id"] = c.Id
    }
    if c.MonitoringFailure != nil {
        structMap["MonitoringFailure"] = c.MonitoringFailure
    }
    if c.Threshold != nil {
        structMap["Threshold"] = c.Threshold
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ClusterConfigStatsRedundancyGroupInfoItem.
// It customizes the JSON unmarshaling process for ClusterConfigStatsRedundancyGroupInfoItem objects.
func (c *ClusterConfigStatsRedundancyGroupInfoItem) UnmarshalJSON(input []byte) error {
    var temp tempClusterConfigStatsRedundancyGroupInfoItem
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "Id", "MonitoringFailure", "Threshold")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.Id = temp.Id
    c.MonitoringFailure = temp.MonitoringFailure
    c.Threshold = temp.Threshold
    return nil
}

// tempClusterConfigStatsRedundancyGroupInfoItem is a temporary struct used for validating the fields of ClusterConfigStatsRedundancyGroupInfoItem.
type tempClusterConfigStatsRedundancyGroupInfoItem  struct {
    Id                *int    `json:"Id,omitempty"`
    MonitoringFailure *string `json:"MonitoringFailure,omitempty"`
    Threshold         *int    `json:"Threshold,omitempty"`
}
