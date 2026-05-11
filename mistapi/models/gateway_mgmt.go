// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// GatewayMgmt represents a GatewayMgmt struct.
// Gateway Management settings
type GatewayMgmt struct {
	// For SSR only, as direct root access is not allowed
	AdminSshkeys []string    `json:"admin_sshkeys,omitempty"`
	AppProbing   *AppProbing `json:"app_probing,omitempty"`
	// Consumes uplink bandwidth, requires WA license
	AppUsage            *bool                           `json:"app_usage,omitempty"`
	AutoSignatureUpdate *GatewayMgmtAutoSignatureUpdate `json:"auto_signature_update,omitempty"`
	// Rollback timer for commit confirmed
	ConfigRevertTimer *int `json:"config_revert_timer,omitempty"`
	// For SSR and SRX, disable console port
	DisableConsole *bool `json:"disable_console,omitempty"`
	// For SSR and SRX, disable management interface
	DisableOob *bool `json:"disable_oob,omitempty"`
	// For SSR and SRX, disable usb interface
	DisableUsb   *bool    `json:"disable_usb,omitempty"`
	FipsEnabled  *bool    `json:"fips_enabled,omitempty"`
	ProbeHosts   []string `json:"probe_hosts,omitempty"`
	ProbeHostsv6 []string `json:"probe_hostsv6,omitempty"`
	// Restrict inbound-traffic to host
	// when enabled, all traffic that is not essential to our operation will be dropped
	// e.g. ntp / dns / traffic to mist will be allowed by default, if dhcpd is enabled, we'll make sure it works
	ProtectRe *ProtectRe `json:"protect_re,omitempty"`
	// SRX only
	RootPassword               *string                `json:"root_password,omitempty"`
	SecurityLogSourceAddress   *string                `json:"security_log_source_address,omitempty"`
	SecurityLogSourceInterface *string                `json:"security_log_source_interface,omitempty"`
	AdditionalProperties       map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for GatewayMgmt,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (g GatewayMgmt) String() string {
	return fmt.Sprintf(
		"GatewayMgmt[AdminSshkeys=%v, AppProbing=%v, AppUsage=%v, AutoSignatureUpdate=%v, ConfigRevertTimer=%v, DisableConsole=%v, DisableOob=%v, DisableUsb=%v, FipsEnabled=%v, ProbeHosts=%v, ProbeHostsv6=%v, ProtectRe=%v, RootPassword=%v, SecurityLogSourceAddress=%v, SecurityLogSourceInterface=%v, AdditionalProperties=%v]",
		g.AdminSshkeys, g.AppProbing, g.AppUsage, g.AutoSignatureUpdate, g.ConfigRevertTimer, g.DisableConsole, g.DisableOob, g.DisableUsb, g.FipsEnabled, g.ProbeHosts, g.ProbeHostsv6, g.ProtectRe, g.RootPassword, g.SecurityLogSourceAddress, g.SecurityLogSourceInterface, g.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for GatewayMgmt.
// It customizes the JSON marshaling process for GatewayMgmt objects.
func (g GatewayMgmt) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(g.AdditionalProperties,
		"admin_sshkeys", "app_probing", "app_usage", "auto_signature_update", "config_revert_timer", "disable_console", "disable_oob", "disable_usb", "fips_enabled", "probe_hosts", "probe_hostsv6", "protect_re", "root_password", "security_log_source_address", "security_log_source_interface"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(g.toMap())
}

// toMap converts the GatewayMgmt object to a map representation for JSON marshaling.
func (g GatewayMgmt) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, g.AdditionalProperties)
	if g.AdminSshkeys != nil {
		structMap["admin_sshkeys"] = g.AdminSshkeys
	}
	if g.AppProbing != nil {
		structMap["app_probing"] = g.AppProbing.toMap()
	}
	if g.AppUsage != nil {
		structMap["app_usage"] = g.AppUsage
	}
	if g.AutoSignatureUpdate != nil {
		structMap["auto_signature_update"] = g.AutoSignatureUpdate.toMap()
	}
	if g.ConfigRevertTimer != nil {
		structMap["config_revert_timer"] = g.ConfigRevertTimer
	}
	if g.DisableConsole != nil {
		structMap["disable_console"] = g.DisableConsole
	}
	if g.DisableOob != nil {
		structMap["disable_oob"] = g.DisableOob
	}
	if g.DisableUsb != nil {
		structMap["disable_usb"] = g.DisableUsb
	}
	if g.FipsEnabled != nil {
		structMap["fips_enabled"] = g.FipsEnabled
	}
	if g.ProbeHosts != nil {
		structMap["probe_hosts"] = g.ProbeHosts
	}
	if g.ProbeHostsv6 != nil {
		structMap["probe_hostsv6"] = g.ProbeHostsv6
	}
	if g.ProtectRe != nil {
		structMap["protect_re"] = g.ProtectRe.toMap()
	}
	if g.RootPassword != nil {
		structMap["root_password"] = g.RootPassword
	}
	if g.SecurityLogSourceAddress != nil {
		structMap["security_log_source_address"] = g.SecurityLogSourceAddress
	}
	if g.SecurityLogSourceInterface != nil {
		structMap["security_log_source_interface"] = g.SecurityLogSourceInterface
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GatewayMgmt.
// It customizes the JSON unmarshaling process for GatewayMgmt objects.
func (g *GatewayMgmt) UnmarshalJSON(input []byte) error {
	var temp tempGatewayMgmt
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "admin_sshkeys", "app_probing", "app_usage", "auto_signature_update", "config_revert_timer", "disable_console", "disable_oob", "disable_usb", "fips_enabled", "probe_hosts", "probe_hostsv6", "protect_re", "root_password", "security_log_source_address", "security_log_source_interface")
	if err != nil {
		return err
	}
	g.AdditionalProperties = additionalProperties

	g.AdminSshkeys = temp.AdminSshkeys
	g.AppProbing = temp.AppProbing
	g.AppUsage = temp.AppUsage
	g.AutoSignatureUpdate = temp.AutoSignatureUpdate
	g.ConfigRevertTimer = temp.ConfigRevertTimer
	g.DisableConsole = temp.DisableConsole
	g.DisableOob = temp.DisableOob
	g.DisableUsb = temp.DisableUsb
	g.FipsEnabled = temp.FipsEnabled
	g.ProbeHosts = temp.ProbeHosts
	g.ProbeHostsv6 = temp.ProbeHostsv6
	g.ProtectRe = temp.ProtectRe
	g.RootPassword = temp.RootPassword
	g.SecurityLogSourceAddress = temp.SecurityLogSourceAddress
	g.SecurityLogSourceInterface = temp.SecurityLogSourceInterface
	return nil
}

// tempGatewayMgmt is a temporary struct used for validating the fields of GatewayMgmt.
type tempGatewayMgmt struct {
	AdminSshkeys               []string                        `json:"admin_sshkeys,omitempty"`
	AppProbing                 *AppProbing                     `json:"app_probing,omitempty"`
	AppUsage                   *bool                           `json:"app_usage,omitempty"`
	AutoSignatureUpdate        *GatewayMgmtAutoSignatureUpdate `json:"auto_signature_update,omitempty"`
	ConfigRevertTimer          *int                            `json:"config_revert_timer,omitempty"`
	DisableConsole             *bool                           `json:"disable_console,omitempty"`
	DisableOob                 *bool                           `json:"disable_oob,omitempty"`
	DisableUsb                 *bool                           `json:"disable_usb,omitempty"`
	FipsEnabled                *bool                           `json:"fips_enabled,omitempty"`
	ProbeHosts                 []string                        `json:"probe_hosts,omitempty"`
	ProbeHostsv6               []string                        `json:"probe_hostsv6,omitempty"`
	ProtectRe                  *ProtectRe                      `json:"protect_re,omitempty"`
	RootPassword               *string                         `json:"root_password,omitempty"`
	SecurityLogSourceAddress   *string                         `json:"security_log_source_address,omitempty"`
	SecurityLogSourceInterface *string                         `json:"security_log_source_interface,omitempty"`
}
