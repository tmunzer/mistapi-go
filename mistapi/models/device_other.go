package models

import (
    "encoding/json"
    "fmt"
    "github.com/google/uuid"
)

// DeviceOther represents a DeviceOther struct.
type DeviceOther struct {
    // When the object has been created, in epoch
    CreatedTime          *float64               `json:"created_time,omitempty"`
    DeviceMac            *string                `json:"device_mac,omitempty"`
    // Unique ID of the object instance in the Mist Organization
    Id                   *uuid.UUID             `json:"id,omitempty"`
    Mac                  *string                `json:"mac,omitempty"`
    Model                *string                `json:"model,omitempty"`
    // When the object has been modified for the last time, in epoch
    ModifiedTime         *float64               `json:"modified_time,omitempty"`
    Name                 *string                `json:"name,omitempty"`
    OrgId                *uuid.UUID             `json:"org_id,omitempty"`
    Serial               *string                `json:"serial,omitempty"`
    SiteId               *uuid.UUID             `json:"site_id,omitempty"`
    State                *string                `json:"state,omitempty"`
    Vendor               *string                `json:"vendor,omitempty"`
    VendorApiId          *string                `json:"vendor_api_id,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for DeviceOther,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (d DeviceOther) String() string {
    return fmt.Sprintf(
    	"DeviceOther[CreatedTime=%v, DeviceMac=%v, Id=%v, Mac=%v, Model=%v, ModifiedTime=%v, Name=%v, OrgId=%v, Serial=%v, SiteId=%v, State=%v, Vendor=%v, VendorApiId=%v, AdditionalProperties=%v]",
    	d.CreatedTime, d.DeviceMac, d.Id, d.Mac, d.Model, d.ModifiedTime, d.Name, d.OrgId, d.Serial, d.SiteId, d.State, d.Vendor, d.VendorApiId, d.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for DeviceOther.
// It customizes the JSON marshaling process for DeviceOther objects.
func (d DeviceOther) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(d.AdditionalProperties,
        "created_time", "device_mac", "id", "mac", "model", "modified_time", "name", "org_id", "serial", "site_id", "state", "vendor", "vendor_api_id"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(d.toMap())
}

// toMap converts the DeviceOther object to a map representation for JSON marshaling.
func (d DeviceOther) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, d.AdditionalProperties)
    if d.CreatedTime != nil {
        structMap["created_time"] = d.CreatedTime
    }
    if d.DeviceMac != nil {
        structMap["device_mac"] = d.DeviceMac
    }
    if d.Id != nil {
        structMap["id"] = d.Id
    }
    if d.Mac != nil {
        structMap["mac"] = d.Mac
    }
    if d.Model != nil {
        structMap["model"] = d.Model
    }
    if d.ModifiedTime != nil {
        structMap["modified_time"] = d.ModifiedTime
    }
    if d.Name != nil {
        structMap["name"] = d.Name
    }
    if d.OrgId != nil {
        structMap["org_id"] = d.OrgId
    }
    if d.Serial != nil {
        structMap["serial"] = d.Serial
    }
    if d.SiteId != nil {
        structMap["site_id"] = d.SiteId
    }
    if d.State != nil {
        structMap["state"] = d.State
    }
    if d.Vendor != nil {
        structMap["vendor"] = d.Vendor
    }
    if d.VendorApiId != nil {
        structMap["vendor_api_id"] = d.VendorApiId
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for DeviceOther.
// It customizes the JSON unmarshaling process for DeviceOther objects.
func (d *DeviceOther) UnmarshalJSON(input []byte) error {
    var temp tempDeviceOther
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "created_time", "device_mac", "id", "mac", "model", "modified_time", "name", "org_id", "serial", "site_id", "state", "vendor", "vendor_api_id")
    if err != nil {
    	return err
    }
    d.AdditionalProperties = additionalProperties
    
    d.CreatedTime = temp.CreatedTime
    d.DeviceMac = temp.DeviceMac
    d.Id = temp.Id
    d.Mac = temp.Mac
    d.Model = temp.Model
    d.ModifiedTime = temp.ModifiedTime
    d.Name = temp.Name
    d.OrgId = temp.OrgId
    d.Serial = temp.Serial
    d.SiteId = temp.SiteId
    d.State = temp.State
    d.Vendor = temp.Vendor
    d.VendorApiId = temp.VendorApiId
    return nil
}

// tempDeviceOther is a temporary struct used for validating the fields of DeviceOther.
type tempDeviceOther  struct {
    CreatedTime  *float64   `json:"created_time,omitempty"`
    DeviceMac    *string    `json:"device_mac,omitempty"`
    Id           *uuid.UUID `json:"id,omitempty"`
    Mac          *string    `json:"mac,omitempty"`
    Model        *string    `json:"model,omitempty"`
    ModifiedTime *float64   `json:"modified_time,omitempty"`
    Name         *string    `json:"name,omitempty"`
    OrgId        *uuid.UUID `json:"org_id,omitempty"`
    Serial       *string    `json:"serial,omitempty"`
    SiteId       *uuid.UUID `json:"site_id,omitempty"`
    State        *string    `json:"state,omitempty"`
    Vendor       *string    `json:"vendor,omitempty"`
    VendorApiId  *string    `json:"vendor_api_id,omitempty"`
}
