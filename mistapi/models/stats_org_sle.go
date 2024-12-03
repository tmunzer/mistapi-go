package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// StatsOrgSle represents a StatsOrgSle struct.
type StatsOrgSle struct {
    Path                 string                  `json:"path"`
    UserMinutes          *StatsOrgSleUserMinutes `json:"user_minutes,omitempty"`
    AdditionalProperties map[string]interface{}  `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for StatsOrgSle.
// It customizes the JSON marshaling process for StatsOrgSle objects.
func (s StatsOrgSle) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "path", "user_minutes"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the StatsOrgSle object to a map representation for JSON marshaling.
func (s StatsOrgSle) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    structMap["path"] = s.Path
    if s.UserMinutes != nil {
        structMap["user_minutes"] = s.UserMinutes.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsOrgSle.
// It customizes the JSON unmarshaling process for StatsOrgSle objects.
func (s *StatsOrgSle) UnmarshalJSON(input []byte) error {
    var temp tempStatsOrgSle
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "path", "user_minutes")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Path = *temp.Path
    s.UserMinutes = temp.UserMinutes
    return nil
}

// tempStatsOrgSle is a temporary struct used for validating the fields of StatsOrgSle.
type tempStatsOrgSle  struct {
    Path        *string                 `json:"path"`
    UserMinutes *StatsOrgSleUserMinutes `json:"user_minutes,omitempty"`
}

func (s *tempStatsOrgSle) validate() error {
    var errs []string
    if s.Path == nil {
        errs = append(errs, "required field `path` is missing for type `stats_org_sle`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
