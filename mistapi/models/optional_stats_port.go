package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// OptionalStatsPort represents a OptionalStatsPort struct.
// Port statistics
type OptionalStatsPort struct {
    // Indicates if interface is active/inactive
    Active               *bool                         `json:"active,omitempty"`
    // if `up`==`true` and has Authenticator role. enum: `authenticated`, `authenticating`, `held`, `init`
    AuthState            *StatsSwitchPortAuthStateEnum `json:"auth_state,omitempty"`
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
    PoeMode              *StatsSwitchPortPoeModeEnum   `json:"poe_mode,omitempty"`
    // is the device attached to POE
    PoeOn                *bool                         `json:"poe_on,omitempty"`
    PortId               string                        `json:"port_id"`
    // interface mac address
    PortMac              string                        `json:"port_mac"`
    // gateway port usage. enum: `lan`
    PortUsage            *StatsSwitchPortPortUsageEnum `json:"port_usage,omitempty"`
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
    StpRole              *StatsSwitchPortStpRoleEnum   `json:"stp_role,omitempty"`
    // if `up`==`true`. enum: `blocking`, `disabled`, `forwarding`, `learning`, `listening`
    StpState             *StatsSwitchPortStpStateEnum  `json:"stp_state,omitempty"`
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
    Type                 *StatsSwitchPortTypeEnum      `json:"type,omitempty"`
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
    AdditionalProperties map[string]interface{}        `json:"_"`
}

