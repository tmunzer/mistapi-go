package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// InstallerSite represents a InstallerSite struct.
type InstallerSite struct {
    Address              string         `json:"address"`
    CountryCode          string         `json:"country_code"`
    Id                   *uuid.UUID     `json:"id,omitempty"`
    Latlng               LatLng         `json:"latlng"`
    Name                 string         `json:"name"`
    RftemplateName       *string        `json:"rftemplate_name,omitempty"`
    SitegroupNames       []string       `json:"sitegroup_names,omitempty"`
    Timezone             *string        `json:"timezone,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for InstallerSite.
// It customizes the JSON marshaling process for InstallerSite objects.
func (i InstallerSite) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(i.toMap())
}

// toMap converts the InstallerSite object to a map representation for JSON marshaling.
func (i InstallerSite) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, i.AdditionalProperties)
    structMap["address"] = i.Address
    structMap["country_code"] = i.CountryCode
    if i.Id != nil {
        structMap["id"] = i.Id
    }
    structMap["latlng"] = i.Latlng.toMap()
    structMap["name"] = i.Name
    if i.RftemplateName != nil {
        structMap["rftemplate_name"] = i.RftemplateName
    }
    if i.SitegroupNames != nil {
        structMap["sitegroup_names"] = i.SitegroupNames
    }
    if i.Timezone != nil {
        structMap["timezone"] = i.Timezone
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for InstallerSite.
// It customizes the JSON unmarshaling process for InstallerSite objects.
func (i *InstallerSite) UnmarshalJSON(input []byte) error {
    var temp tempInstallerSite
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "address", "country_code", "id", "latlng", "name", "rftemplate_name", "sitegroup_names", "timezone")
    if err != nil {
    	return err
    }
    
    i.AdditionalProperties = additionalProperties
    i.Address = *temp.Address
    i.CountryCode = *temp.CountryCode
    i.Id = temp.Id
    i.Latlng = *temp.Latlng
    i.Name = *temp.Name
    i.RftemplateName = temp.RftemplateName
    i.SitegroupNames = temp.SitegroupNames
    i.Timezone = temp.Timezone
    return nil
}

// tempInstallerSite is a temporary struct used for validating the fields of InstallerSite.
type tempInstallerSite  struct {
    Address        *string    `json:"address"`
    CountryCode    *string    `json:"country_code"`
    Id             *uuid.UUID `json:"id,omitempty"`
    Latlng         *LatLng    `json:"latlng"`
    Name           *string    `json:"name"`
    RftemplateName *string    `json:"rftemplate_name,omitempty"`
    SitegroupNames []string   `json:"sitegroup_names,omitempty"`
    Timezone       *string    `json:"timezone,omitempty"`
}

func (i *tempInstallerSite) validate() error {
    var errs []string
    if i.Address == nil {
        errs = append(errs, "required field `address` is missing for type `installer_site`")
    }
    if i.CountryCode == nil {
        errs = append(errs, "required field `country_code` is missing for type `installer_site`")
    }
    if i.Latlng == nil {
        errs = append(errs, "required field `latlng` is missing for type `installer_site`")
    }
    if i.Name == nil {
        errs = append(errs, "required field `name` is missing for type `installer_site`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
