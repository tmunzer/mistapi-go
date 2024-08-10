package models

import (
    "encoding/json"
)

// MxedgeStatsCpuStat represents a MxedgeStatsCpuStat struct.
// CPU/core stats list
type MxedgeStatsCpuStat struct {
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

// MarshalJSON implements the json.Marshaler interface for MxedgeStatsCpuStat.
// It customizes the JSON marshaling process for MxedgeStatsCpuStat objects.
func (m MxedgeStatsCpuStat) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(m.toMap())
}

// toMap converts the MxedgeStatsCpuStat object to a map representation for JSON marshaling.
func (m MxedgeStatsCpuStat) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, m.AdditionalProperties)
    if m.Cpus != nil {
        structMap["cpus"] = m.Cpus
    }
    if m.Idle != nil {
        structMap["idle"] = m.Idle
    }
    if m.Interrupt != nil {
        structMap["interrupt"] = m.Interrupt
    }
    if m.System != nil {
        structMap["system"] = m.System
    }
    if m.Usage != nil {
        structMap["usage"] = m.Usage
    }
    if m.User != nil {
        structMap["user"] = m.User
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MxedgeStatsCpuStat.
// It customizes the JSON unmarshaling process for MxedgeStatsCpuStat objects.
func (m *MxedgeStatsCpuStat) UnmarshalJSON(input []byte) error {
    var temp tempMxedgeStatsCpuStat
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "cpus", "idle", "interrupt", "system", "usage", "user")
    if err != nil {
    	return err
    }
    
    m.AdditionalProperties = additionalProperties
    m.Cpus = temp.Cpus
    m.Idle = temp.Idle
    m.Interrupt = temp.Interrupt
    m.System = temp.System
    m.Usage = temp.Usage
    m.User = temp.User
    return nil
}

// tempMxedgeStatsCpuStat is a temporary struct used for validating the fields of MxedgeStatsCpuStat.
type tempMxedgeStatsCpuStat  struct {
    Cpus      map[string]CpuStat `json:"cpus,omitempty"`
    Idle      *int               `json:"idle,omitempty"`
    Interrupt *int               `json:"interrupt,omitempty"`
    System    *int               `json:"system,omitempty"`
    Usage     *int               `json:"usage,omitempty"`
    User      *int               `json:"user,omitempty"`
}
