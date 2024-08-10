package models

import (
    "encoding/json"
)

// SwitchStatsClientsStatsTotal represents a SwitchStatsClientsStatsTotal struct.
type SwitchStatsClientsStatsTotal struct {
    NumAps               []int          `json:"num_aps,omitempty"`
    NumWiredClients      *int           `json:"num_wired_clients,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SwitchStatsClientsStatsTotal.
// It customizes the JSON marshaling process for SwitchStatsClientsStatsTotal objects.
func (s SwitchStatsClientsStatsTotal) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SwitchStatsClientsStatsTotal object to a map representation for JSON marshaling.
func (s SwitchStatsClientsStatsTotal) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.NumAps != nil {
        structMap["num_aps"] = s.NumAps
    }
    if s.NumWiredClients != nil {
        structMap["num_wired_clients"] = s.NumWiredClients
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SwitchStatsClientsStatsTotal.
// It customizes the JSON unmarshaling process for SwitchStatsClientsStatsTotal objects.
func (s *SwitchStatsClientsStatsTotal) UnmarshalJSON(input []byte) error {
    var temp tempSwitchStatsClientsStatsTotal
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "num_aps", "num_wired_clients")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.NumAps = temp.NumAps
    s.NumWiredClients = temp.NumWiredClients
    return nil
}

// tempSwitchStatsClientsStatsTotal is a temporary struct used for validating the fields of SwitchStatsClientsStatsTotal.
type tempSwitchStatsClientsStatsTotal  struct {
    NumAps          []int `json:"num_aps,omitempty"`
    NumWiredClients *int  `json:"num_wired_clients,omitempty"`
}
