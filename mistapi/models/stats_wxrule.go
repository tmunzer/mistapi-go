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

// StatsWxrule represents a StatsWxrule struct.
// Wxrule statistics
type StatsWxrule struct {
    // enum: `allow`, `block`
    Action               StatsWxruleActionEnum                 `json:"action"`
    ClientMac            []string                              `json:"client_mac"`
    DstAllowWxtags       []uuid.UUID                           `json:"dst_allow_wxtags"`
    DstDenyWxtags        []uuid.UUID                           `json:"dst_deny_wxtags"`
    DstWxtags            []uuid.UUID                           `json:"dst_wxtags"`
    Name                 string                                `json:"name"`
    Order                int                                   `json:"order"`
    SrcWxtags            []uuid.UUID                           `json:"src_wxtags"`
    Usage                map[string]StatsWxruleUsageProperties `json:"usage"`
    AdditionalProperties map[string]interface{}                `json:"_"`
}

// String implements the fmt.Stringer interface for StatsWxrule,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s StatsWxrule) String() string {
    return fmt.Sprintf(
    	"StatsWxrule[Action=%v, ClientMac=%v, DstAllowWxtags=%v, DstDenyWxtags=%v, DstWxtags=%v, Name=%v, Order=%v, SrcWxtags=%v, Usage=%v, AdditionalProperties=%v]",
    	s.Action, s.ClientMac, s.DstAllowWxtags, s.DstDenyWxtags, s.DstWxtags, s.Name, s.Order, s.SrcWxtags, s.Usage, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for StatsWxrule.
// It customizes the JSON marshaling process for StatsWxrule objects.
func (s StatsWxrule) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "action", "client_mac", "dst_allow_wxtags", "dst_deny_wxtags", "dst_wxtags", "name", "order", "src_wxtags", "usage"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the StatsWxrule object to a map representation for JSON marshaling.
func (s StatsWxrule) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    structMap["action"] = s.Action
    structMap["client_mac"] = s.ClientMac
    structMap["dst_allow_wxtags"] = s.DstAllowWxtags
    structMap["dst_deny_wxtags"] = s.DstDenyWxtags
    structMap["dst_wxtags"] = s.DstWxtags
    structMap["name"] = s.Name
    structMap["order"] = s.Order
    structMap["src_wxtags"] = s.SrcWxtags
    structMap["usage"] = s.Usage
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsWxrule.
// It customizes the JSON unmarshaling process for StatsWxrule objects.
func (s *StatsWxrule) UnmarshalJSON(input []byte) error {
    var temp tempStatsWxrule
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "action", "client_mac", "dst_allow_wxtags", "dst_deny_wxtags", "dst_wxtags", "name", "order", "src_wxtags", "usage")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Action = *temp.Action
    s.ClientMac = *temp.ClientMac
    s.DstAllowWxtags = *temp.DstAllowWxtags
    s.DstDenyWxtags = *temp.DstDenyWxtags
    s.DstWxtags = *temp.DstWxtags
    s.Name = *temp.Name
    s.Order = *temp.Order
    s.SrcWxtags = *temp.SrcWxtags
    s.Usage = *temp.Usage
    return nil
}

// tempStatsWxrule is a temporary struct used for validating the fields of StatsWxrule.
type tempStatsWxrule  struct {
    Action         *StatsWxruleActionEnum                 `json:"action"`
    ClientMac      *[]string                              `json:"client_mac"`
    DstAllowWxtags *[]uuid.UUID                           `json:"dst_allow_wxtags"`
    DstDenyWxtags  *[]uuid.UUID                           `json:"dst_deny_wxtags"`
    DstWxtags      *[]uuid.UUID                           `json:"dst_wxtags"`
    Name           *string                                `json:"name"`
    Order          *int                                   `json:"order"`
    SrcWxtags      *[]uuid.UUID                           `json:"src_wxtags"`
    Usage          *map[string]StatsWxruleUsageProperties `json:"usage"`
}

func (s *tempStatsWxrule) validate() error {
    var errs []string
    if s.Action == nil {
        errs = append(errs, "required field `action` is missing for type `stats_wxrule`")
    }
    if s.ClientMac == nil {
        errs = append(errs, "required field `client_mac` is missing for type `stats_wxrule`")
    }
    if s.DstAllowWxtags == nil {
        errs = append(errs, "required field `dst_allow_wxtags` is missing for type `stats_wxrule`")
    }
    if s.DstDenyWxtags == nil {
        errs = append(errs, "required field `dst_deny_wxtags` is missing for type `stats_wxrule`")
    }
    if s.DstWxtags == nil {
        errs = append(errs, "required field `dst_wxtags` is missing for type `stats_wxrule`")
    }
    if s.Name == nil {
        errs = append(errs, "required field `name` is missing for type `stats_wxrule`")
    }
    if s.Order == nil {
        errs = append(errs, "required field `order` is missing for type `stats_wxrule`")
    }
    if s.SrcWxtags == nil {
        errs = append(errs, "required field `src_wxtags` is missing for type `stats_wxrule`")
    }
    if s.Usage == nil {
        errs = append(errs, "required field `usage` is missing for type `stats_wxrule`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
