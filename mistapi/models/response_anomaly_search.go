// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// ResponseAnomalySearch represents a ResponseAnomalySearch struct.
type ResponseAnomalySearch struct {
    End                  int                    `json:"end"`
    Limit                int                    `json:"limit"`
    Page                 int                    `json:"page"`
    Results              []Anomaly              `json:"results"`
    Start                int                    `json:"start"`
    Total                *int                   `json:"total,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseAnomalySearch,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseAnomalySearch) String() string {
    return fmt.Sprintf(
    	"ResponseAnomalySearch[End=%v, Limit=%v, Page=%v, Results=%v, Start=%v, Total=%v, AdditionalProperties=%v]",
    	r.End, r.Limit, r.Page, r.Results, r.Start, r.Total, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseAnomalySearch.
// It customizes the JSON marshaling process for ResponseAnomalySearch objects.
func (r ResponseAnomalySearch) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "end", "limit", "page", "results", "start", "total"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseAnomalySearch object to a map representation for JSON marshaling.
func (r ResponseAnomalySearch) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    structMap["end"] = r.End
    structMap["limit"] = r.Limit
    structMap["page"] = r.Page
    structMap["results"] = r.Results
    structMap["start"] = r.Start
    if r.Total != nil {
        structMap["total"] = r.Total
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseAnomalySearch.
// It customizes the JSON unmarshaling process for ResponseAnomalySearch objects.
func (r *ResponseAnomalySearch) UnmarshalJSON(input []byte) error {
    var temp tempResponseAnomalySearch
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "end", "limit", "page", "results", "start", "total")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.End = *temp.End
    r.Limit = *temp.Limit
    r.Page = *temp.Page
    r.Results = *temp.Results
    r.Start = *temp.Start
    r.Total = temp.Total
    return nil
}

// tempResponseAnomalySearch is a temporary struct used for validating the fields of ResponseAnomalySearch.
type tempResponseAnomalySearch  struct {
    End     *int       `json:"end"`
    Limit   *int       `json:"limit"`
    Page    *int       `json:"page"`
    Results *[]Anomaly `json:"results"`
    Start   *int       `json:"start"`
    Total   *int       `json:"total,omitempty"`
}

func (r *tempResponseAnomalySearch) validate() error {
    var errs []string
    if r.End == nil {
        errs = append(errs, "required field `end` is missing for type `response_anomaly_search`")
    }
    if r.Limit == nil {
        errs = append(errs, "required field `limit` is missing for type `response_anomaly_search`")
    }
    if r.Page == nil {
        errs = append(errs, "required field `page` is missing for type `response_anomaly_search`")
    }
    if r.Results == nil {
        errs = append(errs, "required field `results` is missing for type `response_anomaly_search`")
    }
    if r.Start == nil {
        errs = append(errs, "required field `start` is missing for type `response_anomaly_search`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
