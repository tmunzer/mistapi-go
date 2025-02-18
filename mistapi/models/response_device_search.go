package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// ResponseDeviceSearch represents a ResponseDeviceSearch struct.
type ResponseDeviceSearch struct {
    End                  int                                `json:"end"`
    Limit                int                                `json:"limit"`
    Next                 *string                            `json:"next,omitempty"`
    Results              []ResponseDeviceSearchResultsItems `json:"results"`
    Start                int                                `json:"start"`
    Total                int                                `json:"total"`
    AdditionalProperties map[string]interface{}             `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseDeviceSearch,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseDeviceSearch) String() string {
    return fmt.Sprintf(
    	"ResponseDeviceSearch[End=%v, Limit=%v, Next=%v, Results=%v, Start=%v, Total=%v, AdditionalProperties=%v]",
    	r.End, r.Limit, r.Next, r.Results, r.Start, r.Total, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseDeviceSearch.
// It customizes the JSON marshaling process for ResponseDeviceSearch objects.
func (r ResponseDeviceSearch) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "end", "limit", "next", "results", "start", "total"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseDeviceSearch object to a map representation for JSON marshaling.
func (r ResponseDeviceSearch) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    structMap["end"] = r.End
    structMap["limit"] = r.Limit
    if r.Next != nil {
        structMap["next"] = r.Next
    }
    structMap["results"] = r.Results
    structMap["start"] = r.Start
    structMap["total"] = r.Total
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseDeviceSearch.
// It customizes the JSON unmarshaling process for ResponseDeviceSearch objects.
func (r *ResponseDeviceSearch) UnmarshalJSON(input []byte) error {
    var temp tempResponseDeviceSearch
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "end", "limit", "next", "results", "start", "total")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.End = *temp.End
    r.Limit = *temp.Limit
    r.Next = temp.Next
    r.Results = *temp.Results
    r.Start = *temp.Start
    r.Total = *temp.Total
    return nil
}

// tempResponseDeviceSearch is a temporary struct used for validating the fields of ResponseDeviceSearch.
type tempResponseDeviceSearch  struct {
    End     *int                                `json:"end"`
    Limit   *int                                `json:"limit"`
    Next    *string                             `json:"next,omitempty"`
    Results *[]ResponseDeviceSearchResultsItems `json:"results"`
    Start   *int                                `json:"start"`
    Total   *int                                `json:"total"`
}

func (r *tempResponseDeviceSearch) validate() error {
    var errs []string
    if r.End == nil {
        errs = append(errs, "required field `end` is missing for type `response_device_search`")
    }
    if r.Limit == nil {
        errs = append(errs, "required field `limit` is missing for type `response_device_search`")
    }
    if r.Results == nil {
        errs = append(errs, "required field `results` is missing for type `response_device_search`")
    }
    if r.Start == nil {
        errs = append(errs, "required field `start` is missing for type `response_device_search`")
    }
    if r.Total == nil {
        errs = append(errs, "required field `total` is missing for type `response_device_search`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
