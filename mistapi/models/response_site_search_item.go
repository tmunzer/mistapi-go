package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// ResponseSiteSearchItem represents a ResponseSiteSearchItem struct.
type ResponseSiteSearchItem struct {
    AutoUpgradeEnabled   bool                   `json:"auto_upgrade_enabled"`
    AutoUpgradeVersion   string                 `json:"auto_upgrade_version"`
    CountryCode          *string                `json:"country_code"`
    HoneypotEnabled      bool                   `json:"honeypot_enabled"`
    // Unique ID of the object instance in the Mist Organnization
    Id                   uuid.UUID              `json:"id"`
    Name                 string                 `json:"name"`
    OrgId                uuid.UUID              `json:"org_id"`
    SiteId               uuid.UUID              `json:"site_id"`
    Timestamp            float64                `json:"timestamp"`
    Timezone             string                 `json:"timezone"`
    VnaEnabled           bool                   `json:"vna_enabled"`
    WifiEnabled          bool                   `json:"wifi_enabled"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseSiteSearchItem.
// It customizes the JSON marshaling process for ResponseSiteSearchItem objects.
func (r ResponseSiteSearchItem) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "auto_upgrade_enabled", "auto_upgrade_version", "country_code", "honeypot_enabled", "id", "name", "org_id", "site_id", "timestamp", "timezone", "vna_enabled", "wifi_enabled"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseSiteSearchItem object to a map representation for JSON marshaling.
func (r ResponseSiteSearchItem) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    structMap["auto_upgrade_enabled"] = r.AutoUpgradeEnabled
    structMap["auto_upgrade_version"] = r.AutoUpgradeVersion
    if r.CountryCode != nil {
        structMap["country_code"] = r.CountryCode
    } else {
        structMap["country_code"] = nil
    }
    structMap["honeypot_enabled"] = r.HoneypotEnabled
    structMap["id"] = r.Id
    structMap["name"] = r.Name
    structMap["org_id"] = r.OrgId
    structMap["site_id"] = r.SiteId
    structMap["timestamp"] = r.Timestamp
    structMap["timezone"] = r.Timezone
    structMap["vna_enabled"] = r.VnaEnabled
    structMap["wifi_enabled"] = r.WifiEnabled
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseSiteSearchItem.
// It customizes the JSON unmarshaling process for ResponseSiteSearchItem objects.
func (r *ResponseSiteSearchItem) UnmarshalJSON(input []byte) error {
    var temp tempResponseSiteSearchItem
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "auto_upgrade_enabled", "auto_upgrade_version", "country_code", "honeypot_enabled", "id", "name", "org_id", "site_id", "timestamp", "timezone", "vna_enabled", "wifi_enabled")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.AutoUpgradeEnabled = *temp.AutoUpgradeEnabled
    r.AutoUpgradeVersion = *temp.AutoUpgradeVersion
    r.CountryCode = temp.CountryCode
    r.HoneypotEnabled = *temp.HoneypotEnabled
    r.Id = *temp.Id
    r.Name = *temp.Name
    r.OrgId = *temp.OrgId
    r.SiteId = *temp.SiteId
    r.Timestamp = *temp.Timestamp
    r.Timezone = *temp.Timezone
    r.VnaEnabled = *temp.VnaEnabled
    r.WifiEnabled = *temp.WifiEnabled
    return nil
}

// tempResponseSiteSearchItem is a temporary struct used for validating the fields of ResponseSiteSearchItem.
type tempResponseSiteSearchItem  struct {
    AutoUpgradeEnabled *bool      `json:"auto_upgrade_enabled"`
    AutoUpgradeVersion *string    `json:"auto_upgrade_version"`
    CountryCode        *string    `json:"country_code"`
    HoneypotEnabled    *bool      `json:"honeypot_enabled"`
    Id                 *uuid.UUID `json:"id"`
    Name               *string    `json:"name"`
    OrgId              *uuid.UUID `json:"org_id"`
    SiteId             *uuid.UUID `json:"site_id"`
    Timestamp          *float64   `json:"timestamp"`
    Timezone           *string    `json:"timezone"`
    VnaEnabled         *bool      `json:"vna_enabled"`
    WifiEnabled        *bool      `json:"wifi_enabled"`
}

func (r *tempResponseSiteSearchItem) validate() error {
    var errs []string
    if r.AutoUpgradeEnabled == nil {
        errs = append(errs, "required field `auto_upgrade_enabled` is missing for type `response_site_search_item`")
    }
    if r.AutoUpgradeVersion == nil {
        errs = append(errs, "required field `auto_upgrade_version` is missing for type `response_site_search_item`")
    }
    if r.HoneypotEnabled == nil {
        errs = append(errs, "required field `honeypot_enabled` is missing for type `response_site_search_item`")
    }
    if r.Id == nil {
        errs = append(errs, "required field `id` is missing for type `response_site_search_item`")
    }
    if r.Name == nil {
        errs = append(errs, "required field `name` is missing for type `response_site_search_item`")
    }
    if r.OrgId == nil {
        errs = append(errs, "required field `org_id` is missing for type `response_site_search_item`")
    }
    if r.SiteId == nil {
        errs = append(errs, "required field `site_id` is missing for type `response_site_search_item`")
    }
    if r.Timestamp == nil {
        errs = append(errs, "required field `timestamp` is missing for type `response_site_search_item`")
    }
    if r.Timezone == nil {
        errs = append(errs, "required field `timezone` is missing for type `response_site_search_item`")
    }
    if r.VnaEnabled == nil {
        errs = append(errs, "required field `vna_enabled` is missing for type `response_site_search_item`")
    }
    if r.WifiEnabled == nil {
        errs = append(errs, "required field `wifi_enabled` is missing for type `response_site_search_item`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
