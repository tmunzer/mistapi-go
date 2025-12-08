// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// StatsMxedgePortStatLldpStats represents a StatsMxedgePortStatLldpStats struct.
type StatsMxedgePortStatLldpStats struct {
	ChassisId            *string                `json:"chassis_id,omitempty"`
	MgmtAddr             *string                `json:"mgmt_addr,omitempty"`
	PortDesc             *string                `json:"port_desc,omitempty"`
	PortId               *string                `json:"port_id,omitempty"`
	SystemDesc           *string                `json:"system_desc,omitempty"`
	SystemName           *string                `json:"system_name,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for StatsMxedgePortStatLldpStats,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s StatsMxedgePortStatLldpStats) String() string {
	return fmt.Sprintf(
		"StatsMxedgePortStatLldpStats[ChassisId=%v, MgmtAddr=%v, PortDesc=%v, PortId=%v, SystemDesc=%v, SystemName=%v, AdditionalProperties=%v]",
		s.ChassisId, s.MgmtAddr, s.PortDesc, s.PortId, s.SystemDesc, s.SystemName, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for StatsMxedgePortStatLldpStats.
// It customizes the JSON marshaling process for StatsMxedgePortStatLldpStats objects.
func (s StatsMxedgePortStatLldpStats) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"chassis_id", "mgmt_addr", "port_desc", "port_id", "system_desc", "system_name"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the StatsMxedgePortStatLldpStats object to a map representation for JSON marshaling.
func (s StatsMxedgePortStatLldpStats) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.ChassisId != nil {
		structMap["chassis_id"] = s.ChassisId
	}
	if s.MgmtAddr != nil {
		structMap["mgmt_addr"] = s.MgmtAddr
	}
	if s.PortDesc != nil {
		structMap["port_desc"] = s.PortDesc
	}
	if s.PortId != nil {
		structMap["port_id"] = s.PortId
	}
	if s.SystemDesc != nil {
		structMap["system_desc"] = s.SystemDesc
	}
	if s.SystemName != nil {
		structMap["system_name"] = s.SystemName
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsMxedgePortStatLldpStats.
// It customizes the JSON unmarshaling process for StatsMxedgePortStatLldpStats objects.
func (s *StatsMxedgePortStatLldpStats) UnmarshalJSON(input []byte) error {
	var temp tempStatsMxedgePortStatLldpStats
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "chassis_id", "mgmt_addr", "port_desc", "port_id", "system_desc", "system_name")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.ChassisId = temp.ChassisId
	s.MgmtAddr = temp.MgmtAddr
	s.PortDesc = temp.PortDesc
	s.PortId = temp.PortId
	s.SystemDesc = temp.SystemDesc
	s.SystemName = temp.SystemName
	return nil
}

// tempStatsMxedgePortStatLldpStats is a temporary struct used for validating the fields of StatsMxedgePortStatLldpStats.
type tempStatsMxedgePortStatLldpStats struct {
	ChassisId  *string `json:"chassis_id,omitempty"`
	MgmtAddr   *string `json:"mgmt_addr,omitempty"`
	PortDesc   *string `json:"port_desc,omitempty"`
	PortId     *string `json:"port_id,omitempty"`
	SystemDesc *string `json:"system_desc,omitempty"`
	SystemName *string `json:"system_name,omitempty"`
}
