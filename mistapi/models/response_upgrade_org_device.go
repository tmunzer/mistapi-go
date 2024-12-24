package models

import (
    "encoding/json"
    "fmt"
    "github.com/google/uuid"
)

// ResponseUpgradeOrgDevice represents a ResponseUpgradeOrgDevice struct.
type ResponseUpgradeOrgDevice struct {
    SiteId               *uuid.UUID               `json:"site_id,omitempty"`
    Upgrade              *UpgradeOrgDeviceUpgrade `json:"upgrade,omitempty"`
    AdditionalProperties map[string]interface{}   `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseUpgradeOrgDevice,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseUpgradeOrgDevice) String() string {
    return fmt.Sprintf(
    	"ResponseUpgradeOrgDevice[SiteId=%v, Upgrade=%v, AdditionalProperties=%v]",
    	r.SiteId, r.Upgrade, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseUpgradeOrgDevice.
// It customizes the JSON marshaling process for ResponseUpgradeOrgDevice objects.
func (r ResponseUpgradeOrgDevice) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "site_id", "upgrade"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseUpgradeOrgDevice object to a map representation for JSON marshaling.
func (r ResponseUpgradeOrgDevice) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    if r.SiteId != nil {
        structMap["site_id"] = r.SiteId
    }
    if r.Upgrade != nil {
        structMap["upgrade"] = r.Upgrade.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseUpgradeOrgDevice.
// It customizes the JSON unmarshaling process for ResponseUpgradeOrgDevice objects.
func (r *ResponseUpgradeOrgDevice) UnmarshalJSON(input []byte) error {
    var temp tempResponseUpgradeOrgDevice
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "site_id", "upgrade")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.SiteId = temp.SiteId
    r.Upgrade = temp.Upgrade
    return nil
}

// tempResponseUpgradeOrgDevice is a temporary struct used for validating the fields of ResponseUpgradeOrgDevice.
type tempResponseUpgradeOrgDevice  struct {
    SiteId  *uuid.UUID               `json:"site_id,omitempty"`
    Upgrade *UpgradeOrgDeviceUpgrade `json:"upgrade,omitempty"`
}
