package models

import (
    "encoding/json"
)

// StatsMxedgeCpuStat represents a StatsMxedgeCpuStat struct.
// CPU/core stats list
type StatsMxedgeCpuStat struct {
    Cpus                 map[string]CpuStat `json:"cpus,omitempty"`
    // percentage of Idle, Idle/(Idle + Busy) since last sampling
    Idle                 *int               `json:"idle,omitempty"`
    // percentage of Interrupt, (Irq + SoftIrq)/(Idle + Busy) since last sampling
    Interrupt            *int               `json:"interrupt,omitempty"`
    // percentage of System, System/(Idle + Busy) since last sampling
    System               *int               `json:"system,omitempty"`
    // percentage of load, Busy/(Idle + Busy) since last sampling
    Usage                *int               `json:"usage,omitempty"`
    // percentage of User, User/(Idle + Busy) since last sampling
    User                 *int               `json:"user,omitempty"`
    AdditionalProperties map[string]any     `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for StatsMxedgeCpuStat.
// It customizes the JSON marshaling process for StatsMxedgeCpuStat objects.
func (s StatsMxedgeCpuStat) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the StatsMxedgeCpuStat object to a map representation for JSON marshaling.
func (s StatsMxedgeCpuStat) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
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
    additionalProperties, err := UnmarshalAdditionalProperties(input, "cpus", "idle", "interrupt", "system", "usage", "user")
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
