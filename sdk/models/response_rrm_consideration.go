package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ResponseRrmConsideration represents a ResponseRrmConsideration struct.
type ResponseRrmConsideration struct {
    Results              []RrmConsideration `json:"results"`
    AdditionalProperties map[string]any     `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseRrmConsideration.
// It customizes the JSON marshaling process for ResponseRrmConsideration objects.
func (r ResponseRrmConsideration) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseRrmConsideration object to a map representation for JSON marshaling.
func (r ResponseRrmConsideration) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    structMap["results"] = r.Results
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseRrmConsideration.
// It customizes the JSON unmarshaling process for ResponseRrmConsideration objects.
func (r *ResponseRrmConsideration) UnmarshalJSON(input []byte) error {
    var temp responseRrmConsideration
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "results")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.Results = *temp.Results
    return nil
}

// responseRrmConsideration is a temporary struct used for validating the fields of ResponseRrmConsideration.
type responseRrmConsideration  struct {
    Results *[]RrmConsideration `json:"results"`
}

func (r *responseRrmConsideration) validate() error {
    var errs []string
    if r.Results == nil {
        errs = append(errs, "required field `results` is missing for type `Response_Rrm_Consideration`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
