package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ResponseTunnelSearchResults represents a ResponseTunnelSearchResults struct.
// This is Array of a container for any-of cases.
type ResponseTunnelSearchResults struct {
    value            any
    isMxtunnelStats  bool
    isWanTunnelStats bool
}

// String converts the ResponseTunnelSearchResults object to a string representation.
func (r ResponseTunnelSearchResults) String() string {
    if bytes, err := json.Marshal(r.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for ResponseTunnelSearchResults.
// It customizes the JSON marshaling process for ResponseTunnelSearchResults objects.
func (r ResponseTunnelSearchResults) MarshalJSON() (
    []byte,
    error) {
    if r.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.ResponseTunnelSearchResultsContainer.From*` functions to initialize the ResponseTunnelSearchResults object.")
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseTunnelSearchResults object to a map representation for JSON marshaling.
func (r *ResponseTunnelSearchResults) toMap() any {
    switch obj := r.value.(type) {
    case *MxtunnelStats:
        return obj.toMap()
    case *WanTunnelStats:
        return obj.toMap()
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseTunnelSearchResults.
// It customizes the JSON unmarshaling process for ResponseTunnelSearchResults objects.
func (r *ResponseTunnelSearchResults) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallAnyOf(input,
        NewTypeHolder(&MxtunnelStats{}, false, &r.isMxtunnelStats),
        NewTypeHolder(&WanTunnelStats{}, false, &r.isWanTunnelStats),
    )
    
    r.value = result
    return err
}

func (r *ResponseTunnelSearchResults) AsMxtunnelStats() (
    *MxtunnelStats,
    bool) {
    if !r.isMxtunnelStats {
        return nil, false
    }
    return r.value.(*MxtunnelStats), true
}

func (r *ResponseTunnelSearchResults) AsWanTunnelStats() (
    *WanTunnelStats,
    bool) {
    if !r.isWanTunnelStats {
        return nil, false
    }
    return r.value.(*WanTunnelStats), true
}

// internalResponseTunnelSearchResults represents a responseTunnelSearchResults struct.
// This is Array of a container for any-of cases.
type internalResponseTunnelSearchResults struct {}

var ResponseTunnelSearchResultsContainer internalResponseTunnelSearchResults

// The internalResponseTunnelSearchResults instance, wrapping the provided MxtunnelStats value.
func (r *internalResponseTunnelSearchResults) FromMxtunnelStats(val MxtunnelStats) ResponseTunnelSearchResults {
    return ResponseTunnelSearchResults{value: &val}
}

// The internalResponseTunnelSearchResults instance, wrapping the provided WanTunnelStats value.
func (r *internalResponseTunnelSearchResults) FromWanTunnelStats(val WanTunnelStats) ResponseTunnelSearchResults {
    return ResponseTunnelSearchResults{value: &val}
}
