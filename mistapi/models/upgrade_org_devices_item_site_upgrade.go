// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
)

// UpgradeOrgDevicesItemSiteUpgrade represents a UpgradeOrgDevicesItemSiteUpgrade struct.
type UpgradeOrgDevicesItemSiteUpgrade struct {
	SiteId               *uuid.UUID             `json:"site_id,omitempty"`
	UpgradeId            *uuid.UUID             `json:"upgrade_id,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for UpgradeOrgDevicesItemSiteUpgrade,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UpgradeOrgDevicesItemSiteUpgrade) String() string {
	return fmt.Sprintf(
		"UpgradeOrgDevicesItemSiteUpgrade[SiteId=%v, UpgradeId=%v, AdditionalProperties=%v]",
		u.SiteId, u.UpgradeId, u.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for UpgradeOrgDevicesItemSiteUpgrade.
// It customizes the JSON marshaling process for UpgradeOrgDevicesItemSiteUpgrade objects.
func (u UpgradeOrgDevicesItemSiteUpgrade) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(u.AdditionalProperties,
		"site_id", "upgrade_id"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(u.toMap())
}

// toMap converts the UpgradeOrgDevicesItemSiteUpgrade object to a map representation for JSON marshaling.
func (u UpgradeOrgDevicesItemSiteUpgrade) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, u.AdditionalProperties)
	if u.SiteId != nil {
		structMap["site_id"] = u.SiteId
	}
	if u.UpgradeId != nil {
		structMap["upgrade_id"] = u.UpgradeId
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpgradeOrgDevicesItemSiteUpgrade.
// It customizes the JSON unmarshaling process for UpgradeOrgDevicesItemSiteUpgrade objects.
func (u *UpgradeOrgDevicesItemSiteUpgrade) UnmarshalJSON(input []byte) error {
	var temp tempUpgradeOrgDevicesItemSiteUpgrade
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "site_id", "upgrade_id")
	if err != nil {
		return err
	}
	u.AdditionalProperties = additionalProperties

	u.SiteId = temp.SiteId
	u.UpgradeId = temp.UpgradeId
	return nil
}

// tempUpgradeOrgDevicesItemSiteUpgrade is a temporary struct used for validating the fields of UpgradeOrgDevicesItemSiteUpgrade.
type tempUpgradeOrgDevicesItemSiteUpgrade struct {
	SiteId    *uuid.UUID `json:"site_id,omitempty"`
	UpgradeId *uuid.UUID `json:"upgrade_id,omitempty"`
}
