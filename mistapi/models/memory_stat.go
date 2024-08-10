package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// MemoryStat represents a MemoryStat struct.
// memory usage stat (for virtual chassis, memory usage of master RE)
type MemoryStat struct {
    Usage                float64        `json:"usage"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for MemoryStat.
// It customizes the JSON marshaling process for MemoryStat objects.
func (m MemoryStat) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(m.toMap())
}

// toMap converts the MemoryStat object to a map representation for JSON marshaling.
func (m MemoryStat) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, m.AdditionalProperties)
    structMap["usage"] = m.Usage
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MemoryStat.
// It customizes the JSON unmarshaling process for MemoryStat objects.
func (m *MemoryStat) UnmarshalJSON(input []byte) error {
    var temp tempMemoryStat
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "usage")
    if err != nil {
    	return err
    }
    
    m.AdditionalProperties = additionalProperties
    m.Usage = *temp.Usage
    return nil
}

// tempMemoryStat is a temporary struct used for validating the fields of MemoryStat.
type tempMemoryStat  struct {
    Usage *float64 `json:"usage"`
}

func (m *tempMemoryStat) validate() error {
    var errs []string
    if m.Usage == nil {
        errs = append(errs, "required field `usage` is missing for type `memory_stat`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
