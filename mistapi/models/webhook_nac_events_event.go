package models

import (
	"encoding/json"
	"github.com/google/uuid"
)

// WebhookNacEventsEvent represents a WebhookNacEventsEvent struct.
type WebhookNacEventsEvent struct {
	// random mac
	Ap *string `json:"ap,omitempty"`
	// authentication type, e.g. "eap-tls", "peap-tls", "eap-ttls", "eap-teap", "mab", "psk", "device-auth"
	AuthType *string `json:"auth_type,omitempty"`
	// BSSID
	Bssid *string `json:"bssid,omitempty"`
	// NAC Policy Dry Run Rule ID, if present and matched
	DryrunNacruleId *uuid.UUID `json:"dryrun_nacrule_id,omitempty"`
	// True - if dryrun rule present and matched with priority, False - if not matched or not present
	DryrunNacruleMatched *bool `json:"dryrun_nacrule_matched,omitempty"`
	// SSO ID, if present and used
	IdpId *uuid.UUID `json:"idp_id,omitempty"`
	// IDP returned roles/groups for the user
	IdpRole []string `json:"idp_role,omitempty"`
	// MAC address
	Mac *string `json:"mac,omitempty"`
	// NAC Policy Rule ID, if matched
	NacruleId *uuid.UUID `json:"nacrule_id,omitempty"`
	// NAC Policy Rule Matched
	NacruleMatched *bool `json:"nacrule_matched,omitempty"`
	// vendor of NAS device
	NasVendor *string    `json:"nas_vendor,omitempty"`
	OrgId     *uuid.UUID `json:"org_id,omitempty"`
	// AP MAC
	RandomMac *bool `json:"random_mac,omitempty"`
	// Radius attributes returned by NAC to NAS Devive
	RespAttrs []string `json:"resp_attrs,omitempty"`
	// site id if assigned, null if not assigned
	SiteId *uuid.UUID `json:"site_id,omitempty"`
	// SSID
	Ssid *string `json:"ssid,omitempty"`
	// start time, in epoch
	Timestamp *float64 `json:"timestamp,omitempty"`
	// event type, e.g. NAC_CLIENT_PERMIT
	Type *string `json:"type,omitempty"`
	// Username presented by the client
	Username *string `json:"username,omitempty"`
	// Vlan ID
	Vlan                 *string        `json:"vlan,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for WebhookNacEventsEvent.
// It customizes the JSON marshaling process for WebhookNacEventsEvent objects.
func (w WebhookNacEventsEvent) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(w.toMap())
}

// toMap converts the WebhookNacEventsEvent object to a map representation for JSON marshaling.
func (w WebhookNacEventsEvent) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, w.AdditionalProperties)
	if w.Ap != nil {
		structMap["ap"] = w.Ap
	}
	if w.AuthType != nil {
		structMap["auth_type"] = w.AuthType
	}
	if w.Bssid != nil {
		structMap["bssid"] = w.Bssid
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
	additionalProperties, err := UnmarshalAdditionalProperties(input, "ap", "auth_type", "bssid", "dryrun_nacrule_id", "dryrun_nacrule_matched", "idp_id", "idp_role", "mac", "nacrule_id", "nacrule_matched", "nas_vendor", "org_id", "random_mac", "resp_attrs", "site_id", "ssid", "timestamp", "type", "username", "vlan")
	if err != nil {
		return err
	}

	w.AdditionalProperties = additionalProperties
	w.Ap = temp.Ap
	w.AuthType = temp.AuthType
	w.Bssid = temp.Bssid
	w.DryrunNacruleId = temp.DryrunNacruleId
	w.DryrunNacruleMatched = temp.DryrunNacruleMatched
	w.IdpId = temp.IdpId
	w.IdpRole = temp.IdpRole
	w.Mac = temp.Mac
	w.NacruleId = temp.NacruleId
	w.NacruleMatched = temp.NacruleMatched
	w.NasVendor = temp.NasVendor
	w.OrgId = temp.OrgId
	w.RandomMac = temp.RandomMac
	w.RespAttrs = temp.RespAttrs
	w.SiteId = temp.SiteId
	w.Ssid = temp.Ssid
	w.Timestamp = temp.Timestamp
	w.Type = temp.Type
	w.Username = temp.Username
	w.Vlan = temp.Vlan
	return nil
}

// tempWebhookNacEventsEvent is a temporary struct used for validating the fields of WebhookNacEventsEvent.
type tempWebhookNacEventsEvent struct {
	Ap                   *string    `json:"ap,omitempty"`
	AuthType             *string    `json:"auth_type,omitempty"`
	Bssid                *string    `json:"bssid,omitempty"`
	DryrunNacruleId      *uuid.UUID `json:"dryrun_nacrule_id,omitempty"`
	DryrunNacruleMatched *bool      `json:"dryrun_nacrule_matched,omitempty"`
	IdpId                *uuid.UUID `json:"idp_id,omitempty"`
	IdpRole              []string   `json:"idp_role,omitempty"`
	Mac                  *string    `json:"mac,omitempty"`
	NacruleId            *uuid.UUID `json:"nacrule_id,omitempty"`
	NacruleMatched       *bool      `json:"nacrule_matched,omitempty"`
	NasVendor            *string    `json:"nas_vendor,omitempty"`
	OrgId                *uuid.UUID `json:"org_id,omitempty"`
	RandomMac            *bool      `json:"random_mac,omitempty"`
	RespAttrs            []string   `json:"resp_attrs,omitempty"`
	SiteId               *uuid.UUID `json:"site_id,omitempty"`
	Ssid                 *string    `json:"ssid,omitempty"`
	Timestamp            *float64   `json:"timestamp,omitempty"`
	Type                 *string    `json:"type,omitempty"`
	Username             *string    `json:"username,omitempty"`
	Vlan                 *string    `json:"vlan,omitempty"`
}
