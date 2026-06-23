// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// TunnelConfigIpsecProposal represents a TunnelConfigIpsecProposal struct.
// IPsec proposal settings for custom IPsec tunnels
type TunnelConfigIpsecProposal struct {
	// enum: `md5`, `sha1`, `sha2`
	AuthAlgo *TunnelConfigAuthAlgoEnum `json:"auth_algo,omitempty"`
	// Only if `provider`==`custom-ipsec`. Diffie-Hellman group for IPsec phase 2. enum: `1`, `14`, `15`, `16`, `19`, `2`, `20`, `21`, `24`, `5`. `14` is the default 2048-bit group; `19`, `20`, `21`, and `24` are ECP groups
	DhGroup *TunnelConfigDhGroupEnum `json:"dh_group,omitempty"`
	// enum: `3des`, `aes128`, `aes256`, `aes_gcm128`, `aes_gcm256`
	EncAlgo              Optional[TunnelConfigEncAlgoEnum] `json:"enc_algo"`
	AdditionalProperties map[string]interface{}            `json:"_"`
}

// String implements the fmt.Stringer interface for TunnelConfigIpsecProposal,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (t TunnelConfigIpsecProposal) String() string {
	return fmt.Sprintf(
		"TunnelConfigIpsecProposal[AuthAlgo=%v, DhGroup=%v, EncAlgo=%v, AdditionalProperties=%v]",
		t.AuthAlgo, t.DhGroup, t.EncAlgo, t.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for TunnelConfigIpsecProposal.
// It customizes the JSON marshaling process for TunnelConfigIpsecProposal objects.
func (t TunnelConfigIpsecProposal) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(t.AdditionalProperties,
		"auth_algo", "dh_group", "enc_algo"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(t.toMap())
}

// toMap converts the TunnelConfigIpsecProposal object to a map representation for JSON marshaling.
func (t TunnelConfigIpsecProposal) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, t.AdditionalProperties)
	if t.AuthAlgo != nil {
		structMap["auth_algo"] = t.AuthAlgo
	}
	if t.DhGroup != nil {
		structMap["dh_group"] = t.DhGroup
	}
	if t.EncAlgo.IsValueSet() {
		if t.EncAlgo.Value() != nil {
			structMap["enc_algo"] = t.EncAlgo.Value()
		} else {
			structMap["enc_algo"] = nil
		}
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for TunnelConfigIpsecProposal.
// It customizes the JSON unmarshaling process for TunnelConfigIpsecProposal objects.
func (t *TunnelConfigIpsecProposal) UnmarshalJSON(input []byte) error {
	var temp tempTunnelConfigIpsecProposal
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "auth_algo", "dh_group", "enc_algo")
	if err != nil {
		return err
	}
	t.AdditionalProperties = additionalProperties

	t.AuthAlgo = temp.AuthAlgo
	t.DhGroup = temp.DhGroup
	t.EncAlgo = temp.EncAlgo
	return nil
}

// tempTunnelConfigIpsecProposal is a temporary struct used for validating the fields of TunnelConfigIpsecProposal.
type tempTunnelConfigIpsecProposal struct {
	AuthAlgo *TunnelConfigAuthAlgoEnum         `json:"auth_algo,omitempty"`
	DhGroup  *TunnelConfigDhGroupEnum          `json:"dh_group,omitempty"`
	EncAlgo  Optional[TunnelConfigEncAlgoEnum] `json:"enc_algo"`
}
