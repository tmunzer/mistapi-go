package models

import (
    "encoding/json"
    "fmt"
    "github.com/google/uuid"
)

// Inventory represents a Inventory struct.
type Inventory struct {
    // Only if `type`==`switch` or `type`==`gateway`, whether the switch/gateway is adopted
    Adopted              *bool                  `json:"adopted,omitempty"`
    // Whether the device is connected
    Connected            *bool                  `json:"connected,omitempty"`
    // When the object has been created, in epoch
    CreatedTime          *float64               `json:"created_time,omitempty"`
    // Deviceprofile id if assigned, null if not assigned
    DeviceprofileId      Optional[string]       `json:"deviceprofile_id"`
    // Hostname reported by the device
    Hostname             *string                `json:"hostname,omitempty"`
    // Device hardware revision number
    HwRev                *string                `json:"hw_rev,omitempty"`
    // Unique ID of the object instance in the Mist Organnization
    Id                   *uuid.UUID             `json:"id,omitempty"`
    Jsi                  *bool                  `json:"jsi,omitempty"`
    // Device MAC address
    Mac                  *string                `json:"mac,omitempty"`
    // Device claim code
    Magic                *string                `json:"magic,omitempty"`
    // Device model
    Model                *string                `json:"model,omitempty"`
    // When the object has been modified for the last time, in epoch
    ModifiedTime         *float64               `json:"modified_time,omitempty"`
    // Device name if configured
    Name                 *string                `json:"name,omitempty"`
    OrgId                *uuid.UUID             `json:"org_id,omitempty"`
    // Device serial
    Serial               *string                `json:"serial,omitempty"`
    SiteId               *uuid.UUID             `json:"site_id,omitempty"`
    // Device stock keeping unit
    Sku                  *string                `json:"sku,omitempty"`
    // enum: `ap`, `gateway`, `switch`
    Type                 *DeviceTypeEnum        `json:"type,omitempty"`
    // If `type`==`switch` and device part of a Virtual Chassis, MAC Address of the Virtual Chassis. if `type`==`gateway` and device part of a Clust, MAC Address of the Cluster
    VcMac                *string                `json:"vc_mac,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for Inventory,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (i Inventory) String() string {
    return fmt.Sprintf(
    	"Inventory[Adopted=%v, Connected=%v, CreatedTime=%v, DeviceprofileId=%v, Hostname=%v, HwRev=%v, Id=%v, Jsi=%v, Mac=%v, Magic=%v, Model=%v, ModifiedTime=%v, Name=%v, OrgId=%v, Serial=%v, SiteId=%v, Sku=%v, Type=%v, VcMac=%v, AdditionalProperties=%v]",
    	i.Adopted, i.Connected, i.CreatedTime, i.DeviceprofileId, i.Hostname, i.HwRev, i.Id, i.Jsi, i.Mac, i.Magic, i.Model, i.ModifiedTime, i.Name, i.OrgId, i.Serial, i.SiteId, i.Sku, i.Type, i.VcMac, i.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for Inventory.
// It customizes the JSON marshaling process for Inventory objects.
func (i Inventory) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(i.AdditionalProperties,
        "adopted", "connected", "created_time", "deviceprofile_id", "hostname", "hw_rev", "id", "jsi", "mac", "magic", "model", "modified_time", "name", "org_id", "serial", "site_id", "sku", "type", "vc_mac"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(i.toMap())
}

// toMap converts the Inventory object to a map representation for JSON marshaling.
func (i Inventory) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, i.AdditionalProperties)
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
    var temp tempInventory
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "adopted", "connected", "created_time", "deviceprofile_id", "hostname", "hw_rev", "id", "jsi", "mac", "magic", "model", "modified_time", "name", "org_id", "serial", "site_id", "sku", "type", "vc_mac")
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

// tempInventory is a temporary struct used for validating the fields of Inventory.
type tempInventory  struct {
    Adopted         *bool            `json:"adopted,omitempty"`
    Connected       *bool            `json:"connected,omitempty"`
    CreatedTime     *float64         `json:"created_time,omitempty"`
    DeviceprofileId Optional[string] `json:"deviceprofile_id"`
    Hostname        *string          `json:"hostname,omitempty"`
    HwRev           *string          `json:"hw_rev,omitempty"`
    Id              *uuid.UUID       `json:"id,omitempty"`
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
