package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// LicenseUsageOrg represents a LicenseUsageOrg struct.
type LicenseUsageOrg struct {
    ForSite              *bool          `json:"for_site,omitempty"`
    // Property key is the service name (e.g. "SUB-MAN")
    FullyLoaded          map[string]int `json:"fully_loaded,omitempty"`
    NumDevices           int            `json:"num_devices"`
    SiteId               uuid.UUID      `json:"site_id"`
    // subscriptions and their quantities. Property key is the service name (e.g. "SUB-MAN")
    Usages               map[string]int `json:"usages"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for LicenseUsageOrg.
// It customizes the JSON marshaling process for LicenseUsageOrg objects.
func (l LicenseUsageOrg) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(l.toMap())
}

// toMap converts the LicenseUsageOrg object to a map representation for JSON marshaling.
func (l LicenseUsageOrg) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, l.AdditionalProperties)
    if l.ForSite != nil {
        structMap["for_site"] = l.ForSite
    }
    if l.FullyLoaded != nil {
        structMap["fully_loaded"] = l.FullyLoaded
    }
    structMap["num_devices"] = l.NumDevices
    structMap["site_id"] = l.SiteId
    structMap["usages"] = l.Usages
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for LicenseUsageOrg.
// It customizes the JSON unmarshaling process for LicenseUsageOrg objects.
func (l *LicenseUsageOrg) UnmarshalJSON(input []byte) error {
    var temp licenseUsageOrg
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "for_site", "fully_loaded", "num_devices", "site_id", "usages")
    if err != nil {
    	return err
    }
    
    l.AdditionalProperties = additionalProperties
    l.ForSite = temp.ForSite
    l.FullyLoaded = temp.FullyLoaded
    l.NumDevices = *temp.NumDevices
    l.SiteId = *temp.SiteId
    l.Usages = *temp.Usages
    return nil
}

// licenseUsageOrg is a temporary struct used for validating the fields of LicenseUsageOrg.
type licenseUsageOrg  struct {
    ForSite     *bool           `json:"for_site,omitempty"`
    FullyLoaded map[string]int  `json:"fully_loaded,omitempty"`
    NumDevices  *int            `json:"num_devices"`
    SiteId      *uuid.UUID      `json:"site_id"`
    Usages      *map[string]int `json:"usages"`
}

func (l *licenseUsageOrg) validate() error {
    var errs []string
    if l.NumDevices == nil {
        errs = append(errs, "required field `num_devices` is missing for type `License_Usage_Org`")
    }
    if l.SiteId == nil {
        errs = append(errs, "required field `site_id` is missing for type `License_Usage_Org`")
    }
    if l.Usages == nil {
        errs = append(errs, "required field `usages` is missing for type `License_Usage_Org`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
