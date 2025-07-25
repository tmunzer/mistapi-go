// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// SimpleAlert represents a SimpleAlert struct.
// Set of heuristic rules will be enabled when marvis subscription is not available. It triggers when, in a Z minute window, there are more than Y distinct client encountering over X failures
type SimpleAlert struct {
    ArpFailure           *SimpleAlertArpFailure  `json:"arp_failure,omitempty"`
    DhcpFailure          *SimpleAlertDhcpFailure `json:"dhcp_failure,omitempty"`
    DnsFailure           *SimpleAlertDnsFailure  `json:"dns_failure,omitempty"`
    AdditionalProperties map[string]interface{}  `json:"_"`
}

// String implements the fmt.Stringer interface for SimpleAlert,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SimpleAlert) String() string {
    return fmt.Sprintf(
    	"SimpleAlert[ArpFailure=%v, DhcpFailure=%v, DnsFailure=%v, AdditionalProperties=%v]",
    	s.ArpFailure, s.DhcpFailure, s.DnsFailure, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SimpleAlert.
// It customizes the JSON marshaling process for SimpleAlert objects.
func (s SimpleAlert) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "arp_failure", "dhcp_failure", "dns_failure"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SimpleAlert object to a map representation for JSON marshaling.
func (s SimpleAlert) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.ArpFailure != nil {
        structMap["arp_failure"] = s.ArpFailure.toMap()
    }
    if s.DhcpFailure != nil {
        structMap["dhcp_failure"] = s.DhcpFailure.toMap()
    }
    if s.DnsFailure != nil {
        structMap["dns_failure"] = s.DnsFailure.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SimpleAlert.
// It customizes the JSON unmarshaling process for SimpleAlert objects.
func (s *SimpleAlert) UnmarshalJSON(input []byte) error {
    var temp tempSimpleAlert
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "arp_failure", "dhcp_failure", "dns_failure")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.ArpFailure = temp.ArpFailure
    s.DhcpFailure = temp.DhcpFailure
    s.DnsFailure = temp.DnsFailure
    return nil
}

// tempSimpleAlert is a temporary struct used for validating the fields of SimpleAlert.
type tempSimpleAlert  struct {
    ArpFailure  *SimpleAlertArpFailure  `json:"arp_failure,omitempty"`
    DhcpFailure *SimpleAlertDhcpFailure `json:"dhcp_failure,omitempty"`
    DnsFailure  *SimpleAlertDnsFailure  `json:"dns_failure,omitempty"`
}
