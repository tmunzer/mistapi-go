// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// NacClientCoaResponse represents a NacClientCoaResponse struct.
// Response returned after sending a NAC client CoA command
type NacClientCoaResponse struct {
	// Target AP or switch MAC address for the CoA command
	DeviceMac *string `json:"device_mac,omitempty"`
	// enum: `ap`, `gateway`, `switch`
	DeviceType           *DeviceTypeEnum        `json:"device_type,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for NacClientCoaResponse,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (n NacClientCoaResponse) String() string {
	return fmt.Sprintf(
		"NacClientCoaResponse[DeviceMac=%v, DeviceType=%v, AdditionalProperties=%v]",
		n.DeviceMac, n.DeviceType, n.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for NacClientCoaResponse.
// It customizes the JSON marshaling process for NacClientCoaResponse objects.
func (n NacClientCoaResponse) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(n.AdditionalProperties,
		"device_mac", "device_type"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(n.toMap())
}

// toMap converts the NacClientCoaResponse object to a map representation for JSON marshaling.
func (n NacClientCoaResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, n.AdditionalProperties)
	if n.DeviceMac != nil {
		structMap["device_mac"] = n.DeviceMac
	}
	if n.DeviceType != nil {
		structMap["device_type"] = n.DeviceType
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for NacClientCoaResponse.
// It customizes the JSON unmarshaling process for NacClientCoaResponse objects.
func (n *NacClientCoaResponse) UnmarshalJSON(input []byte) error {
	var temp tempNacClientCoaResponse
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "device_mac", "device_type")
	if err != nil {
		return err
	}
	n.AdditionalProperties = additionalProperties

	n.DeviceMac = temp.DeviceMac
	n.DeviceType = temp.DeviceType
	return nil
}

// tempNacClientCoaResponse is a temporary struct used for validating the fields of NacClientCoaResponse.
type tempNacClientCoaResponse struct {
	DeviceMac  *string         `json:"device_mac,omitempty"`
	DeviceType *DeviceTypeEnum `json:"device_type,omitempty"`
}
