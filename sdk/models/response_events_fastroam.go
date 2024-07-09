package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ResponseEventsFastroam represents a ResponseEventsFastroam struct.
type ResponseEventsFastroam struct {
    End                  int             `json:"end"`
    Limit                int             `json:"limit"`
    // the link to query next set of results. value is null if no next page exists.
    Next                 *string         `json:"next,omitempty"`
    Results              []EventFastroam `json:"results"`
    Start                int             `json:"start"`
    AdditionalProperties map[string]any  `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseEventsFastroam.
// It customizes the JSON marshaling process for ResponseEventsFastroam objects.
func (r ResponseEventsFastroam) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseEventsFastroam object to a map representation for JSON marshaling.
func (r ResponseEventsFastroam) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
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
    var temp responseEventsFastroam
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "end", "limit", "next", "results", "start")
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

// responseEventsFastroam is a temporary struct used for validating the fields of ResponseEventsFastroam.
type responseEventsFastroam  struct {
    End     *int             `json:"end"`
    Limit   *int             `json:"limit"`
    Next    *string          `json:"next,omitempty"`
    Results *[]EventFastroam `json:"results"`
    Start   *int             `json:"start"`
}

func (r *responseEventsFastroam) validate() error {
    var errs []string
    if r.End == nil {
        errs = append(errs, "required field `end` is missing for type `Response_Events_Fastroam`")
    }
    if r.Limit == nil {
        errs = append(errs, "required field `limit` is missing for type `Response_Events_Fastroam`")
    }
    if r.Results == nil {
        errs = append(errs, "required field `results` is missing for type `Response_Events_Fastroam`")
    }
    if r.Start == nil {
        errs = append(errs, "required field `start` is missing for type `Response_Events_Fastroam`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
