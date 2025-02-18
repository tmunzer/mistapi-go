package models

import (
    "encoding/json"
    "fmt"
)

// SearchWanUsage represents a SearchWanUsage struct.
type SearchWanUsage struct {
    Results              []WanUsages            `json:"results,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SearchWanUsage,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SearchWanUsage) String() string {
    return fmt.Sprintf(
    	"SearchWanUsage[Results=%v, AdditionalProperties=%v]",
    	s.Results, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SearchWanUsage.
// It customizes the JSON marshaling process for SearchWanUsage objects.
func (s SearchWanUsage) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "results"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SearchWanUsage object to a map representation for JSON marshaling.
func (s SearchWanUsage) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
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
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "results")
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
