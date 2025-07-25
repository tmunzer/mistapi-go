// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// SiteMxtunnelCluster represents a SiteMxtunnelCluster struct.
type SiteMxtunnelCluster struct {
    Name                 *string                `json:"name,omitempty"`
    TuntermHosts         []string               `json:"tunterm_hosts,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SiteMxtunnelCluster,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SiteMxtunnelCluster) String() string {
    return fmt.Sprintf(
    	"SiteMxtunnelCluster[Name=%v, TuntermHosts=%v, AdditionalProperties=%v]",
    	s.Name, s.TuntermHosts, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SiteMxtunnelCluster.
// It customizes the JSON marshaling process for SiteMxtunnelCluster objects.
func (s SiteMxtunnelCluster) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "name", "tunterm_hosts"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SiteMxtunnelCluster object to a map representation for JSON marshaling.
func (s SiteMxtunnelCluster) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Name != nil {
        structMap["name"] = s.Name
    }
    if s.TuntermHosts != nil {
        structMap["tunterm_hosts"] = s.TuntermHosts
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SiteMxtunnelCluster.
// It customizes the JSON unmarshaling process for SiteMxtunnelCluster objects.
func (s *SiteMxtunnelCluster) UnmarshalJSON(input []byte) error {
    var temp tempSiteMxtunnelCluster
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "name", "tunterm_hosts")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Name = temp.Name
    s.TuntermHosts = temp.TuntermHosts
    return nil
}

// tempSiteMxtunnelCluster is a temporary struct used for validating the fields of SiteMxtunnelCluster.
type tempSiteMxtunnelCluster  struct {
    Name         *string  `json:"name,omitempty"`
    TuntermHosts []string `json:"tunterm_hosts,omitempty"`
}
