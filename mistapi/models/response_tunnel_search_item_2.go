package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ResponseTunnelSearchItem2 represents a ResponseTunnelSearchItem2 struct.
type ResponseTunnelSearchItem2 struct {
    value            any
    isMxtunnelStats  bool
    isWanTunnelStats bool
}

// String converts the ResponseTunnelSearchItem2 object to a string representation.
func (r ResponseTunnelSearchItem2) String() string {
    if bytes, err := json.Marshal(r.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for ResponseTunnelSearchItem2.
// It customizes the JSON marshaling process for ResponseTunnelSearchItem2 objects.
func (r ResponseTunnelSearchItem2) MarshalJSON() (
    []byte,
    error) {
    if r.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.ResponseTunnelSearchItemContainer.From*` functions to initialize the ResponseTunnelSearchItem2 object.")
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseTunnelSearchItem2 object to a map representation for JSON marshaling.
func (r *ResponseTunnelSearchItem2) toMap() any {
    switch obj := r.value.(type) {
    case *MxtunnelStats:
        return obj.toMap()
    case *WanTunnelStats:
        return obj.toMap()
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseTunnelSearchItem2.
// It customizes the JSON unmarshaling process for ResponseTunnelSearchItem2 objects.
func (r *ResponseTunnelSearchItem2) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallAnyOf(input,
        NewTypeHolder(&MxtunnelStats{}, false, &r.isMxtunnelStats),
        NewTypeHolder(&WanTunnelStats{}, false, &r.isWanTunnelStats),
    )
    
    r.value = result
    return err
}

func (r *ResponseTunnelSearchItem2) AsMxtunnelStats() (
    *MxtunnelStats,
    bool) {
    if !r.isMxtunnelStats {
        return nil, false
    }
    return r.value.(*MxtunnelStats), true
}

func (r *ResponseTunnelSearchItem2) AsWanTunnelStats() (
    *WanTunnelStats,
    bool) {
    if !r.isWanTunnelStats {
        return nil, false
    }
    return r.value.(*WanTunnelStats), true
}

// internalResponseTunnelSearchItem2 represents a responseTunnelSearchItem2 struct.
type internalResponseTunnelSearchItem2 struct {}

var ResponseTunnelSearchItemContainer internalResponseTunnelSearchItem2

// The internalResponseTunnelSearchItem2 instance, wrapping the provided MxtunnelStats value.
func (r *internalResponseTunnelSearchItem2) FromMxtunnelStats(val MxtunnelStats) ResponseTunnelSearchItem2 {
    return ResponseTunnelSearchItem2{value: &val}
}

// The internalResponseTunnelSearchItem2 instance, wrapping the provided WanTunnelStats value.
func (r *internalResponseTunnelSearchItem2) FromWanTunnelStats(val WanTunnelStats) ResponseTunnelSearchItem2 {
    return ResponseTunnelSearchItem2{value: &val}
}
