package models

import (
    "encoding/json"
    "fmt"
    "github.com/google/uuid"
)

// ResponseSearchVarItem represents a ResponseSearchVarItem struct.
type ResponseSearchVarItem struct {
    // When the object has been created, in epoch
    CreatedTime          *float64               `json:"created_time,omitempty"`
    // When the object has been modified for the last time, in epoch
    ModifiedTime         *float64               `json:"modified_time,omitempty"`
    OrgId                *uuid.UUID             `json:"org_id,omitempty"`
    SiteId               *uuid.UUID             `json:"site_id,omitempty"`
    Src                  *string                `json:"src,omitempty"`
    Var                  *string                `json:"var,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseSearchVarItem,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseSearchVarItem) String() string {
    return fmt.Sprintf(
    	"ResponseSearchVarItem[CreatedTime=%v, ModifiedTime=%v, OrgId=%v, SiteId=%v, Src=%v, Var=%v, AdditionalProperties=%v]",
    	r.CreatedTime, r.ModifiedTime, r.OrgId, r.SiteId, r.Src, r.Var, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseSearchVarItem.
// It customizes the JSON marshaling process for ResponseSearchVarItem objects.
func (r ResponseSearchVarItem) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "created_time", "modified_time", "org_id", "site_id", "src", "var"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseSearchVarItem object to a map representation for JSON marshaling.
func (r ResponseSearchVarItem) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
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
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "created_time", "modified_time", "org_id", "site_id", "src", "var")
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
