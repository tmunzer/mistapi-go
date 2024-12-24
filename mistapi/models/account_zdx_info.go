package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// AccountZdxInfo represents a AccountZdxInfo struct.
type AccountZdxInfo struct {
    // Generated account id (uuid)
    AccountId            string                 `json:"account_id"`
    // ZDX cloud name
    CloudName            string                 `json:"cloud_name"`
    // Customer account API key ID
    KeyId                string                 `json:"key_id"`
    // Bearer token for the webhook url
    WebhookToken         string                 `json:"webhook_token"`
    // Webhook url to notify Mist about a ZDX alert
    WebhookUrl           string                 `json:"webhook_url"`
    // ZDX organization id
    ZdxOrgId             string                 `json:"zdx_org_id"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for AccountZdxInfo,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a AccountZdxInfo) String() string {
    return fmt.Sprintf(
    	"AccountZdxInfo[AccountId=%v, CloudName=%v, KeyId=%v, WebhookToken=%v, WebhookUrl=%v, ZdxOrgId=%v, AdditionalProperties=%v]",
    	a.AccountId, a.CloudName, a.KeyId, a.WebhookToken, a.WebhookUrl, a.ZdxOrgId, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for AccountZdxInfo.
// It customizes the JSON marshaling process for AccountZdxInfo objects.
func (a AccountZdxInfo) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(a.AdditionalProperties,
        "account_id", "cloud_name", "key_id", "webhook_token", "webhook_url", "zdx_org_id"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(a.toMap())
}

// toMap converts the AccountZdxInfo object to a map representation for JSON marshaling.
func (a AccountZdxInfo) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, a.AdditionalProperties)
    structMap["account_id"] = a.AccountId
    structMap["cloud_name"] = a.CloudName
    structMap["key_id"] = a.KeyId
    structMap["webhook_token"] = a.WebhookToken
    structMap["webhook_url"] = a.WebhookUrl
    structMap["zdx_org_id"] = a.ZdxOrgId
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AccountZdxInfo.
// It customizes the JSON unmarshaling process for AccountZdxInfo objects.
func (a *AccountZdxInfo) UnmarshalJSON(input []byte) error {
    var temp tempAccountZdxInfo
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "account_id", "cloud_name", "key_id", "webhook_token", "webhook_url", "zdx_org_id")
    if err != nil {
    	return err
    }
    a.AdditionalProperties = additionalProperties
    
    a.AccountId = *temp.AccountId
    a.CloudName = *temp.CloudName
    a.KeyId = *temp.KeyId
    a.WebhookToken = *temp.WebhookToken
    a.WebhookUrl = *temp.WebhookUrl
    a.ZdxOrgId = *temp.ZdxOrgId
    return nil
}

// tempAccountZdxInfo is a temporary struct used for validating the fields of AccountZdxInfo.
type tempAccountZdxInfo  struct {
    AccountId    *string `json:"account_id"`
    CloudName    *string `json:"cloud_name"`
    KeyId        *string `json:"key_id"`
    WebhookToken *string `json:"webhook_token"`
    WebhookUrl   *string `json:"webhook_url"`
    ZdxOrgId     *string `json:"zdx_org_id"`
}

func (a *tempAccountZdxInfo) validate() error {
    var errs []string
    if a.AccountId == nil {
        errs = append(errs, "required field `account_id` is missing for type `account_zdx_info`")
    }
    if a.CloudName == nil {
        errs = append(errs, "required field `cloud_name` is missing for type `account_zdx_info`")
    }
    if a.KeyId == nil {
        errs = append(errs, "required field `key_id` is missing for type `account_zdx_info`")
    }
    if a.WebhookToken == nil {
        errs = append(errs, "required field `webhook_token` is missing for type `account_zdx_info`")
    }
    if a.WebhookUrl == nil {
        errs = append(errs, "required field `webhook_url` is missing for type `account_zdx_info`")
    }
    if a.ZdxOrgId == nil {
        errs = append(errs, "required field `zdx_org_id` is missing for type `account_zdx_info`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
