// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// ResponseDeviceEventsSearch represents a ResponseDeviceEventsSearch struct.
type ResponseDeviceEventsSearch struct {
    End                  int                    `json:"end"`
    Limit                int                    `json:"limit"`
    Next                 *string                `json:"next,omitempty"`
    Results              []EventsDeviceAp       `json:"results"`
    Start                int                    `json:"start"`
    Total                int                    `json:"total"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseDeviceEventsSearch,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseDeviceEventsSearch) String() string {
    return fmt.Sprintf(
    	"ResponseDeviceEventsSearch[End=%v, Limit=%v, Next=%v, Results=%v, Start=%v, Total=%v, AdditionalProperties=%v]",
    	r.End, r.Limit, r.Next, r.Results, r.Start, r.Total, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseDeviceEventsSearch.
// It customizes the JSON marshaling process for ResponseDeviceEventsSearch objects.
func (r ResponseDeviceEventsSearch) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "end", "limit", "next", "results", "start", "total"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseDeviceEventsSearch object to a map representation for JSON marshaling.
func (r ResponseDeviceEventsSearch) toMap() map[string]any {
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

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseDeviceEventsSearch.
// It customizes the JSON unmarshaling process for ResponseDeviceEventsSearch objects.
func (r *ResponseDeviceEventsSearch) UnmarshalJSON(input []byte) error {
    var temp tempResponseDeviceEventsSearch
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

// tempResponseDeviceEventsSearch is a temporary struct used for validating the fields of ResponseDeviceEventsSearch.
type tempResponseDeviceEventsSearch  struct {
    End     *int              `json:"end"`
    Limit   *int              `json:"limit"`
    Next    *string           `json:"next,omitempty"`
    Results *[]EventsDeviceAp `json:"results"`
    Start   *int              `json:"start"`
    Total   *int              `json:"total"`
}

func (r *tempResponseDeviceEventsSearch) validate() error {
    var errs []string
    if r.End == nil {
        errs = append(errs, "required field `end` is missing for type `response_device_events_search`")
    }
    if r.Limit == nil {
        errs = append(errs, "required field `limit` is missing for type `response_device_events_search`")
    }
    if r.Results == nil {
        errs = append(errs, "required field `results` is missing for type `response_device_events_search`")
    }
    if r.Start == nil {
        errs = append(errs, "required field `start` is missing for type `response_device_events_search`")
    }
    if r.Total == nil {
        errs = append(errs, "required field `total` is missing for type `response_device_events_search`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
