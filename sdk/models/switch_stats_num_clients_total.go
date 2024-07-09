package models

import (
    "encoding/json"
)

// SwitchStatsNumClientsTotal represents a SwitchStatsNumClientsTotal struct.
type SwitchStatsNumClientsTotal struct {
    NumAps               *string        `json:"num_aps,omitempty"`
    NumClients           *string        `json:"num_clients,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SwitchStatsNumClientsTotal.
// It customizes the JSON marshaling process for SwitchStatsNumClientsTotal objects.
func (s SwitchStatsNumClientsTotal) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SwitchStatsNumClientsTotal object to a map representation for JSON marshaling.
func (s SwitchStatsNumClientsTotal) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.NumAps != nil {
        structMap["num_aps"] = s.NumAps
    }
    if s.NumClients != nil {
        structMap["num_clients"] = s.NumClients
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SwitchStatsNumClientsTotal.
// It customizes the JSON unmarshaling process for SwitchStatsNumClientsTotal objects.
func (s *SwitchStatsNumClientsTotal) UnmarshalJSON(input []byte) error {
    var temp switchStatsNumClientsTotal
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "num_aps", "num_clients")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.NumAps = temp.NumAps
    s.NumClients = temp.NumClients
    return nil
}

// switchStatsNumClientsTotal is a temporary struct used for validating the fields of SwitchStatsNumClientsTotal.
type switchStatsNumClientsTotal  struct {
    NumAps     *string `json:"num_aps,omitempty"`
    NumClients *string `json:"num_clients,omitempty"`
}
