package models

import (
    "encoding/json"
    "fmt"
    "github.com/google/uuid"
)

// SiteMxtunnel represents a SiteMxtunnel struct.
// Site MxTunnel
type SiteMxtunnel struct {
    AdditionalMxtunnels  map[string]SiteMxtunnelAdditionalMxtunnel `json:"additional_mxtunnels,omitempty"`
    // List of subnets where we allow AP to establish Mist Tunnels from
    ApSubnets            []string                                  `json:"ap_subnets,omitempty"`
    // Schedule to preempt ap’s which are not connected to preferred peer
    AutoPreemption       *AutoPreemption                           `json:"auto_preemption,omitempty"`
    // For AP, how to connect to tunterm or radsecproxy
    Clusters             []SiteMxtunnelCluster                     `json:"clusters,omitempty"`
    // When the object has been created, in epoch
    CreatedTime          *float64                                  `json:"created_time,omitempty"`
    Enabled              *bool                                     `json:"enabled,omitempty"`
    ForSite              *bool                                     `json:"for_site,omitempty"`
    // In seconds, used as heartbeat to detect if a tunnel is alive. AP will try another peer after missing N hellos specified by hello_retries
    HelloInterval        *int                                      `json:"hello_interval,omitempty"`
    HelloRetries         *int                                      `json:"hello_retries,omitempty"`
    // Hostnames or IPs where a Mist Tunnel will use as the Peer (i.e. they are reachable from AP)
    Hosts                []string                                  `json:"hosts,omitempty"`
    // Unique ID of the object instance in the Mist Organnization
    Id                   *uuid.UUID                                `json:"id,omitempty"`
    // When the object has been modified for the last time, in epoch
    ModifiedTime         *float64                                  `json:"modified_time,omitempty"`
    // 0 to enable PMTU, 552-1500 to start PMTU with a lower MTU
    Mtu                  *int                                      `json:"mtu,omitempty"`
    OrgId                *uuid.UUID                                `json:"org_id,omitempty"`
    // enum: `ip`, `udp`
    Protocol             *MxtunnelProtocolEnum                     `json:"protocol,omitempty"`
    Radsec               *SiteMxtunnelRadsec                       `json:"radsec,omitempty"`
    SiteId               *uuid.UUID                                `json:"site_id,omitempty"`
    // List of vlan_ids that will be used
    VlanIds              []int                                     `json:"vlan_ids,omitempty"`
    AdditionalProperties map[string]interface{}                    `json:"_"`
}

