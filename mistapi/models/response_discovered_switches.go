// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// ResponseDiscoveredSwitches represents a ResponseDiscoveredSwitches struct.
type ResponseDiscoveredSwitches struct {
    End                  float64                `json:"end"`
    Limit                int                    `json:"limit"`
    Next                 *string                `json:"next,omitempty"`
    Results              []DiscoveredSwitch     `json:"results"`
    Start                float64                `json:"start"`
    Total                int                    `json:"total"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseDiscoveredSwitches,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseDiscoveredSwitches) String() string {
    return fmt.Sprintf(
    	"ResponseDiscoveredSwitches[End=%v, Limit=%v, Next=%v, Results=%v, Start=%v, Total=%v, AdditionalProperties=%v]",
    	r.End, r.Limit, r.Next, r.Results, r.Start, r.Total, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseDiscoveredSwitches.
// It customizes the JSON marshaling process for ResponseDiscoveredSwitches objects.
func (r ResponseDiscoveredSwitches) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "end", "limit", "next", "results", "start", "total"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseDiscoveredSwitches object to a map representation for JSON marshaling.
func (r ResponseDiscoveredSwitches) toMap() map[string]any {
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

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseDiscoveredSwitches.
// It customizes the JSON unmarshaling process for ResponseDiscoveredSwitches objects.
func (r *ResponseDiscoveredSwitches) UnmarshalJSON(input []byte) error {
    var temp tempResponseDiscoveredSwitches
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

// tempResponseDiscoveredSwitches is a temporary struct used for validating the fields of ResponseDiscoveredSwitches.
type tempResponseDiscoveredSwitches  struct {
    End     *float64            `json:"end"`
    Limit   *int                `json:"limit"`
    Next    *string             `json:"next,omitempty"`
    Results *[]DiscoveredSwitch `json:"results"`
    Start   *float64            `json:"start"`
    Total   *int                `json:"total"`
}

func (r *tempResponseDiscoveredSwitches) validate() error {
    var errs []string
    if r.End == nil {
        errs = append(errs, "required field `end` is missing for type `response_discovered_switches`")
    }
    if r.Limit == nil {
        errs = append(errs, "required field `limit` is missing for type `response_discovered_switches`")
    }
    if r.Results == nil {
        errs = append(errs, "required field `results` is missing for type `response_discovered_switches`")
    }
    if r.Start == nil {
        errs = append(errs, "required field `start` is missing for type `response_discovered_switches`")
    }
    if r.Total == nil {
        errs = append(errs, "required field `total` is missing for type `response_discovered_switches`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
