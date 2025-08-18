// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// OrgSettingScep represents a OrgSettingScep struct.
type OrgSettingScep struct {
    // List of SCEP cert providers, e.g. `intune`, `jamf`, `byod`
    CertProviders        []OrgSettingScepCertProviderEnum `json:"cert_providers,omitempty"`
    // Whether SCEP is enabled for this org
    Enable               *bool                            `json:"enable,omitempty"`
    // Whether SCEP is suspended for this org
    Suspended            *bool                            `json:"suspended,omitempty"`
    AdditionalProperties map[string]interface{}           `json:"_"`
}

// String implements the fmt.Stringer interface for OrgSettingScep,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (o OrgSettingScep) String() string {
    return fmt.Sprintf(
    	"OrgSettingScep[CertProviders=%v, Enable=%v, Suspended=%v, AdditionalProperties=%v]",
    	o.CertProviders, o.Enable, o.Suspended, o.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for OrgSettingScep.
// It customizes the JSON marshaling process for OrgSettingScep objects.
func (o OrgSettingScep) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(o.AdditionalProperties,
        "cert_providers", "enable", "suspended"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(o.toMap())
}

// toMap converts the OrgSettingScep object to a map representation for JSON marshaling.
func (o OrgSettingScep) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, o.AdditionalProperties)
    if o.CertProviders != nil {
        structMap["cert_providers"] = o.CertProviders
    }
    if o.Enable != nil {
        structMap["enable"] = o.Enable
    }
    if o.Suspended != nil {
        structMap["suspended"] = o.Suspended
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgSettingScep.
// It customizes the JSON unmarshaling process for OrgSettingScep objects.
func (o *OrgSettingScep) UnmarshalJSON(input []byte) error {
    var temp tempOrgSettingScep
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "cert_providers", "enable", "suspended")
    if err != nil {
    	return err
    }
    o.AdditionalProperties = additionalProperties
    
    o.CertProviders = temp.CertProviders
    o.Enable = temp.Enable
    o.Suspended = temp.Suspended
    return nil
}

// tempOrgSettingScep is a temporary struct used for validating the fields of OrgSettingScep.
type tempOrgSettingScep  struct {
    CertProviders []OrgSettingScepCertProviderEnum `json:"cert_providers,omitempty"`
    Enable        *bool                            `json:"enable,omitempty"`
    Suspended     *bool                            `json:"suspended,omitempty"`
}
