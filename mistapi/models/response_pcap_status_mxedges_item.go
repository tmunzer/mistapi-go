// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// ResponsePcapStatusMxedgesItem represents a ResponsePcapStatusMxedgesItem struct.
type ResponsePcapStatusMxedgesItem struct {
	// Dict of interfaces to capture on, property key is the port name
	Interfaces           map[string]CaptureMxedgeMxedgesInterfaces `json:"interfaces,omitempty"`
	AdditionalProperties map[string]interface{}                    `json:"_"`
}

// String implements the fmt.Stringer interface for ResponsePcapStatusMxedgesItem,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponsePcapStatusMxedgesItem) String() string {
	return fmt.Sprintf(
		"ResponsePcapStatusMxedgesItem[Interfaces=%v, AdditionalProperties=%v]",
		r.Interfaces, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponsePcapStatusMxedgesItem.
// It customizes the JSON marshaling process for ResponsePcapStatusMxedgesItem objects.
func (r ResponsePcapStatusMxedgesItem) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(r.AdditionalProperties,
		"interfaces"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(r.toMap())
}

// toMap converts the ResponsePcapStatusMxedgesItem object to a map representation for JSON marshaling.
func (r ResponsePcapStatusMxedgesItem) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, r.AdditionalProperties)
	if r.Interfaces != nil {
		structMap["interfaces"] = r.Interfaces
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponsePcapStatusMxedgesItem.
// It customizes the JSON unmarshaling process for ResponsePcapStatusMxedgesItem objects.
func (r *ResponsePcapStatusMxedgesItem) UnmarshalJSON(input []byte) error {
	var temp tempResponsePcapStatusMxedgesItem
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "interfaces")
	if err != nil {
		return err
	}
	r.AdditionalProperties = additionalProperties

	r.Interfaces = temp.Interfaces
	return nil
}

// tempResponsePcapStatusMxedgesItem is a temporary struct used for validating the fields of ResponsePcapStatusMxedgesItem.
type tempResponsePcapStatusMxedgesItem struct {
	Interfaces map[string]CaptureMxedgeMxedgesInterfaces `json:"interfaces,omitempty"`
}
