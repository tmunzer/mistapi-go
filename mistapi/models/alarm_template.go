package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// AlarmTemplate represents a AlarmTemplate struct.
// Alarm Template
type AlarmTemplate struct {
    CreatedTime          *float64                     `json:"created_time,omitempty"`
    // Delivery object to configure the alarm delivery
    Delivery             Delivery                     `json:"delivery"`
    Id                   *uuid.UUID                   `json:"id,omitempty"`
    ModifiedTime         *float64                     `json:"modified_time,omitempty"`
    // Some string to name the alarm template
    Name                 *string                      `json:"name,omitempty"`
    OrgId                *uuid.UUID                   `json:"org_id,omitempty"`
    // Alarm Rules object to configure the individual alarm keys/types. Property key is the alarm name.
    Rules                map[string]AlarmTemplateRule `json:"rules"`
    AdditionalProperties map[string]any               `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for AlarmTemplate.
// It customizes the JSON marshaling process for AlarmTemplate objects.
func (a AlarmTemplate) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the AlarmTemplate object to a map representation for JSON marshaling.
func (a AlarmTemplate) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, a.AdditionalProperties)
    if a.CreatedTime != nil {
        structMap["created_time"] = a.CreatedTime
    }
    structMap["delivery"] = a.Delivery.toMap()
    if a.Id != nil {
        structMap["id"] = a.Id
    }
    if a.ModifiedTime != nil {
        structMap["modified_time"] = a.ModifiedTime
    }
    if a.Name != nil {
        structMap["name"] = a.Name
    }
    if a.OrgId != nil {
        structMap["org_id"] = a.OrgId
    }
    structMap["rules"] = a.Rules
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AlarmTemplate.
// It customizes the JSON unmarshaling process for AlarmTemplate objects.
func (a *AlarmTemplate) UnmarshalJSON(input []byte) error {
    var temp tempAlarmTemplate
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "created_time", "delivery", "id", "modified_time", "name", "org_id", "rules")
    if err != nil {
    	return err
    }
    
    a.AdditionalProperties = additionalProperties
    a.CreatedTime = temp.CreatedTime
    a.Delivery = *temp.Delivery
    a.Id = temp.Id
    a.ModifiedTime = temp.ModifiedTime
    a.Name = temp.Name
    a.OrgId = temp.OrgId
    a.Rules = *temp.Rules
    return nil
}

// tempAlarmTemplate is a temporary struct used for validating the fields of AlarmTemplate.
type tempAlarmTemplate  struct {
    CreatedTime  *float64                      `json:"created_time,omitempty"`
    Delivery     *Delivery                     `json:"delivery"`
    Id           *uuid.UUID                    `json:"id,omitempty"`
    ModifiedTime *float64                      `json:"modified_time,omitempty"`
    Name         *string                       `json:"name,omitempty"`
    OrgId        *uuid.UUID                    `json:"org_id,omitempty"`
    Rules        *map[string]AlarmTemplateRule `json:"rules"`
}

func (a *tempAlarmTemplate) validate() error {
    var errs []string
    if a.Delivery == nil {
        errs = append(errs, "required field `delivery` is missing for type `alarm_template`")
    }
    if a.Rules == nil {
        errs = append(errs, "required field `rules` is missing for type `alarm_template`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
