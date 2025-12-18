// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
)

// RoutingPolicyLocalPreference represents a RoutingPolicyLocalPreference struct.
// Optional, for an import policy, local_preference can be changed, value in range 1-4294967294. Can be a Variable (e.g. `{{bgp_as}}`)
type RoutingPolicyLocalPreference struct {
	value    any
	isString bool
	isNumber bool
}

// String implements the fmt.Stringer interface for RoutingPolicyLocalPreference,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r RoutingPolicyLocalPreference) String() string {
	return fmt.Sprintf("%v", r.value)
}

// MarshalJSON implements the json.Marshaler interface for RoutingPolicyLocalPreference.
// It customizes the JSON marshaling process for RoutingPolicyLocalPreference objects.
func (r RoutingPolicyLocalPreference) MarshalJSON() (
	[]byte,
	error) {
	if r.value == nil {
		return nil, errors.New("No underlying type is set. Please use any of the `models.RoutingPolicyLocalPreferenceContainer.From*` functions to initialize the RoutingPolicyLocalPreference object.")
	}
	return json.Marshal(r.toMap())
}

// toMap converts the RoutingPolicyLocalPreference object to a map representation for JSON marshaling.
func (r *RoutingPolicyLocalPreference) toMap() any {
	switch obj := r.value.(type) {
	case *string:
		return *obj
	case *int:
		return *obj
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for RoutingPolicyLocalPreference.
// It customizes the JSON unmarshaling process for RoutingPolicyLocalPreference objects.
func (r *RoutingPolicyLocalPreference) UnmarshalJSON(input []byte) error {
	result, err := UnmarshallAnyOf(input,
		NewTypeHolder(new(string), false, &r.isString),
		NewTypeHolder(new(int), false, &r.isNumber),
	)

	r.value = result
	return err
}

func (r *RoutingPolicyLocalPreference) AsString() (
	*string,
	bool) {
	if !r.isString {
		return nil, false
	}
	return r.value.(*string), true
}

func (r *RoutingPolicyLocalPreference) AsNumber() (
	*int,
	bool) {
	if !r.isNumber {
		return nil, false
	}
	return r.value.(*int), true
}

// internalRoutingPolicyLocalPreference represents a routingPolicyLocalPreference struct.
// Optional, for an import policy, local_preference can be changed, value in range 1-4294967294. Can be a Variable (e.g. `{{bgp_as}}`)
type internalRoutingPolicyLocalPreference struct{}

var RoutingPolicyLocalPreferenceContainer internalRoutingPolicyLocalPreference

// The internalRoutingPolicyLocalPreference instance, wrapping the provided string value.
func (r *internalRoutingPolicyLocalPreference) FromString(val string) RoutingPolicyLocalPreference {
	return RoutingPolicyLocalPreference{value: &val}
}

// The internalRoutingPolicyLocalPreference instance, wrapping the provided int value.
func (r *internalRoutingPolicyLocalPreference) FromNumber(val int) RoutingPolicyLocalPreference {
	return RoutingPolicyLocalPreference{value: &val}
}
