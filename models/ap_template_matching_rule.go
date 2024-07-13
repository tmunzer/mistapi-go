package models

import (
    "encoding/json"
)

// ApTemplateMatchingRule represents a ApTemplateMatchingRule struct.
type ApTemplateMatchingRule struct {
    MatchModel           *string                 `json:"match_model,omitempty"`
    Name                 *string                 `json:"name,omitempty"`
    // Property key is the interface(s) name (e.g. "eth1,eth2")
    PortConfig           map[string]ApPortConfig `json:"port_config,omitempty"`
    AdditionalProperties map[string]any          `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ApTemplateMatchingRule.
// It customizes the JSON marshaling process for ApTemplateMatchingRule objects.
func (a ApTemplateMatchingRule) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the ApTemplateMatchingRule object to a map representation for JSON marshaling.
func (a ApTemplateMatchingRule) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, a.AdditionalProperties)
    if a.MatchModel != nil {
        structMap["match_model"] = a.MatchModel
    }
    if a.Name != nil {
        structMap["name"] = a.Name
    }
    if a.PortConfig != nil {
        structMap["port_config"] = a.PortConfig
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ApTemplateMatchingRule.
// It customizes the JSON unmarshaling process for ApTemplateMatchingRule objects.
func (a *ApTemplateMatchingRule) UnmarshalJSON(input []byte) error {
    var temp apTemplateMatchingRule
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "match_model", "name", "port_config")
    if err != nil {
    	return err
    }
    
    a.AdditionalProperties = additionalProperties
    a.MatchModel = temp.MatchModel
    a.Name = temp.Name
    a.PortConfig = temp.PortConfig
    return nil
}

// apTemplateMatchingRule is a temporary struct used for validating the fields of ApTemplateMatchingRule.
type apTemplateMatchingRule  struct {
    MatchModel *string                 `json:"match_model,omitempty"`
    Name       *string                 `json:"name,omitempty"`
    PortConfig map[string]ApPortConfig `json:"port_config,omitempty"`
}
