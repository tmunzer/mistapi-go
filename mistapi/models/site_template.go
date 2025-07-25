// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// SiteTemplate represents a SiteTemplate struct.
type SiteTemplate struct {
    AutoUpgrade          *SiteTemplateAutoUpgrade `json:"auto_upgrade,omitempty"`
    Name                 *string                  `json:"name,omitempty"`
    // Dictionary of name->value, the vars can then be used in Wlans. This can overwrite those from Site Vars
    Vars                 map[string]string        `json:"vars,omitempty"`
    AdditionalProperties map[string]interface{}   `json:"_"`
}

// String implements the fmt.Stringer interface for SiteTemplate,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SiteTemplate) String() string {
    return fmt.Sprintf(
    	"SiteTemplate[AutoUpgrade=%v, Name=%v, Vars=%v, AdditionalProperties=%v]",
    	s.AutoUpgrade, s.Name, s.Vars, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SiteTemplate.
// It customizes the JSON marshaling process for SiteTemplate objects.
func (s SiteTemplate) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "auto_upgrade", "name", "vars"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SiteTemplate object to a map representation for JSON marshaling.
func (s SiteTemplate) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
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
    var temp tempSiteTemplate
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "auto_upgrade", "name", "vars")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.AutoUpgrade = temp.AutoUpgrade
    s.Name = temp.Name
    s.Vars = temp.Vars
    return nil
}

// tempSiteTemplate is a temporary struct used for validating the fields of SiteTemplate.
type tempSiteTemplate  struct {
    AutoUpgrade *SiteTemplateAutoUpgrade `json:"auto_upgrade,omitempty"`
    Name        *string                  `json:"name,omitempty"`
    Vars        map[string]string        `json:"vars,omitempty"`
}
