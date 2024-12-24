package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// ResponseDiscoveredSwitchMetrics represents a ResponseDiscoveredSwitchMetrics struct.
type ResponseDiscoveredSwitchMetrics struct {
    End                  float64                  `json:"end"`
    Limit                int                      `json:"limit"`
    Next                 *string                  `json:"next,omitempty"`
    Results              []DiscoveredSwitchMetric `json:"results"`
    Start                float64                  `json:"start"`
    Total                int                      `json:"total"`
    AdditionalProperties map[string]interface{}   `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseDiscoveredSwitchMetrics,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseDiscoveredSwitchMetrics) String() string {
    return fmt.Sprintf(
    	"ResponseDiscoveredSwitchMetrics[End=%v, Limit=%v, Next=%v, Results=%v, Start=%v, Total=%v, AdditionalProperties=%v]",
    	r.End, r.Limit, r.Next, r.Results, r.Start, r.Total, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseDiscoveredSwitchMetrics.
// It customizes the JSON marshaling process for ResponseDiscoveredSwitchMetrics objects.
func (r ResponseDiscoveredSwitchMetrics) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "end", "limit", "next", "results", "start", "total"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseDiscoveredSwitchMetrics object to a map representation for JSON marshaling.
func (r ResponseDiscoveredSwitchMetrics) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    structMap["end"] = r.End
    structMap["limit"] = r.Limit
    if r.Next != nil {
        structMap["next"] = r.Next
    }
    structMap["results"] = r.Results
    structMap["start"] = r.Start
    structMap["total"] = r.Total
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseDiscoveredSwitchMetrics.
// It customizes the JSON unmarshaling process for ResponseDiscoveredSwitchMetrics objects.
func (r *ResponseDiscoveredSwitchMetrics) UnmarshalJSON(input []byte) error {
    var temp tempResponseDiscoveredSwitchMetrics
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "end", "limit", "next", "results", "start", "total")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.End = *temp.End
    r.Limit = *temp.Limit
    r.Next = temp.Next
    r.Results = *temp.Results
    r.Start = *temp.Start
    r.Total = *temp.Total
    return nil
}

// tempResponseDiscoveredSwitchMetrics is a temporary struct used for validating the fields of ResponseDiscoveredSwitchMetrics.
type tempResponseDiscoveredSwitchMetrics  struct {
    End     *float64                  `json:"end"`
    Limit   *int                      `json:"limit"`
    Next    *string                   `json:"next,omitempty"`
    Results *[]DiscoveredSwitchMetric `json:"results"`
    Start   *float64                  `json:"start"`
    Total   *int                      `json:"total"`
}

func (r *tempResponseDiscoveredSwitchMetrics) validate() error {
    var errs []string
    if r.End == nil {
        errs = append(errs, "required field `end` is missing for type `response_discovered_switch_metrics`")
    }
    if r.Limit == nil {
        errs = append(errs, "required field `limit` is missing for type `response_discovered_switch_metrics`")
    }
    if r.Results == nil {
        errs = append(errs, "required field `results` is missing for type `response_discovered_switch_metrics`")
    }
    if r.Start == nil {
        errs = append(errs, "required field `start` is missing for type `response_discovered_switch_metrics`")
    }
    if r.Total == nil {
        errs = append(errs, "required field `total` is missing for type `response_discovered_switch_metrics`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
