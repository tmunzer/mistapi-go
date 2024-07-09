package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ResponseSwitchPortSearch represents a ResponseSwitchPortSearch struct.
type ResponseSwitchPortSearch struct {
    End                  int               `json:"end"`
    Limit                int               `json:"limit"`
    Results              []SwitchPortStats `json:"results"`
    Start                int               `json:"start"`
    Total                int               `json:"total"`
    AdditionalProperties map[string]any    `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseSwitchPortSearch.
// It customizes the JSON marshaling process for ResponseSwitchPortSearch objects.
func (r ResponseSwitchPortSearch) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseSwitchPortSearch object to a map representation for JSON marshaling.
func (r ResponseSwitchPortSearch) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    structMap["end"] = r.End
    structMap["limit"] = r.Limit
    structMap["results"] = r.Results
    structMap["start"] = r.Start
    structMap["total"] = r.Total
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseSwitchPortSearch.
// It customizes the JSON unmarshaling process for ResponseSwitchPortSearch objects.
func (r *ResponseSwitchPortSearch) UnmarshalJSON(input []byte) error {
    var temp responseSwitchPortSearch
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "end", "limit", "results", "start", "total")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.End = *temp.End
    r.Limit = *temp.Limit
    r.Results = *temp.Results
    r.Start = *temp.Start
    r.Total = *temp.Total
    return nil
}

// responseSwitchPortSearch is a temporary struct used for validating the fields of ResponseSwitchPortSearch.
type responseSwitchPortSearch  struct {
    End     *int               `json:"end"`
    Limit   *int               `json:"limit"`
    Results *[]SwitchPortStats `json:"results"`
    Start   *int               `json:"start"`
    Total   *int               `json:"total"`
}

func (r *responseSwitchPortSearch) validate() error {
    var errs []string
    if r.End == nil {
        errs = append(errs, "required field `end` is missing for type `Response_Switch_Port_Search`")
    }
    if r.Limit == nil {
        errs = append(errs, "required field `limit` is missing for type `Response_Switch_Port_Search`")
    }
    if r.Results == nil {
        errs = append(errs, "required field `results` is missing for type `Response_Switch_Port_Search`")
    }
    if r.Start == nil {
        errs = append(errs, "required field `start` is missing for type `Response_Switch_Port_Search`")
    }
    if r.Total == nil {
        errs = append(errs, "required field `total` is missing for type `Response_Switch_Port_Search`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
