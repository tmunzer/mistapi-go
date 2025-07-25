// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// AccountOauthConfig represents a AccountOauthConfig struct.
// OAuth linked apps (zoom/teams/intune) account details
type AccountOauthConfig struct {
    // Linked app(zoom/teams/intune) account id
    AccountId            string                 `json:"account_id"`
    // Zoom daily api request quota, https://developers.zoom.us/docs/api/rest/rate-limits/
    MaxDailyApiRequests  *int                   `json:"max_daily_api_requests,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for AccountOauthConfig,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a AccountOauthConfig) String() string {
    return fmt.Sprintf(
    	"AccountOauthConfig[AccountId=%v, MaxDailyApiRequests=%v, AdditionalProperties=%v]",
    	a.AccountId, a.MaxDailyApiRequests, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for AccountOauthConfig.
// It customizes the JSON marshaling process for AccountOauthConfig objects.
func (a AccountOauthConfig) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(a.AdditionalProperties,
        "account_id", "max_daily_api_requests"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(a.toMap())
}

// toMap converts the AccountOauthConfig object to a map representation for JSON marshaling.
func (a AccountOauthConfig) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, a.AdditionalProperties)
    structMap["account_id"] = a.AccountId
    if a.MaxDailyApiRequests != nil {
        structMap["max_daily_api_requests"] = a.MaxDailyApiRequests
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AccountOauthConfig.
// It customizes the JSON unmarshaling process for AccountOauthConfig objects.
func (a *AccountOauthConfig) UnmarshalJSON(input []byte) error {
    var temp tempAccountOauthConfig
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "account_id", "max_daily_api_requests")
    if err != nil {
    	return err
    }
    a.AdditionalProperties = additionalProperties
    
    a.AccountId = *temp.AccountId
    a.MaxDailyApiRequests = temp.MaxDailyApiRequests
    return nil
}

// tempAccountOauthConfig is a temporary struct used for validating the fields of AccountOauthConfig.
type tempAccountOauthConfig  struct {
    AccountId           *string `json:"account_id"`
    MaxDailyApiRequests *int    `json:"max_daily_api_requests,omitempty"`
}

func (a *tempAccountOauthConfig) validate() error {
    var errs []string
    if a.AccountId == nil {
        errs = append(errs, "required field `account_id` is missing for type `account_oauth_config`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
