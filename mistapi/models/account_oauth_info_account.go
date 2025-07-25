// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// AccountOauthInfoAccount represents a AccountOauthInfoAccount struct.
// OAuth linked apps account info
type AccountOauthInfoAccount struct {
    // Linked app account id
    AccountId            *string                `json:"account_id,omitempty"`
    // For Prisma accounts only, tunnel auto probe subnet
    AutoProbeSubnet      *string                `json:"auto_probe_subnet,omitempty"`
    // Customer account Client ID
    ClientId             *string                `json:"client_id,omitempty"`
    // Name of the company whose account mist has subscribed to
    CloudName            *string                `json:"cloud_name,omitempty"`
    // Name of the company whose account mist has subscribed to
    Company              *string                `json:"company,omitempty"`
    // For Prisma accounts only, tunnel probe enable/disable
    EnableProbe          *bool                  `json:"enable_probe,omitempty"`
    // This error is provided when the account fails to fetch token/data
    Error                *string                `json:"error,omitempty"`
    Errors               []string               `json:"errors,omitempty"`
    // Customer account instance URL
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
    // Customer account password instance URL
    Password             *string                `json:"password,omitempty"`
    // For Prisma accounts only
    Region               *string                `json:"region,omitempty"`
    // For Prisma accounts only
    ServiceAccountName   *string                `json:"service_account_name,omitempty"`
    // For Prisma accounts only
    ServiceConnections   []string               `json:"service_connections,omitempty"`
    // Smart group membership for determining compliance status
    SmartgroupName       *string                `json:"smartgroup_name,omitempty"`
    // For Prisma accounts only, Prisma Tenant Service Group id
    TsgId                *string                `json:"tsg_id,omitempty"`
    // Customer account username
    Username             *string                `json:"username,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for AccountOauthInfoAccount,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a AccountOauthInfoAccount) String() string {
    return fmt.Sprintf(
    	"AccountOauthInfoAccount[AccountId=%v, AutoProbeSubnet=%v, ClientId=%v, CloudName=%v, Company=%v, EnableProbe=%v, Error=%v, Errors=%v, InstanceUrl=%v, LastStatus=%v, LastSync=%v, LinkedBy=%v, LinkedTimestamp=%v, MaxDailyApiRequests=%v, Name=%v, Password=%v, Region=%v, ServiceAccountName=%v, ServiceConnections=%v, SmartgroupName=%v, TsgId=%v, Username=%v, AdditionalProperties=%v]",
    	a.AccountId, a.AutoProbeSubnet, a.ClientId, a.CloudName, a.Company, a.EnableProbe, a.Error, a.Errors, a.InstanceUrl, a.LastStatus, a.LastSync, a.LinkedBy, a.LinkedTimestamp, a.MaxDailyApiRequests, a.Name, a.Password, a.Region, a.ServiceAccountName, a.ServiceConnections, a.SmartgroupName, a.TsgId, a.Username, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for AccountOauthInfoAccount.
// It customizes the JSON marshaling process for AccountOauthInfoAccount objects.
func (a AccountOauthInfoAccount) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(a.AdditionalProperties,
        "account_id", "auto_probe_subnet", "client_id", "cloud_name", "company", "enable_probe", "error", "errors", "instance_url", "last_status", "last_sync", "linked_by", "linked_timestamp", "max_daily_api_requests", "name", "password", "region", "service_account_name", "service_connections", "smartgroup_name", "tsg_id", "username"); err != nil {
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
    if a.AutoProbeSubnet != nil {
        structMap["auto_probe_subnet"] = a.AutoProbeSubnet
    }
    if a.ClientId != nil {
        structMap["client_id"] = a.ClientId
    }
    if a.CloudName != nil {
        structMap["cloud_name"] = a.CloudName
    }
    if a.Company != nil {
        structMap["company"] = a.Company
    }
    if a.EnableProbe != nil {
        structMap["enable_probe"] = a.EnableProbe
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
    if a.Region != nil {
        structMap["region"] = a.Region
    }
    if a.ServiceAccountName != nil {
        structMap["service_account_name"] = a.ServiceAccountName
    }
    if a.ServiceConnections != nil {
        structMap["service_connections"] = a.ServiceConnections
    }
    if a.SmartgroupName != nil {
        structMap["smartgroup_name"] = a.SmartgroupName
    }
    if a.TsgId != nil {
        structMap["tsg_id"] = a.TsgId
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
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "account_id", "auto_probe_subnet", "client_id", "cloud_name", "company", "enable_probe", "error", "errors", "instance_url", "last_status", "last_sync", "linked_by", "linked_timestamp", "max_daily_api_requests", "name", "password", "region", "service_account_name", "service_connections", "smartgroup_name", "tsg_id", "username")
    if err != nil {
    	return err
    }
    a.AdditionalProperties = additionalProperties
    
    a.AccountId = temp.AccountId
    a.AutoProbeSubnet = temp.AutoProbeSubnet
    a.ClientId = temp.ClientId
    a.CloudName = temp.CloudName
    a.Company = temp.Company
    a.EnableProbe = temp.EnableProbe
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
    a.Region = temp.Region
    a.ServiceAccountName = temp.ServiceAccountName
    a.ServiceConnections = temp.ServiceConnections
    a.SmartgroupName = temp.SmartgroupName
    a.TsgId = temp.TsgId
    a.Username = temp.Username
    return nil
}

// tempAccountOauthInfoAccount is a temporary struct used for validating the fields of AccountOauthInfoAccount.
type tempAccountOauthInfoAccount  struct {
    AccountId           *string  `json:"account_id,omitempty"`
    AutoProbeSubnet     *string  `json:"auto_probe_subnet,omitempty"`
    ClientId            *string  `json:"client_id,omitempty"`
    CloudName           *string  `json:"cloud_name,omitempty"`
    Company             *string  `json:"company,omitempty"`
    EnableProbe         *bool    `json:"enable_probe,omitempty"`
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
    Region              *string  `json:"region,omitempty"`
    ServiceAccountName  *string  `json:"service_account_name,omitempty"`
    ServiceConnections  []string `json:"service_connections,omitempty"`
    SmartgroupName      *string  `json:"smartgroup_name,omitempty"`
    TsgId               *string  `json:"tsg_id,omitempty"`
    Username            *string  `json:"username,omitempty"`
}
