package models

import (
    "encoding/json"
    "github.com/google/uuid"
)

// OrgServicePolicy represents a OrgServicePolicy struct.
type OrgServicePolicy struct {
    Action               *AllowDenyEnum         `json:"action,omitempty"`
    // For SRX Only
    Appqoe               *ServicePolicyAppqoe   `json:"appqoe,omitempty"`
    CreatedTime          *float64               `json:"created_time,omitempty"`
    Ewf                  []ServicePolicyEwfRule `json:"ewf,omitempty"`
    Id                   *uuid.UUID             `json:"id,omitempty"`
    Idp                  *IdpConfig             `json:"idp,omitempty"`
    // access within the same VRF
    LocalRouting         *bool                  `json:"local_routing,omitempty"`
    ModifiedTime         *float64               `json:"modified_time,omitempty"`
    Name                 *string                `json:"name,omitempty"`
    OrgId                *uuid.UUID             `json:"org_id,omitempty"`
    // by default, we derive all paths available and use them
    // optionally, you can customize by using `path_preference`
    PathPreferences      *string                `json:"path_preferences,omitempty"`
    Services             []string               `json:"services,omitempty"`
    Tenants              []string               `json:"tenants,omitempty"`
    AdditionalProperties map[string]any         `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for OrgServicePolicy.
// It customizes the JSON marshaling process for OrgServicePolicy objects.
func (o OrgServicePolicy) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(o.toMap())
}

// toMap converts the OrgServicePolicy object to a map representation for JSON marshaling.
func (o OrgServicePolicy) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, o.AdditionalProperties)
    if o.Action != nil {
        structMap["action"] = o.Action
    }
    if o.Appqoe != nil {
        structMap["appqoe"] = o.Appqoe.toMap()
    }
    if o.CreatedTime != nil {
        structMap["created_time"] = o.CreatedTime
    }
    if o.Ewf != nil {
        structMap["ewf"] = o.Ewf
    }
    if o.Id != nil {
        structMap["id"] = o.Id
    }
    if o.Idp != nil {
        structMap["idp"] = o.Idp.toMap()
    }
    if o.LocalRouting != nil {
        structMap["local_routing"] = o.LocalRouting
    }
    if o.ModifiedTime != nil {
        structMap["modified_time"] = o.ModifiedTime
    }
    if o.Name != nil {
        structMap["name"] = o.Name
    }
    if o.OrgId != nil {
        structMap["org_id"] = o.OrgId
    }
    if o.PathPreferences != nil {
        structMap["path_preferences"] = o.PathPreferences
    }
    if o.Services != nil {
        structMap["services"] = o.Services
    }
    if o.Tenants != nil {
        structMap["tenants"] = o.Tenants
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgServicePolicy.
// It customizes the JSON unmarshaling process for OrgServicePolicy objects.
func (o *OrgServicePolicy) UnmarshalJSON(input []byte) error {
    var temp orgServicePolicy
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "action", "appqoe", "created_time", "ewf", "id", "idp", "local_routing", "modified_time", "name", "org_id", "path_preferences", "services", "tenants")
    if err != nil {
    	return err
    }
    
    o.AdditionalProperties = additionalProperties
    o.Action = temp.Action
    o.Appqoe = temp.Appqoe
    o.CreatedTime = temp.CreatedTime
    o.Ewf = temp.Ewf
    o.Id = temp.Id
    o.Idp = temp.Idp
    o.LocalRouting = temp.LocalRouting
    o.ModifiedTime = temp.ModifiedTime
    o.Name = temp.Name
    o.OrgId = temp.OrgId
    o.PathPreferences = temp.PathPreferences
    o.Services = temp.Services
    o.Tenants = temp.Tenants
    return nil
}

// orgServicePolicy is a temporary struct used for validating the fields of OrgServicePolicy.
type orgServicePolicy  struct {
    Action          *AllowDenyEnum         `json:"action,omitempty"`
    Appqoe          *ServicePolicyAppqoe   `json:"appqoe,omitempty"`
    CreatedTime     *float64               `json:"created_time,omitempty"`
    Ewf             []ServicePolicyEwfRule `json:"ewf,omitempty"`
    Id              *uuid.UUID             `json:"id,omitempty"`
    Idp             *IdpConfig             `json:"idp,omitempty"`
    LocalRouting    *bool                  `json:"local_routing,omitempty"`
    ModifiedTime    *float64               `json:"modified_time,omitempty"`
    Name            *string                `json:"name,omitempty"`
    OrgId           *uuid.UUID             `json:"org_id,omitempty"`
    PathPreferences *string                `json:"path_preferences,omitempty"`
    Services        []string               `json:"services,omitempty"`
    Tenants         []string               `json:"tenants,omitempty"`
}
