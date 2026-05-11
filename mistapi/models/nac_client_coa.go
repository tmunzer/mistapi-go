// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// NacClientCoa represents a NacClientCoa struct.
type NacClientCoa struct {
	// CoA type to send. enum: `reauth`, `disconnect`
	CoaType              *NacCoaTypeEnum        `json:"coa_type,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for NacClientCoa,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (n NacClientCoa) String() string {
	return fmt.Sprintf(
		"NacClientCoa[CoaType=%v, AdditionalProperties=%v]",
		n.CoaType, n.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for NacClientCoa.
// It customizes the JSON marshaling process for NacClientCoa objects.
func (n NacClientCoa) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(n.AdditionalProperties,
		"coa_type"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(n.toMap())
}

// toMap converts the NacClientCoa object to a map representation for JSON marshaling.
func (n NacClientCoa) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, n.AdditionalProperties)
	if n.CoaType != nil {
		structMap["coa_type"] = n.CoaType
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for NacClientCoa.
// It customizes the JSON unmarshaling process for NacClientCoa objects.
func (n *NacClientCoa) UnmarshalJSON(input []byte) error {
	var temp tempNacClientCoa
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "coa_type")
	if err != nil {
		return err
	}
	n.AdditionalProperties = additionalProperties

	n.CoaType = temp.CoaType
	return nil
}

// tempNacClientCoa is a temporary struct used for validating the fields of NacClientCoa.
type tempNacClientCoa struct {
	CoaType *NacCoaTypeEnum `json:"coa_type,omitempty"`
}
