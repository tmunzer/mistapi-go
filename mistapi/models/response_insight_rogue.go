package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// ResponseInsightRogue represents a ResponseInsightRogue struct.
type ResponseInsightRogue struct {
    End                  int                    `json:"end"`
    Limit                int                    `json:"limit"`
    // Link to next set of results. If more results aren’t present, next is null.
    Next                 *string                `json:"next,omitempty"`
    Results              []InsightRogueAp       `json:"results"`
    Start                int                    `json:"start"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseInsightRogue,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseInsightRogue) String() string {
    return fmt.Sprintf(
    	"ResponseInsightRogue[End=%v, Limit=%v, Next=%v, Results=%v, Start=%v, AdditionalProperties=%v]",
    	r.End, r.Limit, r.Next, r.Results, r.Start, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseInsightRogue.
// It customizes the JSON marshaling process for ResponseInsightRogue objects.
func (r ResponseInsightRogue) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "end", "limit", "next", "results", "start"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseInsightRogue object to a map representation for JSON marshaling.
func (r ResponseInsightRogue) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    structMap["end"] = r.End
    structMap["limit"] = r.Limit
    if r.Next != nil {
        structMap["next"] = r.Next
    }
    structMap["results"] = r.Results
    structMap["start"] = r.Start
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseInsightRogue.
// It customizes the JSON unmarshaling process for ResponseInsightRogue objects.
func (r *ResponseInsightRogue) UnmarshalJSON(input []byte) error {
    var temp tempResponseInsightRogue
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "end", "limit", "next", "results", "start")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.End = *temp.End
    r.Limit = *temp.Limit
    r.Next = temp.Next
    r.Results = *temp.Results
    r.Start = *temp.Start
    return nil
}

// tempResponseInsightRogue is a temporary struct used for validating the fields of ResponseInsightRogue.
type tempResponseInsightRogue  struct {
    End     *int              `json:"end"`
    Limit   *int              `json:"limit"`
    Next    *string           `json:"next,omitempty"`
    Results *[]InsightRogueAp `json:"results"`
    Start   *int              `json:"start"`
}

func (r *tempResponseInsightRogue) validate() error {
    var errs []string
    if r.End == nil {
        errs = append(errs, "required field `end` is missing for type `response_insight_rogue`")
    }
    if r.Limit == nil {
        errs = append(errs, "required field `limit` is missing for type `response_insight_rogue`")
    }
    if r.Results == nil {
        errs = append(errs, "required field `results` is missing for type `response_insight_rogue`")
    }
    if r.Start == nil {
        errs = append(errs, "required field `start` is missing for type `response_insight_rogue`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
