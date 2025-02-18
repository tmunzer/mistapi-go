package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// AccountOauthInfo represents a AccountOauthInfo struct.
type AccountOauthInfo struct {
    // List of linked account details
    Accounts             []AccountOauthInfoAccount `json:"accounts"`
    AuthorizationUrl     *string                   `json:"authorization_url,omitempty"`
    Linked               bool                      `json:"linked"`
    AdditionalProperties map[string]interface{}    `json:"_"`
}

// String implements the fmt.Stringer interface for AccountOauthInfo,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a AccountOauthInfo) String() string {
    return fmt.Sprintf(
    	"AccountOauthInfo[Accounts=%v, AuthorizationUrl=%v, Linked=%v, AdditionalProperties=%v]",
    	a.Accounts, a.AuthorizationUrl, a.Linked, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for AccountOauthInfo.
// It customizes the JSON marshaling process for AccountOauthInfo objects.
func (a AccountOauthInfo) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(a.AdditionalProperties,
        "accounts", "authorization_url", "linked"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(a.toMap())
}

// toMap converts the AccountOauthInfo object to a map representation for JSON marshaling.
func (a AccountOauthInfo) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, a.AdditionalProperties)
    structMap["accounts"] = a.Accounts
    if a.AuthorizationUrl != nil {
        structMap["authorization_url"] = a.AuthorizationUrl
    }
    structMap["linked"] = a.Linked
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
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "accounts", "authorization_url", "linked")
    if err != nil {
    	return err
    }
    a.AdditionalProperties = additionalProperties
    
    a.Accounts = *temp.Accounts
    a.AuthorizationUrl = temp.AuthorizationUrl
    a.Linked = *temp.Linked
    return nil
}

// tempAccountOauthInfo is a temporary struct used for validating the fields of AccountOauthInfo.
type tempAccountOauthInfo  struct {
    Accounts         *[]AccountOauthInfoAccount `json:"accounts"`
    AuthorizationUrl *string                    `json:"authorization_url,omitempty"`
    Linked           *bool                      `json:"linked"`
}

func (a *tempAccountOauthInfo) validate() error {
    var errs []string
    if a.Accounts == nil {
        errs = append(errs, "required field `accounts` is missing for type `account_oauth_info`")
    }
    if a.Linked == nil {
        errs = append(errs, "required field `linked` is missing for type `account_oauth_info`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
