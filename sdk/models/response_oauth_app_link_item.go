package models

import (
    "encoding/json"
)

// ResponseOauthAppLinkItem represents a ResponseOauthAppLinkItem struct.
type ResponseOauthAppLinkItem struct {
    // customer account client id
    ClientId             *string        `json:"client_id,omitempty"`
    // This error is provided when the Jamf account fails to fetch token/data
    Error                *string        `json:"error,omitempty"`
    // customer account Jamf instance URL
    InstanceUrl          *string        `json:"instance_url,omitempty"`
    // Is the last data pull for Jamf account is successful or not
    LastStatus           *string        `json:"last_status,omitempty"`
    // Last data pull timestamp, background jobs that pull Jamf account data
    LastSync             *int64         `json:"last_sync,omitempty"`
    // First name of the user who linked the Jamf account
    LinkedBy             *string        `json:"linked_by,omitempty"`
    // Name of the company whose Jamf account mist has subscribed to
    Name                 *string        `json:"name,omitempty"`
    // smart group membership for determining compliance status
    SmartgroupName       *string        `json:"smartgroup_name,omitempty"`
    // Linked app(zoom/teams/intune) account id
    AccountId            *string        `json:"account_id,omitempty"`
    // Name of the company whose account mist has subscribed to
    Company              *string        `json:"company,omitempty"`
    Errors               []string       `json:"errors,omitempty"`
    // Zoom daily api request quota, https://developers.zoom.us/docs/api/rest/rate-limits/
    MaxDailyApiRequests  *int           `json:"max_daily_api_requests,omitempty"`
    // This error is provided when the VMware account fails to fetch token/data
    LinkedTimestamp      *int           `json:"linked_timestamp,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseOauthAppLinkItem.
// It customizes the JSON marshaling process for ResponseOauthAppLinkItem objects.
func (r ResponseOauthAppLinkItem) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseOauthAppLinkItem object to a map representation for JSON marshaling.
func (r ResponseOauthAppLinkItem) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    if r.ClientId != nil {
        structMap["client_id"] = r.ClientId
    }
    if r.Error != nil {
        structMap["error"] = r.Error
    }
    if r.InstanceUrl != nil {
        structMap["instance_url"] = r.InstanceUrl
    }
    if r.LastStatus != nil {
        structMap["last_status"] = r.LastStatus
    }
    if r.LastSync != nil {
        structMap["last_sync"] = r.LastSync
    }
    if r.LinkedBy != nil {
        structMap["linked_by"] = r.LinkedBy
    }
    if r.Name != nil {
        structMap["name"] = r.Name
    }
    if r.SmartgroupName != nil {
        structMap["smartgroup_name"] = r.SmartgroupName
    }
    if r.AccountId != nil {
        structMap["account_id"] = r.AccountId
    }
    if r.Company != nil {
        structMap["company"] = r.Company
    }
    if r.Errors != nil {
        structMap["errors"] = r.Errors
    }
    if r.MaxDailyApiRequests != nil {
        structMap["max_daily_api_requests"] = r.MaxDailyApiRequests
    }
    if r.LinkedTimestamp != nil {
        structMap["linked_timestamp"] = r.LinkedTimestamp
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseOauthAppLinkItem.
// It customizes the JSON unmarshaling process for ResponseOauthAppLinkItem objects.
func (r *ResponseOauthAppLinkItem) UnmarshalJSON(input []byte) error {
    var temp responseOauthAppLinkItem
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "client_id", "error", "instance_url", "last_status", "last_sync", "linked_by", "name", "smartgroup_name", "account_id", "company", "errors", "max_daily_api_requests", "linked_timestamp")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.ClientId = temp.ClientId
    r.Error = temp.Error
    r.InstanceUrl = temp.InstanceUrl
    r.LastStatus = temp.LastStatus
    r.LastSync = temp.LastSync
    r.LinkedBy = temp.LinkedBy
    r.Name = temp.Name
    r.SmartgroupName = temp.SmartgroupName
    r.AccountId = temp.AccountId
    r.Company = temp.Company
    r.Errors = temp.Errors
    r.MaxDailyApiRequests = temp.MaxDailyApiRequests
    r.LinkedTimestamp = temp.LinkedTimestamp
    return nil
}

// responseOauthAppLinkItem is a temporary struct used for validating the fields of ResponseOauthAppLinkItem.
type responseOauthAppLinkItem  struct {
    ClientId            *string  `json:"client_id,omitempty"`
    Error               *string  `json:"error,omitempty"`
    InstanceUrl         *string  `json:"instance_url,omitempty"`
    LastStatus          *string  `json:"last_status,omitempty"`
    LastSync            *int64   `json:"last_sync,omitempty"`
    LinkedBy            *string  `json:"linked_by,omitempty"`
    Name                *string  `json:"name,omitempty"`
    SmartgroupName      *string  `json:"smartgroup_name,omitempty"`
    AccountId           *string  `json:"account_id,omitempty"`
    Company             *string  `json:"company,omitempty"`
    Errors              []string `json:"errors,omitempty"`
    MaxDailyApiRequests *int     `json:"max_daily_api_requests,omitempty"`
    LinkedTimestamp     *int     `json:"linked_timestamp,omitempty"`
}
