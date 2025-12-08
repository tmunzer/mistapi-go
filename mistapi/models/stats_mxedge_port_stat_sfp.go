// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// StatsMxedgePortStatSfp represents a StatsMxedgePortStatSfp struct.
type StatsMxedgePortStatSfp struct {
	Codes                *string                `json:"codes,omitempty"`
	Mbps                 *int                   `json:"mbps,omitempty"`
	PartNo               *string                `json:"part_no,omitempty"`
	SerialNo             *string                `json:"serial_no,omitempty"`
	Type                 *int                   `json:"type,omitempty"`
	Vendor               *string                `json:"vendor,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for StatsMxedgePortStatSfp,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s StatsMxedgePortStatSfp) String() string {
	return fmt.Sprintf(
		"StatsMxedgePortStatSfp[Codes=%v, Mbps=%v, PartNo=%v, SerialNo=%v, Type=%v, Vendor=%v, AdditionalProperties=%v]",
		s.Codes, s.Mbps, s.PartNo, s.SerialNo, s.Type, s.Vendor, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for StatsMxedgePortStatSfp.
// It customizes the JSON marshaling process for StatsMxedgePortStatSfp objects.
func (s StatsMxedgePortStatSfp) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"codes", "mbps", "part_no", "serial_no", "type", "vendor"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the StatsMxedgePortStatSfp object to a map representation for JSON marshaling.
func (s StatsMxedgePortStatSfp) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.Codes != nil {
		structMap["codes"] = s.Codes
	}
	if s.Mbps != nil {
		structMap["mbps"] = s.Mbps
	}
	if s.PartNo != nil {
		structMap["part_no"] = s.PartNo
	}
	if s.SerialNo != nil {
		structMap["serial_no"] = s.SerialNo
	}
	if s.Type != nil {
		structMap["type"] = s.Type
	}
	if s.Vendor != nil {
		structMap["vendor"] = s.Vendor
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsMxedgePortStatSfp.
// It customizes the JSON unmarshaling process for StatsMxedgePortStatSfp objects.
func (s *StatsMxedgePortStatSfp) UnmarshalJSON(input []byte) error {
	var temp tempStatsMxedgePortStatSfp
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "codes", "mbps", "part_no", "serial_no", "type", "vendor")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.Codes = temp.Codes
	s.Mbps = temp.Mbps
	s.PartNo = temp.PartNo
	s.SerialNo = temp.SerialNo
	s.Type = temp.Type
	s.Vendor = temp.Vendor
	return nil
}

// tempStatsMxedgePortStatSfp is a temporary struct used for validating the fields of StatsMxedgePortStatSfp.
type tempStatsMxedgePortStatSfp struct {
	Codes    *string `json:"codes,omitempty"`
	Mbps     *int    `json:"mbps,omitempty"`
	PartNo   *string `json:"part_no,omitempty"`
	SerialNo *string `json:"serial_no,omitempty"`
	Type     *int    `json:"type,omitempty"`
	Vendor   *string `json:"vendor,omitempty"`
}
