package models

import (
    "encoding/json"
)

// StatsClusterConfigEthernetConnectionItem represents a StatsClusterConfigEthernetConnectionItem struct.
type StatsClusterConfigEthernetConnectionItem struct {
    Name                 *string                `json:"name,omitempty"`
    Status               *string                `json:"status,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for StatsClusterConfigEthernetConnectionItem.
// It customizes the JSON marshaling process for StatsClusterConfigEthernetConnectionItem objects.
func (s StatsClusterConfigEthernetConnectionItem) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "name", "status"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the StatsClusterConfigEthernetConnectionItem object to a map representation for JSON marshaling.
func (s StatsClusterConfigEthernetConnectionItem) toMap() map[string]any {
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

// UnmarshalJSON implements the json.Unmarshaler interface for StatsClusterConfigEthernetConnectionItem.
// It customizes the JSON unmarshaling process for StatsClusterConfigEthernetConnectionItem objects.
func (s *StatsClusterConfigEthernetConnectionItem) UnmarshalJSON(input []byte) error {
    var temp tempStatsClusterConfigEthernetConnectionItem
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

// tempStatsClusterConfigEthernetConnectionItem is a temporary struct used for validating the fields of StatsClusterConfigEthernetConnectionItem.
type tempStatsClusterConfigEthernetConnectionItem  struct {
    Name   *string `json:"name,omitempty"`
    Status *string `json:"status,omitempty"`
}
