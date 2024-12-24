package models

import (
    "encoding/json"
    "fmt"
)

// StatsSwitchClientsStats represents a StatsSwitchClientsStats struct.
type StatsSwitchClientsStats struct {
    Total                *StatsSwitchClientsStatsTotal `json:"total,omitempty"`
    AdditionalProperties map[string]interface{}        `json:"_"`
}

// String implements the fmt.Stringer interface for StatsSwitchClientsStats,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s StatsSwitchClientsStats) String() string {
    return fmt.Sprintf(
    	"StatsSwitchClientsStats[Total=%v, AdditionalProperties=%v]",
    	s.Total, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for StatsSwitchClientsStats.
// It customizes the JSON marshaling process for StatsSwitchClientsStats objects.
func (s StatsSwitchClientsStats) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "total"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the StatsSwitchClientsStats object to a map representation for JSON marshaling.
func (s StatsSwitchClientsStats) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Total != nil {
        structMap["total"] = s.Total.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsSwitchClientsStats.
// It customizes the JSON unmarshaling process for StatsSwitchClientsStats objects.
func (s *StatsSwitchClientsStats) UnmarshalJSON(input []byte) error {
    var temp tempStatsSwitchClientsStats
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "total")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Total = temp.Total
    return nil
}

// tempStatsSwitchClientsStats is a temporary struct used for validating the fields of StatsSwitchClientsStats.
type tempStatsSwitchClientsStats  struct {
    Total *StatsSwitchClientsStatsTotal `json:"total,omitempty"`
}
