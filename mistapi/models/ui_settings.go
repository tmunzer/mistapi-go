package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// UiSettings represents a UiSettings struct.
// UI Settings
type UiSettings struct {
    // when the object has been created, in epoch
    CreatedTime          *float64                    `json:"created_time,omitempty"`
    DefaultScopeId       *string                     `json:"defaultScopeId,omitempty"`
    DefaultScopeType     *string                     `json:"defaultScopeType,omitempty"`
    DefaultTimeRange     *UiSettingsDefaultTimeRange `json:"defaultTimeRange,omitempty"`
    Description          string                      `json:"description"`
    ForSite              *bool                       `json:"for_site,omitempty"`
    // Unique ID of the object instance in the Mist Organnization
    Id                   *uuid.UUID                  `json:"id,omitempty"`
    IsCustomDataboard    *bool                       `json:"isCustomDataboard,omitempty"`
    IsScopeLinked        *bool                       `json:"isScopeLinked,omitempty"`
    IsTimeRangeLinked    *bool                       `json:"isTimeRangeLinked,omitempty"`
    // when the object has been modified for the last time, in epoch
    ModifiedTime         *float64                    `json:"modified_time,omitempty"`
    Name                 *string                     `json:"name,omitempty"`
    OrgId                *uuid.UUID                  `json:"org_id,omitempty"`
    Purpose              string                      `json:"purpose"`
    SiteId               *uuid.UUID                  `json:"site_id,omitempty"`
    Tiles                []UiSettingsTile            `json:"tiles,omitempty"`
    AdditionalProperties map[string]any              `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for UiSettings.
// It customizes the JSON marshaling process for UiSettings objects.
func (u UiSettings) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(u.toMap())
}

// toMap converts the UiSettings object to a map representation for JSON marshaling.
func (u UiSettings) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, u.AdditionalProperties)
    if u.CreatedTime != nil {
        structMap["created_time"] = u.CreatedTime
    }
    if u.DefaultScopeId != nil {
        structMap["defaultScopeId"] = u.DefaultScopeId
    }
    if u.DefaultScopeType != nil {
        structMap["defaultScopeType"] = u.DefaultScopeType
    }
    if u.DefaultTimeRange != nil {
        structMap["defaultTimeRange"] = u.DefaultTimeRange.toMap()
    }
    structMap["description"] = u.Description
    if u.ForSite != nil {
        structMap["for_site"] = u.ForSite
    }
    if u.Id != nil {
        structMap["id"] = u.Id
    }
    if u.IsCustomDataboard != nil {
        structMap["isCustomDataboard"] = u.IsCustomDataboard
    }
    if u.IsScopeLinked != nil {
        structMap["isScopeLinked"] = u.IsScopeLinked
    }
    if u.IsTimeRangeLinked != nil {
        structMap["isTimeRangeLinked"] = u.IsTimeRangeLinked
    }
    if u.ModifiedTime != nil {
        structMap["modified_time"] = u.ModifiedTime
    }
    if u.Name != nil {
        structMap["name"] = u.Name
    }
    if u.OrgId != nil {
        structMap["org_id"] = u.OrgId
    }
    structMap["purpose"] = u.Purpose
    if u.SiteId != nil {
        structMap["site_id"] = u.SiteId
    }
    if u.Tiles != nil {
        structMap["tiles"] = u.Tiles
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UiSettings.
// It customizes the JSON unmarshaling process for UiSettings objects.
func (u *UiSettings) UnmarshalJSON(input []byte) error {
    var temp tempUiSettings
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "created_time", "defaultScopeId", "defaultScopeType", "defaultTimeRange", "description", "for_site", "id", "isCustomDataboard", "isScopeLinked", "isTimeRangeLinked", "modified_time", "name", "org_id", "purpose", "site_id", "tiles")
    if err != nil {
    	return err
    }
    
    u.AdditionalProperties = additionalProperties
    u.CreatedTime = temp.CreatedTime
    u.DefaultScopeId = temp.DefaultScopeId
    u.DefaultScopeType = temp.DefaultScopeType
    u.DefaultTimeRange = temp.DefaultTimeRange
    u.Description = *temp.Description
    u.ForSite = temp.ForSite
    u.Id = temp.Id
    u.IsCustomDataboard = temp.IsCustomDataboard
    u.IsScopeLinked = temp.IsScopeLinked
    u.IsTimeRangeLinked = temp.IsTimeRangeLinked
    u.ModifiedTime = temp.ModifiedTime
    u.Name = temp.Name
    u.OrgId = temp.OrgId
    u.Purpose = *temp.Purpose
    u.SiteId = temp.SiteId
    u.Tiles = temp.Tiles
    return nil
}

// tempUiSettings is a temporary struct used for validating the fields of UiSettings.
type tempUiSettings  struct {
    CreatedTime       *float64                    `json:"created_time,omitempty"`
    DefaultScopeId    *string                     `json:"defaultScopeId,omitempty"`
    DefaultScopeType  *string                     `json:"defaultScopeType,omitempty"`
    DefaultTimeRange  *UiSettingsDefaultTimeRange `json:"defaultTimeRange,omitempty"`
    Description       *string                     `json:"description"`
    ForSite           *bool                       `json:"for_site,omitempty"`
    Id                *uuid.UUID                  `json:"id,omitempty"`
    IsCustomDataboard *bool                       `json:"isCustomDataboard,omitempty"`
    IsScopeLinked     *bool                       `json:"isScopeLinked,omitempty"`
    IsTimeRangeLinked *bool                       `json:"isTimeRangeLinked,omitempty"`
    ModifiedTime      *float64                    `json:"modified_time,omitempty"`
    Name              *string                     `json:"name,omitempty"`
    OrgId             *uuid.UUID                  `json:"org_id,omitempty"`
    Purpose           *string                     `json:"purpose"`
    SiteId            *uuid.UUID                  `json:"site_id,omitempty"`
    Tiles             []UiSettingsTile            `json:"tiles,omitempty"`
}

func (u *tempUiSettings) validate() error {
    var errs []string
    if u.Description == nil {
        errs = append(errs, "required field `description` is missing for type `ui_settings`")
    }
    if u.Purpose == nil {
        errs = append(errs, "required field `purpose` is missing for type `ui_settings`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
