package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// InsightMetrics represents a InsightMetrics struct.
type InsightMetrics struct {
    End                  int                    `json:"end"`
    Interval             int                    `json:"interval"`
    // results depends on the `metric`
    Results              []interface{}          `json:"results"`
    Start                int                    `json:"start"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for InsightMetrics,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (i InsightMetrics) String() string {
    return fmt.Sprintf(
    	"InsightMetrics[End=%v, Interval=%v, Results=%v, Start=%v, AdditionalProperties=%v]",
    	i.End, i.Interval, i.Results, i.Start, i.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for InsightMetrics.
// It customizes the JSON marshaling process for InsightMetrics objects.
func (i InsightMetrics) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(i.AdditionalProperties,
        "end", "interval", "results", "start"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(i.toMap())
}

// toMap converts the InsightMetrics object to a map representation for JSON marshaling.
func (i InsightMetrics) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, i.AdditionalProperties)
    structMap["end"] = i.End
    structMap["interval"] = i.Interval
    structMap["results"] = i.Results
    structMap["start"] = i.Start
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for InsightMetrics.
// It customizes the JSON unmarshaling process for InsightMetrics objects.
func (i *InsightMetrics) UnmarshalJSON(input []byte) error {
    var temp tempInsightMetrics
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "end", "interval", "results", "start")
    if err != nil {
    	return err
    }
    i.AdditionalProperties = additionalProperties
    
    i.End = *temp.End
    i.Interval = *temp.Interval
    i.Results = *temp.Results
    i.Start = *temp.Start
    return nil
}

// tempInsightMetrics is a temporary struct used for validating the fields of InsightMetrics.
type tempInsightMetrics  struct {
    End      *int           `json:"end"`
    Interval *int           `json:"interval"`
    Results  *[]interface{} `json:"results"`
    Start    *int           `json:"start"`
}

func (i *tempInsightMetrics) validate() error {
    var errs []string
    if i.End == nil {
        errs = append(errs, "required field `end` is missing for type `insight_metrics`")
    }
    if i.Interval == nil {
        errs = append(errs, "required field `interval` is missing for type `insight_metrics`")
    }
    if i.Results == nil {
        errs = append(errs, "required field `results` is missing for type `insight_metrics`")
    }
    if i.Start == nil {
        errs = append(errs, "required field `start` is missing for type `insight_metrics`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
