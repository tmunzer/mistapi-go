// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
)

// ResponseUpgradeOrgDevices represents a ResponseUpgradeOrgDevices struct.
type ResponseUpgradeOrgDevices struct {
	// Whether to allow local AP-to-AP FW upgrade
	EnableP2p *bool `json:"enable_p2p,omitempty"`
	// Whether to force upgrade when requested version is same as running version
	Force *bool `json:"force,omitempty"`
	// Unique ID of the object instance in the Mist Organization
	Id *uuid.UUID `json:"id,omitempty"`
	// enum: `big_bang` (upgrade all at once), `canary`, `rrm` (APs only), `serial` (one at a time)
	Strategy *UpgradeDeviceStrategyEnum `json:"strategy,omitempty"`
	// Version to upgrade to
	TargetVersion        *string                    `json:"target_version,omitempty"`
	Upgrades             []UpgradeOrgDevicesUpgrade `json:"upgrades,omitempty"`
	AdditionalProperties map[string]interface{}     `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseUpgradeOrgDevices,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseUpgradeOrgDevices) String() string {
	return fmt.Sprintf(
		"ResponseUpgradeOrgDevices[EnableP2p=%v, Force=%v, Id=%v, Strategy=%v, TargetVersion=%v, Upgrades=%v, AdditionalProperties=%v]",
		r.EnableP2p, r.Force, r.Id, r.Strategy, r.TargetVersion, r.Upgrades, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseUpgradeOrgDevices.
// It customizes the JSON marshaling process for ResponseUpgradeOrgDevices objects.
func (r ResponseUpgradeOrgDevices) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(r.AdditionalProperties,
		"enable_p2p", "force", "id", "strategy", "target_version", "upgrades"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(r.toMap())
}

// toMap converts the ResponseUpgradeOrgDevices object to a map representation for JSON marshaling.
func (r ResponseUpgradeOrgDevices) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, r.AdditionalProperties)
	if r.EnableP2p != nil {
		structMap["enable_p2p"] = r.EnableP2p
	}
	if r.Force != nil {
		structMap["force"] = r.Force
	}
	if r.Id != nil {
		structMap["id"] = r.Id
	}
	if r.Strategy != nil {
		structMap["strategy"] = r.Strategy
	}
	if r.TargetVersion != nil {
		structMap["target_version"] = r.TargetVersion
	}
	if r.Upgrades != nil {
		structMap["upgrades"] = r.Upgrades
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseUpgradeOrgDevices.
// It customizes the JSON unmarshaling process for ResponseUpgradeOrgDevices objects.
func (r *ResponseUpgradeOrgDevices) UnmarshalJSON(input []byte) error {
	var temp tempResponseUpgradeOrgDevices
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "enable_p2p", "force", "id", "strategy", "target_version", "upgrades")
	if err != nil {
		return err
	}
	r.AdditionalProperties = additionalProperties

	r.EnableP2p = temp.EnableP2p
	r.Force = temp.Force
	r.Id = temp.Id
	r.Strategy = temp.Strategy
	r.TargetVersion = temp.TargetVersion
	r.Upgrades = temp.Upgrades
	return nil
}

// tempResponseUpgradeOrgDevices is a temporary struct used for validating the fields of ResponseUpgradeOrgDevices.
type tempResponseUpgradeOrgDevices struct {
	EnableP2p     *bool                      `json:"enable_p2p,omitempty"`
	Force         *bool                      `json:"force,omitempty"`
	Id            *uuid.UUID                 `json:"id,omitempty"`
	Strategy      *UpgradeDeviceStrategyEnum `json:"strategy,omitempty"`
	TargetVersion *string                    `json:"target_version,omitempty"`
	Upgrades      []UpgradeOrgDevicesUpgrade `json:"upgrades,omitempty"`
}
