package models

import (
    "encoding/json"
    "fmt"
    "github.com/google/uuid"
)

// ServicePolicy represents a ServicePolicy struct.
type ServicePolicy struct {
    // enum: `allow`, `deny`
    Action               *AllowDenyEnum         `json:"action,omitempty"`
    // For SRX Only
    Appqoe               *ServicePolicyAppqoe   `json:"appqoe,omitempty"`
    Ewf                  []ServicePolicyEwfRule `json:"ewf,omitempty"`
    Idp                  *IdpConfig             `json:"idp,omitempty"`
    // access within the same VRF
    LocalRouting         *bool                  `json:"local_routing,omitempty"`
    Name                 *string                `json:"name,omitempty"`
    // by default, we derive all paths available and use them. Optionally, you can customize by using `path_preference`
    PathPreference       *string                `json:"path_preference,omitempty"`
    // used to link servicepolicy defined at org level and overwrite some attributes
    ServicepolicyId      *uuid.UUID             `json:"servicepolicy_id,omitempty"`
    Services             []string               `json:"services,omitempty"`
    Tenants              []string               `json:"tenants,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ServicePolicy,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s ServicePolicy) String() string {
    return fmt.Sprintf(
    	"ServicePolicy[Action=%v, Appqoe=%v, Ewf=%v, Idp=%v, LocalRouting=%v, Name=%v, PathPreference=%v, ServicepolicyId=%v, Services=%v, Tenants=%v, AdditionalProperties=%v]",
    	s.Action, s.Appqoe, s.Ewf, s.Idp, s.LocalRouting, s.Name, s.PathPreference, s.ServicepolicyId, s.Services, s.Tenants, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ServicePolicy.
// It customizes the JSON marshaling process for ServicePolicy objects.
func (s ServicePolicy) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "action", "appqoe", "ewf", "idp", "local_routing", "name", "path_preference", "servicepolicy_id", "services", "tenants"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the ServicePolicy object to a map representation for JSON marshaling.
func (s ServicePolicy) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Action != nil {
        structMap["action"] = s.Action
    }
    if s.Appqoe != nil {
        structMap["appqoe"] = s.Appqoe.toMap()
    }
    if s.Ewf != nil {
        structMap["ewf"] = s.Ewf
    }
    if s.Idp != nil {
        structMap["idp"] = s.Idp.toMap()
    }
    if s.LocalRouting != nil {
        structMap["local_routing"] = s.LocalRouting
    }
    if s.Name != nil {
        structMap["name"] = s.Name
    }
    if s.PathPreference != nil {
        structMap["path_preference"] = s.PathPreference
    }
    if s.ServicepolicyId != nil {
        structMap["servicepolicy_id"] = s.ServicepolicyId
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
    var temp tempServicePolicy
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "action", "appqoe", "ewf", "idp", "local_routing", "name", "path_preference", "servicepolicy_id", "services", "tenants")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Action = temp.Action
    s.Appqoe = temp.Appqoe
    s.Ewf = temp.Ewf
    s.Idp = temp.Idp
    s.LocalRouting = temp.LocalRouting
    s.Name = temp.Name
    s.PathPreference = temp.PathPreference
    s.ServicepolicyId = temp.ServicepolicyId
    s.Services = temp.Services
    s.Tenants = temp.Tenants
    return nil
}

// tempServicePolicy is a temporary struct used for validating the fields of ServicePolicy.
type tempServicePolicy  struct {
    Action          *AllowDenyEnum         `json:"action,omitempty"`
    Appqoe          *ServicePolicyAppqoe   `json:"appqoe,omitempty"`
    Ewf             []ServicePolicyEwfRule `json:"ewf,omitempty"`
    Idp             *IdpConfig             `json:"idp,omitempty"`
    LocalRouting    *bool                  `json:"local_routing,omitempty"`
    Name            *string                `json:"name,omitempty"`
    PathPreference  *string                `json:"path_preference,omitempty"`
    ServicepolicyId *uuid.UUID             `json:"servicepolicy_id,omitempty"`
    Services        []string               `json:"services,omitempty"`
    Tenants         []string               `json:"tenants,omitempty"`
}
