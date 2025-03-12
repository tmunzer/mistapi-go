package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// StatsGatewayPort represents a StatsGatewayPort struct.
// Port statistics
type StatsGatewayPort struct {
    // Indicates if interface is active/inactive
    Active               *bool                         `json:"active,omitempty"`
    // if `up`==`true` and has Authenticator role. enum: `authenticated`, `authenticating`, `held`, `init`
    AuthState            *StatsSwitchPortAuthStateEnum `json:"auth_state,omitempty"`
    // Indicates if interface is disabled
    Disabled             *bool                         `json:"disabled,omitempty"`
    ForSite              *bool                         `json:"for_site,omitempty"`
    // Indicates full or half duplex
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
    // Description supplied by the system on the interface E.g. "GigabitEthernet2/0/39"
    NeighborPortDesc     *string                       `json:"neighbor_port_desc,omitempty"`
    // Name supplied by the system on the interface E.g. neighbor system name E.g. "Kumar-Acc-SW.mist.local"
    NeighborSystemName   *string                       `json:"neighbor_system_name,omitempty"`
    // Is the POE configured not be disabled.
    PoeDisabled          *bool                         `json:"poe_disabled,omitempty"`
    // enum: `802.3af`, `802.3at`, `802.3bt`
    PoeMode              *StatsSwitchPortPoeModeEnum   `json:"poe_mode,omitempty"`
    // Is the device attached to POE
    PoeOn                *bool                         `json:"poe_on,omitempty"`
    PortId               string                        `json:"port_id"`
    // Interface mac address
    PortMac              string                        `json:"port_mac"`
    // gateway port usage. enum: `lan`
    PortUsage            *StatsSwitchPortPortUsageEnum `json:"port_usage,omitempty"`
    // Amount of power being used by the interface at the time the command is executed. Unit in watts.
    PowerDraw            *float64                      `json:"power_draw,omitempty"`
    // Broadcast input packets
    RxBcastPkts          *int                          `json:"rx_bcast_pkts,omitempty"`
    // Rate of receiving traffic, bits/seconds, last known
    RxBps                Optional[int64]               `json:"rx_bps"`
    // Amount of traffic received since connection
    RxBytes              Optional[int64]               `json:"rx_bytes"`
    // Input errors
    RxErrors             *int                          `json:"rx_errors,omitempty"`
    // Multicast input packets
    RxMcastPkts          *int                          `json:"rx_mcast_pkts,omitempty"`
    // Amount of packets received since connection
    RxPkts               Optional[int64]               `json:"rx_pkts"`
    // Port speed
    Speed                *int                          `json:"speed,omitempty"`
    // if `up`==`true`. enum: `alternate`, `backup`, `designated`, `root`, `root-prevented`
    StpRole              *StatsSwitchPortStpRoleEnum   `json:"stp_role,omitempty"`
    // if `up`==`true`. enum: `blocking`, `disabled`, `forwarding`, `learning`, `listening`
    StpState             *StatsSwitchPortStpStateEnum  `json:"stp_state,omitempty"`
    // Broadcast output packets
    TxBcastPkts          *int                          `json:"tx_bcast_pkts,omitempty"`
    // Rate of transmitting traffic, bits/seconds, last known
    TxBps                Optional[int64]               `json:"tx_bps"`
    // Amount of traffic sent since connection
    TxBytes              Optional[int64]               `json:"tx_bytes"`
    // Output errors
    TxErrors             *int                          `json:"tx_errors,omitempty"`
    // Multicast output packets
    TxMcastPkts          *int                          `json:"tx_mcast_pkts,omitempty"`
    // Amount of packets sent since connection
    TxPkts               Optional[int64]               `json:"tx_pkts"`
    // device type. enum: `ap`, `ble`, `gateway`, `mxedge`, `nac`, `switch`
    Type                 *StatsSwitchPortTypeEnum      `json:"type,omitempty"`
    // Indicates if interface is unconfigured
    Unconfigured         *bool                         `json:"unconfigured,omitempty"`
    // Indicates if interface is up
    Up                   *bool                         `json:"up,omitempty"`
    // Optic Slot ModelName, Check for null/empty
    XcvrModel            *string                       `json:"xcvr_model,omitempty"`
    // Optic Slot Partnumber, Check for null/empty
    XcvrPartNumber       *string                       `json:"xcvr_part_number,omitempty"`
    // Optic Slot SerialNumber, Check for null/empty
    XcvrSerial           *string                       `json:"xcvr_serial,omitempty"`
    AdditionalProperties map[string]interface{}        `json:"_"`
}

