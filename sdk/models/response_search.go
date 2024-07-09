package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ResponseSearch represents a ResponseSearch struct.
type ResponseSearch struct {
    Limit                int                  `json:"limit"`
    Next                 *string              `json:"next,omitempty"`
    Page                 int                  `json:"page"`
    Results              []ResponseSearchItem `json:"results"`
    Total                int                  `json:"total"`
    AdditionalProperties map[string]any       `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseSearch.
// It customizes the JSON marshaling process for ResponseSearch objects.
func (r ResponseSearch) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseSearch object to a map representation for JSON marshaling.
func (r ResponseSearch) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    structMap["limit"] = r.Limit
    if r.Next != nil {
        structMap["next"] = r.Next
    }
    structMap["page"] = r.Page
    structMap["results"] = r.Results
    structMap["total"] = r.Total
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseSearch.
// It customizes the JSON unmarshaling process for ResponseSearch objects.
func (r *ResponseSearch) UnmarshalJSON(input []byte) error {
    var temp responseSearch
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "limit", "next", "page", "results", "total")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.Limit = *temp.Limit
    r.Next = temp.Next
    r.Page = *temp.Page
    r.Results = *temp.Results
    r.Total = *temp.Total
    return nil
}

// responseSearch is a temporary struct used for validating the fields of ResponseSearch.
type responseSearch  struct {
    Limit   *int                  `json:"limit"`
    Next    *string               `json:"next,omitempty"`
    Page    *int                  `json:"page"`
    Results *[]ResponseSearchItem `json:"results"`
    Total   *int                  `json:"total"`
}

func (r *responseSearch) validate() error {
    var errs []string
    if r.Limit == nil {
        errs = append(errs, "required field `limit` is missing for type `Response_Search`")
    }
    if r.Page == nil {
        errs = append(errs, "required field `page` is missing for type `Response_Search`")
    }
    if r.Results == nil {
        errs = append(errs, "required field `results` is missing for type `Response_Search`")
    }
    if r.Total == nil {
        errs = append(errs, "required field `total` is missing for type `Response_Search`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
