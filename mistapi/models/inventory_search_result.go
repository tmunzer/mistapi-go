package models

import (
    "encoding/json"
    "github.com/google/uuid"
)

// InventorySearchResult represents a InventorySearchResult struct.
type InventorySearchResult struct {
    Mac                  *string                       `json:"mac,omitempty"`
    Master               *bool                         `json:"master,omitempty"`
    Members              []InventorySearchResultMember `json:"members,omitempty"`
    Model                *string                       `json:"model,omitempty"`
    Name                 *string                       `json:"name,omitempty"`
    OrgId                *uuid.UUID                    `json:"org_id,omitempty"`
    Serial               *string                       `json:"serial,omitempty"`
    SiteId               *uuid.UUID                    `json:"site_id,omitempty"`
    Sku                  *string                       `json:"sku,omitempty"`
    Status               *string                       `json:"status,omitempty"`
    // enum: `ap`, `gateway`, `switch`
    Type                 *DeviceTypeEnum               `json:"type,omitempty"`
    VcMac                *string                       `json:"vc_mac,omitempty"`
    Version              *string                       `json:"version,omitempty"`
    AdditionalProperties map[string]any                `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for InventorySearchResult.
// It customizes the JSON marshaling process for InventorySearchResult objects.
func (i InventorySearchResult) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(i.toMap())
}

// toMap converts the InventorySearchResult object to a map representation for JSON marshaling.
func (i InventorySearchResult) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, i.AdditionalProperties)
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
    additionalProperties, err := UnmarshalAdditionalProperties(input, "mac", "master", "members", "model", "name", "org_id", "serial", "site_id", "sku", "status", "type", "vc_mac", "version")
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
type tempInventorySearchResult  struct {
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
    Type    *DeviceTypeEnum               `json:"type,omitempty"`
    VcMac   *string                       `json:"vc_mac,omitempty"`
    Version *string                       `json:"version,omitempty"`
}
