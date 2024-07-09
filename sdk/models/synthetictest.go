package models

import (
    "encoding/json"
)

// Synthetictest represents a Synthetictest struct.
type Synthetictest struct {
    Email                *string        `json:"email,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for Synthetictest.
// It customizes the JSON marshaling process for Synthetictest objects.
func (s Synthetictest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the Synthetictest object to a map representation for JSON marshaling.
func (s Synthetictest) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Email != nil {
        structMap["email"] = s.Email
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Synthetictest.
// It customizes the JSON unmarshaling process for Synthetictest objects.
func (s *Synthetictest) UnmarshalJSON(input []byte) error {
    var temp synthetictest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "email")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.Email = temp.Email
    return nil
}

// synthetictest is a temporary struct used for validating the fields of Synthetictest.
type synthetictest  struct {
    Email *string `json:"email,omitempty"`
}
