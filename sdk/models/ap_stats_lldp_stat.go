package models

import (
    "encoding/json"
)

// ApStatsLldpStat represents a ApStatsLldpStat struct.
// LLDP Stat (neighbor information, power negotiations)
type ApStatsLldpStat struct {
    ChassisId            *string        `json:"chassis_id,omitempty"`
    // whether it support LLDP-MED
    LldpMedSupported     *bool          `json:"lldp_med_supported,omitempty"`
    // switch’s management address (if advertised), can be IPv4, IPv6, or MAC
    MgmtAddr             *string        `json:"mgmt_addr,omitempty"`
    // port description, e.g. “2/20”, “Port 2 on Switch0”
    PortDesc             *string        `json:"port_desc,omitempty"`
    // in mW, provided/allocated by PSE
    PowerAllocated       *float64       `json:"power_allocated,omitempty"`
    // in mW, total power needed by PD
    PowerDraw            *float64       `json:"power_draw,omitempty"`
    // number of negotiations, if it keeps increasing, we don’t have a stable power
    PowerRequestCount    *int           `json:"power_request_count,omitempty"`
    // in mW, the current power requested by PD
    PowerRequested       *float64       `json:"power_requested,omitempty"`
    // description provided by switch, e.g. “HP J9729A 2920-48G-POE+ Switch”
    SystemDesc           *string        `json:"system_desc,omitempty"`
    // name of the switch, e.g. “TC2-OWL-Stack-01”
    SystemName           *string        `json:"system_name,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ApStatsLldpStat.
// It customizes the JSON marshaling process for ApStatsLldpStat objects.
func (a ApStatsLldpStat) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the ApStatsLldpStat object to a map representation for JSON marshaling.
func (a ApStatsLldpStat) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, a.AdditionalProperties)
    if a.ChassisId != nil {
        structMap["chassis_id"] = a.ChassisId
    }
    if a.LldpMedSupported != nil {
        structMap["lldp_med_supported"] = a.LldpMedSupported
    }
    if a.MgmtAddr != nil {
        structMap["mgmt_addr"] = a.MgmtAddr
    }
    if a.PortDesc != nil {
        structMap["port_desc"] = a.PortDesc
    }
    if a.PowerAllocated != nil {
        structMap["power_allocated"] = a.PowerAllocated
    }
    if a.PowerDraw != nil {
        structMap["power_draw"] = a.PowerDraw
    }
    if a.PowerRequestCount != nil {
        structMap["power_request_count"] = a.PowerRequestCount
    }
    if a.PowerRequested != nil {
        structMap["power_requested"] = a.PowerRequested
    }
    if a.SystemDesc != nil {
        structMap["system_desc"] = a.SystemDesc
    }
    if a.SystemName != nil {
        structMap["system_name"] = a.SystemName
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ApStatsLldpStat.
// It customizes the JSON unmarshaling process for ApStatsLldpStat objects.
func (a *ApStatsLldpStat) UnmarshalJSON(input []byte) error {
    var temp apStatsLldpStat
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "chassis_id", "lldp_med_supported", "mgmt_addr", "port_desc", "power_allocated", "power_draw", "power_request_count", "power_requested", "system_desc", "system_name")
    if err != nil {
    	return err
    }
    
    a.AdditionalProperties = additionalProperties
    a.ChassisId = temp.ChassisId
    a.LldpMedSupported = temp.LldpMedSupported
    a.MgmtAddr = temp.MgmtAddr
    a.PortDesc = temp.PortDesc
    a.PowerAllocated = temp.PowerAllocated
    a.PowerDraw = temp.PowerDraw
    a.PowerRequestCount = temp.PowerRequestCount
    a.PowerRequested = temp.PowerRequested
    a.SystemDesc = temp.SystemDesc
    a.SystemName = temp.SystemName
    return nil
}

// apStatsLldpStat is a temporary struct used for validating the fields of ApStatsLldpStat.
type apStatsLldpStat  struct {
    ChassisId         *string  `json:"chassis_id,omitempty"`
    LldpMedSupported  *bool    `json:"lldp_med_supported,omitempty"`
    MgmtAddr          *string  `json:"mgmt_addr,omitempty"`
    PortDesc          *string  `json:"port_desc,omitempty"`
    PowerAllocated    *float64 `json:"power_allocated,omitempty"`
    PowerDraw         *float64 `json:"power_draw,omitempty"`
    PowerRequestCount *int     `json:"power_request_count,omitempty"`
    PowerRequested    *float64 `json:"power_requested,omitempty"`
    SystemDesc        *string  `json:"system_desc,omitempty"`
    SystemName        *string  `json:"system_name,omitempty"`
}
