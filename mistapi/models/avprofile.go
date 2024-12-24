package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/google/uuid"
    "strings"
)

// Avprofile represents a Avprofile struct.
type Avprofile struct {
    // when the object has been created, in epoch
    CreatedTime          *float64                     `json:"created_time,omitempty"`
    // enum: `block`, `permit`
    FallbackAction       *AvprofileFallbackActionEnum `json:"fallback_action,omitempty"`
    // Unique ID of the object instance in the Mist Organnization
    Id                   *uuid.UUID                   `json:"id,omitempty"`
    // in KB
    MexFilesize          *int                         `json:"mex_filesize,omitempty"`
    MimeWhitelist        []string                     `json:"mime_whitelist,omitempty"`
    ModifiedTime         *float64                     `json:"modified_time,omitempty"`
    Name                 string                       `json:"name"`
    OrgId                *uuid.UUID                   `json:"org_id,omitempty"`
    Protocols            []AvprofileProtocolsEnum     `json:"protocols,omitempty"`
    SiteId               *uuid.UUID                   `json:"site_id,omitempty"`
    UrlWhitelist         []string                     `json:"url_whitelist,omitempty"`
    AdditionalProperties map[string]interface{}       `json:"_"`
}

// String implements the fmt.Stringer interface for Avprofile,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a Avprofile) String() string {
    return fmt.Sprintf(
    	"Avprofile[CreatedTime=%v, FallbackAction=%v, Id=%v, MexFilesize=%v, MimeWhitelist=%v, ModifiedTime=%v, Name=%v, OrgId=%v, Protocols=%v, SiteId=%v, UrlWhitelist=%v, AdditionalProperties=%v]",
    	a.CreatedTime, a.FallbackAction, a.Id, a.MexFilesize, a.MimeWhitelist, a.ModifiedTime, a.Name, a.OrgId, a.Protocols, a.SiteId, a.UrlWhitelist, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for Avprofile.
// It customizes the JSON marshaling process for Avprofile objects.
func (a Avprofile) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(a.AdditionalProperties,
        "created_time", "fallback_action", "id", "mex_filesize", "mime_whitelist", "modified_time", "name", "org_id", "protocols", "site_id", "url_whitelist"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(a.toMap())
}

// toMap converts the Avprofile object to a map representation for JSON marshaling.
func (a Avprofile) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, a.AdditionalProperties)
    if a.CreatedTime != nil {
        structMap["created_time"] = a.CreatedTime
    }
    if a.FallbackAction != nil {
        structMap["fallback_action"] = a.FallbackAction
    }
    if a.Id != nil {
        structMap["id"] = a.Id
    }
    if a.MexFilesize != nil {
        structMap["mex_filesize"] = a.MexFilesize
    }
    if a.MimeWhitelist != nil {
        structMap["mime_whitelist"] = a.MimeWhitelist
    }
    if a.ModifiedTime != nil {
        structMap["modified_time"] = a.ModifiedTime
    }
    structMap["name"] = a.Name
    if a.OrgId != nil {
        structMap["org_id"] = a.OrgId
    }
    if a.Protocols != nil {
        structMap["protocols"] = a.Protocols
    }
    if a.SiteId != nil {
        structMap["site_id"] = a.SiteId
    }
    if a.UrlWhitelist != nil {
        structMap["url_whitelist"] = a.UrlWhitelist
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Avprofile.
// It customizes the JSON unmarshaling process for Avprofile objects.
func (a *Avprofile) UnmarshalJSON(input []byte) error {
    var temp tempAvprofile
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "created_time", "fallback_action", "id", "mex_filesize", "mime_whitelist", "modified_time", "name", "org_id", "protocols", "site_id", "url_whitelist")
    if err != nil {
    	return err
    }
    a.AdditionalProperties = additionalProperties
    
    a.CreatedTime = temp.CreatedTime
    a.FallbackAction = temp.FallbackAction
    a.Id = temp.Id
    a.MexFilesize = temp.MexFilesize
    a.MimeWhitelist = temp.MimeWhitelist
    a.ModifiedTime = temp.ModifiedTime
    a.Name = *temp.Name
    a.OrgId = temp.OrgId
    a.Protocols = temp.Protocols
    a.SiteId = temp.SiteId
    a.UrlWhitelist = temp.UrlWhitelist
    return nil
}

// tempAvprofile is a temporary struct used for validating the fields of Avprofile.
type tempAvprofile  struct {
    CreatedTime    *float64                     `json:"created_time,omitempty"`
    FallbackAction *AvprofileFallbackActionEnum `json:"fallback_action,omitempty"`
    Id             *uuid.UUID                   `json:"id,omitempty"`
    MexFilesize    *int                         `json:"mex_filesize,omitempty"`
    MimeWhitelist  []string                     `json:"mime_whitelist,omitempty"`
    ModifiedTime   *float64                     `json:"modified_time,omitempty"`
    Name           *string                      `json:"name"`
    OrgId          *uuid.UUID                   `json:"org_id,omitempty"`
    Protocols      []AvprofileProtocolsEnum     `json:"protocols,omitempty"`
    SiteId         *uuid.UUID                   `json:"site_id,omitempty"`
    UrlWhitelist   []string                     `json:"url_whitelist,omitempty"`
}

func (a *tempAvprofile) validate() error {
    var errs []string
    if a.Name == nil {
        errs = append(errs, "required field `name` is missing for type `avprofile`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
