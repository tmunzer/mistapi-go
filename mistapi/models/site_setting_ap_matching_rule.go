package models

import (
    "encoding/json"
)

// SiteSettingApMatchingRule represents a SiteSettingApMatchingRule struct.
type SiteSettingApMatchingRule struct {
    MatchModel           *string                 `json:"match_model,omitempty"`
    Name                 *string                 `json:"name,omitempty"`
    // Property key is the interface(s) (e.g. "eth1,eth2")
    PortConfig           map[string]ApPortConfig `json:"port_config,omitempty"`
    AdditionalProperties map[string]interface{}  `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SiteSettingApMatchingRule.
// It customizes the JSON marshaling process for SiteSettingApMatchingRule objects.
func (s SiteSettingApMatchingRule) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "match_model", "name", "port_config"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SiteSettingApMatchingRule object to a map representation for JSON marshaling.
func (s SiteSettingApMatchingRule) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.MatchModel != nil {
        structMap["match_model"] = s.MatchModel
    }
    if s.Name != nil {
        structMap["name"] = s.Name
    }
    if s.PortConfig != nil {
        structMap["port_config"] = s.PortConfig
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SiteSettingApMatchingRule.
// It customizes the JSON unmarshaling process for SiteSettingApMatchingRule objects.
func (s *SiteSettingApMatchingRule) UnmarshalJSON(input []byte) error {
    var temp tempSiteSettingApMatchingRule
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "match_model", "name", "port_config")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.MatchModel = temp.MatchModel
    s.Name = temp.Name
    s.PortConfig = temp.PortConfig
    return nil
}

// tempSiteSettingApMatchingRule is a temporary struct used for validating the fields of SiteSettingApMatchingRule.
type tempSiteSettingApMatchingRule  struct {
    MatchModel *string                 `json:"match_model,omitempty"`
    Name       *string                 `json:"name,omitempty"`
    PortConfig map[string]ApPortConfig `json:"port_config,omitempty"`
}
