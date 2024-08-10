package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// SearchWirelssClientSession represents a SearchWirelssClientSession struct.
type SearchWirelssClientSession struct {
    End                  float64                `json:"end"`
    Limit                int                    `json:"limit"`
    Next                 *string                `json:"next,omitempty"`
    Results              []WirelssClientSession `json:"results"`
    Start                float64                `json:"start"`
    Total                int                    `json:"total"`
    AdditionalProperties map[string]any         `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SearchWirelssClientSession.
// It customizes the JSON marshaling process for SearchWirelssClientSession objects.
func (s SearchWirelssClientSession) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SearchWirelssClientSession object to a map representation for JSON marshaling.
func (s SearchWirelssClientSession) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    structMap["end"] = s.End
    structMap["limit"] = s.Limit
    if s.Next != nil {
        structMap["next"] = s.Next
    }
    structMap["results"] = s.Results
    structMap["start"] = s.Start
    structMap["total"] = s.Total
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SearchWirelssClientSession.
// It customizes the JSON unmarshaling process for SearchWirelssClientSession objects.
func (s *SearchWirelssClientSession) UnmarshalJSON(input []byte) error {
    var temp tempSearchWirelssClientSession
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "end", "limit", "next", "results", "start", "total")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.End = *temp.End
    s.Limit = *temp.Limit
    s.Next = temp.Next
    s.Results = *temp.Results
    s.Start = *temp.Start
    s.Total = *temp.Total
    return nil
}

// tempSearchWirelssClientSession is a temporary struct used for validating the fields of SearchWirelssClientSession.
type tempSearchWirelssClientSession  struct {
    End     *float64                `json:"end"`
    Limit   *int                    `json:"limit"`
    Next    *string                 `json:"next,omitempty"`
    Results *[]WirelssClientSession `json:"results"`
    Start   *float64                `json:"start"`
    Total   *int                    `json:"total"`
}

func (s *tempSearchWirelssClientSession) validate() error {
    var errs []string
    if s.End == nil {
        errs = append(errs, "required field `end` is missing for type `search_wirelss_client_session`")
    }
    if s.Limit == nil {
        errs = append(errs, "required field `limit` is missing for type `search_wirelss_client_session`")
    }
    if s.Results == nil {
        errs = append(errs, "required field `results` is missing for type `search_wirelss_client_session`")
    }
    if s.Start == nil {
        errs = append(errs, "required field `start` is missing for type `search_wirelss_client_session`")
    }
    if s.Total == nil {
        errs = append(errs, "required field `total` is missing for type `search_wirelss_client_session`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
