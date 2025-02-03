package models

import (
    "encoding/json"
    "fmt"
)

// SwitchMatching represents a SwitchMatching struct.
// Defines custom switch configuration based on different criterias
type SwitchMatching struct {
    Enable               *bool                  `json:"enable,omitempty"`
    Rules                []SwitchMatchingRule   `json:"rules,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SwitchMatching,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SwitchMatching) String() string {
    return fmt.Sprintf(
    	"SwitchMatching[Enable=%v, Rules=%v, AdditionalProperties=%v]",
    	s.Enable, s.Rules, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SwitchMatching.
// It customizes the JSON marshaling process for SwitchMatching objects.
func (s SwitchMatching) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "enable", "rules"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SwitchMatching object to a map representation for JSON marshaling.
func (s SwitchMatching) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Enable != nil {
        structMap["enable"] = s.Enable
    }
    if s.Rules != nil {
        structMap["rules"] = s.Rules
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SwitchMatching.
// It customizes the JSON unmarshaling process for SwitchMatching objects.
func (s *SwitchMatching) UnmarshalJSON(input []byte) error {
    var temp tempSwitchMatching
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "enable", "rules")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Enable = temp.Enable
    s.Rules = temp.Rules
    return nil
}

// tempSwitchMatching is a temporary struct used for validating the fields of SwitchMatching.
type tempSwitchMatching  struct {
    Enable *bool                `json:"enable,omitempty"`
    Rules  []SwitchMatchingRule `json:"rules,omitempty"`
}
