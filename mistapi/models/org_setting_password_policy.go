package models

import (
    "encoding/json"
)

// OrgSettingPasswordPolicy represents a OrgSettingPasswordPolicy struct.
// password policy
type OrgSettingPasswordPolicy struct {
    // whether the policy is enabled
    Enabled               *bool          `json:"enabled,omitempty"`
    // days, required if password policy is enabled
    Freshness             *int           `json:"freshness,omitempty"`
    // required password length
    MinLength             *int           `json:"min_length,omitempty"`
    // whether to require special character
    RequiresSpecialChar   *bool          `json:"requires_special_char,omitempty"`
    // whether to require two-factor auth
    RequiresTwoFactorAuth *bool          `json:"requires_two_factor_auth,omitempty"`
    AdditionalProperties  map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for OrgSettingPasswordPolicy.
// It customizes the JSON marshaling process for OrgSettingPasswordPolicy objects.
func (o OrgSettingPasswordPolicy) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(o.toMap())
}

// toMap converts the OrgSettingPasswordPolicy object to a map representation for JSON marshaling.
func (o OrgSettingPasswordPolicy) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, o.AdditionalProperties)
    if o.Enabled != nil {
        structMap["enabled"] = o.Enabled
    }
    if o.Freshness != nil {
        structMap["freshness"] = o.Freshness
    }
    if o.MinLength != nil {
        structMap["min_length"] = o.MinLength
    }
    if o.RequiresSpecialChar != nil {
        structMap["requires_special_char"] = o.RequiresSpecialChar
    }
    if o.RequiresTwoFactorAuth != nil {
        structMap["requires_two_factor_auth"] = o.RequiresTwoFactorAuth
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgSettingPasswordPolicy.
// It customizes the JSON unmarshaling process for OrgSettingPasswordPolicy objects.
func (o *OrgSettingPasswordPolicy) UnmarshalJSON(input []byte) error {
    var temp tempOrgSettingPasswordPolicy
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "enabled", "freshness", "min_length", "requires_special_char", "requires_two_factor_auth")
    if err != nil {
    	return err
    }
    
    o.AdditionalProperties = additionalProperties
    o.Enabled = temp.Enabled
    o.Freshness = temp.Freshness
    o.MinLength = temp.MinLength
    o.RequiresSpecialChar = temp.RequiresSpecialChar
    o.RequiresTwoFactorAuth = temp.RequiresTwoFactorAuth
    return nil
}

// tempOrgSettingPasswordPolicy is a temporary struct used for validating the fields of OrgSettingPasswordPolicy.
type tempOrgSettingPasswordPolicy  struct {
    Enabled               *bool `json:"enabled,omitempty"`
    Freshness             *int  `json:"freshness,omitempty"`
    MinLength             *int  `json:"min_length,omitempty"`
    RequiresSpecialChar   *bool `json:"requires_special_char,omitempty"`
    RequiresTwoFactorAuth *bool `json:"requires_two_factor_auth,omitempty"`
}
