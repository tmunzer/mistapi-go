// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// AccountSentineloneConfig represents a AccountSentineloneConfig struct.
// OAuth linked CrowdStrike apps account details
type AccountSentineloneConfig struct {
	// Customer account api_token
	ApiToken string `json:"api_token"`
	// Customer account SentinelOne instance URL
	InstanceUrl          string                 `json:"instance_url"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for AccountSentineloneConfig,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a AccountSentineloneConfig) String() string {
	return fmt.Sprintf(
		"AccountSentineloneConfig[ApiToken=%v, InstanceUrl=%v, AdditionalProperties=%v]",
		a.ApiToken, a.InstanceUrl, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for AccountSentineloneConfig.
// It customizes the JSON marshaling process for AccountSentineloneConfig objects.
func (a AccountSentineloneConfig) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(a.AdditionalProperties,
		"api_token", "instance_url"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(a.toMap())
}

// toMap converts the AccountSentineloneConfig object to a map representation for JSON marshaling.
func (a AccountSentineloneConfig) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, a.AdditionalProperties)
	structMap["api_token"] = a.ApiToken
	structMap["instance_url"] = a.InstanceUrl
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AccountSentineloneConfig.
// It customizes the JSON unmarshaling process for AccountSentineloneConfig objects.
func (a *AccountSentineloneConfig) UnmarshalJSON(input []byte) error {
	var temp tempAccountSentineloneConfig
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "api_token", "instance_url")
	if err != nil {
		return err
	}
	a.AdditionalProperties = additionalProperties

	a.ApiToken = *temp.ApiToken
	a.InstanceUrl = *temp.InstanceUrl
	return nil
}

// tempAccountSentineloneConfig is a temporary struct used for validating the fields of AccountSentineloneConfig.
type tempAccountSentineloneConfig struct {
	ApiToken    *string `json:"api_token"`
	InstanceUrl *string `json:"instance_url"`
}

func (a *tempAccountSentineloneConfig) validate() error {
	var errs []string
	if a.ApiToken == nil {
		errs = append(errs, "required field `api_token` is missing for type `account_sentinelone_config`")
	}
	if a.InstanceUrl == nil {
		errs = append(errs, "required field `instance_url` is missing for type `account_sentinelone_config`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
