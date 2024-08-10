package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ResponseLogSearch represents a ResponseLogSearch struct.
type ResponseLogSearch struct {
    End                  int                     `json:"end"`
    Limit                int                     `json:"limit"`
    Next                 *string                 `json:"next,omitempty"`
    Results              []ResponseLogSearchItem `json:"results"`
    Start                int                     `json:"start"`
    Total                int                     `json:"total"`
    AdditionalProperties map[string]any          `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseLogSearch.
// It customizes the JSON marshaling process for ResponseLogSearch objects.
func (r ResponseLogSearch) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseLogSearch object to a map representation for JSON marshaling.
func (r ResponseLogSearch) toMap() map[string]any {
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

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseLogSearch.
// It customizes the JSON unmarshaling process for ResponseLogSearch objects.
func (r *ResponseLogSearch) UnmarshalJSON(input []byte) error {
    var temp tempResponseLogSearch
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

// tempResponseLogSearch is a temporary struct used for validating the fields of ResponseLogSearch.
type tempResponseLogSearch  struct {
    End     *int                     `json:"end"`
    Limit   *int                     `json:"limit"`
    Next    *string                  `json:"next,omitempty"`
    Results *[]ResponseLogSearchItem `json:"results"`
    Start   *int                     `json:"start"`
    Total   *int                     `json:"total"`
}

func (r *tempResponseLogSearch) validate() error {
    var errs []string
    if r.End == nil {
        errs = append(errs, "required field `end` is missing for type `response_log_search`")
    }
    if r.Limit == nil {
        errs = append(errs, "required field `limit` is missing for type `response_log_search`")
    }
    if r.Results == nil {
        errs = append(errs, "required field `results` is missing for type `response_log_search`")
    }
    if r.Start == nil {
        errs = append(errs, "required field `start` is missing for type `response_log_search`")
    }
    if r.Total == nil {
        errs = append(errs, "required field `total` is missing for type `response_log_search`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
