// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
)

// UpgradeOrgDevicesUpgrade represents a UpgradeOrgDevicesUpgrade struct.
type UpgradeOrgDevicesUpgrade struct {
	SiteId               *uuid.UUID                    `json:"site_id,omitempty"`
	Upgrade              *UpgradeOrgDevicesUpgradeInfo `json:"upgrade,omitempty"`
	AdditionalProperties map[string]interface{}        `json:"_"`
}

// String implements the fmt.Stringer interface for UpgradeOrgDevicesUpgrade,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UpgradeOrgDevicesUpgrade) String() string {
	return fmt.Sprintf(
		"UpgradeOrgDevicesUpgrade[SiteId=%v, Upgrade=%v, AdditionalProperties=%v]",
		u.SiteId, u.Upgrade, u.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for UpgradeOrgDevicesUpgrade.
// It customizes the JSON marshaling process for UpgradeOrgDevicesUpgrade objects.
func (u UpgradeOrgDevicesUpgrade) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(u.AdditionalProperties,
		"site_id", "upgrade"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(u.toMap())
}

// toMap converts the UpgradeOrgDevicesUpgrade object to a map representation for JSON marshaling.
func (u UpgradeOrgDevicesUpgrade) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, u.AdditionalProperties)
	if u.SiteId != nil {
		structMap["site_id"] = u.SiteId
	}
	if u.Upgrade != nil {
		structMap["upgrade"] = u.Upgrade.toMap()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpgradeOrgDevicesUpgrade.
// It customizes the JSON unmarshaling process for UpgradeOrgDevicesUpgrade objects.
func (u *UpgradeOrgDevicesUpgrade) UnmarshalJSON(input []byte) error {
	var temp tempUpgradeOrgDevicesUpgrade
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "site_id", "upgrade")
	if err != nil {
		return err
	}
	u.AdditionalProperties = additionalProperties

	u.SiteId = temp.SiteId
	u.Upgrade = temp.Upgrade
	return nil
}

// tempUpgradeOrgDevicesUpgrade is a temporary struct used for validating the fields of UpgradeOrgDevicesUpgrade.
type tempUpgradeOrgDevicesUpgrade struct {
	SiteId  *uuid.UUID                    `json:"site_id,omitempty"`
	Upgrade *UpgradeOrgDevicesUpgradeInfo `json:"upgrade,omitempty"`
}
