package models

import (
    "encoding/json"
)

// ClusterConfigStatsControlLinkInfo represents a ClusterConfigStatsControlLinkInfo struct.
type ClusterConfigStatsControlLinkInfo struct {
    Name                 *string        `json:"name,omitempty"`
    Status               *string        `json:"status,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ClusterConfigStatsControlLinkInfo.
// It customizes the JSON marshaling process for ClusterConfigStatsControlLinkInfo objects.
func (c ClusterConfigStatsControlLinkInfo) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the ClusterConfigStatsControlLinkInfo object to a map representation for JSON marshaling.
func (c ClusterConfigStatsControlLinkInfo) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    if c.Name != nil {
        structMap["name"] = c.Name
    }
    if c.Status != nil {
        structMap["status"] = c.Status
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ClusterConfigStatsControlLinkInfo.
// It customizes the JSON unmarshaling process for ClusterConfigStatsControlLinkInfo objects.
func (c *ClusterConfigStatsControlLinkInfo) UnmarshalJSON(input []byte) error {
    var temp clusterConfigStatsControlLinkInfo
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "name", "status")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.Name = temp.Name
    c.Status = temp.Status
    return nil
}

// clusterConfigStatsControlLinkInfo is a temporary struct used for validating the fields of ClusterConfigStatsControlLinkInfo.
type clusterConfigStatsControlLinkInfo  struct {
    Name   *string `json:"name,omitempty"`
    Status *string `json:"status,omitempty"`
}
