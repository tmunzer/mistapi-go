package models

import (
    "encoding/json"
    "fmt"
)

// StatsMxedgeCpuStat represents a StatsMxedgeCpuStat struct.
// CPU/core stats list
type StatsMxedgeCpuStat struct {
    Cpus                 map[string]CpuStat     `json:"cpus,omitempty"`
    // Percentage of Idle, Idle/(Idle + Busy) since last sampling
    Idle                 *int                   `json:"idle,omitempty"`
    // Percentage of Interrupt, (Irq + SoftIrq)/(Idle + Busy) since last sampling
    Interrupt            *int                   `json:"interrupt,omitempty"`
    // Percentage of System, System/(Idle + Busy) since last sampling
    System               *int                   `json:"system,omitempty"`
    // Percentage of load, Busy/(Idle + Busy) since last sampling
    Usage                *int                   `json:"usage,omitempty"`
    // Percentage of User, User/(Idle + Busy) since last sampling
    User                 *int                   `json:"user,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for StatsMxedgeCpuStat,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s StatsMxedgeCpuStat) String() string {
    return fmt.Sprintf(
    	"StatsMxedgeCpuStat[Cpus=%v, Idle=%v, Interrupt=%v, System=%v, Usage=%v, User=%v, AdditionalProperties=%v]",
    	s.Cpus, s.Idle, s.Interrupt, s.System, s.Usage, s.User, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for StatsMxedgeCpuStat.
// It customizes the JSON marshaling process for StatsMxedgeCpuStat objects.
func (s StatsMxedgeCpuStat) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "cpus", "idle", "interrupt", "system", "usage", "user"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the StatsMxedgeCpuStat object to a map representation for JSON marshaling.
func (s StatsMxedgeCpuStat) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Cpus != nil {
        structMap["cpus"] = s.Cpus
    }
    if s.Idle != nil {
        structMap["idle"] = s.Idle
    }
    if s.Interrupt != nil {
        structMap["interrupt"] = s.Interrupt
    }
    if s.System != nil {
        structMap["system"] = s.System
    }
    if s.Usage != nil {
        structMap["usage"] = s.Usage
    }
    if s.User != nil {
        structMap["user"] = s.User
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsMxedgeCpuStat.
// It customizes the JSON unmarshaling process for StatsMxedgeCpuStat objects.
func (s *StatsMxedgeCpuStat) UnmarshalJSON(input []byte) error {
    var temp tempStatsMxedgeCpuStat
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "cpus", "idle", "interrupt", "system", "usage", "user")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Cpus = temp.Cpus
    s.Idle = temp.Idle
    s.Interrupt = temp.Interrupt
    s.System = temp.System
    s.Usage = temp.Usage
    s.User = temp.User
    return nil
}

// tempStatsMxedgeCpuStat is a temporary struct used for validating the fields of StatsMxedgeCpuStat.
type tempStatsMxedgeCpuStat  struct {
    Cpus      map[string]CpuStat `json:"cpus,omitempty"`
    Idle      *int               `json:"idle,omitempty"`
    Interrupt *int               `json:"interrupt,omitempty"`
    System    *int               `json:"system,omitempty"`
    Usage     *int               `json:"usage,omitempty"`
    User      *int               `json:"user,omitempty"`
}
