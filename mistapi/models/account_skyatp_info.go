// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// AccountSkyatpInfo represents a AccountSkyatpInfo struct.
// Linked Sky ATP account and realm information
type AccountSkyatpInfo struct {
	// Sky ATP cloud name. enum: `www.amerskyatp.com`, `www.apacskyatp.com`, `www.euroskyatp.com`, `www.canadaskyatp.com`
	CloudName *AccountSkyatpCloudNameEnum `json:"cloud_name,omitempty"`
	// Sky ATP realm linked with this Mist organization
	Realm *string `json:"realm,omitempty"`
	// Sky ATP username configured for the linked realm
	Username             *string                `json:"username,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for AccountSkyatpInfo,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a AccountSkyatpInfo) String() string {
	return fmt.Sprintf(
		"AccountSkyatpInfo[CloudName=%v, Realm=%v, Username=%v, AdditionalProperties=%v]",
		a.CloudName, a.Realm, a.Username, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for AccountSkyatpInfo.
// It customizes the JSON marshaling process for AccountSkyatpInfo objects.
func (a AccountSkyatpInfo) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(a.AdditionalProperties,
		"cloud_name", "realm", "username"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(a.toMap())
}

// toMap converts the AccountSkyatpInfo object to a map representation for JSON marshaling.
func (a AccountSkyatpInfo) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, a.AdditionalProperties)
	if a.CloudName != nil {
		structMap["cloud_name"] = a.CloudName
	}
	if a.Realm != nil {
		structMap["realm"] = a.Realm
	}
	if a.Username != nil {
		structMap["username"] = a.Username
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AccountSkyatpInfo.
// It customizes the JSON unmarshaling process for AccountSkyatpInfo objects.
func (a *AccountSkyatpInfo) UnmarshalJSON(input []byte) error {
	var temp tempAccountSkyatpInfo
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "cloud_name", "realm", "username")
	if err != nil {
		return err
	}
	a.AdditionalProperties = additionalProperties

	a.CloudName = temp.CloudName
	a.Realm = temp.Realm
	a.Username = temp.Username
	return nil
}

// tempAccountSkyatpInfo is a temporary struct used for validating the fields of AccountSkyatpInfo.
type tempAccountSkyatpInfo struct {
	CloudName *AccountSkyatpCloudNameEnum `json:"cloud_name,omitempty"`
	Realm     *string                     `json:"realm,omitempty"`
	Username  *string                     `json:"username,omitempty"`
}
