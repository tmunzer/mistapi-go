package models

import (
    "encoding/json"
)

// ApTemplateMatching represents a ApTemplateMatching struct.
type ApTemplateMatching struct {
    Enabled              *bool                    `json:"enabled,omitempty"`
    Rules                []ApTemplateMatchingRule `json:"rules,omitempty"`
    AdditionalProperties map[string]any           `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ApTemplateMatching.
// It customizes the JSON marshaling process for ApTemplateMatching objects.
func (a ApTemplateMatching) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the ApTemplateMatching object to a map representation for JSON marshaling.
func (a ApTemplateMatching) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, a.AdditionalProperties)
    if a.Enabled != nil {
        structMap["enabled"] = a.Enabled
    }
    if a.Rules != nil {
        structMap["rules"] = a.Rules
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ApTemplateMatching.
// It customizes the JSON unmarshaling process for ApTemplateMatching objects.
func (a *ApTemplateMatching) UnmarshalJSON(input []byte) error {
    var temp apTemplateMatching
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "enabled", "rules")
    if err != nil {
    	return err
    }
    
    a.AdditionalProperties = additionalProperties
    a.Enabled = temp.Enabled
    a.Rules = temp.Rules
    return nil
}

// apTemplateMatching is a temporary struct used for validating the fields of ApTemplateMatching.
type apTemplateMatching  struct {
    Enabled *bool                    `json:"enabled,omitempty"`
    Rules   []ApTemplateMatchingRule `json:"rules,omitempty"`
}
