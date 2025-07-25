// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// SiteMxtunnelAdditionalMxtunnel represents a SiteMxtunnelAdditionalMxtunnel struct.
type SiteMxtunnelAdditionalMxtunnel struct {
    // For AP, how to connect to tunterm or RadSec Proxy
    Clusters             []SiteMxtunnelCluster     `json:"clusters,omitempty"`
    // In seconds, used as heartbeat to detect if a tunnel is alive. AP will try another peer after missing N hellos specified by hello_retries
    HelloInterval        *int                      `json:"hello_interval,omitempty"`
    HelloRetries         *int                      `json:"hello_retries,omitempty"`
    // enum: `ip`, `udp`
    Protocol             *SiteMxtunnelProtocolEnum `json:"protocol,omitempty"`
    VlanIds              []int                     `json:"vlan_ids,omitempty"`
    AdditionalProperties map[string]interface{}    `json:"_"`
}

// String implements the fmt.Stringer interface for SiteMxtunnelAdditionalMxtunnel,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SiteMxtunnelAdditionalMxtunnel) String() string {
    return fmt.Sprintf(
    	"SiteMxtunnelAdditionalMxtunnel[Clusters=%v, HelloInterval=%v, HelloRetries=%v, Protocol=%v, VlanIds=%v, AdditionalProperties=%v]",
    	s.Clusters, s.HelloInterval, s.HelloRetries, s.Protocol, s.VlanIds, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SiteMxtunnelAdditionalMxtunnel.
// It customizes the JSON marshaling process for SiteMxtunnelAdditionalMxtunnel objects.
func (s SiteMxtunnelAdditionalMxtunnel) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "clusters", "hello_interval", "hello_retries", "protocol", "vlan_ids"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SiteMxtunnelAdditionalMxtunnel object to a map representation for JSON marshaling.
func (s SiteMxtunnelAdditionalMxtunnel) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Clusters != nil {
        structMap["clusters"] = s.Clusters
    }
    if s.HelloInterval != nil {
        structMap["hello_interval"] = s.HelloInterval
    }
    if s.HelloRetries != nil {
        structMap["hello_retries"] = s.HelloRetries
    }
    if s.Protocol != nil {
        structMap["protocol"] = s.Protocol
    }
    if s.VlanIds != nil {
        structMap["vlan_ids"] = s.VlanIds
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SiteMxtunnelAdditionalMxtunnel.
// It customizes the JSON unmarshaling process for SiteMxtunnelAdditionalMxtunnel objects.
func (s *SiteMxtunnelAdditionalMxtunnel) UnmarshalJSON(input []byte) error {
    var temp tempSiteMxtunnelAdditionalMxtunnel
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "clusters", "hello_interval", "hello_retries", "protocol", "vlan_ids")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Clusters = temp.Clusters
    s.HelloInterval = temp.HelloInterval
    s.HelloRetries = temp.HelloRetries
    s.Protocol = temp.Protocol
    s.VlanIds = temp.VlanIds
    return nil
}

// tempSiteMxtunnelAdditionalMxtunnel is a temporary struct used for validating the fields of SiteMxtunnelAdditionalMxtunnel.
type tempSiteMxtunnelAdditionalMxtunnel  struct {
    Clusters      []SiteMxtunnelCluster     `json:"clusters,omitempty"`
    HelloInterval *int                      `json:"hello_interval,omitempty"`
    HelloRetries  *int                      `json:"hello_retries,omitempty"`
    Protocol      *SiteMxtunnelProtocolEnum `json:"protocol,omitempty"`
    VlanIds       []int                     `json:"vlan_ids,omitempty"`
}
