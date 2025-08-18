// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// WlanMistNac represents a WlanMistNac struct.
type WlanMistNac struct {
	// How frequently should interim accounting be reported, 60-65535. default is 0 (use one specified in Access-Accept request from Server). Very frequent messages can affect the performance of the radius server, 600 and up is recommended when enabled.
	AcctInterimInterval *int `json:"acct_interim_interval,omitempty"`
	// Radius auth session retries. Following fast timers are set if `fast_dot1x_timers` knob is enabled. "retries" are set to value of `auth_servers_timeout`. "max-requests" is also set when setting `auth_servers_retries` is set to default value to 3.
	AuthServersRetries *int `json:"auth_servers_retries,omitempty"`
	// Radius auth session timeout. Following fast timers are set if `fast_dot1x_timers` knob is enabled. "quite-period" and "transmit-period" are set to half the value of `auth_servers_timeout`. "supplicant-timeout" is also set when setting `auth_servers_timeout` is set to default value of 10.
	AuthServersTimeout *int `json:"auth_servers_timeout,omitempty"`
	// Allows a RADIUS server to dynamically modify the authorization status of a user session.
	CoaEnabled *bool `json:"coa_enabled,omitempty"`
	// the communication port used for “Change of Authorization” (CoA) messages
	CoaPort *int `json:"coa_port,omitempty"`
	// When enabled:
	// * `auth_servers` is ignored
	// * `acct_servers` is ignored
	// * `auth_servers_*` are ignored
	// * `coa_servers` is ignored
	// * `radsec` is ignored
	// * `coa_enabled` is assumed
	Enabled *bool `json:"enabled,omitempty"`
	// If set to true, sets default fast-timers with values calculated from `auth_servers_timeout` and `auth_server_retries`.
	FastDot1xTimers *bool `json:"fast_dot1x_timers,omitempty"`
	// Which network the mist nac server resides in
	Network Optional[string] `json:"network"`
	// In case there is a static IP for this network, we can specify it using source ip
	SourceIp             Optional[string]       `json:"source_ip"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for WlanMistNac,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WlanMistNac) String() string {
	return fmt.Sprintf(
		"WlanMistNac[AcctInterimInterval=%v, AuthServersRetries=%v, AuthServersTimeout=%v, CoaEnabled=%v, CoaPort=%v, Enabled=%v, FastDot1xTimers=%v, Network=%v, SourceIp=%v, AdditionalProperties=%v]",
		w.AcctInterimInterval, w.AuthServersRetries, w.AuthServersTimeout, w.CoaEnabled, w.CoaPort, w.Enabled, w.FastDot1xTimers, w.Network, w.SourceIp, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WlanMistNac.
// It customizes the JSON marshaling process for WlanMistNac objects.
func (w WlanMistNac) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(w.AdditionalProperties,
		"acct_interim_interval", "auth_servers_retries", "auth_servers_timeout", "coa_enabled", "coa_port", "enabled", "fast_dot1x_timers", "network", "source_ip"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(w.toMap())
}

// toMap converts the WlanMistNac object to a map representation for JSON marshaling.
func (w WlanMistNac) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, w.AdditionalProperties)
	if w.AcctInterimInterval != nil {
		structMap["acct_interim_interval"] = w.AcctInterimInterval
	}
	if w.AuthServersRetries != nil {
		structMap["auth_servers_retries"] = w.AuthServersRetries
	}
	if w.AuthServersTimeout != nil {
		structMap["auth_servers_timeout"] = w.AuthServersTimeout
	}
	if w.CoaEnabled != nil {
		structMap["coa_enabled"] = w.CoaEnabled
	}
	if w.CoaPort != nil {
		structMap["coa_port"] = w.CoaPort
	}
	if w.Enabled != nil {
		structMap["enabled"] = w.Enabled
	}
	if w.FastDot1xTimers != nil {
		structMap["fast_dot1x_timers"] = w.FastDot1xTimers
	}
	if w.Network.IsValueSet() {
		if w.Network.Value() != nil {
			structMap["network"] = w.Network.Value()
		} else {
			structMap["network"] = nil
		}
	}
	if w.SourceIp.IsValueSet() {
		if w.SourceIp.Value() != nil {
			structMap["source_ip"] = w.SourceIp.Value()
		} else {
			structMap["source_ip"] = nil
		}
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WlanMistNac.
// It customizes the JSON unmarshaling process for WlanMistNac objects.
func (w *WlanMistNac) UnmarshalJSON(input []byte) error {
	var temp tempWlanMistNac
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "acct_interim_interval", "auth_servers_retries", "auth_servers_timeout", "coa_enabled", "coa_port", "enabled", "fast_dot1x_timers", "network", "source_ip")
	if err != nil {
		return err
	}
	w.AdditionalProperties = additionalProperties

	w.AcctInterimInterval = temp.AcctInterimInterval
	w.AuthServersRetries = temp.AuthServersRetries
	w.AuthServersTimeout = temp.AuthServersTimeout
	w.CoaEnabled = temp.CoaEnabled
	w.CoaPort = temp.CoaPort
	w.Enabled = temp.Enabled
	w.FastDot1xTimers = temp.FastDot1xTimers
	w.Network = temp.Network
	w.SourceIp = temp.SourceIp
	return nil
}

// tempWlanMistNac is a temporary struct used for validating the fields of WlanMistNac.
type tempWlanMistNac struct {
	AcctInterimInterval *int             `json:"acct_interim_interval,omitempty"`
	AuthServersRetries  *int             `json:"auth_servers_retries,omitempty"`
	AuthServersTimeout  *int             `json:"auth_servers_timeout,omitempty"`
	CoaEnabled          *bool            `json:"coa_enabled,omitempty"`
	CoaPort             *int             `json:"coa_port,omitempty"`
	Enabled             *bool            `json:"enabled,omitempty"`
	FastDot1xTimers     *bool            `json:"fast_dot1x_timers,omitempty"`
	Network             Optional[string] `json:"network"`
	SourceIp            Optional[string] `json:"source_ip"`
}
