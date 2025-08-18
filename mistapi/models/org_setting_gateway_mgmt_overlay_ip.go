// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// OrgSettingGatewayMgmtOverlayIp represents a OrgSettingGatewayMgmtOverlayIp struct.
type OrgSettingGatewayMgmtOverlayIp struct {
	// When it's going overlay, a routable IP to overlay will be required
	Ip *string `json:"ip,omitempty"`
	// For SSR HA cluster, another IP for node1 will be required, too
	Node1Ip              *string                `json:"node1_ip,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for OrgSettingGatewayMgmtOverlayIp,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (o OrgSettingGatewayMgmtOverlayIp) String() string {
	return fmt.Sprintf(
		"OrgSettingGatewayMgmtOverlayIp[Ip=%v, Node1Ip=%v, AdditionalProperties=%v]",
		o.Ip, o.Node1Ip, o.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for OrgSettingGatewayMgmtOverlayIp.
// It customizes the JSON marshaling process for OrgSettingGatewayMgmtOverlayIp objects.
func (o OrgSettingGatewayMgmtOverlayIp) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(o.AdditionalProperties,
		"ip", "node1_ip"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(o.toMap())
}

// toMap converts the OrgSettingGatewayMgmtOverlayIp object to a map representation for JSON marshaling.
func (o OrgSettingGatewayMgmtOverlayIp) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, o.AdditionalProperties)
	if o.Ip != nil {
		structMap["ip"] = o.Ip
	}
	if o.Node1Ip != nil {
		structMap["node1_ip"] = o.Node1Ip
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgSettingGatewayMgmtOverlayIp.
// It customizes the JSON unmarshaling process for OrgSettingGatewayMgmtOverlayIp objects.
func (o *OrgSettingGatewayMgmtOverlayIp) UnmarshalJSON(input []byte) error {
	var temp tempOrgSettingGatewayMgmtOverlayIp
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "ip", "node1_ip")
	if err != nil {
		return err
	}
	o.AdditionalProperties = additionalProperties

	o.Ip = temp.Ip
	o.Node1Ip = temp.Node1Ip
	return nil
}

// tempOrgSettingGatewayMgmtOverlayIp is a temporary struct used for validating the fields of OrgSettingGatewayMgmtOverlayIp.
type tempOrgSettingGatewayMgmtOverlayIp struct {
	Ip      *string `json:"ip,omitempty"`
	Node1Ip *string `json:"node1_ip,omitempty"`
}
