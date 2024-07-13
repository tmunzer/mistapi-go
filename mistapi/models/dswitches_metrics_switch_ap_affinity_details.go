package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// DswitchesMetricsSwitchApAffinityDetails represents a DswitchesMetricsSwitchApAffinityDetails struct.
type DswitchesMetricsSwitchApAffinityDetails struct {
    SystemName           []string       `json:"system_name"`
    Threshold            float64        `json:"threshold"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for DswitchesMetricsSwitchApAffinityDetails.
// It customizes the JSON marshaling process for DswitchesMetricsSwitchApAffinityDetails objects.
func (d DswitchesMetricsSwitchApAffinityDetails) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(d.toMap())
}

// toMap converts the DswitchesMetricsSwitchApAffinityDetails object to a map representation for JSON marshaling.
func (d DswitchesMetricsSwitchApAffinityDetails) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, d.AdditionalProperties)
    structMap["system_name"] = d.SystemName
    structMap["threshold"] = d.Threshold
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for DswitchesMetricsSwitchApAffinityDetails.
// It customizes the JSON unmarshaling process for DswitchesMetricsSwitchApAffinityDetails objects.
func (d *DswitchesMetricsSwitchApAffinityDetails) UnmarshalJSON(input []byte) error {
    var temp dswitchesMetricsSwitchApAffinityDetails
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "system_name", "threshold")
    if err != nil {
    	return err
    }
    
    d.AdditionalProperties = additionalProperties
    d.SystemName = *temp.SystemName
    d.Threshold = *temp.Threshold
    return nil
}

// dswitchesMetricsSwitchApAffinityDetails is a temporary struct used for validating the fields of DswitchesMetricsSwitchApAffinityDetails.
type dswitchesMetricsSwitchApAffinityDetails  struct {
    SystemName *[]string `json:"system_name"`
    Threshold  *float64  `json:"threshold"`
}

func (d *dswitchesMetricsSwitchApAffinityDetails) validate() error {
    var errs []string
    if d.SystemName == nil {
        errs = append(errs, "required field `system_name` is missing for type `Dswitches_Metrics_Switch_Ap_Affinity_Details`")
    }
    if d.Threshold == nil {
        errs = append(errs, "required field `threshold` is missing for type `Dswitches_Metrics_Switch_Ap_Affinity_Details`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}