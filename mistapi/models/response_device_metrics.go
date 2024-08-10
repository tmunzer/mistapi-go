package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ResponseDeviceMetrics represents a ResponseDeviceMetrics struct.
type ResponseDeviceMetrics struct {
    End                  int                                 `json:"end"`
    Interval             int                                 `json:"interval"`
    Results              []ResponseDeviceMetricsResultsItems `json:"results"`
    Rt                   []string                            `json:"rt,omitempty"`
    Start                int                                 `json:"start"`
    AdditionalProperties map[string]any                      `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseDeviceMetrics.
// It customizes the JSON marshaling process for ResponseDeviceMetrics objects.
func (r ResponseDeviceMetrics) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseDeviceMetrics object to a map representation for JSON marshaling.
func (r ResponseDeviceMetrics) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    structMap["end"] = r.End
    structMap["interval"] = r.Interval
    structMap["results"] = r.Results
    if r.Rt != nil {
        structMap["rt"] = r.Rt
    }
    structMap["start"] = r.Start
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseDeviceMetrics.
// It customizes the JSON unmarshaling process for ResponseDeviceMetrics objects.
func (r *ResponseDeviceMetrics) UnmarshalJSON(input []byte) error {
    var temp tempResponseDeviceMetrics
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "end", "interval", "results", "rt", "start")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.End = *temp.End
    r.Interval = *temp.Interval
    r.Results = *temp.Results
    r.Rt = temp.Rt
    r.Start = *temp.Start
    return nil
}

// tempResponseDeviceMetrics is a temporary struct used for validating the fields of ResponseDeviceMetrics.
type tempResponseDeviceMetrics  struct {
    End      *int                                 `json:"end"`
    Interval *int                                 `json:"interval"`
    Results  *[]ResponseDeviceMetricsResultsItems `json:"results"`
    Rt       []string                             `json:"rt,omitempty"`
    Start    *int                                 `json:"start"`
}

func (r *tempResponseDeviceMetrics) validate() error {
    var errs []string
    if r.End == nil {
        errs = append(errs, "required field `end` is missing for type `response_device_metrics`")
    }
    if r.Interval == nil {
        errs = append(errs, "required field `interval` is missing for type `response_device_metrics`")
    }
    if r.Results == nil {
        errs = append(errs, "required field `results` is missing for type `response_device_metrics`")
    }
    if r.Start == nil {
        errs = append(errs, "required field `start` is missing for type `response_device_metrics`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
