package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// Vpn represents a Vpn struct.
type Vpn struct {
    CreatedTime          *float64           `json:"created_time,omitempty"`
    Id                   *uuid.UUID         `json:"id,omitempty"`
    ModifiedTime         *float64           `json:"modified_time,omitempty"`
    Name                 string             `json:"name"`
    OrgId                *uuid.UUID         `json:"org_id,omitempty"`
    Paths                map[string]VpnPath `json:"paths"`
    AdditionalProperties map[string]any     `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for Vpn.
// It customizes the JSON marshaling process for Vpn objects.
func (v Vpn) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(v.toMap())
}

// toMap converts the Vpn object to a map representation for JSON marshaling.
func (v Vpn) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, v.AdditionalProperties)
    if v.CreatedTime != nil {
        structMap["created_time"] = v.CreatedTime
    }
    if v.Id != nil {
        structMap["id"] = v.Id
    }
    if v.ModifiedTime != nil {
        structMap["modified_time"] = v.ModifiedTime
    }
    structMap["name"] = v.Name
    if v.OrgId != nil {
        structMap["org_id"] = v.OrgId
    }
    structMap["paths"] = v.Paths
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Vpn.
// It customizes the JSON unmarshaling process for Vpn objects.
func (v *Vpn) UnmarshalJSON(input []byte) error {
    var temp vpn
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "created_time", "id", "modified_time", "name", "org_id", "paths")
    if err != nil {
    	return err
    }
    
    v.AdditionalProperties = additionalProperties
    v.CreatedTime = temp.CreatedTime
    v.Id = temp.Id
    v.ModifiedTime = temp.ModifiedTime
    v.Name = *temp.Name
    v.OrgId = temp.OrgId
    v.Paths = *temp.Paths
    return nil
}

// vpn is a temporary struct used for validating the fields of Vpn.
type vpn  struct {
    CreatedTime  *float64            `json:"created_time,omitempty"`
    Id           *uuid.UUID          `json:"id,omitempty"`
    ModifiedTime *float64            `json:"modified_time,omitempty"`
    Name         *string             `json:"name"`
    OrgId        *uuid.UUID          `json:"org_id,omitempty"`
    Paths        *map[string]VpnPath `json:"paths"`
}

func (v *vpn) validate() error {
    var errs []string
    if v.Name == nil {
        errs = append(errs, "required field `name` is missing for type `Vpn`")
    }
    if v.Paths == nil {
        errs = append(errs, "required field `paths` is missing for type `Vpn`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
