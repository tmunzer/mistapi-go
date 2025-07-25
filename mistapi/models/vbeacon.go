// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
    "github.com/google/uuid"
)

// Vbeacon represents a Vbeacon struct.
// vBeacon
type Vbeacon struct {
    // When the object has been created, in epoch
    CreatedTime          *float64                `json:"created_time,omitempty"`
    ForSite              *bool                   `json:"for_site,omitempty"`
    // Unique ID of the object instance in the Mist Organization
    Id                   *uuid.UUID              `json:"id,omitempty"`
    // Bluetooth tag major
    Major                *int                    `json:"major,omitempty"`
    // Map where the device belongs to
    MapId                *uuid.UUID              `json:"map_id,omitempty"`
    // Message that can be displayed when the sdkclient gets near the vbeacon
    Message              *string                 `json:"message,omitempty"`
    // Bluetooth tag minor
    Minor                *int                    `json:"minor,omitempty"`
    // When the object has been modified for the last time, in epoch
    ModifiedTime         *float64                `json:"modified_time,omitempty"`
    // Name / label of the device
    Name                 *string                 `json:"name,omitempty"`
    OrgId                *uuid.UUID              `json:"org_id,omitempty"`
    // Required if `power_mode`==`custom`, -30 - 100, in dBm. For default power_mode, power = 4 dBm.
    Power                *int                    `json:"power,omitempty"`
    // enum: `custom`, `default`
    PowerMode            *BleConfigPowerModeEnum `json:"power_mode,omitempty"`
    SiteId               *uuid.UUID              `json:"site_id,omitempty"`
    // URL to show, optional
    Url                  *string                 `json:"url,omitempty"`
    // Bluetooth tag UUID
    Uuid                 *uuid.UUID              `json:"uuid,omitempty"`
    // Name to be used in wayfinding_path or wayfinding_grid blob
    WayfindingNodename   *string                 `json:"wayfinding_nodename,omitempty"`
    // X in pixel
    X                    *float64                `json:"x,omitempty"`
    // Y in pixel
    Y                    *float64                `json:"y,omitempty"`
    AdditionalProperties map[string]interface{}  `json:"_"`
}

// String implements the fmt.Stringer interface for Vbeacon,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (v Vbeacon) String() string {
    return fmt.Sprintf(
    	"Vbeacon[CreatedTime=%v, ForSite=%v, Id=%v, Major=%v, MapId=%v, Message=%v, Minor=%v, ModifiedTime=%v, Name=%v, OrgId=%v, Power=%v, PowerMode=%v, SiteId=%v, Url=%v, Uuid=%v, WayfindingNodename=%v, X=%v, Y=%v, AdditionalProperties=%v]",
    	v.CreatedTime, v.ForSite, v.Id, v.Major, v.MapId, v.Message, v.Minor, v.ModifiedTime, v.Name, v.OrgId, v.Power, v.PowerMode, v.SiteId, v.Url, v.Uuid, v.WayfindingNodename, v.X, v.Y, v.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for Vbeacon.
// It customizes the JSON marshaling process for Vbeacon objects.
func (v Vbeacon) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(v.AdditionalProperties,
        "created_time", "for_site", "id", "major", "map_id", "message", "minor", "modified_time", "name", "org_id", "power", "power_mode", "site_id", "url", "uuid", "wayfinding_nodename", "x", "y"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(v.toMap())
}

// toMap converts the Vbeacon object to a map representation for JSON marshaling.
func (v Vbeacon) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, v.AdditionalProperties)
    if v.CreatedTime != nil {
        structMap["created_time"] = v.CreatedTime
    }
    if v.ForSite != nil {
        structMap["for_site"] = v.ForSite
    }
    if v.Id != nil {
        structMap["id"] = v.Id
    }
    if v.Major != nil {
        structMap["major"] = v.Major
    }
    if v.MapId != nil {
        structMap["map_id"] = v.MapId
    }
    if v.Message != nil {
        structMap["message"] = v.Message
    }
    if v.Minor != nil {
        structMap["minor"] = v.Minor
    }
    if v.ModifiedTime != nil {
        structMap["modified_time"] = v.ModifiedTime
    }
    if v.Name != nil {
        structMap["name"] = v.Name
    }
    if v.OrgId != nil {
        structMap["org_id"] = v.OrgId
    }
    if v.Power != nil {
        structMap["power"] = v.Power
    }
    if v.PowerMode != nil {
        structMap["power_mode"] = v.PowerMode
    }
    if v.SiteId != nil {
        structMap["site_id"] = v.SiteId
    }
    if v.Url != nil {
        structMap["url"] = v.Url
    }
    if v.Uuid != nil {
        structMap["uuid"] = v.Uuid
    }
    if v.WayfindingNodename != nil {
        structMap["wayfinding_nodename"] = v.WayfindingNodename
    }
    if v.X != nil {
        structMap["x"] = v.X
    }
    if v.Y != nil {
        structMap["y"] = v.Y
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Vbeacon.
// It customizes the JSON unmarshaling process for Vbeacon objects.
func (v *Vbeacon) UnmarshalJSON(input []byte) error {
    var temp tempVbeacon
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "created_time", "for_site", "id", "major", "map_id", "message", "minor", "modified_time", "name", "org_id", "power", "power_mode", "site_id", "url", "uuid", "wayfinding_nodename", "x", "y")
    if err != nil {
    	return err
    }
    v.AdditionalProperties = additionalProperties
    
    v.CreatedTime = temp.CreatedTime
    v.ForSite = temp.ForSite
    v.Id = temp.Id
    v.Major = temp.Major
    v.MapId = temp.MapId
    v.Message = temp.Message
    v.Minor = temp.Minor
    v.ModifiedTime = temp.ModifiedTime
    v.Name = temp.Name
    v.OrgId = temp.OrgId
    v.Power = temp.Power
    v.PowerMode = temp.PowerMode
    v.SiteId = temp.SiteId
    v.Url = temp.Url
    v.Uuid = temp.Uuid
    v.WayfindingNodename = temp.WayfindingNodename
    v.X = temp.X
    v.Y = temp.Y
    return nil
}

// tempVbeacon is a temporary struct used for validating the fields of Vbeacon.
type tempVbeacon  struct {
    CreatedTime        *float64                `json:"created_time,omitempty"`
    ForSite            *bool                   `json:"for_site,omitempty"`
    Id                 *uuid.UUID              `json:"id,omitempty"`
    Major              *int                    `json:"major,omitempty"`
    MapId              *uuid.UUID              `json:"map_id,omitempty"`
    Message            *string                 `json:"message,omitempty"`
    Minor              *int                    `json:"minor,omitempty"`
    ModifiedTime       *float64                `json:"modified_time,omitempty"`
    Name               *string                 `json:"name,omitempty"`
    OrgId              *uuid.UUID              `json:"org_id,omitempty"`
    Power              *int                    `json:"power,omitempty"`
    PowerMode          *BleConfigPowerModeEnum `json:"power_mode,omitempty"`
    SiteId             *uuid.UUID              `json:"site_id,omitempty"`
    Url                *string                 `json:"url,omitempty"`
    Uuid               *uuid.UUID              `json:"uuid,omitempty"`
    WayfindingNodename *string                 `json:"wayfinding_nodename,omitempty"`
    X                  *float64                `json:"x,omitempty"`
    Y                  *float64                `json:"y,omitempty"`
}
