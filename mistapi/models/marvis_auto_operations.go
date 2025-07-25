// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// MarvisAutoOperations represents a MarvisAutoOperations struct.
type MarvisAutoOperations struct {
    BouncePortForAbnormalPoeClient         *bool                  `json:"bounce_port_for_abnormal_poe_client,omitempty"`
    DisablePortWhenDdosProtocolViolation   *bool                  `json:"disable_port_when_ddos_protocol_violation,omitempty"`
    DisablePortWhenRogueDhcpServerDetected *bool                  `json:"disable_port_when_rogue_dhcp_server_detected,omitempty"`
    AdditionalProperties                   map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for MarvisAutoOperations,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m MarvisAutoOperations) String() string {
    return fmt.Sprintf(
    	"MarvisAutoOperations[BouncePortForAbnormalPoeClient=%v, DisablePortWhenDdosProtocolViolation=%v, DisablePortWhenRogueDhcpServerDetected=%v, AdditionalProperties=%v]",
    	m.BouncePortForAbnormalPoeClient, m.DisablePortWhenDdosProtocolViolation, m.DisablePortWhenRogueDhcpServerDetected, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for MarvisAutoOperations.
// It customizes the JSON marshaling process for MarvisAutoOperations objects.
func (m MarvisAutoOperations) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(m.AdditionalProperties,
        "bounce_port_for_abnormal_poe_client", "disable_port_when_ddos_protocol_violation", "disable_port_when_rogue_dhcp_server_detected"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(m.toMap())
}

// toMap converts the MarvisAutoOperations object to a map representation for JSON marshaling.
func (m MarvisAutoOperations) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, m.AdditionalProperties)
    if m.BouncePortForAbnormalPoeClient != nil {
        structMap["bounce_port_for_abnormal_poe_client"] = m.BouncePortForAbnormalPoeClient
    }
    if m.DisablePortWhenDdosProtocolViolation != nil {
        structMap["disable_port_when_ddos_protocol_violation"] = m.DisablePortWhenDdosProtocolViolation
    }
    if m.DisablePortWhenRogueDhcpServerDetected != nil {
        structMap["disable_port_when_rogue_dhcp_server_detected"] = m.DisablePortWhenRogueDhcpServerDetected
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
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "bounce_port_for_abnormal_poe_client", "disable_port_when_ddos_protocol_violation", "disable_port_when_rogue_dhcp_server_detected")
    if err != nil {
    	return err
    }
    m.AdditionalProperties = additionalProperties
    
    m.BouncePortForAbnormalPoeClient = temp.BouncePortForAbnormalPoeClient
    m.DisablePortWhenDdosProtocolViolation = temp.DisablePortWhenDdosProtocolViolation
    m.DisablePortWhenRogueDhcpServerDetected = temp.DisablePortWhenRogueDhcpServerDetected
    return nil
}

// tempMarvisAutoOperations is a temporary struct used for validating the fields of MarvisAutoOperations.
type tempMarvisAutoOperations  struct {
    BouncePortForAbnormalPoeClient         *bool `json:"bounce_port_for_abnormal_poe_client,omitempty"`
    DisablePortWhenDdosProtocolViolation   *bool `json:"disable_port_when_ddos_protocol_violation,omitempty"`
    DisablePortWhenRogueDhcpServerDetected *bool `json:"disable_port_when_rogue_dhcp_server_detected,omitempty"`
}
