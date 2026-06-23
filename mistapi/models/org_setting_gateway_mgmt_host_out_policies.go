// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// OrgSettingGatewayMgmtHostOutPolicies represents a OrgSettingGatewayMgmtHostOutPolicies struct.
// Optional path preferences for gateway-originated management traffic; ECMP is used across available paths when no preference is specified
type OrgSettingGatewayMgmtHostOutPolicies struct {
	// Host-out path policy for gateway-originated management traffic
	Dns *GatewayMgmtHostOutPolicy `json:"dns,omitempty"`
	// Host-out path policy for gateway-originated management traffic
	Ntp *GatewayMgmtHostOutPolicy `json:"ntp,omitempty"`
	// Host-out path policy for gateway syslog traffic
	Syslog               *GatewayMgmtHostOutPolicySyslog `json:"syslog,omitempty"`
	AdditionalProperties map[string]interface{}          `json:"_"`
}

// String implements the fmt.Stringer interface for OrgSettingGatewayMgmtHostOutPolicies,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (o OrgSettingGatewayMgmtHostOutPolicies) String() string {
	return fmt.Sprintf(
		"OrgSettingGatewayMgmtHostOutPolicies[Dns=%v, Ntp=%v, Syslog=%v, AdditionalProperties=%v]",
		o.Dns, o.Ntp, o.Syslog, o.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for OrgSettingGatewayMgmtHostOutPolicies.
// It customizes the JSON marshaling process for OrgSettingGatewayMgmtHostOutPolicies objects.
func (o OrgSettingGatewayMgmtHostOutPolicies) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(o.AdditionalProperties,
		"dns", "ntp", "syslog"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(o.toMap())
}

// toMap converts the OrgSettingGatewayMgmtHostOutPolicies object to a map representation for JSON marshaling.
func (o OrgSettingGatewayMgmtHostOutPolicies) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, o.AdditionalProperties)
	if o.Dns != nil {
		structMap["dns"] = o.Dns.toMap()
	}
	if o.Ntp != nil {
		structMap["ntp"] = o.Ntp.toMap()
	}
	if o.Syslog != nil {
		structMap["syslog"] = o.Syslog.toMap()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgSettingGatewayMgmtHostOutPolicies.
// It customizes the JSON unmarshaling process for OrgSettingGatewayMgmtHostOutPolicies objects.
func (o *OrgSettingGatewayMgmtHostOutPolicies) UnmarshalJSON(input []byte) error {
	var temp tempOrgSettingGatewayMgmtHostOutPolicies
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "dns", "ntp", "syslog")
	if err != nil {
		return err
	}
	o.AdditionalProperties = additionalProperties

	o.Dns = temp.Dns
	o.Ntp = temp.Ntp
	o.Syslog = temp.Syslog
	return nil
}

// tempOrgSettingGatewayMgmtHostOutPolicies is a temporary struct used for validating the fields of OrgSettingGatewayMgmtHostOutPolicies.
type tempOrgSettingGatewayMgmtHostOutPolicies struct {
	Dns    *GatewayMgmtHostOutPolicy       `json:"dns,omitempty"`
	Ntp    *GatewayMgmtHostOutPolicy       `json:"ntp,omitempty"`
	Syslog *GatewayMgmtHostOutPolicySyslog `json:"syslog,omitempty"`
}