// String implements the fmt.Stringer interface for StatsGatewayPort,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s StatsGatewayPort) String() string {
    return fmt.Sprintf(
    	"StatsGatewayPort[Active=%v, AuthState=%v, Disabled=%v, ForSite=%v, FullDuplex=%v, Jitter=%v, Latency=%v, Loss=%v, LteIccid=%v, LteImei=%v, LteImsi=%v, MacCount=%v, MacLimit=%v, NeighborMac=%v, NeighborPortDesc=%v, NeighborSystemName=%v, PoeDisabled=%v, PoeMode=%v, PoeOn=%v, PortId=%v, PortMac=%v, PortUsage=%v, PowerDraw=%v, RxBcastPkts=%v, RxBps=%v, RxBytes=%v, RxErrors=%v, RxMcastPkts=%v, RxPkts=%v, Speed=%v, StpRole=%v, StpState=%v, TxBcastPkts=%v, TxBps=%v, TxBytes=%v, TxErrors=%v, TxMcastPkts=%v, TxPkts=%v, Type=%v, Unconfigured=%v, Up=%v, XcvrModel=%v, XcvrPartNumber=%v, XcvrSerial=%v, AdditionalProperties=%v]",
    	s.Active, s.AuthState, s.Disabled, s.ForSite, s.FullDuplex, s.Jitter, s.Latency, s.Loss, s.LteIccid, s.LteImei, s.LteImsi, s.MacCount, s.MacLimit, s.NeighborMac, s.NeighborPortDesc, s.NeighborSystemName, s.PoeDisabled, s.PoeMode, s.PoeOn, s.PortId, s.PortMac, s.PortUsage, s.PowerDraw, s.RxBcastPkts, s.RxBps, s.RxBytes, s.RxErrors, s.RxMcastPkts, s.RxPkts, s.Speed, s.StpRole, s.StpState, s.TxBcastPkts, s.TxBps, s.TxBytes, s.TxErrors, s.TxMcastPkts, s.TxPkts, s.Type, s.Unconfigured, s.Up, s.XcvrModel, s.XcvrPartNumber, s.XcvrSerial, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for StatsGatewayPort.
// It customizes the JSON marshaling process for StatsGatewayPort objects.
func (s StatsGatewayPort) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "active", "auth_state", "disabled", "for_site", "full_duplex", "jitter", "latency", "loss", "lte_iccid", "lte_imei", "lte_imsi", "mac_count", "mac_limit", "neighbor_mac", "neighbor_port_desc", "neighbor_system_name", "poe_disabled", "poe_mode", "poe_on", "port_id", "port_mac", "port_usage", "power_draw", "rx_bcast_pkts", "rx_bps", "rx_bytes", "rx_errors", "rx_mcast_pkts", "rx_pkts", "speed", "stp_role", "stp_state", "tx_bcast_pkts", "tx_bps", "tx_bytes", "tx_errors", "tx_mcast_pkts", "tx_pkts", "type", "unconfigured", "up", "xcvr_model", "xcvr_part_number", "xcvr_serial"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the StatsGatewayPort object to a map representation for JSON marshaling.
func (s StatsGatewayPort) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Active != nil {
        structMap["active"] = s.Active
    }
    if s.AuthState != nil {
        structMap["auth_state"] = s.AuthState
    }
    if s.Disabled != nil {
        structMap["disabled"] = s.Disabled
    }
    if s.ForSite != nil {
        structMap["for_site"] = s.ForSite
    }
    if s.FullDuplex != nil {
        structMap["full_duplex"] = s.FullDuplex
    }
    if s.Jitter != nil {
        structMap["jitter"] = s.Jitter
    }
    if s.Latency != nil {
        structMap["latency"] = s.Latency
    }
    if s.Loss != nil {
        structMap["loss"] = s.Loss
    }
    if s.LteIccid.IsValueSet() {
        if s.LteIccid.Value() != nil {
            structMap["lte_iccid"] = s.LteIccid.Value()
        } else {
            structMap["lte_iccid"] = nil
        }
    }
    if s.LteImei.IsValueSet() {
        if s.LteImei.Value() != nil {
            structMap["lte_imei"] = s.LteImei.Value()
        } else {
            structMap["lte_imei"] = nil
        }
    }
    if s.LteImsi.IsValueSet() {
        if s.LteImsi.Value() != nil {
            structMap["lte_imsi"] = s.LteImsi.Value()
        } else {
            structMap["lte_imsi"] = nil
        }
    }
    if s.MacCount != nil {
        structMap["mac_count"] = s.MacCount
    }
    if s.MacLimit != nil {
        structMap["mac_limit"] = s.MacLimit
    }
    structMap["neighbor_mac"] = s.NeighborMac
    if s.NeighborPortDesc != nil {
        structMap["neighbor_port_desc"] = s.NeighborPortDesc
    }
    if s.NeighborSystemName != nil {
        structMap["neighbor_system_name"] = s.NeighborSystemName
    }
    if s.PoeDisabled != nil {
        structMap["poe_disabled"] = s.PoeDisabled
    }
    if s.PoeMode != nil {
        structMap["poe_mode"] = s.PoeMode
    }
    if s.PoeOn != nil {
        structMap["poe_on"] = s.PoeOn
    }
    structMap["port_id"] = s.PortId
    structMap["port_mac"] = s.PortMac
    if s.PortUsage != nil {
        structMap["port_usage"] = s.PortUsage
    }
    if s.PowerDraw != nil {
        structMap["power_draw"] = s.PowerDraw
    }
    if s.RxBcastPkts != nil {
        structMap["rx_bcast_pkts"] = s.RxBcastPkts
    }
    if s.RxBps.IsValueSet() {
        if s.RxBps.Value() != nil {
            structMap["rx_bps"] = s.RxBps.Value()
        } else {
            structMap["rx_bps"] = nil
        }
    }
    if s.RxBytes.IsValueSet() {
        if s.RxBytes.Value() != nil {
            structMap["rx_bytes"] = s.RxBytes.Value()
        } else {
            structMap["rx_bytes"] = nil
        }
    }
    if s.RxErrors != nil {
        structMap["rx_errors"] = s.RxErrors
    }
    if s.RxMcastPkts != nil {
        structMap["rx_mcast_pkts"] = s.RxMcastPkts
    }
    if s.RxPkts.IsValueSet() {
        if s.RxPkts.Value() != nil {
            structMap["rx_pkts"] = s.RxPkts.Value()
        } else {
            structMap["rx_pkts"] = nil
        }
    }
    if s.Speed != nil {
        structMap["speed"] = s.Speed
    }
    if s.StpRole != nil {
        structMap["stp_role"] = s.StpRole
    }
    if s.StpState != nil {
        structMap["stp_state"] = s.StpState
    }
    if s.TxBcastPkts != nil {
        structMap["tx_bcast_pkts"] = s.TxBcastPkts
    }
    if s.TxBps.IsValueSet() {
        if s.TxBps.Value() != nil {
            structMap["tx_bps"] = s.TxBps.Value()
        } else {
            structMap["tx_bps"] = nil
        }
    }
    if s.TxBytes.IsValueSet() {
        if s.TxBytes.Value() != nil {
            structMap["tx_bytes"] = s.TxBytes.Value()
        } else {
            structMap["tx_bytes"] = nil
        }
    }
    if s.TxErrors != nil {
        structMap["tx_errors"] = s.TxErrors
    }
    if s.TxMcastPkts != nil {
        structMap["tx_mcast_pkts"] = s.TxMcastPkts
    }
    if s.TxPkts.IsValueSet() {
        if s.TxPkts.Value() != nil {
            structMap["tx_pkts"] = s.TxPkts.Value()
        } else {
            structMap["tx_pkts"] = nil
        }
    }
    if s.Type != nil {
        structMap["type"] = s.Type
    }
    if s.Unconfigured != nil {
        structMap["unconfigured"] = s.Unconfigured
    }
    if s.Up != nil {
        structMap["up"] = s.Up
    }
    if s.XcvrModel != nil {
        structMap["xcvr_model"] = s.XcvrModel
    }
    if s.XcvrPartNumber != nil {
        structMap["xcvr_part_number"] = s.XcvrPartNumber
    }
    if s.XcvrSerial != nil {
        structMap["xcvr_serial"] = s.XcvrSerial
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsGatewayPort.
// It customizes the JSON unmarshaling process for StatsGatewayPort objects.
func (s *StatsGatewayPort) UnmarshalJSON(input []byte) error {
    var temp tempStatsGatewayPort
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "active", "auth_state", "disabled", "for_site", "full_duplex", "jitter", "latency", "loss", "lte_iccid", "lte_imei", "lte_imsi", "mac_count", "mac_limit", "neighbor_mac", "neighbor_port_desc", "neighbor_system_name", "poe_disabled", "poe_mode", "poe_on", "port_id", "port_mac", "port_usage", "power_draw", "rx_bcast_pkts", "rx_bps", "rx_bytes", "rx_errors", "rx_mcast_pkts", "rx_pkts", "speed", "stp_role", "stp_state", "tx_bcast_pkts", "tx_bps", "tx_bytes", "tx_errors", "tx_mcast_pkts", "tx_pkts", "type", "unconfigured", "up", "xcvr_model", "xcvr_part_number", "xcvr_serial")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Active = temp.Active
    s.AuthState = temp.AuthState
    s.Disabled = temp.Disabled
    s.ForSite = temp.ForSite
    s.FullDuplex = temp.FullDuplex
    s.Jitter = temp.Jitter
    s.Latency = temp.Latency
    s.Loss = temp.Loss
    s.LteIccid = temp.LteIccid
    s.LteImei = temp.LteImei
    s.LteImsi = temp.LteImsi
    s.MacCount = temp.MacCount
    s.MacLimit = temp.MacLimit
    s.NeighborMac = *temp.NeighborMac
    s.NeighborPortDesc = temp.NeighborPortDesc
    s.NeighborSystemName = temp.NeighborSystemName
    s.PoeDisabled = temp.PoeDisabled
    s.PoeMode = temp.PoeMode
    s.PoeOn = temp.PoeOn
    s.PortId = *temp.PortId
    s.PortMac = *temp.PortMac
    s.PortUsage = temp.PortUsage
    s.PowerDraw = temp.PowerDraw
    s.RxBcastPkts = temp.RxBcastPkts
    s.RxBps = temp.RxBps
    s.RxBytes = temp.RxBytes
    s.RxErrors = temp.RxErrors
    s.RxMcastPkts = temp.RxMcastPkts
    s.RxPkts = temp.RxPkts
    s.Speed = temp.Speed
    s.StpRole = temp.StpRole
    s.StpState = temp.StpState
    s.TxBcastPkts = temp.TxBcastPkts
    s.TxBps = temp.TxBps
    s.TxBytes = temp.TxBytes
    s.TxErrors = temp.TxErrors
    s.TxMcastPkts = temp.TxMcastPkts
    s.TxPkts = temp.TxPkts
    s.Type = temp.Type
    s.Unconfigured = temp.Unconfigured
    s.Up = temp.Up
    s.XcvrModel = temp.XcvrModel
    s.XcvrPartNumber = temp.XcvrPartNumber
    s.XcvrSerial = temp.XcvrSerial
    return nil
}

// tempStatsGatewayPort is a temporary struct used for validating the fields of StatsGatewayPort.
type tempStatsGatewayPort  struct {
    Active             *bool                         `json:"active,omitempty"`
    AuthState          *StatsSwitchPortAuthStateEnum `json:"auth_state,omitempty"`
    Disabled           *bool                         `json:"disabled,omitempty"`
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
    RxBps              Optional[int64]               `json:"rx_bps"`
    RxBytes            Optional[int64]               `json:"rx_bytes"`
    RxErrors           *int                          `json:"rx_errors,omitempty"`
    RxMcastPkts        *int                          `json:"rx_mcast_pkts,omitempty"`
    RxPkts             Optional[int64]               `json:"rx_pkts"`
    Speed              *int                          `json:"speed,omitempty"`
    StpRole            *StatsSwitchPortStpRoleEnum   `json:"stp_role,omitempty"`
    StpState           *StatsSwitchPortStpStateEnum  `json:"stp_state,omitempty"`
    TxBcastPkts        *int                          `json:"tx_bcast_pkts,omitempty"`
    TxBps              Optional[int64]               `json:"tx_bps"`
    TxBytes            Optional[int64]               `json:"tx_bytes"`
    TxErrors           *int                          `json:"tx_errors,omitempty"`
    TxMcastPkts        *int                          `json:"tx_mcast_pkts,omitempty"`
    TxPkts             Optional[int64]               `json:"tx_pkts"`
    Type               *StatsSwitchPortTypeEnum      `json:"type,omitempty"`
    Unconfigured       *bool                         `json:"unconfigured,omitempty"`
    Up                 *bool                         `json:"up,omitempty"`
    XcvrModel          *string                       `json:"xcvr_model,omitempty"`
    XcvrPartNumber     *string                       `json:"xcvr_part_number,omitempty"`
    XcvrSerial         *string                       `json:"xcvr_serial,omitempty"`
}

func (s *tempStatsGatewayPort) validate() error {
    var errs []string
    if s.NeighborMac == nil {
        errs = append(errs, "required field `neighbor_mac` is missing for type `stats_gateway_port`")
    }
    if s.PortId == nil {
        errs = append(errs, "required field `port_id` is missing for type `stats_gateway_port`")
    }
    if s.PortMac == nil {
        errs = append(errs, "required field `port_mac` is missing for type `stats_gateway_port`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
