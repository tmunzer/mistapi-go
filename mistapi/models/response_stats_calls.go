package models

import (
    "encoding/json"
)

// ResponseStatsCalls represents a ResponseStatsCalls struct.
type ResponseStatsCalls struct {
    End                  *float64               `json:"end,omitempty"`
    Limit                *int                   `json:"limit,omitempty"`
    Next                 *string                `json:"next,omitempty"`
    Results              []StatsCall            `json:"results,omitempty"`
    Start                *float64               `json:"start,omitempty"`
    Total                *int                   `json:"total,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseStatsCalls.
// It customizes the JSON marshaling process for ResponseStatsCalls objects.
func (r ResponseStatsCalls) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "end", "limit", "next", "results", "start", "total"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseStatsCalls object to a map representation for JSON marshaling.
func (r ResponseStatsCalls) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
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

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseStatsCalls.
// It customizes the JSON unmarshaling process for ResponseStatsCalls objects.
func (r *ResponseStatsCalls) UnmarshalJSON(input []byte) error {
    var temp tempResponseStatsCalls
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "end", "limit", "next", "results", "start", "total")
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

// tempResponseStatsCalls is a temporary struct used for validating the fields of ResponseStatsCalls.
type tempResponseStatsCalls  struct {
    End     *float64    `json:"end,omitempty"`
    Limit   *int        `json:"limit,omitempty"`
    Next    *string     `json:"next,omitempty"`
    Results []StatsCall `json:"results,omitempty"`
    Start   *float64    `json:"start,omitempty"`
    Total   *int        `json:"total,omitempty"`
}
