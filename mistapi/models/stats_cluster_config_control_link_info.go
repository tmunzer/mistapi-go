package models

import (
    "encoding/json"
)

// StatsClusterConfigControlLinkInfo represents a StatsClusterConfigControlLinkInfo struct.
type StatsClusterConfigControlLinkInfo struct {
    Name                 *string                `json:"name,omitempty"`
    Status               *string                `json:"status,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for StatsClusterConfigControlLinkInfo.
// It customizes the JSON marshaling process for StatsClusterConfigControlLinkInfo objects.
func (s StatsClusterConfigControlLinkInfo) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "name", "status"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the StatsClusterConfigControlLinkInfo object to a map representation for JSON marshaling.
func (s StatsClusterConfigControlLinkInfo) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Name != nil {
        structMap["name"] = s.Name
    }
    if s.Status != nil {
        structMap["status"] = s.Status
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsClusterConfigControlLinkInfo.
// It customizes the JSON unmarshaling process for StatsClusterConfigControlLinkInfo objects.
func (s *StatsClusterConfigControlLinkInfo) UnmarshalJSON(input []byte) error {
    var temp tempStatsClusterConfigControlLinkInfo
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "name", "status")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Name = temp.Name
    s.Status = temp.Status
    return nil
}

// tempStatsClusterConfigControlLinkInfo is a temporary struct used for validating the fields of StatsClusterConfigControlLinkInfo.
type tempStatsClusterConfigControlLinkInfo  struct {
    Name   *string `json:"name,omitempty"`
    Status *string `json:"status,omitempty"`
}
