// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// UtilsZigbeeKick represents a UtilsZigbeeKick struct.
// Request body for kicking one or more Zigbee clients from an AP
type UtilsZigbeeKick struct {
	// One or more Zigbee EUI-64 (8-byte) MACs. Accepts colon-separated (`00:17:7a:01:06:0c:ae:9f`) or plain hex (`00177a01060cae9f`). Must be non-empty.
	Macs                 []string               `json:"macs"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for UtilsZigbeeKick,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UtilsZigbeeKick) String() string {
	return fmt.Sprintf(
		"UtilsZigbeeKick[Macs=%v, AdditionalProperties=%v]",
		u.Macs, u.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for UtilsZigbeeKick.
// It customizes the JSON marshaling process for UtilsZigbeeKick objects.
func (u UtilsZigbeeKick) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(u.AdditionalProperties,
		"macs"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(u.toMap())
}

// toMap converts the UtilsZigbeeKick object to a map representation for JSON marshaling.
func (u UtilsZigbeeKick) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, u.AdditionalProperties)
	structMap["macs"] = u.Macs
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UtilsZigbeeKick.
// It customizes the JSON unmarshaling process for UtilsZigbeeKick objects.
func (u *UtilsZigbeeKick) UnmarshalJSON(input []byte) error {
	var temp tempUtilsZigbeeKick
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "macs")
	if err != nil {
		return err
	}
	u.AdditionalProperties = additionalProperties

	u.Macs = *temp.Macs
	return nil
}

// tempUtilsZigbeeKick is a temporary struct used for validating the fields of UtilsZigbeeKick.
type tempUtilsZigbeeKick struct {
	Macs *[]string `json:"macs"`
}

func (u *tempUtilsZigbeeKick) validate() error {
	var errs []string
	if u.Macs == nil {
		errs = append(errs, "required field `macs` is missing for type `utils_zigbee_kick`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
