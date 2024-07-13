package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// OrgSleStatUserMinutes represents a OrgSleStatUserMinutes struct.
type OrgSleStatUserMinutes struct {
    Ok                   float64        `json:"ok"`
    Total                float64        `json:"total"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for OrgSleStatUserMinutes.
// It customizes the JSON marshaling process for OrgSleStatUserMinutes objects.
func (o OrgSleStatUserMinutes) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(o.toMap())
}

// toMap converts the OrgSleStatUserMinutes object to a map representation for JSON marshaling.
func (o OrgSleStatUserMinutes) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, o.AdditionalProperties)
    structMap["ok"] = o.Ok
    structMap["total"] = o.Total
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgSleStatUserMinutes.
// It customizes the JSON unmarshaling process for OrgSleStatUserMinutes objects.
func (o *OrgSleStatUserMinutes) UnmarshalJSON(input []byte) error {
    var temp orgSleStatUserMinutes
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "ok", "total")
    if err != nil {
    	return err
    }
    
    o.AdditionalProperties = additionalProperties
    o.Ok = *temp.Ok
    o.Total = *temp.Total
    return nil
}

// orgSleStatUserMinutes is a temporary struct used for validating the fields of OrgSleStatUserMinutes.
type orgSleStatUserMinutes  struct {
    Ok    *float64 `json:"ok"`
    Total *float64 `json:"total"`
}

func (o *orgSleStatUserMinutes) validate() error {
    var errs []string
    if o.Ok == nil {
        errs = append(errs, "required field `ok` is missing for type `Org_Sle_Stat_User_Minutes`")
    }
    if o.Total == nil {
        errs = append(errs, "required field `total` is missing for type `Org_Sle_Stat_User_Minutes`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
