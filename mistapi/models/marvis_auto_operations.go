// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// MarvisAutoOperations represents a MarvisAutoOperations struct.
type MarvisAutoOperations struct {
	ApInsufficientCapacity                 *bool                  `json:"ap_insufficient_capacity,omitempty"`
	ApLoop                                 *bool                  `json:"ap_loop,omitempty"`
	ApNonCompliant                         *bool                  `json:"ap_non_compliant,omitempty"`
	BouncePortForAbnormalPoeClient         *bool                  `json:"bounce_port_for_abnormal_poe_client,omitempty"`
	DisablePortWhenDdosProtocolViolation   *bool                  `json:"disable_port_when_ddos_protocol_violation,omitempty"`
	DisablePortWhenRogueDhcpServerDetected *bool                  `json:"disable_port_when_rogue_dhcp_server_detected,omitempty"`
	GatewayNonCompliant                    *bool                  `json:"gateway_non_compliant,omitempty"`
	SwitchMisconfiguredPort                *bool                  `json:"switch_misconfigured_port,omitempty"`
	SwitchPortStuck                        *bool                  `json:"switch_port_stuck,omitempty"`
	AdditionalProperties                   map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for MarvisAutoOperations,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m MarvisAutoOperations) String() string {
	return fmt.Sprintf(
		"MarvisAutoOperations[ApInsufficientCapacity=%v, ApLoop=%v, ApNonCompliant=%v, BouncePortForAbnormalPoeClient=%v, DisablePortWhenDdosProtocolViolation=%v, DisablePortWhenRogueDhcpServerDetected=%v, GatewayNonCompliant=%v, SwitchMisconfiguredPort=%v, SwitchPortStuck=%v, AdditionalProperties=%v]",
		m.ApInsufficientCapacity, m.ApLoop, m.ApNonCompliant, m.BouncePortForAbnormalPoeClient, m.DisablePortWhenDdosProtocolViolation, m.DisablePortWhenRogueDhcpServerDetected, m.GatewayNonCompliant, m.SwitchMisconfiguredPort, m.SwitchPortStuck, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for MarvisAutoOperations.
// It customizes the JSON marshaling process for MarvisAutoOperations objects.
func (m MarvisAutoOperations) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(m.AdditionalProperties,
		"ap_insufficient_capacity", "ap_loop", "ap_non_compliant", "bounce_port_for_abnormal_poe_client", "disable_port_when_ddos_protocol_violation", "disable_port_when_rogue_dhcp_server_detected", "gateway_non_compliant", "switch_misconfigured_port", "switch_port_stuck"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(m.toMap())
}

// toMap converts the MarvisAutoOperations object to a map representation for JSON marshaling.
func (m MarvisAutoOperations) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, m.AdditionalProperties)
	if m.ApInsufficientCapacity != nil {
		structMap["ap_insufficient_capacity"] = m.ApInsufficientCapacity
	}
	if m.ApLoop != nil {
		structMap["ap_loop"] = m.ApLoop
	}
	if m.ApNonCompliant != nil {
		structMap["ap_non_compliant"] = m.ApNonCompliant
	}
	if m.BouncePortForAbnormalPoeClient != nil {
		structMap["bounce_port_for_abnormal_poe_client"] = m.BouncePortForAbnormalPoeClient
	}
	if m.DisablePortWhenDdosProtocolViolation != nil {
		structMap["disable_port_when_ddos_protocol_violation"] = m.DisablePortWhenDdosProtocolViolation
	}
	if m.DisablePortWhenRogueDhcpServerDetected != nil {
		structMap["disable_port_when_rogue_dhcp_server_detected"] = m.DisablePortWhenRogueDhcpServerDetected
	}
	if m.GatewayNonCompliant != nil {
		structMap["gateway_non_compliant"] = m.GatewayNonCompliant
	}
	if m.SwitchMisconfiguredPort != nil {
		structMap["switch_misconfigured_port"] = m.SwitchMisconfiguredPort
	}
	if m.SwitchPortStuck != nil {
		structMap["switch_port_stuck"] = m.SwitchPortStuck
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MarvisAutoOperations.
// It customizes the JSON unmarshaling process for MarvisAutoOperations objects.
func (m *MarvisAutoOperations) UnmarshalJSON(input []byte) error {
	var temp tempMarvisAutoOperations
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "ap_insufficient_capacity", "ap_loop", "ap_non_compliant", "bounce_port_for_abnormal_poe_client", "disable_port_when_ddos_protocol_violation", "disable_port_when_rogue_dhcp_server_detected", "gateway_non_compliant", "switch_misconfigured_port", "switch_port_stuck")
	if err != nil {
		return err
	}
	m.AdditionalProperties = additionalProperties

	m.ApInsufficientCapacity = temp.ApInsufficientCapacity
	m.ApLoop = temp.ApLoop
	m.ApNonCompliant = temp.ApNonCompliant
	m.BouncePortForAbnormalPoeClient = temp.BouncePortForAbnormalPoeClient
	m.DisablePortWhenDdosProtocolViolation = temp.DisablePortWhenDdosProtocolViolation
	m.DisablePortWhenRogueDhcpServerDetected = temp.DisablePortWhenRogueDhcpServerDetected
	m.GatewayNonCompliant = temp.GatewayNonCompliant
	m.SwitchMisconfiguredPort = temp.SwitchMisconfiguredPort
	m.SwitchPortStuck = temp.SwitchPortStuck
	return nil
}

// tempMarvisAutoOperations is a temporary struct used for validating the fields of MarvisAutoOperations.
type tempMarvisAutoOperations struct {
	ApInsufficientCapacity                 *bool `json:"ap_insufficient_capacity,omitempty"`
	ApLoop                                 *bool `json:"ap_loop,omitempty"`
	ApNonCompliant                         *bool `json:"ap_non_compliant,omitempty"`
	BouncePortForAbnormalPoeClient         *bool `json:"bounce_port_for_abnormal_poe_client,omitempty"`
	DisablePortWhenDdosProtocolViolation   *bool `json:"disable_port_when_ddos_protocol_violation,omitempty"`
	DisablePortWhenRogueDhcpServerDetected *bool `json:"disable_port_when_rogue_dhcp_server_detected,omitempty"`
	GatewayNonCompliant                    *bool `json:"gateway_non_compliant,omitempty"`
	SwitchMisconfiguredPort                *bool `json:"switch_misconfigured_port,omitempty"`
	SwitchPortStuck                        *bool `json:"switch_port_stuck,omitempty"`
}
