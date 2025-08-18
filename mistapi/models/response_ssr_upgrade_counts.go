// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// ResponseSsrUpgradeCounts represents a ResponseSsrUpgradeCounts struct.
type ResponseSsrUpgradeCounts struct {
	Failed               int                    `json:"failed"`
	Queued               int                    `json:"queued"`
	Success              int                    `json:"success"`
	Upgrading            int                    `json:"upgrading"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseSsrUpgradeCounts,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseSsrUpgradeCounts) String() string {
	return fmt.Sprintf(
		"ResponseSsrUpgradeCounts[Failed=%v, Queued=%v, Success=%v, Upgrading=%v, AdditionalProperties=%v]",
		r.Failed, r.Queued, r.Success, r.Upgrading, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseSsrUpgradeCounts.
// It customizes the JSON marshaling process for ResponseSsrUpgradeCounts objects.
func (r ResponseSsrUpgradeCounts) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(r.AdditionalProperties,
		"failed", "queued", "success", "upgrading"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(r.toMap())
}

// toMap converts the ResponseSsrUpgradeCounts object to a map representation for JSON marshaling.
func (r ResponseSsrUpgradeCounts) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, r.AdditionalProperties)
	structMap["failed"] = r.Failed
	structMap["queued"] = r.Queued
	structMap["success"] = r.Success
	structMap["upgrading"] = r.Upgrading
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseSsrUpgradeCounts.
// It customizes the JSON unmarshaling process for ResponseSsrUpgradeCounts objects.
func (r *ResponseSsrUpgradeCounts) UnmarshalJSON(input []byte) error {
	var temp tempResponseSsrUpgradeCounts
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "failed", "queued", "success", "upgrading")
	if err != nil {
		return err
	}
	r.AdditionalProperties = additionalProperties

	r.Failed = *temp.Failed
	r.Queued = *temp.Queued
	r.Success = *temp.Success
	r.Upgrading = *temp.Upgrading
	return nil
}

// tempResponseSsrUpgradeCounts is a temporary struct used for validating the fields of ResponseSsrUpgradeCounts.
type tempResponseSsrUpgradeCounts struct {
	Failed    *int `json:"failed"`
	Queued    *int `json:"queued"`
	Success   *int `json:"success"`
	Upgrading *int `json:"upgrading"`
}

func (r *tempResponseSsrUpgradeCounts) validate() error {
	var errs []string
	if r.Failed == nil {
		errs = append(errs, "required field `failed` is missing for type `response_ssr_upgrade_counts`")
	}
	if r.Queued == nil {
		errs = append(errs, "required field `queued` is missing for type `response_ssr_upgrade_counts`")
	}
	if r.Success == nil {
		errs = append(errs, "required field `success` is missing for type `response_ssr_upgrade_counts`")
	}
	if r.Upgrading == nil {
		errs = append(errs, "required field `upgrading` is missing for type `response_ssr_upgrade_counts`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
