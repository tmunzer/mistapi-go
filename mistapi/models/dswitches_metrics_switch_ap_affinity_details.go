// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// DswitchesMetricsSwitchApAffinityDetails represents a DswitchesMetricsSwitchApAffinityDetails struct.
type DswitchesMetricsSwitchApAffinityDetails struct {
	SystemName           []string               `json:"system_name"`
	Threshold            float64                `json:"threshold"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for DswitchesMetricsSwitchApAffinityDetails,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (d DswitchesMetricsSwitchApAffinityDetails) String() string {
	return fmt.Sprintf(
		"DswitchesMetricsSwitchApAffinityDetails[SystemName=%v, Threshold=%v, AdditionalProperties=%v]",
		d.SystemName, d.Threshold, d.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for DswitchesMetricsSwitchApAffinityDetails.
// It customizes the JSON marshaling process for DswitchesMetricsSwitchApAffinityDetails objects.
func (d DswitchesMetricsSwitchApAffinityDetails) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(d.AdditionalProperties,
		"system_name", "threshold"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(d.toMap())
}

// toMap converts the DswitchesMetricsSwitchApAffinityDetails object to a map representation for JSON marshaling.
func (d DswitchesMetricsSwitchApAffinityDetails) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, d.AdditionalProperties)
	structMap["system_name"] = d.SystemName
	structMap["threshold"] = d.Threshold
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for DswitchesMetricsSwitchApAffinityDetails.
// It customizes the JSON unmarshaling process for DswitchesMetricsSwitchApAffinityDetails objects.
func (d *DswitchesMetricsSwitchApAffinityDetails) UnmarshalJSON(input []byte) error {
	var temp tempDswitchesMetricsSwitchApAffinityDetails
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "system_name", "threshold")
	if err != nil {
		return err
	}
	d.AdditionalProperties = additionalProperties

	d.SystemName = *temp.SystemName
	d.Threshold = *temp.Threshold
	return nil
}

// tempDswitchesMetricsSwitchApAffinityDetails is a temporary struct used for validating the fields of DswitchesMetricsSwitchApAffinityDetails.
type tempDswitchesMetricsSwitchApAffinityDetails struct {
	SystemName *[]string `json:"system_name"`
	Threshold  *float64  `json:"threshold"`
}

func (d *tempDswitchesMetricsSwitchApAffinityDetails) validate() error {
	var errs []string
	if d.SystemName == nil {
		errs = append(errs, "required field `system_name` is missing for type `dswitches_metrics_switch_ap_affinity_details`")
	}
	if d.Threshold == nil {
		errs = append(errs, "required field `threshold` is missing for type `dswitches_metrics_switch_ap_affinity_details`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
