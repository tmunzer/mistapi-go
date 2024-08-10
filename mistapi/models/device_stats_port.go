package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// DeviceStatsPort represents a DeviceStatsPort struct.
type DeviceStatsPort struct {
    // Indicates if interface is active/inactive
    Active               *bool                         `json:"active,omitempty"`
    // if `up`==`true` and has Authenticator role. enum: `authenticated`, `authenticating`, `held`, `init`
    AuthState            *SwitchPortStatsAuthStateEnum `json:"auth_state,omitempty"`
    ForSite              *bool                         `json:"for_site,omitempty"`
    // indicates full or half duplex
    FullDuplex           *bool                         `json:"full_duplex,omitempty"`
    // Last sampled jitter of the interface
    Jitter               *float64                      `json:"jitter,omitempty"`
    // Last sampled latency of the interface
    Latency              *float64                      `json:"latency,omitempty"`
    // Last sampled loss of the interface
    Loss                 *float64                      `json:"loss,omitempty"`
    // LTE ICCID value, Check for null/empty
    LteIccid             Optional[string]              `json:"lte_iccid"`
    // LTE IMEI value, Check for null/empty
    LteImei              Optional[string]              `json:"lte_imei"`
    // LTE IMSI value, Check for null/empty
    LteImsi              Optional[string]              `json:"lte_imsi"`
    // Number of mac addresses in the forwarding table
    MacCount             *int                          `json:"mac_count,omitempty"`
    // Limit on number of dynamically learned macs
    MacLimit             *int                          `json:"mac_limit,omitempty"`
    // chassis identifier of the chassis type listed
    NeighborMac          string                        `json:"neighbor_mac"`
    // description supplied by the system on the interface E.g. “GigabitEthernet2/0/39”
    NeighborPortDesc     *string                       `json:"neighbor_port_desc,omitempty"`
    // name supplied by the system on the interface E.g. neighbor system name E.g. “Kumar-Acc-SW.mist.local”
    NeighborSystemName   *string                       `json:"neighbor_system_name,omitempty"`
    // is the POE configured not be disabled.
    PoeDisabled          *bool                         `json:"poe_disabled,omitempty"`
    // enum: `802.3af`, `802.3at`, `802.3bt`
    PoeMode              *SwitchPortStatsPoeModeEnum   `json:"poe_mode,omitempty"`
    // is the device attached to POE
    PoeOn                *bool                         `json:"poe_on,omitempty"`
    PortId               string                        `json:"port_id"`
    // interface mac address
    PortMac              string                        `json:"port_mac"`
    // gateway port usage. enum: `lan`
    PortUsage            *SwitchPortStatsPortUsageEnum `json:"port_usage,omitempty"`
    // Amount of power being used by the interface at the time the command is executed. Unit in watts.
    PowerDraw            *float64                      `json:"power_draw,omitempty"`
    // Broadcast input packets
    RxBcastPkts          *int                          `json:"rx_bcast_pkts,omitempty"`
    // Input rate
    RxBps                *int                          `json:"rx_bps,omitempty"`
    // rx bytes
    RxBytes              int64                         `json:"rx_bytes"`
    // Input errors
    RxErrors             *int                          `json:"rx_errors,omitempty"`
    // Multicast input packets
    RxMcastPkts          *int                          `json:"rx_mcast_pkts,omitempty"`
    // rx packets
    RxPkts               int                           `json:"rx_pkts"`
    // port speed
    Speed                *int                          `json:"speed,omitempty"`
    // if `up`==`true`. enum: `alternate`, `backup`, `designated`, `root`, `root-prevented`
    StpRole              *SwitchPortStatsStpRoleEnum   `json:"stp_role,omitempty"`
    // if `up`==`true`. enum: `blocking`, `disabled`, `forwarding`, `learning`, `listening`
    StpState             *SwitchPortStatsStpStateEnum  `json:"stp_state,omitempty"`
    // Broadcast output packets
    TxBcastPkts          *int                          `json:"tx_bcast_pkts,omitempty"`
    // Output rate
    TxBps                *int                          `json:"tx_bps,omitempty"`
    // tx bytes
    TxBytes              int64                         `json:"tx_bytes"`
    // Output errors
    TxErrors             *int                          `json:"tx_errors,omitempty"`
    // Multicast output packets
    TxMcastPkts          *int                          `json:"tx_mcast_pkts,omitempty"`
    // tx packets
    TxPkts               int                           `json:"tx_pkts"`
    // device type. enum: `ap`, `ble`, `gateway`, `mxedge`, `nac`, `switch`
    Type                 *SwitchPortStatsTypeEnum      `json:"type,omitempty"`
    // indicates if interface is unconfigured
    Unconfigured         *bool                         `json:"unconfigured,omitempty"`
    // indicates if interface is up
    Up                   *bool                         `json:"up,omitempty"`
    // Optic Slot ModelName, Check for null/empty
    XcvrModel            *string                       `json:"xcvr_model,omitempty"`
    // Optic Slot Partnumber, Check for null/empty
    XcvrPartNumber       *string                       `json:"xcvr_part_number,omitempty"`
    // Optic Slot SerialNumber, Check for null/empty
    XcvrSerial           *string                       `json:"xcvr_serial,omitempty"`
    AdditionalProperties map[string]any                `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for DeviceStatsPort.
// It customizes the JSON marshaling process for DeviceStatsPort objects.
func (d DeviceStatsPort) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(d.toMap())
}

// toMap converts the DeviceStatsPort object to a map representation for JSON marshaling.
func (d DeviceStatsPort) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, d.AdditionalProperties)
    if d.Active != nil {
        structMap["active"] = d.Active
    }
    if d.AuthState != nil {
        structMap["auth_state"] = d.AuthState
    }
    if d.ForSite != nil {
        structMap["for_site"] = d.ForSite
    }
    if d.FullDuplex != nil {
        structMap["full_duplex"] = d.FullDuplex
    }
    if d.Jitter != nil {
        structMap["jitter"] = d.Jitter
    }
    if d.Latency != nil {
        structMap["latency"] = d.Latency
    }
    if d.Loss != nil {
        structMap["loss"] = d.Loss
    }
    if d.LteIccid.IsValueSet() {
        if d.LteIccid.Value() != nil {
            structMap["lte_iccid"] = d.LteIccid.Value()
        } else {
            structMap["lte_iccid"] = nil
        }
    }
    if d.LteImei.IsValueSet() {
        if d.LteImei.Value() != nil {
            structMap["lte_imei"] = d.LteImei.Value()
        } else {
            structMap["lte_imei"] = nil
        }
    }
    if d.LteImsi.IsValueSet() {
        if d.LteImsi.Value() != nil {
            structMap["lte_imsi"] = d.LteImsi.Value()
        } else {
            structMap["lte_imsi"] = nil
        }
    }
    if d.MacCount != nil {
        structMap["mac_count"] = d.MacCount
    }
    if d.MacLimit != nil {
        structMap["mac_limit"] = d.MacLimit
    }
    structMap["neighbor_mac"] = d.NeighborMac
    if d.NeighborPortDesc != nil {
        structMap["neighbor_port_desc"] = d.NeighborPortDesc
    }
    if d.NeighborSystemName != nil {
        structMap["neighbor_system_name"] = d.NeighborSystemName
    }
    if d.PoeDisabled != nil {
        structMap["poe_disabled"] = d.PoeDisabled
    }
    if d.PoeMode != nil {
        structMap["poe_mode"] = d.PoeMode
    }
    if d.PoeOn != nil {
        structMap["poe_on"] = d.PoeOn
    }
    structMap["port_id"] = d.PortId
    structMap["port_mac"] = d.PortMac
    if d.PortUsage != nil {
        structMap["port_usage"] = d.PortUsage
    }
    if d.PowerDraw != nil {
        structMap["power_draw"] = d.PowerDraw
    }
    if d.RxBcastPkts != nil {
        structMap["rx_bcast_pkts"] = d.RxBcastPkts
    }
    if d.RxBps != nil {
        structMap["rx_bps"] = d.RxBps
    }
    structMap["rx_bytes"] = d.RxBytes
    if d.RxErrors != nil {
        structMap["rx_errors"] = d.RxErrors
    }
    if d.RxMcastPkts != nil {
        structMap["rx_mcast_pkts"] = d.RxMcastPkts
    }
    structMap["rx_pkts"] = d.RxPkts
    if d.Speed != nil {
        structMap["speed"] = d.Speed
    }
    if d.StpRole != nil {
        structMap["stp_role"] = d.StpRole
    }
    if d.StpState != nil {
        structMap["stp_state"] = d.StpState
    }
    if d.TxBcastPkts != nil {
        structMap["tx_bcast_pkts"] = d.TxBcastPkts
    }
    if d.TxBps != nil {
        structMap["tx_bps"] = d.TxBps
    }
    structMap["tx_bytes"] = d.TxBytes
    if d.TxErrors != nil {
        structMap["tx_errors"] = d.TxErrors
    }
    if d.TxMcastPkts != nil {
        structMap["tx_mcast_pkts"] = d.TxMcastPkts
    }
    structMap["tx_pkts"] = d.TxPkts
    if d.Type != nil {
        structMap["type"] = d.Type
    }
    if d.Unconfigured != nil {
        structMap["unconfigured"] = d.Unconfigured
    }
    if d.Up != nil {
        structMap["up"] = d.Up
    }
    if d.XcvrModel != nil {
        structMap["xcvr_model"] = d.XcvrModel
    }
    if d.XcvrPartNumber != nil {
        structMap["xcvr_part_number"] = d.XcvrPartNumber
    }
    if d.XcvrSerial != nil {
        structMap["xcvr_serial"] = d.XcvrSerial
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for DeviceStatsPort.
// It customizes the JSON unmarshaling process for DeviceStatsPort objects.
func (d *DeviceStatsPort) UnmarshalJSON(input []byte) error {
    var temp tempDeviceStatsPort
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "active", "auth_state", "for_site", "full_duplex", "jitter", "latency", "loss", "lte_iccid", "lte_imei", "lte_imsi", "mac_count", "mac_limit", "neighbor_mac", "neighbor_port_desc", "neighbor_system_name", "poe_disabled", "poe_mode", "poe_on", "port_id", "port_mac", "port_usage", "power_draw", "rx_bcast_pkts", "rx_bps", "rx_bytes", "rx_errors", "rx_mcast_pkts", "rx_pkts", "speed", "stp_role", "stp_state", "tx_bcast_pkts", "tx_bps", "tx_bytes", "tx_errors", "tx_mcast_pkts", "tx_pkts", "type", "unconfigured", "up", "xcvr_model", "xcvr_part_number", "xcvr_serial")
    if err != nil {
    	return err
    }
    
    d.AdditionalProperties = additionalProperties
    d.Active = temp.Active
    d.AuthState = temp.AuthState
    d.ForSite = temp.ForSite
    d.FullDuplex = temp.FullDuplex
    d.Jitter = temp.Jitter
    d.Latency = temp.Latency
    d.Loss = temp.Loss
    d.LteIccid = temp.LteIccid
    d.LteImei = temp.LteImei
    d.LteImsi = temp.LteImsi
    d.MacCount = temp.MacCount
    d.MacLimit = temp.MacLimit
    d.NeighborMac = *temp.NeighborMac
    d.NeighborPortDesc = temp.NeighborPortDesc
    d.NeighborSystemName = temp.NeighborSystemName
    d.PoeDisabled = temp.PoeDisabled
    d.PoeMode = temp.PoeMode
    d.PoeOn = temp.PoeOn
    d.PortId = *temp.PortId
    d.PortMac = *temp.PortMac
    d.PortUsage = temp.PortUsage
    d.PowerDraw = temp.PowerDraw
    d.RxBcastPkts = temp.RxBcastPkts
    d.RxBps = temp.RxBps
    d.RxBytes = *temp.RxBytes
    d.RxErrors = temp.RxErrors
    d.RxMcastPkts = temp.RxMcastPkts
    d.RxPkts = *temp.RxPkts
    d.Speed = temp.Speed
    d.StpRole = temp.StpRole
    d.StpState = temp.StpState
    d.TxBcastPkts = temp.TxBcastPkts
    d.TxBps = temp.TxBps
    d.TxBytes = *temp.TxBytes
    d.TxErrors = temp.TxErrors
    d.TxMcastPkts = temp.TxMcastPkts
    d.TxPkts = *temp.TxPkts
    d.Type = temp.Type
    d.Unconfigured = temp.Unconfigured
    d.Up = temp.Up
    d.XcvrModel = temp.XcvrModel
    d.XcvrPartNumber = temp.XcvrPartNumber
    d.XcvrSerial = temp.XcvrSerial
    return nil
}

// tempDeviceStatsPort is a temporary struct used for validating the fields of DeviceStatsPort.
type tempDeviceStatsPort  struct {
    Active             *bool                         `json:"active,omitempty"`
    AuthState          *SwitchPortStatsAuthStateEnum `json:"auth_state,omitempty"`
    ForSite            *bool                         `json:"for_site,omitempty"`
    FullDuplex         *bool                         `json:"full_duplex,omitempty"`
    Jitter             *float64                      `json:"jitter,omitempty"`
    Latency            *float64                      `json:"latency,omitempty"`
    Loss               *float64                      `json:"loss,omitempty"`
    LteIccid           Optional[string]              `json:"lte_iccid"`
    LteImei            Optional[string]              `json:"lte_imei"`
    LteImsi            Optional[string]              `json:"lte_imsi"`
    MacCount           *int                          `json:"mac_count,omitempty"`
    MacLimit           *int                          `json:"mac_limit,omitempty"`
    NeighborMac        *string                       `json:"neighbor_mac"`
    NeighborPortDesc   *string                       `json:"neighbor_port_desc,omitempty"`
    NeighborSystemName *string                       `json:"neighbor_system_name,omitempty"`
    PoeDisabled        *bool                         `json:"poe_disabled,omitempty"`
    PoeMode            *SwitchPortStatsPoeModeEnum   `json:"poe_mode,omitempty"`
    PoeOn              *bool                         `json:"poe_on,omitempty"`
    PortId             *string                       `json:"port_id"`
    PortMac            *string                       `json:"port_mac"`
    PortUsage          *SwitchPortStatsPortUsageEnum `json:"port_usage,omitempty"`
    PowerDraw          *float64                      `json:"power_draw,omitempty"`
    RxBcastPkts        *int                          `json:"rx_bcast_pkts,omitempty"`
    RxBps              *int                          `json:"rx_bps,omitempty"`
    RxBytes            *int64                        `json:"rx_bytes"`
    RxErrors           *int                          `json:"rx_errors,omitempty"`
    RxMcastPkts        *int                          `json:"rx_mcast_pkts,omitempty"`
    RxPkts             *int                          `json:"rx_pkts"`
    Speed              *int                          `json:"speed,omitempty"`
    StpRole            *SwitchPortStatsStpRoleEnum   `json:"stp_role,omitempty"`
    StpState           *SwitchPortStatsStpStateEnum  `json:"stp_state,omitempty"`
    TxBcastPkts        *int                          `json:"tx_bcast_pkts,omitempty"`
    TxBps              *int                          `json:"tx_bps,omitempty"`
    TxBytes            *int64                        `json:"tx_bytes"`
    TxErrors           *int                          `json:"tx_errors,omitempty"`
    TxMcastPkts        *int                          `json:"tx_mcast_pkts,omitempty"`
    TxPkts             *int                          `json:"tx_pkts"`
    Type               *SwitchPortStatsTypeEnum      `json:"type,omitempty"`
    Unconfigured       *bool                         `json:"unconfigured,omitempty"`
    Up                 *bool                         `json:"up,omitempty"`
    XcvrModel          *string                       `json:"xcvr_model,omitempty"`
    XcvrPartNumber     *string                       `json:"xcvr_part_number,omitempty"`
    XcvrSerial         *string                       `json:"xcvr_serial,omitempty"`
}

func (d *tempDeviceStatsPort) validate() error {
    var errs []string
    if d.NeighborMac == nil {
        errs = append(errs, "required field `neighbor_mac` is missing for type `device_stats_port`")
    }
    if d.PortId == nil {
        errs = append(errs, "required field `port_id` is missing for type `device_stats_port`")
    }
    if d.PortMac == nil {
        errs = append(errs, "required field `port_mac` is missing for type `device_stats_port`")
    }
    if d.RxBytes == nil {
        errs = append(errs, "required field `rx_bytes` is missing for type `device_stats_port`")
    }
    if d.RxPkts == nil {
        errs = append(errs, "required field `rx_pkts` is missing for type `device_stats_port`")
    }
    if d.TxBytes == nil {
        errs = append(errs, "required field `tx_bytes` is missing for type `device_stats_port`")
    }
    if d.TxPkts == nil {
        errs = append(errs, "required field `tx_pkts` is missing for type `device_stats_port`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
