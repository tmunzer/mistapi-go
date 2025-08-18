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

// ResponseSsrUpgrade represents a ResponseSsrUpgrade struct.
type ResponseSsrUpgrade struct {
	Channel    string                   `json:"channel"`
	Counts     ResponseSsrUpgradeCounts `json:"counts"`
	DeviceType string                   `json:"device_type"`
	// Unique ID of the object instance in the Mist Organization
	Id                   uuid.UUID              `json:"id"`
	Status               string                 `json:"status"`
	Strategy             string                 `json:"strategy"`
	Versions             map[string]string      `json:"versions"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseSsrUpgrade,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseSsrUpgrade) String() string {
	return fmt.Sprintf(
		"ResponseSsrUpgrade[Channel=%v, Counts=%v, DeviceType=%v, Id=%v, Status=%v, Strategy=%v, Versions=%v, AdditionalProperties=%v]",
		r.Channel, r.Counts, r.DeviceType, r.Id, r.Status, r.Strategy, r.Versions, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseSsrUpgrade.
// It customizes the JSON marshaling process for ResponseSsrUpgrade objects.
func (r ResponseSsrUpgrade) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(r.AdditionalProperties,
		"channel", "counts", "device_type", "id", "status", "strategy", "versions"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(r.toMap())
}

// toMap converts the ResponseSsrUpgrade object to a map representation for JSON marshaling.
func (r ResponseSsrUpgrade) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, r.AdditionalProperties)
	structMap["channel"] = r.Channel
	structMap["counts"] = r.Counts.toMap()
	structMap["device_type"] = r.DeviceType
	structMap["id"] = r.Id
	structMap["status"] = r.Status
	structMap["strategy"] = r.Strategy
	structMap["versions"] = r.Versions
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseSsrUpgrade.
// It customizes the JSON unmarshaling process for ResponseSsrUpgrade objects.
func (r *ResponseSsrUpgrade) UnmarshalJSON(input []byte) error {
	var temp tempResponseSsrUpgrade
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "channel", "counts", "device_type", "id", "status", "strategy", "versions")
	if err != nil {
		return err
	}
	r.AdditionalProperties = additionalProperties

	r.Channel = *temp.Channel
	r.Counts = *temp.Counts
	r.DeviceType = *temp.DeviceType
	r.Id = *temp.Id
	r.Status = *temp.Status
	r.Strategy = *temp.Strategy
	r.Versions = *temp.Versions
	return nil
}

// tempResponseSsrUpgrade is a temporary struct used for validating the fields of ResponseSsrUpgrade.
type tempResponseSsrUpgrade struct {
	Channel    *string                   `json:"channel"`
	Counts     *ResponseSsrUpgradeCounts `json:"counts"`
	DeviceType *string                   `json:"device_type"`
	Id         *uuid.UUID                `json:"id"`
	Status     *string                   `json:"status"`
	Strategy   *string                   `json:"strategy"`
	Versions   *map[string]string        `json:"versions"`
}

func (r *tempResponseSsrUpgrade) validate() error {
	var errs []string
	if r.Channel == nil {
		errs = append(errs, "required field `channel` is missing for type `response_ssr_upgrade`")
	}
	if r.Counts == nil {
		errs = append(errs, "required field `counts` is missing for type `response_ssr_upgrade`")
	}
	if r.DeviceType == nil {
		errs = append(errs, "required field `device_type` is missing for type `response_ssr_upgrade`")
	}
	if r.Id == nil {
		errs = append(errs, "required field `id` is missing for type `response_ssr_upgrade`")
	}
	if r.Status == nil {
		errs = append(errs, "required field `status` is missing for type `response_ssr_upgrade`")
	}
	if r.Strategy == nil {
		errs = append(errs, "required field `strategy` is missing for type `response_ssr_upgrade`")
	}
	if r.Versions == nil {
		errs = append(errs, "required field `versions` is missing for type `response_ssr_upgrade`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
