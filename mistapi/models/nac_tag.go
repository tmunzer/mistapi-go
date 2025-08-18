// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/google/uuid"
    "strings"
)

// NacTag represents a NacTag struct.
type NacTag struct {
    // Can be set to true to allow the override by usermac result
    AllowUsermacOverride *bool                   `json:"allow_usermac_override,omitempty"`
    // When the object has been created, in epoch
    CreatedTime          *float64                `json:"created_time,omitempty"`
    // If `type`==`egress_vlan_names`, list of egress vlans to return
    EgressVlanNames      []string                `json:"egress_vlan_names,omitempty"`
    // If `type`==`gbp_tag`
    GbpTag               *NacTagGbpTag           `json:"gbp_tag,omitempty"`
    // Unique ID of the object instance in the Mist Organization
    Id                   *uuid.UUID              `json:"id,omitempty"`
    // if `type`==`match`. enum: `cert_cn`, `cert_issuer`, `cert_san`, `cert_serial`, `cert_sub`, `cert_template`, `client_mac`, `hostname`, `idp_role`, `ingress_vlan`, `mdm_status`, `nas_ip`, `radius_group`, `realm`, `ssid`, `user_name`, `usermac_label`
    Match                *NacTagMatchEnum        `json:"match,omitempty"`
    // This field is applicable only when `type`==`match`
    // * `false`: means it is sufficient to match any of the values (i.e., match-any behavior)
    // * `true`: means all values should be matched (i.e., match-all behavior)
    // Currently it makes sense to set this field to `true` only if the `match`==`idp_role` or `match`==`usermac_label`
    MatchAll             *bool                   `json:"match_all,omitempty"`
    // When the object has been modified for the last time, in epoch
    ModifiedTime         *float64                `json:"modified_time,omitempty"`
    // If `type`==`redirect_guest_portal`, the ID of the guest portal to redirect to
    NacportalId          *uuid.UUID              `json:"nacportal_id,omitempty"`
    Name                 string                  `json:"name"`
    OrgId                *uuid.UUID              `json:"org_id,omitempty"`
    // If `type`==`radius_attrs`, user can specify a list of one or more standard attributes in the field "radius_attrs".
    // It is the responsibility of the user to provide a syntactically correct string, otherwise it may not work as expected.
    // Note that it is allowed to have more than one radius_attrs in the result of a given rule.
    RadiusAttrs          []string                `json:"radius_attrs,omitempty"`
    // If `type`==`radius_group`
    RadiusGroup          *string                 `json:"radius_group,omitempty"`
    // If `type`==`radius_vendor_attrs`, user can specify a list of one or more vendor-specific attributes in the field "radius_vendor_attrs".
    // It is the responsibility of the user to provide a syntactically correct string, otherwise it may not work as expected.
    // Note that it is allowed to have more than one radius_vendor_attrs in the result of a given rule.
    RadiusVendorAttrs    []string                `json:"radius_vendor_attrs,omitempty"`
    // If `type`==`session_timeout, in seconds
    SessionTimeout       *int                    `json:"session_timeout,omitempty"`
    // enum: `egress_vlan_names`, `gbp_tag`, `match`, `radius_attrs`, `radius_group`, `radius_vendor_attrs`, `redirect_guest_portal`, `session_timeout`, `username_attr`, `vlan`
    Type                 NacTagTypeEnum          `json:"type"`
    // enum: `automatic`, `cn`, `dns`, `email`, `upn`
    UsernameAttr         *NacTagUsernameAttrEnum `json:"username_attr,omitempty"`
    // If `type`==`match`
    Values               []string                `json:"values,omitempty"`
    // If `type`==`vlan`
    Vlan                 *string                 `json:"vlan,omitempty"`
    AdditionalProperties map[string]interface{}  `json:"_"`
}

