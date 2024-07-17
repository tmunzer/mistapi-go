package models

import (
    "encoding/json"
    "github.com/google/uuid"
)

// ServicePolicy represents a ServicePolicy struct.
type ServicePolicy struct {
    Action               *AllowDenyEnum         `json:"action,omitempty"`
    // For SRX Only
    Appqoe               *ServicePolicyAppqoe   `json:"appqoe,omitempty"`
    CreatedTime          *float64               `json:"created_time,omitempty"`
    Ewf                  []ServicePolicyEwfRule `json:"ewf,omitempty"`
    // used to link servicepolicy defined at org level and overwrite some attributes
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

// MarshalJSON implements the json.Marshaler interface for ServicePolicy.
// It customizes the JSON marshaling process for ServicePolicy objects.
func (s ServicePolicy) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the ServicePolicy object to a map representation for JSON marshaling.
func (s ServicePolicy) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Action != nil {
        structMap["action"] = s.Action
    }
    if s.Appqoe != nil {
        structMap["appqoe"] = s.Appqoe.toMap()
    }
    if s.CreatedTime != nil {
        structMap["created_time"] = s.CreatedTime
    }
    if s.Ewf != nil {
        structMap["ewf"] = s.Ewf
    }
    if s.Id != nil {
        structMap["id"] = s.Id
    }
    if s.Idp != nil {
        structMap["idp"] = s.Idp.toMap()
    }
    if s.LocalRouting != nil {
        structMap["local_routing"] = s.LocalRouting
    }
    if s.ModifiedTime != nil {
        structMap["modified_time"] = s.ModifiedTime
    }
    if s.Name != nil {
        structMap["name"] = s.Name
    }
    if s.OrgId != nil {
        structMap["org_id"] = s.OrgId
    }
    if s.PathPreferences != nil {
        structMap["path_preferences"] = s.PathPreferences
    }
    if s.Services != nil {
        structMap["services"] = s.Services
    }
    if s.Tenants != nil {
        structMap["tenants"] = s.Tenants
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ServicePolicy.
// It customizes the JSON unmarshaling process for ServicePolicy objects.
func (s *ServicePolicy) UnmarshalJSON(input []byte) error {
    var temp servicePolicy
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "action", "appqoe", "created_time", "ewf", "id", "idp", "local_routing", "modified_time", "name", "org_id", "path_preferences", "services", "tenants")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.Action = temp.Action
    s.Appqoe = temp.Appqoe
    s.CreatedTime = temp.CreatedTime
    s.Ewf = temp.Ewf
    s.Id = temp.Id
    s.Idp = temp.Idp
    s.LocalRouting = temp.LocalRouting
    s.ModifiedTime = temp.ModifiedTime
    s.Name = temp.Name
    s.OrgId = temp.OrgId
    s.PathPreferences = temp.PathPreferences
    s.Services = temp.Services
    s.Tenants = temp.Tenants
    return nil
}

// servicePolicy is a temporary struct used for validating the fields of ServicePolicy.
type servicePolicy  struct {
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
