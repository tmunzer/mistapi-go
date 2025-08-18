// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// OspfArea represents a OspfArea struct.
// Property key is the OSPF Area (Area should be a number (0-255) / IP address)
type OspfArea struct {
	IncludeLoopback *bool                       `json:"include_loopback,omitempty"`
	Networks        map[string]OspfAreasNetwork `json:"networks,omitempty"`
	// OSPF type. enum: `default`, `nssa`, `stub`
	Type                 *OspfAreaTypeEnum      `json:"type,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for OspfArea,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (o OspfArea) String() string {
	return fmt.Sprintf(
		"OspfArea[IncludeLoopback=%v, Networks=%v, Type=%v, AdditionalProperties=%v]",
		o.IncludeLoopback, o.Networks, o.Type, o.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for OspfArea.
// It customizes the JSON marshaling process for OspfArea objects.
func (o OspfArea) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(o.AdditionalProperties,
		"include_loopback", "networks", "type"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(o.toMap())
}

// toMap converts the OspfArea object to a map representation for JSON marshaling.
func (o OspfArea) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, o.AdditionalProperties)
	if o.IncludeLoopback != nil {
		structMap["include_loopback"] = o.IncludeLoopback
	}
	if o.Networks != nil {
		structMap["networks"] = o.Networks
	}
	if o.Type != nil {
		structMap["type"] = o.Type
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OspfArea.
// It customizes the JSON unmarshaling process for OspfArea objects.
func (o *OspfArea) UnmarshalJSON(input []byte) error {
	var temp tempOspfArea
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "include_loopback", "networks", "type")
	if err != nil {
		return err
	}
	o.AdditionalProperties = additionalProperties

	o.IncludeLoopback = temp.IncludeLoopback
	o.Networks = temp.Networks
	o.Type = temp.Type
	return nil
}

// tempOspfArea is a temporary struct used for validating the fields of OspfArea.
type tempOspfArea struct {
	IncludeLoopback *bool                       `json:"include_loopback,omitempty"`
	Networks        map[string]OspfAreasNetwork `json:"networks,omitempty"`
	Type            *OspfAreaTypeEnum           `json:"type,omitempty"`
}
