package models

import (
    "encoding/json"
)

// CpuStat represents a CpuStat struct.
type CpuStat struct {
    // Percentage of CPU time that is idle
    Idle                 Optional[float64] `json:"idle"`
    // Percentage of CPU time being used by interrupts
    Interrupt            Optional[float64] `json:"interrupt"`
    // Load averages for the last 1, 5, and 15 minutes
    LoadAvg              []float64         `json:"load_avg,omitempty"`
    // Percentage of CPU time being used by system processes
    System               Optional[float64] `json:"system"`
    // Percentage of CPU time being used by user processe
    User                 Optional[float64] `json:"user"`
    AdditionalProperties map[string]any    `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CpuStat.
// It customizes the JSON marshaling process for CpuStat objects.
func (c CpuStat) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CpuStat object to a map representation for JSON marshaling.
func (c CpuStat) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    if c.Idle.IsValueSet() {
        if c.Idle.Value() != nil {
            structMap["idle"] = c.Idle.Value()
        } else {
            structMap["idle"] = nil
        }
    }
    if c.Interrupt.IsValueSet() {
        if c.Interrupt.Value() != nil {
            structMap["interrupt"] = c.Interrupt.Value()
        } else {
            structMap["interrupt"] = nil
        }
    }
    if c.LoadAvg != nil {
        structMap["load_avg"] = c.LoadAvg
    }
    if c.System.IsValueSet() {
        if c.System.Value() != nil {
            structMap["system"] = c.System.Value()
        } else {
            structMap["system"] = nil
        }
    }
    if c.User.IsValueSet() {
        if c.User.Value() != nil {
            structMap["user"] = c.User.Value()
        } else {
            structMap["user"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CpuStat.
// It customizes the JSON unmarshaling process for CpuStat objects.
func (c *CpuStat) UnmarshalJSON(input []byte) error {
    var temp cpuStat
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "idle", "interrupt", "load_avg", "system", "user")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.Idle = temp.Idle
    c.Interrupt = temp.Interrupt
    c.LoadAvg = temp.LoadAvg
    c.System = temp.System
    c.User = temp.User
    return nil
}

// cpuStat is a temporary struct used for validating the fields of CpuStat.
type cpuStat  struct {
    Idle      Optional[float64] `json:"idle"`
    Interrupt Optional[float64] `json:"interrupt"`
    LoadAvg   []float64         `json:"load_avg,omitempty"`
    System    Optional[float64] `json:"system"`
    User      Optional[float64] `json:"user"`
}