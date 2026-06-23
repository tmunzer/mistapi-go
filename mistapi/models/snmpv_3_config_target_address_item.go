// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// Snmpv3ConfigTargetAddressItem represents a Snmpv3ConfigTargetAddressItem struct.
// SNMPv3 notification target address entry
type Snmpv3ConfigTargetAddressItem struct {
	// IP address or hostname of the SNMP target
	Address *string `json:"address,omitempty"`
	// Mask applied to the SNMP target address
	AddressMask *string `json:"address_mask,omitempty"`
	// UDP port used by the SNMP target
	Port Optional[string] `json:"port"`
	// Set of notification tags for this target address; use spaces between multiple tags
	TagList *string `json:"tag_list,omitempty"`
	// Name of the SNMP target address entry
	TargetAddressName *string `json:"target_address_name,omitempty"`
	// Target parameter profile referenced by this target address
	TargetParameters     *string                `json:"target_parameters,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for Snmpv3ConfigTargetAddressItem,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s Snmpv3ConfigTargetAddressItem) String() string {
	return fmt.Sprintf(
		"Snmpv3ConfigTargetAddressItem[Address=%v, AddressMask=%v, Port=%v, TagList=%v, TargetAddressName=%v, TargetParameters=%v, AdditionalProperties=%v]",
		s.Address, s.AddressMask, s.Port, s.TagList, s.TargetAddressName, s.TargetParameters, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for Snmpv3ConfigTargetAddressItem.
// It customizes the JSON marshaling process for Snmpv3ConfigTargetAddressItem objects.
func (s Snmpv3ConfigTargetAddressItem) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"address", "address_mask", "port", "tag_list", "target_address_name", "target_parameters"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the Snmpv3ConfigTargetAddressItem object to a map representation for JSON marshaling.
func (s Snmpv3ConfigTargetAddressItem) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.Address != nil {
		structMap["address"] = s.Address
	}
	if s.AddressMask != nil {
		structMap["address_mask"] = s.AddressMask
	}
	if s.Port.IsValueSet() {
		if s.Port.Value() != nil {
			structMap["port"] = s.Port.Value()
		} else {
			structMap["port"] = nil
		}
	}
	if s.TagList != nil {
		structMap["tag_list"] = s.TagList
	}
	if s.TargetAddressName != nil {
		structMap["target_address_name"] = s.TargetAddressName
	}
	if s.TargetParameters != nil {
		structMap["target_parameters"] = s.TargetParameters
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Snmpv3ConfigTargetAddressItem.
// It customizes the JSON unmarshaling process for Snmpv3ConfigTargetAddressItem objects.
func (s *Snmpv3ConfigTargetAddressItem) UnmarshalJSON(input []byte) error {
	var temp tempSnmpv3ConfigTargetAddressItem
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "address", "address_mask", "port", "tag_list", "target_address_name", "target_parameters")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.Address = temp.Address
	s.AddressMask = temp.AddressMask
	s.Port = temp.Port
	s.TagList = temp.TagList
	s.TargetAddressName = temp.TargetAddressName
	s.TargetParameters = temp.TargetParameters
	return nil
}

// tempSnmpv3ConfigTargetAddressItem is a temporary struct used for validating the fields of Snmpv3ConfigTargetAddressItem.
type tempSnmpv3ConfigTargetAddressItem struct {
	Address           *string          `json:"address,omitempty"`
	AddressMask       *string          `json:"address_mask,omitempty"`
	Port              Optional[string] `json:"port"`
	TagList           *string          `json:"tag_list,omitempty"`
	TargetAddressName *string          `json:"target_address_name,omitempty"`
	TargetParameters  *string          `json:"target_parameters,omitempty"`
}