// String implements the fmt.Stringer interface for OptionalStatsPort,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (o OptionalStatsPort) String() string {
    return fmt.Sprintf(
    	"OptionalStatsPort[Active=%v, AuthState=%v, ForSite=%v, FullDuplex=%v, Jitter=%v, Latency=%v, Loss=%v, LteIccid=%v, LteImei=%v, LteImsi=%v, MacCount=%v, MacLimit=%v, NeighborMac=%v, NeighborPortDesc=%v, NeighborSystemName=%v, PoeDisabled=%v, PoeMode=%v, PoeOn=%v, PortId=%v, PortMac=%v, PortUsage=%v, PowerDraw=%v, RxBcastPkts=%v, RxBps=%v, RxBytes=%v, RxErrors=%v, RxMcastPkts=%v, RxPkts=%v, Speed=%v, StpRole=%v, StpState=%v, TxBcastPkts=%v, TxBps=%v, TxBytes=%v, TxErrors=%v, TxMcastPkts=%v, TxPkts=%v, Type=%v, Unconfigured=%v, Up=%v, XcvrModel=%v, XcvrPartNumber=%v, XcvrSerial=%v, AdditionalProperties=%v]",
    	o.Active, o.AuthState, o.ForSite, o.FullDuplex, o.Jitter, o.Latency, o.Loss, o.LteIccid, o.LteImei, o.LteImsi, o.MacCount, o.MacLimit, o.NeighborMac, o.NeighborPortDesc, o.NeighborSystemName, o.PoeDisabled, o.PoeMode, o.PoeOn, o.PortId, o.PortMac, o.PortUsage, o.PowerDraw, o.RxBcastPkts, o.RxBps, o.RxBytes, o.RxErrors, o.RxMcastPkts, o.RxPkts, o.Speed, o.StpRole, o.StpState, o.TxBcastPkts, o.TxBps, o.TxBytes, o.TxErrors, o.TxMcastPkts, o.TxPkts, o.Type, o.Unconfigured, o.Up, o.XcvrModel, o.XcvrPartNumber, o.XcvrSerial, o.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for OptionalStatsPort.
// It customizes the JSON marshaling process for OptionalStatsPort objects.
func (o OptionalStatsPort) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(o.AdditionalProperties,
        "active", "auth_state", "for_site", "full_duplex", "jitter", "latency", "loss", "lte_iccid", "lte_imei", "lte_imsi", "mac_count", "mac_limit", "neighbor_mac", "neighbor_port_desc", "neighbor_system_name", "poe_disabled", "poe_mode", "poe_on", "port_id", "port_mac", "port_usage", "power_draw", "rx_bcast_pkts", "rx_bps", "rx_bytes", "rx_errors", "rx_mcast_pkts", "rx_pkts", "speed", "stp_role", "stp_state", "tx_bcast_pkts", "tx_bps", "tx_bytes", "tx_errors", "tx_mcast_pkts", "tx_pkts", "type", "unconfigured", "up", "xcvr_model", "xcvr_part_number", "xcvr_serial"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(o.toMap())
}

// toMap converts the OptionalStatsPort object to a map representation for JSON marshaling.
func (o OptionalStatsPort) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, o.AdditionalProperties)
    if o.Active != nil {
        structMap["active"] = o.Active
    }
    if o.AuthState != nil {
        structMap["auth_state"] = o.AuthState
    }
    if o.ForSite != nil {
        structMap["for_site"] = o.ForSite
    }
    if o.FullDuplex != nil {
        structMap["full_duplex"] = o.FullDuplex
    }
    if o.Jitter != nil {
        structMap["jitter"] = o.Jitter
    }
    if o.Latency != nil {
        structMap["latency"] = o.Latency
    }
    if o.Loss != nil {
        structMap["loss"] = o.Loss
    }
    if o.LteIccid.IsValueSet() {
        if o.LteIccid.Value() != nil {
            structMap["lte_iccid"] = o.LteIccid.Value()
        } else {
            structMap["lte_iccid"] = nil
        }
    }
    if o.LteImei.IsValueSet() {
        if o.LteImei.Value() != nil {
            structMap["lte_imei"] = o.LteImei.Value()
        } else {
            structMap["lte_imei"] = nil
        }
    }
    if o.LteImsi.IsValueSet() {
        if o.LteImsi.Value() != nil {
            structMap["lte_imsi"] = o.LteImsi.Value()
        } else {
            structMap["lte_imsi"] = nil
        }
    }
    if o.MacCount != nil {
        structMap["mac_count"] = o.MacCount
    }
    if o.MacLimit != nil {
        structMap["mac_limit"] = o.MacLimit
    }
    structMap["neighbor_mac"] = o.NeighborMac
    if o.NeighborPortDesc != nil {
        structMap["neighbor_port_desc"] = o.NeighborPortDesc
    }
    if o.NeighborSystemName != nil {
        structMap["neighbor_system_name"] = o.NeighborSystemName
    }
    if o.PoeDisabled != nil {
        structMap["poe_disabled"] = o.PoeDisabled
    }
    if o.PoeMode != nil {
        structMap["poe_mode"] = o.PoeMode
    }
    if o.PoeOn != nil {
        structMap["poe_on"] = o.PoeOn
    }
    structMap["port_id"] = o.PortId
    structMap["port_mac"] = o.PortMac
    if o.PortUsage != nil {
        structMap["port_usage"] = o.PortUsage
    }
    if o.PowerDraw != nil {
        structMap["power_draw"] = o.PowerDraw
    }
    if o.RxBcastPkts != nil {
        structMap["rx_bcast_pkts"] = o.RxBcastPkts
    }
    if o.RxBps != nil {
        structMap["rx_bps"] = o.RxBps
    }
    structMap["rx_bytes"] = o.RxBytes
    if o.RxErrors != nil {
        structMap["rx_errors"] = o.RxErrors
    }
    if o.RxMcastPkts != nil {
        structMap["rx_mcast_pkts"] = o.RxMcastPkts
    }
    structMap["rx_pkts"] = o.RxPkts
    if o.Speed != nil {
        structMap["speed"] = o.Speed
    }
    if o.StpRole != nil {
        structMap["stp_role"] = o.StpRole
    }
    if o.StpState != nil {
        structMap["stp_state"] = o.StpState
    }
    if o.TxBcastPkts != nil {
        structMap["tx_bcast_pkts"] = o.TxBcastPkts
    }
    if o.TxBps != nil {
        structMap["tx_bps"] = o.TxBps
    }
    structMap["tx_bytes"] = o.TxBytes
    if o.TxErrors != nil {
        structMap["tx_errors"] = o.TxErrors
    }
    if o.TxMcastPkts != nil {
        structMap["tx_mcast_pkts"] = o.TxMcastPkts
    }
    structMap["tx_pkts"] = o.TxPkts
    if o.Type != nil {
        structMap["type"] = o.Type
    }
    if o.Unconfigured != nil {
        structMap["unconfigured"] = o.Unconfigured
    }
    if o.Up != nil {
        structMap["up"] = o.Up
    }
    if o.XcvrModel != nil {
        structMap["xcvr_model"] = o.XcvrModel
    }
    if o.XcvrPartNumber != nil {
        structMap["xcvr_part_number"] = o.XcvrPartNumber
    }
    if o.XcvrSerial != nil {
        structMap["xcvr_serial"] = o.XcvrSerial
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OptionalStatsPort.
// It customizes the JSON unmarshaling process for OptionalStatsPort objects.
func (o *OptionalStatsPort) UnmarshalJSON(input []byte) error {
    var temp tempOptionalStatsPort
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "active", "auth_state", "for_site", "full_duplex", "jitter", "latency", "loss", "lte_iccid", "lte_imei", "lte_imsi", "mac_count", "mac_limit", "neighbor_mac", "neighbor_port_desc", "neighbor_system_name", "poe_disabled", "poe_mode", "poe_on", "port_id", "port_mac", "port_usage", "power_draw", "rx_bcast_pkts", "rx_bps", "rx_bytes", "rx_errors", "rx_mcast_pkts", "rx_pkts", "speed", "stp_role", "stp_state", "tx_bcast_pkts", "tx_bps", "tx_bytes", "tx_errors", "tx_mcast_pkts", "tx_pkts", "type", "unconfigured", "up", "xcvr_model", "xcvr_part_number", "xcvr_serial")
    if err != nil {
    	return err
    }
    o.AdditionalProperties = additionalProperties
    
    o.Active = temp.Active
    o.AuthState = temp.AuthState
    o.ForSite = temp.ForSite
    o.FullDuplex = temp.FullDuplex
    o.Jitter = temp.Jitter
    o.Latency = temp.Latency
    o.Loss = temp.Loss
    o.LteIccid = temp.LteIccid
    o.LteImei = temp.LteImei
    o.LteImsi = temp.LteImsi
    o.MacCount = temp.MacCount
    o.MacLimit = temp.MacLimit
    o.NeighborMac = *temp.NeighborMac
    o.NeighborPortDesc = temp.NeighborPortDesc
    o.NeighborSystemName = temp.NeighborSystemName
    o.PoeDisabled = temp.PoeDisabled
    o.PoeMode = temp.PoeMode
    o.PoeOn = temp.PoeOn
    o.PortId = *temp.PortId
    o.PortMac = *temp.PortMac
    o.PortUsage = temp.PortUsage
    o.PowerDraw = temp.PowerDraw
    o.RxBcastPkts = temp.RxBcastPkts
    o.RxBps = temp.RxBps
    o.RxBytes = *temp.RxBytes
    o.RxErrors = temp.RxErrors
    o.RxMcastPkts = temp.RxMcastPkts
    o.RxPkts = *temp.RxPkts
    o.Speed = temp.Speed
    o.StpRole = temp.StpRole
    o.StpState = temp.StpState
    o.TxBcastPkts = temp.TxBcastPkts
    o.TxBps = temp.TxBps
    o.TxBytes = *temp.TxBytes
    o.TxErrors = temp.TxErrors
    o.TxMcastPkts = temp.TxMcastPkts
    o.TxPkts = *temp.TxPkts
    o.Type = temp.Type
    o.Unconfigured = temp.Unconfigured
    o.Up = temp.Up
    o.XcvrModel = temp.XcvrModel
    o.XcvrPartNumber = temp.XcvrPartNumber
    o.XcvrSerial = temp.XcvrSerial
    return nil
}

