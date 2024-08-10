package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ResponseDeviceMetricsResultsItems represents a ResponseDeviceMetricsResultsItems struct.
type ResponseDeviceMetricsResultsItems struct {
    value    any
    isString bool
    isNumber bool
}

// String converts the ResponseDeviceMetricsResultsItems object to a string representation.
func (r ResponseDeviceMetricsResultsItems) String() string {
    if bytes, err := json.Marshal(r.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for ResponseDeviceMetricsResultsItems.
// It customizes the JSON marshaling process for ResponseDeviceMetricsResultsItems objects.
func (r ResponseDeviceMetricsResultsItems) MarshalJSON() (
    []byte,
    error) {
    if r.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.ResponseDeviceMetricsResultsItemsContainer.From*` functions to initialize the ResponseDeviceMetricsResultsItems object.")
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseDeviceMetricsResultsItems object to a map representation for JSON marshaling.
func (r *ResponseDeviceMetricsResultsItems) toMap() any {
    switch obj := r.value.(type) {
    case *string:
        return *obj
    case *int:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseDeviceMetricsResultsItems.
// It customizes the JSON unmarshaling process for ResponseDeviceMetricsResultsItems objects.
func (r *ResponseDeviceMetricsResultsItems) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(string), false, &r.isString),
        NewTypeHolder(new(int), false, &r.isNumber),
    )
    
    r.value = result
    return err
}

func (r *ResponseDeviceMetricsResultsItems) AsString() (
    *string,
    bool) {
    if !r.isString {
        return nil, false
    }
    return r.value.(*string), true
}

func (r *ResponseDeviceMetricsResultsItems) AsNumber() (
    *int,
    bool) {
    if !r.isNumber {
        return nil, false
    }
    return r.value.(*int), true
}

// internalResponseDeviceMetricsResultsItems represents a responseDeviceMetricsResultsItems struct.
type internalResponseDeviceMetricsResultsItems struct {}

var ResponseDeviceMetricsResultsItemsContainer internalResponseDeviceMetricsResultsItems

// The internalResponseDeviceMetricsResultsItems instance, wrapping the provided string value.
func (r *internalResponseDeviceMetricsResultsItems) FromString(val string) ResponseDeviceMetricsResultsItems {
    return ResponseDeviceMetricsResultsItems{value: &val}
}

// The internalResponseDeviceMetricsResultsItems instance, wrapping the provided int value.
func (r *internalResponseDeviceMetricsResultsItems) FromNumber(val int) ResponseDeviceMetricsResultsItems {
    return ResponseDeviceMetricsResultsItems{value: &val}
}
