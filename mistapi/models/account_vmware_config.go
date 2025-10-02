// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// AccountVmwareConfig represents a AccountVmwareConfig struct.
type AccountVmwareConfig struct {
	// Customer account Client ID
	ClientId string `json:"client_id"`
	// Customer account Client Secret
	ClientSecret string `json:"client_secret"`
	// Customer account VMware instance URL
	InstanceUrl string `json:"instance_url"`
	// Enables or disables the webhook integration
	WebhookEnabled       bool                   `json:"webhook_enabled"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for AccountVmwareConfig,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a AccountVmwareConfig) String() string {
	return fmt.Sprintf(
		"AccountVmwareConfig[ClientId=%v, ClientSecret=%v, InstanceUrl=%v, WebhookEnabled=%v, AdditionalProperties=%v]",
		a.ClientId, a.ClientSecret, a.InstanceUrl, a.WebhookEnabled, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for AccountVmwareConfig.
// It customizes the JSON marshaling process for AccountVmwareConfig objects.
func (a AccountVmwareConfig) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(a.AdditionalProperties,
		"client_id", "client_secret", "instance_url", "webhook_enabled"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(a.toMap())
}

// toMap converts the AccountVmwareConfig object to a map representation for JSON marshaling.
func (a AccountVmwareConfig) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, a.AdditionalProperties)
	structMap["client_id"] = a.ClientId
	structMap["client_secret"] = a.ClientSecret
	structMap["instance_url"] = a.InstanceUrl
	structMap["webhook_enabled"] = a.WebhookEnabled
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AccountVmwareConfig.
// It customizes the JSON unmarshaling process for AccountVmwareConfig objects.
func (a *AccountVmwareConfig) UnmarshalJSON(input []byte) error {
	var temp tempAccountVmwareConfig
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "client_id", "client_secret", "instance_url", "webhook_enabled")
	if err != nil {
		return err
	}
	a.AdditionalProperties = additionalProperties

	a.ClientId = *temp.ClientId
	a.ClientSecret = *temp.ClientSecret
	a.InstanceUrl = *temp.InstanceUrl
	a.WebhookEnabled = *temp.WebhookEnabled
	return nil
}

// tempAccountVmwareConfig is a temporary struct used for validating the fields of AccountVmwareConfig.
type tempAccountVmwareConfig struct {
	ClientId       *string `json:"client_id"`
	ClientSecret   *string `json:"client_secret"`
	InstanceUrl    *string `json:"instance_url"`
	WebhookEnabled *bool   `json:"webhook_enabled"`
}

func (a *tempAccountVmwareConfig) validate() error {
	var errs []string
	if a.ClientId == nil {
		errs = append(errs, "required field `client_id` is missing for type `account_vmware_config`")
	}
	if a.ClientSecret == nil {
		errs = append(errs, "required field `client_secret` is missing for type `account_vmware_config`")
	}
	if a.InstanceUrl == nil {
		errs = append(errs, "required field `instance_url` is missing for type `account_vmware_config`")
	}
	if a.WebhookEnabled == nil {
		errs = append(errs, "required field `webhook_enabled` is missing for type `account_vmware_config`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
