// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// OrgSettingSecurity represents a OrgSettingSecurity struct.
type OrgSettingSecurity struct {
    // Whether to disable local SSH (by default, local SSH is enabled with allow_mist in Org is enabled
    DisableLocalSsh      *bool                  `json:"disable_local_ssh,omitempty"`
    // password required to zeroize devices (FIPS) on site level
    FipsZeroizePassword  *string                `json:"fips_zeroize_password,omitempty"`
    // Whether to allow certain SSH keys to SSH into the AP (see Site:Setting)
    LimitSshAccess       *bool                  `json:"limit_ssh_access,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for OrgSettingSecurity,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (o OrgSettingSecurity) String() string {
    return fmt.Sprintf(
    	"OrgSettingSecurity[DisableLocalSsh=%v, FipsZeroizePassword=%v, LimitSshAccess=%v, AdditionalProperties=%v]",
    	o.DisableLocalSsh, o.FipsZeroizePassword, o.LimitSshAccess, o.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for OrgSettingSecurity.
// It customizes the JSON marshaling process for OrgSettingSecurity objects.
func (o OrgSettingSecurity) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(o.AdditionalProperties,
        "disable_local_ssh", "fips_zeroize_password", "limit_ssh_access"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(o.toMap())
}

// toMap converts the OrgSettingSecurity object to a map representation for JSON marshaling.
func (o OrgSettingSecurity) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, o.AdditionalProperties)
    if o.DisableLocalSsh != nil {
        structMap["disable_local_ssh"] = o.DisableLocalSsh
    }
    if o.FipsZeroizePassword != nil {
        structMap["fips_zeroize_password"] = o.FipsZeroizePassword
    }
    if o.LimitSshAccess != nil {
        structMap["limit_ssh_access"] = o.LimitSshAccess
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgSettingSecurity.
// It customizes the JSON unmarshaling process for OrgSettingSecurity objects.
func (o *OrgSettingSecurity) UnmarshalJSON(input []byte) error {
    var temp tempOrgSettingSecurity
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "disable_local_ssh", "fips_zeroize_password", "limit_ssh_access")
    if err != nil {
    	return err
    }
    o.AdditionalProperties = additionalProperties
    
    o.DisableLocalSsh = temp.DisableLocalSsh
    o.FipsZeroizePassword = temp.FipsZeroizePassword
    o.LimitSshAccess = temp.LimitSshAccess
    return nil
}

// tempOrgSettingSecurity is a temporary struct used for validating the fields of OrgSettingSecurity.
type tempOrgSettingSecurity  struct {
    DisableLocalSsh     *bool   `json:"disable_local_ssh,omitempty"`
    FipsZeroizePassword *string `json:"fips_zeroize_password,omitempty"`
    LimitSshAccess      *bool   `json:"limit_ssh_access,omitempty"`
}
