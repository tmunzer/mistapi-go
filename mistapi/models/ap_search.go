package models

import (
	"encoding/json"
	"github.com/google/uuid"
)

// ApSearch represents a ApSearch struct.
type ApSearch struct {
	// Bandwith of band_24
	Band24Bandwith *string `json:"band_24_bandwith,omitempty"`
	// Channel of band_24
	Band24Channel *int `json:"band_24_channel,omitempty"`
	Band24Power   *int `json:"band_24_power,omitempty"`
	// Bandwith of band_5
	Band5Bandwith *string `json:"band_5_bandwith,omitempty"`
	// Channel of band_5
	Band5Channel  *int    `json:"band_5_channel,omitempty"`
	Band5Power    *int    `json:"band_5_power,omitempty"`
	Band6Bandwith *string `json:"band_6_bandwith,omitempty"`
	// Channel of band_6
	Band6Channel *int `json:"band_6_channel,omitempty"`
	Band6Power   *int `json:"band_6_power,omitempty"`
	// Port speed of eth0
	Eth0PortSpeed *int    `json:"eth0_port_speed,omitempty"`
	ExtIp         *string `json:"ext_ip,omitempty"`
	// partial / full hostname
	Hostname []string `json:"hostname,omitempty"`
	// ip address
	Ip *string `json:"ip,omitempty"`
	// LLDP management ip address
	LldpMgmtAddr *string `json:"lldp_mgmt_addr,omitempty"`
	LldpPortDesc *string `json:"lldp_port_desc,omitempty"`
	// LLDP port id
	LldpPortId         *string `json:"lldp_port_id,omitempty"`
	LldpPowerAllocated *int    `json:"lldp_power_allocated,omitempty"`
	LldpPowerDraw      *int    `json:"lldp_power_draw,omitempty"`
	// LLDP system description
	LldpSystemDesc *string `json:"lldp_system_desc,omitempty"`
	// LLDP system name
	LldpSystemName *string `json:"lldp_system_name,omitempty"`
	// device model
	Mac   *string `json:"mac,omitempty"`
	Model *string `json:"model,omitempty"`
	// Mist Edge id, if AP is connecting to a Mist Edge
	MxedgeId *string `json:"mxedge_id,omitempty"`
	// MxTunnel status
	MxtunnelStatus   *string    `json:"mxtunnel_status,omitempty"`
	OrgId            *uuid.UUID `json:"org_id,omitempty"`
	PowerConstrained *bool      `json:"power_constrained,omitempty"`
	SiteId           *uuid.UUID `json:"site_id,omitempty"`
	Sku              *string    `json:"sku,omitempty"`
	Timestamp        *float64   `json:"timestamp,omitempty"`
	Uptime           *int       `json:"uptime,omitempty"`
	// version
	Version              *string        `json:"version,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ApSearch.
// It customizes the JSON marshaling process for ApSearch objects.
func (a ApSearch) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(a.toMap())
}

// toMap converts the ApSearch object to a map representation for JSON marshaling.
func (a ApSearch) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, a.AdditionalProperties)
	if a.Band24Bandwith != nil {
		structMap["band_24_bandwith"] = a.Band24Bandwith
	}
	if a.Band24Channel != nil {
		structMap["band_24_channel"] = a.Band24Channel
	}
	if a.Band24Power != nil {
		structMap["band_24_power"] = a.Band24Power
	}
	if a.Band5Bandwith != nil {
		structMap["band_5_bandwith"] = a.Band5Bandwith
	}
	if a.Band5Channel != nil {
		structMap["band_5_channel"] = a.Band5Channel
	}
	if a.Band5Power != nil {
		structMap["band_5_power"] = a.Band5Power
	}
	if a.Band6Bandwith != nil {
		structMap["band_6_bandwith"] = a.Band6Bandwith
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
	if a.Ip != nil {
		structMap["ip"] = a.Ip
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
	if a.MxtunnelStatus != nil {
		structMap["mxtunnel_status"] = a.MxtunnelStatus
	}
	if a.OrgId != nil {
		structMap["org_id"] = a.OrgId
	}
	if a.PowerConstrained != nil {
		structMap["power_constrained"] = a.PowerConstrained
	}
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
	additionalProperties, err := UnmarshalAdditionalProperties(input, "band_24_bandwith", "band_24_channel", "band_24_power", "band_5_bandwith", "band_5_channel", "band_5_power", "band_6_bandwith", "band_6_channel", "band_6_power", "eth0_port_speed", "ext_ip", "hostname", "ip", "lldp_mgmt_addr", "lldp_port_desc", "lldp_port_id", "lldp_power_allocated", "lldp_power_draw", "lldp_system_desc", "lldp_system_name", "mac", "model", "mxedge_id", "mxtunnel_status", "org_id", "power_constrained", "site_id", "sku", "timestamp", "uptime", "version")
	if err != nil {
		return err
	}

	a.AdditionalProperties = additionalProperties
	a.Band24Bandwith = temp.Band24Bandwith
	a.Band24Channel = temp.Band24Channel
	a.Band24Power = temp.Band24Power
	a.Band5Bandwith = temp.Band5Bandwith
	a.Band5Channel = temp.Band5Channel
	a.Band5Power = temp.Band5Power
	a.Band6Bandwith = temp.Band6Bandwith
	a.Band6Channel = temp.Band6Channel
	a.Band6Power = temp.Band6Power
	a.Eth0PortSpeed = temp.Eth0PortSpeed
	a.ExtIp = temp.ExtIp
	a.Hostname = temp.Hostname
	a.Ip = temp.Ip
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
	a.MxtunnelStatus = temp.MxtunnelStatus
	a.OrgId = temp.OrgId
	a.PowerConstrained = temp.PowerConstrained
	a.SiteId = temp.SiteId
	a.Sku = temp.Sku
	a.Timestamp = temp.Timestamp
	a.Uptime = temp.Uptime
	a.Version = temp.Version
	return nil
}

// tempApSearch is a temporary struct used for validating the fields of ApSearch.
type tempApSearch struct {
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
}
