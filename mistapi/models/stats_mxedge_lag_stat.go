package models

import (
    "encoding/json"
    "fmt"
)

// StatsMxedgeLagStat represents a StatsMxedgeLagStat struct.
type StatsMxedgeLagStat struct {
    // list of ports active on the LAG defined by the LACP
    ActivePorts          []string               `json:"active_ports,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for StatsMxedgeLagStat,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s StatsMxedgeLagStat) String() string {
    return fmt.Sprintf(
    	"StatsMxedgeLagStat[ActivePorts=%v, AdditionalProperties=%v]",
    	s.ActivePorts, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for StatsMxedgeLagStat.
// It customizes the JSON marshaling process for StatsMxedgeLagStat objects.
func (s StatsMxedgeLagStat) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "active_ports"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the StatsMxedgeLagStat object to a map representation for JSON marshaling.
func (s StatsMxedgeLagStat) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.ActivePorts != nil {
        structMap["active_ports"] = s.ActivePorts
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsMxedgeLagStat.
// It customizes the JSON unmarshaling process for StatsMxedgeLagStat objects.
func (s *StatsMxedgeLagStat) UnmarshalJSON(input []byte) error {
    var temp tempStatsMxedgeLagStat
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "active_ports")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.ActivePorts = temp.ActivePorts
    return nil
}

// tempStatsMxedgeLagStat is a temporary struct used for validating the fields of StatsMxedgeLagStat.
type tempStatsMxedgeLagStat  struct {
    ActivePorts []string `json:"active_ports,omitempty"`
}
