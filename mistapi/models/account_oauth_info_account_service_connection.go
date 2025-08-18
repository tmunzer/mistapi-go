// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// AccountOauthInfoAccountServiceConnection represents a AccountOauthInfoAccountServiceConnection struct.
type AccountOauthInfoAccountServiceConnection struct {
    // Region of the service connection
    Region               *string                `json:"region,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for AccountOauthInfoAccountServiceConnection,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a AccountOauthInfoAccountServiceConnection) String() string {
    return fmt.Sprintf(
    	"AccountOauthInfoAccountServiceConnection[Region=%v, AdditionalProperties=%v]",
    	a.Region, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for AccountOauthInfoAccountServiceConnection.
// It customizes the JSON marshaling process for AccountOauthInfoAccountServiceConnection objects.
func (a AccountOauthInfoAccountServiceConnection) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(a.AdditionalProperties,
        "region"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(a.toMap())
}

// toMap converts the AccountOauthInfoAccountServiceConnection object to a map representation for JSON marshaling.
func (a AccountOauthInfoAccountServiceConnection) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, a.AdditionalProperties)
    if a.Region != nil {
        structMap["region"] = a.Region
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AccountOauthInfoAccountServiceConnection.
// It customizes the JSON unmarshaling process for AccountOauthInfoAccountServiceConnection objects.
func (a *AccountOauthInfoAccountServiceConnection) UnmarshalJSON(input []byte) error {
    var temp tempAccountOauthInfoAccountServiceConnection
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "region")
    if err != nil {
    	return err
    }
    a.AdditionalProperties = additionalProperties
    
    a.Region = temp.Region
    return nil
}

// tempAccountOauthInfoAccountServiceConnection is a temporary struct used for validating the fields of AccountOauthInfoAccountServiceConnection.
type tempAccountOauthInfoAccountServiceConnection  struct {
    Region *string `json:"region,omitempty"`
}
