// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// AlarmSearchResult represents a AlarmSearchResult struct.
type AlarmSearchResult struct {
    // Component of the alarm
    Component            *string                `json:"component,omitempty"`
    End                  int                    `json:"end"`
    Limit                int                    `json:"limit"`
    Next                 *string                `json:"next,omitempty"`
    Page                 *int                   `json:"page,omitempty"`
    Results              []Alarm                `json:"results"`
    Start                int                    `json:"start"`
    Total                int                    `json:"total"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for AlarmSearchResult,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a AlarmSearchResult) String() string {
    return fmt.Sprintf(
    	"AlarmSearchResult[Component=%v, End=%v, Limit=%v, Next=%v, Page=%v, Results=%v, Start=%v, Total=%v, AdditionalProperties=%v]",
    	a.Component, a.End, a.Limit, a.Next, a.Page, a.Results, a.Start, a.Total, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for AlarmSearchResult.
// It customizes the JSON marshaling process for AlarmSearchResult objects.
func (a AlarmSearchResult) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(a.AdditionalProperties,
        "component", "end", "limit", "next", "page", "results", "start", "total"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(a.toMap())
}

// toMap converts the AlarmSearchResult object to a map representation for JSON marshaling.
func (a AlarmSearchResult) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, a.AdditionalProperties)
    if a.Component != nil {
        structMap["component"] = a.Component
    }
    structMap["end"] = a.End
    structMap["limit"] = a.Limit
    if a.Next != nil {
        structMap["next"] = a.Next
    }
    if a.Page != nil {
        structMap["page"] = a.Page
    }
    structMap["results"] = a.Results
    structMap["start"] = a.Start
    structMap["total"] = a.Total
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AlarmSearchResult.
// It customizes the JSON unmarshaling process for AlarmSearchResult objects.
func (a *AlarmSearchResult) UnmarshalJSON(input []byte) error {
    var temp tempAlarmSearchResult
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "component", "end", "limit", "next", "page", "results", "start", "total")
    if err != nil {
    	return err
    }
    a.AdditionalProperties = additionalProperties
    
    a.Component = temp.Component
    a.End = *temp.End
    a.Limit = *temp.Limit
    a.Next = temp.Next
    a.Page = temp.Page
    a.Results = *temp.Results
    a.Start = *temp.Start
    a.Total = *temp.Total
    return nil
}

// tempAlarmSearchResult is a temporary struct used for validating the fields of AlarmSearchResult.
type tempAlarmSearchResult  struct {
    Component *string  `json:"component,omitempty"`
    End       *int     `json:"end"`
    Limit     *int     `json:"limit"`
    Next      *string  `json:"next,omitempty"`
    Page      *int     `json:"page,omitempty"`
    Results   *[]Alarm `json:"results"`
    Start     *int     `json:"start"`
    Total     *int     `json:"total"`
}

func (a *tempAlarmSearchResult) validate() error {
    var errs []string
    if a.End == nil {
        errs = append(errs, "required field `end` is missing for type `alarm_search_result`")
    }
    if a.Limit == nil {
        errs = append(errs, "required field `limit` is missing for type `alarm_search_result`")
    }
    if a.Results == nil {
        errs = append(errs, "required field `results` is missing for type `alarm_search_result`")
    }
    if a.Start == nil {
        errs = append(errs, "required field `start` is missing for type `alarm_search_result`")
    }
    if a.Total == nil {
        errs = append(errs, "required field `total` is missing for type `alarm_search_result`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
