package models

import (
    "encoding/json"
    "fmt"
)

// ResponseClientNacSearch represents a ResponseClientNacSearch struct.
type ResponseClientNacSearch struct {
    End                  *int                   `json:"end,omitempty"`
    Limit                *int                   `json:"limit,omitempty"`
    Results              []ClientNac            `json:"results,omitempty"`
    Start                *int                   `json:"start,omitempty"`
    Total                *int                   `json:"total,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseClientNacSearch,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseClientNacSearch) String() string {
    return fmt.Sprintf(
    	"ResponseClientNacSearch[End=%v, Limit=%v, Results=%v, Start=%v, Total=%v, AdditionalProperties=%v]",
    	r.End, r.Limit, r.Results, r.Start, r.Total, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseClientNacSearch.
// It customizes the JSON marshaling process for ResponseClientNacSearch objects.
func (r ResponseClientNacSearch) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "end", "limit", "results", "start", "total"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseClientNacSearch object to a map representation for JSON marshaling.
func (r ResponseClientNacSearch) toMap() map[string]any {
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

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseClientNacSearch.
// It customizes the JSON unmarshaling process for ResponseClientNacSearch objects.
func (r *ResponseClientNacSearch) UnmarshalJSON(input []byte) error {
    var temp tempResponseClientNacSearch
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

// tempResponseClientNacSearch is a temporary struct used for validating the fields of ResponseClientNacSearch.
type tempResponseClientNacSearch  struct {
    End     *int        `json:"end,omitempty"`
    Limit   *int        `json:"limit,omitempty"`
    Results []ClientNac `json:"results,omitempty"`
    Start   *int        `json:"start,omitempty"`
    Total   *int        `json:"total,omitempty"`
}
