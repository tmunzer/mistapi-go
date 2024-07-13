package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// Orggroup represents a Orggroup struct.
// Organizations Group
type Orggroup struct {
    CreatedTime          *float64       `json:"created_time,omitempty"`
    Id                   *uuid.UUID     `json:"id,omitempty"`
    ModifiedTime         *float64       `json:"modified_time,omitempty"`
    MspId                *uuid.UUID     `json:"msp_id,omitempty"`
    Name                 string         `json:"name"`
    OrgIds               []uuid.UUID    `json:"org_ids,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for Orggroup.
// It customizes the JSON marshaling process for Orggroup objects.
func (o Orggroup) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(o.toMap())
}

// toMap converts the Orggroup object to a map representation for JSON marshaling.
func (o Orggroup) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, o.AdditionalProperties)
    if o.CreatedTime != nil {
        structMap["created_time"] = o.CreatedTime
    }
    if o.Id != nil {
        structMap["id"] = o.Id
    }
    if o.ModifiedTime != nil {
        structMap["modified_time"] = o.ModifiedTime
    }
    if o.MspId != nil {
        structMap["msp_id"] = o.MspId
    }
    structMap["name"] = o.Name
    if o.OrgIds != nil {
        structMap["org_ids"] = o.OrgIds
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Orggroup.
// It customizes the JSON unmarshaling process for Orggroup objects.
func (o *Orggroup) UnmarshalJSON(input []byte) error {
    var temp orggroup
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "created_time", "id", "modified_time", "msp_id", "name", "org_ids")
    if err != nil {
    	return err
    }
    
    o.AdditionalProperties = additionalProperties
    o.CreatedTime = temp.CreatedTime
    o.Id = temp.Id
    o.ModifiedTime = temp.ModifiedTime
    o.MspId = temp.MspId
    o.Name = *temp.Name
    o.OrgIds = temp.OrgIds
    return nil
}

// orggroup is a temporary struct used for validating the fields of Orggroup.
type orggroup  struct {
    CreatedTime  *float64    `json:"created_time,omitempty"`
    Id           *uuid.UUID  `json:"id,omitempty"`
    ModifiedTime *float64    `json:"modified_time,omitempty"`
    MspId        *uuid.UUID  `json:"msp_id,omitempty"`
    Name         *string     `json:"name"`
    OrgIds       []uuid.UUID `json:"org_ids,omitempty"`
}

func (o *orggroup) validate() error {
    var errs []string
    if o.Name == nil {
        errs = append(errs, "required field `name` is missing for type `Orggroup`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
