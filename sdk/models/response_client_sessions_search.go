package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ResponseClientSessionsSearch represents a ResponseClientSessionsSearch struct.
type ResponseClientSessionsSearch struct {
    End                  int                                `json:"end"`
    Limit                int                                `json:"limit"`
    Next                 *string                            `json:"next,omitempty"`
    Results              []ResponseClientSessionsSearchItem `json:"results"`
    Start                int                                `json:"start"`
    Total                int                                `json:"total"`
    AdditionalProperties map[string]any                     `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseClientSessionsSearch.
// It customizes the JSON marshaling process for ResponseClientSessionsSearch objects.
func (r ResponseClientSessionsSearch) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseClientSessionsSearch object to a map representation for JSON marshaling.
func (r ResponseClientSessionsSearch) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
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

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseClientSessionsSearch.
// It customizes the JSON unmarshaling process for ResponseClientSessionsSearch objects.
func (r *ResponseClientSessionsSearch) UnmarshalJSON(input []byte) error {
    var temp responseClientSessionsSearch
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "end", "limit", "next", "results", "start", "total")
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

// responseClientSessionsSearch is a temporary struct used for validating the fields of ResponseClientSessionsSearch.
type responseClientSessionsSearch  struct {
    End     *int                                `json:"end"`
    Limit   *int                                `json:"limit"`
    Next    *string                             `json:"next,omitempty"`
    Results *[]ResponseClientSessionsSearchItem `json:"results"`
    Start   *int                                `json:"start"`
    Total   *int                                `json:"total"`
}

func (r *responseClientSessionsSearch) validate() error {
    var errs []string
    if r.End == nil {
        errs = append(errs, "required field `end` is missing for type `Response_Client_Sessions_Search`")
    }
    if r.Limit == nil {
        errs = append(errs, "required field `limit` is missing for type `Response_Client_Sessions_Search`")
    }
    if r.Results == nil {
        errs = append(errs, "required field `results` is missing for type `Response_Client_Sessions_Search`")
    }
    if r.Start == nil {
        errs = append(errs, "required field `start` is missing for type `Response_Client_Sessions_Search`")
    }
    if r.Total == nil {
        errs = append(errs, "required field `total` is missing for type `Response_Client_Sessions_Search`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
