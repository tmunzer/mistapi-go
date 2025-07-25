// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// DswitchesMetricsSwitchApAffinity represents a DswitchesMetricsSwitchApAffinity struct.
type DswitchesMetricsSwitchApAffinity struct {
    Details              DswitchesMetricsSwitchApAffinityDetails `json:"details"`
    Score                float64                                 `json:"score"`
    AdditionalProperties map[string]interface{}                  `json:"_"`
}

// String implements the fmt.Stringer interface for DswitchesMetricsSwitchApAffinity,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (d DswitchesMetricsSwitchApAffinity) String() string {
    return fmt.Sprintf(
    	"DswitchesMetricsSwitchApAffinity[Details=%v, Score=%v, AdditionalProperties=%v]",
    	d.Details, d.Score, d.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for DswitchesMetricsSwitchApAffinity.
// It customizes the JSON marshaling process for DswitchesMetricsSwitchApAffinity objects.
func (d DswitchesMetricsSwitchApAffinity) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(d.AdditionalProperties,
        "details", "score"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(d.toMap())
}

// toMap converts the DswitchesMetricsSwitchApAffinity object to a map representation for JSON marshaling.
func (d DswitchesMetricsSwitchApAffinity) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, d.AdditionalProperties)
    structMap["details"] = d.Details.toMap()
    structMap["score"] = d.Score
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for DswitchesMetricsSwitchApAffinity.
// It customizes the JSON unmarshaling process for DswitchesMetricsSwitchApAffinity objects.
func (d *DswitchesMetricsSwitchApAffinity) UnmarshalJSON(input []byte) error {
    var temp tempDswitchesMetricsSwitchApAffinity
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "details", "score")
    if err != nil {
    	return err
    }
    d.AdditionalProperties = additionalProperties
    
    d.Details = *temp.Details
    d.Score = *temp.Score
    return nil
}

// tempDswitchesMetricsSwitchApAffinity is a temporary struct used for validating the fields of DswitchesMetricsSwitchApAffinity.
type tempDswitchesMetricsSwitchApAffinity  struct {
    Details *DswitchesMetricsSwitchApAffinityDetails `json:"details"`
    Score   *float64                                 `json:"score"`
}

func (d *tempDswitchesMetricsSwitchApAffinity) validate() error {
    var errs []string
    if d.Details == nil {
        errs = append(errs, "required field `details` is missing for type `dswitches_metrics_switch_ap_affinity`")
    }
    if d.Score == nil {
        errs = append(errs, "required field `score` is missing for type `dswitches_metrics_switch_ap_affinity`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
