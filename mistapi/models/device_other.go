package models

import (
    "encoding/json"
    "github.com/google/uuid"
)

// DeviceOther represents a DeviceOther struct.
type DeviceOther struct {
    CreatedTime          *float64       `json:"created_time,omitempty"`
    DeviceMac            *string        `json:"device_mac,omitempty"`
    Id                   *string        `json:"id,omitempty"`
    Mac                  *string        `json:"mac,omitempty"`
    Model                *string        `json:"model,omitempty"`
    ModifiedTime         *float64       `json:"modified_time,omitempty"`
    Name                 *string        `json:"name,omitempty"`
    OrgId                *uuid.UUID     `json:"org_id,omitempty"`
    Serial               *string        `json:"serial,omitempty"`
    SiteId               *uuid.UUID     `json:"site_id,omitempty"`
    Vendor               *string        `json:"vendor,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for DeviceOther.
// It customizes the JSON marshaling process for DeviceOther objects.
func (d DeviceOther) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(d.toMap())
}

// toMap converts the DeviceOther object to a map representation for JSON marshaling.
func (d DeviceOther) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, d.AdditionalProperties)
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
    if d.Vendor != nil {
        structMap["vendor"] = d.Vendor
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for DeviceOther.
// It customizes the JSON unmarshaling process for DeviceOther objects.
func (d *DeviceOther) UnmarshalJSON(input []byte) error {
    var temp deviceOther
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "created_time", "device_mac", "id", "mac", "model", "modified_time", "name", "org_id", "serial", "site_id", "vendor")
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
    d.Vendor = temp.Vendor
    return nil
}

// deviceOther is a temporary struct used for validating the fields of DeviceOther.
type deviceOther  struct {
    CreatedTime  *float64   `json:"created_time,omitempty"`
    DeviceMac    *string    `json:"device_mac,omitempty"`
    Id           *string    `json:"id,omitempty"`
    Mac          *string    `json:"mac,omitempty"`
    Model        *string    `json:"model,omitempty"`
    ModifiedTime *float64   `json:"modified_time,omitempty"`
    Name         *string    `json:"name,omitempty"`
    OrgId        *uuid.UUID `json:"org_id,omitempty"`
    Serial       *string    `json:"serial,omitempty"`
    SiteId       *uuid.UUID `json:"site_id,omitempty"`
    Vendor       *string    `json:"vendor,omitempty"`
}
