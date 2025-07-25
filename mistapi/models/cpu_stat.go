// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// CpuStat represents a CpuStat struct.
type CpuStat struct {
    // Percentage of CPU time that is idle
    Idle                 Optional[float64]      `json:"idle"`
    // Percentage of CPU time being used by interrupts
    Interrupt            Optional[float64]      `json:"interrupt"`
    // Load averages for the last 1, 5, and 15 minutes
    LoadAvg              []float64              `json:"load_avg,omitempty"`
    // Percentage of CPU time being used by system processes
    System               Optional[float64]      `json:"system"`
    // Percentage of CPU time being used by user processes
    User                 Optional[float64]      `json:"user"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for CpuStat,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CpuStat) String() string {
    return fmt.Sprintf(
    	"CpuStat[Idle=%v, Interrupt=%v, LoadAvg=%v, System=%v, User=%v, AdditionalProperties=%v]",
    	c.Idle, c.Interrupt, c.LoadAvg, c.System, c.User, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CpuStat.
// It customizes the JSON marshaling process for CpuStat objects.
func (c CpuStat) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "idle", "interrupt", "load_avg", "system", "user"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CpuStat object to a map representation for JSON marshaling.
func (c CpuStat) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
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
    var temp tempCpuStat
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "idle", "interrupt", "load_avg", "system", "user")
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

// tempCpuStat is a temporary struct used for validating the fields of CpuStat.
type tempCpuStat  struct {
    Idle      Optional[float64] `json:"idle"`
    Interrupt Optional[float64] `json:"interrupt"`
    LoadAvg   []float64         `json:"load_avg,omitempty"`
    System    Optional[float64] `json:"system"`
    User      Optional[float64] `json:"user"`
}
