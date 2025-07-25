// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// ResponseRrmNeighbors represents a ResponseRrmNeighbors struct.
type ResponseRrmNeighbors struct {
    End                  int                    `json:"end"`
    Limit                int                    `json:"limit"`
    // Link to query next set of results. value is null if no next page exists.
    Next                 *string                `json:"next,omitempty"`
    Results              []RrmNeighbors         `json:"results"`
    Start                int                    `json:"start"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseRrmNeighbors,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseRrmNeighbors) String() string {
    return fmt.Sprintf(
    	"ResponseRrmNeighbors[End=%v, Limit=%v, Next=%v, Results=%v, Start=%v, AdditionalProperties=%v]",
    	r.End, r.Limit, r.Next, r.Results, r.Start, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseRrmNeighbors.
// It customizes the JSON marshaling process for ResponseRrmNeighbors objects.
func (r ResponseRrmNeighbors) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "end", "limit", "next", "results", "start"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseRrmNeighbors object to a map representation for JSON marshaling.
func (r ResponseRrmNeighbors) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    structMap["end"] = r.End
    structMap["limit"] = r.Limit
    if r.Next != nil {
        structMap["next"] = r.Next
    }
    structMap["results"] = r.Results
    structMap["start"] = r.Start
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseRrmNeighbors.
// It customizes the JSON unmarshaling process for ResponseRrmNeighbors objects.
func (r *ResponseRrmNeighbors) UnmarshalJSON(input []byte) error {
    var temp tempResponseRrmNeighbors
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "end", "limit", "next", "results", "start")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.End = *temp.End
    r.Limit = *temp.Limit
    r.Next = temp.Next
    r.Results = *temp.Results
    r.Start = *temp.Start
    return nil
}

// tempResponseRrmNeighbors is a temporary struct used for validating the fields of ResponseRrmNeighbors.
type tempResponseRrmNeighbors  struct {
    End     *int            `json:"end"`
    Limit   *int            `json:"limit"`
    Next    *string         `json:"next,omitempty"`
    Results *[]RrmNeighbors `json:"results"`
    Start   *int            `json:"start"`
}

func (r *tempResponseRrmNeighbors) validate() error {
    var errs []string
    if r.End == nil {
        errs = append(errs, "required field `end` is missing for type `response_rrm_neighbors`")
    }
    if r.Limit == nil {
        errs = append(errs, "required field `limit` is missing for type `response_rrm_neighbors`")
    }
    if r.Results == nil {
        errs = append(errs, "required field `results` is missing for type `response_rrm_neighbors`")
    }
    if r.Start == nil {
        errs = append(errs, "required field `start` is missing for type `response_rrm_neighbors`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
