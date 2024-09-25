package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// OrgApitoken represents a OrgApitoken struct.
// Org API Token
// **Note:**
// `privilege` field is required to create the object, but may not be 
// returned in the POST API Response (only in the afterward GET)
type OrgApitoken struct {
    // email of the token creator / null if creator is deleted
    CreatedBy            Optional[string]  `json:"created_by"`
    CreatedTime          *float64          `json:"created_time,omitempty"`
    Id                   *uuid.UUID        `json:"id,omitempty"`
    Key                  *string           `json:"key,omitempty"`
    LastUsed             Optional[float64] `json:"last_used"`
    // name of the token
    Name                 string            `json:"name"`
    OrgId                *uuid.UUID        `json:"org_id,omitempty"`
    // list of privileges the token has on the orgs/sites
    Privileges           []PrivilegeOrg    `json:"privileges,omitempty"`
    // list of allowed IP addresses from where the token can be used from. At most 10 IP addresses can be specified, cannot be changed once the API Token is created.
    SrcIps               []string          `json:"src_ips,omitempty"`
    AdditionalProperties map[string]any    `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for OrgApitoken.
// It customizes the JSON marshaling process for OrgApitoken objects.
func (o OrgApitoken) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(o.toMap())
}

// toMap converts the OrgApitoken object to a map representation for JSON marshaling.
func (o OrgApitoken) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, o.AdditionalProperties)
    if o.CreatedBy.IsValueSet() {
        if o.CreatedBy.Value() != nil {
            structMap["created_by"] = o.CreatedBy.Value()
        } else {
            structMap["created_by"] = nil
        }
    }
    if o.CreatedTime != nil {
        structMap["created_time"] = o.CreatedTime
    }
    if o.Id != nil {
        structMap["id"] = o.Id
    }
    if o.Key != nil {
        structMap["key"] = o.Key
    }
    if o.LastUsed.IsValueSet() {
        if o.LastUsed.Value() != nil {
            structMap["last_used"] = o.LastUsed.Value()
        } else {
            structMap["last_used"] = nil
        }
    }
    structMap["name"] = o.Name
    if o.OrgId != nil {
        structMap["org_id"] = o.OrgId
    }
    if o.Privileges != nil {
        structMap["privileges"] = o.Privileges
    }
    if o.SrcIps != nil {
        structMap["src_ips"] = o.SrcIps
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgApitoken.
// It customizes the JSON unmarshaling process for OrgApitoken objects.
func (o *OrgApitoken) UnmarshalJSON(input []byte) error {
    var temp tempOrgApitoken
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "created_by", "created_time", "id", "key", "last_used", "name", "org_id", "privileges", "src_ips")
    if err != nil {
    	return err
    }
    
    o.AdditionalProperties = additionalProperties
    o.CreatedBy = temp.CreatedBy
    o.CreatedTime = temp.CreatedTime
    o.Id = temp.Id
    o.Key = temp.Key
    o.LastUsed = temp.LastUsed
    o.Name = *temp.Name
    o.OrgId = temp.OrgId
    o.Privileges = temp.Privileges
    o.SrcIps = temp.SrcIps
    return nil
}

// tempOrgApitoken is a temporary struct used for validating the fields of OrgApitoken.
type tempOrgApitoken  struct {
    CreatedBy   Optional[string]  `json:"created_by"`
    CreatedTime *float64          `json:"created_time,omitempty"`
    Id          *uuid.UUID        `json:"id,omitempty"`
    Key         *string           `json:"key,omitempty"`
    LastUsed    Optional[float64] `json:"last_used"`
    Name        *string           `json:"name"`
    OrgId       *uuid.UUID        `json:"org_id,omitempty"`
    Privileges  []PrivilegeOrg    `json:"privileges,omitempty"`
    SrcIps      []string          `json:"src_ips,omitempty"`
}

func (o *tempOrgApitoken) validate() error {
    var errs []string
    if o.Name == nil {
        errs = append(errs, "required field `name` is missing for type `org_apitoken`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
