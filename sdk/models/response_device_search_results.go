package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ResponseDeviceSearchResults represents a ResponseDeviceSearchResults struct.
// This is Array of a container for any-of cases.
type ResponseDeviceSearchResults struct {
    value           any
    isApSearch      bool
    isSwitchSearch  bool
    isGatewaySearch bool
}

// String converts the ResponseDeviceSearchResults object to a string representation.
func (r ResponseDeviceSearchResults) String() string {
    if bytes, err := json.Marshal(r.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for ResponseDeviceSearchResults.
// It customizes the JSON marshaling process for ResponseDeviceSearchResults objects.
func (r ResponseDeviceSearchResults) MarshalJSON() (
    []byte,
    error) {
    if r.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.ResponseDeviceSearchResultsContainer.From*` functions to initialize the ResponseDeviceSearchResults object.")
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseDeviceSearchResults object to a map representation for JSON marshaling.
func (r *ResponseDeviceSearchResults) toMap() any {
    switch obj := r.value.(type) {
    case *ApSearch:
        return obj.toMap()
    case *SwitchSearch:
        return obj.toMap()
    case *GatewaySearch:
        return obj.toMap()
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseDeviceSearchResults.
// It customizes the JSON unmarshaling process for ResponseDeviceSearchResults objects.
func (r *ResponseDeviceSearchResults) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallAnyOf(input,
        NewTypeHolder(&ApSearch{}, false, &r.isApSearch),
        NewTypeHolder(&SwitchSearch{}, false, &r.isSwitchSearch),
        NewTypeHolder(&GatewaySearch{}, false, &r.isGatewaySearch),
    )
    
    r.value = result
    return err
}

func (r *ResponseDeviceSearchResults) AsApSearch() (
    *ApSearch,
    bool) {
    if !r.isApSearch {
        return nil, false
    }
    return r.value.(*ApSearch), true
}

func (r *ResponseDeviceSearchResults) AsSwitchSearch() (
    *SwitchSearch,
    bool) {
    if !r.isSwitchSearch {
        return nil, false
    }
    return r.value.(*SwitchSearch), true
}

func (r *ResponseDeviceSearchResults) AsGatewaySearch() (
    *GatewaySearch,
    bool) {
    if !r.isGatewaySearch {
        return nil, false
    }
    return r.value.(*GatewaySearch), true
}

// internalResponseDeviceSearchResults represents a responseDeviceSearchResults struct.
// This is Array of a container for any-of cases.
type internalResponseDeviceSearchResults struct {}

var ResponseDeviceSearchResultsContainer internalResponseDeviceSearchResults

// The internalResponseDeviceSearchResults instance, wrapping the provided ApSearch value.
func (r *internalResponseDeviceSearchResults) FromApSearch(val ApSearch) ResponseDeviceSearchResults {
    return ResponseDeviceSearchResults{value: &val}
}

// The internalResponseDeviceSearchResults instance, wrapping the provided SwitchSearch value.
func (r *internalResponseDeviceSearchResults) FromSwitchSearch(val SwitchSearch) ResponseDeviceSearchResults {
    return ResponseDeviceSearchResults{value: &val}
}

// The internalResponseDeviceSearchResults instance, wrapping the provided GatewaySearch value.
func (r *internalResponseDeviceSearchResults) FromGatewaySearch(val GatewaySearch) ResponseDeviceSearchResults {
    return ResponseDeviceSearchResults{value: &val}
}
