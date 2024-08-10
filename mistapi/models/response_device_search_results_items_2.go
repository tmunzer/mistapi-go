package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ResponseDeviceSearchResultsItems2 represents a ResponseDeviceSearchResultsItems2 struct.
type ResponseDeviceSearchResultsItems2 struct {
    value           any
    isApSearch      bool
    isSwitchSearch  bool
    isGatewaySearch bool
}

// String converts the ResponseDeviceSearchResultsItems2 object to a string representation.
func (r ResponseDeviceSearchResultsItems2) String() string {
    if bytes, err := json.Marshal(r.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for ResponseDeviceSearchResultsItems2.
// It customizes the JSON marshaling process for ResponseDeviceSearchResultsItems2 objects.
func (r ResponseDeviceSearchResultsItems2) MarshalJSON() (
    []byte,
    error) {
    if r.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.ResponseDeviceSearchResultsItemsContainer.From*` functions to initialize the ResponseDeviceSearchResultsItems2 object.")
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseDeviceSearchResultsItems2 object to a map representation for JSON marshaling.
func (r *ResponseDeviceSearchResultsItems2) toMap() any {
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

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseDeviceSearchResultsItems2.
// It customizes the JSON unmarshaling process for ResponseDeviceSearchResultsItems2 objects.
func (r *ResponseDeviceSearchResultsItems2) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallAnyOf(input,
        NewTypeHolder(&ApSearch{}, false, &r.isApSearch),
        NewTypeHolder(&SwitchSearch{}, false, &r.isSwitchSearch),
        NewTypeHolder(&GatewaySearch{}, false, &r.isGatewaySearch),
    )
    
    r.value = result
    return err
}

func (r *ResponseDeviceSearchResultsItems2) AsApSearch() (
    *ApSearch,
    bool) {
    if !r.isApSearch {
        return nil, false
    }
    return r.value.(*ApSearch), true
}

func (r *ResponseDeviceSearchResultsItems2) AsSwitchSearch() (
    *SwitchSearch,
    bool) {
    if !r.isSwitchSearch {
        return nil, false
    }
    return r.value.(*SwitchSearch), true
}

func (r *ResponseDeviceSearchResultsItems2) AsGatewaySearch() (
    *GatewaySearch,
    bool) {
    if !r.isGatewaySearch {
        return nil, false
    }
    return r.value.(*GatewaySearch), true
}

// internalResponseDeviceSearchResultsItems2 represents a responseDeviceSearchResultsItems2 struct.
type internalResponseDeviceSearchResultsItems2 struct {}

var ResponseDeviceSearchResultsItemsContainer internalResponseDeviceSearchResultsItems2

// The internalResponseDeviceSearchResultsItems2 instance, wrapping the provided ApSearch value.
func (r *internalResponseDeviceSearchResultsItems2) FromApSearch(val ApSearch) ResponseDeviceSearchResultsItems2 {
    return ResponseDeviceSearchResultsItems2{value: &val}
}

// The internalResponseDeviceSearchResultsItems2 instance, wrapping the provided SwitchSearch value.
func (r *internalResponseDeviceSearchResultsItems2) FromSwitchSearch(val SwitchSearch) ResponseDeviceSearchResultsItems2 {
    return ResponseDeviceSearchResultsItems2{value: &val}
}

// The internalResponseDeviceSearchResultsItems2 instance, wrapping the provided GatewaySearch value.
func (r *internalResponseDeviceSearchResultsItems2) FromGatewaySearch(val GatewaySearch) ResponseDeviceSearchResultsItems2 {
    return ResponseDeviceSearchResultsItems2{value: &val}
}
