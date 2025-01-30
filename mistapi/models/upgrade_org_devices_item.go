package models

import (
    "encoding/json"
    "fmt"
    "github.com/google/uuid"
)

// UpgradeOrgDevicesItem represents a UpgradeOrgDevicesItem struct.
type UpgradeOrgDevicesItem struct {
    // Unique ID of the object instance in the Mist Organnization
    Id                   *uuid.UUID                         `json:"id,omitempty"`
    SiteUpgrades         []UpgradeOrgDevicesItemSiteUpgrade `json:"site_upgrades,omitempty"`
    AdditionalProperties map[string]interface{}             `json:"_"`
}

// String implements the fmt.Stringer interface for UpgradeOrgDevicesItem,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UpgradeOrgDevicesItem) String() string {
    return fmt.Sprintf(
    	"UpgradeOrgDevicesItem[Id=%v, SiteUpgrades=%v, AdditionalProperties=%v]",
    	u.Id, u.SiteUpgrades, u.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for UpgradeOrgDevicesItem.
// It customizes the JSON marshaling process for UpgradeOrgDevicesItem objects.
func (u UpgradeOrgDevicesItem) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(u.AdditionalProperties,
        "id", "site_upgrades"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UpgradeOrgDevicesItem object to a map representation for JSON marshaling.
func (u UpgradeOrgDevicesItem) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, u.AdditionalProperties)
    if u.Id != nil {
        structMap["id"] = u.Id
    }
    if u.SiteUpgrades != nil {
        structMap["site_upgrades"] = u.SiteUpgrades
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpgradeOrgDevicesItem.
// It customizes the JSON unmarshaling process for UpgradeOrgDevicesItem objects.
func (u *UpgradeOrgDevicesItem) UnmarshalJSON(input []byte) error {
    var temp tempUpgradeOrgDevicesItem
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "id", "site_upgrades")
    if err != nil {
    	return err
    }
    u.AdditionalProperties = additionalProperties
    
    u.Id = temp.Id
    u.SiteUpgrades = temp.SiteUpgrades
    return nil
}

// tempUpgradeOrgDevicesItem is a temporary struct used for validating the fields of UpgradeOrgDevicesItem.
type tempUpgradeOrgDevicesItem  struct {
    Id           *uuid.UUID                         `json:"id,omitempty"`
    SiteUpgrades []UpgradeOrgDevicesItemSiteUpgrade `json:"site_upgrades,omitempty"`
}
