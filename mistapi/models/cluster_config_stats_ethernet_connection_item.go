package models

import (
    "encoding/json"
)

// ClusterConfigStatsEthernetConnectionItem represents a ClusterConfigStatsEthernetConnectionItem struct.
type ClusterConfigStatsEthernetConnectionItem struct {
    Name                 *string        `json:"name,omitempty"`
    Status               *string        `json:"status,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ClusterConfigStatsEthernetConnectionItem.
// It customizes the JSON marshaling process for ClusterConfigStatsEthernetConnectionItem objects.
func (c ClusterConfigStatsEthernetConnectionItem) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the ClusterConfigStatsEthernetConnectionItem object to a map representation for JSON marshaling.
func (c ClusterConfigStatsEthernetConnectionItem) toMap() map[string]any {
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

// UnmarshalJSON implements the json.Unmarshaler interface for ClusterConfigStatsEthernetConnectionItem.
// It customizes the JSON unmarshaling process for ClusterConfigStatsEthernetConnectionItem objects.
func (c *ClusterConfigStatsEthernetConnectionItem) UnmarshalJSON(input []byte) error {
    var temp tempClusterConfigStatsEthernetConnectionItem
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

// tempClusterConfigStatsEthernetConnectionItem is a temporary struct used for validating the fields of ClusterConfigStatsEthernetConnectionItem.
type tempClusterConfigStatsEthernetConnectionItem  struct {
    Name   *string `json:"name,omitempty"`
    Status *string `json:"status,omitempty"`
}
