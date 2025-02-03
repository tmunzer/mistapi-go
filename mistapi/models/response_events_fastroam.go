package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// ResponseEventsFastroam represents a ResponseEventsFastroam struct.
type ResponseEventsFastroam struct {
    End                  int                    `json:"end"`
    Limit                int                    `json:"limit"`
    // Link to query next set of results. value is null if no next page exists.
    Next                 *string                `json:"next,omitempty"`
    Results              []EventFastroam        `json:"results"`
    Start                int                    `json:"start"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseEventsFastroam,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseEventsFastroam) String() string {
    return fmt.Sprintf(
    	"ResponseEventsFastroam[End=%v, Limit=%v, Next=%v, Results=%v, Start=%v, AdditionalProperties=%v]",
    	r.End, r.Limit, r.Next, r.Results, r.Start, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseEventsFastroam.
// It customizes the JSON marshaling process for ResponseEventsFastroam objects.
func (r ResponseEventsFastroam) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "end", "limit", "next", "results", "start"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseEventsFastroam object to a map representation for JSON marshaling.
func (r ResponseEventsFastroam) toMap() map[string]any {
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

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseEventsFastroam.
// It customizes the JSON unmarshaling process for ResponseEventsFastroam objects.
func (r *ResponseEventsFastroam) UnmarshalJSON(input []byte) error {
    var temp tempResponseEventsFastroam
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

// tempResponseEventsFastroam is a temporary struct used for validating the fields of ResponseEventsFastroam.
type tempResponseEventsFastroam  struct {
    End     *int             `json:"end"`
    Limit   *int             `json:"limit"`
    Next    *string          `json:"next,omitempty"`
    Results *[]EventFastroam `json:"results"`
    Start   *int             `json:"start"`
}

func (r *tempResponseEventsFastroam) validate() error {
    var errs []string
    if r.End == nil {
        errs = append(errs, "required field `end` is missing for type `response_events_fastroam`")
    }
    if r.Limit == nil {
        errs = append(errs, "required field `limit` is missing for type `response_events_fastroam`")
    }
    if r.Results == nil {
        errs = append(errs, "required field `results` is missing for type `response_events_fastroam`")
    }
    if r.Start == nil {
        errs = append(errs, "required field `start` is missing for type `response_events_fastroam`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
