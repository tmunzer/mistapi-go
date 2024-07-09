package models

import (
    "encoding/json"
)

// ResponseEventsOrgsSearch represents a ResponseEventsOrgsSearch struct.
type ResponseEventsOrgsSearch struct {
    End                  *int           `json:"end,omitempty"`
    Limit                *int           `json:"limit,omitempty"`
    Results              []OrgEvent     `json:"results,omitempty"`
    Start                *int           `json:"start,omitempty"`
    Total                *int           `json:"total,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseEventsOrgsSearch.
// It customizes the JSON marshaling process for ResponseEventsOrgsSearch objects.
func (r ResponseEventsOrgsSearch) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseEventsOrgsSearch object to a map representation for JSON marshaling.
func (r ResponseEventsOrgsSearch) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
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

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseEventsOrgsSearch.
// It customizes the JSON unmarshaling process for ResponseEventsOrgsSearch objects.
func (r *ResponseEventsOrgsSearch) UnmarshalJSON(input []byte) error {
    var temp responseEventsOrgsSearch
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "end", "limit", "results", "start", "total")
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

// responseEventsOrgsSearch is a temporary struct used for validating the fields of ResponseEventsOrgsSearch.
type responseEventsOrgsSearch  struct {
    End     *int       `json:"end,omitempty"`
    Limit   *int       `json:"limit,omitempty"`
    Results []OrgEvent `json:"results,omitempty"`
    Start   *int       `json:"start,omitempty"`
    Total   *int       `json:"total,omitempty"`
}
