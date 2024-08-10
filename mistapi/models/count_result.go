package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// CountResult represents a CountResult struct.
type CountResult struct {
    Count                int            `json:"count"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CountResult.
// It customizes the JSON marshaling process for CountResult objects.
func (c CountResult) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CountResult object to a map representation for JSON marshaling.
func (c CountResult) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["count"] = c.Count
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CountResult.
// It customizes the JSON unmarshaling process for CountResult objects.
func (c *CountResult) UnmarshalJSON(input []byte) error {
    var temp tempCountResult
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "count")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.Count = *temp.Count
    return nil
}

// tempCountResult is a temporary struct used for validating the fields of CountResult.
type tempCountResult  struct {
    Count *int `json:"count"`
}

func (c *tempCountResult) validate() error {
    var errs []string
    if c.Count == nil {
        errs = append(errs, "required field `count` is missing for type `count_result`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
