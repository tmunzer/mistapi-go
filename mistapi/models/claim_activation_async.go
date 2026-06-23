// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// ClaimActivationAsync represents a ClaimActivationAsync struct.
// Request to schedule an asynchronous inventory claim
type ClaimActivationAsync struct {
	// Activation code to claim
	Code string `json:"code"`
	// enum: `ap`, `gateway`, `switch`
	DeviceType *DeviceTypeDefaultApEnum `json:"device_type,omitempty"`
	// Claim scope for async inventory claiming. enum: `all`, `inventory`
	Type                 ClaimTypeAsyncEnum     `json:"type"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ClaimActivationAsync,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c ClaimActivationAsync) String() string {
	return fmt.Sprintf(
		"ClaimActivationAsync[Code=%v, DeviceType=%v, Type=%v, AdditionalProperties=%v]",
		c.Code, c.DeviceType, c.Type, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ClaimActivationAsync.
// It customizes the JSON marshaling process for ClaimActivationAsync objects.
func (c ClaimActivationAsync) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(c.AdditionalProperties,
		"code", "device_type", "type"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(c.toMap())
}

// toMap converts the ClaimActivationAsync object to a map representation for JSON marshaling.
func (c ClaimActivationAsync) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, c.AdditionalProperties)
	structMap["code"] = c.Code
	if c.DeviceType != nil {
		structMap["device_type"] = c.DeviceType
	}
	structMap["type"] = c.Type
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ClaimActivationAsync.
// It customizes the JSON unmarshaling process for ClaimActivationAsync objects.
func (c *ClaimActivationAsync) UnmarshalJSON(input []byte) error {
	var temp tempClaimActivationAsync
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "code", "device_type", "type")
	if err != nil {
		return err
	}
	c.AdditionalProperties = additionalProperties

	c.Code = *temp.Code
	c.DeviceType = temp.DeviceType
	c.Type = *temp.Type
	return nil
}

// tempClaimActivationAsync is a temporary struct used for validating the fields of ClaimActivationAsync.
type tempClaimActivationAsync struct {
	Code       *string                  `json:"code"`
	DeviceType *DeviceTypeDefaultApEnum `json:"device_type,omitempty"`
	Type       *ClaimTypeAsyncEnum      `json:"type"`
}

func (c *tempClaimActivationAsync) validate() error {
	var errs []string
	if c.Code == nil {
		errs = append(errs, "required field `code` is missing for type `claim_activation_async`")
	}
	if c.Type == nil {
		errs = append(errs, "required field `type` is missing for type `claim_activation_async`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
