package models

import (
    "encoding/json"
)

// ConfigSwitchLocalAccountsUser represents a ConfigSwitchLocalAccountsUser struct.
type ConfigSwitchLocalAccountsUser struct {
    Password             *string                                `json:"password,omitempty"`
    // enum: `admin`, `helpdesk`, `none`, `read`
    Role                 *ConfigSwitchLocalAccountsUserRoleEnum `json:"role,omitempty"`
    AdditionalProperties map[string]any                         `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ConfigSwitchLocalAccountsUser.
// It customizes the JSON marshaling process for ConfigSwitchLocalAccountsUser objects.
func (c ConfigSwitchLocalAccountsUser) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the ConfigSwitchLocalAccountsUser object to a map representation for JSON marshaling.
func (c ConfigSwitchLocalAccountsUser) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
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
    var temp configSwitchLocalAccountsUser
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "password", "role")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.Password = temp.Password
    c.Role = temp.Role
    return nil
}

// configSwitchLocalAccountsUser is a temporary struct used for validating the fields of ConfigSwitchLocalAccountsUser.
type configSwitchLocalAccountsUser  struct {
    Password *string                                `json:"password,omitempty"`
    Role     *ConfigSwitchLocalAccountsUserRoleEnum `json:"role,omitempty"`
}
