package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// ResponseSwitchPortSearch represents a ResponseSwitchPortSearch struct.
type ResponseSwitchPortSearch struct {
    End                  int                    `json:"end"`
    Limit                int                    `json:"limit"`
    Results              []StatsSwitchPort      `json:"results"`
    Start                int                    `json:"start"`
    Total                int                    `json:"total"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseSwitchPortSearch,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseSwitchPortSearch) String() string {
    return fmt.Sprintf(
    	"ResponseSwitchPortSearch[End=%v, Limit=%v, Results=%v, Start=%v, Total=%v, AdditionalProperties=%v]",
    	r.End, r.Limit, r.Results, r.Start, r.Total, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseSwitchPortSearch.
// It customizes the JSON marshaling process for ResponseSwitchPortSearch objects.
func (r ResponseSwitchPortSearch) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "end", "limit", "results", "start", "total"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseSwitchPortSearch object to a map representation for JSON marshaling.
func (r ResponseSwitchPortSearch) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
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
    var temp tempResponseSwitchPortSearch
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "end", "limit", "results", "start", "total")
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

// tempResponseSwitchPortSearch is a temporary struct used for validating the fields of ResponseSwitchPortSearch.
type tempResponseSwitchPortSearch  struct {
    End     *int               `json:"end"`
    Limit   *int               `json:"limit"`
    Results *[]StatsSwitchPort `json:"results"`
    Start   *int               `json:"start"`
    Total   *int               `json:"total"`
}

func (r *tempResponseSwitchPortSearch) validate() error {
    var errs []string
    if r.End == nil {
        errs = append(errs, "required field `end` is missing for type `response_switch_port_search`")
    }
    if r.Limit == nil {
        errs = append(errs, "required field `limit` is missing for type `response_switch_port_search`")
    }
    if r.Results == nil {
        errs = append(errs, "required field `results` is missing for type `response_switch_port_search`")
    }
    if r.Start == nil {
        errs = append(errs, "required field `start` is missing for type `response_switch_port_search`")
    }
    if r.Total == nil {
        errs = append(errs, "required field `total` is missing for type `response_switch_port_search`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
