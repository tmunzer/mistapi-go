package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// DswitchesMetricsPoeCompliance represents a DswitchesMetricsPoeCompliance struct.
type DswitchesMetricsPoeCompliance struct {
    Details              DswitchesMetricsPoeComplianceDetails `json:"details"`
    Score                float64                              `json:"score"`
    AdditionalProperties map[string]interface{}               `json:"_"`
}

// String implements the fmt.Stringer interface for DswitchesMetricsPoeCompliance,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (d DswitchesMetricsPoeCompliance) String() string {
    return fmt.Sprintf(
    	"DswitchesMetricsPoeCompliance[Details=%v, Score=%v, AdditionalProperties=%v]",
    	d.Details, d.Score, d.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for DswitchesMetricsPoeCompliance.
// It customizes the JSON marshaling process for DswitchesMetricsPoeCompliance objects.
func (d DswitchesMetricsPoeCompliance) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(d.AdditionalProperties,
        "details", "score"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(d.toMap())
}

// toMap converts the DswitchesMetricsPoeCompliance object to a map representation for JSON marshaling.
func (d DswitchesMetricsPoeCompliance) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, d.AdditionalProperties)
    structMap["details"] = d.Details.toMap()
    structMap["score"] = d.Score
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for DswitchesMetricsPoeCompliance.
// It customizes the JSON unmarshaling process for DswitchesMetricsPoeCompliance objects.
func (d *DswitchesMetricsPoeCompliance) UnmarshalJSON(input []byte) error {
    var temp tempDswitchesMetricsPoeCompliance
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

// tempDswitchesMetricsPoeCompliance is a temporary struct used for validating the fields of DswitchesMetricsPoeCompliance.
type tempDswitchesMetricsPoeCompliance  struct {
    Details *DswitchesMetricsPoeComplianceDetails `json:"details"`
    Score   *float64                              `json:"score"`
}

func (d *tempDswitchesMetricsPoeCompliance) validate() error {
    var errs []string
    if d.Details == nil {
        errs = append(errs, "required field `details` is missing for type `dswitches_metrics_poe_compliance`")
    }
    if d.Score == nil {
        errs = append(errs, "required field `score` is missing for type `dswitches_metrics_poe_compliance`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