// String implements the fmt.Stringer interface for NacTag,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (n NacTag) String() string {
    return fmt.Sprintf(
    	"NacTag[AllowUsermacOverride=%v, CreatedTime=%v, EgressVlanNames=%v, GbpTag=%v, Id=%v, Match=%v, MatchAll=%v, ModifiedTime=%v, NacportalId=%v, Name=%v, OrgId=%v, RadiusAttrs=%v, RadiusGroup=%v, RadiusVendorAttrs=%v, SessionTimeout=%v, Type=%v, UsernameAttr=%v, Values=%v, Vlan=%v, AdditionalProperties=%v]",
    	n.AllowUsermacOverride, n.CreatedTime, n.EgressVlanNames, n.GbpTag, n.Id, n.Match, n.MatchAll, n.ModifiedTime, n.NacportalId, n.Name, n.OrgId, n.RadiusAttrs, n.RadiusGroup, n.RadiusVendorAttrs, n.SessionTimeout, n.Type, n.UsernameAttr, n.Values, n.Vlan, n.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for NacTag.
// It customizes the JSON marshaling process for NacTag objects.
func (n NacTag) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(n.AdditionalProperties,
        "allow_usermac_override", "created_time", "egress_vlan_names", "gbp_tag", "id", "match", "match_all", "modified_time", "nacportal_id", "name", "org_id", "radius_attrs", "radius_group", "radius_vendor_attrs", "session_timeout", "type", "username_attr", "values", "vlan"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(n.toMap())
}

// toMap converts the NacTag object to a map representation for JSON marshaling.
func (n NacTag) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, n.AdditionalProperties)
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
        structMap["gbp_tag"] = n.GbpTag.toMap()
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
    if n.NacportalId != nil {
        structMap["nacportal_id"] = n.NacportalId
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
    if n.UsernameAttr != nil {
        structMap["username_attr"] = n.UsernameAttr
    }
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
    var temp tempNacTag
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "allow_usermac_override", "created_time", "egress_vlan_names", "gbp_tag", "id", "match", "match_all", "modified_time", "nacportal_id", "name", "org_id", "radius_attrs", "radius_group", "radius_vendor_attrs", "session_timeout", "type", "username_attr", "values", "vlan")
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
    n.NacportalId = temp.NacportalId
    n.Name = *temp.Name
    n.OrgId = temp.OrgId
    n.RadiusAttrs = temp.RadiusAttrs
    n.RadiusGroup = temp.RadiusGroup
    n.RadiusVendorAttrs = temp.RadiusVendorAttrs
    n.SessionTimeout = temp.SessionTimeout
    n.Type = *temp.Type
    n.UsernameAttr = temp.UsernameAttr
    n.Values = temp.Values
    n.Vlan = temp.Vlan
    return nil
}

// tempNacTag is a temporary struct used for validating the fields of NacTag.
type tempNacTag  struct {
    AllowUsermacOverride *bool                   `json:"allow_usermac_override,omitempty"`
    CreatedTime          *float64                `json:"created_time,omitempty"`
    EgressVlanNames      []string                `json:"egress_vlan_names,omitempty"`
    GbpTag               *NacTagGbpTag           `json:"gbp_tag,omitempty"`
    Id                   *uuid.UUID              `json:"id,omitempty"`
    Match                *NacTagMatchEnum        `json:"match,omitempty"`
    MatchAll             *bool                   `json:"match_all,omitempty"`
    ModifiedTime         *float64                `json:"modified_time,omitempty"`
    NacportalId          *uuid.UUID              `json:"nacportal_id,omitempty"`
    Name                 *string                 `json:"name"`
    OrgId                *uuid.UUID              `json:"org_id,omitempty"`
    RadiusAttrs          []string                `json:"radius_attrs,omitempty"`
    RadiusGroup          *string                 `json:"radius_group,omitempty"`
    RadiusVendorAttrs    []string                `json:"radius_vendor_attrs,omitempty"`
    SessionTimeout       *int                    `json:"session_timeout,omitempty"`
    Type                 *NacTagTypeEnum         `json:"type"`
    UsernameAttr         *NacTagUsernameAttrEnum `json:"username_attr,omitempty"`
    Values               []string                `json:"values,omitempty"`
    Vlan                 *string                 `json:"vlan,omitempty"`
}

func (n *tempNacTag) validate() error {
    var errs []string
    if n.Name == nil {
        errs = append(errs, "required field `name` is missing for type `nac_tag`")
    }
    if n.Type == nil {
        errs = append(errs, "required field `type` is missing for type `nac_tag`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
