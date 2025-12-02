// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
)

// NacClientEvent represents a NacClientEvent struct.
type NacClientEvent struct {
	// AP mac
	Ap *string `json:"ap,omitempty"`
	// enum: `cert`, `device-auth`, `eap-teap`, `eap-tls`, `eap-ttls`, `idp`, `mab`, `eap-peap`
	AuthType *NacAuthTypeEnum `json:"auth_type,omitempty"`
	// BSSID
	Bssid *string `json:"bssid,omitempty"`
	// Type of network access. enum: `wireless`, `wired`, `vty`
	ClientType *NacAccessTypeEnum `json:"client_type,omitempty"`
	// MAC Address of the device (AP, Switch) the client is connected to
	DeviceMac *string `json:"device_mac,omitempty"`
	// NAC Policy Dry Run Rule ID, if present and matched
	DryrunNacruleId *uuid.UUID `json:"dryrun_nacrule_id,omitempty"`
	// `true` if dryrun rule present and matched with priority, `false` if not matched or not present
	DryrunNacruleMatched *bool `json:"dryrun_nacrule_matched,omitempty"`
	// If IDP is used, the id of the IDP configuration used
	IdpId   *uuid.UUID `json:"idp_id,omitempty"`
	IdpRole []string   `json:"idp_role,omitempty"`
	// If IDP is used, the username presented to the Identity Provider
	IdpUsername *string `json:"idp_username,omitempty"`
	// Client MAC address
	Mac *string `json:"mac,omitempty"`
	// Mist Edge ID used to connect to cloud
	MxedgeId *string `json:"mxedge_id,omitempty"`
	// NAC Policy Rule ID, if matched
	NacruleId *uuid.UUID `json:"nacrule_id,omitempty"`
	// NAC Policy Rule Matched
	NacruleMatched *bool `json:"nacrule_matched,omitempty"`
	// Vendor name of the NAS
	NasVendor *string    `json:"nas_vendor,omitempty"`
	OrgId     *uuid.UUID `json:"org_id,omitempty"`
	// Port ID where the NAC client event occurred
	PortId *string `json:"port_id,omitempty"`
	// Type of network access. enum: `wireless`, `wired`, `vty`
	PortType *NacAccessTypeEnum `json:"port_type,omitempty"`
	// Whether the client is using randomized MAC Address or not
	RandomMac *RandomMacEnum `json:"random_mac,omitempty"`
	// List of Radius AVP returned by the Authentication Server
	RespAttrs []string   `json:"resp_attrs,omitempty"`
	SiteId    *uuid.UUID `json:"site_id,omitempty"`
	// SSIDs the client was connecting to
	Ssid *string `json:"ssid,omitempty"`
	// Epoch (seconds)
	Timestamp *float64 `json:"timestamp,omitempty"`
	// Event type, e.g. NAC_CLIENT_PERMIT. Use the [List NAC Events Definitions]($e/Constants%20Events/listNacEventsDefinitions) endpoint to get the full list of available values.
	Type *string `json:"type,omitempty"`
	// Labels derived from usermac entry
	UsermacLabel []string `json:"usermac_label,omitempty"`
	// username assigned to the client
	Username *string `json:"username,omitempty"`
	// vlan that assigned to the client
	Vlan *string `json:"vlan,omitempty"`
	// Vlan source, e.g. "nactag", "usermac"
	VlanSource           *string                `json:"vlan_source,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for NacClientEvent,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (n NacClientEvent) String() string {
	return fmt.Sprintf(
		"NacClientEvent[Ap=%v, AuthType=%v, Bssid=%v, ClientType=%v, DeviceMac=%v, DryrunNacruleId=%v, DryrunNacruleMatched=%v, IdpId=%v, IdpRole=%v, IdpUsername=%v, Mac=%v, MxedgeId=%v, NacruleId=%v, NacruleMatched=%v, NasVendor=%v, OrgId=%v, PortId=%v, PortType=%v, RandomMac=%v, RespAttrs=%v, SiteId=%v, Ssid=%v, Timestamp=%v, Type=%v, UsermacLabel=%v, Username=%v, Vlan=%v, VlanSource=%v, AdditionalProperties=%v]",
		n.Ap, n.AuthType, n.Bssid, n.ClientType, n.DeviceMac, n.DryrunNacruleId, n.DryrunNacruleMatched, n.IdpId, n.IdpRole, n.IdpUsername, n.Mac, n.MxedgeId, n.NacruleId, n.NacruleMatched, n.NasVendor, n.OrgId, n.PortId, n.PortType, n.RandomMac, n.RespAttrs, n.SiteId, n.Ssid, n.Timestamp, n.Type, n.UsermacLabel, n.Username, n.Vlan, n.VlanSource, n.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for NacClientEvent.
// It customizes the JSON marshaling process for NacClientEvent objects.
func (n NacClientEvent) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(n.AdditionalProperties,
		"ap", "auth_type", "bssid", "client_type", "device_mac", "dryrun_nacrule_id", "dryrun_nacrule_matched", "idp_id", "idp_role", "idp_username", "mac", "mxedge_id", "nacrule_id", "nacrule_matched", "nas_vendor", "org_id", "port_id", "port_type", "random_mac", "resp_attrs", "site_id", "ssid", "timestamp", "type", "usermac_label", "username", "vlan", "vlan_source"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(n.toMap())
}

// toMap converts the NacClientEvent object to a map representation for JSON marshaling.
func (n NacClientEvent) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, n.AdditionalProperties)
	if n.Ap != nil {
		structMap["ap"] = n.Ap
	}
	if n.AuthType != nil {
		structMap["auth_type"] = n.AuthType
	}
	if n.Bssid != nil {
		structMap["bssid"] = n.Bssid
	}
	if n.ClientType != nil {
		structMap["client_type"] = n.ClientType
	}
	if n.DeviceMac != nil {
		structMap["device_mac"] = n.DeviceMac
	}
	if n.DryrunNacruleId != nil {
		structMap["dryrun_nacrule_id"] = n.DryrunNacruleId
	}
	if n.DryrunNacruleMatched != nil {
		structMap["dryrun_nacrule_matched"] = n.DryrunNacruleMatched
	}
	if n.IdpId != nil {
		structMap["idp_id"] = n.IdpId
	}
	if n.IdpRole != nil {
		structMap["idp_role"] = n.IdpRole
	}
	if n.IdpUsername != nil {
		structMap["idp_username"] = n.IdpUsername
	}
	if n.Mac != nil {
		structMap["mac"] = n.Mac
	}
	if n.MxedgeId != nil {
		structMap["mxedge_id"] = n.MxedgeId
	}
	if n.NacruleId != nil {
		structMap["nacrule_id"] = n.NacruleId
	}
	if n.NacruleMatched != nil {
		structMap["nacrule_matched"] = n.NacruleMatched
	}
	if n.NasVendor != nil {
		structMap["nas_vendor"] = n.NasVendor
	}
	if n.OrgId != nil {
		structMap["org_id"] = n.OrgId
	}
	if n.PortId != nil {
		structMap["port_id"] = n.PortId
	}
	if n.PortType != nil {
		structMap["port_type"] = n.PortType
	}
	if n.RandomMac != nil {
		structMap["random_mac"] = n.RandomMac
	}
	if n.RespAttrs != nil {
		structMap["resp_attrs"] = n.RespAttrs
	}
	if n.SiteId != nil {
		structMap["site_id"] = n.SiteId
	}
	if n.Ssid != nil {
		structMap["ssid"] = n.Ssid
	}
	if n.Timestamp != nil {
		structMap["timestamp"] = n.Timestamp
	}
	if n.Type != nil {
		structMap["type"] = n.Type
	}
	if n.UsermacLabel != nil {
		structMap["usermac_label"] = n.UsermacLabel
	}
	if n.Username != nil {
		structMap["username"] = n.Username
	}
	if n.Vlan != nil {
		structMap["vlan"] = n.Vlan
	}
	if n.VlanSource != nil {
		structMap["vlan_source"] = n.VlanSource
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for NacClientEvent.
// It customizes the JSON unmarshaling process for NacClientEvent objects.
func (n *NacClientEvent) UnmarshalJSON(input []byte) error {
	var temp tempNacClientEvent
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "ap", "auth_type", "bssid", "client_type", "device_mac", "dryrun_nacrule_id", "dryrun_nacrule_matched", "idp_id", "idp_role", "idp_username", "mac", "mxedge_id", "nacrule_id", "nacrule_matched", "nas_vendor", "org_id", "port_id", "port_type", "random_mac", "resp_attrs", "site_id", "ssid", "timestamp", "type", "usermac_label", "username", "vlan", "vlan_source")
	if err != nil {
		return err
	}
	n.AdditionalProperties = additionalProperties

	n.Ap = temp.Ap
	n.AuthType = temp.AuthType
	n.Bssid = temp.Bssid
	n.ClientType = temp.ClientType
	n.DeviceMac = temp.DeviceMac
	n.DryrunNacruleId = temp.DryrunNacruleId
	n.DryrunNacruleMatched = temp.DryrunNacruleMatched
	n.IdpId = temp.IdpId
	n.IdpRole = temp.IdpRole
	n.IdpUsername = temp.IdpUsername
	n.Mac = temp.Mac
	n.MxedgeId = temp.MxedgeId
	n.NacruleId = temp.NacruleId
	n.NacruleMatched = temp.NacruleMatched
	n.NasVendor = temp.NasVendor
	n.OrgId = temp.OrgId
	n.PortId = temp.PortId
	n.PortType = temp.PortType
	n.RandomMac = temp.RandomMac
	n.RespAttrs = temp.RespAttrs
	n.SiteId = temp.SiteId
	n.Ssid = temp.Ssid
	n.Timestamp = temp.Timestamp
	n.Type = temp.Type
	n.UsermacLabel = temp.UsermacLabel
	n.Username = temp.Username
	n.Vlan = temp.Vlan
	n.VlanSource = temp.VlanSource
	return nil
}

// tempNacClientEvent is a temporary struct used for validating the fields of NacClientEvent.
type tempNacClientEvent struct {
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
	MxedgeId             *string            `json:"mxedge_id,omitempty"`
	NacruleId            *uuid.UUID         `json:"nacrule_id,omitempty"`
	NacruleMatched       *bool              `json:"nacrule_matched,omitempty"`
	NasVendor            *string            `json:"nas_vendor,omitempty"`
	OrgId                *uuid.UUID         `json:"org_id,omitempty"`
	PortId               *string            `json:"port_id,omitempty"`
	PortType             *NacAccessTypeEnum `json:"port_type,omitempty"`
	RandomMac            *RandomMacEnum     `json:"random_mac,omitempty"`
	RespAttrs            []string           `json:"resp_attrs,omitempty"`
	SiteId               *uuid.UUID         `json:"site_id,omitempty"`
	Ssid                 *string            `json:"ssid,omitempty"`
	Timestamp            *float64           `json:"timestamp,omitempty"`
	Type                 *string            `json:"type,omitempty"`
	UsermacLabel         []string           `json:"usermac_label,omitempty"`
	Username             *string            `json:"username,omitempty"`
	Vlan                 *string            `json:"vlan,omitempty"`
	VlanSource           *string            `json:"vlan_source,omitempty"`
}
