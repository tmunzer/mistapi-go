// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
)

// JsInventoryItem represents a JsInventoryItem struct.
type JsInventoryItem struct {
	// Indicates if the device is claimed by any org
	Claimed *bool `json:"claimed,omitempty"`
	// Name of the device
	DeviceName *string `json:"device_name,omitempty"`
	// End of life time
	EolTime *int `json:"eol_time,omitempty"`
	// End of support time
	EosTime *int `json:"eos_time,omitempty"`
	// Indicates if the device is covered under active support contract
	HasSupport *bool `json:"has_support,omitempty"`
	// Indicates whether it is Master
	Master *bool `json:"master,omitempty"`
	// Model of the install base inventory
	Model *string    `json:"model,omitempty"`
	OrgId *uuid.UUID `json:"org_id,omitempty"`
	// Serial Number of the inventory
	Serial *string `json:"serial,omitempty"`
	// Serviceable device stock
	Sku *string `json:"sku,omitempty"`
	// Status of the connected device
	Status *string `json:"status,omitempty"`
	// Suggested SW version
	SuggestedVersion *string `json:"suggested_version,omitempty"`
	// enum: `ap`, `gateway`, `switch`
	Type *DeviceTypeEnum `json:"type,omitempty"`
	// SW version running
	Version *string `json:"version,omitempty"`
	// End of Service of version
	VersionEosTime *int `json:"version_eos_time,omitempty"`
	// FRS date of the version
	VersionTime *int `json:"version_time,omitempty"`
	// warranty description
	Warranty *string `json:"warranty,omitempty"`
	// Time when warranty needs to be renewed
	WarrantyTime *int `json:"warranty_time,omitempty"`
	// Warranty type for Juniper Support Insight (JSI) devices. The warranty type
	// is used to determine the support level and duration of the warranty for the
	// device. enum:
	// * WTY00001: Standard Hardware Warranty
	// * WTY00002: Enhanced Hardware Warranty
	// * WTY00003: Dead On Arrival Warranty
	// * WTY00004: Limited Lifetime Warranty
	// * WTY00005: Software Warranty
	// * WTY00006: Limited Lifetime Warranty for WLA
	// * WTY00007: Warranty-JCPO EOL (DOA Not Included)
	// * WTY00008: MIST Enhanced Hardware Warranty
	// * WTY00009: MIST Standard Warranty
	// * WTY00099: Determine Lifetime warranty
	WarrantyType         *JsiWarrantyTypeEnum   `json:"warranty_type,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for JsInventoryItem,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (j JsInventoryItem) String() string {
	return fmt.Sprintf(
		"JsInventoryItem[Claimed=%v, DeviceName=%v, EolTime=%v, EosTime=%v, HasSupport=%v, Master=%v, Model=%v, OrgId=%v, Serial=%v, Sku=%v, Status=%v, SuggestedVersion=%v, Type=%v, Version=%v, VersionEosTime=%v, VersionTime=%v, Warranty=%v, WarrantyTime=%v, WarrantyType=%v, AdditionalProperties=%v]",
		j.Claimed, j.DeviceName, j.EolTime, j.EosTime, j.HasSupport, j.Master, j.Model, j.OrgId, j.Serial, j.Sku, j.Status, j.SuggestedVersion, j.Type, j.Version, j.VersionEosTime, j.VersionTime, j.Warranty, j.WarrantyTime, j.WarrantyType, j.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for JsInventoryItem.
// It customizes the JSON marshaling process for JsInventoryItem objects.
func (j JsInventoryItem) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(j.AdditionalProperties,
		"claimed", "device_name", "eol_time", "eos_time", "has_support", "master", "model", "org_id", "serial", "sku", "status", "suggested_version", "type", "version", "version_eos_time", "version_time", "warranty", "warranty_time", "warranty_type"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(j.toMap())
}

// toMap converts the JsInventoryItem object to a map representation for JSON marshaling.
func (j JsInventoryItem) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, j.AdditionalProperties)
	if j.Claimed != nil {
		structMap["claimed"] = j.Claimed
	}
	if j.DeviceName != nil {
		structMap["device_name"] = j.DeviceName
	}
	if j.EolTime != nil {
		structMap["eol_time"] = j.EolTime
	}
	if j.EosTime != nil {
		structMap["eos_time"] = j.EosTime
	}
	if j.HasSupport != nil {
		structMap["has_support"] = j.HasSupport
	}
	if j.Master != nil {
		structMap["master"] = j.Master
	}
	if j.Model != nil {
		structMap["model"] = j.Model
	}
	if j.OrgId != nil {
		structMap["org_id"] = j.OrgId
	}
	if j.Serial != nil {
		structMap["serial"] = j.Serial
	}
	if j.Sku != nil {
		structMap["sku"] = j.Sku
	}
	if j.Status != nil {
		structMap["status"] = j.Status
	}
	if j.SuggestedVersion != nil {
		structMap["suggested_version"] = j.SuggestedVersion
	}
	if j.Type != nil {
		structMap["type"] = j.Type
	}
	if j.Version != nil {
		structMap["version"] = j.Version
	}
	if j.VersionEosTime != nil {
		structMap["version_eos_time"] = j.VersionEosTime
	}
	if j.VersionTime != nil {
		structMap["version_time"] = j.VersionTime
	}
	if j.Warranty != nil {
		structMap["warranty"] = j.Warranty
	}
	if j.WarrantyTime != nil {
		structMap["warranty_time"] = j.WarrantyTime
	}
	if j.WarrantyType != nil {
		structMap["warranty_type"] = j.WarrantyType
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for JsInventoryItem.
// It customizes the JSON unmarshaling process for JsInventoryItem objects.
func (j *JsInventoryItem) UnmarshalJSON(input []byte) error {
	var temp tempJsInventoryItem
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "claimed", "device_name", "eol_time", "eos_time", "has_support", "master", "model", "org_id", "serial", "sku", "status", "suggested_version", "type", "version", "version_eos_time", "version_time", "warranty", "warranty_time", "warranty_type")
	if err != nil {
		return err
	}
	j.AdditionalProperties = additionalProperties

	j.Claimed = temp.Claimed
	j.DeviceName = temp.DeviceName
	j.EolTime = temp.EolTime
	j.EosTime = temp.EosTime
	j.HasSupport = temp.HasSupport
	j.Master = temp.Master
	j.Model = temp.Model
	j.OrgId = temp.OrgId
	j.Serial = temp.Serial
	j.Sku = temp.Sku
	j.Status = temp.Status
	j.SuggestedVersion = temp.SuggestedVersion
	j.Type = temp.Type
	j.Version = temp.Version
	j.VersionEosTime = temp.VersionEosTime
	j.VersionTime = temp.VersionTime
	j.Warranty = temp.Warranty
	j.WarrantyTime = temp.WarrantyTime
	j.WarrantyType = temp.WarrantyType
	return nil
}

// tempJsInventoryItem is a temporary struct used for validating the fields of JsInventoryItem.
type tempJsInventoryItem struct {
	Claimed          *bool                `json:"claimed,omitempty"`
	DeviceName       *string              `json:"device_name,omitempty"`
	EolTime          *int                 `json:"eol_time,omitempty"`
	EosTime          *int                 `json:"eos_time,omitempty"`
	HasSupport       *bool                `json:"has_support,omitempty"`
	Master           *bool                `json:"master,omitempty"`
	Model            *string              `json:"model,omitempty"`
	OrgId            *uuid.UUID           `json:"org_id,omitempty"`
	Serial           *string              `json:"serial,omitempty"`
	Sku              *string              `json:"sku,omitempty"`
	Status           *string              `json:"status,omitempty"`
	SuggestedVersion *string              `json:"suggested_version,omitempty"`
	Type             *DeviceTypeEnum      `json:"type,omitempty"`
	Version          *string              `json:"version,omitempty"`
	VersionEosTime   *int                 `json:"version_eos_time,omitempty"`
	VersionTime      *int                 `json:"version_time,omitempty"`
	Warranty         *string              `json:"warranty,omitempty"`
	WarrantyTime     *int                 `json:"warranty_time,omitempty"`
	WarrantyType     *JsiWarrantyTypeEnum `json:"warranty_type,omitempty"`
}
