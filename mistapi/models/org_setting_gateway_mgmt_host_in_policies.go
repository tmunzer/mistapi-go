// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// OrgSettingGatewayMgmtHostInPolicies represents a OrgSettingGatewayMgmtHostInPolicies struct.
type OrgSettingGatewayMgmtHostInPolicies struct {
	Icmp                 *OrgSettingGatewayMgmtHostInPolicy `json:"icmp,omitempty"`
	Snmp                 *OrgSettingGatewayMgmtHostInPolicy `json:"snmp,omitempty"`
	AdditionalProperties map[string]interface{}             `json:"_"`
}

// String implements the fmt.Stringer interface for OrgSettingGatewayMgmtHostInPolicies,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (o OrgSettingGatewayMgmtHostInPolicies) String() string {
	return fmt.Sprintf(
		"OrgSettingGatewayMgmtHostInPolicies[Icmp=%v, Snmp=%v, AdditionalProperties=%v]",
		o.Icmp, o.Snmp, o.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for OrgSettingGatewayMgmtHostInPolicies.
// It customizes the JSON marshaling process for OrgSettingGatewayMgmtHostInPolicies objects.
func (o OrgSettingGatewayMgmtHostInPolicies) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(o.AdditionalProperties,
		"icmp", "snmp"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(o.toMap())
}

// toMap converts the OrgSettingGatewayMgmtHostInPolicies object to a map representation for JSON marshaling.
func (o OrgSettingGatewayMgmtHostInPolicies) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, o.AdditionalProperties)
	if o.Icmp != nil {
		structMap["icmp"] = o.Icmp.toMap()
	}
	if o.Snmp != nil {
		structMap["snmp"] = o.Snmp.toMap()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgSettingGatewayMgmtHostInPolicies.
// It customizes the JSON unmarshaling process for OrgSettingGatewayMgmtHostInPolicies objects.
func (o *OrgSettingGatewayMgmtHostInPolicies) UnmarshalJSON(input []byte) error {
	var temp tempOrgSettingGatewayMgmtHostInPolicies
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "icmp", "snmp")
	if err != nil {
		return err
	}
	o.AdditionalProperties = additionalProperties

	o.Icmp = temp.Icmp
	o.Snmp = temp.Snmp
	return nil
}

// tempOrgSettingGatewayMgmtHostInPolicies is a temporary struct used for validating the fields of OrgSettingGatewayMgmtHostInPolicies.
type tempOrgSettingGatewayMgmtHostInPolicies struct {
	Icmp *OrgSettingGatewayMgmtHostInPolicy `json:"icmp,omitempty"`
	Snmp *OrgSettingGatewayMgmtHostInPolicy `json:"snmp,omitempty"`
}
