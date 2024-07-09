package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// ApTemplate represents a ApTemplate struct.
type ApTemplate struct {
    ApMatching           ApTemplateMatching `json:"ap_matching"`
    CreatedTime          *float64           `json:"created_time,omitempty"`
    ForSite              *bool              `json:"for_site,omitempty"`
    Id                   *uuid.UUID         `json:"id,omitempty"`
    ModifiedTime         *float64           `json:"modified_time,omitempty"`
    OrgId                *uuid.UUID         `json:"org_id,omitempty"`
    SiteId               *uuid.UUID         `json:"site_id,omitempty"`
    Wifi                 *ApTemplateWifi    `json:"wifi,omitempty"`
    AdditionalProperties map[string]any     `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ApTemplate.
// It customizes the JSON marshaling process for ApTemplate objects.
func (a ApTemplate) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the ApTemplate object to a map representation for JSON marshaling.
func (a ApTemplate) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, a.AdditionalProperties)
    structMap["ap_matching"] = a.ApMatching.toMap()
    if a.CreatedTime != nil {
        structMap["created_time"] = a.CreatedTime
    }
    if a.ForSite != nil {
        structMap["for_site"] = a.ForSite
    }
    if a.Id != nil {
        structMap["id"] = a.Id
    }
    if a.ModifiedTime != nil {
        structMap["modified_time"] = a.ModifiedTime
    }
    if a.OrgId != nil {
        structMap["org_id"] = a.OrgId
    }
    if a.SiteId != nil {
        structMap["site_id"] = a.SiteId
    }
    if a.Wifi != nil {
        structMap["wifi"] = a.Wifi.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ApTemplate.
// It customizes the JSON unmarshaling process for ApTemplate objects.
func (a *ApTemplate) UnmarshalJSON(input []byte) error {
    var temp apTemplate
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "ap_matching", "created_time", "for_site", "id", "modified_time", "org_id", "site_id", "wifi")
    if err != nil {
    	return err
    }
    
    a.AdditionalProperties = additionalProperties
    a.ApMatching = *temp.ApMatching
    a.CreatedTime = temp.CreatedTime
    a.ForSite = temp.ForSite
    a.Id = temp.Id
    a.ModifiedTime = temp.ModifiedTime
    a.OrgId = temp.OrgId
    a.SiteId = temp.SiteId
    a.Wifi = temp.Wifi
    return nil
}

// apTemplate is a temporary struct used for validating the fields of ApTemplate.
type apTemplate  struct {
    ApMatching   *ApTemplateMatching `json:"ap_matching"`
    CreatedTime  *float64            `json:"created_time,omitempty"`
    ForSite      *bool               `json:"for_site,omitempty"`
    Id           *uuid.UUID          `json:"id,omitempty"`
    ModifiedTime *float64            `json:"modified_time,omitempty"`
    OrgId        *uuid.UUID          `json:"org_id,omitempty"`
    SiteId       *uuid.UUID          `json:"site_id,omitempty"`
    Wifi         *ApTemplateWifi     `json:"wifi,omitempty"`
}

func (a *apTemplate) validate() error {
    var errs []string
    if a.ApMatching == nil {
        errs = append(errs, "required field `ap_matching` is missing for type `Ap_Template`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
