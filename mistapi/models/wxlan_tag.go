package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// WxlanTag represents a WxlanTag struct.
// WxLAN Tag
// * type:
// ** client: created manually (e.g. on wireless client table, when they spot a device of interest, they can create a wxlan tag for it
// ** resource: created automatically when we discover a network resource
// ** subnet: create automatically when a subnet is discovered
// * match:
// ** wlan_id, ap_id: values are a list of Wlan / Device ids
// ** client_mac: values are a list of MAC addresses
// * radius_group: this is a smart tag that matches RADIUS-Filter-ID, Airespace-ACL-Name (VendorID=14179, VendorType=6) / Aruba-User-Role (VendorID=14823, VendorType=1)
// * radius_username: this matches the ATTR-User-Name(1)
// * radius_class: thie matches the ATTR-Class(25)
// * radius_attr: the values are [ “6=1”, “26=10.2.3.4” ], this support other RADIUS attributes where we know the type
// * radius_vendor: the values are [ “14179.10=1”, “14178.16=1.2.3.4” ], this matches vendor attributes and will be dynamically evaluated
type WxlanTag struct {
    CreatedTime          *float64               `json:"created_time,omitempty"`
    ForSite              *bool                  `json:"for_site,omitempty"`
    Id                   *uuid.UUID             `json:"id,omitempty"`
    LastIps              []string               `json:"last_ips,omitempty"`
    Mac                  Optional[string]       `json:"mac"`
    Match                *WxlanTagMatchEnum     `json:"match,omitempty"`
    ModifiedTime         *float64               `json:"modified_time,omitempty"`
    // The name
    Name                 string                 `json:"name"`
    Op                   *WxlanTagOperationEnum `json:"op,omitempty"`
    OrgId                *uuid.UUID             `json:"org_id,omitempty"`
    ResourceMac          Optional[string]       `json:"resource_mac"`
    Services             []string               `json:"services,omitempty"`
    SiteId               *uuid.UUID             `json:"site_id,omitempty"`
    // if `type`==`specs`
    Specs                []WxlanTagSpec         `json:"specs,omitempty"`
    Subnet               *string                `json:"subnet,omitempty"`
    Type                 WxlanTagTypeEnum       `json:"type"`
    // if `type`!=`vlan_id` and `type`!=`specs`, list of values to match
    Values               []string               `json:"values,omitempty"`
    // if `type`==`vlan_id`
    VlanId               *WxlanTagVlanId        `json:"vlan_id,omitempty"`
    AdditionalProperties map[string]any         `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for WxlanTag.
// It customizes the JSON marshaling process for WxlanTag objects.
func (w WxlanTag) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(w.toMap())
}

// toMap converts the WxlanTag object to a map representation for JSON marshaling.
func (w WxlanTag) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, w.AdditionalProperties)
    if w.CreatedTime != nil {
        structMap["created_time"] = w.CreatedTime
    }
    if w.ForSite != nil {
        structMap["for_site"] = w.ForSite
    }
    if w.Id != nil {
        structMap["id"] = w.Id
    }
    if w.LastIps != nil {
        structMap["last_ips"] = w.LastIps
    }
    if w.Mac.IsValueSet() {
        if w.Mac.Value() != nil {
            structMap["mac"] = w.Mac.Value()
        } else {
            structMap["mac"] = nil
        }
    }
    if w.Match != nil {
        structMap["match"] = w.Match
    }
    if w.ModifiedTime != nil {
        structMap["modified_time"] = w.ModifiedTime
    }
    structMap["name"] = w.Name
    if w.Op != nil {
        structMap["op"] = w.Op
    }
    if w.OrgId != nil {
        structMap["org_id"] = w.OrgId
    }
    if w.ResourceMac.IsValueSet() {
        if w.ResourceMac.Value() != nil {
            structMap["resource_mac"] = w.ResourceMac.Value()
        } else {
            structMap["resource_mac"] = nil
        }
    }
    if w.Services != nil {
        structMap["services"] = w.Services
    }
    if w.SiteId != nil {
        structMap["site_id"] = w.SiteId
    }
    if w.Specs != nil {
        structMap["specs"] = w.Specs
    }
    if w.Subnet != nil {
        structMap["subnet"] = w.Subnet
    }
    structMap["type"] = w.Type
    if w.Values != nil {
        structMap["values"] = w.Values
    }
    if w.VlanId != nil {
        structMap["vlan_id"] = w.VlanId.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WxlanTag.
// It customizes the JSON unmarshaling process for WxlanTag objects.
func (w *WxlanTag) UnmarshalJSON(input []byte) error {
    var temp wxlanTag
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "created_time", "for_site", "id", "last_ips", "mac", "match", "modified_time", "name", "op", "org_id", "resource_mac", "services", "site_id", "specs", "subnet", "type", "values", "vlan_id")
    if err != nil {
    	return err
    }
    
    w.AdditionalProperties = additionalProperties
    w.CreatedTime = temp.CreatedTime
    w.ForSite = temp.ForSite
    w.Id = temp.Id
    w.LastIps = temp.LastIps
    w.Mac = temp.Mac
    w.Match = temp.Match
    w.ModifiedTime = temp.ModifiedTime
    w.Name = *temp.Name
    w.Op = temp.Op
    w.OrgId = temp.OrgId
    w.ResourceMac = temp.ResourceMac
    w.Services = temp.Services
    w.SiteId = temp.SiteId
    w.Specs = temp.Specs
    w.Subnet = temp.Subnet
    w.Type = *temp.Type
    w.Values = temp.Values
    w.VlanId = temp.VlanId
    return nil
}

// wxlanTag is a temporary struct used for validating the fields of WxlanTag.
type wxlanTag  struct {
    CreatedTime  *float64               `json:"created_time,omitempty"`
    ForSite      *bool                  `json:"for_site,omitempty"`
    Id           *uuid.UUID             `json:"id,omitempty"`
    LastIps      []string               `json:"last_ips,omitempty"`
    Mac          Optional[string]       `json:"mac"`
    Match        *WxlanTagMatchEnum     `json:"match,omitempty"`
    ModifiedTime *float64               `json:"modified_time,omitempty"`
    Name         *string                `json:"name"`
    Op           *WxlanTagOperationEnum `json:"op,omitempty"`
    OrgId        *uuid.UUID             `json:"org_id,omitempty"`
    ResourceMac  Optional[string]       `json:"resource_mac"`
    Services     []string               `json:"services,omitempty"`
    SiteId       *uuid.UUID             `json:"site_id,omitempty"`
    Specs        []WxlanTagSpec         `json:"specs,omitempty"`
    Subnet       *string                `json:"subnet,omitempty"`
    Type         *WxlanTagTypeEnum      `json:"type"`
    Values       []string               `json:"values,omitempty"`
    VlanId       *WxlanTagVlanId        `json:"vlan_id,omitempty"`
}

func (w *wxlanTag) validate() error {
    var errs []string
    if w.Name == nil {
        errs = append(errs, "required field `name` is missing for type `Wxlan_Tag`")
    }
    if w.Type == nil {
        errs = append(errs, "required field `type` is missing for type `Wxlan_Tag`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
