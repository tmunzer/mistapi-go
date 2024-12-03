package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ResponseSsoFailureSearch represents a ResponseSsoFailureSearch struct.
type ResponseSsoFailureSearch struct {
    Results              []ResponseSsoFailureSearchItem `json:"results"`
    AdditionalProperties map[string]interface{}         `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseSsoFailureSearch.
// It customizes the JSON marshaling process for ResponseSsoFailureSearch objects.
func (r ResponseSsoFailureSearch) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "results"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseSsoFailureSearch object to a map representation for JSON marshaling.
func (r ResponseSsoFailureSearch) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    structMap["results"] = r.Results
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseSsoFailureSearch.
// It customizes the JSON unmarshaling process for ResponseSsoFailureSearch objects.
func (r *ResponseSsoFailureSearch) UnmarshalJSON(input []byte) error {
    var temp tempResponseSsoFailureSearch
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "results")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.Results = *temp.Results
    return nil
}

// tempResponseSsoFailureSearch is a temporary struct used for validating the fields of ResponseSsoFailureSearch.
type tempResponseSsoFailureSearch  struct {
    Results *[]ResponseSsoFailureSearchItem `json:"results"`
}

func (r *tempResponseSsoFailureSearch) validate() error {
    var errs []string
    if r.Results == nil {
        errs = append(errs, "required field `results` is missing for type `response_sso_failure_search`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
