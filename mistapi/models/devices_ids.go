// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"strings"
)

// DevicesIds represents a DevicesIds struct.
type DevicesIds struct {
	DeviceIds            []uuid.UUID            `json:"device_ids"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for DevicesIds,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (d DevicesIds) String() string {
	return fmt.Sprintf(
		"DevicesIds[DeviceIds=%v, AdditionalProperties=%v]",
		d.DeviceIds, d.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for DevicesIds.
// It customizes the JSON marshaling process for DevicesIds objects.
func (d DevicesIds) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(d.AdditionalProperties,
		"device_ids"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(d.toMap())
}

// toMap converts the DevicesIds object to a map representation for JSON marshaling.
func (d DevicesIds) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, d.AdditionalProperties)
	structMap["device_ids"] = d.DeviceIds
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for DevicesIds.
// It customizes the JSON unmarshaling process for DevicesIds objects.
func (d *DevicesIds) UnmarshalJSON(input []byte) error {
	var temp tempDevicesIds
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "device_ids")
	if err != nil {
		return err
	}
	d.AdditionalProperties = additionalProperties

	d.DeviceIds = *temp.DeviceIds
	return nil
}

// tempDevicesIds is a temporary struct used for validating the fields of DevicesIds.
type tempDevicesIds struct {
	DeviceIds *[]uuid.UUID `json:"device_ids"`
}

func (d *tempDevicesIds) validate() error {
	var errs []string
	if d.DeviceIds == nil {
		errs = append(errs, "required field `device_ids` is missing for type `devices_ids`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
