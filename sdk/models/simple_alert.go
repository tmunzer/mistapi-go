package models

import (
    "encoding/json"
)

// SimpleAlert represents a SimpleAlert struct.
// Set of heuristic rules will be enabled when marvis subscription is not available.
// It triggers when, in a Z minute window, there are more than Y distinct client encountring over X failures
type SimpleAlert struct {
    ArpFailure           *SimpleAlertArpFailure  `json:"arp_failure,omitempty"`
    DhcpFailure          *SimpleAlertDhcpFailure `json:"dhcp_failure,omitempty"`
    DnsFailure           *SimpleAlertDnsFailure  `json:"dns_failure,omitempty"`
    AdditionalProperties map[string]any          `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SimpleAlert.
// It customizes the JSON marshaling process for SimpleAlert objects.
func (s SimpleAlert) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SimpleAlert object to a map representation for JSON marshaling.
func (s SimpleAlert) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
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
    var temp simpleAlert
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "arp_failure", "dhcp_failure", "dns_failure")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.ArpFailure = temp.ArpFailure
    s.DhcpFailure = temp.DhcpFailure
    s.DnsFailure = temp.DnsFailure
    return nil
}

// simpleAlert is a temporary struct used for validating the fields of SimpleAlert.
type simpleAlert  struct {
    ArpFailure  *SimpleAlertArpFailure  `json:"arp_failure,omitempty"`
    DhcpFailure *SimpleAlertDhcpFailure `json:"dhcp_failure,omitempty"`
    DnsFailure  *SimpleAlertDnsFailure  `json:"dns_failure,omitempty"`
}
