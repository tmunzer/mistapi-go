package models

import (
    "encoding/json"
)

// StatsApLldpStat represents a StatsApLldpStat struct.
// LLDP Stat (neighbor information, power negotiations)
type StatsApLldpStat struct {
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
    // number of negotiations, if it keeps increasing, we don’ t have a stable power
    PowerRequestCount    Optional[int]     `json:"power_request_count"`
    // in mW, the current power requested by PD
    PowerRequested       Optional[float64] `json:"power_requested"`
    // description provided by switch
    SystemDesc           Optional[string]  `json:"system_desc"`
    // name of the switch
    SystemName           Optional[string]  `json:"system_name"`
    AdditionalProperties map[string]any    `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for StatsApLldpStat.
// It customizes the JSON marshaling process for StatsApLldpStat objects.
func (s StatsApLldpStat) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the StatsApLldpStat object to a map representation for JSON marshaling.
func (s StatsApLldpStat) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.ChassisId.IsValueSet() {
        if s.ChassisId.Value() != nil {
            structMap["chassis_id"] = s.ChassisId.Value()
        } else {
            structMap["chassis_id"] = nil
        }
    }
    if s.LldpMedSupported.IsValueSet() {
        if s.LldpMedSupported.Value() != nil {
            structMap["lldp_med_supported"] = s.LldpMedSupported.Value()
        } else {
            structMap["lldp_med_supported"] = nil
        }
    }
    if s.MgmtAddr.IsValueSet() {
        if s.MgmtAddr.Value() != nil {
            structMap["mgmt_addr"] = s.MgmtAddr.Value()
        } else {
            structMap["mgmt_addr"] = nil
        }
    }
    if s.MgmtAddrs != nil {
        structMap["mgmt_addrs"] = s.MgmtAddrs
    }
    if s.PortDesc.IsValueSet() {
        if s.PortDesc.Value() != nil {
            structMap["port_desc"] = s.PortDesc.Value()
        } else {
            structMap["port_desc"] = nil
        }
    }
    if s.PortId.IsValueSet() {
        if s.PortId.Value() != nil {
            structMap["port_id"] = s.PortId.Value()
        } else {
            structMap["port_id"] = nil
        }
    }
    if s.PowerAllocated.IsValueSet() {
        if s.PowerAllocated.Value() != nil {
            structMap["power_allocated"] = s.PowerAllocated.Value()
        } else {
            structMap["power_allocated"] = nil
        }
    }
    if s.PowerDraw.IsValueSet() {
        if s.PowerDraw.Value() != nil {
            structMap["power_draw"] = s.PowerDraw.Value()
        } else {
            structMap["power_draw"] = nil
        }
    }
    if s.PowerRequestCount.IsValueSet() {
        if s.PowerRequestCount.Value() != nil {
            structMap["power_request_count"] = s.PowerRequestCount.Value()
        } else {
            structMap["power_request_count"] = nil
        }
    }
    if s.PowerRequested.IsValueSet() {
        if s.PowerRequested.Value() != nil {
            structMap["power_requested"] = s.PowerRequested.Value()
        } else {
            structMap["power_requested"] = nil
        }
    }
    if s.SystemDesc.IsValueSet() {
        if s.SystemDesc.Value() != nil {
            structMap["system_desc"] = s.SystemDesc.Value()
        } else {
            structMap["system_desc"] = nil
        }
    }
    if s.SystemName.IsValueSet() {
        if s.SystemName.Value() != nil {
            structMap["system_name"] = s.SystemName.Value()
        } else {
            structMap["system_name"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsApLldpStat.
// It customizes the JSON unmarshaling process for StatsApLldpStat objects.
func (s *StatsApLldpStat) UnmarshalJSON(input []byte) error {
    var temp tempStatsApLldpStat
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "chassis_id", "lldp_med_supported", "mgmt_addr", "mgmt_addrs", "port_desc", "port_id", "power_allocated", "power_draw", "power_request_count", "power_requested", "system_desc", "system_name")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.ChassisId = temp.ChassisId
    s.LldpMedSupported = temp.LldpMedSupported
    s.MgmtAddr = temp.MgmtAddr
    s.MgmtAddrs = temp.MgmtAddrs
    s.PortDesc = temp.PortDesc
    s.PortId = temp.PortId
    s.PowerAllocated = temp.PowerAllocated
    s.PowerDraw = temp.PowerDraw
    s.PowerRequestCount = temp.PowerRequestCount
    s.PowerRequested = temp.PowerRequested
    s.SystemDesc = temp.SystemDesc
    s.SystemName = temp.SystemName
    return nil
}

// tempStatsApLldpStat is a temporary struct used for validating the fields of StatsApLldpStat.
type tempStatsApLldpStat  struct {
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
