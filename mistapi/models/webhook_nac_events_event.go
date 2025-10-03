// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
)

// WebhookNacEventsEvent represents a WebhookNacEventsEvent struct.
type WebhookNacEventsEvent struct {
	// AP mac
	Ap *string `json:"ap,omitempty"`
	// enum: `cert`, `device-auth`, `eap-teap`, `eap-tls`, `eap-ttls`, `idp`, `mab`, `eap-peap`
	AuthType *NacAuthTypeEnum `json:"auth_type,omitempty"`
	// BSSID
	Bssid *string `json:"bssid,omitempty"`
	// Type of network access. enum: `wireless`, `wired`
	ClientType *NacAccessTypeEnum `json:"client_type,omitempty"`
	// MAC Address of the device (AP, Switch) the client is connected to
	DeviceMac *string `json:"device_mac,omitempty"`
	// NAC Policy Dry Run Rule ID, if present and matched
	DryrunNacruleId *uuid.UUID `json:"dryrun_nacrule_id,omitempty"`
	// `true` if dryrun rule present and matched with priority, `false` if not matched or not present
	DryrunNacruleMatched *bool `json:"dryrun_nacrule_matched,omitempty"`
	// If IDP is used, the id of the IDP configuration used
	IdpId *uuid.UUID `json:"idp_id,omitempty"`
	// IDP returned roles/groups for the user
	IdpRole []string `json:"idp_role,omitempty"`
	// If IDP is used, the username presented to the Identity Provider
	IdpUsername *string `json:"idp_username,omitempty"`
	// Client MAC address
	Mac *string `json:"mac,omitempty"`
	// NAC Policy Rule ID, if matched
	NacruleId *uuid.UUID `json:"nacrule_id,omitempty"`
	// NAC Policy Rule Matched
	NacruleMatched *bool `json:"nacrule_matched,omitempty"`
	// Vendor name of the NAS
	NasVendor *string    `json:"nas_vendor,omitempty"`
	OrgId     *uuid.UUID `json:"org_id,omitempty"`
	// Port-ids the client was connected to  for the specified duration
	PortId []string `json:"port_id,omitempty"`
	// Whether the client is using randomized MAC Address or not
	RandomMac *bool `json:"random_mac,omitempty"`
	// List of Radius AVP returned by the Authentication Server
	RespAttrs []string   `json:"resp_attrs,omitempty"`
	SiteId    *uuid.UUID `json:"site_id,omitempty"`
	// SSIDs the client was connecting to
	Ssid *string `json:"ssid,omitempty"`
	// Epoch (seconds)
	Timestamp *float64 `json:"timestamp,omitempty"`
	// Event type, e.g. NAC_CLIENT_PERMIT. Use the [List NAC Events Definitions](/#operations/listNacEventsDefinitions) endpoint to get the full list of available values.
	Type *string `json:"type,omitempty"`
	// username assigned to the client
	Username *string `json:"username,omitempty"`
	// vlan that assigned to the client
	Vlan *string `json:"vlan,omitempty"`
	// Vlan source, e.g. "nactag", "usermac"
	VlanSource           *string                `json:"vlan_source,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for WebhookNacEventsEvent,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WebhookNacEventsEvent) String() string {
	return fmt.Sprintf(
		"WebhookNacEventsEvent[Ap=%v, AuthType=%v, Bssid=%v, ClientType=%v, DeviceMac=%v, DryrunNacruleId=%v, DryrunNacruleMatched=%v, IdpId=%v, IdpRole=%v, IdpUsername=%v, Mac=%v, NacruleId=%v, NacruleMatched=%v, NasVendor=%v, OrgId=%v, PortId=%v, RandomMac=%v, RespAttrs=%v, SiteId=%v, Ssid=%v, Timestamp=%v, Type=%v, Username=%v, Vlan=%v, VlanSource=%v, AdditionalProperties=%v]",
		w.Ap, w.AuthType, w.Bssid, w.ClientType, w.DeviceMac, w.DryrunNacruleId, w.DryrunNacruleMatched, w.IdpId, w.IdpRole, w.IdpUsername, w.Mac, w.NacruleId, w.NacruleMatched, w.NasVendor, w.OrgId, w.PortId, w.RandomMac, w.RespAttrs, w.SiteId, w.Ssid, w.Timestamp, w.Type, w.Username, w.Vlan, w.VlanSource, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WebhookNacEventsEvent.
// It customizes the JSON marshaling process for WebhookNacEventsEvent objects.
func (w WebhookNacEventsEvent) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(w.AdditionalProperties,
		"ap", "auth_type", "bssid", "client_type", "device_mac", "dryrun_nacrule_id", "dryrun_nacrule_matched", "idp_id", "idp_role", "idp_username", "mac", "nacrule_id", "nacrule_matched", "nas_vendor", "org_id", "port_id", "random_mac", "resp_attrs", "site_id", "ssid", "timestamp", "type", "username", "vlan", "vlan_source"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(w.toMap())
}

// toMap converts the WebhookNacEventsEvent object to a map representation for JSON marshaling.
func (w WebhookNacEventsEvent) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, w.AdditionalProperties)
	if w.Ap != nil {
		structMap["ap"] = w.Ap
	}
	if w.AuthType != nil {
		structMap["auth_type"] = w.AuthType
	}
	if w.Bssid != nil {
		structMap["bssid"] = w.Bssid
	}
	if w.ClientType != nil {
		structMap["client_type"] = w.ClientType
	}
	if w.DeviceMac != nil {
		structMap["device_mac"] = w.DeviceMac
	}
	if w.DryrunNacruleId != nil {
		structMap["dryrun_nacrule_id"] = w.DryrunNacruleId
	}
	if w.DryrunNacruleMatched != nil {
		structMap["dryrun_nacrule_matched"] = w.DryrunNacruleMatched
	}
	if w.IdpId != nil {
		structMap["idp_id"] = w.IdpId
	}
	if w.IdpRole != nil {
		structMap["idp_role"] = w.IdpRole
	}
	if w.IdpUsername != nil {
		structMap["idp_username"] = w.IdpUsername
	}
	if w.Mac != nil {
		structMap["mac"] = w.Mac
	}
	if w.NacruleId != nil {
		structMap["nacrule_id"] = w.NacruleId
	}
	if w.NacruleMatched != nil {
		structMap["nacrule_matched"] = w.NacruleMatched
	}
	if w.NasVendor != nil {
		structMap["nas_vendor"] = w.NasVendor
	}
	if w.OrgId != nil {
		structMap["org_id"] = w.OrgId
	}
	if w.PortId != nil {
		structMap["port_id"] = w.PortId
	}
	if w.RandomMac != nil {
		structMap["random_mac"] = w.RandomMac
	}
	if w.RespAttrs != nil {
		structMap["resp_attrs"] = w.RespAttrs
	}
	if w.SiteId != nil {
		structMap["site_id"] = w.SiteId
	}
	if w.Ssid != nil {
		structMap["ssid"] = w.Ssid
	}
	if w.Timestamp != nil {
		structMap["timestamp"] = w.Timestamp
	}
	if w.Type != nil {
		structMap["type"] = w.Type
	}
	if w.Username != nil {
		structMap["username"] = w.Username
	}
	if w.Vlan != nil {
		structMap["vlan"] = w.Vlan
	}
	if w.VlanSource != nil {
		structMap["vlan_source"] = w.VlanSource
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookNacEventsEvent.
// It customizes the JSON unmarshaling process for WebhookNacEventsEvent objects.
func (w *WebhookNacEventsEvent) UnmarshalJSON(input []byte) error {
	var temp tempWebhookNacEventsEvent
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "ap", "auth_type", "bssid", "client_type", "device_mac", "dryrun_nacrule_id", "dryrun_nacrule_matched", "idp_id", "idp_role", "idp_username", "mac", "nacrule_id", "nacrule_matched", "nas_vendor", "org_id", "port_id", "random_mac", "resp_attrs", "site_id", "ssid", "timestamp", "type", "username", "vlan", "vlan_source")
	if err != nil {
		return err
	}
	w.AdditionalProperties = additionalProperties

	w.Ap = temp.Ap
	w.AuthType = temp.AuthType
	w.Bssid = temp.Bssid
	w.ClientType = temp.ClientType
	w.DeviceMac = temp.DeviceMac
	w.DryrunNacruleId = temp.DryrunNacruleId
	w.DryrunNacruleMatched = temp.DryrunNacruleMatched
	w.IdpId = temp.IdpId
	w.IdpRole = temp.IdpRole
	w.IdpUsername = temp.IdpUsername
	w.Mac = temp.Mac
	w.NacruleId = temp.NacruleId
	w.NacruleMatched = temp.NacruleMatched
	w.NasVendor = temp.NasVendor
	w.OrgId = temp.OrgId
	w.PortId = temp.PortId
	w.RandomMac = temp.RandomMac
	w.RespAttrs = temp.RespAttrs
	w.SiteId = temp.SiteId
	w.Ssid = temp.Ssid
	w.Timestamp = temp.Timestamp
	w.Type = temp.Type
	w.Username = temp.Username
	w.Vlan = temp.Vlan
	w.VlanSource = temp.VlanSource
	return nil
}

// tempWebhookNacEventsEvent is a temporary struct used for validating the fields of WebhookNacEventsEvent.
type tempWebhookNacEventsEvent struct {
	Ap                   *string            `json:"ap,omitempty"`
	AuthType             *NacAuthTypeEnum   `json:"auth_type,omitempty"`
	Bssid                *string            `json:"bssid,omitempty"`
	ClientType           *NacAccessTypeEnum `json:"client_type,omitempty"`
	DeviceMac            *string            `json:"device_mac,omitempty"`
	DryrunNacruleId      *uuid.UUID         `json:"dryrun_nacrule_id,omitempty"`
	DryrunNacruleMatched *bool              `json:"dryrun_nacrule_matched,omitempty"`
	IdpId                *uuid.UUID         `json:"idp_id,omitempty"`
	IdpRole              []string           `json:"idp_role,omitempty"`
	IdpUsername          *string            `json:"idp_username,omitempty"`
	Mac                  *string            `json:"mac,omitempty"`
	NacruleId            *uuid.UUID         `json:"nacrule_id,omitempty"`
	NacruleMatched       *bool              `json:"nacrule_matched,omitempty"`
	NasVendor            *string            `json:"nas_vendor,omitempty"`
	OrgId                *uuid.UUID         `json:"org_id,omitempty"`
	PortId               []string           `json:"port_id,omitempty"`
	RandomMac            *bool              `json:"random_mac,omitempty"`
	RespAttrs            []string           `json:"resp_attrs,omitempty"`
	SiteId               *uuid.UUID         `json:"site_id,omitempty"`
	Ssid                 *string            `json:"ssid,omitempty"`
	Timestamp            *float64           `json:"timestamp,omitempty"`
	Type                 *string            `json:"type,omitempty"`
	Username             *string            `json:"username,omitempty"`
	Vlan                 *string            `json:"vlan,omitempty"`
	VlanSource           *string            `json:"vlan_source,omitempty"`
}
