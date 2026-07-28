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

// ResponseSsrUpgradeStatus represents a ResponseSsrUpgradeStatus struct.
// Detailed status for an SSR firmware upgrade job
type ResponseSsrUpgradeStatus struct {
	// Firmware release channel used for the SSR upgrade
	Channel string `json:"channel"`
	// Type of devices targeted by the SSR upgrade
	DeviceType *string `json:"device_type,omitempty"`
	// Whether the upgrade was forced even when the requested version matched the running version
	Force *bool `json:"force,omitempty"`
	// Unique ID of the object instance in the Mist Organization
	Id uuid.UUID `json:"id"`
	// Current status of the SSR upgrade job
	Status string `json:"status"`
	// Upgrade strategy used by the SSR upgrade job
	Strategy *string `json:"strategy,omitempty"`
	// SSR device IDs grouped by upgrade status
	Targets ResponseSsrUpgradeStatusTargets `json:"targets"`
	// SSR firmware versions included in the upgrade job
	Versions             interface{}            `json:"versions"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseSsrUpgradeStatus,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseSsrUpgradeStatus) String() string {
	return fmt.Sprintf(
		"ResponseSsrUpgradeStatus[Channel=%v, DeviceType=%v, Force=%v, Id=%v, Status=%v, Strategy=%v, Targets=%v, Versions=%v, AdditionalProperties=%v]",
		r.Channel, r.DeviceType, r.Force, r.Id, r.Status, r.Strategy, r.Targets, r.Versions, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseSsrUpgradeStatus.
// It customizes the JSON marshaling process for ResponseSsrUpgradeStatus objects.
func (r ResponseSsrUpgradeStatus) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(r.AdditionalProperties,
		"channel", "device_type", "force", "id", "status", "strategy", "targets", "versions"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(r.toMap())
}

// toMap converts the ResponseSsrUpgradeStatus object to a map representation for JSON marshaling.
func (r ResponseSsrUpgradeStatus) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, r.AdditionalProperties)
	structMap["channel"] = r.Channel
	if r.DeviceType != nil {
		structMap["device_type"] = r.DeviceType
	}
	if r.Force != nil {
		structMap["force"] = r.Force
	}
	structMap["id"] = r.Id
	structMap["status"] = r.Status
	if r.Strategy != nil {
		structMap["strategy"] = r.Strategy
	}
	structMap["targets"] = r.Targets.toMap()
	structMap["versions"] = r.Versions
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseSsrUpgradeStatus.
// It customizes the JSON unmarshaling process for ResponseSsrUpgradeStatus objects.
func (r *ResponseSsrUpgradeStatus) UnmarshalJSON(input []byte) error {
	var temp tempResponseSsrUpgradeStatus
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "channel", "device_type", "force", "id", "status", "strategy", "targets", "versions")
	if err != nil {
		return err
	}
	r.AdditionalProperties = additionalProperties

	r.Channel = *temp.Channel
	r.DeviceType = temp.DeviceType
	r.Force = temp.Force
	r.Id = *temp.Id
	r.Status = *temp.Status
	r.Strategy = temp.Strategy
	r.Targets = *temp.Targets
	r.Versions = *temp.Versions
	return nil
}

// tempResponseSsrUpgradeStatus is a temporary struct used for validating the fields of ResponseSsrUpgradeStatus.
type tempResponseSsrUpgradeStatus struct {
	Channel    *string                          `json:"channel"`
	DeviceType *string                          `json:"device_type,omitempty"`
	Force      *bool                            `json:"force,omitempty"`
	Id         *uuid.UUID                       `json:"id"`
	Status     *string                          `json:"status"`
	Strategy   *string                          `json:"strategy,omitempty"`
	Targets    *ResponseSsrUpgradeStatusTargets `json:"targets"`
	Versions   *interface{}                     `json:"versions"`
}

func (r *tempResponseSsrUpgradeStatus) validate() error {
	var errs []string
	if r.Channel == nil {
		errs = append(errs, "required field `channel` is missing for type `response_ssr_upgrade_status`")
	}
	if r.Id == nil {
		errs = append(errs, "required field `id` is missing for type `response_ssr_upgrade_status`")
	}
	if r.Status == nil {
		errs = append(errs, "required field `status` is missing for type `response_ssr_upgrade_status`")
	}
	if r.Targets == nil {
		errs = append(errs, "required field `targets` is missing for type `response_ssr_upgrade_status`")
	}
	if r.Versions == nil {
		errs = append(errs, "required field `versions` is missing for type `response_ssr_upgrade_status`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
