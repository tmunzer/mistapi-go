package models

import (
    "encoding/json"
    "fmt"
    "github.com/google/uuid"
)

// Mxtunnel represents a Mxtunnel struct.
// MxTunnel
type Mxtunnel struct {
    // list of anchor mxtunnels used for forming edge to edge tunnels
    AnchorMxtunnelIds    []uuid.UUID            `json:"anchor_mxtunnel_ids,omitempty"`
    // schedule to preempt ap’s which are not connected to preferred peer
    AutoPreemption       *AutoPreemption        `json:"auto_preemption,omitempty"`
    // when the object has been created, in epoch
    CreatedTime          *float64               `json:"created_time,omitempty"`
    ForSite              *bool                  `json:"for_site,omitempty"`
    // in seconds, used as heartbeat to detect if a tunnel is alive. AP will try another peer after missing N hellos specified by `hello_retries`.
    HelloInterval        Optional[int]          `json:"hello_interval"`
    HelloRetries         Optional[int]          `json:"hello_retries"`
    // Unique ID of the object instance in the Mist Organnization
    Id                   *uuid.UUID             `json:"id,omitempty"`
    Ipsec                *MxtunnelIpsec         `json:"ipsec,omitempty"`
    // when the object has been modified for the last time, in epoch
    ModifiedTime         *float64               `json:"modified_time,omitempty"`
    // 0 to enable PMTU, 552-1500 to start PMTU with a lower MTU
    Mtu                  *int                   `json:"mtu,omitempty"`
    // list of mxclusters to deploy this tunnel to
    MxclusterIds         []uuid.UUID            `json:"mxcluster_ids,omitempty"`
    Name                 Optional[string]       `json:"name"`
    OrgId                *uuid.UUID             `json:"org_id,omitempty"`
    // enum: `ip`, `udp`
    Protocol             *MxtunnelProtocolEnum  `json:"protocol,omitempty"`
    SiteId               *uuid.UUID             `json:"site_id,omitempty"`
    // list of vlan_ids that will be used
    VlanIds              []int                  `json:"vlan_ids,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for Mxtunnel,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m Mxtunnel) String() string {
    return fmt.Sprintf(
    	"Mxtunnel[AnchorMxtunnelIds=%v, AutoPreemption=%v, CreatedTime=%v, ForSite=%v, HelloInterval=%v, HelloRetries=%v, Id=%v, Ipsec=%v, ModifiedTime=%v, Mtu=%v, MxclusterIds=%v, Name=%v, OrgId=%v, Protocol=%v, SiteId=%v, VlanIds=%v, AdditionalProperties=%v]",
    	m.AnchorMxtunnelIds, m.AutoPreemption, m.CreatedTime, m.ForSite, m.HelloInterval, m.HelloRetries, m.Id, m.Ipsec, m.ModifiedTime, m.Mtu, m.MxclusterIds, m.Name, m.OrgId, m.Protocol, m.SiteId, m.VlanIds, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for Mxtunnel.
// It customizes the JSON marshaling process for Mxtunnel objects.
func (m Mxtunnel) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(m.AdditionalProperties,
        "anchor_mxtunnel_ids", "auto_preemption", "created_time", "for_site", "hello_interval", "hello_retries", "id", "ipsec", "modified_time", "mtu", "mxcluster_ids", "name", "org_id", "protocol", "site_id", "vlan_ids"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(m.toMap())
}

// toMap converts the Mxtunnel object to a map representation for JSON marshaling.
func (m Mxtunnel) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, m.AdditionalProperties)
    if m.AnchorMxtunnelIds != nil {
        structMap["anchor_mxtunnel_ids"] = m.AnchorMxtunnelIds
    }
    if m.AutoPreemption != nil {
        structMap["auto_preemption"] = m.AutoPreemption.toMap()
    }
    if m.CreatedTime != nil {
        structMap["created_time"] = m.CreatedTime
    }
    if m.ForSite != nil {
        structMap["for_site"] = m.ForSite
    }
    if m.HelloInterval.IsValueSet() {
        if m.HelloInterval.Value() != nil {
            structMap["hello_interval"] = m.HelloInterval.Value()
        } else {
            structMap["hello_interval"] = nil
        }
    }
    if m.HelloRetries.IsValueSet() {
        if m.HelloRetries.Value() != nil {
            structMap["hello_retries"] = m.HelloRetries.Value()
        } else {
            structMap["hello_retries"] = nil
        }
    }
    if m.Id != nil {
        structMap["id"] = m.Id
    }
    if m.Ipsec != nil {
        structMap["ipsec"] = m.Ipsec.toMap()
    }
    if m.ModifiedTime != nil {
        structMap["modified_time"] = m.ModifiedTime
    }
    if m.Mtu != nil {
        structMap["mtu"] = m.Mtu
    }
    if m.MxclusterIds != nil {
        structMap["mxcluster_ids"] = m.MxclusterIds
    }
    if m.Name.IsValueSet() {
        if m.Name.Value() != nil {
            structMap["name"] = m.Name.Value()
        } else {
            structMap["name"] = nil
        }
    }
    if m.OrgId != nil {
        structMap["org_id"] = m.OrgId
    }
    if m.Protocol != nil {
        structMap["protocol"] = m.Protocol
    }
    if m.SiteId != nil {
        structMap["site_id"] = m.SiteId
    }
    if m.VlanIds != nil {
        structMap["vlan_ids"] = m.VlanIds
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Mxtunnel.
// It customizes the JSON unmarshaling process for Mxtunnel objects.
func (m *Mxtunnel) UnmarshalJSON(input []byte) error {
    var temp tempMxtunnel
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "anchor_mxtunnel_ids", "auto_preemption", "created_time", "for_site", "hello_interval", "hello_retries", "id", "ipsec", "modified_time", "mtu", "mxcluster_ids", "name", "org_id", "protocol", "site_id", "vlan_ids")
    if err != nil {
    	return err
    }
    m.AdditionalProperties = additionalProperties
    
    m.AnchorMxtunnelIds = temp.AnchorMxtunnelIds
    m.AutoPreemption = temp.AutoPreemption
    m.CreatedTime = temp.CreatedTime
    m.ForSite = temp.ForSite
    m.HelloInterval = temp.HelloInterval
    m.HelloRetries = temp.HelloRetries
    m.Id = temp.Id
    m.Ipsec = temp.Ipsec
    m.ModifiedTime = temp.ModifiedTime
    m.Mtu = temp.Mtu
    m.MxclusterIds = temp.MxclusterIds
    m.Name = temp.Name
    m.OrgId = temp.OrgId
    m.Protocol = temp.Protocol
    m.SiteId = temp.SiteId
    m.VlanIds = temp.VlanIds
    return nil
}

// tempMxtunnel is a temporary struct used for validating the fields of Mxtunnel.
type tempMxtunnel  struct {
    AnchorMxtunnelIds []uuid.UUID           `json:"anchor_mxtunnel_ids,omitempty"`
    AutoPreemption    *AutoPreemption       `json:"auto_preemption,omitempty"`
    CreatedTime       *float64              `json:"created_time,omitempty"`
    ForSite           *bool                 `json:"for_site,omitempty"`
    HelloInterval     Optional[int]         `json:"hello_interval"`
    HelloRetries      Optional[int]         `json:"hello_retries"`
    Id                *uuid.UUID            `json:"id,omitempty"`
    Ipsec             *MxtunnelIpsec        `json:"ipsec,omitempty"`
    ModifiedTime      *float64              `json:"modified_time,omitempty"`
    Mtu               *int                  `json:"mtu,omitempty"`
    MxclusterIds      []uuid.UUID           `json:"mxcluster_ids,omitempty"`
    Name              Optional[string]      `json:"name"`
    OrgId             *uuid.UUID            `json:"org_id,omitempty"`
    Protocol          *MxtunnelProtocolEnum `json:"protocol,omitempty"`
    SiteId            *uuid.UUID            `json:"site_id,omitempty"`
    VlanIds           []int                 `json:"vlan_ids,omitempty"`
}
