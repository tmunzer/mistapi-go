package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// DswitchesMetricsSwitchApAffinity represents a DswitchesMetricsSwitchApAffinity struct.
type DswitchesMetricsSwitchApAffinity struct {
    Details              DswitchesMetricsSwitchApAffinityDetails `json:"details"`
    Score                float64                                 `json:"score"`
    AdditionalProperties map[string]any                          `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for DswitchesMetricsSwitchApAffinity.
// It customizes the JSON marshaling process for DswitchesMetricsSwitchApAffinity objects.
func (d DswitchesMetricsSwitchApAffinity) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(d.toMap())
}

// toMap converts the DswitchesMetricsSwitchApAffinity object to a map representation for JSON marshaling.
func (d DswitchesMetricsSwitchApAffinity) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, d.AdditionalProperties)
    structMap["details"] = d.Details.toMap()
    structMap["score"] = d.Score
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for DswitchesMetricsSwitchApAffinity.
// It customizes the JSON unmarshaling process for DswitchesMetricsSwitchApAffinity objects.
func (d *DswitchesMetricsSwitchApAffinity) UnmarshalJSON(input []byte) error {
    var temp dswitchesMetricsSwitchApAffinity
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "details", "score")
    if err != nil {
    	return err
    }
    
    d.AdditionalProperties = additionalProperties
    d.Details = *temp.Details
    d.Score = *temp.Score
    return nil
}

// dswitchesMetricsSwitchApAffinity is a temporary struct used for validating the fields of DswitchesMetricsSwitchApAffinity.
type dswitchesMetricsSwitchApAffinity  struct {
    Details *DswitchesMetricsSwitchApAffinityDetails `json:"details"`
    Score   *float64                                 `json:"score"`
}

func (d *dswitchesMetricsSwitchApAffinity) validate() error {
    var errs []string
    if d.Details == nil {
        errs = append(errs, "required field `details` is missing for type `Dswitches_Metrics_Switch_Ap_Affinity`")
    }
    if d.Score == nil {
        errs = append(errs, "required field `score` is missing for type `Dswitches_Metrics_Switch_Ap_Affinity`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
