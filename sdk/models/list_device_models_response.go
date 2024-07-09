package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ListDeviceModelsResponse represents a ListDeviceModelsResponse struct.
// This is Array of a container for any-of cases.
type ListDeviceModelsResponse struct {
    value                any
    isConstDeviceAp      bool
    isConstDeviceSwitch  bool
    isConstDeviceGateway bool
    isConstDeviceUnknown bool
}

// String converts the ListDeviceModelsResponse object to a string representation.
func (l ListDeviceModelsResponse) String() string {
    if bytes, err := json.Marshal(l.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for ListDeviceModelsResponse.
// It customizes the JSON marshaling process for ListDeviceModelsResponse objects.
func (l ListDeviceModelsResponse) MarshalJSON() (
    []byte,
    error) {
    if l.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.ListDeviceModelsResponseContainer.From*` functions to initialize the ListDeviceModelsResponse object.")
    }
    return json.Marshal(l.toMap())
}

// toMap converts the ListDeviceModelsResponse object to a map representation for JSON marshaling.
func (l *ListDeviceModelsResponse) toMap() any {
    switch obj := l.value.(type) {
    case *ConstDeviceAp:
        return obj.toMap()
    case *ConstDeviceSwitch:
        return obj.toMap()
    case *ConstDeviceGateway:
        return obj.toMap()
    case *ConstDeviceUnknown:
        return obj.toMap()
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListDeviceModelsResponse.
// It customizes the JSON unmarshaling process for ListDeviceModelsResponse objects.
func (l *ListDeviceModelsResponse) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallAnyOf(input,
        NewTypeHolder(&ConstDeviceAp{}, false, &l.isConstDeviceAp),
        NewTypeHolder(&ConstDeviceSwitch{}, false, &l.isConstDeviceSwitch),
        NewTypeHolder(&ConstDeviceGateway{}, false, &l.isConstDeviceGateway),
        NewTypeHolder(&ConstDeviceUnknown{}, false, &l.isConstDeviceUnknown),
    )
    
    l.value = result
    return err
}

func (l *ListDeviceModelsResponse) AsConstDeviceAp() (
    *ConstDeviceAp,
    bool) {
    if !l.isConstDeviceAp {
        return nil, false
    }
    return l.value.(*ConstDeviceAp), true
}

func (l *ListDeviceModelsResponse) AsConstDeviceSwitch() (
    *ConstDeviceSwitch,
    bool) {
    if !l.isConstDeviceSwitch {
        return nil, false
    }
    return l.value.(*ConstDeviceSwitch), true
}

func (l *ListDeviceModelsResponse) AsConstDeviceGateway() (
    *ConstDeviceGateway,
    bool) {
    if !l.isConstDeviceGateway {
        return nil, false
    }
    return l.value.(*ConstDeviceGateway), true
}

func (l *ListDeviceModelsResponse) AsConstDeviceUnknown() (
    *ConstDeviceUnknown,
    bool) {
    if !l.isConstDeviceUnknown {
        return nil, false
    }
    return l.value.(*ConstDeviceUnknown), true
}

// internalListDeviceModelsResponse represents a listDeviceModelsResponse struct.
// This is Array of a container for any-of cases.
type internalListDeviceModelsResponse struct {}

var ListDeviceModelsResponseContainer internalListDeviceModelsResponse

// The internalListDeviceModelsResponse instance, wrapping the provided ConstDeviceAp value.
func (l *internalListDeviceModelsResponse) FromConstDeviceAp(val ConstDeviceAp) ListDeviceModelsResponse {
    return ListDeviceModelsResponse{value: &val}
}

// The internalListDeviceModelsResponse instance, wrapping the provided ConstDeviceSwitch value.
func (l *internalListDeviceModelsResponse) FromConstDeviceSwitch(val ConstDeviceSwitch) ListDeviceModelsResponse {
    return ListDeviceModelsResponse{value: &val}
}

// The internalListDeviceModelsResponse instance, wrapping the provided ConstDeviceGateway value.
func (l *internalListDeviceModelsResponse) FromConstDeviceGateway(val ConstDeviceGateway) ListDeviceModelsResponse {
    return ListDeviceModelsResponse{value: &val}
}

// The internalListDeviceModelsResponse instance, wrapping the provided ConstDeviceUnknown value.
func (l *internalListDeviceModelsResponse) FromConstDeviceUnknown(val ConstDeviceUnknown) ListDeviceModelsResponse {
    return ListDeviceModelsResponse{value: &val}
}
