package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// StatsWrule represents a StatsWrule struct.
// Wxrule statistics
type StatsWrule struct {
    // enum: `allow`, `block`
    Action               StatsWruleActionEnum                 `json:"action"`
    ClientMac            []string                             `json:"client_mac"`
    DstAllowWxtags       []uuid.UUID                          `json:"dst_allow_wxtags"`
    DstDenyWxtags        []uuid.UUID                          `json:"dst_deny_wxtags"`
    DstWxtags            []uuid.UUID                          `json:"dst_wxtags"`
    Name                 string                               `json:"name"`
    Order                int                                  `json:"order"`
    SrcWxtags            []uuid.UUID                          `json:"src_wxtags"`
    Usage                map[string]StatsWruleUsageProperties `json:"usage"`
    AdditionalProperties map[string]any                       `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for StatsWrule.
// It customizes the JSON marshaling process for StatsWrule objects.
func (s StatsWrule) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the StatsWrule object to a map representation for JSON marshaling.
func (s StatsWrule) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
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

// UnmarshalJSON implements the json.Unmarshaler interface for StatsWrule.
// It customizes the JSON unmarshaling process for StatsWrule objects.
func (s *StatsWrule) UnmarshalJSON(input []byte) error {
    var temp tempStatsWrule
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

// tempStatsWrule is a temporary struct used for validating the fields of StatsWrule.
type tempStatsWrule  struct {
    Action         *StatsWruleActionEnum                 `json:"action"`
    ClientMac      *[]string                             `json:"client_mac"`
    DstAllowWxtags *[]uuid.UUID                          `json:"dst_allow_wxtags"`
    DstDenyWxtags  *[]uuid.UUID                          `json:"dst_deny_wxtags"`
    DstWxtags      *[]uuid.UUID                          `json:"dst_wxtags"`
    Name           *string                               `json:"name"`
    Order          *int                                  `json:"order"`
    SrcWxtags      *[]uuid.UUID                          `json:"src_wxtags"`
    Usage          *map[string]StatsWruleUsageProperties `json:"usage"`
}

func (s *tempStatsWrule) validate() error {
    var errs []string
    if s.Action == nil {
        errs = append(errs, "required field `action` is missing for type `stats_wrule`")
    }
    if s.ClientMac == nil {
        errs = append(errs, "required field `client_mac` is missing for type `stats_wrule`")
    }
    if s.DstAllowWxtags == nil {
        errs = append(errs, "required field `dst_allow_wxtags` is missing for type `stats_wrule`")
    }
    if s.DstDenyWxtags == nil {
        errs = append(errs, "required field `dst_deny_wxtags` is missing for type `stats_wrule`")
    }
    if s.DstWxtags == nil {
        errs = append(errs, "required field `dst_wxtags` is missing for type `stats_wrule`")
    }
    if s.Name == nil {
        errs = append(errs, "required field `name` is missing for type `stats_wrule`")
    }
    if s.Order == nil {
        errs = append(errs, "required field `order` is missing for type `stats_wrule`")
    }
    if s.SrcWxtags == nil {
        errs = append(errs, "required field `src_wxtags` is missing for type `stats_wrule`")
    }
    if s.Usage == nil {
        errs = append(errs, "required field `usage` is missing for type `stats_wrule`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
