// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// StatsSwitchClientsStatsTotal represents a StatsSwitchClientsStatsTotal struct.
type StatsSwitchClientsStatsTotal struct {
    NumAps               []int                  `json:"num_aps,omitempty"`
    NumWiredClients      *int                   `json:"num_wired_clients,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for StatsSwitchClientsStatsTotal,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s StatsSwitchClientsStatsTotal) String() string {
    return fmt.Sprintf(
    	"StatsSwitchClientsStatsTotal[NumAps=%v, NumWiredClients=%v, AdditionalProperties=%v]",
    	s.NumAps, s.NumWiredClients, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for StatsSwitchClientsStatsTotal.
// It customizes the JSON marshaling process for StatsSwitchClientsStatsTotal objects.
func (s StatsSwitchClientsStatsTotal) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "num_aps", "num_wired_clients"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the StatsSwitchClientsStatsTotal object to a map representation for JSON marshaling.
func (s StatsSwitchClientsStatsTotal) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.NumAps != nil {
        structMap["num_aps"] = s.NumAps
    }
    if s.NumWiredClients != nil {
        structMap["num_wired_clients"] = s.NumWiredClients
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsSwitchClientsStatsTotal.
// It customizes the JSON unmarshaling process for StatsSwitchClientsStatsTotal objects.
func (s *StatsSwitchClientsStatsTotal) UnmarshalJSON(input []byte) error {
    var temp tempStatsSwitchClientsStatsTotal
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "num_aps", "num_wired_clients")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.NumAps = temp.NumAps
    s.NumWiredClients = temp.NumWiredClients
    return nil
}

// tempStatsSwitchClientsStatsTotal is a temporary struct used for validating the fields of StatsSwitchClientsStatsTotal.
type tempStatsSwitchClientsStatsTotal  struct {
    NumAps          []int `json:"num_aps,omitempty"`
    NumWiredClients *int  `json:"num_wired_clients,omitempty"`
}
