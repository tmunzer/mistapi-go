// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// OrgSettingGatewayMgmtHostInPolicy represents a OrgSettingGatewayMgmtHostInPolicy struct.
type OrgSettingGatewayMgmtHostInPolicy struct {
	Tenants              []string               `json:"tenants,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for OrgSettingGatewayMgmtHostInPolicy,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (o OrgSettingGatewayMgmtHostInPolicy) String() string {
	return fmt.Sprintf(
		"OrgSettingGatewayMgmtHostInPolicy[Tenants=%v, AdditionalProperties=%v]",
		o.Tenants, o.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for OrgSettingGatewayMgmtHostInPolicy.
// It customizes the JSON marshaling process for OrgSettingGatewayMgmtHostInPolicy objects.
func (o OrgSettingGatewayMgmtHostInPolicy) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(o.AdditionalProperties,
		"tenants"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(o.toMap())
}

// toMap converts the OrgSettingGatewayMgmtHostInPolicy object to a map representation for JSON marshaling.
func (o OrgSettingGatewayMgmtHostInPolicy) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, o.AdditionalProperties)
	if o.Tenants != nil {
		structMap["tenants"] = o.Tenants
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgSettingGatewayMgmtHostInPolicy.
// It customizes the JSON unmarshaling process for OrgSettingGatewayMgmtHostInPolicy objects.
func (o *OrgSettingGatewayMgmtHostInPolicy) UnmarshalJSON(input []byte) error {
	var temp tempOrgSettingGatewayMgmtHostInPolicy
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "tenants")
	if err != nil {
		return err
	}
	o.AdditionalProperties = additionalProperties

	o.Tenants = temp.Tenants
	return nil
}

// tempOrgSettingGatewayMgmtHostInPolicy is a temporary struct used for validating the fields of OrgSettingGatewayMgmtHostInPolicy.
type tempOrgSettingGatewayMgmtHostInPolicy struct {
	Tenants []string `json:"tenants,omitempty"`
}
