// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// AccountZdxConfig represents a AccountZdxConfig struct.
// OAuth linked ZDX apps account details
type AccountZdxConfig struct {
    // ZDX cloud name. Refer https://help.zscaler.com/zdx/getting-started-zdx-api for ZDX cloud name
    CloudName            *string                `json:"cloud_name,omitempty"`
    // Customer account API key ID
    KeyId                string                 `json:"key_id"`
    // Customer account API key Secret
    KeySecret            string                 `json:"key_secret"`
    // ZDX organization id
    ZdxOrgId             string                 `json:"zdx_org_id"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for AccountZdxConfig,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a AccountZdxConfig) String() string {
    return fmt.Sprintf(
    	"AccountZdxConfig[CloudName=%v, KeyId=%v, KeySecret=%v, ZdxOrgId=%v, AdditionalProperties=%v]",
    	a.CloudName, a.KeyId, a.KeySecret, a.ZdxOrgId, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for AccountZdxConfig.
// It customizes the JSON marshaling process for AccountZdxConfig objects.
func (a AccountZdxConfig) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(a.AdditionalProperties,
        "cloud_name", "key_id", "key_secret", "zdx_org_id"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(a.toMap())
}

// toMap converts the AccountZdxConfig object to a map representation for JSON marshaling.
func (a AccountZdxConfig) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, a.AdditionalProperties)
    if a.CloudName != nil {
        structMap["cloud_name"] = a.CloudName
    }
    structMap["key_id"] = a.KeyId
    structMap["key_secret"] = a.KeySecret
    structMap["zdx_org_id"] = a.ZdxOrgId
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AccountZdxConfig.
// It customizes the JSON unmarshaling process for AccountZdxConfig objects.
func (a *AccountZdxConfig) UnmarshalJSON(input []byte) error {
    var temp tempAccountZdxConfig
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "cloud_name", "key_id", "key_secret", "zdx_org_id")
    if err != nil {
    	return err
    }
    a.AdditionalProperties = additionalProperties
    
    a.CloudName = temp.CloudName
    a.KeyId = *temp.KeyId
    a.KeySecret = *temp.KeySecret
    a.ZdxOrgId = *temp.ZdxOrgId
    return nil
}

// tempAccountZdxConfig is a temporary struct used for validating the fields of AccountZdxConfig.
type tempAccountZdxConfig  struct {
    CloudName *string `json:"cloud_name,omitempty"`
    KeyId     *string `json:"key_id"`
    KeySecret *string `json:"key_secret"`
    ZdxOrgId  *string `json:"zdx_org_id"`
}

func (a *tempAccountZdxConfig) validate() error {
    var errs []string
    if a.KeyId == nil {
        errs = append(errs, "required field `key_id` is missing for type `account_zdx_config`")
    }
    if a.KeySecret == nil {
        errs = append(errs, "required field `key_secret` is missing for type `account_zdx_config`")
    }
    if a.ZdxOrgId == nil {
        errs = append(errs, "required field `zdx_org_id` is missing for type `account_zdx_config`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
