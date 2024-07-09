package models

import (
    "encoding/json"
    "github.com/google/uuid"
)

// OrgDeviceUpgrade represents a OrgDeviceUpgrade struct.
type OrgDeviceUpgrade struct {
    Id                   *uuid.UUID                    `json:"id,omitempty"`
    SiteUpgrades         []OrgDeviceUpgradeSiteUpgrade `json:"site_upgrades,omitempty"`
    AdditionalProperties map[string]any                `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for OrgDeviceUpgrade.
// It customizes the JSON marshaling process for OrgDeviceUpgrade objects.
func (o OrgDeviceUpgrade) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(o.toMap())
}

// toMap converts the OrgDeviceUpgrade object to a map representation for JSON marshaling.
func (o OrgDeviceUpgrade) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, o.AdditionalProperties)
    if o.Id != nil {
        structMap["id"] = o.Id
    }
    if o.SiteUpgrades != nil {
        structMap["site_upgrades"] = o.SiteUpgrades
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgDeviceUpgrade.
// It customizes the JSON unmarshaling process for OrgDeviceUpgrade objects.
func (o *OrgDeviceUpgrade) UnmarshalJSON(input []byte) error {
    var temp orgDeviceUpgrade
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "id", "site_upgrades")
    if err != nil {
    	return err
    }
    
    o.AdditionalProperties = additionalProperties
    o.Id = temp.Id
    o.SiteUpgrades = temp.SiteUpgrades
    return nil
}

// orgDeviceUpgrade is a temporary struct used for validating the fields of OrgDeviceUpgrade.
type orgDeviceUpgrade  struct {
    Id           *uuid.UUID                    `json:"id,omitempty"`
    SiteUpgrades []OrgDeviceUpgradeSiteUpgrade `json:"site_upgrades,omitempty"`
}
