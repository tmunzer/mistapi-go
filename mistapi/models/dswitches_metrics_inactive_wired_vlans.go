package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// DswitchesMetricsInactiveWiredVlans represents a DswitchesMetricsInactiveWiredVlans struct.
type DswitchesMetricsInactiveWiredVlans struct {
    Details              interface{}            `json:"details"`
    Score                float64                `json:"score"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for DswitchesMetricsInactiveWiredVlans,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (d DswitchesMetricsInactiveWiredVlans) String() string {
    return fmt.Sprintf(
    	"DswitchesMetricsInactiveWiredVlans[Details=%v, Score=%v, AdditionalProperties=%v]",
    	d.Details, d.Score, d.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for DswitchesMetricsInactiveWiredVlans.
// It customizes the JSON marshaling process for DswitchesMetricsInactiveWiredVlans objects.
func (d DswitchesMetricsInactiveWiredVlans) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(d.AdditionalProperties,
        "details", "score"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(d.toMap())
}

// toMap converts the DswitchesMetricsInactiveWiredVlans object to a map representation for JSON marshaling.
func (d DswitchesMetricsInactiveWiredVlans) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, d.AdditionalProperties)
    structMap["details"] = d.Details
    structMap["score"] = d.Score
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for DswitchesMetricsInactiveWiredVlans.
// It customizes the JSON unmarshaling process for DswitchesMetricsInactiveWiredVlans objects.
func (d *DswitchesMetricsInactiveWiredVlans) UnmarshalJSON(input []byte) error {
    var temp tempDswitchesMetricsInactiveWiredVlans
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

// tempDswitchesMetricsInactiveWiredVlans is a temporary struct used for validating the fields of DswitchesMetricsInactiveWiredVlans.
type tempDswitchesMetricsInactiveWiredVlans  struct {
    Details *interface{} `json:"details"`
    Score   *float64     `json:"score"`
}

func (d *tempDswitchesMetricsInactiveWiredVlans) validate() error {
    var errs []string
    if d.Details == nil {
        errs = append(errs, "required field `details` is missing for type `dswitches_metrics_inactive_wired_vlans`")
    }
    if d.Score == nil {
        errs = append(errs, "required field `score` is missing for type `dswitches_metrics_inactive_wired_vlans`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
