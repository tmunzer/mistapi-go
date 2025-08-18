// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// JunosOtherIpConfig represents a JunosOtherIpConfig struct.
// Optional, if it's required to have switch's L3 presence on a network/vlan
type JunosOtherIpConfig struct {
	// For EVPN, if anycast is desired
	EvpnAnycast *bool `json:"evpn_anycast,omitempty"`
	// Required if `type`==`static`
	Ip *string `json:"ip,omitempty"`
	// Required if `type6`==`static`
	Ip6 *string `json:"ip6,omitempty"`
	// Optional, `subnet` from `network` definition will be used if defined
	Netmask *string `json:"netmask,omitempty"`
	// Optional, `subnet` from `network` definition will be used if defined
	Netmask6 *string `json:"netmask6,omitempty"`
	// enum: `dhcp`, `static`
	Type *IpTypeEnum `json:"type,omitempty"`
	// enum: `autoconf`, `dhcp`, `disabled`, `static`
	Type6                *IpType6Enum           `json:"type6,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for JunosOtherIpConfig,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (j JunosOtherIpConfig) String() string {
	return fmt.Sprintf(
		"JunosOtherIpConfig[EvpnAnycast=%v, Ip=%v, Ip6=%v, Netmask=%v, Netmask6=%v, Type=%v, Type6=%v, AdditionalProperties=%v]",
		j.EvpnAnycast, j.Ip, j.Ip6, j.Netmask, j.Netmask6, j.Type, j.Type6, j.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for JunosOtherIpConfig.
// It customizes the JSON marshaling process for JunosOtherIpConfig objects.
func (j JunosOtherIpConfig) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(j.AdditionalProperties,
		"evpn_anycast", "ip", "ip6", "netmask", "netmask6", "type", "type6"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(j.toMap())
}

// toMap converts the JunosOtherIpConfig object to a map representation for JSON marshaling.
func (j JunosOtherIpConfig) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, j.AdditionalProperties)
	if j.EvpnAnycast != nil {
		structMap["evpn_anycast"] = j.EvpnAnycast
	}
	if j.Ip != nil {
		structMap["ip"] = j.Ip
	}
	if j.Ip6 != nil {
		structMap["ip6"] = j.Ip6
	}
	if j.Netmask != nil {
		structMap["netmask"] = j.Netmask
	}
	if j.Netmask6 != nil {
		structMap["netmask6"] = j.Netmask6
	}
	if j.Type != nil {
		structMap["type"] = j.Type
	}
	if j.Type6 != nil {
		structMap["type6"] = j.Type6
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for JunosOtherIpConfig.
// It customizes the JSON unmarshaling process for JunosOtherIpConfig objects.
func (j *JunosOtherIpConfig) UnmarshalJSON(input []byte) error {
	var temp tempJunosOtherIpConfig
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "evpn_anycast", "ip", "ip6", "netmask", "netmask6", "type", "type6")
	if err != nil {
		return err
	}
	j.AdditionalProperties = additionalProperties

	j.EvpnAnycast = temp.EvpnAnycast
	j.Ip = temp.Ip
	j.Ip6 = temp.Ip6
	j.Netmask = temp.Netmask
	j.Netmask6 = temp.Netmask6
	j.Type = temp.Type
	j.Type6 = temp.Type6
	return nil
}

// tempJunosOtherIpConfig is a temporary struct used for validating the fields of JunosOtherIpConfig.
type tempJunosOtherIpConfig struct {
	EvpnAnycast *bool        `json:"evpn_anycast,omitempty"`
	Ip          *string      `json:"ip,omitempty"`
	Ip6         *string      `json:"ip6,omitempty"`
	Netmask     *string      `json:"netmask,omitempty"`
	Netmask6    *string      `json:"netmask6,omitempty"`
	Type        *IpTypeEnum  `json:"type,omitempty"`
	Type6       *IpType6Enum `json:"type6,omitempty"`
}
