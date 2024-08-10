package models

import (
    "encoding/json"
)

// SwitchStatsClientsStats represents a SwitchStatsClientsStats struct.
type SwitchStatsClientsStats struct {
    Total                *SwitchStatsClientsStatsTotal `json:"total,omitempty"`
    AdditionalProperties map[string]any                `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SwitchStatsClientsStats.
// It customizes the JSON marshaling process for SwitchStatsClientsStats objects.
func (s SwitchStatsClientsStats) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SwitchStatsClientsStats object to a map representation for JSON marshaling.
func (s SwitchStatsClientsStats) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Total != nil {
        structMap["total"] = s.Total.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SwitchStatsClientsStats.
// It customizes the JSON unmarshaling process for SwitchStatsClientsStats objects.
func (s *SwitchStatsClientsStats) UnmarshalJSON(input []byte) error {
    var temp tempSwitchStatsClientsStats
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "total")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.Total = temp.Total
    return nil
}

// tempSwitchStatsClientsStats is a temporary struct used for validating the fields of SwitchStatsClientsStats.
type tempSwitchStatsClientsStats  struct {
    Total *SwitchStatsClientsStatsTotal `json:"total,omitempty"`
}
