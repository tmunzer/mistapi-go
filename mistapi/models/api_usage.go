package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// ApiUsage represents a ApiUsage struct.
type ApiUsage struct {
    // max number of request permitted
    RequestLimit         int                    `json:"request_limit"`
    // num of request made in the current hour
    Requests             int                    `json:"requests"`
    Seconds              *float64               `json:"seconds,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ApiUsage,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a ApiUsage) String() string {
    return fmt.Sprintf(
    	"ApiUsage[RequestLimit=%v, Requests=%v, Seconds=%v, AdditionalProperties=%v]",
    	a.RequestLimit, a.Requests, a.Seconds, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ApiUsage.
// It customizes the JSON marshaling process for ApiUsage objects.
func (a ApiUsage) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(a.AdditionalProperties,
        "request_limit", "requests", "seconds"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(a.toMap())
}

// toMap converts the ApiUsage object to a map representation for JSON marshaling.
func (a ApiUsage) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, a.AdditionalProperties)
    structMap["request_limit"] = a.RequestLimit
    structMap["requests"] = a.Requests
    if a.Seconds != nil {
        structMap["seconds"] = a.Seconds
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ApiUsage.
// It customizes the JSON unmarshaling process for ApiUsage objects.
func (a *ApiUsage) UnmarshalJSON(input []byte) error {
    var temp tempApiUsage
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "request_limit", "requests", "seconds")
    if err != nil {
    	return err
    }
    a.AdditionalProperties = additionalProperties
    
    a.RequestLimit = *temp.RequestLimit
    a.Requests = *temp.Requests
    a.Seconds = temp.Seconds
    return nil
}

// tempApiUsage is a temporary struct used for validating the fields of ApiUsage.
type tempApiUsage  struct {
    RequestLimit *int     `json:"request_limit"`
    Requests     *int     `json:"requests"`
    Seconds      *float64 `json:"seconds,omitempty"`
}

func (a *tempApiUsage) validate() error {
    var errs []string
    if a.RequestLimit == nil {
        errs = append(errs, "required field `request_limit` is missing for type `api_usage`")
    }
    if a.Requests == nil {
        errs = append(errs, "required field `requests` is missing for type `api_usage`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
