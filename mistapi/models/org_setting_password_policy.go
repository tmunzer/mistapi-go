// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// OrgSettingPasswordPolicy represents a OrgSettingPasswordPolicy struct.
// password policy
type OrgSettingPasswordPolicy struct {
	// Whether the policy is enabled
	Enabled *bool `json:"enabled,omitempty"`
	// Password expiry in days. Password Expiry Notice banner will display in the UI 14 days before expiration
	ExpiryInDays *int `json:"expiry_in_days,omitempty"`
	// Required password length
	MinLength *int `json:"min_length,omitempty"`
	// Whether to require special character
	RequiresSpecialChar *bool `json:"requires_special_char,omitempty"`
	// Whether to require two-factor auth
	RequiresTwoFactorAuth *bool                  `json:"requires_two_factor_auth,omitempty"`
	AdditionalProperties  map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for OrgSettingPasswordPolicy,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (o OrgSettingPasswordPolicy) String() string {
	return fmt.Sprintf(
		"OrgSettingPasswordPolicy[Enabled=%v, ExpiryInDays=%v, MinLength=%v, RequiresSpecialChar=%v, RequiresTwoFactorAuth=%v, AdditionalProperties=%v]",
		o.Enabled, o.ExpiryInDays, o.MinLength, o.RequiresSpecialChar, o.RequiresTwoFactorAuth, o.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for OrgSettingPasswordPolicy.
// It customizes the JSON marshaling process for OrgSettingPasswordPolicy objects.
func (o OrgSettingPasswordPolicy) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(o.AdditionalProperties,
		"enabled", "expiry_in_days", "min_length", "requires_special_char", "requires_two_factor_auth"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(o.toMap())
}

// toMap converts the OrgSettingPasswordPolicy object to a map representation for JSON marshaling.
func (o OrgSettingPasswordPolicy) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, o.AdditionalProperties)
	if o.Enabled != nil {
		structMap["enabled"] = o.Enabled
	}
	if o.ExpiryInDays != nil {
		structMap["expiry_in_days"] = o.ExpiryInDays
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
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "enabled", "expiry_in_days", "min_length", "requires_special_char", "requires_two_factor_auth")
	if err != nil {
		return err
	}
	o.AdditionalProperties = additionalProperties

	o.Enabled = temp.Enabled
	o.ExpiryInDays = temp.ExpiryInDays
	o.MinLength = temp.MinLength
	o.RequiresSpecialChar = temp.RequiresSpecialChar
	o.RequiresTwoFactorAuth = temp.RequiresTwoFactorAuth
	return nil
}

// tempOrgSettingPasswordPolicy is a temporary struct used for validating the fields of OrgSettingPasswordPolicy.
type tempOrgSettingPasswordPolicy struct {
	Enabled               *bool `json:"enabled,omitempty"`
	ExpiryInDays          *int  `json:"expiry_in_days,omitempty"`
	MinLength             *int  `json:"min_length,omitempty"`
	RequiresSpecialChar   *bool `json:"requires_special_char,omitempty"`
	RequiresTwoFactorAuth *bool `json:"requires_two_factor_auth,omitempty"`
}
