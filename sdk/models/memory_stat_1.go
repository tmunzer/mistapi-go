package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// MemoryStat1 represents a MemoryStat1 struct.
// device mac
type MemoryStat1 struct {
    Usage                float64        `json:"usage"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for MemoryStat1.
// It customizes the JSON marshaling process for MemoryStat1 objects.
func (m MemoryStat1) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(m.toMap())
}

// toMap converts the MemoryStat1 object to a map representation for JSON marshaling.
func (m MemoryStat1) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, m.AdditionalProperties)
    structMap["usage"] = m.Usage
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MemoryStat1.
// It customizes the JSON unmarshaling process for MemoryStat1 objects.
func (m *MemoryStat1) UnmarshalJSON(input []byte) error {
    var temp memoryStat1
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

// memoryStat1 is a temporary struct used for validating the fields of MemoryStat1.
type memoryStat1  struct {
    Usage *float64 `json:"usage"`
}

func (m *memoryStat1) validate() error {
    var errs []string
    if m.Usage == nil {
        errs = append(errs, "required field `usage` is missing for type `Memory_Stat1`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
