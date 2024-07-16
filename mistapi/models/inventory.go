package models

import (
    "encoding/json"
    "github.com/google/uuid"
)

// Inventory represents a Inventory struct.
type Inventory struct {
    // only if `type`==`switch` or `type`==`gateway`
    // whether the switch/gateway is adopted
    Adopted              *bool            `json:"adopted,omitempty"`
    // whether the device is connected
    Connected            *bool            `json:"connected,omitempty"`
    // inventory created time, in epoch
    CreatedTime          *float64         `json:"created_time,omitempty"`
    // deviceprofile id if assigned, null if not assigned
    DeviceprofileId      Optional[string] `json:"deviceprofile_id"`
    // hostname reported by the device
    Hostname             *string          `json:"hostname,omitempty"`
    // device hardware revision number
    HwRev                *string          `json:"hw_rev,omitempty"`
    // device id
    Id                   *string          `json:"id,omitempty"`
    Jsi                  *bool            `json:"jsi,omitempty"`
    // device MAC address
    Mac                  *string          `json:"mac,omitempty"`
    // device claim code
    Magic                *string          `json:"magic,omitempty"`
    // device model
    Model                *string          `json:"model,omitempty"`
    // inventory last modified time, in epoch
    ModifiedTime         *float64         `json:"modified_time,omitempty"`
    // device name if configured
    Name                 *string          `json:"name,omitempty"`
    OrgId                *uuid.UUID       `json:"org_id,omitempty"`
    // device serial
    Serial               *string          `json:"serial,omitempty"`
    SiteId               *uuid.UUID       `json:"site_id,omitempty"`
    // device stock keeping unit
    Sku                  *string          `json:"sku,omitempty"`
    Type                 *DeviceTypeEnum  `json:"type,omitempty"`
    // only if `type`==`switch`, MAC Address of the Virtual Chassis
    VcMac                *string          `json:"vc_mac,omitempty"`
    AdditionalProperties map[string]any   `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for Inventory.
// It customizes the JSON marshaling process for Inventory objects.
func (i Inventory) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(i.toMap())
}

// toMap converts the Inventory object to a map representation for JSON marshaling.
func (i Inventory) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, i.AdditionalProperties)
    if i.Adopted != nil {
        structMap["adopted"] = i.Adopted
    }
    if i.Connected != nil {
        structMap["connected"] = i.Connected
    }
    if i.CreatedTime != nil {
        structMap["created_time"] = i.CreatedTime
    }
    if i.DeviceprofileId.IsValueSet() {
        if i.DeviceprofileId.Value() != nil {
            structMap["deviceprofile_id"] = i.DeviceprofileId.Value()
        } else {
            structMap["deviceprofile_id"] = nil
        }
    }
    if i.Hostname != nil {
        structMap["hostname"] = i.Hostname
    }
    if i.HwRev != nil {
        structMap["hw_rev"] = i.HwRev
    }
    if i.Id != nil {
        structMap["id"] = i.Id
    }
    if i.Jsi != nil {
        structMap["jsi"] = i.Jsi
    }
    if i.Mac != nil {
        structMap["mac"] = i.Mac
    }
    if i.Magic != nil {
        structMap["magic"] = i.Magic
    }
    if i.Model != nil {
        structMap["model"] = i.Model
    }
    if i.ModifiedTime != nil {
        structMap["modified_time"] = i.ModifiedTime
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
    if i.Type != nil {
        structMap["type"] = i.Type
    }
    if i.VcMac != nil {
        structMap["vc_mac"] = i.VcMac
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Inventory.
// It customizes the JSON unmarshaling process for Inventory objects.
func (i *Inventory) UnmarshalJSON(input []byte) error {
    var temp inventory
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "adopted", "connected", "created_time", "deviceprofile_id", "hostname", "hw_rev", "id", "jsi", "mac", "magic", "model", "modified_time", "name", "org_id", "serial", "site_id", "sku", "type", "vc_mac")
    if err != nil {
    	return err
    }
    
    i.AdditionalProperties = additionalProperties
    i.Adopted = temp.Adopted
    i.Connected = temp.Connected
    i.CreatedTime = temp.CreatedTime
    i.DeviceprofileId = temp.DeviceprofileId
    i.Hostname = temp.Hostname
    i.HwRev = temp.HwRev
    i.Id = temp.Id
    i.Jsi = temp.Jsi
    i.Mac = temp.Mac
    i.Magic = temp.Magic
    i.Model = temp.Model
    i.ModifiedTime = temp.ModifiedTime
    i.Name = temp.Name
    i.OrgId = temp.OrgId
    i.Serial = temp.Serial
    i.SiteId = temp.SiteId
    i.Sku = temp.Sku
    i.Type = temp.Type
    i.VcMac = temp.VcMac
    return nil
}

// inventory is a temporary struct used for validating the fields of Inventory.
type inventory  struct {
    Adopted         *bool            `json:"adopted,omitempty"`
    Connected       *bool            `json:"connected,omitempty"`
    CreatedTime     *float64         `json:"created_time,omitempty"`
    DeviceprofileId Optional[string] `json:"deviceprofile_id"`
    Hostname        *string          `json:"hostname,omitempty"`
    HwRev           *string          `json:"hw_rev,omitempty"`
    Id              *string          `json:"id,omitempty"`
    Jsi             *bool            `json:"jsi,omitempty"`
    Mac             *string          `json:"mac,omitempty"`
    Magic           *string          `json:"magic,omitempty"`
    Model           *string          `json:"model,omitempty"`
    ModifiedTime    *float64         `json:"modified_time,omitempty"`
    Name            *string          `json:"name,omitempty"`
    OrgId           *uuid.UUID       `json:"org_id,omitempty"`
    Serial          *string          `json:"serial,omitempty"`
    SiteId          *uuid.UUID       `json:"site_id,omitempty"`
    Sku             *string          `json:"sku,omitempty"`
    Type            *DeviceTypeEnum  `json:"type,omitempty"`
    VcMac           *string          `json:"vc_mac,omitempty"`
}
