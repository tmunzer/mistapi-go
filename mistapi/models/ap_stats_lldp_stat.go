package models

import (
    "encoding/json"
)

// ApStatsLldpStat represents a ApStatsLldpStat struct.
// LLDP Stat (neighbor information, power negotiations)
type ApStatsLldpStat struct {
    ChassisId            Optional[string]  `json:"chassis_id"`
    // whether it support LLDP-MED
    LldpMedSupported     Optional[bool]    `json:"lldp_med_supported"`
    // switch’s management address (if advertised), can be IPv4, IPv6, or MAC
    MgmtAddr             Optional[string]  `json:"mgmt_addr"`
    MgmtAddrs            []string          `json:"mgmt_addrs,omitempty"`
    // ge-0/0/4
    PortDesc             Optional[string]  `json:"port_desc"`
    PortId               Optional[string]  `json:"port_id"`
    // in mW, provided/allocated by PSE
    PowerAllocated       Optional[float64] `json:"power_allocated"`
    // in mW, total power needed by PD
    PowerDraw            Optional[float64] `json:"power_draw"`
    // number of negotiations, if it keeps increasing, we don’t have a stable power
    PowerRequestCount    Optional[int]     `json:"power_request_count"`
    // in mW, the current power requested by PD
    PowerRequested       Optional[float64] `json:"power_requested"`
    // description provided by switch
    SystemDesc           Optional[string]  `json:"system_desc"`
    // name of the switch
    SystemName           Optional[string]  `json:"system_name"`
    AdditionalProperties map[string]any    `json:"_"`
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
    if a.ChassisId.IsValueSet() {
        if a.ChassisId.Value() != nil {
            structMap["chassis_id"] = a.ChassisId.Value()
        } else {
            structMap["chassis_id"] = nil
        }
    }
    if a.LldpMedSupported.IsValueSet() {
        if a.LldpMedSupported.Value() != nil {
            structMap["lldp_med_supported"] = a.LldpMedSupported.Value()
        } else {
            structMap["lldp_med_supported"] = nil
        }
    }
    if a.MgmtAddr.IsValueSet() {
        if a.MgmtAddr.Value() != nil {
            structMap["mgmt_addr"] = a.MgmtAddr.Value()
        } else {
            structMap["mgmt_addr"] = nil
        }
    }
    if a.MgmtAddrs != nil {
        structMap["mgmt_addrs"] = a.MgmtAddrs
    }
    if a.PortDesc.IsValueSet() {
        if a.PortDesc.Value() != nil {
            structMap["port_desc"] = a.PortDesc.Value()
        } else {
            structMap["port_desc"] = nil
        }
    }
    if a.PortId.IsValueSet() {
        if a.PortId.Value() != nil {
            structMap["port_id"] = a.PortId.Value()
        } else {
            structMap["port_id"] = nil
        }
    }
    if a.PowerAllocated.IsValueSet() {
        if a.PowerAllocated.Value() != nil {
            structMap["power_allocated"] = a.PowerAllocated.Value()
        } else {
            structMap["power_allocated"] = nil
        }
    }
    if a.PowerDraw.IsValueSet() {
        if a.PowerDraw.Value() != nil {
            structMap["power_draw"] = a.PowerDraw.Value()
        } else {
            structMap["power_draw"] = nil
        }
    }
    if a.PowerRequestCount.IsValueSet() {
        if a.PowerRequestCount.Value() != nil {
            structMap["power_request_count"] = a.PowerRequestCount.Value()
        } else {
            structMap["power_request_count"] = nil
        }
    }
    if a.PowerRequested.IsValueSet() {
        if a.PowerRequested.Value() != nil {
            structMap["power_requested"] = a.PowerRequested.Value()
        } else {
            structMap["power_requested"] = nil
        }
    }
    if a.SystemDesc.IsValueSet() {
        if a.SystemDesc.Value() != nil {
            structMap["system_desc"] = a.SystemDesc.Value()
        } else {
            structMap["system_desc"] = nil
        }
    }
    if a.SystemName.IsValueSet() {
        if a.SystemName.Value() != nil {
            structMap["system_name"] = a.SystemName.Value()
        } else {
            structMap["system_name"] = nil
        }
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
    additionalProperties, err := UnmarshalAdditionalProperties(input, "chassis_id", "lldp_med_supported", "mgmt_addr", "mgmt_addrs", "port_desc", "port_id", "power_allocated", "power_draw", "power_request_count", "power_requested", "system_desc", "system_name")
    if err != nil {
    	return err
    }
    
    a.AdditionalProperties = additionalProperties
    a.ChassisId = temp.ChassisId
    a.LldpMedSupported = temp.LldpMedSupported
    a.MgmtAddr = temp.MgmtAddr
    a.MgmtAddrs = temp.MgmtAddrs
    a.PortDesc = temp.PortDesc
    a.PortId = temp.PortId
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
    ChassisId         Optional[string]  `json:"chassis_id"`
    LldpMedSupported  Optional[bool]    `json:"lldp_med_supported"`
    MgmtAddr          Optional[string]  `json:"mgmt_addr"`
    MgmtAddrs         []string          `json:"mgmt_addrs,omitempty"`
    PortDesc          Optional[string]  `json:"port_desc"`
    PortId            Optional[string]  `json:"port_id"`
    PowerAllocated    Optional[float64] `json:"power_allocated"`
    PowerDraw         Optional[float64] `json:"power_draw"`
    PowerRequestCount Optional[int]     `json:"power_request_count"`
    PowerRequested    Optional[float64] `json:"power_requested"`
    SystemDesc        Optional[string]  `json:"system_desc"`
    SystemName        Optional[string]  `json:"system_name"`
}
