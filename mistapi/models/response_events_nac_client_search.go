package models

import (
    "encoding/json"
)

// ResponseEventsNacClientSearch represents a ResponseEventsNacClientSearch struct.
type ResponseEventsNacClientSearch struct {
    End                  *int                   `json:"end,omitempty"`
    Limit                *int                   `json:"limit,omitempty"`
    Results              []EventNacClient       `json:"results,omitempty"`
    Start                *int                   `json:"start,omitempty"`
    Total                *int                   `json:"total,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseEventsNacClientSearch.
// It customizes the JSON marshaling process for ResponseEventsNacClientSearch objects.
func (r ResponseEventsNacClientSearch) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "end", "limit", "results", "start", "total"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseEventsNacClientSearch object to a map representation for JSON marshaling.
func (r ResponseEventsNacClientSearch) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    if r.End != nil {
        structMap["end"] = r.End
    }
    if r.Limit != nil {
        structMap["limit"] = r.Limit
    }
    if r.Results != nil {
        structMap["results"] = r.Results
    }
    if r.Start != nil {
        structMap["start"] = r.Start
    }
    if r.Total != nil {
        structMap["total"] = r.Total
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseEventsNacClientSearch.
// It customizes the JSON unmarshaling process for ResponseEventsNacClientSearch objects.
func (r *ResponseEventsNacClientSearch) UnmarshalJSON(input []byte) error {
    var temp tempResponseEventsNacClientSearch
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "end", "limit", "results", "start", "total")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.End = temp.End
    r.Limit = temp.Limit
    r.Results = temp.Results
    r.Start = temp.Start
    r.Total = temp.Total
    return nil
}

// tempResponseEventsNacClientSearch is a temporary struct used for validating the fields of ResponseEventsNacClientSearch.
type tempResponseEventsNacClientSearch  struct {
    End     *int             `json:"end,omitempty"`
    Limit   *int             `json:"limit,omitempty"`
    Results []EventNacClient `json:"results,omitempty"`
    Start   *int             `json:"start,omitempty"`
    Total   *int             `json:"total,omitempty"`
}
