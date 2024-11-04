package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// Template represents a Template struct.
// Template
type Template struct {
    // where this template should be applied to, can be org_id, site_ids, sitegroup_ids
    Applies               *TemplateApplies    `json:"applies,omitempty"`
    // when the object has been created, in epoch
    CreatedTime           *float64            `json:"created_time,omitempty"`
    // list of Device Profile ids
    DeviceprofileIds      []uuid.UUID         `json:"deviceprofile_ids,omitempty"`
    // where this template should not be applied to (takes precedence)
    Exceptions            *TemplateExceptions `json:"exceptions,omitempty"`
    // whether to further filter by Device Profile
    FilterByDeviceprofile *bool               `json:"filter_by_deviceprofile,omitempty"`
    // Unique ID of the object instance in the Mist Organnization
    Id                    *uuid.UUID          `json:"id,omitempty"`
    // when the object has been modified for the last time, in epoch
    ModifiedTime          *float64            `json:"modified_time,omitempty"`
    Name                  string              `json:"name"`
    OrgId                 *uuid.UUID          `json:"org_id,omitempty"`
    AdditionalProperties  map[string]any      `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for Template.
// It customizes the JSON marshaling process for Template objects.
func (t Template) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(t.toMap())
}

// toMap converts the Template object to a map representation for JSON marshaling.
func (t Template) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, t.AdditionalProperties)
    if t.Applies != nil {
        structMap["applies"] = t.Applies.toMap()
    }
    if t.CreatedTime != nil {
        structMap["created_time"] = t.CreatedTime
    }
    if t.DeviceprofileIds != nil {
        structMap["deviceprofile_ids"] = t.DeviceprofileIds
    }
    if t.Exceptions != nil {
        structMap["exceptions"] = t.Exceptions.toMap()
    }
    if t.FilterByDeviceprofile != nil {
        structMap["filter_by_deviceprofile"] = t.FilterByDeviceprofile
    }
    if t.Id != nil {
        structMap["id"] = t.Id
    }
    if t.ModifiedTime != nil {
        structMap["modified_time"] = t.ModifiedTime
    }
    structMap["name"] = t.Name
    if t.OrgId != nil {
        structMap["org_id"] = t.OrgId
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Template.
// It customizes the JSON unmarshaling process for Template objects.
func (t *Template) UnmarshalJSON(input []byte) error {
    var temp tempTemplate
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "applies", "created_time", "deviceprofile_ids", "exceptions", "filter_by_deviceprofile", "id", "modified_time", "name", "org_id")
    if err != nil {
    	return err
    }
    
    t.AdditionalProperties = additionalProperties
    t.Applies = temp.Applies
    t.CreatedTime = temp.CreatedTime
    t.DeviceprofileIds = temp.DeviceprofileIds
    t.Exceptions = temp.Exceptions
    t.FilterByDeviceprofile = temp.FilterByDeviceprofile
    t.Id = temp.Id
    t.ModifiedTime = temp.ModifiedTime
    t.Name = *temp.Name
    t.OrgId = temp.OrgId
    return nil
}

// tempTemplate is a temporary struct used for validating the fields of Template.
type tempTemplate  struct {
    Applies               *TemplateApplies    `json:"applies,omitempty"`
    CreatedTime           *float64            `json:"created_time,omitempty"`
    DeviceprofileIds      []uuid.UUID         `json:"deviceprofile_ids,omitempty"`
    Exceptions            *TemplateExceptions `json:"exceptions,omitempty"`
    FilterByDeviceprofile *bool               `json:"filter_by_deviceprofile,omitempty"`
    Id                    *uuid.UUID          `json:"id,omitempty"`
    ModifiedTime          *float64            `json:"modified_time,omitempty"`
    Name                  *string             `json:"name"`
    OrgId                 *uuid.UUID          `json:"org_id,omitempty"`
}

func (t *tempTemplate) validate() error {
    var errs []string
    if t.Name == nil {
        errs = append(errs, "required field `name` is missing for type `template`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
