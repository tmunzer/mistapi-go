// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// JsInventorySearch represents a JsInventorySearch struct.
type JsInventorySearch struct {
    // Offset to end at
    End                  *int                   `json:"end,omitempty"`
    // Number of results to return
    Limit                *int                   `json:"limit,omitempty"`
    Results              []JsInventoryItem      `json:"results,omitempty"`
    // Offset to start from
    Start                *int                   `json:"start,omitempty"`
    // Total number of results
    Total                *int                   `json:"total,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for JsInventorySearch,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (j JsInventorySearch) String() string {
    return fmt.Sprintf(
    	"JsInventorySearch[End=%v, Limit=%v, Results=%v, Start=%v, Total=%v, AdditionalProperties=%v]",
    	j.End, j.Limit, j.Results, j.Start, j.Total, j.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for JsInventorySearch.
// It customizes the JSON marshaling process for JsInventorySearch objects.
func (j JsInventorySearch) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(j.AdditionalProperties,
        "end", "limit", "results", "start", "total"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(j.toMap())
}

// toMap converts the JsInventorySearch object to a map representation for JSON marshaling.
func (j JsInventorySearch) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, j.AdditionalProperties)
    if j.End != nil {
        structMap["end"] = j.End
    }
    if j.Limit != nil {
        structMap["limit"] = j.Limit
    }
    if j.Results != nil {
        structMap["results"] = j.Results
    }
    if j.Start != nil {
        structMap["start"] = j.Start
    }
    if j.Total != nil {
        structMap["total"] = j.Total
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for JsInventorySearch.
// It customizes the JSON unmarshaling process for JsInventorySearch objects.
func (j *JsInventorySearch) UnmarshalJSON(input []byte) error {
    var temp tempJsInventorySearch
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "end", "limit", "results", "start", "total")
    if err != nil {
    	return err
    }
    j.AdditionalProperties = additionalProperties
    
    j.End = temp.End
    j.Limit = temp.Limit
    j.Results = temp.Results
    j.Start = temp.Start
    j.Total = temp.Total
    return nil
}

// tempJsInventorySearch is a temporary struct used for validating the fields of JsInventorySearch.
type tempJsInventorySearch  struct {
    End     *int              `json:"end,omitempty"`
    Limit   *int              `json:"limit,omitempty"`
    Results []JsInventoryItem `json:"results,omitempty"`
    Start   *int              `json:"start,omitempty"`
    Total   *int              `json:"total,omitempty"`
}
