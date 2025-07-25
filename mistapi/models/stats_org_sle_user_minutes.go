// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// StatsOrgSleUserMinutes represents a StatsOrgSleUserMinutes struct.
type StatsOrgSleUserMinutes struct {
    Ok                   float64                `json:"ok"`
    Total                float64                `json:"total"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for StatsOrgSleUserMinutes,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s StatsOrgSleUserMinutes) String() string {
    return fmt.Sprintf(
    	"StatsOrgSleUserMinutes[Ok=%v, Total=%v, AdditionalProperties=%v]",
    	s.Ok, s.Total, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for StatsOrgSleUserMinutes.
// It customizes the JSON marshaling process for StatsOrgSleUserMinutes objects.
func (s StatsOrgSleUserMinutes) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "ok", "total"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the StatsOrgSleUserMinutes object to a map representation for JSON marshaling.
func (s StatsOrgSleUserMinutes) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    structMap["ok"] = s.Ok
    structMap["total"] = s.Total
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsOrgSleUserMinutes.
// It customizes the JSON unmarshaling process for StatsOrgSleUserMinutes objects.
func (s *StatsOrgSleUserMinutes) UnmarshalJSON(input []byte) error {
    var temp tempStatsOrgSleUserMinutes
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "ok", "total")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Ok = *temp.Ok
    s.Total = *temp.Total
    return nil
}

// tempStatsOrgSleUserMinutes is a temporary struct used for validating the fields of StatsOrgSleUserMinutes.
type tempStatsOrgSleUserMinutes  struct {
    Ok    *float64 `json:"ok"`
    Total *float64 `json:"total"`
}

func (s *tempStatsOrgSleUserMinutes) validate() error {
    var errs []string
    if s.Ok == nil {
        errs = append(errs, "required field `ok` is missing for type `stats_org_sle_user_minutes`")
    }
    if s.Total == nil {
        errs = append(errs, "required field `total` is missing for type `stats_org_sle_user_minutes`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
