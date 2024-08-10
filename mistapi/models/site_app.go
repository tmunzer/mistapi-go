package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// SiteApp represents a SiteApp struct.
type SiteApp struct {
    Group                string         `json:"group"`
    Key                  string         `json:"key"`
    Name                 string         `json:"name"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SiteApp.
// It customizes the JSON marshaling process for SiteApp objects.
func (s SiteApp) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SiteApp object to a map representation for JSON marshaling.
func (s SiteApp) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    structMap["group"] = s.Group
    structMap["key"] = s.Key
    structMap["name"] = s.Name
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SiteApp.
// It customizes the JSON unmarshaling process for SiteApp objects.
func (s *SiteApp) UnmarshalJSON(input []byte) error {
    var temp tempSiteApp
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "group", "key", "name")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.Group = *temp.Group
    s.Key = *temp.Key
    s.Name = *temp.Name
    return nil
}

// tempSiteApp is a temporary struct used for validating the fields of SiteApp.
type tempSiteApp  struct {
    Group *string `json:"group"`
    Key   *string `json:"key"`
    Name  *string `json:"name"`
}

func (s *tempSiteApp) validate() error {
    var errs []string
    if s.Group == nil {
        errs = append(errs, "required field `group` is missing for type `site_app`")
    }
    if s.Key == nil {
        errs = append(errs, "required field `key` is missing for type `site_app`")
    }
    if s.Name == nil {
        errs = append(errs, "required field `name` is missing for type `site_app`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
