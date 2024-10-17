package models

import (
	"encoding/json"
	"errors"
	"strings"
)

// ResponseTunnelSearchItem represents a ResponseTunnelSearchItem struct.
type ResponseTunnelSearchItem struct {
	value            any
	isStatsMxtunnel  bool
	isStatsWanTunnel bool
}

// String converts the ResponseTunnelSearchItem object to a string representation.
func (r ResponseTunnelSearchItem) String() string {
	if bytes, err := json.Marshal(r.value); err == nil {
		return strings.Trim(string(bytes), "\"")
	}
	return ""
}

// MarshalJSON implements the json.Marshaler interface for ResponseTunnelSearchItem.
// It customizes the JSON marshaling process for ResponseTunnelSearchItem objects.
func (r ResponseTunnelSearchItem) MarshalJSON() (
	[]byte,
	error) {
	if r.value == nil {
		return nil, errors.New("No underlying type is set. Please use any of the `models.ResponseTunnelSearchItemContainer.From*` functions to initialize the ResponseTunnelSearchItem object.")
	}
	return json.Marshal(r.toMap())
}

// toMap converts the ResponseTunnelSearchItem object to a map representation for JSON marshaling.
func (r *ResponseTunnelSearchItem) toMap() any {
	switch obj := r.value.(type) {
	case *StatsMxtunnel:
		return obj.toMap()
	case *StatsWanTunnel:
		return obj.toMap()
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseTunnelSearchItem.
// It customizes the JSON unmarshaling process for ResponseTunnelSearchItem objects.
func (r *ResponseTunnelSearchItem) UnmarshalJSON(input []byte) error {
	result, err := UnmarshallOneOf(input,
		NewTypeHolder(&StatsMxtunnel{}, false, &r.isStatsMxtunnel),
		NewTypeHolder(&StatsWanTunnel{}, false, &r.isStatsWanTunnel),
	)

	r.value = result
	return err
}

func (r *ResponseTunnelSearchItem) AsStatsMxtunnel() (
	*StatsMxtunnel,
	bool) {
	if !r.isStatsMxtunnel {
		return nil, false
	}
	return r.value.(*StatsMxtunnel), true
}

func (r *ResponseTunnelSearchItem) AsStatsWanTunnel() (
	*StatsWanTunnel,
	bool) {
	if !r.isStatsWanTunnel {
		return nil, false
	}
	return r.value.(*StatsWanTunnel), true
}

// internalResponseTunnelSearchItem represents a responseTunnelSearchItem struct.
type internalResponseTunnelSearchItem struct{}

var ResponseTunnelSearchItemContainer internalResponseTunnelSearchItem

// The internalResponseTunnelSearchItem instance, wrapping the provided StatsMxtunnel value.
func (r *internalResponseTunnelSearchItem) FromStatsMxtunnel(val StatsMxtunnel) ResponseTunnelSearchItem {
	return ResponseTunnelSearchItem{value: &val}
}

// The internalResponseTunnelSearchItem instance, wrapping the provided StatsWanTunnel value.
func (r *internalResponseTunnelSearchItem) FromStatsWanTunnel(val StatsWanTunnel) ResponseTunnelSearchItem {
	return ResponseTunnelSearchItem{value: &val}
}
