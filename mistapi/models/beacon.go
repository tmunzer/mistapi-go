// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
)

// Beacon represents a Beacon struct.
// Beacon
type Beacon struct {
	// When the object has been created, in epoch
	CreatedTime *float64 `json:"created_time,omitempty"`
	// Eddystone-UID instance (6 bytes) in hexstring format
	EddystoneInstance *string `json:"eddystone_instance,omitempty"`
	// Eddystone-UID namespace (10 bytes) in hexstring format
	EddystoneNamespace *string `json:"eddystone_namespace,omitempty"`
	// Eddystone-URL url
	EddystoneUrl *string `json:"eddystone_url,omitempty"`
	ForSite      *bool   `json:"for_site,omitempty"`
	// Major number for iBeacon
	IbeaconMajor Optional[int] `json:"ibeacon_major"`
	// Minor number for iBeacon
	IbeaconMinor Optional[int]       `json:"ibeacon_minor"`
	IbeaconUuid  Optional[uuid.UUID] `json:"ibeacon_uuid"`
	// Unique ID of the object instance in the Mist Organization
	Id *uuid.UUID `json:"id,omitempty"`
	// Optional, MAC of the beacon, currently used only to identify battery voltage
	Mac *string `json:"mac,omitempty"`
	// Map where the device belongs to
	MapId *uuid.UUID `json:"map_id,omitempty"`
	// When the object has been modified for the last time, in epoch
	ModifiedTime *float64 `json:"modified_time,omitempty"`
	// Name / label of the device
	Name  *string    `json:"name,omitempty"`
	OrgId *uuid.UUID `json:"org_id,omitempty"`
	// In dBm
	Power  *int       `json:"power,omitempty"`
	SiteId *uuid.UUID `json:"site_id,omitempty"`
	// enum: `eddystone-uid`, `eddystone-url`, `ibeacon`
	Type *BeaconTypeEnum `json:"type,omitempty"`
	// X in pixel
	X *float64 `json:"x,omitempty"`
	// Y in pixel
	Y                    *float64               `json:"y,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for Beacon,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (b Beacon) String() string {
	return fmt.Sprintf(
		"Beacon[CreatedTime=%v, EddystoneInstance=%v, EddystoneNamespace=%v, EddystoneUrl=%v, ForSite=%v, IbeaconMajor=%v, IbeaconMinor=%v, IbeaconUuid=%v, Id=%v, Mac=%v, MapId=%v, ModifiedTime=%v, Name=%v, OrgId=%v, Power=%v, SiteId=%v, Type=%v, X=%v, Y=%v, AdditionalProperties=%v]",
		b.CreatedTime, b.EddystoneInstance, b.EddystoneNamespace, b.EddystoneUrl, b.ForSite, b.IbeaconMajor, b.IbeaconMinor, b.IbeaconUuid, b.Id, b.Mac, b.MapId, b.ModifiedTime, b.Name, b.OrgId, b.Power, b.SiteId, b.Type, b.X, b.Y, b.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for Beacon.
// It customizes the JSON marshaling process for Beacon objects.
func (b Beacon) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(b.AdditionalProperties,
		"created_time", "eddystone_instance", "eddystone_namespace", "eddystone_url", "for_site", "ibeacon_major", "ibeacon_minor", "ibeacon_uuid", "id", "mac", "map_id", "modified_time", "name", "org_id", "power", "site_id", "type", "x", "y"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(b.toMap())
}

// toMap converts the Beacon object to a map representation for JSON marshaling.
func (b Beacon) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, b.AdditionalProperties)
	if b.CreatedTime != nil {
		structMap["created_time"] = b.CreatedTime
	}
	if b.EddystoneInstance != nil {
		structMap["eddystone_instance"] = b.EddystoneInstance
	}
	if b.EddystoneNamespace != nil {
		structMap["eddystone_namespace"] = b.EddystoneNamespace
	}
	if b.EddystoneUrl != nil {
		structMap["eddystone_url"] = b.EddystoneUrl
	}
	if b.ForSite != nil {
		structMap["for_site"] = b.ForSite
	}
	if b.IbeaconMajor.IsValueSet() {
		if b.IbeaconMajor.Value() != nil {
			structMap["ibeacon_major"] = b.IbeaconMajor.Value()
		} else {
			structMap["ibeacon_major"] = nil
		}
	}
	if b.IbeaconMinor.IsValueSet() {
		if b.IbeaconMinor.Value() != nil {
			structMap["ibeacon_minor"] = b.IbeaconMinor.Value()
		} else {
			structMap["ibeacon_minor"] = nil
		}
	}
	if b.IbeaconUuid.IsValueSet() {
		if b.IbeaconUuid.Value() != nil {
			structMap["ibeacon_uuid"] = b.IbeaconUuid.Value()
		} else {
			structMap["ibeacon_uuid"] = nil
		}
	}
	if b.Id != nil {
		structMap["id"] = b.Id
	}
	if b.Mac != nil {
		structMap["mac"] = b.Mac
	}
	if b.MapId != nil {
		structMap["map_id"] = b.MapId
	}
	if b.ModifiedTime != nil {
		structMap["modified_time"] = b.ModifiedTime
	}
	if b.Name != nil {
		structMap["name"] = b.Name
	}
	if b.OrgId != nil {
		structMap["org_id"] = b.OrgId
	}
	if b.Power != nil {
		structMap["power"] = b.Power
	}
	if b.SiteId != nil {
		structMap["site_id"] = b.SiteId
	}
	if b.Type != nil {
		structMap["type"] = b.Type
	}
	if b.X != nil {
		structMap["x"] = b.X
	}
	if b.Y != nil {
		structMap["y"] = b.Y
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Beacon.
// It customizes the JSON unmarshaling process for Beacon objects.
func (b *Beacon) UnmarshalJSON(input []byte) error {
	var temp tempBeacon
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "created_time", "eddystone_instance", "eddystone_namespace", "eddystone_url", "for_site", "ibeacon_major", "ibeacon_minor", "ibeacon_uuid", "id", "mac", "map_id", "modified_time", "name", "org_id", "power", "site_id", "type", "x", "y")
	if err != nil {
		return err
	}
	b.AdditionalProperties = additionalProperties

	b.CreatedTime = temp.CreatedTime
	b.EddystoneInstance = temp.EddystoneInstance
	b.EddystoneNamespace = temp.EddystoneNamespace
	b.EddystoneUrl = temp.EddystoneUrl
	b.ForSite = temp.ForSite
	b.IbeaconMajor = temp.IbeaconMajor
	b.IbeaconMinor = temp.IbeaconMinor
	b.IbeaconUuid = temp.IbeaconUuid
	b.Id = temp.Id
	b.Mac = temp.Mac
	b.MapId = temp.MapId
	b.ModifiedTime = temp.ModifiedTime
	b.Name = temp.Name
	b.OrgId = temp.OrgId
	b.Power = temp.Power
	b.SiteId = temp.SiteId
	b.Type = temp.Type
	b.X = temp.X
	b.Y = temp.Y
	return nil
}

// tempBeacon is a temporary struct used for validating the fields of Beacon.
type tempBeacon struct {
	CreatedTime        *float64            `json:"created_time,omitempty"`
	EddystoneInstance  *string             `json:"eddystone_instance,omitempty"`
	EddystoneNamespace *string             `json:"eddystone_namespace,omitempty"`
	EddystoneUrl       *string             `json:"eddystone_url,omitempty"`
	ForSite            *bool               `json:"for_site,omitempty"`
	IbeaconMajor       Optional[int]       `json:"ibeacon_major"`
	IbeaconMinor       Optional[int]       `json:"ibeacon_minor"`
	IbeaconUuid        Optional[uuid.UUID] `json:"ibeacon_uuid"`
	Id                 *uuid.UUID          `json:"id,omitempty"`
	Mac                *string             `json:"mac,omitempty"`
	MapId              *uuid.UUID          `json:"map_id,omitempty"`
	ModifiedTime       *float64            `json:"modified_time,omitempty"`
	Name               *string             `json:"name,omitempty"`
	OrgId              *uuid.UUID          `json:"org_id,omitempty"`
	Power              *int                `json:"power,omitempty"`
	SiteId             *uuid.UUID          `json:"site_id,omitempty"`
	Type               *BeaconTypeEnum     `json:"type,omitempty"`
	X                  *float64            `json:"x,omitempty"`
	Y                  *float64            `json:"y,omitempty"`
}