// tempOptionalStatsPort is a temporary struct used for validating the fields of OptionalStatsPort.
type tempOptionalStatsPort  struct {
    Active             *bool                         `json:"active,omitempty"`
    AuthState          *StatsSwitchPortAuthStateEnum `json:"auth_state,omitempty"`
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
    PoeMode            *StatsSwitchPortPoeModeEnum   `json:"poe_mode,omitempty"`
    PoeOn              *bool                         `json:"poe_on,omitempty"`
    PortId             *string                       `json:"port_id"`
    PortMac            *string                       `json:"port_mac"`
    PortUsage          *StatsSwitchPortPortUsageEnum `json:"port_usage,omitempty"`
    PowerDraw          *float64                      `json:"power_draw,omitempty"`
    RxBcastPkts        *int                          `json:"rx_bcast_pkts,omitempty"`
    RxBps              *int                          `json:"rx_bps,omitempty"`
    RxBytes            *int64                        `json:"rx_bytes"`
    RxErrors           *int                          `json:"rx_errors,omitempty"`
    RxMcastPkts        *int                          `json:"rx_mcast_pkts,omitempty"`
    RxPkts             *int                          `json:"rx_pkts"`
    Speed              *int                          `json:"speed,omitempty"`
    StpRole            *StatsSwitchPortStpRoleEnum   `json:"stp_role,omitempty"`
    StpState           *StatsSwitchPortStpStateEnum  `json:"stp_state,omitempty"`
    TxBcastPkts        *int                          `json:"tx_bcast_pkts,omitempty"`
    TxBps              *int                          `json:"tx_bps,omitempty"`
    TxBytes            *int64                        `json:"tx_bytes"`
    TxErrors           *int                          `json:"tx_errors,omitempty"`
    TxMcastPkts        *int                          `json:"tx_mcast_pkts,omitempty"`
    TxPkts             *int                          `json:"tx_pkts"`
    Type               *StatsSwitchPortTypeEnum      `json:"type,omitempty"`
    Unconfigured       *bool                         `json:"unconfigured,omitempty"`
    Up                 *bool                         `json:"up,omitempty"`
    XcvrModel          *string                       `json:"xcvr_model,omitempty"`
    XcvrPartNumber     *string                       `json:"xcvr_part_number,omitempty"`
    XcvrSerial         *string                       `json:"xcvr_serial,omitempty"`
}

func (o *tempOptionalStatsPort) validate() error {
    var errs []string
    if o.NeighborMac == nil {
        errs = append(errs, "required field `neighbor_mac` is missing for type `optional_stats_port`")
    }
    if o.PortId == nil {
        errs = append(errs, "required field `port_id` is missing for type `optional_stats_port`")
    }
    if o.PortMac == nil {
        errs = append(errs, "required field `port_mac` is missing for type `optional_stats_port`")
    }
    if o.RxBytes == nil {
        errs = append(errs, "required field `rx_bytes` is missing for type `optional_stats_port`")
    }
    if o.RxPkts == nil {
        errs = append(errs, "required field `rx_pkts` is missing for type `optional_stats_port`")
    }
    if o.TxBytes == nil {
        errs = append(errs, "required field `tx_bytes` is missing for type `optional_stats_port`")
    }
    if o.TxPkts == nil {
        errs = append(errs, "required field `tx_pkts` is missing for type `optional_stats_port`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
