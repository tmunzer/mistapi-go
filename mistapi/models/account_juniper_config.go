// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// AccountJuniperConfig represents a AccountJuniperConfig struct.
type AccountJuniperConfig struct {
	// Customer account password
	Password string `json:"password"`
	// Customer account user name
	Username             string                 `json:"username"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for AccountJuniperConfig,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a AccountJuniperConfig) String() string {
	return fmt.Sprintf(
		"AccountJuniperConfig[Password=%v, Username=%v, AdditionalProperties=%v]",
		a.Password, a.Username, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for AccountJuniperConfig.
// It customizes the JSON marshaling process for AccountJuniperConfig objects.
func (a AccountJuniperConfig) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(a.AdditionalProperties,
		"password", "username"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(a.toMap())
}

// toMap converts the AccountJuniperConfig object to a map representation for JSON marshaling.
func (a AccountJuniperConfig) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, a.AdditionalProperties)
	structMap["password"] = a.Password
	structMap["username"] = a.Username
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AccountJuniperConfig.
// It customizes the JSON unmarshaling process for AccountJuniperConfig objects.
func (a *AccountJuniperConfig) UnmarshalJSON(input []byte) error {
	var temp tempAccountJuniperConfig
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "password", "username")
	if err != nil {
		return err
	}
	a.AdditionalProperties = additionalProperties

	a.Password = *temp.Password
	a.Username = *temp.Username
	return nil
}

// tempAccountJuniperConfig is a temporary struct used for validating the fields of AccountJuniperConfig.
type tempAccountJuniperConfig struct {
	Password *string `json:"password"`
	Username *string `json:"username"`
}

func (a *tempAccountJuniperConfig) validate() error {
	var errs []string
	if a.Password == nil {
		errs = append(errs, "required field `password` is missing for type `account_juniper_config`")
	}
	if a.Username == nil {
		errs = append(errs, "required field `username` is missing for type `account_juniper_config`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
