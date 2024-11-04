package models

import (
    "encoding/json"
    "github.com/google/uuid"
)

// ResponseSearchVarItem represents a ResponseSearchVarItem struct.
type ResponseSearchVarItem struct {
    // when the object has been created, in epoch
    CreatedTime          *float64       `json:"created_time,omitempty"`
    // when the object has been modified for the last time, in epoch
    ModifiedTime         *float64       `json:"modified_time,omitempty"`
    OrgId                *uuid.UUID     `json:"org_id,omitempty"`
    SiteId               *uuid.UUID     `json:"site_id,omitempty"`
    Src                  *string        `json:"src,omitempty"`
    Var                  *string        `json:"var,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseSearchVarItem.
// It customizes the JSON marshaling process for ResponseSearchVarItem objects.
func (r ResponseSearchVarItem) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseSearchVarItem object to a map representation for JSON marshaling.
func (r ResponseSearchVarItem) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    if r.CreatedTime != nil {
        structMap["created_time"] = r.CreatedTime
    }
    if r.ModifiedTime != nil {
        structMap["modified_time"] = r.ModifiedTime
    }
    if r.OrgId != nil {
        structMap["org_id"] = r.OrgId
    }
    if r.SiteId != nil {
        structMap["site_id"] = r.SiteId
    }
    if r.Src != nil {
        structMap["src"] = r.Src
    }
    if r.Var != nil {
        structMap["var"] = r.Var
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseSearchVarItem.
// It customizes the JSON unmarshaling process for ResponseSearchVarItem objects.
func (r *ResponseSearchVarItem) UnmarshalJSON(input []byte) error {
    var temp tempResponseSearchVarItem
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "created_time", "modified_time", "org_id", "site_id", "src", "var")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.CreatedTime = temp.CreatedTime
    r.ModifiedTime = temp.ModifiedTime
    r.OrgId = temp.OrgId
    r.SiteId = temp.SiteId
    r.Src = temp.Src
    r.Var = temp.Var
    return nil
}

// tempResponseSearchVarItem is a temporary struct used for validating the fields of ResponseSearchVarItem.
type tempResponseSearchVarItem  struct {
    CreatedTime  *float64   `json:"created_time,omitempty"`
    ModifiedTime *float64   `json:"modified_time,omitempty"`
    OrgId        *uuid.UUID `json:"org_id,omitempty"`
    SiteId       *uuid.UUID `json:"site_id,omitempty"`
    Src          *string    `json:"src,omitempty"`
    Var          *string    `json:"var,omitempty"`
}
