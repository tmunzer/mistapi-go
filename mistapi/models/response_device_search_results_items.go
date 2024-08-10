package models

import (
    "encoding/json"
    "github.com/google/uuid"
)

// ResponseDeviceSearchResultsItems represents a ResponseDeviceSearchResultsItems struct.
type ResponseDeviceSearchResultsItems struct {
    // Bandwith of band_24
    Band24Bandwith       *string        `json:"band_24_bandwith,omitempty"`
    // Channel of band_24
    Band24Channel        *int           `json:"band_24_channel,omitempty"`
    Band24Power          *int           `json:"band_24_power,omitempty"`
    // Bandwith of band_5
    Band5Bandwith        *string        `json:"band_5_bandwith,omitempty"`
    // Channel of band_5
    Band5Channel         *int           `json:"band_5_channel,omitempty"`
    Band5Power           *int           `json:"band_5_power,omitempty"`
    Band6Bandwith        *string        `json:"band_6_bandwith,omitempty"`
    // Channel of band_6
    Band6Channel         *int           `json:"band_6_channel,omitempty"`
    Band6Power           *int           `json:"band_6_power,omitempty"`
    // Port speed of eth0
    Eth0PortSpeed        *int           `json:"eth0_port_speed,omitempty"`
    ExtIp                *string        `json:"ext_ip,omitempty"`
    // partial / full hostname
    Hostname             []string       `json:"hostname,omitempty"`
    // ip address
    Ip                   *string        `json:"ip,omitempty"`
    // LLDP management ip address
    LldpMgmtAddr         *string        `json:"lldp_mgmt_addr,omitempty"`
    LldpPortDesc         *string        `json:"lldp_port_desc,omitempty"`
    // LLDP port id
    LldpPortId           *string        `json:"lldp_port_id,omitempty"`
    LldpPowerAllocated   *int           `json:"lldp_power_allocated,omitempty"`
    LldpPowerDraw        *int           `json:"lldp_power_draw,omitempty"`
    // LLDP system description
    LldpSystemDesc       *string        `json:"lldp_system_desc,omitempty"`
    // LLDP system name
    LldpSystemName       *string        `json:"lldp_system_name,omitempty"`
    // device model
    Mac                  *string        `json:"mac,omitempty"`
    Model                *string        `json:"model,omitempty"`
    // Mist Edge id, if AP is connecting to a Mist Edge
    MxedgeId             *string        `json:"mxedge_id,omitempty"`
    // MxTunnel status
    MxtunnelStatus       *string        `json:"mxtunnel_status,omitempty"`
    OrgId                *uuid.UUID     `json:"org_id,omitempty"`
    PowerConstrained     *bool          `json:"power_constrained,omitempty"`
    SiteId               *uuid.UUID     `json:"site_id,omitempty"`
    Sku                  *string        `json:"sku,omitempty"`
    Timestamp            *float64       `json:"timestamp,omitempty"`
    Uptime               *int           `json:"uptime,omitempty"`
    // version
    Version              *string        `json:"version,omitempty"`
    LastHostname         *string        `json:"last_hostname,omitempty"`
    NumMembers           *int           `json:"num_members,omitempty"`
    Type                 *string        `json:"type,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseDeviceSearchResultsItems.
// It customizes the JSON marshaling process for ResponseDeviceSearchResultsItems objects.
func (r ResponseDeviceSearchResultsItems) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseDeviceSearchResultsItems object to a map representation for JSON marshaling.
func (r ResponseDeviceSearchResultsItems) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    if r.Band24Bandwith != nil {
        structMap["band_24_bandwith"] = r.Band24Bandwith
    }
    if r.Band24Channel != nil {
        structMap["band_24_channel"] = r.Band24Channel
    }
    if r.Band24Power != nil {
        structMap["band_24_power"] = r.Band24Power
    }
    if r.Band5Bandwith != nil {
        structMap["band_5_bandwith"] = r.Band5Bandwith
    }
    if r.Band5Channel != nil {
        structMap["band_5_channel"] = r.Band5Channel
    }
    if r.Band5Power != nil {
        structMap["band_5_power"] = r.Band5Power
    }
    if r.Band6Bandwith != nil {
        structMap["band_6_bandwith"] = r.Band6Bandwith
    }
    if r.Band6Channel != nil {
        structMap["band_6_channel"] = r.Band6Channel
    }
    if r.Band6Power != nil {
        structMap["band_6_power"] = r.Band6Power
    }
    if r.Eth0PortSpeed != nil {
        structMap["eth0_port_speed"] = r.Eth0PortSpeed
    }
    if r.ExtIp != nil {
        structMap["ext_ip"] = r.ExtIp
    }
    if r.Hostname != nil {
        structMap["hostname"] = r.Hostname
    }
    if r.Ip != nil {
        structMap["ip"] = r.Ip
    }
    if r.LldpMgmtAddr != nil {
        structMap["lldp_mgmt_addr"] = r.LldpMgmtAddr
    }
    if r.LldpPortDesc != nil {
        structMap["lldp_port_desc"] = r.LldpPortDesc
    }
    if r.LldpPortId != nil {
        structMap["lldp_port_id"] = r.LldpPortId
    }
    if r.LldpPowerAllocated != nil {
        structMap["lldp_power_allocated"] = r.LldpPowerAllocated
    }
    if r.LldpPowerDraw != nil {
        structMap["lldp_power_draw"] = r.LldpPowerDraw
    }
    if r.LldpSystemDesc != nil {
        structMap["lldp_system_desc"] = r.LldpSystemDesc
    }
    if r.LldpSystemName != nil {
        structMap["lldp_system_name"] = r.LldpSystemName
    }
    if r.Mac != nil {
        structMap["mac"] = r.Mac
    }
    if r.Model != nil {
        structMap["model"] = r.Model
    }
    if r.MxedgeId != nil {
        structMap["mxedge_id"] = r.MxedgeId
    }
    if r.MxtunnelStatus != nil {
        structMap["mxtunnel_status"] = r.MxtunnelStatus
    }
    if r.OrgId != nil {
        structMap["org_id"] = r.OrgId
    }
    if r.PowerConstrained != nil {
        structMap["power_constrained"] = r.PowerConstrained
    }
    if r.SiteId != nil {
        structMap["site_id"] = r.SiteId
    }
    if r.Sku != nil {
        structMap["sku"] = r.Sku
    }
    if r.Timestamp != nil {
        structMap["timestamp"] = r.Timestamp
    }
    if r.Uptime != nil {
        structMap["uptime"] = r.Uptime
    }
    if r.Version != nil {
        structMap["version"] = r.Version
    }
    if r.LastHostname != nil {
        structMap["last_hostname"] = r.LastHostname
    }
    if r.NumMembers != nil {
        structMap["num_members"] = r.NumMembers
    }
    if r.Type != nil {
        structMap["type"] = r.Type
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseDeviceSearchResultsItems.
// It customizes the JSON unmarshaling process for ResponseDeviceSearchResultsItems objects.
func (r *ResponseDeviceSearchResultsItems) UnmarshalJSON(input []byte) error {
    var temp tempResponseDeviceSearchResultsItems
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "band_24_bandwith", "band_24_channel", "band_24_power", "band_5_bandwith", "band_5_channel", "band_5_power", "band_6_bandwith", "band_6_channel", "band_6_power", "eth0_port_speed", "ext_ip", "hostname", "ip", "lldp_mgmt_addr", "lldp_port_desc", "lldp_port_id", "lldp_power_allocated", "lldp_power_draw", "lldp_system_desc", "lldp_system_name", "mac", "model", "mxedge_id", "mxtunnel_status", "org_id", "power_constrained", "site_id", "sku", "timestamp", "uptime", "version", "last_hostname", "num_members", "type")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.Band24Bandwith = temp.Band24Bandwith
    r.Band24Channel = temp.Band24Channel
    r.Band24Power = temp.Band24Power
    r.Band5Bandwith = temp.Band5Bandwith
    r.Band5Channel = temp.Band5Channel
    r.Band5Power = temp.Band5Power
    r.Band6Bandwith = temp.Band6Bandwith
    r.Band6Channel = temp.Band6Channel
    r.Band6Power = temp.Band6Power
    r.Eth0PortSpeed = temp.Eth0PortSpeed
    r.ExtIp = temp.ExtIp
    r.Hostname = temp.Hostname
    r.Ip = temp.Ip
    r.LldpMgmtAddr = temp.LldpMgmtAddr
    r.LldpPortDesc = temp.LldpPortDesc
    r.LldpPortId = temp.LldpPortId
    r.LldpPowerAllocated = temp.LldpPowerAllocated
    r.LldpPowerDraw = temp.LldpPowerDraw
    r.LldpSystemDesc = temp.LldpSystemDesc
    r.LldpSystemName = temp.LldpSystemName
    r.Mac = temp.Mac
    r.Model = temp.Model
    r.MxedgeId = temp.MxedgeId
    r.MxtunnelStatus = temp.MxtunnelStatus
    r.OrgId = temp.OrgId
    r.PowerConstrained = temp.PowerConstrained
    r.SiteId = temp.SiteId
    r.Sku = temp.Sku
    r.Timestamp = temp.Timestamp
    r.Uptime = temp.Uptime
    r.Version = temp.Version
    r.LastHostname = temp.LastHostname
    r.NumMembers = temp.NumMembers
    r.Type = temp.Type
    return nil
}

// tempResponseDeviceSearchResultsItems is a temporary struct used for validating the fields of ResponseDeviceSearchResultsItems.
type tempResponseDeviceSearchResultsItems  struct {
    Band24Bandwith     *string    `json:"band_24_bandwith,omitempty"`
    Band24Channel      *int       `json:"band_24_channel,omitempty"`
    Band24Power        *int       `json:"band_24_power,omitempty"`
    Band5Bandwith      *string    `json:"band_5_bandwith,omitempty"`
    Band5Channel       *int       `json:"band_5_channel,omitempty"`
    Band5Power         *int       `json:"band_5_power,omitempty"`
    Band6Bandwith      *string    `json:"band_6_bandwith,omitempty"`
    Band6Channel       *int       `json:"band_6_channel,omitempty"`
    Band6Power         *int       `json:"band_6_power,omitempty"`
    Eth0PortSpeed      *int       `json:"eth0_port_speed,omitempty"`
    ExtIp              *string    `json:"ext_ip,omitempty"`
    Hostname           []string   `json:"hostname,omitempty"`
    Ip                 *string    `json:"ip,omitempty"`
    LldpMgmtAddr       *string    `json:"lldp_mgmt_addr,omitempty"`
    LldpPortDesc       *string    `json:"lldp_port_desc,omitempty"`
    LldpPortId         *string    `json:"lldp_port_id,omitempty"`
    LldpPowerAllocated *int       `json:"lldp_power_allocated,omitempty"`
    LldpPowerDraw      *int       `json:"lldp_power_draw,omitempty"`
    LldpSystemDesc     *string    `json:"lldp_system_desc,omitempty"`
    LldpSystemName     *string    `json:"lldp_system_name,omitempty"`
    Mac                *string    `json:"mac,omitempty"`
    Model              *string    `json:"model,omitempty"`
    MxedgeId           *string    `json:"mxedge_id,omitempty"`
    MxtunnelStatus     *string    `json:"mxtunnel_status,omitempty"`
    OrgId              *uuid.UUID `json:"org_id,omitempty"`
    PowerConstrained   *bool      `json:"power_constrained,omitempty"`
    SiteId             *uuid.UUID `json:"site_id,omitempty"`
    Sku                *string    `json:"sku,omitempty"`
    Timestamp          *float64   `json:"timestamp,omitempty"`
    Uptime             *int       `json:"uptime,omitempty"`
    Version            *string    `json:"version,omitempty"`
    LastHostname       *string    `json:"last_hostname,omitempty"`
    NumMembers         *int       `json:"num_members,omitempty"`
    Type               *string    `json:"type,omitempty"`
}
