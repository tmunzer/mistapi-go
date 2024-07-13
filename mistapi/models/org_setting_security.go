package models

import (
    "encoding/json"
)

// OrgSettingSecurity represents a OrgSettingSecurity struct.
type OrgSettingSecurity struct {
    // whether to disable local SSH (by default, local SSH is enabled with allow_mist in Org is enabled
    DisableLocalSsh      *bool          `json:"disable_local_ssh,omitempty"`
    // password required to zeroize devices (FIPS) on site level
    FipsZeroizePassword  *string        `json:"fips_zeroize_password,omitempty"`
    // whether to allow certain SSH keys to SSH into the AP (see Site:Setting)
    LimitSshAccess       *bool          `json:"limit_ssh_access,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for OrgSettingSecurity.
// It customizes the JSON marshaling process for OrgSettingSecurity objects.
func (o OrgSettingSecurity) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(o.toMap())
}

// toMap converts the OrgSettingSecurity object to a map representation for JSON marshaling.
func (o OrgSettingSecurity) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, o.AdditionalProperties)
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
    var temp orgSettingSecurity
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "disable_local_ssh", "fips_zeroize_password", "limit_ssh_access")
    if err != nil {
    	return err
    }
    
    o.AdditionalProperties = additionalProperties
    o.DisableLocalSsh = temp.DisableLocalSsh
    o.FipsZeroizePassword = temp.FipsZeroizePassword
    o.LimitSshAccess = temp.LimitSshAccess
    return nil
}

// orgSettingSecurity is a temporary struct used for validating the fields of OrgSettingSecurity.
type orgSettingSecurity  struct {
    DisableLocalSsh     *bool   `json:"disable_local_ssh,omitempty"`
    FipsZeroizePassword *string `json:"fips_zeroize_password,omitempty"`
    LimitSshAccess      *bool   `json:"limit_ssh_access,omitempty"`
}