// String implements the fmt.Stringer interface for SiteMxtunnel,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SiteMxtunnel) String() string {
    return fmt.Sprintf(
    	"SiteMxtunnel[AdditionalMxtunnels=%v, ApSubnets=%v, AutoPreemption=%v, Clusters=%v, CreatedTime=%v, Enabled=%v, ForSite=%v, HelloInterval=%v, HelloRetries=%v, Hosts=%v, Id=%v, ModifiedTime=%v, Mtu=%v, OrgId=%v, Protocol=%v, Radsec=%v, SiteId=%v, VlanIds=%v, AdditionalProperties=%v]",
    	s.AdditionalMxtunnels, s.ApSubnets, s.AutoPreemption, s.Clusters, s.CreatedTime, s.Enabled, s.ForSite, s.HelloInterval, s.HelloRetries, s.Hosts, s.Id, s.ModifiedTime, s.Mtu, s.OrgId, s.Protocol, s.Radsec, s.SiteId, s.VlanIds, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SiteMxtunnel.
// It customizes the JSON marshaling process for SiteMxtunnel objects.
func (s SiteMxtunnel) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "additional_mxtunnels", "ap_subnets", "auto_preemption", "clusters", "created_time", "enabled", "for_site", "hello_interval", "hello_retries", "hosts", "id", "modified_time", "mtu", "org_id", "protocol", "radsec", "site_id", "vlan_ids"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SiteMxtunnel object to a map representation for JSON marshaling.
func (s SiteMxtunnel) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.AdditionalMxtunnels != nil {
        structMap["additional_mxtunnels"] = s.AdditionalMxtunnels
    }
    if s.ApSubnets != nil {
        structMap["ap_subnets"] = s.ApSubnets
    }
    if s.AutoPreemption != nil {
        structMap["auto_preemption"] = s.AutoPreemption.toMap()
    }
    if s.Clusters != nil {
        structMap["clusters"] = s.Clusters
    }
    if s.CreatedTime != nil {
        structMap["created_time"] = s.CreatedTime
    }
    if s.Enabled != nil {
        structMap["enabled"] = s.Enabled
    }
    if s.ForSite != nil {
        structMap["for_site"] = s.ForSite
    }
    if s.HelloInterval != nil {
        structMap["hello_interval"] = s.HelloInterval
    }
    if s.HelloRetries != nil {
        structMap["hello_retries"] = s.HelloRetries
    }
    if s.Hosts != nil {
        structMap["hosts"] = s.Hosts
    }
    if s.Id != nil {
        structMap["id"] = s.Id
    }
    if s.ModifiedTime != nil {
        structMap["modified_time"] = s.ModifiedTime
    }
    if s.Mtu != nil {
        structMap["mtu"] = s.Mtu
    }
    if s.OrgId != nil {
        structMap["org_id"] = s.OrgId
    }
    if s.Protocol != nil {
        structMap["protocol"] = s.Protocol
    }
    if s.Radsec != nil {
        structMap["radsec"] = s.Radsec.toMap()
    }
    if s.SiteId != nil {
        structMap["site_id"] = s.SiteId
    }
    if s.VlanIds != nil {
        structMap["vlan_ids"] = s.VlanIds
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SiteMxtunnel.
// It customizes the JSON unmarshaling process for SiteMxtunnel objects.
func (s *SiteMxtunnel) UnmarshalJSON(input []byte) error {
    var temp tempSiteMxtunnel
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "additional_mxtunnels", "ap_subnets", "auto_preemption", "clusters", "created_time", "enabled", "for_site", "hello_interval", "hello_retries", "hosts", "id", "modified_time", "mtu", "org_id", "protocol", "radsec", "site_id", "vlan_ids")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.AdditionalMxtunnels = temp.AdditionalMxtunnels
    s.ApSubnets = temp.ApSubnets
    s.AutoPreemption = temp.AutoPreemption
    s.Clusters = temp.Clusters
    s.CreatedTime = temp.CreatedTime
    s.Enabled = temp.Enabled
    s.ForSite = temp.ForSite
    s.HelloInterval = temp.HelloInterval
    s.HelloRetries = temp.HelloRetries
    s.Hosts = temp.Hosts
    s.Id = temp.Id
    s.ModifiedTime = temp.ModifiedTime
    s.Mtu = temp.Mtu
    s.OrgId = temp.OrgId
    s.Protocol = temp.Protocol
    s.Radsec = temp.Radsec
    s.SiteId = temp.SiteId
    s.VlanIds = temp.VlanIds
    return nil
}

// tempSiteMxtunnel is a temporary struct used for validating the fields of SiteMxtunnel.
type tempSiteMxtunnel  struct {
    AdditionalMxtunnels map[string]SiteMxtunnelAdditionalMxtunnel `json:"additional_mxtunnels,omitempty"`
    ApSubnets           []string                                  `json:"ap_subnets,omitempty"`
    AutoPreemption      *AutoPreemption                           `json:"auto_preemption,omitempty"`
    Clusters            []SiteMxtunnelCluster                     `json:"clusters,omitempty"`
    CreatedTime         *float64                                  `json:"created_time,omitempty"`
    Enabled             *bool                                     `json:"enabled,omitempty"`
    ForSite             *bool                                     `json:"for_site,omitempty"`
    HelloInterval       *int                                      `json:"hello_interval,omitempty"`
    HelloRetries        *int                                      `json:"hello_retries,omitempty"`
    Hosts               []string                                  `json:"hosts,omitempty"`
    Id                  *uuid.UUID                                `json:"id,omitempty"`
    ModifiedTime        *float64                                  `json:"modified_time,omitempty"`
    Mtu                 *int                                      `json:"mtu,omitempty"`
    OrgId               *uuid.UUID                                `json:"org_id,omitempty"`
    Protocol            *MxtunnelProtocolEnum                     `json:"protocol,omitempty"`
    Radsec              *SiteMxtunnelRadsec                       `json:"radsec,omitempty"`
    SiteId              *uuid.UUID                                `json:"site_id,omitempty"`
    VlanIds             []int                                     `json:"vlan_ids,omitempty"`
}
