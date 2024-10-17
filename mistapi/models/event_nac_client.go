package models

import (
	"encoding/json"
	"github.com/google/uuid"
)

// EventNacClient represents a EventNacClient struct.
type EventNacClient struct {
	Ap *string `json:"ap,omitempty"`
	// authentication type, e.g. "eap-tls", "peap-tls", "eap-ttls", "eap-teap", "mab", "psk", "device-auth"
	AuthType  *string `json:"auth_type,omitempty"`
	Bssid     *string `json:"bssid,omitempty"`
	DeviceMac *string `json:"device_mac,omitempty"`
	// NAC Policy Dry Run Rule ID, if present and matched
	DryrunNacruleId      *uuid.UUID `json:"dryrun_nacrule_id,omitempty"`
	DryrunNacruleMatched *bool      `json:"dryrun_nacrule_matched,omitempty"`
	IdpId                *uuid.UUID `json:"idp_id,omitempty"`
	IdpRole              []string   `json:"idp_role,omitempty"`
	// Username presented to the Identity Provider
	IdpUsername *string `json:"idp_username,omitempty"`
	// Client MAC address
	Mac *string `json:"mac,omitempty"`
	// Mist Edge ID used to connect to cloud
	MxedgeId *string `json:"mxedge_id,omitempty"`
	// NAC Policy Rule ID, if matched
	NacruleId      *uuid.UUID `json:"nacrule_id,omitempty"`
	NacruleMatched *bool      `json:"nacrule_matched,omitempty"`
	NasVendor      *string    `json:"nas_vendor,omitempty"`
	OrgId          *uuid.UUID `json:"org_id,omitempty"`
	PortId         *string    `json:"port_id,omitempty"`
	// Type of client i.e wired, wireless
	PortType  *string    `json:"port_type,omitempty"`
	RandomMac *bool      `json:"random_mac,omitempty"`
	RespAttrs []string   `json:"resp_attrs,omitempty"`
	SiteId    *uuid.UUID `json:"site_id,omitempty"`
	Ssid      *string    `json:"ssid,omitempty"`
	Timestamp *float64   `json:"timestamp,omitempty"`
	// event type, e.g. NAC_CLIENT_PERMIT
	Type *string `json:"type,omitempty"`
	// labels derived from usermac entry
	UsermacLabel []string `json:"usermac_label,omitempty"`
	// Username presented by the client
	Username *string `json:"username,omitempty"`
	// Vlan ID
	Vlan *string `json:"vlan,omitempty"`
	// Vlan source, e.g. "nactag", "usermac"
	VlanSource           *string        `json:"vlan_source,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for EventNacClient.
// It customizes the JSON marshaling process for EventNacClient objects.
func (e EventNacClient) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(e.toMap())
}

// toMap converts the EventNacClient object to a map representation for JSON marshaling.
func (e EventNacClient) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, e.AdditionalProperties)
	if e.Ap != nil {
		structMap["ap"] = e.Ap
	}
	if e.AuthType != nil {
		structMap["auth_type"] = e.AuthType
	}
	if e.Bssid != nil {
		structMap["bssid"] = e.Bssid
	}
	if e.DeviceMac != nil {
		structMap["device_mac"] = e.DeviceMac
	}
	if e.DryrunNacruleId != nil {
		structMap["dryrun_nacrule_id"] = e.DryrunNacruleId
	}
	if e.DryrunNacruleMatched != nil {
		structMap["dryrun_nacrule_matched"] = e.DryrunNacruleMatched
	}
	if e.IdpId != nil {
		structMap["idp_id"] = e.IdpId
	}
	if e.IdpRole != nil {
		structMap["idp_role"] = e.IdpRole
	}
	if e.IdpUsername != nil {
		structMap["idp_username"] = e.IdpUsername
	}
	if e.Mac != nil {
		structMap["mac"] = e.Mac
	}
	if e.MxedgeId != nil {
		structMap["mxedge_id"] = e.MxedgeId
	}
	if e.NacruleId != nil {
		structMap["nacrule_id"] = e.NacruleId
	}
	if e.NacruleMatched != nil {
		structMap["nacrule_matched"] = e.NacruleMatched
	}
	if e.NasVendor != nil {
		structMap["nas_vendor"] = e.NasVendor
	}
	if e.OrgId != nil {
		structMap["org_id"] = e.OrgId
	}
	if e.PortId != nil {
		structMap["port_id"] = e.PortId
	}
	if e.PortType != nil {
		structMap["port_type"] = e.PortType
	}
	if e.RandomMac != nil {
		structMap["random_mac"] = e.RandomMac
	}
	if e.RespAttrs != nil {
		structMap["resp_attrs"] = e.RespAttrs
	}
	if e.SiteId != nil {
		structMap["site_id"] = e.SiteId
	}
	if e.Ssid != nil {
		structMap["ssid"] = e.Ssid
	}
	if e.Timestamp != nil {
		structMap["timestamp"] = e.Timestamp
	}
	if e.Type != nil {
		structMap["type"] = e.Type
	}
	if e.UsermacLabel != nil {
		structMap["usermac_label"] = e.UsermacLabel
	}
	if e.Username != nil {
		structMap["username"] = e.Username
	}
	if e.Vlan != nil {
		structMap["vlan"] = e.Vlan
	}
	if e.VlanSource != nil {
		structMap["vlan_source"] = e.VlanSource
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for EventNacClient.
// It customizes the JSON unmarshaling process for EventNacClient objects.
func (e *EventNacClient) UnmarshalJSON(input []byte) error {
	var temp tempEventNacClient
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "ap", "auth_type", "bssid", "device_mac", "dryrun_nacrule_id", "dryrun_nacrule_matched", "idp_id", "idp_role", "idp_username", "mac", "mxedge_id", "nacrule_id", "nacrule_matched", "nas_vendor", "org_id", "port_id", "port_type", "random_mac", "resp_attrs", "site_id", "ssid", "timestamp", "type", "usermac_label", "username", "vlan", "vlan_source")
	if err != nil {
		return err
	}

	e.AdditionalProperties = additionalProperties
	e.Ap = temp.Ap
	e.AuthType = temp.AuthType
	e.Bssid = temp.Bssid
	e.DeviceMac = temp.DeviceMac
	e.DryrunNacruleId = temp.DryrunNacruleId
	e.DryrunNacruleMatched = temp.DryrunNacruleMatched
	e.IdpId = temp.IdpId
	e.IdpRole = temp.IdpRole
	e.IdpUsername = temp.IdpUsername
	e.Mac = temp.Mac
	e.MxedgeId = temp.MxedgeId
	e.NacruleId = temp.NacruleId
	e.NacruleMatched = temp.NacruleMatched
	e.NasVendor = temp.NasVendor
	e.OrgId = temp.OrgId
	e.PortId = temp.PortId
	e.PortType = temp.PortType
	e.RandomMac = temp.RandomMac
	e.RespAttrs = temp.RespAttrs
	e.SiteId = temp.SiteId
	e.Ssid = temp.Ssid
	e.Timestamp = temp.Timestamp
	e.Type = temp.Type
	e.UsermacLabel = temp.UsermacLabel
	e.Username = temp.Username
	e.Vlan = temp.Vlan
	e.VlanSource = temp.VlanSource
	return nil
}

// tempEventNacClient is a temporary struct used for validating the fields of EventNacClient.
type tempEventNacClient struct {
	Ap                   *string    `json:"ap,omitempty"`
	AuthType             *string    `json:"auth_type,omitempty"`
	Bssid                *string    `json:"bssid,omitempty"`
	DeviceMac            *string    `json:"device_mac,omitempty"`
	DryrunNacruleId      *uuid.UUID `json:"dryrun_nacrule_id,omitempty"`
	DryrunNacruleMatched *bool      `json:"dryrun_nacrule_matched,omitempty"`
	IdpId                *uuid.UUID `json:"idp_id,omitempty"`
	IdpRole              []string   `json:"idp_role,omitempty"`
	IdpUsername          *string    `json:"idp_username,omitempty"`
	Mac                  *string    `json:"mac,omitempty"`
	MxedgeId             *string    `json:"mxedge_id,omitempty"`
	NacruleId            *uuid.UUID `json:"nacrule_id,omitempty"`
	NacruleMatched       *bool      `json:"nacrule_matched,omitempty"`
	NasVendor            *string    `json:"nas_vendor,omitempty"`
	OrgId                *uuid.UUID `json:"org_id,omitempty"`
	PortId               *string    `json:"port_id,omitempty"`
	PortType             *string    `json:"port_type,omitempty"`
	RandomMac            *bool      `json:"random_mac,omitempty"`
	RespAttrs            []string   `json:"resp_attrs,omitempty"`
	SiteId               *uuid.UUID `json:"site_id,omitempty"`
	Ssid                 *string    `json:"ssid,omitempty"`
	Timestamp            *float64   `json:"timestamp,omitempty"`
	Type                 *string    `json:"type,omitempty"`
	UsermacLabel         []string   `json:"usermac_label,omitempty"`
	Username             *string    `json:"username,omitempty"`
	Vlan                 *string    `json:"vlan,omitempty"`
	VlanSource           *string    `json:"vlan_source,omitempty"`
}
