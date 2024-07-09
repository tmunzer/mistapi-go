package models

import (
    "encoding/json"
)

// SiteTemplate represents a SiteTemplate struct.
type SiteTemplate struct {
    AutoUpgrade          *SiteTemplateAutoUpgrade `json:"auto_upgrade,omitempty"`
    Name                 *string                  `json:"name,omitempty"`
    // a dictionary of name->value, the vars can then be used in Wlans. This can overwrite those from Site Vars
    Vars                 map[string]string        `json:"vars,omitempty"`
    AdditionalProperties map[string]any           `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SiteTemplate.
// It customizes the JSON marshaling process for SiteTemplate objects.
func (s SiteTemplate) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SiteTemplate object to a map representation for JSON marshaling.
func (s SiteTemplate) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.AutoUpgrade != nil {
        structMap["auto_upgrade"] = s.AutoUpgrade.toMap()
    }
    if s.Name != nil {
        structMap["name"] = s.Name
    }
    if s.Vars != nil {
        structMap["vars"] = s.Vars
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SiteTemplate.
// It customizes the JSON unmarshaling process for SiteTemplate objects.
func (s *SiteTemplate) UnmarshalJSON(input []byte) error {
    var temp siteTemplate
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "auto_upgrade", "name", "vars")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.AutoUpgrade = temp.AutoUpgrade
    s.Name = temp.Name
    s.Vars = temp.Vars
    return nil
}

// siteTemplate is a temporary struct used for validating the fields of SiteTemplate.
type siteTemplate  struct {
    AutoUpgrade *SiteTemplateAutoUpgrade `json:"auto_upgrade,omitempty"`
    Name        *string                  `json:"name,omitempty"`
    Vars        map[string]string        `json:"vars,omitempty"`
}