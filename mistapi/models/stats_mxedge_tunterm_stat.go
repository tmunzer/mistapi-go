package models

import (
    "encoding/json"
    "fmt"
)

// StatsMxedgeTuntermStat represents a StatsMxedgeTuntermStat struct.
type StatsMxedgeTuntermStat struct {
    MonitoringFailed     *bool                  `json:"monitoring_failed,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for StatsMxedgeTuntermStat,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s StatsMxedgeTuntermStat) String() string {
    return fmt.Sprintf(
    	"StatsMxedgeTuntermStat[MonitoringFailed=%v, AdditionalProperties=%v]",
    	s.MonitoringFailed, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for StatsMxedgeTuntermStat.
// It customizes the JSON marshaling process for StatsMxedgeTuntermStat objects.
func (s StatsMxedgeTuntermStat) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "monitoring_failed"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the StatsMxedgeTuntermStat object to a map representation for JSON marshaling.
func (s StatsMxedgeTuntermStat) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.MonitoringFailed != nil {
        structMap["monitoring_failed"] = s.MonitoringFailed
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsMxedgeTuntermStat.
// It customizes the JSON unmarshaling process for StatsMxedgeTuntermStat objects.
func (s *StatsMxedgeTuntermStat) UnmarshalJSON(input []byte) error {
    var temp tempStatsMxedgeTuntermStat
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "monitoring_failed")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.MonitoringFailed = temp.MonitoringFailed
    return nil
}

// tempStatsMxedgeTuntermStat is a temporary struct used for validating the fields of StatsMxedgeTuntermStat.
type tempStatsMxedgeTuntermStat  struct {
    MonitoringFailed *bool `json:"monitoring_failed,omitempty"`
}
