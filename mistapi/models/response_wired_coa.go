// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
)

// ResponseWiredCoa represents a ResponseWiredCoa struct.
type ResponseWiredCoa struct {
	DeviceMac            *string                `json:"device_mac,omitempty"`
	PortId               *string                `json:"port_id,omitempty"`
	Session              *uuid.UUID             `json:"session,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseWiredCoa,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseWiredCoa) String() string {
	return fmt.Sprintf(
		"ResponseWiredCoa[DeviceMac=%v, PortId=%v, Session=%v, AdditionalProperties=%v]",
		r.DeviceMac, r.PortId, r.Session, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseWiredCoa.
// It customizes the JSON marshaling process for ResponseWiredCoa objects.
func (r ResponseWiredCoa) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(r.AdditionalProperties,
		"device_mac", "port_id", "session"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(r.toMap())
}

// toMap converts the ResponseWiredCoa object to a map representation for JSON marshaling.
func (r ResponseWiredCoa) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, r.AdditionalProperties)
	if r.DeviceMac != nil {
		structMap["device_mac"] = r.DeviceMac
	}
	if r.PortId != nil {
		structMap["port_id"] = r.PortId
	}
	if r.Session != nil {
		structMap["session"] = r.Session
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseWiredCoa.
// It customizes the JSON unmarshaling process for ResponseWiredCoa objects.
func (r *ResponseWiredCoa) UnmarshalJSON(input []byte) error {
	var temp tempResponseWiredCoa
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "device_mac", "port_id", "session")
	if err != nil {
		return err
	}
	r.AdditionalProperties = additionalProperties

	r.DeviceMac = temp.DeviceMac
	r.PortId = temp.PortId
	r.Session = temp.Session
	return nil
}

// tempResponseWiredCoa is a temporary struct used for validating the fields of ResponseWiredCoa.
type tempResponseWiredCoa struct {
	DeviceMac *string    `json:"device_mac,omitempty"`
	PortId    *string    `json:"port_id,omitempty"`
	Session   *uuid.UUID `json:"session,omitempty"`
}
