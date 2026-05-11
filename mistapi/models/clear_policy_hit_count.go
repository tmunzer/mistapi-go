// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// ClearPolicyHitCount represents a ClearPolicyHitCount struct.
type ClearPolicyHitCount struct {
	PolicyName           string                 `json:"policy_name"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ClearPolicyHitCount,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c ClearPolicyHitCount) String() string {
	return fmt.Sprintf(
		"ClearPolicyHitCount[PolicyName=%v, AdditionalProperties=%v]",
		c.PolicyName, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ClearPolicyHitCount.
// It customizes the JSON marshaling process for ClearPolicyHitCount objects.
func (c ClearPolicyHitCount) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(c.AdditionalProperties,
		"policy_name"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(c.toMap())
}

// toMap converts the ClearPolicyHitCount object to a map representation for JSON marshaling.
func (c ClearPolicyHitCount) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, c.AdditionalProperties)
	structMap["policy_name"] = c.PolicyName
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ClearPolicyHitCount.
// It customizes the JSON unmarshaling process for ClearPolicyHitCount objects.
func (c *ClearPolicyHitCount) UnmarshalJSON(input []byte) error {
	var temp tempClearPolicyHitCount
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "policy_name")
	if err != nil {
		return err
	}
	c.AdditionalProperties = additionalProperties

	c.PolicyName = *temp.PolicyName
	return nil
}

// tempClearPolicyHitCount is a temporary struct used for validating the fields of ClearPolicyHitCount.
type tempClearPolicyHitCount struct {
	PolicyName *string `json:"policy_name"`
}

func (c *tempClearPolicyHitCount) validate() error {
	var errs []string
	if c.PolicyName == nil {
		errs = append(errs, "required field `policy_name` is missing for type `clear_policy_hit_count`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
