package models

import (
    "encoding/json"
)

// ResponseCountMarvisActions represents a ResponseCountMarvisActions struct.
type ResponseCountMarvisActions struct {
    Distinct             *string                            `json:"distinct,omitempty"`
    Limit                *int                               `json:"limit,omitempty"`
    Results              []ResponseCountMarvisActionsResult `json:"results,omitempty"`
    Total                *int                               `json:"total,omitempty"`
    AdditionalProperties map[string]any                     `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseCountMarvisActions.
// It customizes the JSON marshaling process for ResponseCountMarvisActions objects.
func (r ResponseCountMarvisActions) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseCountMarvisActions object to a map representation for JSON marshaling.
func (r ResponseCountMarvisActions) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    if r.Distinct != nil {
        structMap["distinct"] = r.Distinct
    }
    if r.Limit != nil {
        structMap["limit"] = r.Limit
    }
    if r.Results != nil {
        structMap["results"] = r.Results
    }
    if r.Total != nil {
        structMap["total"] = r.Total
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseCountMarvisActions.
// It customizes the JSON unmarshaling process for ResponseCountMarvisActions objects.
func (r *ResponseCountMarvisActions) UnmarshalJSON(input []byte) error {
    var temp responseCountMarvisActions
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "distinct", "limit", "results", "total")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.Distinct = temp.Distinct
    r.Limit = temp.Limit
    r.Results = temp.Results
    r.Total = temp.Total
    return nil
}

// responseCountMarvisActions is a temporary struct used for validating the fields of ResponseCountMarvisActions.
type responseCountMarvisActions  struct {
    Distinct *string                            `json:"distinct,omitempty"`
    Limit    *int                               `json:"limit,omitempty"`
    Results  []ResponseCountMarvisActionsResult `json:"results,omitempty"`
    Total    *int                               `json:"total,omitempty"`
}
