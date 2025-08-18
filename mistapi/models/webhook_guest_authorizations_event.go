// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
)

// WebhookGuestAuthorizationsEvent represents a WebhookGuestAuthorizationsEvent struct.
type WebhookGuestAuthorizationsEvent struct {
	// mac address of the AP the guest is connected to
	Ap *string `json:"ap,omitempty"`
	// authentication method used
	AuthMethod *string `json:"auth_method,omitempty"`
	// expiry time for guest
	AuthorizedExpiringTime *int `json:"authorized_expiring_time,omitempty"`
	// time of authorization of guest
	AuthorizedTime *int `json:"authorized_time,omitempty"`
	// carrier used when authentication by free cell provider
	Carrier *string `json:"carrier,omitempty"`
	// client mac
	Client *string `json:"client,omitempty"`
	// guest company
	Company *string `json:"company,omitempty"`
	// guest email
	Email *string `json:"email,omitempty"`
	// field1 value
	Field1 *string `json:"field1,omitempty"`
	// field2 value
	Field2 *string `json:"field2,omitempty"`
	// field3 value
	Field3 *string `json:"field3,omitempty"`
	// field4 value
	Field4 *string `json:"field4,omitempty"`
	// guest mobile number
	Mobile *string `json:"mobile,omitempty"`
	// guest name
	Name   *string    `json:"name,omitempty"`
	OrgId  *uuid.UUID `json:"org_id,omitempty"`
	SiteId *uuid.UUID `json:"site_id,omitempty"`
	// sms gateway used via text auth paid service
	SmsGateway *string `json:"sms_gateway,omitempty"`
	// guest sponsor email
	SponsorEmail *string `json:"sponsor_email,omitempty"`
	// ssid
	Ssid *string `json:"ssid,omitempty"`
	// wlan id
	WlanId               *string                `json:"wlan_id,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for WebhookGuestAuthorizationsEvent,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WebhookGuestAuthorizationsEvent) String() string {
	return fmt.Sprintf(
		"WebhookGuestAuthorizationsEvent[Ap=%v, AuthMethod=%v, AuthorizedExpiringTime=%v, AuthorizedTime=%v, Carrier=%v, Client=%v, Company=%v, Email=%v, Field1=%v, Field2=%v, Field3=%v, Field4=%v, Mobile=%v, Name=%v, OrgId=%v, SiteId=%v, SmsGateway=%v, SponsorEmail=%v, Ssid=%v, WlanId=%v, AdditionalProperties=%v]",
		w.Ap, w.AuthMethod, w.AuthorizedExpiringTime, w.AuthorizedTime, w.Carrier, w.Client, w.Company, w.Email, w.Field1, w.Field2, w.Field3, w.Field4, w.Mobile, w.Name, w.OrgId, w.SiteId, w.SmsGateway, w.SponsorEmail, w.Ssid, w.WlanId, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WebhookGuestAuthorizationsEvent.
// It customizes the JSON marshaling process for WebhookGuestAuthorizationsEvent objects.
func (w WebhookGuestAuthorizationsEvent) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(w.AdditionalProperties,
		"ap", "auth_method", "authorized_expiring_time", "authorized_time", "carrier", "client", "company", "email", "field1", "field2", "field3", "field4", "mobile", "name", "org_id", "site_id", "sms_gateway", "sponsor_email", "ssid", "wlan_id"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(w.toMap())
}

// toMap converts the WebhookGuestAuthorizationsEvent object to a map representation for JSON marshaling.
func (w WebhookGuestAuthorizationsEvent) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, w.AdditionalProperties)
	if w.Ap != nil {
		structMap["ap"] = w.Ap
	}
	if w.AuthMethod != nil {
		structMap["auth_method"] = w.AuthMethod
	}
	if w.AuthorizedExpiringTime != nil {
		structMap["authorized_expiring_time"] = w.AuthorizedExpiringTime
	}
	if w.AuthorizedTime != nil {
		structMap["authorized_time"] = w.AuthorizedTime
	}
	if w.Carrier != nil {
		structMap["carrier"] = w.Carrier
	}
	if w.Client != nil {
		structMap["client"] = w.Client
	}
	if w.Company != nil {
		structMap["company"] = w.Company
	}
	if w.Email != nil {
		structMap["email"] = w.Email
	}
	if w.Field1 != nil {
		structMap["field1"] = w.Field1
	}
	if w.Field2 != nil {
		structMap["field2"] = w.Field2
	}
	if w.Field3 != nil {
		structMap["field3"] = w.Field3
	}
	if w.Field4 != nil {
		structMap["field4"] = w.Field4
	}
	if w.Mobile != nil {
		structMap["mobile"] = w.Mobile
	}
	if w.Name != nil {
		structMap["name"] = w.Name
	}
	if w.OrgId != nil {
		structMap["org_id"] = w.OrgId
	}
	if w.SiteId != nil {
		structMap["site_id"] = w.SiteId
	}
	if w.SmsGateway != nil {
		structMap["sms_gateway"] = w.SmsGateway
	}
	if w.SponsorEmail != nil {
		structMap["sponsor_email"] = w.SponsorEmail
	}
	if w.Ssid != nil {
		structMap["ssid"] = w.Ssid
	}
	if w.WlanId != nil {
		structMap["wlan_id"] = w.WlanId
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookGuestAuthorizationsEvent.
// It customizes the JSON unmarshaling process for WebhookGuestAuthorizationsEvent objects.
func (w *WebhookGuestAuthorizationsEvent) UnmarshalJSON(input []byte) error {
	var temp tempWebhookGuestAuthorizationsEvent
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "ap", "auth_method", "authorized_expiring_time", "authorized_time", "carrier", "client", "company", "email", "field1", "field2", "field3", "field4", "mobile", "name", "org_id", "site_id", "sms_gateway", "sponsor_email", "ssid", "wlan_id")
	if err != nil {
		return err
	}
	w.AdditionalProperties = additionalProperties

	w.Ap = temp.Ap
	w.AuthMethod = temp.AuthMethod
	w.AuthorizedExpiringTime = temp.AuthorizedExpiringTime
	w.AuthorizedTime = temp.AuthorizedTime
	w.Carrier = temp.Carrier
	w.Client = temp.Client
	w.Company = temp.Company
	w.Email = temp.Email
	w.Field1 = temp.Field1
	w.Field2 = temp.Field2
	w.Field3 = temp.Field3
	w.Field4 = temp.Field4
	w.Mobile = temp.Mobile
	w.Name = temp.Name
	w.OrgId = temp.OrgId
	w.SiteId = temp.SiteId
	w.SmsGateway = temp.SmsGateway
	w.SponsorEmail = temp.SponsorEmail
	w.Ssid = temp.Ssid
	w.WlanId = temp.WlanId
	return nil
}

// tempWebhookGuestAuthorizationsEvent is a temporary struct used for validating the fields of WebhookGuestAuthorizationsEvent.
type tempWebhookGuestAuthorizationsEvent struct {
	Ap                     *string    `json:"ap,omitempty"`
	AuthMethod             *string    `json:"auth_method,omitempty"`
	AuthorizedExpiringTime *int       `json:"authorized_expiring_time,omitempty"`
	AuthorizedTime         *int       `json:"authorized_time,omitempty"`
	Carrier                *string    `json:"carrier,omitempty"`
	Client                 *string    `json:"client,omitempty"`
	Company                *string    `json:"company,omitempty"`
	Email                  *string    `json:"email,omitempty"`
	Field1                 *string    `json:"field1,omitempty"`
	Field2                 *string    `json:"field2,omitempty"`
	Field3                 *string    `json:"field3,omitempty"`
	Field4                 *string    `json:"field4,omitempty"`
	Mobile                 *string    `json:"mobile,omitempty"`
	Name                   *string    `json:"name,omitempty"`
	OrgId                  *uuid.UUID `json:"org_id,omitempty"`
	SiteId                 *uuid.UUID `json:"site_id,omitempty"`
	SmsGateway             *string    `json:"sms_gateway,omitempty"`
	SponsorEmail           *string    `json:"sponsor_email,omitempty"`
	Ssid                   *string    `json:"ssid,omitempty"`
	WlanId                 *string    `json:"wlan_id,omitempty"`
}
