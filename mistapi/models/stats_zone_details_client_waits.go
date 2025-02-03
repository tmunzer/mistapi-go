package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// StatsZoneDetailsClientWaits represents a StatsZoneDetailsClientWaits struct.
// Client wait time right now
type StatsZoneDetailsClientWaits struct {
    // Average wait time in seconds
    Avg                  int                    `json:"avg"`
    // Longest wait time in seconds
    Max                  int                    `json:"max"`
    // Shortest wait time in seconds
    Min                  int                    `json:"min"`
    // 95th percentile of all the wait time(s)
    P95                  int                    `json:"p95"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for StatsZoneDetailsClientWaits,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s StatsZoneDetailsClientWaits) String() string {
    return fmt.Sprintf(
    	"StatsZoneDetailsClientWaits[Avg=%v, Max=%v, Min=%v, P95=%v, AdditionalProperties=%v]",
    	s.Avg, s.Max, s.Min, s.P95, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for StatsZoneDetailsClientWaits.
// It customizes the JSON marshaling process for StatsZoneDetailsClientWaits objects.
func (s StatsZoneDetailsClientWaits) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "avg", "max", "min", "p95"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the StatsZoneDetailsClientWaits object to a map representation for JSON marshaling.
func (s StatsZoneDetailsClientWaits) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    structMap["avg"] = s.Avg
    structMap["max"] = s.Max
    structMap["min"] = s.Min
    structMap["p95"] = s.P95
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsZoneDetailsClientWaits.
// It customizes the JSON unmarshaling process for StatsZoneDetailsClientWaits objects.
func (s *StatsZoneDetailsClientWaits) UnmarshalJSON(input []byte) error {
    var temp tempStatsZoneDetailsClientWaits
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "avg", "max", "min", "p95")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Avg = *temp.Avg
    s.Max = *temp.Max
    s.Min = *temp.Min
    s.P95 = *temp.P95
    return nil
}

// tempStatsZoneDetailsClientWaits is a temporary struct used for validating the fields of StatsZoneDetailsClientWaits.
type tempStatsZoneDetailsClientWaits  struct {
    Avg *int `json:"avg"`
    Max *int `json:"max"`
    Min *int `json:"min"`
    P95 *int `json:"p95"`
}

func (s *tempStatsZoneDetailsClientWaits) validate() error {
    var errs []string
    if s.Avg == nil {
        errs = append(errs, "required field `avg` is missing for type `stats_zone_details_client_waits`")
    }
    if s.Max == nil {
        errs = append(errs, "required field `max` is missing for type `stats_zone_details_client_waits`")
    }
    if s.Min == nil {
        errs = append(errs, "required field `min` is missing for type `stats_zone_details_client_waits`")
    }
    if s.P95 == nil {
        errs = append(errs, "required field `p95` is missing for type `stats_zone_details_client_waits`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
