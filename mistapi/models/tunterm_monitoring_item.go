// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// TuntermMonitoringItem represents a TuntermMonitoringItem struct.
type TuntermMonitoringItem struct {
	// Can be ip, ipv6, hostname
	Host *string `json:"host,omitempty"`
	// When `protocol`==`tcp`
	Port *int `json:"port,omitempty"`
	// enum: `arp`, `ping`, `tcp`
	Protocol *TuntermMonitoringProtocolEnum `json:"protocol,omitempty"`
	// Optional source for the monitoring check, vlan_id configured in tunterm_other_ip_configs
	SrcVlanId            *int                   `json:"src_vlan_id,omitempty"`
	Timeout              *int                   `json:"timeout,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for TuntermMonitoringItem,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (t TuntermMonitoringItem) String() string {
	return fmt.Sprintf(
		"TuntermMonitoringItem[Host=%v, Port=%v, Protocol=%v, SrcVlanId=%v, Timeout=%v, AdditionalProperties=%v]",
		t.Host, t.Port, t.Protocol, t.SrcVlanId, t.Timeout, t.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for TuntermMonitoringItem.
// It customizes the JSON marshaling process for TuntermMonitoringItem objects.
func (t TuntermMonitoringItem) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(t.AdditionalProperties,
		"host", "port", "protocol", "src_vlan_id", "timeout"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(t.toMap())
}

// toMap converts the TuntermMonitoringItem object to a map representation for JSON marshaling.
func (t TuntermMonitoringItem) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, t.AdditionalProperties)
	if t.Host != nil {
		structMap["host"] = t.Host
	}
	if t.Port != nil {
		structMap["port"] = t.Port
	}
	if t.Protocol != nil {
		structMap["protocol"] = t.Protocol
	}
	if t.SrcVlanId != nil {
		structMap["src_vlan_id"] = t.SrcVlanId
	}
	if t.Timeout != nil {
		structMap["timeout"] = t.Timeout
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for TuntermMonitoringItem.
// It customizes the JSON unmarshaling process for TuntermMonitoringItem objects.
func (t *TuntermMonitoringItem) UnmarshalJSON(input []byte) error {
	var temp tempTuntermMonitoringItem
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "host", "port", "protocol", "src_vlan_id", "timeout")
	if err != nil {
		return err
	}
	t.AdditionalProperties = additionalProperties

	t.Host = temp.Host
	t.Port = temp.Port
	t.Protocol = temp.Protocol
	t.SrcVlanId = temp.SrcVlanId
	t.Timeout = temp.Timeout
	return nil
}

// tempTuntermMonitoringItem is a temporary struct used for validating the fields of TuntermMonitoringItem.
type tempTuntermMonitoringItem struct {
	Host      *string                        `json:"host,omitempty"`
	Port      *int                           `json:"port,omitempty"`
	Protocol  *TuntermMonitoringProtocolEnum `json:"protocol,omitempty"`
	SrcVlanId *int                           `json:"src_vlan_id,omitempty"`
	Timeout   *int                           `json:"timeout,omitempty"`
}
