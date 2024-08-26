package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ResponseEventsSearch represents a ResponseEventsSearch struct.
type ResponseEventsSearch struct {
    End                  int            `json:"end"`
    Limit                int            `json:"limit"`
    Next                 *string        `json:"next,omitempty"`
    Results              []EventsClient `json:"results"`
    Start                int            `json:"start"`
    Total                int            `json:"total"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseEventsSearch.
// It customizes the JSON marshaling process for ResponseEventsSearch objects.
func (r ResponseEventsSearch) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseEventsSearch object to a map representation for JSON marshaling.
func (r ResponseEventsSearch) toMap() map[string]any {
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

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseEventsSearch.
// It customizes the JSON unmarshaling process for ResponseEventsSearch objects.
func (r *ResponseEventsSearch) UnmarshalJSON(input []byte) error {
    var temp tempResponseEventsSearch
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

// tempResponseEventsSearch is a temporary struct used for validating the fields of ResponseEventsSearch.
type tempResponseEventsSearch  struct {
    End     *int            `json:"end"`
    Limit   *int            `json:"limit"`
    Next    *string         `json:"next,omitempty"`
    Results *[]EventsClient `json:"results"`
    Start   *int            `json:"start"`
    Total   *int            `json:"total"`
}

func (r *tempResponseEventsSearch) validate() error {
    var errs []string
    if r.End == nil {
        errs = append(errs, "required field `end` is missing for type `response_events_search`")
    }
    if r.Limit == nil {
        errs = append(errs, "required field `limit` is missing for type `response_events_search`")
    }
    if r.Results == nil {
        errs = append(errs, "required field `results` is missing for type `response_events_search`")
    }
    if r.Start == nil {
        errs = append(errs, "required field `start` is missing for type `response_events_search`")
    }
    if r.Total == nil {
        errs = append(errs, "required field `total` is missing for type `response_events_search`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
