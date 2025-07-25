// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// StatsGatewayCluster represents a StatsGatewayCluster struct.
type StatsGatewayCluster struct {
    State                Optional[string]       `json:"state"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for StatsGatewayCluster,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s StatsGatewayCluster) String() string {
    return fmt.Sprintf(
    	"StatsGatewayCluster[State=%v, AdditionalProperties=%v]",
    	s.State, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for StatsGatewayCluster.
// It customizes the JSON marshaling process for StatsGatewayCluster objects.
func (s StatsGatewayCluster) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "state"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the StatsGatewayCluster object to a map representation for JSON marshaling.
func (s StatsGatewayCluster) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.State.IsValueSet() {
        if s.State.Value() != nil {
            structMap["state"] = s.State.Value()
        } else {
            structMap["state"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsGatewayCluster.
// It customizes the JSON unmarshaling process for StatsGatewayCluster objects.
func (s *StatsGatewayCluster) UnmarshalJSON(input []byte) error {
    var temp tempStatsGatewayCluster
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "state")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.State = temp.State
    return nil
}

// tempStatsGatewayCluster is a temporary struct used for validating the fields of StatsGatewayCluster.
type tempStatsGatewayCluster  struct {
    State Optional[string] `json:"state"`
}
