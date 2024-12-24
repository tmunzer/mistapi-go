package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// ResponseEventsRrm represents a ResponseEventsRrm struct.
type ResponseEventsRrm struct {
    End                  int                    `json:"end"`
    Limit                int                    `json:"limit"`
    // the link to query next set of results. value is null if no next page exists.
    Next                 *string                `json:"next,omitempty"`
    Results              []RrmEvent             `json:"results"`
    Start                int                    `json:"start"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseEventsRrm,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseEventsRrm) String() string {
    return fmt.Sprintf(
    	"ResponseEventsRrm[End=%v, Limit=%v, Next=%v, Results=%v, Start=%v, AdditionalProperties=%v]",
    	r.End, r.Limit, r.Next, r.Results, r.Start, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseEventsRrm.
// It customizes the JSON marshaling process for ResponseEventsRrm objects.
func (r ResponseEventsRrm) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "end", "limit", "next", "results", "start"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseEventsRrm object to a map representation for JSON marshaling.
func (r ResponseEventsRrm) toMap() map[string]any {
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

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseEventsRrm.
// It customizes the JSON unmarshaling process for ResponseEventsRrm objects.
func (r *ResponseEventsRrm) UnmarshalJSON(input []byte) error {
    var temp tempResponseEventsRrm
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

// tempResponseEventsRrm is a temporary struct used for validating the fields of ResponseEventsRrm.
type tempResponseEventsRrm  struct {
    End     *int        `json:"end"`
    Limit   *int        `json:"limit"`
    Next    *string     `json:"next,omitempty"`
    Results *[]RrmEvent `json:"results"`
    Start   *int        `json:"start"`
}

func (r *tempResponseEventsRrm) validate() error {
    var errs []string
    if r.End == nil {
        errs = append(errs, "required field `end` is missing for type `response_events_rrm`")
    }
    if r.Limit == nil {
        errs = append(errs, "required field `limit` is missing for type `response_events_rrm`")
    }
    if r.Results == nil {
        errs = append(errs, "required field `results` is missing for type `response_events_rrm`")
    }
    if r.Start == nil {
        errs = append(errs, "required field `start` is missing for type `response_events_rrm`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
