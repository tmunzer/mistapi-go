package models

import (
    "encoding/json"
)

// AccountOauthInfoAccount represents a AccountOauthInfoAccount struct.
// OAuth linked apps account info
type AccountOauthInfoAccount struct {
    // Linked app(zoom/teams/intune) account id
    AccountId            *string                `json:"account_id,omitempty"`
    // customer account Client ID
    ClientId             *string                `json:"client_id,omitempty"`
    // Name of the company whose account mist has subscribed to
    Company              *string                `json:"company,omitempty"`
    // This error is provided when the account fails to fetch token/data
    Error                *string                `json:"error,omitempty"`
    Errors               []string               `json:"errors,omitempty"`
    // customer account instance URL
    InstanceUrl          *string                `json:"instance_url,omitempty"`
    // Is the last data pull for account is successful or not
    LastStatus           *string                `json:"last_status,omitempty"`
    // Last data pull timestamp, background jobs that pull account data
    LastSync             *int64                 `json:"last_sync,omitempty"`
    // First name of the user who linked the account
    LinkedBy             *string                `json:"linked_by,omitempty"`
    LinkedTimestamp      *float64               `json:"linked_timestamp,omitempty"`
    // Zoom daily api request quota, https://developers.zoom.us/docs/api/rest/rate-limits/
    MaxDailyApiRequests  *int                   `json:"max_daily_api_requests,omitempty"`
    // Name of the company whose account mist has subscribed to
    Name                 *string                `json:"name,omitempty"`
    // customer account password instance URL
    Password             *string                `json:"password,omitempty"`
    // smart group membership for determining compliance status
    SmartgroupName       *string                `json:"smartgroup_name,omitempty"`
    // customer account username
    Username             *string                `json:"username,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for AccountOauthInfoAccount.
// It customizes the JSON marshaling process for AccountOauthInfoAccount objects.
func (a AccountOauthInfoAccount) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(a.AdditionalProperties,
        "account_id", "client_id", "company", "error", "errors", "instance_url", "last_status", "last_sync", "linked_by", "linked_timestamp", "max_daily_api_requests", "name", "password", "smartgroup_name", "username"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(a.toMap())
}

// toMap converts the AccountOauthInfoAccount object to a map representation for JSON marshaling.
func (a AccountOauthInfoAccount) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, a.AdditionalProperties)
    if a.AccountId != nil {
        structMap["account_id"] = a.AccountId
    }
    if a.ClientId != nil {
        structMap["client_id"] = a.ClientId
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
    if a.InstanceUrl != nil {
        structMap["instance_url"] = a.InstanceUrl
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
    if a.LinkedTimestamp != nil {
        structMap["linked_timestamp"] = a.LinkedTimestamp
    }
    if a.MaxDailyApiRequests != nil {
        structMap["max_daily_api_requests"] = a.MaxDailyApiRequests
    }
    if a.Name != nil {
        structMap["name"] = a.Name
    }
    if a.Password != nil {
        structMap["password"] = a.Password
    }
    if a.SmartgroupName != nil {
        structMap["smartgroup_name"] = a.SmartgroupName
    }
    if a.Username != nil {
        structMap["username"] = a.Username
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AccountOauthInfoAccount.
// It customizes the JSON unmarshaling process for AccountOauthInfoAccount objects.
func (a *AccountOauthInfoAccount) UnmarshalJSON(input []byte) error {
    var temp tempAccountOauthInfoAccount
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "account_id", "client_id", "company", "error", "errors", "instance_url", "last_status", "last_sync", "linked_by", "linked_timestamp", "max_daily_api_requests", "name", "password", "smartgroup_name", "username")
    if err != nil {
    	return err
    }
    a.AdditionalProperties = additionalProperties
    
    a.AccountId = temp.AccountId
    a.ClientId = temp.ClientId
    a.Company = temp.Company
    a.Error = temp.Error
    a.Errors = temp.Errors
    a.InstanceUrl = temp.InstanceUrl
    a.LastStatus = temp.LastStatus
    a.LastSync = temp.LastSync
    a.LinkedBy = temp.LinkedBy
    a.LinkedTimestamp = temp.LinkedTimestamp
    a.MaxDailyApiRequests = temp.MaxDailyApiRequests
    a.Name = temp.Name
    a.Password = temp.Password
    a.SmartgroupName = temp.SmartgroupName
    a.Username = temp.Username
    return nil
}

// tempAccountOauthInfoAccount is a temporary struct used for validating the fields of AccountOauthInfoAccount.
type tempAccountOauthInfoAccount  struct {
    AccountId           *string  `json:"account_id,omitempty"`
    ClientId            *string  `json:"client_id,omitempty"`
    Company             *string  `json:"company,omitempty"`
    Error               *string  `json:"error,omitempty"`
    Errors              []string `json:"errors,omitempty"`
    InstanceUrl         *string  `json:"instance_url,omitempty"`
    LastStatus          *string  `json:"last_status,omitempty"`
    LastSync            *int64   `json:"last_sync,omitempty"`
    LinkedBy            *string  `json:"linked_by,omitempty"`
    LinkedTimestamp     *float64 `json:"linked_timestamp,omitempty"`
    MaxDailyApiRequests *int     `json:"max_daily_api_requests,omitempty"`
    Name                *string  `json:"name,omitempty"`
    Password            *string  `json:"password,omitempty"`
    SmartgroupName      *string  `json:"smartgroup_name,omitempty"`
    Username            *string  `json:"username,omitempty"`
}
