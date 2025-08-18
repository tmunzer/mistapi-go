// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// ApClientBridgeAuth represents a ApClientBridgeAuth struct.
type ApClientBridgeAuth struct {
	Psk *string `json:"psk,omitempty"`
	// wpa2-AES/CCMPp is assumed when `type`==`psk`. enum: `open`, `psk`
	Type                 *ApClientBridgeAuthTypeEnum `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}      `json:"_"`
}

// String implements the fmt.Stringer interface for ApClientBridgeAuth,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a ApClientBridgeAuth) String() string {
	return fmt.Sprintf(
		"ApClientBridgeAuth[Psk=%v, Type=%v, AdditionalProperties=%v]",
		a.Psk, a.Type, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ApClientBridgeAuth.
// It customizes the JSON marshaling process for ApClientBridgeAuth objects.
func (a ApClientBridgeAuth) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(a.AdditionalProperties,
		"psk", "type"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(a.toMap())
}

// toMap converts the ApClientBridgeAuth object to a map representation for JSON marshaling.
func (a ApClientBridgeAuth) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, a.AdditionalProperties)
	if a.Psk != nil {
		structMap["psk"] = a.Psk
	}
	if a.Type != nil {
		structMap["type"] = a.Type
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ApClientBridgeAuth.
// It customizes the JSON unmarshaling process for ApClientBridgeAuth objects.
func (a *ApClientBridgeAuth) UnmarshalJSON(input []byte) error {
	var temp tempApClientBridgeAuth
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "psk", "type")
	if err != nil {
		return err
	}
	a.AdditionalProperties = additionalProperties

	a.Psk = temp.Psk
	a.Type = temp.Type
	return nil
}

// tempApClientBridgeAuth is a temporary struct used for validating the fields of ApClientBridgeAuth.
type tempApClientBridgeAuth struct {
	Psk  *string                     `json:"psk,omitempty"`
	Type *ApClientBridgeAuthTypeEnum `json:"type,omitempty"`
}
