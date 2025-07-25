// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// ApTemplateMatchingRule represents a ApTemplateMatchingRule struct.
type ApTemplateMatchingRule struct {
    MatchModel           *string                 `json:"match_model,omitempty"`
    Name                 *string                 `json:"name,omitempty"`
    // Property key is the interface(s) name (e.g. "eth1,eth2")
    PortConfig           map[string]ApPortConfig `json:"port_config,omitempty"`
    AdditionalProperties map[string]interface{}  `json:"_"`
}

// String implements the fmt.Stringer interface for ApTemplateMatchingRule,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a ApTemplateMatchingRule) String() string {
    return fmt.Sprintf(
    	"ApTemplateMatchingRule[MatchModel=%v, Name=%v, PortConfig=%v, AdditionalProperties=%v]",
    	a.MatchModel, a.Name, a.PortConfig, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ApTemplateMatchingRule.
// It customizes the JSON marshaling process for ApTemplateMatchingRule objects.
func (a ApTemplateMatchingRule) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(a.AdditionalProperties,
        "match_model", "name", "port_config"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(a.toMap())
}

// toMap converts the ApTemplateMatchingRule object to a map representation for JSON marshaling.
func (a ApTemplateMatchingRule) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, a.AdditionalProperties)
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
    var temp tempApTemplateMatchingRule
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "match_model", "name", "port_config")
    if err != nil {
    	return err
    }
    a.AdditionalProperties = additionalProperties
    
    a.MatchModel = temp.MatchModel
    a.Name = temp.Name
    a.PortConfig = temp.PortConfig
    return nil
}

// tempApTemplateMatchingRule is a temporary struct used for validating the fields of ApTemplateMatchingRule.
type tempApTemplateMatchingRule  struct {
    MatchModel *string                 `json:"match_model,omitempty"`
    Name       *string                 `json:"name,omitempty"`
    PortConfig map[string]ApPortConfig `json:"port_config,omitempty"`
}
