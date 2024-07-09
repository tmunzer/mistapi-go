package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// NacTag represents a NacTag struct.
type NacTag struct {
    // can be set to true to allow the override by usermac result
    AllowUsermacOverride *bool            `json:"allow_usermac_override,omitempty"`
    CreatedTime          *float64         `json:"created_time,omitempty"`
    // if `type`==`egress_vlan_names`, list of egress vlans to return
    EgressVlanNames      []string         `json:"egress_vlan_names,omitempty"`
    // if `type`==`gbp_tag`
    GbpTag               *int             `json:"gbp_tag,omitempty"`
    Id                   *uuid.UUID       `json:"id,omitempty"`
    // if `type`==`match`
    Match                *NacTagMatchEnum `json:"match,omitempty"`
    // This field is applicable only when `type`==`match`
    // * `false`: means it is sufficient to match any of the values (i.e., match-any behavior)
    // * `true`: means all values should be matched (i.e., match-all behavior)
    // Currently it makes sense to set this field to `true` only if the `match`==`idp_role` or `match`==`usermac_label`
    MatchAll             *bool            `json:"match_all,omitempty"`
    ModifiedTime         *float64         `json:"modified_time,omitempty"`
    Name                 string           `json:"name"`
    OrgId                *uuid.UUID       `json:"org_id,omitempty"`
    // if `type`==`radius_attrs`, user can specify a list of one or more standard attributes in the field "radius_attrs". 
    // It is the responsibility of the user to provide a syntactically correct string, otherwise it may not work as expected.
    // Note that it is allowed to have more than one radius_attrs in the result of a given rule.
    RadiusAttrs          []string         `json:"radius_attrs,omitempty"`
    // if `type`==`radius_group`
    RadiusGroup          *string          `json:"radius_group,omitempty"`
    // if `type`==`radius_vendor_attrs`, user can specify a list of one or more vendor-specific attributes in the field "radius_vendor_attrs". 
    // It is the responsibility of the user to provide a syntactically correct string, otherwise it may not work as expected.
    // Note that it is allowed to have more than one radius_vendor_attrs in the result of a given rule.
    RadiusVendorAttrs    []string         `json:"radius_vendor_attrs,omitempty"`
    // if `type`==`session_timeout, in seconds
    SessionTimeout       *int             `json:"session_timeout,omitempty"`
    Type                 NacTagTypeEnum   `json:"type"`
    // if `type`==`match`
    Values               []string         `json:"values,omitempty"`
    // if `type`==`vlan`
    Vlan                 *string          `json:"vlan,omitempty"`
    AdditionalProperties map[string]any   `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for NacTag.
// It customizes the JSON marshaling process for NacTag objects.
func (n NacTag) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(n.toMap())
}

// toMap converts the NacTag object to a map representation for JSON marshaling.
func (n NacTag) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, n.AdditionalProperties)
    if n.AllowUsermacOverride != nil {
        structMap["allow_usermac_override"] = n.AllowUsermacOverride
    }
    if n.CreatedTime != nil {
        structMap["created_time"] = n.CreatedTime
    }
    if n.EgressVlanNames != nil {
        structMap["egress_vlan_names"] = n.EgressVlanNames
    }
    if n.GbpTag != nil {
        structMap["gbp_tag"] = n.GbpTag
    }
    if n.Id != nil {
        structMap["id"] = n.Id
    }
    if n.Match != nil {
        structMap["match"] = n.Match
    }
    if n.MatchAll != nil {
        structMap["match_all"] = n.MatchAll
    }
    if n.ModifiedTime != nil {
        structMap["modified_time"] = n.ModifiedTime
    }
    structMap["name"] = n.Name
    if n.OrgId != nil {
        structMap["org_id"] = n.OrgId
    }
    if n.RadiusAttrs != nil {
        structMap["radius_attrs"] = n.RadiusAttrs
    }
    if n.RadiusGroup != nil {
        structMap["radius_group"] = n.RadiusGroup
    }
    if n.RadiusVendorAttrs != nil {
        structMap["radius_vendor_attrs"] = n.RadiusVendorAttrs
    }
    if n.SessionTimeout != nil {
        structMap["session_timeout"] = n.SessionTimeout
    }
    structMap["type"] = n.Type
    if n.Values != nil {
        structMap["values"] = n.Values
    }
    if n.Vlan != nil {
        structMap["vlan"] = n.Vlan
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for NacTag.
// It customizes the JSON unmarshaling process for NacTag objects.
func (n *NacTag) UnmarshalJSON(input []byte) error {
    var temp nacTag
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "allow_usermac_override", "created_time", "egress_vlan_names", "gbp_tag", "id", "match", "match_all", "modified_time", "name", "org_id", "radius_attrs", "radius_group", "radius_vendor_attrs", "session_timeout", "type", "values", "vlan")
    if err != nil {
    	return err
    }
    
    n.AdditionalProperties = additionalProperties
    n.AllowUsermacOverride = temp.AllowUsermacOverride
    n.CreatedTime = temp.CreatedTime
    n.EgressVlanNames = temp.EgressVlanNames
    n.GbpTag = temp.GbpTag
    n.Id = temp.Id
    n.Match = temp.Match
    n.MatchAll = temp.MatchAll
    n.ModifiedTime = temp.ModifiedTime
    n.Name = *temp.Name
    n.OrgId = temp.OrgId
    n.RadiusAttrs = temp.RadiusAttrs
    n.RadiusGroup = temp.RadiusGroup
    n.RadiusVendorAttrs = temp.RadiusVendorAttrs
    n.SessionTimeout = temp.SessionTimeout
    n.Type = *temp.Type
    n.Values = temp.Values
    n.Vlan = temp.Vlan
    return nil
}

// nacTag is a temporary struct used for validating the fields of NacTag.
type nacTag  struct {
    AllowUsermacOverride *bool            `json:"allow_usermac_override,omitempty"`
    CreatedTime          *float64         `json:"created_time,omitempty"`
    EgressVlanNames      []string         `json:"egress_vlan_names,omitempty"`
    GbpTag               *int             `json:"gbp_tag,omitempty"`
    Id                   *uuid.UUID       `json:"id,omitempty"`
    Match                *NacTagMatchEnum `json:"match,omitempty"`
    MatchAll             *bool            `json:"match_all,omitempty"`
    ModifiedTime         *float64         `json:"modified_time,omitempty"`
    Name                 *string          `json:"name"`
    OrgId                *uuid.UUID       `json:"org_id,omitempty"`
    RadiusAttrs          []string         `json:"radius_attrs,omitempty"`
    RadiusGroup          *string          `json:"radius_group,omitempty"`
    RadiusVendorAttrs    []string         `json:"radius_vendor_attrs,omitempty"`
    SessionTimeout       *int             `json:"session_timeout,omitempty"`
    Type                 *NacTagTypeEnum  `json:"type"`
    Values               []string         `json:"values,omitempty"`
    Vlan                 *string          `json:"vlan,omitempty"`
}

func (n *nacTag) validate() error {
    var errs []string
    if n.Name == nil {
        errs = append(errs, "required field `name` is missing for type `Nac_Tag`")
    }
    if n.Type == nil {
        errs = append(errs, "required field `type` is missing for type `Nac_Tag`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
