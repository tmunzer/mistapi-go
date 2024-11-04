package models

import (
    "encoding/json"
)

// ResponseCountMarvisActionsResult represents a ResponseCountMarvisActionsResult struct.
type ResponseCountMarvisActionsResult struct {
    Count                *int           `json:"count,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseCountMarvisActionsResult.
// It customizes the JSON marshaling process for ResponseCountMarvisActionsResult objects.
func (r ResponseCountMarvisActionsResult) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseCountMarvisActionsResult object to a map representation for JSON marshaling.
func (r ResponseCountMarvisActionsResult) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    if r.Count != nil {
        structMap["count"] = r.Count
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseCountMarvisActionsResult.
// It customizes the JSON unmarshaling process for ResponseCountMarvisActionsResult objects.
func (r *ResponseCountMarvisActionsResult) UnmarshalJSON(input []byte) error {
    var temp tempResponseCountMarvisActionsResult
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "count")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.Count = temp.Count
    return nil
}

// tempResponseCountMarvisActionsResult is a temporary struct used for validating the fields of ResponseCountMarvisActionsResult.
type tempResponseCountMarvisActionsResult  struct {
    Count *int `json:"count,omitempty"`
}
