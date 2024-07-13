package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// DswitchesMetricsInactiveWiredVlans represents a DswitchesMetricsInactiveWiredVlans struct.
type DswitchesMetricsInactiveWiredVlans struct {
    Details              interface{}    `json:"details"`
    Score                float64        `json:"score"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for DswitchesMetricsInactiveWiredVlans.
// It customizes the JSON marshaling process for DswitchesMetricsInactiveWiredVlans objects.
func (d DswitchesMetricsInactiveWiredVlans) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(d.toMap())
}

// toMap converts the DswitchesMetricsInactiveWiredVlans object to a map representation for JSON marshaling.
func (d DswitchesMetricsInactiveWiredVlans) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, d.AdditionalProperties)
    structMap["details"] = d.Details
    structMap["score"] = d.Score
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for DswitchesMetricsInactiveWiredVlans.
// It customizes the JSON unmarshaling process for DswitchesMetricsInactiveWiredVlans objects.
func (d *DswitchesMetricsInactiveWiredVlans) UnmarshalJSON(input []byte) error {
    var temp dswitchesMetricsInactiveWiredVlans
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

// dswitchesMetricsInactiveWiredVlans is a temporary struct used for validating the fields of DswitchesMetricsInactiveWiredVlans.
type dswitchesMetricsInactiveWiredVlans  struct {
    Details *interface{} `json:"details"`
    Score   *float64     `json:"score"`
}

func (d *dswitchesMetricsInactiveWiredVlans) validate() error {
    var errs []string
    if d.Details == nil {
        errs = append(errs, "required field `details` is missing for type `Dswitches_Metrics_Inactive_Wired_Vlans`")
    }
    if d.Score == nil {
        errs = append(errs, "required field `score` is missing for type `Dswitches_Metrics_Inactive_Wired_Vlans`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
