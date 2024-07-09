package models

import (
    "encoding/json"
)

// ResponseZoneSearch represents a ResponseZoneSearch struct.
type ResponseZoneSearch struct {
    End                  *float64                 `json:"end,omitempty"`
    Limit                *int                     `json:"limit,omitempty"`
    Next                 *string                  `json:"next,omitempty"`
    Results              []ResponseZoneSearchItem `json:"results,omitempty"`
    Start                *float64                 `json:"start,omitempty"`
    Total                *int                     `json:"total,omitempty"`
    AdditionalProperties map[string]any           `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseZoneSearch.
// It customizes the JSON marshaling process for ResponseZoneSearch objects.
func (r ResponseZoneSearch) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseZoneSearch object to a map representation for JSON marshaling.
func (r ResponseZoneSearch) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    if r.End != nil {
        structMap["end"] = r.End
    }
    if r.Limit != nil {
        structMap["limit"] = r.Limit
    }
    if r.Next != nil {
        structMap["next"] = r.Next
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

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseZoneSearch.
// It customizes the JSON unmarshaling process for ResponseZoneSearch objects.
func (r *ResponseZoneSearch) UnmarshalJSON(input []byte) error {
    var temp responseZoneSearch
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "end", "limit", "next", "results", "start", "total")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.End = temp.End
    r.Limit = temp.Limit
    r.Next = temp.Next
    r.Results = temp.Results
    r.Start = temp.Start
    r.Total = temp.Total
    return nil
}

// responseZoneSearch is a temporary struct used for validating the fields of ResponseZoneSearch.
type responseZoneSearch  struct {
    End     *float64                 `json:"end,omitempty"`
    Limit   *int                     `json:"limit,omitempty"`
    Next    *string                  `json:"next,omitempty"`
    Results []ResponseZoneSearchItem `json:"results,omitempty"`
    Start   *float64                 `json:"start,omitempty"`
    Total   *int                     `json:"total,omitempty"`
}
