package models

import (
    "encoding/json"
    "fmt"
)

// ResponseTroubleshoot represents a ResponseTroubleshoot struct.
type ResponseTroubleshoot struct {
    End                  *int                       `json:"end,omitempty"`
    Results              []ResponseTroubleshootItem `json:"results,omitempty"`
    Start                *int                       `json:"start,omitempty"`
    AdditionalProperties map[string]interface{}     `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseTroubleshoot,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseTroubleshoot) String() string {
    return fmt.Sprintf(
    	"ResponseTroubleshoot[End=%v, Results=%v, Start=%v, AdditionalProperties=%v]",
    	r.End, r.Results, r.Start, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseTroubleshoot.
// It customizes the JSON marshaling process for ResponseTroubleshoot objects.
func (r ResponseTroubleshoot) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "end", "results", "start"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseTroubleshoot object to a map representation for JSON marshaling.
func (r ResponseTroubleshoot) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    if r.End != nil {
        structMap["end"] = r.End
    }
    if r.Results != nil {
        structMap["results"] = r.Results
    }
    if r.Start != nil {
        structMap["start"] = r.Start
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseTroubleshoot.
// It customizes the JSON unmarshaling process for ResponseTroubleshoot objects.
func (r *ResponseTroubleshoot) UnmarshalJSON(input []byte) error {
    var temp tempResponseTroubleshoot
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "end", "results", "start")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.End = temp.End
    r.Results = temp.Results
    r.Start = temp.Start
    return nil
}

// tempResponseTroubleshoot is a temporary struct used for validating the fields of ResponseTroubleshoot.
type tempResponseTroubleshoot  struct {
    End     *int                       `json:"end,omitempty"`
    Results []ResponseTroubleshootItem `json:"results,omitempty"`
    Start   *int                       `json:"start,omitempty"`
}
