package models

import (
    "encoding/json"
)

// ResponsePskPortalLogsSearch represents a ResponsePskPortalLogsSearch struct.
type ResponsePskPortalLogsSearch struct {
    End                  *int                              `json:"end,omitempty"`
    Limit                *int                              `json:"limit,omitempty"`
    Results              []ResponsePskPortalLogsSearchItem `json:"results,omitempty"`
    Start                *int                              `json:"start,omitempty"`
    Total                *int                              `json:"total,omitempty"`
    AdditionalProperties map[string]any                    `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponsePskPortalLogsSearch.
// It customizes the JSON marshaling process for ResponsePskPortalLogsSearch objects.
func (r ResponsePskPortalLogsSearch) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the ResponsePskPortalLogsSearch object to a map representation for JSON marshaling.
func (r ResponsePskPortalLogsSearch) toMap() map[string]any {
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

// UnmarshalJSON implements the json.Unmarshaler interface for ResponsePskPortalLogsSearch.
// It customizes the JSON unmarshaling process for ResponsePskPortalLogsSearch objects.
func (r *ResponsePskPortalLogsSearch) UnmarshalJSON(input []byte) error {
    var temp tempResponsePskPortalLogsSearch
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

// tempResponsePskPortalLogsSearch is a temporary struct used for validating the fields of ResponsePskPortalLogsSearch.
type tempResponsePskPortalLogsSearch  struct {
    End     *int                              `json:"end,omitempty"`
    Limit   *int                              `json:"limit,omitempty"`
    Results []ResponsePskPortalLogsSearchItem `json:"results,omitempty"`
    Start   *int                              `json:"start,omitempty"`
    Total   *int                              `json:"total,omitempty"`
}
