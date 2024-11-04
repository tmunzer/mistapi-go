package models

import (
    "encoding/json"
)

// SearchWanUsage represents a SearchWanUsage struct.
type SearchWanUsage struct {
    Results              []WanUsages    `json:"results,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SearchWanUsage.
// It customizes the JSON marshaling process for SearchWanUsage objects.
func (s SearchWanUsage) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SearchWanUsage object to a map representation for JSON marshaling.
func (s SearchWanUsage) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Results != nil {
        structMap["results"] = s.Results
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SearchWanUsage.
// It customizes the JSON unmarshaling process for SearchWanUsage objects.
func (s *SearchWanUsage) UnmarshalJSON(input []byte) error {
    var temp tempSearchWanUsage
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "results")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.Results = temp.Results
    return nil
}

// tempSearchWanUsage is a temporary struct used for validating the fields of SearchWanUsage.
type tempSearchWanUsage  struct {
    Results []WanUsages `json:"results,omitempty"`
}
