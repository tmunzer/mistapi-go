package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ResponseDeviceMetricsResults represents a ResponseDeviceMetricsResults struct.
// This is Array of a container for one-of cases.
type ResponseDeviceMetricsResults struct {
    value    any
    isString bool
    isNumber bool
}

// String converts the ResponseDeviceMetricsResults object to a string representation.
func (r ResponseDeviceMetricsResults) String() string {
    if bytes, err := json.Marshal(r.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for ResponseDeviceMetricsResults.
// It customizes the JSON marshaling process for ResponseDeviceMetricsResults objects.
func (r ResponseDeviceMetricsResults) MarshalJSON() (
    []byte,
    error) {
    if r.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.ResponseDeviceMetricsResultsContainer.From*` functions to initialize the ResponseDeviceMetricsResults object.")
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseDeviceMetricsResults object to a map representation for JSON marshaling.
func (r *ResponseDeviceMetricsResults) toMap() any {
    switch obj := r.value.(type) {
    case *string:
        return *obj
    case *int:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseDeviceMetricsResults.
// It customizes the JSON unmarshaling process for ResponseDeviceMetricsResults objects.
func (r *ResponseDeviceMetricsResults) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(string), false, &r.isString),
        NewTypeHolder(new(int), false, &r.isNumber),
    )
    
    r.value = result
    return err
}

func (r *ResponseDeviceMetricsResults) AsString() (
    *string,
    bool) {
    if !r.isString {
        return nil, false
    }
    return r.value.(*string), true
}

func (r *ResponseDeviceMetricsResults) AsNumber() (
    *int,
    bool) {
    if !r.isNumber {
        return nil, false
    }
    return r.value.(*int), true
}

// internalResponseDeviceMetricsResults represents a responseDeviceMetricsResults struct.
// This is Array of a container for one-of cases.
type internalResponseDeviceMetricsResults struct {}

var ResponseDeviceMetricsResultsContainer internalResponseDeviceMetricsResults

// The internalResponseDeviceMetricsResults instance, wrapping the provided string value.
func (r *internalResponseDeviceMetricsResults) FromString(val string) ResponseDeviceMetricsResults {
    return ResponseDeviceMetricsResults{value: &val}
}

// The internalResponseDeviceMetricsResults instance, wrapping the provided int value.
func (r *internalResponseDeviceMetricsResults) FromNumber(val int) ResponseDeviceMetricsResults {
    return ResponseDeviceMetricsResults{value: &val}
}
