// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
)

// Guest represents a Guest struct.
// Guest authorization record at site scope
type Guest struct {
	// If `auth_method`==`email`, the email address where the authorization code has been sent to
	AccessCodeEmail *string `json:"access_code_email,omitempty"`
	// MAC address of the AP used during guest registration
	ApMac *string `json:"ap_mac,omitempty"`
	// Guest authentication method used for the authorization
	AuthMethod *string `json:"auth_method,omitempty"`
	// Whether the guest is currently authorized
	Authorized *bool `json:"authorized,omitempty"`
	// Unix timestamp when the guest authorization expires
	AuthorizedExpiringTime *float64 `json:"authorized_expiring_time,omitempty"`
	// Unix timestamp when the guest was authorized
	AuthorizedTime *float64 `json:"authorized_time,omitempty"`
	// Optional company name provided by the guest during registration
	Company *string `json:"company,omitempty"`
	// Optional email address provided by the guest during registration
	Email *string `json:"email,omitempty"`
	// Optional custom field 1 value provided by the guest during registration
	Field1 *string `json:"field1,omitempty"`
	// Optional custom field 2 value provided by the guest during registration
	Field2 *string `json:"field2,omitempty"`
	// Optional custom field 3 value provided by the guest during registration
	Field3 *string `json:"field3,omitempty"`
	// Optional custom field 4 value provided by the guest during registration
	Field4 *string `json:"field4,omitempty"`
	// Device MAC address captured during guest registration
	Mac *string `json:"mac,omitempty"`
	// Authorization duration, in minutes. Default is 1440 minutes (1 day), maximum is 259200 (180 days)
	Minutes *int `json:"minutes,omitempty"`
	// Optional name provided by the guest during registration
	Name *string `json:"name,omitempty"`
	// Whether the guest device used a randomized MAC address to connect to the SSID
	RandomMac *bool `json:"random_mac,omitempty"`
	// Name of the SSID
	Ssid *string `json:"ssid,omitempty"`
	// Identifier of the WLAN used for the guest authorization
	WlanId               *uuid.UUID             `json:"wlan_id,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for Guest,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (g Guest) String() string {
	return fmt.Sprintf(
		"Guest[AccessCodeEmail=%v, ApMac=%v, AuthMethod=%v, Authorized=%v, AuthorizedExpiringTime=%v, AuthorizedTime=%v, Company=%v, Email=%v, Field1=%v, Field2=%v, Field3=%v, Field4=%v, Mac=%v, Minutes=%v, Name=%v, RandomMac=%v, Ssid=%v, WlanId=%v, AdditionalProperties=%v]",
		g.AccessCodeEmail, g.ApMac, g.AuthMethod, g.Authorized, g.AuthorizedExpiringTime, g.AuthorizedTime, g.Company, g.Email, g.Field1, g.Field2, g.Field3, g.Field4, g.Mac, g.Minutes, g.Name, g.RandomMac, g.Ssid, g.WlanId, g.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for Guest.
// It customizes the JSON marshaling process for Guest objects.
func (g Guest) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(g.AdditionalProperties,
		"access_code_email", "ap_mac", "auth_method", "authorized", "authorized_expiring_time", "authorized_time", "company", "email", "field1", "field2", "field3", "field4", "mac", "minutes", "name", "random_mac", "ssid", "wlan_id"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(g.toMap())
}

// toMap converts the Guest object to a map representation for JSON marshaling.
func (g Guest) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, g.AdditionalProperties)
	if g.AccessCodeEmail != nil {
		structMap["access_code_email"] = g.AccessCodeEmail
	}
	if g.ApMac != nil {
		structMap["ap_mac"] = g.ApMac
	}
	if g.AuthMethod != nil {
		structMap["auth_method"] = g.AuthMethod
	}
	if g.Authorized != nil {
		structMap["authorized"] = g.Authorized
	}
	if g.AuthorizedExpiringTime != nil {
		structMap["authorized_expiring_time"] = g.AuthorizedExpiringTime
	}
	if g.AuthorizedTime != nil {
		structMap["authorized_time"] = g.AuthorizedTime
	}
	if g.Company != nil {
		structMap["company"] = g.Company
	}
	if g.Email != nil {
		structMap["email"] = g.Email
	}
	if g.Field1 != nil {
		structMap["field1"] = g.Field1
	}
	if g.Field2 != nil {
		structMap["field2"] = g.Field2
	}
	if g.Field3 != nil {
		structMap["field3"] = g.Field3
	}
	if g.Field4 != nil {
		structMap["field4"] = g.Field4
	}
	if g.Mac != nil {
		structMap["mac"] = g.Mac
	}
	if g.Minutes != nil {
		structMap["minutes"] = g.Minutes
	}
	if g.Name != nil {
		structMap["name"] = g.Name
	}
	if g.RandomMac != nil {
		structMap["random_mac"] = g.RandomMac
	}
	if g.Ssid != nil {
		structMap["ssid"] = g.Ssid
	}
	if g.WlanId != nil {
		structMap["wlan_id"] = g.WlanId
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Guest.
// It customizes the JSON unmarshaling process for Guest objects.
func (g *Guest) UnmarshalJSON(input []byte) error {
	var temp tempGuest
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "access_code_email", "ap_mac", "auth_method", "authorized", "authorized_expiring_time", "authorized_time", "company", "email", "field1", "field2", "field3", "field4", "mac", "minutes", "name", "random_mac", "ssid", "wlan_id")
	if err != nil {
		return err
	}
	g.AdditionalProperties = additionalProperties

	g.AccessCodeEmail = temp.AccessCodeEmail
	g.ApMac = temp.ApMac
	g.AuthMethod = temp.AuthMethod
	g.Authorized = temp.Authorized
	g.AuthorizedExpiringTime = temp.AuthorizedExpiringTime
	g.AuthorizedTime = temp.AuthorizedTime
	g.Company = temp.Company
	g.Email = temp.Email
	g.Field1 = temp.Field1
	g.Field2 = temp.Field2
	g.Field3 = temp.Field3
	g.Field4 = temp.Field4
	g.Mac = temp.Mac
	g.Minutes = temp.Minutes
	g.Name = temp.Name
	g.RandomMac = temp.RandomMac
	g.Ssid = temp.Ssid
	g.WlanId = temp.WlanId
	return nil
}

// tempGuest is a temporary struct used for validating the fields of Guest.
type tempGuest struct {
	AccessCodeEmail        *string    `json:"access_code_email,omitempty"`
	ApMac                  *string    `json:"ap_mac,omitempty"`
	AuthMethod             *string    `json:"auth_method,omitempty"`
	Authorized             *bool      `json:"authorized,omitempty"`
	AuthorizedExpiringTime *float64   `json:"authorized_expiring_time,omitempty"`
	AuthorizedTime         *float64   `json:"authorized_time,omitempty"`
	Company                *string    `json:"company,omitempty"`
	Email                  *string    `json:"email,omitempty"`
	Field1                 *string    `json:"field1,omitempty"`
	Field2                 *string    `json:"field2,omitempty"`
	Field3                 *string    `json:"field3,omitempty"`
	Field4                 *string    `json:"field4,omitempty"`
	Mac                    *string    `json:"mac,omitempty"`
	Minutes                *int       `json:"minutes,omitempty"`
	Name                   *string    `json:"name,omitempty"`
	RandomMac              *bool      `json:"random_mac,omitempty"`
	Ssid                   *string    `json:"ssid,omitempty"`
	WlanId                 *uuid.UUID `json:"wlan_id,omitempty"`
}
