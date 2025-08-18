// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
)

// OrgServicePolicy represents a OrgServicePolicy struct.
type OrgServicePolicy struct {
	// For SRX Only
	Aamw *ServicePolicyAamw `json:"aamw,omitempty"`
	// enum: `allow`, `deny`
	Action *AllowDenyEnum `json:"action,omitempty"`
	// For SRX-only
	Antivirus *ServicePolicyAntivirus `json:"antivirus,omitempty"`
	// For SRX Only
	Appqoe *ServicePolicyAppqoe `json:"appqoe,omitempty"`
	// When the object has been created, in epoch
	CreatedTime *float64               `json:"created_time,omitempty"`
	Ewf         []ServicePolicyEwfRule `json:"ewf,omitempty"`
	// Unique ID of the object instance in the Mist Organization
	Id  *uuid.UUID `json:"id,omitempty"`
	Idp *IdpConfig `json:"idp,omitempty"`
	// access within the same VRF
	LocalRouting *bool `json:"local_routing,omitempty"`
	// When the object has been modified for the last time, in epoch
	ModifiedTime *float64   `json:"modified_time,omitempty"`
	Name         *string    `json:"name,omitempty"`
	OrgId        *uuid.UUID `json:"org_id,omitempty"`
	// By default, we derive all paths available and use them, optionally, you can customize by using `path_preference`
	PathPreference *string `json:"path_preference,omitempty"`
	// For SRX Only
	Secintel *ServicePolicySecintel `json:"secintel,omitempty"`
	Services []string               `json:"services,omitempty"`
	// For SRX-only
	SslProxy             *ServicePolicySslProxy `json:"ssl_proxy,omitempty"`
	Tenants              []string               `json:"tenants,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for OrgServicePolicy,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (o OrgServicePolicy) String() string {
	return fmt.Sprintf(
		"OrgServicePolicy[Aamw=%v, Action=%v, Antivirus=%v, Appqoe=%v, CreatedTime=%v, Ewf=%v, Id=%v, Idp=%v, LocalRouting=%v, ModifiedTime=%v, Name=%v, OrgId=%v, PathPreference=%v, Secintel=%v, Services=%v, SslProxy=%v, Tenants=%v, AdditionalProperties=%v]",
		o.Aamw, o.Action, o.Antivirus, o.Appqoe, o.CreatedTime, o.Ewf, o.Id, o.Idp, o.LocalRouting, o.ModifiedTime, o.Name, o.OrgId, o.PathPreference, o.Secintel, o.Services, o.SslProxy, o.Tenants, o.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for OrgServicePolicy.
// It customizes the JSON marshaling process for OrgServicePolicy objects.
func (o OrgServicePolicy) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(o.AdditionalProperties,
		"aamw", "action", "antivirus", "appqoe", "created_time", "ewf", "id", "idp", "local_routing", "modified_time", "name", "org_id", "path_preference", "secintel", "services", "ssl_proxy", "tenants"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(o.toMap())
}

// toMap converts the OrgServicePolicy object to a map representation for JSON marshaling.
func (o OrgServicePolicy) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, o.AdditionalProperties)
	if o.Aamw != nil {
		structMap["aamw"] = o.Aamw.toMap()
	}
	if o.Action != nil {
		structMap["action"] = o.Action
	}
	if o.Antivirus != nil {
		structMap["antivirus"] = o.Antivirus.toMap()
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
	if o.PathPreference != nil {
		structMap["path_preference"] = o.PathPreference
	}
	if o.Secintel != nil {
		structMap["secintel"] = o.Secintel.toMap()
	}
	if o.Services != nil {
		structMap["services"] = o.Services
	}
	if o.SslProxy != nil {
		structMap["ssl_proxy"] = o.SslProxy.toMap()
	}
	if o.Tenants != nil {
		structMap["tenants"] = o.Tenants
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgServicePolicy.
// It customizes the JSON unmarshaling process for OrgServicePolicy objects.
func (o *OrgServicePolicy) UnmarshalJSON(input []byte) error {
	var temp tempOrgServicePolicy
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "aamw", "action", "antivirus", "appqoe", "created_time", "ewf", "id", "idp", "local_routing", "modified_time", "name", "org_id", "path_preference", "secintel", "services", "ssl_proxy", "tenants")
	if err != nil {
		return err
	}
	o.AdditionalProperties = additionalProperties

	o.Aamw = temp.Aamw
	o.Action = temp.Action
	o.Antivirus = temp.Antivirus
	o.Appqoe = temp.Appqoe
	o.CreatedTime = temp.CreatedTime
	o.Ewf = temp.Ewf
	o.Id = temp.Id
	o.Idp = temp.Idp
	o.LocalRouting = temp.LocalRouting
	o.ModifiedTime = temp.ModifiedTime
	o.Name = temp.Name
	o.OrgId = temp.OrgId
	o.PathPreference = temp.PathPreference
	o.Secintel = temp.Secintel
	o.Services = temp.Services
	o.SslProxy = temp.SslProxy
	o.Tenants = temp.Tenants
	return nil
}

// tempOrgServicePolicy is a temporary struct used for validating the fields of OrgServicePolicy.
type tempOrgServicePolicy struct {
	Aamw           *ServicePolicyAamw      `json:"aamw,omitempty"`
	Action         *AllowDenyEnum          `json:"action,omitempty"`
	Antivirus      *ServicePolicyAntivirus `json:"antivirus,omitempty"`
	Appqoe         *ServicePolicyAppqoe    `json:"appqoe,omitempty"`
	CreatedTime    *float64                `json:"created_time,omitempty"`
	Ewf            []ServicePolicyEwfRule  `json:"ewf,omitempty"`
	Id             *uuid.UUID              `json:"id,omitempty"`
	Idp            *IdpConfig              `json:"idp,omitempty"`
	LocalRouting   *bool                   `json:"local_routing,omitempty"`
	ModifiedTime   *float64                `json:"modified_time,omitempty"`
	Name           *string                 `json:"name,omitempty"`
	OrgId          *uuid.UUID              `json:"org_id,omitempty"`
	PathPreference *string                 `json:"path_preference,omitempty"`
	Secintel       *ServicePolicySecintel  `json:"secintel,omitempty"`
	Services       []string                `json:"services,omitempty"`
	SslProxy       *ServicePolicySslProxy  `json:"ssl_proxy,omitempty"`
	Tenants        []string                `json:"tenants,omitempty"`
}
