// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// DswitchesMetricsPoeComplianceDetails represents a DswitchesMetricsPoeComplianceDetails struct.
type DswitchesMetricsPoeComplianceDetails struct {
    TotalAps             int                    `json:"total_aps"`
    TotalPower           float64                `json:"total_power"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for DswitchesMetricsPoeComplianceDetails,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (d DswitchesMetricsPoeComplianceDetails) String() string {
    return fmt.Sprintf(
    	"DswitchesMetricsPoeComplianceDetails[TotalAps=%v, TotalPower=%v, AdditionalProperties=%v]",
    	d.TotalAps, d.TotalPower, d.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for DswitchesMetricsPoeComplianceDetails.
// It customizes the JSON marshaling process for DswitchesMetricsPoeComplianceDetails objects.
func (d DswitchesMetricsPoeComplianceDetails) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(d.AdditionalProperties,
        "total_aps", "total_power"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(d.toMap())
}

// toMap converts the DswitchesMetricsPoeComplianceDetails object to a map representation for JSON marshaling.
func (d DswitchesMetricsPoeComplianceDetails) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, d.AdditionalProperties)
    structMap["total_aps"] = d.TotalAps
    structMap["total_power"] = d.TotalPower
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for DswitchesMetricsPoeComplianceDetails.
// It customizes the JSON unmarshaling process for DswitchesMetricsPoeComplianceDetails objects.
func (d *DswitchesMetricsPoeComplianceDetails) UnmarshalJSON(input []byte) error {
    var temp tempDswitchesMetricsPoeComplianceDetails
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "total_aps", "total_power")
    if err != nil {
    	return err
    }
    d.AdditionalProperties = additionalProperties
    
    d.TotalAps = *temp.TotalAps
    d.TotalPower = *temp.TotalPower
    return nil
}

// tempDswitchesMetricsPoeComplianceDetails is a temporary struct used for validating the fields of DswitchesMetricsPoeComplianceDetails.
type tempDswitchesMetricsPoeComplianceDetails  struct {
    TotalAps   *int     `json:"total_aps"`
    TotalPower *float64 `json:"total_power"`
}

func (d *tempDswitchesMetricsPoeComplianceDetails) validate() error {
    var errs []string
    if d.TotalAps == nil {
        errs = append(errs, "required field `total_aps` is missing for type `dswitches_metrics_poe_compliance_details`")
    }
    if d.TotalPower == nil {
        errs = append(errs, "required field `total_power` is missing for type `dswitches_metrics_poe_compliance_details`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
