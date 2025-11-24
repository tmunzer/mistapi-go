// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// DevicesGbpTag represents a DevicesGbpTag struct.
type DevicesGbpTag struct {
	GbpTag               int                    `json:"gbp_tag"`
	Macs                 []string               `json:"macs"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for DevicesGbpTag,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (d DevicesGbpTag) String() string {
	return fmt.Sprintf(
		"DevicesGbpTag[GbpTag=%v, Macs=%v, AdditionalProperties=%v]",
		d.GbpTag, d.Macs, d.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for DevicesGbpTag.
// It customizes the JSON marshaling process for DevicesGbpTag objects.
func (d DevicesGbpTag) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(d.AdditionalProperties,
		"gbp_tag", "macs"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(d.toMap())
}

// toMap converts the DevicesGbpTag object to a map representation for JSON marshaling.
func (d DevicesGbpTag) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, d.AdditionalProperties)
	structMap["gbp_tag"] = d.GbpTag
	structMap["macs"] = d.Macs
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for DevicesGbpTag.
// It customizes the JSON unmarshaling process for DevicesGbpTag objects.
func (d *DevicesGbpTag) UnmarshalJSON(input []byte) error {
	var temp tempDevicesGbpTag
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "gbp_tag", "macs")
	if err != nil {
		return err
	}
	d.AdditionalProperties = additionalProperties

	d.GbpTag = *temp.GbpTag
	d.Macs = *temp.Macs
	return nil
}

// tempDevicesGbpTag is a temporary struct used for validating the fields of DevicesGbpTag.
type tempDevicesGbpTag struct {
	GbpTag *int      `json:"gbp_tag"`
	Macs   *[]string `json:"macs"`
}

func (d *tempDevicesGbpTag) validate() error {
	var errs []string
	if d.GbpTag == nil {
		errs = append(errs, "required field `gbp_tag` is missing for type `devices_gbp_tag`")
	}
	if d.Macs == nil {
		errs = append(errs, "required field `macs` is missing for type `devices_gbp_tag`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
