// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// NacPortalGuestPortal represents a NacPortalGuestPortal struct.
// Guest portal configuration when `type`==`guest_portal`. If
// * `auth`==`none`, the user is presented with a terms of service and can click and continue.
// * `auth`==`external`, the user is redirected to an external URL for authentication.
// * `auth`==`multi`, the user is presented with a choice of authentication methods:
// - social logins: facebook / google / amazon / microsoft / azure
// - sponsor
// - sms: supported provider: twillio
// - email
// - sso
// - userpass: pre created guest list
type NacPortalGuestPortal struct {
	// Guest portal authentication type. enum: `external`, `multi`, `none`
	Auth *NacPortalGuestPortalAuthEnum `json:"auth,omitempty"`
	// If `auth`==`none` or `auth`==`multi`, whether to expire the guest after a certain time
	Expire *int `json:"expire,omitempty"`
	// If `auth`==`external`, the URL to redirect the user to for authentication
	ExternalPortalUrl *string `json:"external_portal_url,omitempty"`
	ForceReconnect    *bool   `json:"force_reconnect,omitempty"`
	// If `auth`==`none` or `auth`==`multi`, whether to forward the user to the guest portal after authentication
	Forward *bool `json:"forward,omitempty"`
	// If `auth`==`none` or `auth`==`multi`, URL to forward the user to after authentication
	ForwardUrl *string `json:"forward_url,omitempty"`
	// List of hostnames without http(s):// (matched by substring)
	PortalAllowedHostnames []string `json:"portal_allowed_hostnames,omitempty"`
	// List of CIDRs
	PortalAllowedSubnets []string `json:"portal_allowed_subnets,omitempty"`
	// List of hostnames without http(s):// (matched by substring), this takes precedence over portal_allowed_hostnames
	PortalDeniedHostnames []string `json:"portal_denied_hostnames,omitempty"`
	// If `auth`==`none` or `auth`==`multi`, whether to show the privacy policy
	Privacy              *bool                  `json:"privacy,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for NacPortalGuestPortal,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (n NacPortalGuestPortal) String() string {
	return fmt.Sprintf(
		"NacPortalGuestPortal[Auth=%v, Expire=%v, ExternalPortalUrl=%v, ForceReconnect=%v, Forward=%v, ForwardUrl=%v, PortalAllowedHostnames=%v, PortalAllowedSubnets=%v, PortalDeniedHostnames=%v, Privacy=%v, AdditionalProperties=%v]",
		n.Auth, n.Expire, n.ExternalPortalUrl, n.ForceReconnect, n.Forward, n.ForwardUrl, n.PortalAllowedHostnames, n.PortalAllowedSubnets, n.PortalDeniedHostnames, n.Privacy, n.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for NacPortalGuestPortal.
// It customizes the JSON marshaling process for NacPortalGuestPortal objects.
func (n NacPortalGuestPortal) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(n.AdditionalProperties,
		"auth", "expire", "external_portal_url", "force_reconnect", "forward", "forward_url", "portal_allowed_hostnames", "portal_allowed_subnets", "portal_denied_hostnames", "privacy"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(n.toMap())
}

// toMap converts the NacPortalGuestPortal object to a map representation for JSON marshaling.
func (n NacPortalGuestPortal) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, n.AdditionalProperties)
	if n.Auth != nil {
		structMap["auth"] = n.Auth
	}
	if n.Expire != nil {
		structMap["expire"] = n.Expire
	}
	if n.ExternalPortalUrl != nil {
		structMap["external_portal_url"] = n.ExternalPortalUrl
	}
	if n.ForceReconnect != nil {
		structMap["force_reconnect"] = n.ForceReconnect
	}
	if n.Forward != nil {
		structMap["forward"] = n.Forward
	}
	if n.ForwardUrl != nil {
		structMap["forward_url"] = n.ForwardUrl
	}
	if n.PortalAllowedHostnames != nil {
		structMap["portal_allowed_hostnames"] = n.PortalAllowedHostnames
	}
	if n.PortalAllowedSubnets != nil {
		structMap["portal_allowed_subnets"] = n.PortalAllowedSubnets
	}
	if n.PortalDeniedHostnames != nil {
		structMap["portal_denied_hostnames"] = n.PortalDeniedHostnames
	}
	if n.Privacy != nil {
		structMap["privacy"] = n.Privacy
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for NacPortalGuestPortal.
// It customizes the JSON unmarshaling process for NacPortalGuestPortal objects.
func (n *NacPortalGuestPortal) UnmarshalJSON(input []byte) error {
	var temp tempNacPortalGuestPortal
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "auth", "expire", "external_portal_url", "force_reconnect", "forward", "forward_url", "portal_allowed_hostnames", "portal_allowed_subnets", "portal_denied_hostnames", "privacy")
	if err != nil {
		return err
	}
	n.AdditionalProperties = additionalProperties

	n.Auth = temp.Auth
	n.Expire = temp.Expire
	n.ExternalPortalUrl = temp.ExternalPortalUrl
	n.ForceReconnect = temp.ForceReconnect
	n.Forward = temp.Forward
	n.ForwardUrl = temp.ForwardUrl
	n.PortalAllowedHostnames = temp.PortalAllowedHostnames
	n.PortalAllowedSubnets = temp.PortalAllowedSubnets
	n.PortalDeniedHostnames = temp.PortalDeniedHostnames
	n.Privacy = temp.Privacy
	return nil
}

// tempNacPortalGuestPortal is a temporary struct used for validating the fields of NacPortalGuestPortal.
type tempNacPortalGuestPortal struct {
	Auth                   *NacPortalGuestPortalAuthEnum `json:"auth,omitempty"`
	Expire                 *int                          `json:"expire,omitempty"`
	ExternalPortalUrl      *string                       `json:"external_portal_url,omitempty"`
	ForceReconnect         *bool                         `json:"force_reconnect,omitempty"`
	Forward                *bool                         `json:"forward,omitempty"`
	ForwardUrl             *string                       `json:"forward_url,omitempty"`
	PortalAllowedHostnames []string                      `json:"portal_allowed_hostnames,omitempty"`
	PortalAllowedSubnets   []string                      `json:"portal_allowed_subnets,omitempty"`
	PortalDeniedHostnames  []string                      `json:"portal_denied_hostnames,omitempty"`
	Privacy                *bool                         `json:"privacy,omitempty"`
}
