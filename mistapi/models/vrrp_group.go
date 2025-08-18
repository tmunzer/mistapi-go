// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// VrrpGroup represents a VrrpGroup struct.
// Junos VRRP group
type VrrpGroup struct {
	// If `auth_type`==`md5`
	AuthKey *string `json:"auth_key,omitempty"`
	// If `auth_type`==`simple`
	AuthPassword *string `json:"auth_password,omitempty"`
	// enum: `md5`, `simple`
	AuthType *VrrpGroupAuthTypeEnum `json:"auth_type,omitempty"`
	// Property key is the network name
	Networks             map[string]VrrpGroupNetwork `json:"networks,omitempty"`
	AdditionalProperties map[string]interface{}      `json:"_"`
}

// String implements the fmt.Stringer interface for VrrpGroup,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (v VrrpGroup) String() string {
	return fmt.Sprintf(
		"VrrpGroup[AuthKey=%v, AuthPassword=%v, AuthType=%v, Networks=%v, AdditionalProperties=%v]",
		v.AuthKey, v.AuthPassword, v.AuthType, v.Networks, v.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for VrrpGroup.
// It customizes the JSON marshaling process for VrrpGroup objects.
func (v VrrpGroup) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(v.AdditionalProperties,
		"auth_key", "auth_password", "auth_type", "networks"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(v.toMap())
}

// toMap converts the VrrpGroup object to a map representation for JSON marshaling.
func (v VrrpGroup) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, v.AdditionalProperties)
	if v.AuthKey != nil {
		structMap["auth_key"] = v.AuthKey
	}
	if v.AuthPassword != nil {
		structMap["auth_password"] = v.AuthPassword
	}
	if v.AuthType != nil {
		structMap["auth_type"] = v.AuthType
	}
	if v.Networks != nil {
		structMap["networks"] = v.Networks
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for VrrpGroup.
// It customizes the JSON unmarshaling process for VrrpGroup objects.
func (v *VrrpGroup) UnmarshalJSON(input []byte) error {
	var temp tempVrrpGroup
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "auth_key", "auth_password", "auth_type", "networks")
	if err != nil {
		return err
	}
	v.AdditionalProperties = additionalProperties

	v.AuthKey = temp.AuthKey
	v.AuthPassword = temp.AuthPassword
	v.AuthType = temp.AuthType
	v.Networks = temp.Networks
	return nil
}

// tempVrrpGroup is a temporary struct used for validating the fields of VrrpGroup.
type tempVrrpGroup struct {
	AuthKey      *string                     `json:"auth_key,omitempty"`
	AuthPassword *string                     `json:"auth_password,omitempty"`
	AuthType     *VrrpGroupAuthTypeEnum      `json:"auth_type,omitempty"`
	Networks     map[string]VrrpGroupNetwork `json:"networks,omitempty"`
}
