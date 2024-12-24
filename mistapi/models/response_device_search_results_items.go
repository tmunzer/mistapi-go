package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// ResponseDeviceSearchResultsItems represents a ResponseDeviceSearchResultsItems struct.
type ResponseDeviceSearchResultsItems struct {
    value           any
    isApSearch      bool
    isSwitchSearch  bool
    isGatewaySearch bool
}

// String implements the fmt.Stringer interface for ResponseDeviceSearchResultsItems,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseDeviceSearchResultsItems) String() string {
    return fmt.Sprintf("%v", r.value)
}

// MarshalJSON implements the json.Marshaler interface for ResponseDeviceSearchResultsItems.
// It customizes the JSON marshaling process for ResponseDeviceSearchResultsItems objects.
func (r ResponseDeviceSearchResultsItems) MarshalJSON() (
    []byte,
    error) {
    if r.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.ResponseDeviceSearchResultsItemsContainer.From*` functions to initialize the ResponseDeviceSearchResultsItems object.")
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseDeviceSearchResultsItems object to a map representation for JSON marshaling.
func (r *ResponseDeviceSearchResultsItems) toMap() any {
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

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseDeviceSearchResultsItems.
// It customizes the JSON unmarshaling process for ResponseDeviceSearchResultsItems objects.
func (r *ResponseDeviceSearchResultsItems) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(&ApSearch{}, false, &r.isApSearch),
        NewTypeHolder(&SwitchSearch{}, false, &r.isSwitchSearch),
        NewTypeHolder(&GatewaySearch{}, false, &r.isGatewaySearch),
    )
    
    r.value = result
    return err
}

func (r *ResponseDeviceSearchResultsItems) AsApSearch() (
    *ApSearch,
    bool) {
    if !r.isApSearch {
        return nil, false
    }
    return r.value.(*ApSearch), true
}

func (r *ResponseDeviceSearchResultsItems) AsSwitchSearch() (
    *SwitchSearch,
    bool) {
    if !r.isSwitchSearch {
        return nil, false
    }
    return r.value.(*SwitchSearch), true
}

func (r *ResponseDeviceSearchResultsItems) AsGatewaySearch() (
    *GatewaySearch,
    bool) {
    if !r.isGatewaySearch {
        return nil, false
    }
    return r.value.(*GatewaySearch), true
}

// internalResponseDeviceSearchResultsItems represents a responseDeviceSearchResultsItems struct.
type internalResponseDeviceSearchResultsItems struct {}

var ResponseDeviceSearchResultsItemsContainer internalResponseDeviceSearchResultsItems

// The internalResponseDeviceSearchResultsItems instance, wrapping the provided ApSearch value.
func (r *internalResponseDeviceSearchResultsItems) FromApSearch(val ApSearch) ResponseDeviceSearchResultsItems {
    return ResponseDeviceSearchResultsItems{value: &val}
}

// The internalResponseDeviceSearchResultsItems instance, wrapping the provided SwitchSearch value.
func (r *internalResponseDeviceSearchResultsItems) FromSwitchSearch(val SwitchSearch) ResponseDeviceSearchResultsItems {
    return ResponseDeviceSearchResultsItems{value: &val}
}

// The internalResponseDeviceSearchResultsItems instance, wrapping the provided GatewaySearch value.
func (r *internalResponseDeviceSearchResultsItems) FromGatewaySearch(val GatewaySearch) ResponseDeviceSearchResultsItems {
    return ResponseDeviceSearchResultsItems{value: &val}
}
