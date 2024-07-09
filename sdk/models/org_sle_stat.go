package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// OrgSleStat represents a OrgSleStat struct.
type OrgSleStat struct {
    Path                 string                 `json:"path"`
    UserMinutes          *OrgSleStatUserMinutes `json:"user_minutes,omitempty"`
    AdditionalProperties map[string]any         `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for OrgSleStat.
// It customizes the JSON marshaling process for OrgSleStat objects.
func (o OrgSleStat) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(o.toMap())
}

// toMap converts the OrgSleStat object to a map representation for JSON marshaling.
func (o OrgSleStat) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, o.AdditionalProperties)
    structMap["path"] = o.Path
    if o.UserMinutes != nil {
        structMap["user_minutes"] = o.UserMinutes.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgSleStat.
// It customizes the JSON unmarshaling process for OrgSleStat objects.
func (o *OrgSleStat) UnmarshalJSON(input []byte) error {
    var temp orgSleStat
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "path", "user_minutes")
    if err != nil {
    	return err
    }
    
    o.AdditionalProperties = additionalProperties
    o.Path = *temp.Path
    o.UserMinutes = temp.UserMinutes
    return nil
}

// orgSleStat is a temporary struct used for validating the fields of OrgSleStat.
type orgSleStat  struct {
    Path        *string                `json:"path"`
    UserMinutes *OrgSleStatUserMinutes `json:"user_minutes,omitempty"`
}

func (o *orgSleStat) validate() error {
    var errs []string
    if o.Path == nil {
        errs = append(errs, "required field `path` is missing for type `Org_Sle_Stat`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
