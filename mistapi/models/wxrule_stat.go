package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// WxruleStat represents a WxruleStat struct.
// Wxrule statistics
type WxruleStat struct {
    // enum: `allow`, `block`
    Action               WxruleStatsActionEnum                 `json:"action"`
    ClientMac            []string                              `json:"client_mac"`
    DstAllowWxtags       []uuid.UUID                           `json:"dst_allow_wxtags"`
    DstDenyWxtags        []uuid.UUID                           `json:"dst_deny_wxtags"`
    DstWxtags            []uuid.UUID                           `json:"dst_wxtags"`
    Name                 string                                `json:"name"`
    Order                int                                   `json:"order"`
    SrcWxtags            []uuid.UUID                           `json:"src_wxtags"`
    Usage                map[string]WxruleStatsUsageProperties `json:"usage"`
    AdditionalProperties map[string]any                        `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for WxruleStat.
// It customizes the JSON marshaling process for WxruleStat objects.
func (w WxruleStat) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(w.toMap())
}

// toMap converts the WxruleStat object to a map representation for JSON marshaling.
func (w WxruleStat) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, w.AdditionalProperties)
    structMap["action"] = w.Action
    structMap["client_mac"] = w.ClientMac
    structMap["dst_allow_wxtags"] = w.DstAllowWxtags
    structMap["dst_deny_wxtags"] = w.DstDenyWxtags
    structMap["dst_wxtags"] = w.DstWxtags
    structMap["name"] = w.Name
    structMap["order"] = w.Order
    structMap["src_wxtags"] = w.SrcWxtags
    structMap["usage"] = w.Usage
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WxruleStat.
// It customizes the JSON unmarshaling process for WxruleStat objects.
func (w *WxruleStat) UnmarshalJSON(input []byte) error {
    var temp tempWxruleStat
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "action", "client_mac", "dst_allow_wxtags", "dst_deny_wxtags", "dst_wxtags", "name", "order", "src_wxtags", "usage")
    if err != nil {
    	return err
    }
    
    w.AdditionalProperties = additionalProperties
    w.Action = *temp.Action
    w.ClientMac = *temp.ClientMac
    w.DstAllowWxtags = *temp.DstAllowWxtags
    w.DstDenyWxtags = *temp.DstDenyWxtags
    w.DstWxtags = *temp.DstWxtags
    w.Name = *temp.Name
    w.Order = *temp.Order
    w.SrcWxtags = *temp.SrcWxtags
    w.Usage = *temp.Usage
    return nil
}

// tempWxruleStat is a temporary struct used for validating the fields of WxruleStat.
type tempWxruleStat  struct {
    Action         *WxruleStatsActionEnum                 `json:"action"`
    ClientMac      *[]string                              `json:"client_mac"`
    DstAllowWxtags *[]uuid.UUID                           `json:"dst_allow_wxtags"`
    DstDenyWxtags  *[]uuid.UUID                           `json:"dst_deny_wxtags"`
    DstWxtags      *[]uuid.UUID                           `json:"dst_wxtags"`
    Name           *string                                `json:"name"`
    Order          *int                                   `json:"order"`
    SrcWxtags      *[]uuid.UUID                           `json:"src_wxtags"`
    Usage          *map[string]WxruleStatsUsageProperties `json:"usage"`
}

func (w *tempWxruleStat) validate() error {
    var errs []string
    if w.Action == nil {
        errs = append(errs, "required field `action` is missing for type `wxrule_stat`")
    }
    if w.ClientMac == nil {
        errs = append(errs, "required field `client_mac` is missing for type `wxrule_stat`")
    }
    if w.DstAllowWxtags == nil {
        errs = append(errs, "required field `dst_allow_wxtags` is missing for type `wxrule_stat`")
    }
    if w.DstDenyWxtags == nil {
        errs = append(errs, "required field `dst_deny_wxtags` is missing for type `wxrule_stat`")
    }
    if w.DstWxtags == nil {
        errs = append(errs, "required field `dst_wxtags` is missing for type `wxrule_stat`")
    }
    if w.Name == nil {
        errs = append(errs, "required field `name` is missing for type `wxrule_stat`")
    }
    if w.Order == nil {
        errs = append(errs, "required field `order` is missing for type `wxrule_stat`")
    }
    if w.SrcWxtags == nil {
        errs = append(errs, "required field `src_wxtags` is missing for type `wxrule_stat`")
    }
    if w.Usage == nil {
        errs = append(errs, "required field `usage` is missing for type `wxrule_stat`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
