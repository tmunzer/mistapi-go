// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// StatsApLldpStat represents a StatsApLldpStat struct.
// LLDP neighbor information and power negotiations. For backward compatibility, when multiple neighbors exist, only information from the first neighbor is displayed.
type StatsApLldpStat struct {
	ChassisId Optional[string] `json:"chassis_id"`
	// Whether it support LLDP-MED
	LldpMedSupported Optional[bool] `json:"lldp_med_supported"`
	// Management IP address of the switch
	MgmtAddr Optional[string] `json:"mgmt_addr"`
	// List of management IP addresses (IPv4 and IPv6)
	MgmtAddrs []string `json:"mgmt_addrs,omitempty"`
	// Port description, e.g. “2/20”, “Port 2 on Switch0”
	PortDesc Optional[string] `json:"port_desc"`
	// Port identifier
	PortId Optional[string] `json:"port_id"`
	// In mW, power allocated by PSE
	PowerAllocated Optional[float64] `json:"power_allocated"`
	// In mW, total Power Avail at AP from pwr source
	PowerAvail *int `json:"power_avail,omitempty"`
	// In mW, surplus if positive or deficit if negative
	PowerBudget *int `json:"power_budget,omitempty"`
	// Whether power is insufficient
	PowerConstrained *bool `json:"power_constrained,omitempty"`
	// In mW, total power needed by PD
	PowerDraw Optional[float64] `json:"power_draw"`
	// In mW, total Power needed incl Peripherals
	PowerNeeded *int `json:"power_needed,omitempty"`
	// Constrained mode
	PowerOpmode *string `json:"power_opmode,omitempty"`
	// Number of negotiations, if it keeps increasing, we don’ t have a stable power
	PowerRequestCount Optional[int] `json:"power_request_count"`
	// In mW, power requested by PD
	PowerRequested Optional[float64] `json:"power_requested"`
	// Single power source (DC Input / PoE 802.3at / PoE 802.3af / PoE 802.3bt / MULTI-PD / LLDP / ? (unknown)).
	PowerSrc *string `json:"power_src,omitempty"`
	// List of management IP addresses (IPv4 and IPv6)
	PowerSrcs []string `json:"power_srcs,omitempty"`
	// Description provided by switch
	SystemDesc Optional[string] `json:"system_desc"`
	// Name of the switch
	SystemName           Optional[string]       `json:"system_name"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for StatsApLldpStat,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s StatsApLldpStat) String() string {
	return fmt.Sprintf(
		"StatsApLldpStat[ChassisId=%v, LldpMedSupported=%v, MgmtAddr=%v, MgmtAddrs=%v, PortDesc=%v, PortId=%v, PowerAllocated=%v, PowerAvail=%v, PowerBudget=%v, PowerConstrained=%v, PowerDraw=%v, PowerNeeded=%v, PowerOpmode=%v, PowerRequestCount=%v, PowerRequested=%v, PowerSrc=%v, PowerSrcs=%v, SystemDesc=%v, SystemName=%v, AdditionalProperties=%v]",
		s.ChassisId, s.LldpMedSupported, s.MgmtAddr, s.MgmtAddrs, s.PortDesc, s.PortId, s.PowerAllocated, s.PowerAvail, s.PowerBudget, s.PowerConstrained, s.PowerDraw, s.PowerNeeded, s.PowerOpmode, s.PowerRequestCount, s.PowerRequested, s.PowerSrc, s.PowerSrcs, s.SystemDesc, s.SystemName, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for StatsApLldpStat.
// It customizes the JSON marshaling process for StatsApLldpStat objects.
func (s StatsApLldpStat) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"chassis_id", "lldp_med_supported", "mgmt_addr", "mgmt_addrs", "port_desc", "port_id", "power_allocated", "power_avail", "power_budget", "power_constrained", "power_draw", "power_needed", "power_opmode", "power_request_count", "power_requested", "power_src", "power_srcs", "system_desc", "system_name"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the StatsApLldpStat object to a map representation for JSON marshaling.
func (s StatsApLldpStat) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
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
	if s.PowerAvail != nil {
		structMap["power_avail"] = s.PowerAvail
	}
	if s.PowerBudget != nil {
		structMap["power_budget"] = s.PowerBudget
	}
	if s.PowerConstrained != nil {
		structMap["power_constrained"] = s.PowerConstrained
	}
	if s.PowerDraw.IsValueSet() {
		if s.PowerDraw.Value() != nil {
			structMap["power_draw"] = s.PowerDraw.Value()
		} else {
			structMap["power_draw"] = nil
		}
	}
	if s.PowerNeeded != nil {
		structMap["power_needed"] = s.PowerNeeded
	}
	if s.PowerOpmode != nil {
		structMap["power_opmode"] = s.PowerOpmode
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
	if s.PowerSrc != nil {
		structMap["power_src"] = s.PowerSrc
	}
	if s.PowerSrcs != nil {
		structMap["power_srcs"] = s.PowerSrcs
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
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "chassis_id", "lldp_med_supported", "mgmt_addr", "mgmt_addrs", "port_desc", "port_id", "power_allocated", "power_avail", "power_budget", "power_constrained", "power_draw", "power_needed", "power_opmode", "power_request_count", "power_requested", "power_src", "power_srcs", "system_desc", "system_name")
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
	s.PowerAvail = temp.PowerAvail
	s.PowerBudget = temp.PowerBudget
	s.PowerConstrained = temp.PowerConstrained
	s.PowerDraw = temp.PowerDraw
	s.PowerNeeded = temp.PowerNeeded
	s.PowerOpmode = temp.PowerOpmode
	s.PowerRequestCount = temp.PowerRequestCount
	s.PowerRequested = temp.PowerRequested
	s.PowerSrc = temp.PowerSrc
	s.PowerSrcs = temp.PowerSrcs
	s.SystemDesc = temp.SystemDesc
	s.SystemName = temp.SystemName
	return nil
}

// tempStatsApLldpStat is a temporary struct used for validating the fields of StatsApLldpStat.
type tempStatsApLldpStat struct {
	ChassisId         Optional[string]  `json:"chassis_id"`
	LldpMedSupported  Optional[bool]    `json:"lldp_med_supported"`
	MgmtAddr          Optional[string]  `json:"mgmt_addr"`
	MgmtAddrs         []string          `json:"mgmt_addrs,omitempty"`
	PortDesc          Optional[string]  `json:"port_desc"`
	PortId            Optional[string]  `json:"port_id"`
	PowerAllocated    Optional[float64] `json:"power_allocated"`
	PowerAvail        *int              `json:"power_avail,omitempty"`
	PowerBudget       *int              `json:"power_budget,omitempty"`
	PowerConstrained  *bool             `json:"power_constrained,omitempty"`
	PowerDraw         Optional[float64] `json:"power_draw"`
	PowerNeeded       *int              `json:"power_needed,omitempty"`
	PowerOpmode       *string           `json:"power_opmode,omitempty"`
	PowerRequestCount Optional[int]     `json:"power_request_count"`
	PowerRequested    Optional[float64] `json:"power_requested"`
	PowerSrc          *string           `json:"power_src,omitempty"`
	PowerSrcs         []string          `json:"power_srcs,omitempty"`
	SystemDesc        Optional[string]  `json:"system_desc"`
	SystemName        Optional[string]  `json:"system_name"`
}
