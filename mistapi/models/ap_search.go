package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/google/uuid"
    "strings"
)

// ApSearch represents a ApSearch struct.
type ApSearch struct {
    // bandwidth of band_24
    Band24Bandwidth      *string                `json:"band_24_bandwidth,omitempty"`
    // Channel of band_24
    Band24Channel        *int                   `json:"band_24_channel,omitempty"`
    Band24Power          *int                   `json:"band_24_power,omitempty"`
    // bandwidth of band_5
    Band5Bandwidth       *string                `json:"band_5_bandwidth,omitempty"`
    // Channel of band_5
    Band5Channel         *int                   `json:"band_5_channel,omitempty"`
    Band5Power           *int                   `json:"band_5_power,omitempty"`
    Band6Bandwidth       *string                `json:"band_6_bandwidth,omitempty"`
    // Channel of band_6
    Band6Channel         *int                   `json:"band_6_channel,omitempty"`
    Band6Power           *int                   `json:"band_6_power,omitempty"`
    // Port speed of eth0
    Eth0PortSpeed        *int                   `json:"eth0_port_speed,omitempty"`
    ExtIp                *string                `json:"ext_ip,omitempty"`
    // partial / full hostname
    Hostname             []string               `json:"hostname,omitempty"`
    InactiveWiredVlans   []int                  `json:"inactive_wired_vlans,omitempty"`
    // ip address
    Ip                   *string                `json:"ip,omitempty"`
    LastHostname         *string                `json:"last_hostname,omitempty"`
    // LLDP management ip address
    LldpMgmtAddr         *string                `json:"lldp_mgmt_addr,omitempty"`
    LldpPortDesc         *string                `json:"lldp_port_desc,omitempty"`
    // LLDP port id
    LldpPortId           *string                `json:"lldp_port_id,omitempty"`
    LldpPowerAllocated   *int                   `json:"lldp_power_allocated,omitempty"`
    LldpPowerDraw        *int                   `json:"lldp_power_draw,omitempty"`
    // LLDP system description
    LldpSystemDesc       *string                `json:"lldp_system_desc,omitempty"`
    // LLDP system name
    LldpSystemName       *string                `json:"lldp_system_name,omitempty"`
    // device model
    Mac                  *string                `json:"mac,omitempty"`
    Model                *string                `json:"model,omitempty"`
    // Mist Edge id, if AP is connecting to a Mist Edge
    MxedgeId             *string                `json:"mxedge_id,omitempty"`
    // Comma separated list of Mist Edge ids, if AP is connecting to a Mist Edge
    MxedgeIds            *string                `json:"mxedge_ids,omitempty"`
    // MxTunnel status
    MxtunnelStatus       string                 `json:"mxtunnel_status"`
    OrgId                *uuid.UUID             `json:"org_id,omitempty"`
    PowerConstrained     bool                   `json:"power_constrained"`
    PowerOpmode          string                 `json:"power_opmode"`
    SiteId               *uuid.UUID             `json:"site_id,omitempty"`
    Sku                  *string                `json:"sku,omitempty"`
    Timestamp            *float64               `json:"timestamp,omitempty"`
    Uptime               *int                   `json:"uptime,omitempty"`
    // version
    Version              *string                `json:"version,omitempty"`
    Wlans                []ApSearchWlan         `json:"wlans"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ApSearch,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a ApSearch) String() string {
    return fmt.Sprintf(
    	"ApSearch[Band24Bandwidth=%v, Band24Channel=%v, Band24Power=%v, Band5Bandwidth=%v, Band5Channel=%v, Band5Power=%v, Band6Bandwidth=%v, Band6Channel=%v, Band6Power=%v, Eth0PortSpeed=%v, ExtIp=%v, Hostname=%v, InactiveWiredVlans=%v, Ip=%v, LastHostname=%v, LldpMgmtAddr=%v, LldpPortDesc=%v, LldpPortId=%v, LldpPowerAllocated=%v, LldpPowerDraw=%v, LldpSystemDesc=%v, LldpSystemName=%v, Mac=%v, Model=%v, MxedgeId=%v, MxedgeIds=%v, MxtunnelStatus=%v, OrgId=%v, PowerConstrained=%v, PowerOpmode=%v, SiteId=%v, Sku=%v, Timestamp=%v, Uptime=%v, Version=%v, Wlans=%v, AdditionalProperties=%v]",
    	a.Band24Bandwidth, a.Band24Channel, a.Band24Power, a.Band5Bandwidth, a.Band5Channel, a.Band5Power, a.Band6Bandwidth, a.Band6Channel, a.Band6Power, a.Eth0PortSpeed, a.ExtIp, a.Hostname, a.InactiveWiredVlans, a.Ip, a.LastHostname, a.LldpMgmtAddr, a.LldpPortDesc, a.LldpPortId, a.LldpPowerAllocated, a.LldpPowerDraw, a.LldpSystemDesc, a.LldpSystemName, a.Mac, a.Model, a.MxedgeId, a.MxedgeIds, a.MxtunnelStatus, a.OrgId, a.PowerConstrained, a.PowerOpmode, a.SiteId, a.Sku, a.Timestamp, a.Uptime, a.Version, a.Wlans, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ApSearch.
// It customizes the JSON marshaling process for ApSearch objects.
func (a ApSearch) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(a.AdditionalProperties,
        "band_24_bandwidth", "band_24_channel", "band_24_power", "band_5_bandwidth", "band_5_channel", "band_5_power", "band_6_bandwidth", "band_6_channel", "band_6_power", "eth0_port_speed", "ext_ip", "hostname", "inactive_wired_vlans", "ip", "last_hostname", "lldp_mgmt_addr", "lldp_port_desc", "lldp_port_id", "lldp_power_allocated", "lldp_power_draw", "lldp_system_desc", "lldp_system_name", "mac", "model", "mxedge_id", "mxedge_ids", "mxtunnel_status", "org_id", "power_constrained", "power_opmode", "site_id", "sku", "timestamp", "uptime", "version", "wlans"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(a.toMap())
}

// toMap converts the ApSearch object to a map representation for JSON marshaling.
func (a ApSearch) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, a.AdditionalProperties)
    if a.Band24Bandwidth != nil {
        structMap["band_24_bandwidth"] = a.Band24Bandwidth
    }
    if a.Band24Channel != nil {
        structMap["band_24_channel"] = a.Band24Channel
    }
    if a.Band24Power != nil {
        structMap["band_24_power"] = a.Band24Power
    }
    if a.Band5Bandwidth != nil {
        structMap["band_5_bandwidth"] = a.Band5Bandwidth
    }
    if a.Band5Channel != nil {
        structMap["band_5_channel"] = a.Band5Channel
    }
    if a.Band5Power != nil {
        structMap["band_5_power"] = a.Band5Power
    }
    if a.Band6Bandwidth != nil {
        structMap["band_6_bandwidth"] = a.Band6Bandwidth
    }
    if a.Band6Channel != nil {
        structMap["band_6_channel"] = a.Band6Channel
    }
    if a.Band6Power != nil {
        structMap["band_6_power"] = a.Band6Power
    }
    if a.Eth0PortSpeed != nil {
        structMap["eth0_port_speed"] = a.Eth0PortSpeed
    }
    if a.ExtIp != nil {
        structMap["ext_ip"] = a.ExtIp
    }
    if a.Hostname != nil {
        structMap["hostname"] = a.Hostname
    }
    if a.InactiveWiredVlans != nil {
        structMap["inactive_wired_vlans"] = a.InactiveWiredVlans
    }
    if a.Ip != nil {
        structMap["ip"] = a.Ip
    }
    if a.LastHostname != nil {
        structMap["last_hostname"] = a.LastHostname
    }
    if a.LldpMgmtAddr != nil {
        structMap["lldp_mgmt_addr"] = a.LldpMgmtAddr
    }
    if a.LldpPortDesc != nil {
        structMap["lldp_port_desc"] = a.LldpPortDesc
    }
    if a.LldpPortId != nil {
        structMap["lldp_port_id"] = a.LldpPortId
    }
    if a.LldpPowerAllocated != nil {
        structMap["lldp_power_allocated"] = a.LldpPowerAllocated
    }
    if a.LldpPowerDraw != nil {
        structMap["lldp_power_draw"] = a.LldpPowerDraw
    }
    if a.LldpSystemDesc != nil {
        structMap["lldp_system_desc"] = a.LldpSystemDesc
    }
    if a.LldpSystemName != nil {
        structMap["lldp_system_name"] = a.LldpSystemName
    }
    if a.Mac != nil {
        structMap["mac"] = a.Mac
    }
    if a.Model != nil {
        structMap["model"] = a.Model
    }
    if a.MxedgeId != nil {
        structMap["mxedge_id"] = a.MxedgeId
    }
    if a.MxedgeIds != nil {
        structMap["mxedge_ids"] = a.MxedgeIds
    }
    structMap["mxtunnel_status"] = a.MxtunnelStatus
    if a.OrgId != nil {
        structMap["org_id"] = a.OrgId
    }
    structMap["power_constrained"] = a.PowerConstrained
    structMap["power_opmode"] = a.PowerOpmode
    if a.SiteId != nil {
        structMap["site_id"] = a.SiteId
    }
    if a.Sku != nil {
        structMap["sku"] = a.Sku
    }
    if a.Timestamp != nil {
        structMap["timestamp"] = a.Timestamp
    }
    if a.Uptime != nil {
        structMap["uptime"] = a.Uptime
    }
    if a.Version != nil {
        structMap["version"] = a.Version
    }
    structMap["wlans"] = a.Wlans
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ApSearch.
// It customizes the JSON unmarshaling process for ApSearch objects.
func (a *ApSearch) UnmarshalJSON(input []byte) error {
    var temp tempApSearch
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "band_24_bandwidth", "band_24_channel", "band_24_power", "band_5_bandwidth", "band_5_channel", "band_5_power", "band_6_bandwidth", "band_6_channel", "band_6_power", "eth0_port_speed", "ext_ip", "hostname", "inactive_wired_vlans", "ip", "last_hostname", "lldp_mgmt_addr", "lldp_port_desc", "lldp_port_id", "lldp_power_allocated", "lldp_power_draw", "lldp_system_desc", "lldp_system_name", "mac", "model", "mxedge_id", "mxedge_ids", "mxtunnel_status", "org_id", "power_constrained", "power_opmode", "site_id", "sku", "timestamp", "uptime", "version", "wlans")
    if err != nil {
    	return err
    }
    a.AdditionalProperties = additionalProperties
    
    a.Band24Bandwidth = temp.Band24Bandwidth
    a.Band24Channel = temp.Band24Channel
    a.Band24Power = temp.Band24Power
    a.Band5Bandwidth = temp.Band5Bandwidth
    a.Band5Channel = temp.Band5Channel
    a.Band5Power = temp.Band5Power
    a.Band6Bandwidth = temp.Band6Bandwidth
    a.Band6Channel = temp.Band6Channel
    a.Band6Power = temp.Band6Power
    a.Eth0PortSpeed = temp.Eth0PortSpeed
    a.ExtIp = temp.ExtIp
    a.Hostname = temp.Hostname
    a.InactiveWiredVlans = temp.InactiveWiredVlans
    a.Ip = temp.Ip
    a.LastHostname = temp.LastHostname
    a.LldpMgmtAddr = temp.LldpMgmtAddr
    a.LldpPortDesc = temp.LldpPortDesc
    a.LldpPortId = temp.LldpPortId
    a.LldpPowerAllocated = temp.LldpPowerAllocated
    a.LldpPowerDraw = temp.LldpPowerDraw
    a.LldpSystemDesc = temp.LldpSystemDesc
    a.LldpSystemName = temp.LldpSystemName
    a.Mac = temp.Mac
    a.Model = temp.Model
    a.MxedgeId = temp.MxedgeId
    a.MxedgeIds = temp.MxedgeIds
    a.MxtunnelStatus = *temp.MxtunnelStatus
    a.OrgId = temp.OrgId
    a.PowerConstrained = *temp.PowerConstrained
    a.PowerOpmode = *temp.PowerOpmode
    a.SiteId = temp.SiteId
    a.Sku = temp.Sku
    a.Timestamp = temp.Timestamp
    a.Uptime = temp.Uptime
    a.Version = temp.Version
    a.Wlans = *temp.Wlans
    return nil
}

// tempApSearch is a temporary struct used for validating the fields of ApSearch.
type tempApSearch  struct {
    Band24Bandwidth    *string         `json:"band_24_bandwidth,omitempty"`
    Band24Channel      *int            `json:"band_24_channel,omitempty"`
    Band24Power        *int            `json:"band_24_power,omitempty"`
    Band5Bandwidth     *string         `json:"band_5_bandwidth,omitempty"`
    Band5Channel       *int            `json:"band_5_channel,omitempty"`
    Band5Power         *int            `json:"band_5_power,omitempty"`
    Band6Bandwidth     *string         `json:"band_6_bandwidth,omitempty"`
    Band6Channel       *int            `json:"band_6_channel,omitempty"`
    Band6Power         *int            `json:"band_6_power,omitempty"`
    Eth0PortSpeed      *int            `json:"eth0_port_speed,omitempty"`
    ExtIp              *string         `json:"ext_ip,omitempty"`
    Hostname           []string        `json:"hostname,omitempty"`
    InactiveWiredVlans []int           `json:"inactive_wired_vlans,omitempty"`
    Ip                 *string         `json:"ip,omitempty"`
    LastHostname       *string         `json:"last_hostname,omitempty"`
    LldpMgmtAddr       *string         `json:"lldp_mgmt_addr,omitempty"`
    LldpPortDesc       *string         `json:"lldp_port_desc,omitempty"`
    LldpPortId         *string         `json:"lldp_port_id,omitempty"`
    LldpPowerAllocated *int            `json:"lldp_power_allocated,omitempty"`
    LldpPowerDraw      *int            `json:"lldp_power_draw,omitempty"`
    LldpSystemDesc     *string         `json:"lldp_system_desc,omitempty"`
    LldpSystemName     *string         `json:"lldp_system_name,omitempty"`
    Mac                *string         `json:"mac,omitempty"`
    Model              *string         `json:"model,omitempty"`
    MxedgeId           *string         `json:"mxedge_id,omitempty"`
    MxedgeIds          *string         `json:"mxedge_ids,omitempty"`
    MxtunnelStatus     *string         `json:"mxtunnel_status"`
    OrgId              *uuid.UUID      `json:"org_id,omitempty"`
    PowerConstrained   *bool           `json:"power_constrained"`
    PowerOpmode        *string         `json:"power_opmode"`
    SiteId             *uuid.UUID      `json:"site_id,omitempty"`
    Sku                *string         `json:"sku,omitempty"`
    Timestamp          *float64        `json:"timestamp,omitempty"`
    Uptime             *int            `json:"uptime,omitempty"`
    Version            *string         `json:"version,omitempty"`
    Wlans              *[]ApSearchWlan `json:"wlans"`
}

func (a *tempApSearch) validate() error {
    var errs []string
    if a.MxtunnelStatus == nil {
        errs = append(errs, "required field `mxtunnel_status` is missing for type `ap_search`")
    }
    if a.PowerConstrained == nil {
        errs = append(errs, "required field `power_constrained` is missing for type `ap_search`")
    }
    if a.PowerOpmode == nil {
        errs = append(errs, "required field `power_opmode` is missing for type `ap_search`")
    }
    if a.Wlans == nil {
        errs = append(errs, "required field `wlans` is missing for type `ap_search`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
