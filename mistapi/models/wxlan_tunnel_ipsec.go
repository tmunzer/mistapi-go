// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// WxlanTunnelIpsec represents a WxlanTunnelIpsec struct.
// IPSec-related configurations; requires DMVPN be enabled
type WxlanTunnelIpsec struct {
    // Whether ipsec is enabled, requires DMVPN be enabled
    Enabled              *bool                  `json:"enabled,omitempty"`
    // IPSec pre-shared key
    Psk                  string                 `json:"psk"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for WxlanTunnelIpsec,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WxlanTunnelIpsec) String() string {
    return fmt.Sprintf(
    	"WxlanTunnelIpsec[Enabled=%v, Psk=%v, AdditionalProperties=%v]",
    	w.Enabled, w.Psk, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WxlanTunnelIpsec.
// It customizes the JSON marshaling process for WxlanTunnelIpsec objects.
func (w WxlanTunnelIpsec) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(w.AdditionalProperties,
        "enabled", "psk"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WxlanTunnelIpsec object to a map representation for JSON marshaling.
func (w WxlanTunnelIpsec) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, w.AdditionalProperties)
    if w.Enabled != nil {
        structMap["enabled"] = w.Enabled
    }
    structMap["psk"] = w.Psk
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WxlanTunnelIpsec.
// It customizes the JSON unmarshaling process for WxlanTunnelIpsec objects.
func (w *WxlanTunnelIpsec) UnmarshalJSON(input []byte) error {
    var temp tempWxlanTunnelIpsec
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "enabled", "psk")
    if err != nil {
    	return err
    }
    w.AdditionalProperties = additionalProperties
    
    w.Enabled = temp.Enabled
    w.Psk = *temp.Psk
    return nil
}

// tempWxlanTunnelIpsec is a temporary struct used for validating the fields of WxlanTunnelIpsec.
type tempWxlanTunnelIpsec  struct {
    Enabled *bool   `json:"enabled,omitempty"`
    Psk     *string `json:"psk"`
}

func (w *tempWxlanTunnelIpsec) validate() error {
    var errs []string
    if w.Psk == nil {
        errs = append(errs, "required field `psk` is missing for type `wxlan_tunnel_ipsec`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
