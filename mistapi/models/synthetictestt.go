package models

import (
    "encoding/json"
    "fmt"
)

// Synthetictest represents a Synthetictest struct.
type Synthetictest struct {
    Email                *string                `json:"email,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for Synthetictest,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s Synthetictest) String() string {
    return fmt.Sprintf(
    	"Synthetictest[Email=%v, AdditionalProperties=%v]",
    	s.Email, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for Synthetictest.
// It customizes the JSON marshaling process for Synthetictest objects.
func (s Synthetictest) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "email"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the Synthetictest object to a map representation for JSON marshaling.
func (s Synthetictest) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Email != nil {
        structMap["email"] = s.Email
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Synthetictest.
// It customizes the JSON unmarshaling process for Synthetictest objects.
func (s *Synthetictest) UnmarshalJSON(input []byte) error {
    var temp tempSynthetictest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "email")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Email = temp.Email
    return nil
}

// tempSynthetictest is a temporary struct used for validating the fields of Synthetictest.
type tempSynthetictest  struct {
    Email *string `json:"email,omitempty"`
}
