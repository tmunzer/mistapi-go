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

// UtilsDevicesRestartMulti represents a UtilsDevicesRestartMulti struct.
type UtilsDevicesRestartMulti struct {
	DeviceIds []uuid.UUID `json:"device_ids"`
	// only for SSR: if node is not present, both nodes are restarted. For other devices: node should not be present
	Node                 *string                `json:"node,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for UtilsDevicesRestartMulti,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UtilsDevicesRestartMulti) String() string {
	return fmt.Sprintf(
		"UtilsDevicesRestartMulti[DeviceIds=%v, Node=%v, AdditionalProperties=%v]",
		u.DeviceIds, u.Node, u.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for UtilsDevicesRestartMulti.
// It customizes the JSON marshaling process for UtilsDevicesRestartMulti objects.
func (u UtilsDevicesRestartMulti) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(u.AdditionalProperties,
		"device_ids", "node"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(u.toMap())
}

// toMap converts the UtilsDevicesRestartMulti object to a map representation for JSON marshaling.
func (u UtilsDevicesRestartMulti) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, u.AdditionalProperties)
	structMap["device_ids"] = u.DeviceIds
	if u.Node != nil {
		structMap["node"] = u.Node
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UtilsDevicesRestartMulti.
// It customizes the JSON unmarshaling process for UtilsDevicesRestartMulti objects.
func (u *UtilsDevicesRestartMulti) UnmarshalJSON(input []byte) error {
	var temp tempUtilsDevicesRestartMulti
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "device_ids", "node")
	if err != nil {
		return err
	}
	u.AdditionalProperties = additionalProperties

	u.DeviceIds = *temp.DeviceIds
	u.Node = temp.Node
	return nil
}

// tempUtilsDevicesRestartMulti is a temporary struct used for validating the fields of UtilsDevicesRestartMulti.
type tempUtilsDevicesRestartMulti struct {
	DeviceIds *[]uuid.UUID `json:"device_ids"`
	Node      *string      `json:"node,omitempty"`
}

func (u *tempUtilsDevicesRestartMulti) validate() error {
	var errs []string
	if u.DeviceIds == nil {
		errs = append(errs, "required field `device_ids` is missing for type `utils_devices_restart_multi`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
