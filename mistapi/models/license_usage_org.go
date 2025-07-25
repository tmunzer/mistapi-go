// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/google/uuid"
    "strings"
)

// LicenseUsageOrg represents a LicenseUsageOrg struct.
type LicenseUsageOrg struct {
    ForSite              *bool                  `json:"for_site,omitempty"`
    // Maximum number of licenses that may be required if the service is enabled on all the Organization Devices. Property key is the service name (e.g. "SUB-MAN").
    FullyLoaded          map[string]int         `json:"fully_loaded,omitempty"`
    NumDevices           int                    `json:"num_devices"`
    SiteId               uuid.UUID              `json:"site_id"`
    // Number of licenses currently consumed. Property key is license type (e.g. SUB-MAN).
    Summary              map[string]int         `json:"summary,omitempty"`
    // Number of available licenes. Property key is the service name (e.g. "SUB-MAN"). name (e.g. "SUB-MAN")
    Usages               map[string]int         `json:"usages"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for LicenseUsageOrg,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (l LicenseUsageOrg) String() string {
    return fmt.Sprintf(
    	"LicenseUsageOrg[ForSite=%v, FullyLoaded=%v, NumDevices=%v, SiteId=%v, Summary=%v, Usages=%v, AdditionalProperties=%v]",
    	l.ForSite, l.FullyLoaded, l.NumDevices, l.SiteId, l.Summary, l.Usages, l.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for LicenseUsageOrg.
// It customizes the JSON marshaling process for LicenseUsageOrg objects.
func (l LicenseUsageOrg) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(l.AdditionalProperties,
        "for_site", "fully_loaded", "num_devices", "site_id", "summary", "usages"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(l.toMap())
}

// toMap converts the LicenseUsageOrg object to a map representation for JSON marshaling.
func (l LicenseUsageOrg) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, l.AdditionalProperties)
    if l.ForSite != nil {
        structMap["for_site"] = l.ForSite
    }
    if l.FullyLoaded != nil {
        structMap["fully_loaded"] = l.FullyLoaded
    }
    structMap["num_devices"] = l.NumDevices
    structMap["site_id"] = l.SiteId
    if l.Summary != nil {
        structMap["summary"] = l.Summary
    }
    structMap["usages"] = l.Usages
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for LicenseUsageOrg.
// It customizes the JSON unmarshaling process for LicenseUsageOrg objects.
func (l *LicenseUsageOrg) UnmarshalJSON(input []byte) error {
    var temp tempLicenseUsageOrg
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "for_site", "fully_loaded", "num_devices", "site_id", "summary", "usages")
    if err != nil {
    	return err
    }
    l.AdditionalProperties = additionalProperties
    
    l.ForSite = temp.ForSite
    l.FullyLoaded = temp.FullyLoaded
    l.NumDevices = *temp.NumDevices
    l.SiteId = *temp.SiteId
    l.Summary = temp.Summary
    l.Usages = *temp.Usages
    return nil
}

// tempLicenseUsageOrg is a temporary struct used for validating the fields of LicenseUsageOrg.
type tempLicenseUsageOrg  struct {
    ForSite     *bool           `json:"for_site,omitempty"`
    FullyLoaded map[string]int  `json:"fully_loaded,omitempty"`
    NumDevices  *int            `json:"num_devices"`
    SiteId      *uuid.UUID      `json:"site_id"`
    Summary     map[string]int  `json:"summary,omitempty"`
    Usages      *map[string]int `json:"usages"`
}

func (l *tempLicenseUsageOrg) validate() error {
    var errs []string
    if l.NumDevices == nil {
        errs = append(errs, "required field `num_devices` is missing for type `license_usage_org`")
    }
    if l.SiteId == nil {
        errs = append(errs, "required field `site_id` is missing for type `license_usage_org`")
    }
    if l.Usages == nil {
        errs = append(errs, "required field `usages` is missing for type `license_usage_org`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
