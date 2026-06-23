// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
)

// InventorySearchResult represents a InventorySearchResult struct.
// Inventory record returned by inventory search
type InventorySearchResult struct {
	// Device MAC address for this inventory search result
	Mac *string `json:"mac,omitempty"`
	// Whether this search result represents the master member of a Virtual Chassis
	Master *bool `json:"master,omitempty"`
	// Virtual Chassis members included in an inventory search result
	Members []InventorySearchResultMember `json:"members,omitempty"`
	// Device model for this inventory search result
	Model *string `json:"model,omitempty"`
	// Configured device name in this inventory search result
	Name *string `json:"name,omitempty"`
	// Unique identifier of a Mist organization
	OrgId *uuid.UUID `json:"org_id,omitempty"`
	// Device serial number for this inventory search result
	Serial *string `json:"serial,omitempty"`
	// Unique identifier of a Mist site
	SiteId *uuid.UUID `json:"site_id,omitempty"`
	// Device SKU for this inventory search result
	Sku *string `json:"sku,omitempty"`
	// Current inventory status for the device
	Status *string `json:"status,omitempty"`
	// enum: `ap`, `gateway`, `switch`
	Type *DeviceTypeDefaultApEnum `json:"type,omitempty"`
	// Virtual Chassis MAC address for this inventory search result
	VcMac *string `json:"vc_mac,omitempty"`
	// Software version reported for this inventory search result
	Version              *string                `json:"version,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for InventorySearchResult,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (i InventorySearchResult) String() string {
	return fmt.Sprintf(
		"InventorySearchResult[Mac=%v, Master=%v, Members=%v, Model=%v, Name=%v, OrgId=%v, Serial=%v, SiteId=%v, Sku=%v, Status=%v, Type=%v, VcMac=%v, Version=%v, AdditionalProperties=%v]",
		i.Mac, i.Master, i.Members, i.Model, i.Name, i.OrgId, i.Serial, i.SiteId, i.Sku, i.Status, i.Type, i.VcMac, i.Version, i.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for InventorySearchResult.
// It customizes the JSON marshaling process for InventorySearchResult objects.
func (i InventorySearchResult) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(i.AdditionalProperties,
		"mac", "master", "members", "model", "name", "org_id", "serial", "site_id", "sku", "status", "type", "vc_mac", "version"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(i.toMap())
}

// toMap converts the InventorySearchResult object to a map representation for JSON marshaling.
func (i InventorySearchResult) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, i.AdditionalProperties)
	if i.Mac != nil {
		structMap["mac"] = i.Mac
	}
	if i.Master != nil {
		structMap["master"] = i.Master
	}
	if i.Members != nil {
		structMap["members"] = i.Members
	}
	if i.Model != nil {
		structMap["model"] = i.Model
	}
	if i.Name != nil {
		structMap["name"] = i.Name
	}
	if i.OrgId != nil {
		structMap["org_id"] = i.OrgId
	}
	if i.Serial != nil {
		structMap["serial"] = i.Serial
	}
	if i.SiteId != nil {
		structMap["site_id"] = i.SiteId
	}
	if i.Sku != nil {
		structMap["sku"] = i.Sku
	}
	if i.Status != nil {
		structMap["status"] = i.Status
	}
	if i.Type != nil {
		structMap["type"] = i.Type
	}
	if i.VcMac != nil {
		structMap["vc_mac"] = i.VcMac
	}
	if i.Version != nil {
		structMap["version"] = i.Version
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for InventorySearchResult.
// It customizes the JSON unmarshaling process for InventorySearchResult objects.
func (i *InventorySearchResult) UnmarshalJSON(input []byte) error {
	var temp tempInventorySearchResult
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "mac", "master", "members", "model", "name", "org_id", "serial", "site_id", "sku", "status", "type", "vc_mac", "version")
	if err != nil {
		return err
	}
	i.AdditionalProperties = additionalProperties

	i.Mac = temp.Mac
	i.Master = temp.Master
	i.Members = temp.Members
	i.Model = temp.Model
	i.Name = temp.Name
	i.OrgId = temp.OrgId
	i.Serial = temp.Serial
	i.SiteId = temp.SiteId
	i.Sku = temp.Sku
	i.Status = temp.Status
	i.Type = temp.Type
	i.VcMac = temp.VcMac
	i.Version = temp.Version
	return nil
}

// tempInventorySearchResult is a temporary struct used for validating the fields of InventorySearchResult.
type tempInventorySearchResult struct {
	Mac     *string                       `json:"mac,omitempty"`
	Master  *bool                         `json:"master,omitempty"`
	Members []InventorySearchResultMember `json:"members,omitempty"`
	Model   *string                       `json:"model,omitempty"`
	Name    *string                       `json:"name,omitempty"`
	OrgId   *uuid.UUID                    `json:"org_id,omitempty"`
	Serial  *string                       `json:"serial,omitempty"`
	SiteId  *uuid.UUID                    `json:"site_id,omitempty"`
	Sku     *string                       `json:"sku,omitempty"`
	Status  *string                       `json:"status,omitempty"`
	Type    *DeviceTypeDefaultApEnum      `json:"type,omitempty"`
	VcMac   *string                       `json:"vc_mac,omitempty"`
	Version *string                       `json:"version,omitempty"`
}
