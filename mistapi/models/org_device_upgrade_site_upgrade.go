package models

import (
    "encoding/json"
    "github.com/google/uuid"
)

// OrgDeviceUpgradeSiteUpgrade represents a OrgDeviceUpgradeSiteUpgrade struct.
type OrgDeviceUpgradeSiteUpgrade struct {
    SiteId               *uuid.UUID             `json:"site_id,omitempty"`
    UpgradeId            *uuid.UUID             `json:"upgrade_id,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for OrgDeviceUpgradeSiteUpgrade.
// It customizes the JSON marshaling process for OrgDeviceUpgradeSiteUpgrade objects.
func (o OrgDeviceUpgradeSiteUpgrade) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(o.AdditionalProperties,
        "site_id", "upgrade_id"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(o.toMap())
}

// toMap converts the OrgDeviceUpgradeSiteUpgrade object to a map representation for JSON marshaling.
func (o OrgDeviceUpgradeSiteUpgrade) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, o.AdditionalProperties)
    if o.SiteId != nil {
        structMap["site_id"] = o.SiteId
    }
    if o.UpgradeId != nil {
        structMap["upgrade_id"] = o.UpgradeId
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgDeviceUpgradeSiteUpgrade.
// It customizes the JSON unmarshaling process for OrgDeviceUpgradeSiteUpgrade objects.
func (o *OrgDeviceUpgradeSiteUpgrade) UnmarshalJSON(input []byte) error {
    var temp tempOrgDeviceUpgradeSiteUpgrade
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "site_id", "upgrade_id")
    if err != nil {
    	return err
    }
    o.AdditionalProperties = additionalProperties
    
    o.SiteId = temp.SiteId
    o.UpgradeId = temp.UpgradeId
    return nil
}

// tempOrgDeviceUpgradeSiteUpgrade is a temporary struct used for validating the fields of OrgDeviceUpgradeSiteUpgrade.
type tempOrgDeviceUpgradeSiteUpgrade  struct {
    SiteId    *uuid.UUID `json:"site_id,omitempty"`
    UpgradeId *uuid.UUID `json:"upgrade_id,omitempty"`
}
