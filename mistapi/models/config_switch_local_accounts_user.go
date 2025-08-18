// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// ConfigSwitchLocalAccountsUser represents a ConfigSwitchLocalAccountsUser struct.
type ConfigSwitchLocalAccountsUser struct {
	Password *string `json:"password,omitempty"`
	// enum: `admin`, `helpdesk`, `none`, `read`
	Role                 *ConfigSwitchLocalAccountsUserRoleEnum `json:"role,omitempty"`
	AdditionalProperties map[string]interface{}                 `json:"_"`
}

// String implements the fmt.Stringer interface for ConfigSwitchLocalAccountsUser,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c ConfigSwitchLocalAccountsUser) String() string {
	return fmt.Sprintf(
		"ConfigSwitchLocalAccountsUser[Password=%v, Role=%v, AdditionalProperties=%v]",
		c.Password, c.Role, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ConfigSwitchLocalAccountsUser.
// It customizes the JSON marshaling process for ConfigSwitchLocalAccountsUser objects.
func (c ConfigSwitchLocalAccountsUser) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(c.AdditionalProperties,
		"password", "role"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(c.toMap())
}

// toMap converts the ConfigSwitchLocalAccountsUser object to a map representation for JSON marshaling.
func (c ConfigSwitchLocalAccountsUser) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, c.AdditionalProperties)
	if c.Password != nil {
		structMap["password"] = c.Password
	}
	if c.Role != nil {
		structMap["role"] = c.Role
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ConfigSwitchLocalAccountsUser.
// It customizes the JSON unmarshaling process for ConfigSwitchLocalAccountsUser objects.
func (c *ConfigSwitchLocalAccountsUser) UnmarshalJSON(input []byte) error {
	var temp tempConfigSwitchLocalAccountsUser
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "password", "role")
	if err != nil {
		return err
	}
	c.AdditionalProperties = additionalProperties

	c.Password = temp.Password
	c.Role = temp.Role
	return nil
}

// tempConfigSwitchLocalAccountsUser is a temporary struct used for validating the fields of ConfigSwitchLocalAccountsUser.
type tempConfigSwitchLocalAccountsUser struct {
	Password *string                                `json:"password,omitempty"`
	Role     *ConfigSwitchLocalAccountsUserRoleEnum `json:"role,omitempty"`
}
