package models

import (
	"encoding/json"
)

// AccountOauthInfo represents a AccountOauthInfo struct.
// OAuth linked apps (zoom/teams/intune) account details
type AccountOauthInfo struct {
	// Linked app(zoom/teams/intune) account id
	AccountId *string `json:"account_id,omitempty"`
	// Name of the company whose account mist has subscribed to
	Company *string `json:"company,omitempty"`
	// This error is provided when the account fails to fetch token/data
	Error  *string  `json:"error,omitempty"`
	Errors []string `json:"errors,omitempty"`
	// Is the last data pull for account is successful or not
	LastStatus *string `json:"last_status,omitempty"`
	// Last data pull timestamp, background jobs that pull account data
	LastSync *int64 `json:"last_sync,omitempty"`
	// First name of the user who linked the account
	LinkedBy *string `json:"linked_by,omitempty"`
	// Zoom daily api request quota, https://developers.zoom.us/docs/api/rest/rate-limits/
	MaxDailyApiRequests  *int           `json:"max_daily_api_requests,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for AccountOauthInfo.
// It customizes the JSON marshaling process for AccountOauthInfo objects.
func (a AccountOauthInfo) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(a.toMap())
}

// toMap converts the AccountOauthInfo object to a map representation for JSON marshaling.
func (a AccountOauthInfo) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, a.AdditionalProperties)
	if a.AccountId != nil {
		structMap["account_id"] = a.AccountId
	}
	if a.Company != nil {
		structMap["company"] = a.Company
	}
	if a.Error != nil {
		structMap["error"] = a.Error
	}
	if a.Errors != nil {
		structMap["errors"] = a.Errors
	}
	if a.LastStatus != nil {
		structMap["last_status"] = a.LastStatus
	}
	if a.LastSync != nil {
		structMap["last_sync"] = a.LastSync
	}
	if a.LinkedBy != nil {
		structMap["linked_by"] = a.LinkedBy
	}
	if a.MaxDailyApiRequests != nil {
		structMap["max_daily_api_requests"] = a.MaxDailyApiRequests
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AccountOauthInfo.
// It customizes the JSON unmarshaling process for AccountOauthInfo objects.
func (a *AccountOauthInfo) UnmarshalJSON(input []byte) error {
	var temp tempAccountOauthInfo
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "account_id", "company", "error", "errors", "last_status", "last_sync", "linked_by", "max_daily_api_requests")
	if err != nil {
		return err
	}

	a.AdditionalProperties = additionalProperties
	a.AccountId = temp.AccountId
	a.Company = temp.Company
	a.Error = temp.Error
	a.Errors = temp.Errors
	a.LastStatus = temp.LastStatus
	a.LastSync = temp.LastSync
	a.LinkedBy = temp.LinkedBy
	a.MaxDailyApiRequests = temp.MaxDailyApiRequests
	return nil
}

// tempAccountOauthInfo is a temporary struct used for validating the fields of AccountOauthInfo.
type tempAccountOauthInfo struct {
	AccountId           *string  `json:"account_id,omitempty"`
	Company             *string  `json:"company,omitempty"`
	Error               *string  `json:"error,omitempty"`
	Errors              []string `json:"errors,omitempty"`
	LastStatus          *string  `json:"last_status,omitempty"`
	LastSync            *int64   `json:"last_sync,omitempty"`
	LinkedBy            *string  `json:"linked_by,omitempty"`
	MaxDailyApiRequests *int     `json:"max_daily_api_requests,omitempty"`
}
