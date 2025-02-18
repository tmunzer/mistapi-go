package models

import (
    "encoding/json"
    "fmt"
)

// ResponseEventsPathSearch represents a ResponseEventsPathSearch struct.
type ResponseEventsPathSearch struct {
    End                  *int                   `json:"end,omitempty"`
    Limit                *int                   `json:"limit,omitempty"`
    Results              []ServicePathEvent     `json:"results,omitempty"`
    Start                *int                   `json:"start,omitempty"`
    Total                *int                   `json:"total,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseEventsPathSearch,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseEventsPathSearch) String() string {
    return fmt.Sprintf(
    	"ResponseEventsPathSearch[End=%v, Limit=%v, Results=%v, Start=%v, Total=%v, AdditionalProperties=%v]",
    	r.End, r.Limit, r.Results, r.Start, r.Total, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseEventsPathSearch.
// It customizes the JSON marshaling process for ResponseEventsPathSearch objects.
func (r ResponseEventsPathSearch) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "end", "limit", "results", "start", "total"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseEventsPathSearch object to a map representation for JSON marshaling.
func (r ResponseEventsPathSearch) toMap() map[string]any {
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

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseEventsPathSearch.
// It customizes the JSON unmarshaling process for ResponseEventsPathSearch objects.
func (r *ResponseEventsPathSearch) UnmarshalJSON(input []byte) error {
    var temp tempResponseEventsPathSearch
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

// tempResponseEventsPathSearch is a temporary struct used for validating the fields of ResponseEventsPathSearch.
type tempResponseEventsPathSearch  struct {
    End     *int               `json:"end,omitempty"`
    Limit   *int               `json:"limit,omitempty"`
    Results []ServicePathEvent `json:"results,omitempty"`
    Start   *int               `json:"start,omitempty"`
    Total   *int               `json:"total,omitempty"`
}
