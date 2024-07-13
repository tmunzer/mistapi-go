package models

import (
    "encoding/json"
)

// SwitchMatching represents a SwitchMatching struct.
// Switch template
type SwitchMatching struct {
    Enable               *bool                `json:"enable,omitempty"`
    Rules                []SwitchMatchingRule `json:"rules,omitempty"`
    AdditionalProperties map[string]any       `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SwitchMatching.
// It customizes the JSON marshaling process for SwitchMatching objects.
func (s SwitchMatching) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SwitchMatching object to a map representation for JSON marshaling.
func (s SwitchMatching) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
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
    var temp switchMatching
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "enable", "rules")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.Enable = temp.Enable
    s.Rules = temp.Rules
    return nil
}

// switchMatching is a temporary struct used for validating the fields of SwitchMatching.
type switchMatching  struct {
    Enable *bool                `json:"enable,omitempty"`
    Rules  []SwitchMatchingRule `json:"rules,omitempty"`
}
