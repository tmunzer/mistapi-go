// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
    "github.com/google/uuid"
)

// ServicePolicy represents a ServicePolicy struct.
type ServicePolicy struct {
    // enum: `allow`, `deny`
    Action               *AllowDenyEnum          `json:"action,omitempty"`
    // For SRX-only
    Antivirus            *ServicePolicyAntivirus `json:"antivirus,omitempty"`
    // For SRX Only
    Appqoe               *ServicePolicyAppqoe    `json:"appqoe,omitempty"`
    Ewf                  []ServicePolicyEwfRule  `json:"ewf,omitempty"`
    Idp                  *IdpConfig              `json:"idp,omitempty"`
    // access within the same VRF
    LocalRouting         *bool                   `json:"local_routing,omitempty"`
    Name                 *string                 `json:"name,omitempty"`
    // By default, we derive all paths available and use them. Optionally, you can customize by using `path_preference`
    PathPreference       *string                 `json:"path_preference,omitempty"`
    // For SRX Only
    Secintel             *ServicePolicySecintel  `json:"secintel,omitempty"`
    // Used to link servicepolicy defined at org level and overwrite some attributes
    ServicepolicyId      *uuid.UUID              `json:"servicepolicy_id,omitempty"`
    Services             []string                `json:"services,omitempty"`
    // For SRX-only
    SslProxy             *ServicePolicySslProxy  `json:"ssl_proxy,omitempty"`
    Tenants              []string                `json:"tenants,omitempty"`
    AdditionalProperties map[string]interface{}  `json:"_"`
}

// String implements the fmt.Stringer interface for ServicePolicy,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s ServicePolicy) String() string {
    return fmt.Sprintf(
    	"ServicePolicy[Action=%v, Antivirus=%v, Appqoe=%v, Ewf=%v, Idp=%v, LocalRouting=%v, Name=%v, PathPreference=%v, Secintel=%v, ServicepolicyId=%v, Services=%v, SslProxy=%v, Tenants=%v, AdditionalProperties=%v]",
    	s.Action, s.Antivirus, s.Appqoe, s.Ewf, s.Idp, s.LocalRouting, s.Name, s.PathPreference, s.Secintel, s.ServicepolicyId, s.Services, s.SslProxy, s.Tenants, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ServicePolicy.
// It customizes the JSON marshaling process for ServicePolicy objects.
func (s ServicePolicy) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "action", "antivirus", "appqoe", "ewf", "idp", "local_routing", "name", "path_preference", "secintel", "servicepolicy_id", "services", "ssl_proxy", "tenants"); err != nil {
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
    if s.Antivirus != nil {
        structMap["antivirus"] = s.Antivirus.toMap()
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
    if s.Secintel != nil {
        structMap["secintel"] = s.Secintel.toMap()
    }
    if s.ServicepolicyId != nil {
        structMap["servicepolicy_id"] = s.ServicepolicyId
    }
    if s.Services != nil {
        structMap["services"] = s.Services
    }
    if s.SslProxy != nil {
        structMap["ssl_proxy"] = s.SslProxy.toMap()
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
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "action", "antivirus", "appqoe", "ewf", "idp", "local_routing", "name", "path_preference", "secintel", "servicepolicy_id", "services", "ssl_proxy", "tenants")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Action = temp.Action
    s.Antivirus = temp.Antivirus
    s.Appqoe = temp.Appqoe
    s.Ewf = temp.Ewf
    s.Idp = temp.Idp
    s.LocalRouting = temp.LocalRouting
    s.Name = temp.Name
    s.PathPreference = temp.PathPreference
    s.Secintel = temp.Secintel
    s.ServicepolicyId = temp.ServicepolicyId
    s.Services = temp.Services
    s.SslProxy = temp.SslProxy
    s.Tenants = temp.Tenants
    return nil
}

// tempServicePolicy is a temporary struct used for validating the fields of ServicePolicy.
type tempServicePolicy  struct {
    Action          *AllowDenyEnum          `json:"action,omitempty"`
    Antivirus       *ServicePolicyAntivirus `json:"antivirus,omitempty"`
    Appqoe          *ServicePolicyAppqoe    `json:"appqoe,omitempty"`
    Ewf             []ServicePolicyEwfRule  `json:"ewf,omitempty"`
    Idp             *IdpConfig              `json:"idp,omitempty"`
    LocalRouting    *bool                   `json:"local_routing,omitempty"`
    Name            *string                 `json:"name,omitempty"`
    PathPreference  *string                 `json:"path_preference,omitempty"`
    Secintel        *ServicePolicySecintel  `json:"secintel,omitempty"`
    ServicepolicyId *uuid.UUID              `json:"servicepolicy_id,omitempty"`
    Services        []string                `json:"services,omitempty"`
    SslProxy        *ServicePolicySslProxy  `json:"ssl_proxy,omitempty"`
    Tenants         []string                `json:"tenants,omitempty"`
}
