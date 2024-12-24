package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/google/uuid"
    "strings"
)

// WxlanRule represents a WxlanRule struct.
// WXlan
type WxlanRule struct {
    // type of action, allow / block. enum: `allow`, `block`
    Action               *WxlanRuleActionEnum   `json:"action,omitempty"`
    ApplyTags            []string               `json:"apply_tags,omitempty"`
    // blocked apps (always blocking, ignoring action), the key of Get Application List
    BlockedApps          []string               `json:"blocked_apps,omitempty"`
    // when the object has been created, in epoch
    CreatedTime          *float64               `json:"created_time,omitempty"`
    // List of WxTag UUID to indicate these tags are allowed access
    DstAllowWxtags       []string               `json:"dst_allow_wxtags"`
    // List of WxTag UUID to indicate these tags are blocked access
    DstDenyWxtags        []string               `json:"dst_deny_wxtags"`
    // List of WxTag UUID
    DstWxtags            []string               `json:"dst_wxtags,omitempty"`
    Enabled              *bool                  `json:"enabled,omitempty"`
    ForSite              *bool                  `json:"for_site,omitempty"`
    // Unique ID of the object instance in the Mist Organnization
    Id                   *uuid.UUID             `json:"id,omitempty"`
    // when the object has been modified for the last time, in epoch
    ModifiedTime         *float64               `json:"modified_time,omitempty"`
    // the order how rules would be looked up, > 0 and bigger order got matched first, -1 means LAST, uniqueness not checked
    Order                int                    `json:"order"`
    OrgId                *uuid.UUID             `json:"org_id,omitempty"`
    SiteId               *uuid.UUID             `json:"site_id,omitempty"`
    // List of WxTag UUID to determine if this rule would match
    SrcWxtags            []string               `json:"src_wxtags"`
    // Only for Org Level WxRule
    TemplateId           *uuid.UUID             `json:"template_id,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for WxlanRule,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WxlanRule) String() string {
    return fmt.Sprintf(
    	"WxlanRule[Action=%v, ApplyTags=%v, BlockedApps=%v, CreatedTime=%v, DstAllowWxtags=%v, DstDenyWxtags=%v, DstWxtags=%v, Enabled=%v, ForSite=%v, Id=%v, ModifiedTime=%v, Order=%v, OrgId=%v, SiteId=%v, SrcWxtags=%v, TemplateId=%v, AdditionalProperties=%v]",
    	w.Action, w.ApplyTags, w.BlockedApps, w.CreatedTime, w.DstAllowWxtags, w.DstDenyWxtags, w.DstWxtags, w.Enabled, w.ForSite, w.Id, w.ModifiedTime, w.Order, w.OrgId, w.SiteId, w.SrcWxtags, w.TemplateId, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WxlanRule.
// It customizes the JSON marshaling process for WxlanRule objects.
func (w WxlanRule) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(w.AdditionalProperties,
        "action", "apply_tags", "blocked_apps", "created_time", "dst_allow_wxtags", "dst_deny_wxtags", "dst_wxtags", "enabled", "for_site", "id", "modified_time", "order", "org_id", "site_id", "src_wxtags", "template_id"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WxlanRule object to a map representation for JSON marshaling.
func (w WxlanRule) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, w.AdditionalProperties)
    if w.Action != nil {
        structMap["action"] = w.Action
    }
    if w.ApplyTags != nil {
        structMap["apply_tags"] = w.ApplyTags
    }
    if w.BlockedApps != nil {
        structMap["blocked_apps"] = w.BlockedApps
    }
    if w.CreatedTime != nil {
        structMap["created_time"] = w.CreatedTime
    }
    structMap["dst_allow_wxtags"] = w.DstAllowWxtags
    structMap["dst_deny_wxtags"] = w.DstDenyWxtags
    if w.DstWxtags != nil {
        structMap["dst_wxtags"] = w.DstWxtags
    }
    if w.Enabled != nil {
        structMap["enabled"] = w.Enabled
    }
    if w.ForSite != nil {
        structMap["for_site"] = w.ForSite
    }
    if w.Id != nil {
        structMap["id"] = w.Id
    }
    if w.ModifiedTime != nil {
        structMap["modified_time"] = w.ModifiedTime
    }
    structMap["order"] = w.Order
    if w.OrgId != nil {
        structMap["org_id"] = w.OrgId
    }
    if w.SiteId != nil {
        structMap["site_id"] = w.SiteId
    }
    structMap["src_wxtags"] = w.SrcWxtags
    if w.TemplateId != nil {
        structMap["template_id"] = w.TemplateId
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WxlanRule.
// It customizes the JSON unmarshaling process for WxlanRule objects.
func (w *WxlanRule) UnmarshalJSON(input []byte) error {
    var temp tempWxlanRule
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "action", "apply_tags", "blocked_apps", "created_time", "dst_allow_wxtags", "dst_deny_wxtags", "dst_wxtags", "enabled", "for_site", "id", "modified_time", "order", "org_id", "site_id", "src_wxtags", "template_id")
    if err != nil {
    	return err
    }
    w.AdditionalProperties = additionalProperties
    
    w.Action = temp.Action
    w.ApplyTags = temp.ApplyTags
    w.BlockedApps = temp.BlockedApps
    w.CreatedTime = temp.CreatedTime
    w.DstAllowWxtags = *temp.DstAllowWxtags
    w.DstDenyWxtags = *temp.DstDenyWxtags
    w.DstWxtags = temp.DstWxtags
    w.Enabled = temp.Enabled
    w.ForSite = temp.ForSite
    w.Id = temp.Id
    w.ModifiedTime = temp.ModifiedTime
    w.Order = *temp.Order
    w.OrgId = temp.OrgId
    w.SiteId = temp.SiteId
    w.SrcWxtags = *temp.SrcWxtags
    w.TemplateId = temp.TemplateId
    return nil
}

// tempWxlanRule is a temporary struct used for validating the fields of WxlanRule.
type tempWxlanRule  struct {
    Action         *WxlanRuleActionEnum `json:"action,omitempty"`
    ApplyTags      []string             `json:"apply_tags,omitempty"`
    BlockedApps    []string             `json:"blocked_apps,omitempty"`
    CreatedTime    *float64             `json:"created_time,omitempty"`
    DstAllowWxtags *[]string            `json:"dst_allow_wxtags"`
    DstDenyWxtags  *[]string            `json:"dst_deny_wxtags"`
    DstWxtags      []string             `json:"dst_wxtags,omitempty"`
    Enabled        *bool                `json:"enabled,omitempty"`
    ForSite        *bool                `json:"for_site,omitempty"`
    Id             *uuid.UUID           `json:"id,omitempty"`
    ModifiedTime   *float64             `json:"modified_time,omitempty"`
    Order          *int                 `json:"order"`
    OrgId          *uuid.UUID           `json:"org_id,omitempty"`
    SiteId         *uuid.UUID           `json:"site_id,omitempty"`
    SrcWxtags      *[]string            `json:"src_wxtags"`
    TemplateId     *uuid.UUID           `json:"template_id,omitempty"`
}

func (w *tempWxlanRule) validate() error {
    var errs []string
    if w.DstAllowWxtags == nil {
        errs = append(errs, "required field `dst_allow_wxtags` is missing for type `wxlan_rule`")
    }
    if w.DstDenyWxtags == nil {
        errs = append(errs, "required field `dst_deny_wxtags` is missing for type `wxlan_rule`")
    }
    if w.Order == nil {
        errs = append(errs, "required field `order` is missing for type `wxlan_rule`")
    }
    if w.SrcWxtags == nil {
        errs = append(errs, "required field `src_wxtags` is missing for type `wxlan_rule`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
